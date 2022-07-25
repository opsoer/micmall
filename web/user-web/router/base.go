package router

import (
	"github.com/gin-gonic/gin"
	"user-web/api"
	"user-web/middlewares"
)

func InitBaseRouter(Router *gin.RouterGroup){
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha).Use(middlewares.Trace())
	}
}
