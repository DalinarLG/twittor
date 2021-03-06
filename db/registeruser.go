package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//RegisterUser insert a new user to the db
func RegisterUser(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twitter").Collection("user")

	u.Password, _ = EncryptPass(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
