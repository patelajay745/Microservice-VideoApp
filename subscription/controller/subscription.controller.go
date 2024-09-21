package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/subscription/config"
	"github.com/patelajay745/Microservice-VideoApp/subscription/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var subscriptionCollection *mongo.Collection = config.GetCollection(config.DB, "subscriptions")

func GetSubscribedChannels(c *gin.Context) {

	userId, _ := c.Get("_id")

	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))

	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{{"subscriber", userIdObj}}},
		},
		{
			{"$lookup", bson.D{
				{"from", "users"},
				{"localField", "channel"},
				{"foreignField", "_id"},
				{"as", "channel"},
			}},
		},
		{
			{"$unwind", bson.D{
				{"path", "$channel"},
				{"preserveNullAndEmptyArrays", true},
			}},
		},
		{
			{"$project", bson.D{
				{"subscriber", 1},
				{"channel.userName", 1},
				{"channel.avatar", 1},
				{"channel.fullName", 1},
			}},
		},
	}

	cursor, err := subscriptionCollection.Aggregate(c, pipeline)

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	var subscriptions []bson.M
	if err = cursor.All(c, &subscriptions); err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	if len(subscriptions) == 0 {
		c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "No Subscription found",
		})
		return
	}

	c.JSON(200, utils.ResSubscription{
		StatusCode: 200,
		Data:       subscriptions,
		Message:    "Subscription are fetched",
		Success:    true,
	})

}

func ToggleSubscription(c *gin.Context) {

}

func GetUserChannelSubscribers(c *gin.Context) {

}
