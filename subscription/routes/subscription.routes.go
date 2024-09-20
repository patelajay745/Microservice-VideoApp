package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/subscription/controller"
	"github.com/patelajay745/Microservice-VideoApp/subscription/middleware"
)

func SetUpRouter(router *gin.Engine) {
	subscription := router.Group("/api/v1/subscription")
	{
		subscription.GET("/c", middleware.VerifyJWT(), func(c *gin.Context) {
			controller.GetSubscribedChannels(c)
		})
		subscription.POST("/c/:channelId", middleware.VerifyJWT(), func(c *gin.Context) {
			controller.ToggleSubscription(c)
		})
		subscription.GET("/u/:subscriptionId", middleware.VerifyJWT(), func(c *gin.Context) {
			controller.GetUserChannelSubscribers(c)
		})
	}
}
