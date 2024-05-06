package main

import (
	"github.com/gin-gonic/gin"

	route "wordly/api/route"
)

func main() {
	server := gin.Default()

	wordly := server.Group("/wordly")
	{
		server.Group("/user", route.UserRoute(wordly))
		server.Group("/quiz", route.QuizRoute(wordly))
	}

	server.Run("6969")
}
