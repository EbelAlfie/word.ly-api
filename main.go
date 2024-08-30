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

	server.Use(middleware.CORSMiddleware())

	wordly := server.Group("/wordly")
	{
		public := wordly.Group("")

		route.UserRoute(public.Group("/user"))

		private := wordly.Group("")
		private.Use(middleware.JwtAuthMiddleware(secret))

		route.QuizRoute(private.Group("/quiz"))
	}

	if err := server.Run("localhost:3030"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
