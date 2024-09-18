package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/tweet/models"
	"github.com/patelajay745/Microservice-VideoApp/tweet/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTweet(c *gin.Context, client *mongo.Client) {

	userId, _ := c.Get("_id")

	var tweet models.Tweet
	if err := c.Bind(&tweet); err != nil {

		log.Fatal(err)
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	tweet.ID = primitive.NewObjectID()
	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))
	tweet.Owner = &userIdObj

	tweetCollection := client.Database("videoapp").Collection("tweets")

	_, err := tweetCollection.InsertOne(c.Request.Context(), tweet)

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	c.JSON(201, utils.ResTweets{
		StatusCode: 201,
		Data:       []models.Tweet{tweet},
		Message:    "Tweet has been added succesfully",
		Success:    true,
	})

}

func GetUserTweets(c *gin.Context, client *mongo.Client) {

}

func DeleteTweet(c *gin.Context, client *mongo.Client) {

}

func UpdateTweet(c *gin.Context, client *mongo.Client) {

}
