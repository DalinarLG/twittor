package models

type Relation struct {
	UserID     string `bson:"userid" json:"userId"`
	FollowerID string `bson:"followerid" json:"followerId"`
}
