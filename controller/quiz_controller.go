package controller

import (
	"github.com/gin-gonic/gin"

	repository "wordly/api/repository"
)

type QuizControllerImpl struct {
	repo domain.QuizRepository
}

func CreateQuizController(repo domain.QuizRepository) domain.QuizController {
	return &QuizControllerImpl {
		repo
	}
}

func (cont *domain.QuizController) GetQuizes(context *gin.Context) {
	repo := repository.CreateQuizRepository() 
}
