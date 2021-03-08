package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN connection
var MongoCN = connectDB()

//Connect client database
var clientOptions = options.Client().ApplyURI("mongodb+srv://dalinar:guerrero1015@cluster0.stgqr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection Ok")
	return client
}

//CheckConnection function
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}

//RegisterUser registers new user in database
