package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/like/models"
	"github.com/patelajay745/Microservice-VideoApp/like/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ToggleVideoLike(c *gin.Context, client *mongo.Client) {

	videoId := c.Param("videoId")
	userId, _ := c.Get("_id")

	if videoId == "" {
		c.JSON(400, utils.ResError{
			Success: false,
			Error:   fmt.Errorf("VideoId is required"),
		})
		return
	}

	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))
	videoIdObj, _ := primitive.ObjectIDFromHex(videoId)

	likeCollection := client.Database("videoapp").Collection("likes")

	//find if use has liked video
	var existingLike models.Like
	err := likeCollection.FindOne(c.Request.Context(), bson.M{
		"video":   videoIdObj,
		"likedBy": userIdObj,
	}).Decode(&existingLike)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			//no matching document found, so create new document
			newLike := models.Like{
				ID:        primitive.NewObjectID(),
				Video:     &videoIdObj,
				LikedBy:   userIdObj,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			likeCollection.InsertOne(c.Request.Context(), newLike)

			c.JSON(200, utils.ResLikes{
				StatusCode: 200,
				Data:       []models.Like{newLike},
				Message:    "liked",
				Success:    true,
			})
		} else {
			// Some other error occurred
			c.JSON(500, utils.ResError{
				Success: false,
				Error:   err,
			})
			return
		}
	} else {
		//matching document found
		likeCollection.DeleteOne(c.Request.Context(), bson.M{
			"video":   videoIdObj,
			"likedBy": userIdObj,
		})
		c.JSON(200, utils.ResLikes{
			StatusCode: 200,
			Data:       []models.Like{},
			Message:    "unliked",
			Success:    true,
		})
		return
	}
}

func ToggleCommentLike(c *gin.Context, client *mongo.Client) {

	commentId := c.Param("commentId")
	userId, _ := c.Get("_id")
	if commentId == "" {
		c.JSON(400, utils.ResError{
			Success: false,
			Error:   fmt.Errorf("CommentId is required"),
		})
		return
	}

	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))
	commentIdObj, _ := primitive.ObjectIDFromHex(commentId)

	likeCollection := client.Database("videoapp").Collection("likes")

	//find if user has liked the comment
	var existingLike models.Like

	err := likeCollection.FindOne(c.Request.Context(), bson.M{
		"comment": commentIdObj,
		"likedBy": userIdObj,
	}).Decode(&existingLike)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			newLike := models.Like{
				ID:        primitive.NewObjectID(),
				Comment:   &commentIdObj,
				LikedBy:   userIdObj,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			likeCollection.InsertOne(c.Request.Context(), newLike)
			c.JSON(200, utils.ResLikes{
				StatusCode: 200,
				Data:       []models.Like{newLike},
				Message:    "liked",
				Success:    true,
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

	likeCollection.DeleteOne(c.Request.Context(), bson.M{
		"comment": commentIdObj,
		"likedBy": userIdObj,
	})

	c.JSON(200, utils.ResLikes{
		StatusCode: 200,
		Data:       []models.Like{},
		Message:    "unliked",
		Success:    true,
	})
}

func ToggelTweetLike(c *gin.Context, client *mongo.Client) {

	tweetId := c.Param("tweetId")
	userId, _ := c.Get("_id")

	if tweetId == "" {

		c.JSON(400, utils.ResError{
			Success: false,
			Error:   fmt.Errorf("tweetId is required"),
		})
		return
	}

	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))
	tweetIdObj, _ := primitive.ObjectIDFromHex(tweetId)

	likeCollection := client.Database("videoapp").Collection("likes")

	// find if user has liked the tweet
	var existingLike models.Like
	err := likeCollection.FindOne(c.Request.Context(), bson.M{
		"tweet":   tweetIdObj,
		"likedBy": userIdObj,
	}).Decode(&existingLike)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No matching document found
			newLike := models.Like{
				ID:        primitive.NewObjectID(),
				Tweet:     &tweetIdObj,
				LikedBy:   userIdObj,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			likeCollection.InsertOne(c.Request.Context(), newLike)

			//send new liked tweet as response
			c.JSON(200, utils.ResLikes{
				StatusCode: 200,
				Data:       []models.Like{newLike},
				Message:    "liked",
				Success:    true,
			})
			return
		} else {
			// Some other error occurred
			c.JSON(500, utils.ResError{
				Success: false,
				Error:   err,
			})
			return
		}
	}
	// Matching document found
	likeCollection.DeleteOne(c.Request.Context(), bson.M{
		"tweet":   tweetIdObj,
		"likedBy": userIdObj,
	})

	c.JSON(200, utils.ResLikes{
		StatusCode: 200,
		Data:       []models.Like{},
		Message:    "unliked",
		Success:    true,
	})

}

func GetLikedVideos(c *gin.Context, client *mongo.Client) {

	likedVideos := []models.Like{}

	likeCollection := client.Database("videoapp").Collection("likes")

	userId, _ := c.Get("_id")

	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))

	cursor, err := likeCollection.Find(c.Request.Context(), bson.M{
		"likedBy": userIdObj,
		"video": bson.M{
			"$exists": true,
		},
	})

	fmt.Println("cursor", cursor)

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	err = cursor.All(c.Request.Context(), &likedVideos)

	if err != nil {
		c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
		return
	}

	c.JSON(200, utils.ResLikes{
		StatusCode: 200,
		Data:       likedVideos,
		Message:    "Liked Videos are retrieved successfully",
		Success:    true,
	})
}
