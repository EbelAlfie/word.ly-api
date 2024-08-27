package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"wordly/api/middleware"
	route "wordly/api/route"
)

func main() {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("SECRET")

	server := gin.Default()

	gin.SetMode(gin.DebugMode)

	wordly := server.Group("/wordly")
	{
		wordly.Group("/user", route.UserRoute(wordly))
		wordly.Use(middleware.JwtAuthMiddleware(secret))
		wordly.Group("/quiz", route.QuizRoute(wordly))
	}

	server.Run(":3030")
}
