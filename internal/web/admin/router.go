package admin

import (
	"database/sql"
	"net/http"

	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
)

type AdminMux struct {
	*http.ServeMux
	conn *sql.DB
}

func NewAdminMux(r2 *pkg.R2, conn *sql.DB) http.Handler {
	mux := AdminMux{
		ServeMux: http.NewServeMux(),
		conn:     conn,
	}

	// routes
	mux.HandleFunc("GET /{act}", func(w http.ResponseWriter, r *http.Request) {
		act := r.PathValue("act")
		views.Index(components.Admin(act)).Render(r.Context(), w)
	})

	return http.StripPrefix("/admin", mux)
}
