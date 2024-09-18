package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tweet struct {
	ID      primitive.ObjectID  `bson:"_id" json:"_id" form:"_id"`
	Content string              `bson:"content" json:"content" form:"content"`
	Owner   *primitive.ObjectID `bson:"owner" json:"owner" form:"owner"`
}
