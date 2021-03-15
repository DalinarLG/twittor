package routers

import (
	"encoding/json"
	"net/http"
	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"

)
//Register function to register a new user
func Register(w http.ResponseWriter, r *http.Request){
	var user  models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error data"+err.Error(), 400)
		return
	}

	if len(user.Email)==0{
		http.Error(w, "Email required",400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password must have 6 characters at least",400)
		return
	}

	_, found, _ := db.CheckEmail(user.Email)
	if found {
		http.Error(w, "User already Exists", 400)
		return
	}

	_, status, err := db.RegisterUser(user)
	if err != nil {
		http.Error(w, "Error registering user "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Something wrong happened ", 400)
		return 
	}

	w.WriteHeader(http.StatusCreated)
}