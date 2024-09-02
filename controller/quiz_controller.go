package controller

import (
	"fmt"
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
	var requestBody domain.QuizRequest
	contentErr := context.ShouldBind(&requestBody)
	if contentErr != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Body must not be empty"})
		return
	}

	reqErr := validateQuizRequest(*&requestBody)

	if reqErr != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: reqErr.Error()})
	}

	cont.repository.InsertQuiz(requestBody)

	context.JSON(http.StatusOK, domain.SuccessResponse{Message: "Success"})
}

func validateQuizRequest(request domain.QuizRequest) error {
	if request.Soal == "" {
		return fmt.Errorf("question must not be empty")
	}

	if request.Benar == "" {
		return fmt.Errorf("answer must not be empty")
	}

	if len(request.Jawaban) == 0 {
		return fmt.Errorf("choices must not be empty")
	}

	if request.Tips == "" {
		return fmt.Errorf("tips must not be empty")
	}
	return nil
}
