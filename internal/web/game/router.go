package game

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/byterotom/infinity-play/pkg"
)

type GameMux struct {
	*http.ServeMux
	r2   *pkg.R2
	conn *sql.DB
	tmpl *template.Template
}

func NewGameMux(r2 *pkg.R2, conn *sql.DB, tmpl *template.Template) http.Handler {
	mux := GameMux{
		ServeMux: http.NewServeMux(),
		r2:       r2,
		conn:     conn,
		tmpl:     tmpl,
	}

	// routes
	mux.HandleFunc("POST /upload", mux.uploadGame)
	mux.HandleFunc("GET /{game_name}", mux.getGameData)
	mux.HandleFunc("GET /file/{game_id}", mux.getGameFile)
	mux.HandleFunc("DELETE /{game_name}", mux.deleteGame)

	return http.StripPrefix("/game", mux)
}
