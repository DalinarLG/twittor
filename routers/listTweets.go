package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DalinarLG/twittor/db"
)

func ListTweets(w http.ResponseWriter, r *http.Request ){
	id := r.URL.Query().Get("id")
	if len(id)<1{
		http.Error(w, "id is needed", 400)
		return
	}

	if len(r.URL.Query().Get("pag"))<1{
		http.Error(w, "page is needed", 400)
		return
	}
	
	pages, err := strconv.Atoi(r.URL.Query().Get("pag"))
	if err != nil {
		http.Error(w, "Error with page number", 400)
		return
	}

	pag := int64(pages)

	tweets, status := db.ListTweets(id, pag)
	if !status {
		http.Error(w, "Error looking tweets", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweets)

}