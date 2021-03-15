package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/jwt"
	"github.com/DalinarLG/twittor/models"
)

//LoginRouter router
func LoginRouter(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type","Application/Json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Email or password invalid "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0{
		http.Error(w, "Email is required ", 400)
		return
	}

	if len(user.Password) == 0{
		http.Error(w, "Password is required ", 400)
		return
	}

	result,status := db.Login(user.Email, user.Password)
	if !status {
		http.Error(w, "Email or password invalid ", 400)
		return
	}

	jwtk, err := jwt.GenerateToken(result)
	if err != nil {
		http.Error(w, "Error generating token "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token : jwtk,
	}

	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//set cookie on user browser

	expirationTime := time.Now().Add(24 *time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name: "Token",
		Value: jwtk,
		Expires: expirationTime,
	})

}