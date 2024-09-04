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
	requestParam := context.Request.URL.Query().Get("quizType")
	quizType := domain.ParseToEnum(requestParam)

	quizes, quizError := cont.repository.GetQuiz(quizType)
	if quizError != nil {
		context.JSON(http.StatusNotFound, domain.ErrorResponse{Message: quizError.Error()})
		return
	}

	context.JSON(http.StatusOK, quizes)
}

func (cont *QuizControllerImpl) GetQuizesByUserId(context *gin.Context) {
	teacherId := context.GetString("x-user-id")

	quiz, quizError := cont.repository.GetQuizesByUserId(teacherId)
	if quizError != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: quizError.Error()})
		return
	}

	context.JSON(http.StatusOK, quiz)
}

func (cont *QuizControllerImpl) GetQuizDetail(context *gin.Context) {
	requestParam := context.Request.URL.Query().Get("quizId")
	if requestParam == "" {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Quiz Id is required"})
		return
	}

	quiz, quizError := cont.repository.GetQuizDetail(requestParam)
	if quizError != nil {
		context.JSON(http.StatusNotFound, domain.ErrorResponse{Message: quizError.Error()})
		return
	}

	context.JSON(http.StatusOK, quiz)
}

func (cont *QuizControllerImpl) UpdateQuiz(context *gin.Context) {
	var requestBody domain.QuizRequest
	contentErr := context.ShouldBind(&requestBody)

	if contentErr != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Body must not be empty"})
		return
	}

	reqErr := validateQuizRequest(requestBody)

	if reqErr != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: reqErr.Error()})
		return
	}

	reqErr = validateForUpdate(requestBody)

	if reqErr != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: reqErr.Error()})
		return
	}

	err := cont.repository.UpdateQuiz(requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, domain.SuccessResponse{Message: "Success"})
}

func (cont *QuizControllerImpl) InsertQuiz(context *gin.Context) {
	var requestBody domain.QuizRequest
	contentErr := context.ShouldBind(&requestBody)

	if contentErr != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Body must not be empty"})
		return
	}

	reqErr := validateQuizRequest(requestBody)

	if reqErr != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: reqErr.Error()})
		return
	}

	teacherID := context.GetString("x-user-id")

	err := cont.repository.InsertQuiz(teacherID, requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, domain.SuccessResponse{Message: "Success"})
}

func validateQuizRequest(request domain.QuizRequest) error {
	if request.Question == "" {
		return fmt.Errorf("question must not be empty")
	}

	if request.CorrectAnswer == "" {
		return fmt.Errorf("answer must not be empty")
	}

	if len(request.Choices) == 0 {
		return fmt.Errorf("choices must not be empty")
	}

	if request.Hint == "" {
		return fmt.Errorf("tips must not be empty")
	}
	return nil
}

func validateForUpdate(request domain.QuizRequest) error {
	if request.QuizId == "" {
		return fmt.Errorf("QuizId is empty")
	}
	if request.ChoiceId == "" {
		return fmt.Errorf("ChoiceId is empty")
	}
	return nil
}
