package middlewares

import (
	"net/http"

	"github.com/DalinarLG/twittor/routers"
)

func ValidateJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, _, _, err := routers.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Tokern error"+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)

	}
}
