package db

import (
	"context"
	"time"

	"github.com/DalinarLG/twittor/models"
)

func DelRelation(re models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := MongoCN.Database("twitter").Collection("relation")

	_, err := col.DeleteOne(ctx, re)
	if err != nil {
		return false, err
	}

	return true, nil
}
