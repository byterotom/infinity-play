package web

import (
	"embed"
	"net/http"

	"github.com/byterotom/infinity-play/internal/web/admin"
	"github.com/byterotom/infinity-play/internal/web/game"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/jackc/pgx/v5"
)

type InfinityMux struct {
	*http.ServeMux
	conn *pgx.Conn
}

func NewInfinityMux(r2 *pkg.R2, conn *pgx.Conn, staticFs *embed.FS) http.Handler {
	mux := &InfinityMux{
		ServeMux: http.NewServeMux(),
		conn:     conn,
	}

	mux.HandleFunc("/", mux.home)
	mux.HandleFunc("GET /category/{cat}", mux.category)
	mux.Handle("/game/", game.NewGameMux(r2, conn))
	mux.Handle("/admin/", admin.NewAdminMux(r2, conn))

	// satic file server
	fs := http.FileServer(http.FS(staticFs))
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Service-Worker-Allowed", "/")
		fs.ServeHTTP(w, r)
	})

	return mux
}
