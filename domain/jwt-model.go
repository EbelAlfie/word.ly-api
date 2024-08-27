package domain

import (
	jwt "github.com/golang-jwt/jwt"
)

type JwtModel struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}
