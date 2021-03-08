package middlewares

import (
	"net/http"
	"github.com/DalinarLG/twittor/db"
)

//CheckDb check database
func CheckDb(next http.HandlerFunc)http.HandlerFunc{
	return func (w http.ResponseWriter , r *http.Request){
		if db.CheckConnection() == 0 {
			http.Error(w, "Database connection lost", 500)
			return
		}

		next.ServeHTTP(w, r)
	}


}