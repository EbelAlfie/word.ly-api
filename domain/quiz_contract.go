package domain

import (
	"github.com/gin-gonic/gin"
)

type QuizRepository interface {
	GetCerpen(quizType string) (*QuizModel, error)
}

type QuizController interface {
	GetCerpen(context *gin.Context)
}
