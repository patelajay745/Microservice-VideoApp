package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Subscriber primitive.ObjectID `bson:"subscriber" json:"subscriber"`
	Channel    primitive.ObjectID `bson:"channel" json:"channel"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}
