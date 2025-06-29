package admin

import (
	"fmt"
	"net/http"

	"github.com/byterotom/infinity-play/internal/auth"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
)

func loggedIn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cook, err := r.Cookie("infinity")
		redirectTo := fmt.Sprintf("/admin/%s", r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			views.Index(components.Login(redirectTo)).Render(r.Context(), w)
			return
		}
		tokenString := cook.Value
		if _, err := auth.ValidateJwt(tokenString); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			views.Index(components.Login(redirectTo)).Render(r.Context(), w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
