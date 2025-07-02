package game

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/byterotom/infinity-play/internal/db/dbgen"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
	"github.com/jackc/pgx/v5"
)

func (mux *GameMux) uploadGame(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	r.ParseMultipartForm(10 << 20)

	var arg dbgen.AddGameParams
	var err error

	tx, err := mux.conn.BeginTx(ctx, pgx.TxOptions{})
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
			tx.Rollback(ctx)
			http.Error(w, err.Error(), 500)
			return
		}
		tx.Commit(ctx)
	}()

	arg.Name = strings.ToLower(r.FormValue("name"))
	arg.Description = r.FormValue("description")
	t := r.FormValue("technology")

	if t == "html" {
		arg.Technology = dbgen.TechHtml
	} else {
		arg.Technology = dbgen.TechFlash
	}
	tags := strings.Split(strings.ReplaceAll(r.FormValue("tags"), " ", ""), ",")

	var buf bytes.Buffer
	file, _, err := r.FormFile("game_file")
	if err != nil {
		return
	}

	tee := io.TeeReader(file, &buf)
	arg.ID, err = pkg.HashWithReader(tee)
	if err != nil {
		return
	}

	err = q.AddGame(ctx, arg)
	if err != nil {
		return
	}

	for _, tag := range tags {
		err = q.AddNewTags(ctx, tag)
		if err != nil {
			return
		}
		tid, err := q.GetTagIdByName(ctx, tag)
		if err != nil {
			return
		}
		err = q.AddGameTags(ctx, dbgen.AddGameTagsParams{GameID: arg.ID, TagID: tid})
		if err != nil {
			return
		}
	}

	// game_file upload
	var gameFileKey string
	if t == "html" {
		gameFileKey = fmt.Sprintf("%s/game_file.zip", arg.ID)
	} else {
		gameFileKey = fmt.Sprintf("%s/game_file.swf", arg.ID)
	}
	err = mux.r2.Upload(gameFileKey, &buf)
	if err != nil {
		return
	}

	// thumbnail upload
	thumbnail, _, err := r.FormFile("thumbnail")
	thumbnailKey := fmt.Sprintf("%s/thumbnail", arg.ID)
	if err != nil {
		return
	}
	err = mux.r2.Upload(thumbnailKey, thumbnail)
	if err != nil {
		return
	}

	// gif upload
	gif, _, err := r.FormFile("gif")
	gifKey := fmt.Sprintf("%s/gif.gif", arg.ID)
	if err != nil {
		return
	}
	err = mux.r2.Upload(gifKey, gif)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "Uploaded %s", arg.Name)
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
	if pkg.IsHTMXRequest(r) {
		components.Game(&game).Render(r.Context(), w)
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
	if fileType != "swf" && fileType != "thumbnail" && fileType != "gif" && fileType != "html" {
		return
	}

	var fileKey string

	switch fileType {
	case "swf":
		fileKey = fmt.Sprintf("%s/game_file.swf", gameId)
		w.Header().Set("Content-Type", "application/x-shockwave-flash")
	case "html":
		fileKey = fmt.Sprintf("%s/game_file.zip", gameId)
		w.Header().Set("Content-Type", "application/zip")
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

	err = q.DeleteGameById(ctx, id)
	if err != nil {
		return
	}

	mux.r2.Delete(id)
}

func (mux *GameMux) vote(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	gameId := r.PathValue("game_id")
	voteType := r.URL.Query().Get("v")

	q := dbgen.New(mux.conn)

	err = q.VoteGameById(ctx, gameId)
	if err != nil {
		return
	}

	if voteType == "like" {
		err = q.LikeGameById(ctx, gameId)
		if err != nil {
			return
		}
	}

}
