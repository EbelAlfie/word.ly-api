package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twitchyliquid64/golang-asm/obj"

	domain "wordly/api/domain"
)

type UserControllerImpl struct {
	repo domain.UserRepository
}

func CreateUserController() domain.UserController {
	return &UserControllerImpl{}
}

func (cont *UserControllerImpl) Register(c *gin.Context) {
	var requestBody domain.RegisterRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, {Message: err.Error()})
		return
	}

}

func (cont *UserControllerImpl) Login(c *gin.Context) {
	var requestBody domain.LoginRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, {Message: err.Error()})
		return
	}

}
