package domain

import (
	"github.com/gin-gonic/gin"
)

type QuizRepository interface {
	GetQuiz(quiz QuizType) ([]QuizModel, error)
	GetQuizesByUserId(teacherId string) ([]QuizModel, error)
	GetQuizDetail(teacherId string) (*QuizModel, error)
	UpdateQuiz(request QuizRequest) error
	InsertQuiz(teacherId string, request QuizRequest) error
}

type QuizController interface {
	GetQuiz(context *gin.Context)
	GetQuizesByUserId(context *gin.Context)
	GetQuizDetail(context *gin.Context)
	UpdateQuiz(context *gin.Context)
	InsertQuiz(context *gin.Context)
}
