package routers

import (
	"net/http"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"
)

func InsertRelation(w http.ResponseWriter, r *http.Request) {
	var re models.Relation
	fid := r.URL.Query().Get("fid")
	if len(fid) < 1 {
		http.Error(w, "follower id is required", 400)
		return
	}

	re.UserID = UserID
	re.FollowerID = fid

	status, err := db.InsertRelation(re)
	if !status {
		http.Error(w, "could not create relation", 400)
		return
	}
	if err != nil {
		http.Error(w, "could not create relation "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
