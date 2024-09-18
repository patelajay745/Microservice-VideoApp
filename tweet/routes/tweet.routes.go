package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/tweet/controller"
	"github.com/patelajay745/Microservice-VideoApp/tweet/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRouter(router *gin.Engine, client *mongo.Client) {

	tweet := router.Group("/api/v1/tweet")
	{
		tweet.POST("/", middleware.VerifyJWT(), func(c *gin.Context) { controller.CreateTweet(c, client) })

		tweet.GET("/user/:userID", middleware.VerifyJWT(), func(c *gin.Context) { controller.GetUserTweets(c, client) })

		tweet.DELETE("/:tweetId", middleware.VerifyJWT(), func(c *gin.Context) { controller.DeleteTweet(c, client) })

		tweet.PATCH("/:tweetId", middleware.VerifyJWT(), func(c *gin.Context) { controller.UpdateTweet(c, client) })

	}

}
