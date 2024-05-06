package route

import (
	"github.com/gin-gonic/gin"

	controller "wordly/api/controller"
)

func QuizRoute(group *gin.RouterGroup): gin.HandlerFunc {
	quizController := controller.CreateQuizController()
	return gin.HandlerFunc(func(ctx *gin.Context) {
		group.POST("/all-quiz", quizController.GetQuizes)
	})
}