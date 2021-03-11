package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
)

//RegisterUser insert a new user to the db
func RegisterUser(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twitter").Collection("user")

	u.Password, _ = EncryptPass(u.Password)

	_, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//ObjID, _ := result.InsertedID(primitive.ObjectID)
	return "ok", true, nil

}
