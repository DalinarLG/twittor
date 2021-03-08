package main

import (
	"log"
	"github.com/DalinarLG/twittor/handlers"
	"github.com/DalinarLG/twittor/db"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("No database connection")
		return
	}
	handlers.Handlers()
}
