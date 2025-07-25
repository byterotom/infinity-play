package web

import (
	"context"
	"errors"
	"net/http"

	"github.com/byterotom/infinity-play/internal/db/dbgen"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
)

func (mux *InfinityMux) category(w http.ResponseWriter, r *http.Request) {

	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	var cats map[string]bool = map[string]bool{
		"action":   true,
		"racing":   true,
		"shooting": true,
		"sports":   true,
		"strategy": true,
		"puzzle":   true,
		"io":       true,
		"2-player": true,
	}
	cat := r.PathValue("cat")
	if !cats[cat] {
		err = errors.New("NOT FOUND")
		return
	}

	q := dbgen.New(mux.conn)

	games, err := q.GetGamesByTag(context.TODO(), cat)
	if err != nil {
		return
	}
	if pkg.IsHTMXRequest(r) {
		components.Tag(cat, false, false, games).Render(r.Context(), w)
		return
	}
	views.Index(components.Tag(cat, false, false, games)).Render(r.Context(), w)
}

func (mux *InfinityMux) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	var games map[string][]dbgen.Game = make(map[string][]dbgen.Game)

	q := dbgen.New(mux.conn)

	games["new"], err = q.GetNewGames(context.TODO())
	if err != nil {
		return
	}
	games["top rated"], err = q.GetTopRatedGames(context.TODO())
	if err != nil {
		return
	}
	games["popular"], err = q.GetPopularGames(context.TODO())
	if err != nil {
		return
	}
	if pkg.IsHTMXRequest(r) {
		components.Home(games).Render(r.Context(), w)
		return
	}
	views.Index(components.Home(games)).Render(r.Context(), w)
}

func (mux *InfinityMux) searchGame(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	pattern := r.URL.Query().Get("q")

	q := dbgen.New(mux.conn)

	games, err := q.GetGamesByPattern(ctx, pattern)
	if err != nil {
		return
	}
	del := r.URL.Query().Get("d") == "1"
	if pkg.IsHTMXRequest(r) {
		components.Tag(pattern, true, del, games).Render(r.Context(), w)
	} else {
		views.Index(components.Tag(pattern, true, del, games)).Render(r.Context(), w)
	}
}