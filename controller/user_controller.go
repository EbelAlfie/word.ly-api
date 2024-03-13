package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	domain "wordly/api/domain"
)

type UserControllerImpl struct {
	repo domain.UserRepository
}

func (cont *UserControllerImpl) Login(c *gin.Context) {
	var requestBody domain.AuthModel
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

}
