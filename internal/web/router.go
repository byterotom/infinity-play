package web

import (
	"context"
	"errors"
	"net/http"

	"github.com/byterotom/infinity-play/internal/db/dbgen"
	"github.com/byterotom/infinity-play/internal/web/admin"
	"github.com/byterotom/infinity-play/internal/web/game"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
	"github.com/jackc/pgx/v5"
)

type InfinityMux struct {
	*http.ServeMux
	conn *pgx.Conn
}

func NewInfinityMux(r2 *pkg.R2, conn *pgx.Conn) http.Handler {
	mux := &InfinityMux{
		ServeMux: http.NewServeMux(),
		conn:     conn,
	}

	mux.HandleFunc("/", mux.home)
	mux.HandleFunc("GET /category/{cat}", mux.category)
	mux.Handle("/game/", game.NewGameMux(r2, conn))
	mux.Handle("/admin/", admin.NewAdminMux(r2, conn))

	return mux
}

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

	err = views.Index(components.Tag(cat, false, games)).Render(r.Context(), w)
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

	err = views.Index(components.Home(games)).Render(r.Context(), w)
}
