package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/subscription/config"
	"github.com/patelajay745/Microservice-VideoApp/subscription/models"
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
	userId, _ := c.Get("_id")

	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))

	channelId := c.Param("channelId")

	channelIdObj, _ := primitive.ObjectIDFromHex(channelId)

	var subscription []bson.M
	err := subscriptionCollection.FindOne(c, bson.M{"channel": channelIdObj}).Decode(&subscription)

	if err != nil {
		if err == mongo.ErrNoDocuments {

			newSubsciprion := models.Subscription{
				ID:         primitive.NewObjectID(),
				Subscriber: userIdObj,
				Channel:    channelIdObj,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}

			subscriptionCollection.InsertOne(c, newSubsciprion)

			c.JSON(200, utils.ResMessage{
				Message: "Subscription has been added",
				Success: true,
			})
			return
		} else {
			c.JSON(500, utils.ResError{
				Success: false,
				Error:   err,
			})

			return
		}
	}

	// Matching document found
	if len(subscription) > 1 {
		subscriptionCollection.DeleteOne(c, bson.M{
			"channel":    channelIdObj,
			"subscriber": userIdObj,
		})

		c.JSON(200, utils.ResMessage{
			Message: "Unsubscribed successfully",
			Success: true,
		})
	}

}

func GetUserChannelSubscribers(c *gin.Context) {

	subscriptionId := c.Param("subscriptionId")

	if subscriptionId == "" {
		c.JSON(400, utils.ResError{
			Success: false,
			Error:   fmt.Errorf("subscriptionId is required"),
		})
		return
	}

	var subscribers []bson.M
	cursor, err := subscriptionCollection.Find(c, bson.M{"channel": subscriptionId})

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	err = cursor.All(c.Request.Context(), &subscribers)

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	c.JSON(200, utils.ResSubscription{
		StatusCode: 200,
		Data:       subscribers,
		Message:    "Subscribers are fetched",
		Success:    true,
	})

}
