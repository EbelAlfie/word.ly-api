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

	userRepository := cont.repository

	errorRepo := userRepository.Register(requestBody)
	if errorRepo != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Success"})
}

func (cont *UserControllerImpl) Login(c *gin.Context) {
	var requestBody domain.LoginRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		fmt.Print("HELLO")
		return
	}

	userRepository := cont.repository
	userData, err := userRepository.Login()
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if userData == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "You must be registered"})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Message: "token"})
	}
}
