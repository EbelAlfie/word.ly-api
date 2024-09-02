package route

import (
	"github.com/gin-gonic/gin"

	controller "wordly/api/controller"
	repository "wordly/api/repository"
)

func QuizRoute(group *gin.RouterGroup) {
	quizRepository := repository.CreateQuizRepository()
	quizController := controller.CreateQuizController(quizRepository)

	group.GET("/quizes", quizController.GetQuiz)
	group.POST("/add-quiz", quizController.InsertQuiz)
	group.GET("/quiz-detail", quizController.GetQuizDetail)
}
