package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"
)


func QueryRelation(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if len(id)<1{
		http.Error(w, "id is required", 400)
		return
	}

	var t models.Relation
	t.UserID = UserID
	t.FollowerID = id

	status, err := db.QueryRelation(t)
	if !status {
		status = false
	}
	if err != nil{
		status = false
	}

	response := models.ResponseRelation{
		Status: status,
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}