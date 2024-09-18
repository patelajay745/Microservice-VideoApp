package controller

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/tweet/models"
	"github.com/patelajay745/Microservice-VideoApp/tweet/utils"
	"go.mongodb.org/mongo-driver/bson"
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

	userID := c.Param("userID")

	userIDObj, _ := primitive.ObjectIDFromHex(userID)

	var tweets []models.Tweet

	tweetCollection := client.Database("videoapp").Collection("tweets")

	cursor, err := tweetCollection.Find(c.Request.Context(), bson.M{
		"owner": userIDObj,
	})

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	if err = cursor.All(c.Request.Context(), &tweets); err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	if len(tweets) == 0 {
		c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "No Tweets found",
		})
		return
	}

	c.JSON(200, utils.ResTweets{
		StatusCode: 200,
		Data:       tweets,
		Message:    "Tweets are fetched",
		Success:    true,
	})

}

func DeleteTweet(c *gin.Context, client *mongo.Client) {

	tweetId := c.Param("tweetId")

	if tweetId == "" {
		c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "Please provide tweetId",
		})
		return
	}

	tweetIdObj, _ := primitive.ObjectIDFromHex(tweetId)

	tweetCollection := client.Database("videoapp").Collection("tweets")

	result, err := tweetCollection.DeleteOne(c.Request.Context(), bson.M{
		"_id": tweetIdObj,
	})

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "No Tweets found",
		})
		return
	}

	c.JSON(200, utils.ResMessage{
		Success: true,
		Message: "Tweet has been deleted",
	})

}

func UpdateTweet(c *gin.Context, client *mongo.Client) {

	tweetId := c.Param("tweetId")

	if tweetId == "" {
		c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "Please provide tweetId",
		})
		return
	}

	tweetIdObj, _ := primitive.ObjectIDFromHex(tweetId)

	var newTweet models.Tweet
	if err := c.Bind(&newTweet); err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	//check if new content is provided
	if newTweet.Content == "" {
		c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "Please provide content",
		})
		return
	}

	tweetCollection := client.Database("videoapp").Collection("tweets")

	var oldTweet models.Tweet
	tweetCollection.FindOne(c.Request.Context(), bson.M{"_id": tweetIdObj}).Decode(&oldTweet)

	//check if old tweet is found
	if oldTweet.Content == "" {
		c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "No Tweet found",
		})
		return
	}

	oldTweet.Content = newTweet.Content
	oldTweet.UpdatedAt = time.Now()

	_, err := tweetCollection.ReplaceOne(c.Request.Context(), bson.M{
		"_id": tweetIdObj,
	}, oldTweet)

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	c.JSON(200, utils.ResTweets{
		Success:    true,
		Message:    "Tweet has been updated",
		Data:       []models.Tweet{oldTweet},
		StatusCode: 200,
	})

}
