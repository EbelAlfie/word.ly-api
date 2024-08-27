package middleware

import (
	"net/http"
	"strings"
	"wordly/api/domain"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		raw := strings.Split(authHeader, " ")
		//Create token
		if len(raw) == 2 {
			token := raw[1]
			isAuthorized, authError := CheckAuthorized(token, secret)
			if isAuthorized && authError == nil {
				userId, err := ExtractUID(token, secret)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Unauthorized"})
					ctx.Abort()
				} else {
					ctx.Set("x-user-id", userId)
					ctx.Next()
				}
			}
		}
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Unauthorized"})
		ctx.Abort()
	}
}
