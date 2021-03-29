package routers

import (
	"net/http"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"
)

func DelRelation(w http.ResponseWriter, r *http.Request) {
	var re models.Relation
	fid := r.URL.Query().Get("fid")
	if len(fid) < 1 {
		http.Error(w, "follower id is required", 400)
		return
	}

	re.UserID = UserID
	re.FollowerID = fid

	status, err := db.DelRelation(re)
	if !status {
		http.Error(w, "error deleting relation", 400)
		return
	}

	if err != nil {
		http.Error(w, "error deleting relation "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
