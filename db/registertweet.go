package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func RegisterTweet(t models.Tweet)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twitter").Collection("tweet")
	//tweet := bson.M{}
	result, err := col.InsertOne(ctx, t)
	if err != nil {
		return "",false, err
	}

	objId :=  result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
