package domain

import (
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Register(RegisterRequest) (*AuthResponse, error)
	Login() (*UserData, error)
}

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
