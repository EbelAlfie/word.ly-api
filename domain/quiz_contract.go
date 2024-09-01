package domain

import (
	"github.com/gin-gonic/gin"
)

type QuizRepository interface {
	GetQuiz(quiz QuizType) (*QuizModel, error)
	GetDetailByUserId(teacherId string) (*QuizModel, error)
	UpdateQuiz() (*QuizModel, error)
	InsertQuiz() error
}

type QuizController interface {
	GetQuiz(context *gin.Context)
	GetDetailByUserId(context *gin.Context)
	UpdateQuiz(context *gin.Context)
	InsertQuiz(context *gin.Context)
}
