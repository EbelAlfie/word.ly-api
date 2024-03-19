package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	domain "wordly/api/domain"
)

type UserControllerImpl struct {
	repo domain.UserRepository
}

func CreateUserController(repo domain.UserRepositoryImpl) domain.UserController {
	return &UserControllerImpl{
		repo
	}
}

func (cont *UserControllerImpl) Register(c *gin.Context) {
	var requestBody domain.RegisterRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, {Message: err.Error()})
		return
	}

	userRepository := domain.CreateUserRepo() 

	if err:= userRepository.Register() != nil {
		c.JSON(http.StatusBadRequest, {Message: err.Error()})
		return
	}

	c.JSON(http.StatusOk, {Message: "Success"})
}

func (cont *UserControllerImpl) Login(c *gin.Context) {
	var requestBody domain.LoginRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, {Message: err.Error()})
		return
	}

	userRepository := domain.CreateUserRepo() 
	if err := userRepository.Login() != nil {
		c.JSON(http.StatusNotFound, {Message: err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, {Token: "erewfwffwerweddwefewr"})
}
