package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID        primitive.ObjectID  `bson:"_id"`
	Video     *primitive.ObjectID `bson:"video,omitempty"`
	Tweet     *primitive.ObjectID `bson:"tweet,omitempty"`
	Comment   *primitive.ObjectID `bson:"comment,omitempty"`
	LikedBy   primitive.ObjectID  `bson:"likedBy"`
	CreatedAt time.Time           `bson:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt"`
}
