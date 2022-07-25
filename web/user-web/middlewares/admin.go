package middlewares

import (
	"github.com/gin-gonic/gin"
	"user-web/models"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc{
	//验证管理员权限
	return func(ctx *gin.Context){
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg":"无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
