package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/patelajay745/Microservice-VideoApp/comment/controller"
	"github.com/patelajay745/Microservice-VideoApp/comment/middleware"
)

func SetUpRouter(e *echo.Echo) {

	comment := e.Group("/api/v1/comments", middleware.VerifyJWT())

	comment.GET("/:videoId", controller.GetVideoComments)
	comment.POST("/:videoId", controller.AddComment)
	comment.DELETE("/c/:commentId", controller.DeleteComment)
	comment.PATCH("/c/:commentId", controller.UpdateComment)

}
