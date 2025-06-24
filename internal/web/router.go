package web

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

type InfinityMux struct {
	*http.ServeMux
	Tmpl *template.Template
}

func NewInfinityMux(content *embed.FS) *InfinityMux {
	mux := &InfinityMux{
		ServeMux: http.NewServeMux(),
		Tmpl:     template.Must(template.ParseFS(content, "templates/index.html", "templates/partials/navbar.html")),
	}
	mux.initializeLayout()
	return mux
}

func (mux *InfinityMux) initializeLayout() {
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

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err := mux.Tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Template execution error:", err)
		}
	})
}
