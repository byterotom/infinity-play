package game

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/byterotom/infinity-play/internal/db/dbgen"
	"github.com/byterotom/infinity-play/pkg"
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
		file, _, _ := r.FormFile("game_file")

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
}

func (mux *GameMux) getGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GAME PAGE!")
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

	id, err := q.GetIdByName(ctx, gameName)
	if err != nil {
		return
	}

	_, err = q.DeleteById(ctx, id)
	if err != nil {
		return
	}

	mux.r2.Delete(id)
}
