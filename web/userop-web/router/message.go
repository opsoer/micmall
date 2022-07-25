package router

import (
	"github.com/gin-gonic/gin"
	"userop-web/api/message"
	"userop-web/middlewares"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth()).Use(middlewares.Trace())
	{
		MessageRouter.GET("", message.List)
		MessageRouter.POST("", message.New)
	}
}