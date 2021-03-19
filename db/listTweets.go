package db

import (
	"context"
	"log"
	"time"

	"github.com/DalinarLG/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListTweets(id string, pag int64) ([]models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var tweets []models.Tweet
	col := MongoCN.Database("twitter").Collection("tweet")
	cond := bson.M{"userid": id}
	opt := options.Find()
	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "Date", Value: -1}})
	opt.SetSkip((pag - 1) * 20)

	cursor, err := col.Find(ctx, cond, opt)
	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}

	for cursor.Next(context.TODO()) {
		var tweet models.Tweet
		err := cursor.Decode(&tweet)
		if err != nil {
			return tweets, false
		}

		tweets = append(tweets, tweet)
	}
	return tweets, true

}
