package route

import (
	"github.com/gin-gonic/gin"

	controller "wordly/api/controller"
	repository "wordly/api/repository"
)

func QuizRoute(group *gin.RouterGroup) gin.HandlerFunc {
	quizRepository := repository.CreateQuizRepository()
	quizController := controller.CreateQuizController(quizRepository)
	return gin.HandlerFunc(func(ctx *gin.Context) {
		group.POST("/all-quiz", quizController.GetQuizes)
	})
}
