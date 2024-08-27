package domain

import (
	"github.com/gin-gonic/gin"
)

type QuizRepository interface {
	GetQuizes(quizType string)
}

type QuizController interface {
	GetQuizes(context *gin.Context)
}
