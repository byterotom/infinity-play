package admin

import (
	"net/http"

	"github.com/byterotom/infinity-play/internal/auth"
	"github.com/byterotom/infinity-play/views"
	"github.com/byterotom/infinity-play/views/components"
)

func loggedIn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cook, err := r.Cookie("infinity")
		if err != nil {
			http.SetCookie(w, &http.Cookie{
				Name:     "infinity",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			})
			views.Index(components.Login()).Render(r.Context(), w)
			return
		}
		tokenString := cook.Value
		if _, err := auth.ValidateJwt(tokenString); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.SetCookie(w, &http.Cookie{
				Name:     "infinity",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			})
			views.Index(components.Login()).Render(r.Context(), w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
