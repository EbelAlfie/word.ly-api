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
		server.Group("/user", route.UserRoute(wordly))
		server.Group("/quiz", route.QuizRoute(wordly))
	}

	server.Run(":8080")
}
