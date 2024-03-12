package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(group *gin.RouterGroup) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		group.POST("/login")
	})
}
