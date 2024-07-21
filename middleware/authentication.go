package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string, c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	token := strings.Split(authHeader, " ")
	//Create token
}
