package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"
)

func RegisterTweet(w http.ResponseWriter, r *http.Request) {
	var t models.Tweet
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Something happend "+err.Error(), 400)
		return
	}

	if len(t.Body) == 0 {
		http.Error(w, "Must write something ", 400)
		return
	}

	tweet := models.Tweet{
		Body: t.Body,
		UserId: UserID,
		Date: time.Now(),
	}
	_, status, err := db.RegisterTweet(tweet)
	if !status {
		http.Error(w, "Error trying to insert tweet", 400)
		return
	}

	if err != nil {
		http.Error(w, "something happened "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
