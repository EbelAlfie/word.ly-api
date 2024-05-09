package main

import (
	"github.com/gin-gonic/gin"

	route "wordly/api/route"
)

func main() {
	server := gin.Default()

	gin.SetMode(gin.DebugMode)

	wordly := server.Group("/wordly")
	{
		wordly.Group("/user", route.UserRoute(wordly))
		wordly.Group("/quiz", route.QuizRoute(wordly))
	}

	server.Run(":3030")
}
