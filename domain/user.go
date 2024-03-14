package domain

import (
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Register()
	Login()
}

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
