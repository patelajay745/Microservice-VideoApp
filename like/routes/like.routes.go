package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/like/controller"
	"github.com/patelajay745/Microservice-VideoApp/like/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(router *gin.Engine, client *mongo.Client) {

	like := router.Group("/api/v1/likes")
	{
		like.POST("/toggle/v/:videoId", middleware.VerifyJWT(), func(c *gin.Context) {
			controller.ToggleVideoLike(c, client)
		})
		like.POST("/toggle/c/:commentId", middleware.VerifyJWT(), func(c *gin.Context) {
			controller.ToggleCommentLike(c, client)
		})
		like.POST("/toggle/t/:tweetId", middleware.VerifyJWT(), func(c *gin.Context) {

			controller.ToggelTweetLike(c, client)
		})
		like.GET("/videos", middleware.VerifyJWT(), func(c *gin.Context) {
			controller.GetLikedVideos(c, client)
		})
	}

}
