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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":"+PORT, handler)

}
