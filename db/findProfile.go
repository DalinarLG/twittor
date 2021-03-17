package db

import (
	"context"
	"log"
	"time"

	"github.com/DalinarLG/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(id string) (models.User, error) {	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	col := MongoCN.Database("twitter").Collection("user")

	var profile models.User
	objId, _ := primitive.ObjectIDFromHex(id)	
	cond := bson.M{"_id": objId}
	err := col.FindOne(ctx, cond).Decode(&profile)	
	profile.Password = ""
	if err != nil {
		log.Println("Error buscando perfil" + err.Error())
		return profile, err
	}

	return profile, nil
}
