package main

import (
	"github.com/gin-gonic/gin"

	user "wordly/api/route"
)

func main() {
	server := gin.Default()

	wordly := server.Group("/wordly")
	{
		server.Group("/user", user.UserRoute(wordly))
	}

	server.Run("6969")
}
