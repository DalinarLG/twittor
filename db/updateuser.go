package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twitter").Collection("user")
	record := make(map[string]interface{})

	if len(user.Name) > 0 {
		record["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		record["lastname"] = user.Lastname
	}

	record["birthday"] = user.Birthday

	if len(user.Avatar) > 0 {
		record["avatar"] = user.Avatar
	}

	if len(user.Bio) > 0 {
		record["bio"] = user.Bio
	}

	if len(user.Location) > 0 {
		record["location"] = user.Location
	}

	updateString := bson.M{
		"$set": record,
	}

	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": id}}
	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
