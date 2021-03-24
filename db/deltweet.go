package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DelTweet(id, userid string)error{
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel ()
	col := MongoCN.Database("twitter").Collection("tweet")
	objId,_ := primitive.ObjectIDFromHex(id)
	cond := bson.M{
		"_id":objId,
		"userid":userid,
	}

	_, err := col.DeleteOne(ctx, cond)
	return err
}