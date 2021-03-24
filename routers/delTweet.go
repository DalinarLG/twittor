package routers

import (
	"net/http"

	"github.com/DalinarLG/twittor/db"
)

func DelTweet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "id is required", 400)
		return
	}

	err := db.DelTweet(id, UserID)
	if err != nil {
		http.Error(w, "Error deleting tweet "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
