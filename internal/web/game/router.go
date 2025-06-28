package game

import (
	"net/http"

	"github.com/byterotom/infinity-play/pkg"
	"github.com/jackc/pgx/v5"
)

type GameMux struct {
	*http.ServeMux
	r2   *pkg.R2
	conn *pgx.Conn
}

func NewGameMux(r2 *pkg.R2, conn *pgx.Conn) http.Handler {
	mux := GameMux{
		ServeMux: http.NewServeMux(),
		r2:       r2,
		conn:     conn,
	}

	// routes
	mux.HandleFunc("POST /upload", mux.uploadGame)
	mux.HandleFunc("GET /{game_name}", mux.getGameData)
	mux.HandleFunc("DELETE /{game_name}", mux.deleteGame)
	mux.HandleFunc("GET /search", mux.searchGame)
	mux.HandleFunc("GET /{file_type}/{game_id}", mux.getGameFile)

	return http.StripPrefix("/game", mux)
}
