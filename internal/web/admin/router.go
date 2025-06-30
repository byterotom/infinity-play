package admin

import (
	"net/http"

	"github.com/byterotom/infinity-play/pkg"
	"github.com/jackc/pgx/v5"
)

type AdminMux struct {
	*http.ServeMux
	conn *pgx.Conn
}

func NewAdminMux(r2 *pkg.R2, conn *pgx.Conn) http.Handler {
	mux := AdminMux{
		ServeMux: http.NewServeMux(),
		conn:     conn,
	}

	// routes
	mux.HandleFunc("GET /", loggedIn(formHandler))
	mux.HandleFunc("GET /{act}", loggedIn(formHandler))
	mux.HandleFunc("POST /login", mux.login)
	mux.HandleFunc("GET /logout", mux.logout)

	return http.StripPrefix("/admin", mux)
}
