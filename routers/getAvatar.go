package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/DalinarLG/twittor/db"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "id is required", 400)
		return
	}

	user, err := db.FindProfile(id)
	if err != nil {
		http.Error(w, "error "+err.Error(), 400)
		return
	}

	file, err := os.Open("/uploads/avatar" + user.Avatar)
	if err != nil {
		http.Error(w, "error "+err.Error(), 400)
		return
	}
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "error getting image "+err.Error(), 400)
	}

}
