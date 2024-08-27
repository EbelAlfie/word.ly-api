package route

import (
	"github.com/gin-gonic/gin"

	controller "wordly/api/controller"
	repository "wordly/api/repository"
)

func UserRoute(group *gin.RouterGroup) {
	userRepo := repository.CreateUserRepo()
	userController := controller.CreateUserController(userRepo)

	group.POST("/login", userController.Login)
	group.POST("/register", userController.Register)
}
