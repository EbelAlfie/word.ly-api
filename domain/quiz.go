package domain

import {
	"github.com/gin-gonic/gin"
}


type QuizRepository interface {
	func CreateOrGetGameRoom()
}

type QuizController interface {
	func CreateOrGetGameRoom(c *gin.Context)
}