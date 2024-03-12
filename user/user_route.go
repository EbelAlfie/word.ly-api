package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(group *gin.RouterGroup) gin.HandlerFunc {
	group.POST("/login")
}
