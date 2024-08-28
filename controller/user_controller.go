package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	domain "wordly/api/domain"
)

type UserControllerImpl struct {
	repository domain.UserRepository
}

func CreateUserController(repo domain.UserRepository) domain.UserController {
	return &UserControllerImpl{
		repository: repo,
	}
}

func (cont *UserControllerImpl) Register(c *gin.Context) {
	var requestBody domain.RegisterRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fieldErr := validateRegistration(requestBody)

	if fieldErr != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: fieldErr.Error()})
		return
	}

	userRepository := cont.repository

	authResponse, errorRepo := userRepository.Register(requestBody)
	if errorRepo != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: errorRepo.Error()})
		return
	}

	c.JSON(http.StatusOK, *authResponse)
}

func (cont *UserControllerImpl) Login(c *gin.Context) {
	var requestBody domain.LoginRequest
	contentErr := c.ShouldBind(&requestBody)
	if contentErr != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: contentErr.Error()})
		return
	}

	fieldErr := validateLoginReq(requestBody)

	if fieldErr != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: fieldErr.Error()})
		return
	}

	userRepository := cont.repository
	userToken, err := userRepository.Login()
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, userToken)
}

func (cont *UserControllerImpl) GetProfile(c *gin.Context) {

}

func validateRegistration(request domain.RegisterRequest) error {
	if request.Email == "" {
		return fmt.Errorf("Email must not be empty")
	}

	if request.Username == "" {
		return fmt.Errorf("Username must not be empty")
	}

	if request.Password == "" {
		return fmt.Errorf("Password is required")
	}
	return nil
}

func validateLoginReq(request domain.LoginRequest) error {
	if request.Username == "" {
		return fmt.Errorf("Username is empty")
	}

	if request.Password == "" {
		return fmt.Errorf("Password is required")
	}
	return nil
}
