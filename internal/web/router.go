package web

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/byterotom/infinity-play/internal/web/game"
	"github.com/byterotom/infinity-play/pkg"
)

type InfinityMux struct {
	*http.ServeMux
	tmpl *template.Template
}

func NewInfinityMux(content *embed.FS, r2 *pkg.R2, conn *sql.DB) http.Handler {
	mux := &InfinityMux{
		ServeMux: http.NewServeMux(),
		tmpl:     template.Must(template.ParseFS(content, "templates/index.html", "templates/game.html", "templates/partials/navbar.html")),
	}

	mux.Handle("/game/", game.NewGameMux(r2, conn, mux.tmpl))
	mux.setupLayout()

	return mux
}

func (mux *InfinityMux) setupLayout() {

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err := mux.tmpl.ExecuteTemplate(w, "index", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Template execution error:", err)
		}
	})
}
