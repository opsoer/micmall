package router

import (
	"github.com/gin-gonic/gin"
	"userop-web/api/user_fav"
	"userop-web/middlewares"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	UserFavRouter := Router.Group("userfavs").Use(middlewares.Trace())
	{
		UserFavRouter.DELETE("/:id", middlewares.JWTAuth(), user_fav.Delete)
		UserFavRouter.GET("/:id", middlewares.JWTAuth(), user_fav.Detail)
		UserFavRouter.POST("", middlewares.JWTAuth(), user_fav.New)
		UserFavRouter.GET("", middlewares.JWTAuth(), user_fav.List)
	}
}