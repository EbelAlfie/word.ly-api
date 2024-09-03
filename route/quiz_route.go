package route

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	controller "wordly/api/controller"
	"wordly/api/middleware"
	repository "wordly/api/repository"
)

func QuizRoute(group *gin.RouterGroup) {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("SECRET")

	quizRepository := repository.CreateQuizRepository()
	quizController := controller.CreateQuizController(quizRepository)

	public := group.Group("")

	public.GET("/quizes", quizController.GetQuiz)

	private := group.Group("")
	private.Use(middleware.JwtAuthMiddleware(secret))

	private.POST("/add-quiz", quizController.InsertQuiz)
	private.GET("/quiz-detail", quizController.GetQuizDetail)
	private.GET("/my-quiz", quizController.GetQuizesByUserId)
	private.POST("/update", quizController.UpdateQuiz)
}
