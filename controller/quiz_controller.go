package controller

import (
	"net/http"

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

func (cont *QuizControllerImpl) GetQuiz(context *gin.Context) {
	requestParam := context.Params.ByName("quizType")
	quizType := domain.ParseToEnum(requestParam)

	quizes, quizError := cont.repository.GetQuiz(quizType)
	if quizError != nil {
		context.JSON(http.StatusNotFound, domain.ErrorResponse{Message: quizError.Error()})
		return
	}

	context.JSON(http.StatusOK, quizes)
}

func (cont *QuizControllerImpl) GetQuizesByUserId(context *gin.Context) {
	requestParam := context.Params.ByName("userId")
	if requestParam == "" {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User Id is required"})
	}

	quiz, quizError := cont.repository.GetQuizesByUserId(requestParam)
	if quizError != nil {
		context.JSON(http.StatusNotFound, domain.ErrorResponse{Message: quizError.Error()})
		return
	}

	context.JSON(http.StatusOK, quiz)
}

func (cont *QuizControllerImpl) GetQuizDetail(context *gin.Context) {
	requestParam := context.Params.ByName("quizId")
	if requestParam == "" {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Quiz Id is required"})
	}

	quiz, quizError := cont.repository.GetQuizesByUserId(requestParam)
	if quizError != nil {
		context.JSON(http.StatusNotFound, domain.ErrorResponse{Message: quizError.Error()})
		return
	}

	context.JSON(http.StatusOK, quiz)
}

func (cont *QuizControllerImpl) UpdateQuiz(context *gin.Context) {

}

func (cont *QuizControllerImpl) InsertQuiz(context *gin.Context) {

}
