package handlers

import (
	"net/http"
	"os"

	"github.com/DalinarLG/twittor/middlewares"
	"github.com/DalinarLG/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers endpoints
func Handlers() {
	r := mux.NewRouter()

	r.HandleFunc("/api/register", middlewares.CheckDb(routers.Register)).Methods("POST")
	r.HandleFunc("/api/login", middlewares.CheckDb(routers.LoginRouter)).Methods("GET")
	r.HandleFunc("/api/profile", middlewares.CheckDb(middlewares.ValidateJwt(routers.ViewProfile))).Methods("GET")
	r.HandleFunc("/api/update", middlewares.CheckDb(middlewares.ValidateJwt(routers.UpdateUser))).Methods("PUT")
	r.HandleFunc("/api/tweet", middlewares.CheckDb(middlewares.ValidateJwt(routers.RegisterTweet))).Methods("POST")
	r.HandleFunc("/api/listweets", middlewares.CheckDb(middlewares.ValidateJwt(routers.ListTweets)))
	r.HandleFunc("/api/deltweet", middlewares.CheckDb(middlewares.ValidateJwt(routers.DelTweet))).Methods("DELETE")
	r.HandleFunc("/api/uploadavatar", middlewares.CheckDb(middlewares.ValidateJwt(routers.UploadAvatar))).Methods("POST")
	r.HandleFunc("/api/getavatar", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetAvatar))).Methods("GET")
	r.HandleFunc("/api/insertrelation", middlewares.CheckDb(middlewares.ValidateJwt(routers.InsertRelation))).Methods("POST")
	r.HandleFunc("/api/delrelation", middlewares.CheckDb(middlewares.ValidateJwt(routers.DelRelation))).Methods("DELETE")
	r.HandleFunc("/api/queryrelation", middlewares.CheckDb(middlewares.ValidateJwt(routers.QueryRelation))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":"+PORT, handler)

}
