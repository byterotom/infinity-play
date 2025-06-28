package game

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/byterotom/infinity-play/internal/db/dbgen"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
)

func (mux *GameMux) uploadGame(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	r.ParseMultipartForm(10 << 20)

	var arg dbgen.AddGameParams
	var err error

	tx, err := mux.conn.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	q := dbgen.New(mux.conn).WithTx(tx)

	defer func() {
		if err != nil {
			if arg.ID != "" {
				mux.r2.Delete(arg.ID)
			}
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}
		tx.Commit()
	}()
	arg.Name = r.FormValue("name")
	arg.Description = r.FormValue("description")
	arg.Technology = r.FormValue("technology")

	var buf bytes.Buffer
	if arg.Technology == "flash" {
		file, _, err := r.FormFile("game_file")
		if err != nil {
			return
		}
		tee := io.TeeReader(file, &buf)

		arg.ID, err = pkg.HashWithReader(tee)
		if err != nil {
			return
		}
	} else {
		arg.GameUrl = sql.NullString{
			String: r.FormValue("game_url"),
			Valid:  true,
		}

		arg.ID = pkg.HashWithString(arg.GameUrl.String)
	}

	_, err = q.AddGame(ctx, arg)
	if err != nil {
		return
	}

	if arg.Technology == "flash" {
		gameFileKey := fmt.Sprintf("%s/game_file.swf", arg.ID)

		err = mux.r2.Upload(gameFileKey, &buf)
		if err != nil {
			return
		}
	}

	thumbnail, _, err := r.FormFile("thumbnail")
	thumbnailKey := fmt.Sprintf("%s/thumbnail", arg.ID)
	if err != nil {
		return
	}

	err = mux.r2.Upload(thumbnailKey, thumbnail)
	if err != nil {
		return
	}

	gif, _, err := r.FormFile("gif")
	gifKey := fmt.Sprintf("%s/gif.gif", arg.ID)
	if err != nil {
		return
	}

	err = mux.r2.Upload(gifKey, gif)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "Uploaded")
}

func (mux *GameMux) getGameData(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	gameName := r.PathValue("game_name")

	q := dbgen.New(mux.conn)

	game, err := q.GetGameByName(ctx, gameName)
	if err != nil {
		return
	}
	views.Index(components.Game(&game)).Render(r.Context(), w)
}

func (mux *GameMux) getGameFile(w http.ResponseWriter, r *http.Request) {

	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	gameId := r.PathValue("game_id")
	fileType := r.PathValue("file_type")
	if fileType != "swf" && fileType != "thumbnail" && fileType != "gif" {
		return
	}

	var fileKey string

	switch fileType {
	case "swf":
		fileKey = fmt.Sprintf("%s/game_file.swf", gameId)
		w.Header().Set("Content-Type", "application/x-shockwave-flash")
	case "gif":
		fileKey = fmt.Sprintf("%s/gif.gif", gameId)
		w.Header().Set("Content-Type", "image/gif")
	default:
		fileKey = fmt.Sprintf("%s/thumbnail", gameId)
		w.Header().Set("Content-Type", "image/jpeg")
	}

	obj, err := mux.r2.Get(fileKey)
	if err != nil {
		return
	}
	defer obj.Close()

	_, err = io.Copy(w, obj)

	if err != nil {
		log.Println("Stream error:", err)
		return
	}
}

func (mux *GameMux) deleteGame(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	gameName := r.PathValue("game_name")

	q := dbgen.New(mux.conn)

	id, err := q.GetGameIdByName(ctx, gameName)
	if err != nil {
		return
	}

	_, err = q.DeleteGameById(ctx, id)
	if err != nil {
		return
	}

	mux.r2.Delete(id)
}
