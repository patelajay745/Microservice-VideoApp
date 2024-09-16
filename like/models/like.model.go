package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Video     *primitive.ObjectID `bson:"video,omitempty" json:"video,omitempty"`
	Tweet     *primitive.ObjectID `bson:"tweet,omitempty" json:"tweet,omitempty"`
	Comment   *primitive.ObjectID `bson:"comment,omitempty" json:"comment,omitempty"`
	LikedBy   primitive.ObjectID  `bson:"likedBy" json:"likedBy"`
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
}
