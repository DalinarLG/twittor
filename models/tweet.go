package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Body   string             `bson:"body" json:"body,omitempty"`
	Date   time.Time          `bson:"time" json:"time,omitempty"`
	UserId string             `bson:"userid" json:"userid,omitempty"`
}
