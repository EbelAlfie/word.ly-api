package domain

import (
	"github.com/gin-gonic/gin"
)

type QuizRepository interface {
	GetQuizes()
}

type QuizController interface {
	GetQuizes(context *gin.Context)
}
