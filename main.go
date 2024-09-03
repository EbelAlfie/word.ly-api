package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"wordly/api/middleware"
	route "wordly/api/route"
)

func main() {
	server := gin.Default()

	gin.SetMode(gin.DebugMode)

	server.Use(middleware.CORSMiddleware())

	wordly := server.Group("/wordly")
	{
		route.UserRoute(wordly.Group("/user"))
		route.QuizRoute(wordly.Group("/quiz"))
	}

	if err := server.Run("localhost:3030"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
