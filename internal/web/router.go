package web

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/byterotom/infinity-play/internal/web/game"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
)

type InfinityMux struct {
	*http.ServeMux
}

func NewInfinityMux(r2 *pkg.R2, conn *sql.DB) http.Handler {
	mux := &InfinityMux{
		ServeMux: http.NewServeMux(),
	}

	mux.Handle("/game/", game.NewGameMux(r2, conn))
	mux.setupLayout()

	return mux
}

func (mux *InfinityMux) setupLayout() {

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err := views.Index(nil).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Template execution error:", err)
		}
	})
}
