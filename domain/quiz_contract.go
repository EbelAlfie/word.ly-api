package domain

import (
	"github.com/gin-gonic/gin"
)

type QuizRepository interface {
	GetCerpen() (*QuizModel, error)
	GetKalimatEfektif() (*QuizModel, error)
	UpdateSoal() error
}

type QuizController interface {
	GetCerpen(context *gin.Context)
	GetKalimatEfektif(context *gin.Context)
	UpdateSoal(context *gin.Context)
}
