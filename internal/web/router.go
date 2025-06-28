package web

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/byterotom/infinity-play/internal/db/dbgen"
	"github.com/byterotom/infinity-play/internal/web/admin"
	"github.com/byterotom/infinity-play/internal/web/game"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
)

type InfinityMux struct {
	*http.ServeMux
	conn *sql.DB
}

func NewInfinityMux(r2 *pkg.R2, conn *sql.DB) http.Handler {
	mux := &InfinityMux{
		ServeMux: http.NewServeMux(),
		conn:     conn,
	}

	mux.Handle("/game/", game.NewGameMux(r2, conn))
	mux.Handle("/admin/", admin.NewAdminMux(r2, conn))
	mux.HandleFunc("/", mux.home)

	return mux
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
