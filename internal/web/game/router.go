package game

import (
	"database/sql"
	"net/http"

	"github.com/byterotom/infinity-play/pkg"
)

type GameMux struct {
	*http.ServeMux
	r2   *pkg.R2
	conn *sql.DB
}

func NewGameMux(r2 *pkg.R2, conn *sql.DB) http.Handler {
	mux := GameMux{
		ServeMux: http.NewServeMux(),
		r2:       r2,
		conn:     conn,
	}

	// routes
	mux.HandleFunc("POST /upload", mux.uploadGame)
	mux.HandleFunc("GET /{game_name}", mux.getGame)
	mux.HandleFunc("DELETE /{game_name}", mux.deleteGame)

	return http.StripPrefix("/game", mux)
}
