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
		tmpl:     template.Must(template.ParseFS(content, "templates/index.html", "templates/partials/navbar.html")),
	}

	mux.Handle("/game/", game.NewGameMux(r2, conn))
	mux.setupLayout()

	return mux
}

func (mux *InfinityMux) setupLayout() {
	type Category struct {
		Slug  string
		Label string
	}

	categories := []Category{
		{Slug: "new", Label: "NEW"},
		{Slug: "popular", Label: "POPULAR"},
		{Slug: "action", Label: "ACTION"},
		{Slug: "racing", Label: "RACING"},
		{Slug: "shooting", Label: "SHOOTING"},
		{Slug: "sports", Label: "SPORTS"},
		{Slug: "strategy", Label: "STRATEGY"},
		{Slug: "puzzle", Label: "PUZZLE"},
		// {Slug: "io", Label: ".IO"},
		{Slug: "2-player", Label: "2 PLAYER"},
	}

	data := struct{ Categories []Category }{
		Categories: categories,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err := mux.tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Template execution error:", err)
		}
	})
}
