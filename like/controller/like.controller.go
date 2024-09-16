package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/like/models"
	"github.com/patelajay745/Microservice-VideoApp/like/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ToggleVideoLike() {

}

func ToggleCommentLike() {

}

func ToggelTweetLike() {

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
