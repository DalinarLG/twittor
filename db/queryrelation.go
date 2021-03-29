package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)


func QueryRelation(t models.Relation)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15 *time.Second)
	defer cancel()

	col := MongoCN.Database("twitter").Collection("relation")
	cond := bson.M{
		"userid": t.UserID,
		"followerid":t.FollowerID,
	}

	var result models.Relation
	err := col.FindOne(ctx, cond).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}