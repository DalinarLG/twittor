package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "something happened "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.UpdateUser(user, UserID)
	if err != nil {
		http.Error(w, "Error updating user "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Error updating user ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
