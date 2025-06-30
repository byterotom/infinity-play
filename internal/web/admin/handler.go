package admin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/byterotom/infinity-play/internal/auth"
	"github.com/byterotom/infinity-play/internal/db/dbgen"
	"github.com/byterotom/infinity-play/pkg"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	act := r.PathValue("act")
	views.Index(components.Admin(act)).Render(r.Context(), w)
}

func (mux *AdminMux) login(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()

	r.ParseMultipartForm(10 << 20)
	username := r.FormValue("username")
	password := pkg.HashWithString(r.FormValue("password"))

	q := dbgen.New(mux.conn)

	_, err := q.GetAdminIdByCredentials(ctx, dbgen.GetAdminIdByCredentialsParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}
	tokenString, err := auth.IssueJwt(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "infinity",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (mux *AdminMux) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "infinity",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
