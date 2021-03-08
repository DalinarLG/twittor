package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckEmail checks if email is already inserted
func CheckEmail(email string)(models.User, bool, string){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	var result models.User
	
	db := MongoCN.Database("twitter")
	col := db.Collection("user")
	cond := bson.M{"email":email}
	err := col.FindOne(ctx, cond).Decode(&result)

	id := result.ID.Hex()

	if err != nil {
		return result, false, id
	}

	return result, true, id


}