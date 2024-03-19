package user

import (
	"github.com/gin-gonic/gin"

	controller "wordly/api/controller"
)

func UserRoute(group *gin.RouterGroup) gin.HandlerFunc {
	userController := controller.CreateUserController()
	return gin.HandlerFunc(func(ctx *gin.Context) {
		group.POST("/login", userController.Login)
		group.POST("/register", userController.Register)
	})
}
