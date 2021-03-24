package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	ext := strings.Split(handler.Filename, ".")[1]
	path := "uploads/avatar" + UserID + "." + ext

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error uploading image "+err.Error(), 400)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error uploading image "+err.Error(), 400)
		return
	}

	var user models.User
	user.Avatar = UserID + "." + ext

	status, err := db.UpdateUser(user, UserID)
	if !status {
		http.Error(w, "error updateing info", 400)
		return
	}
	if err != nil {
		http.Error(w, "error updateing info "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
