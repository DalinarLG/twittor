package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
)

func InsertRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twitter").Collection("relation")

	_, err := col.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}

	return true, nil

}
