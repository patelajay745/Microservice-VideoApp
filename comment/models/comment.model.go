package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Content   string              `bson:"content" json:"content" form:"content"`
	Video     *primitive.ObjectID `bson:"video" json:"video" form:"video"`
	Owner     *primitive.ObjectID `bson:"owner" json:"owner" form:"owner"`
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
}
