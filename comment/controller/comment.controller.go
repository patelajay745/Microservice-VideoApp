package controller

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patelajay745/Microservice-VideoApp/comment/config"
	"github.com/patelajay745/Microservice-VideoApp/comment/models"
	"github.com/patelajay745/Microservice-VideoApp/comment/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var commentCollection *mongo.Collection = config.GetCollection(config.DB, "comments")

func GetVideoComments(c echo.Context) error {

	videoId := c.Param("videoId")

	if videoId == "" {
		return c.JSON(400, utils.ResError{
			Success: false,
			Error:   fmt.Errorf("VideoId is required"),
		})
	}

	videoIdObj, _ := primitive.ObjectIDFromHex(videoId)

	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{{"video", videoIdObj}}},
		},
		{
			{"$lookup", bson.D{
				{"from", "users"},
				{"localField", "owner"},
				{"foreignField", "_id"},
				{"as", "owner"},
			}},
		},
		{
			{"$unwind", bson.D{
				{"path", "$ownerData"},
				{"preserveNullAndEmptyArrays", true},
			}},
		},
		{
			{"$project", bson.D{
				{"content", 1},
				{"createdAt", 1},
				{"updatedAt", 1},
				{"video", 1},
				{"owner._id", 1},
				{"owner.fullName", 1},
				{"owner.avatar", 1},
			}},
		},
	}

	cursor, err := commentCollection.Aggregate(c.Request().Context(), pipeline)

	if err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	var comments []bson.M
	if err = cursor.All(c.Request().Context(), &comments); err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	if len(comments) == 0 {
		return c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "No Comment found",
		})
	}

	return c.JSON(200, utils.ResComment{
		StatusCode: 200,
		Data:       comments,
		Message:    "Comments are fetched",
		Success:    true,
	})
}

func AddComment(c echo.Context) error {

	videoId := c.Param("videoId")

	if videoId == "" {
		return c.JSON(400, utils.ResError{
			Success: false,
			Error:   fmt.Errorf("VideoId is required"),
		})
	}

	videoIdObj, _ := primitive.ObjectIDFromHex(videoId)

	var comment models.Comment

	if err := c.Bind(&comment); err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})

	}

	comment.ID = primitive.NewObjectID()
	userId := c.Get("_id")
	userIdObj, _ := primitive.ObjectIDFromHex(userId.(string))
	comment.Owner = &userIdObj
	comment.Video = &videoIdObj
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	_, err := commentCollection.InsertOne(c.Request().Context(), comment)

	if err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	commentBson, err := bson.Marshal(comment)
	if err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	var commentMap bson.M
	if err := bson.Unmarshal(commentBson, &commentMap); err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	return c.JSON(201, utils.ResComment{
		StatusCode: 201,
		Data:       []bson.M{commentMap},
		Message:    "Comment has been added succesfully",
		Success:    true,
	})
}

func DeleteComment(c echo.Context) error {

	commentId := c.Param("commentId")

	if commentId == "" {
		return c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "Please provide commentId",
		})
	}
	commentIdObj, _ := primitive.ObjectIDFromHex(commentId)

	result, err := commentCollection.DeleteOne(c.Request().Context(), bson.M{
		"_id": commentIdObj,
	})
	if err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}
	if result.DeletedCount == 0 {
		return c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "No comment found",
		})
	}

	return c.JSON(200, utils.ResMessage{
		Success: true,
		Message: "comment has been deleted",
	})
}

func UpdateComment(c echo.Context) error {

	commentId := c.Param("commentId")
	if commentId == "" {
		return c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "Please provide commentId",
		})
	}
	commentIdObj, _ := primitive.ObjectIDFromHex(commentId)
	var newComment models.Comment
	if err := c.Bind(&newComment); err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	//check if new content is provided
	if newComment.Content == "" {
		return c.JSON(400, utils.ResMessage{
			Success: false,
			Message: "Please provide content",
		})
	}

	var oldComment models.Comment

	err := commentCollection.FindOne(c.Request().Context(), bson.M{
		"_id": commentIdObj,
	}).Decode(&oldComment)

	//check if old comment is found
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(400, utils.ResMessage{
				Success: false,
				Message: "No comment found",
			})
		}
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	oldComment.Content = newComment.Content
	oldComment.UpdatedAt = time.Now()

	_, err = commentCollection.ReplaceOne(c.Request().Context(), bson.M{
		"_id": commentIdObj,
	}, oldComment)

	if err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	oldCommentBson, err := bson.Marshal(oldComment)
	if err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	var oldCommentMap bson.M

	if err := bson.Unmarshal(oldCommentBson, &oldCommentMap); err != nil {
		return c.JSON(500, utils.ResError{
			Success: false,
			Error:   err,
		})
	}

	return c.JSON(200, utils.ResComment{
		Success:    true,
		Message:    "Tweet has been updated",
		Data:       []bson.M{oldCommentMap},
		StatusCode: 200,
	})
}
