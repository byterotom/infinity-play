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
	mux.HandleFunc("POST /upload", authenticate(mux.uploadGame))
	mux.HandleFunc("DELETE /{game_name}", authenticate(mux.deleteGame))
	mux.HandleFunc("GET /{game_name}", mux.getGameData)
	mux.HandleFunc("GET /{file_type}/{game_id}", mux.getGameFile)
	mux.HandleFunc("GET /search", mux.searchGame)

	return http.StripPrefix("/game", mux)
}
