package route

import (
	"github.com/gin-gonic/gin"

	controller "wordly/api/controller"
	repository "wordly/api/repository"
)

func UserRoute(group *gin.RouterGroup) gin.HandlerFunc {
	userRepo := repository.CreateUserRepo()
	userController := controller.CreateUserController(userRepo)
	return gin.HandlerFunc(func(ctx *gin.Context) {
		group.POST("/login", userController.Login)
		group.POST("/register", userController.Register)
	})
}
