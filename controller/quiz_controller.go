package controller

import (
	"github.com/gin-gonic/gin"

	domain "wordly/api/domain"
)

type QuizControllerImpl struct {
	repository domain.QuizRepository
}

func CreateQuizController(repo domain.QuizRepository) domain.QuizController {
	return &QuizControllerImpl{
		repository: repo,
	}
}

func (cont *QuizControllerImpl) GetCerpen(context *gin.Context) {

}

func (cont *QuizControllerImpl) GetKalimatEfektif(context *gin.Context) {

}

func (cont *QuizControllerImpl) UpdateSoal(context *gin.Context) {
	
}
