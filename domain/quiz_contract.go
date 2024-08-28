package domain

import (
	"github.com/gin-gonic/gin"
)

type QuizRepository interface {
	GetCerpen() (*QuizModel, error)
	GetKalimatEfektif() (*QuizModel, error)
}

type QuizController interface {
	GetCerpen(context *gin.Context)
	GetKalimatEfektif(context *gin.Context)
}
