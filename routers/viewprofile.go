package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DalinarLG/twittor/db"
	
)

func ViewProfile(w http.ResponseWriter, r*http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "param id is needed", http.StatusBadRequest)
		return
	}	

	user, err := db.FindProfile(ID)
	if err != nil{
		http.Error(w, "error searching profile"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}