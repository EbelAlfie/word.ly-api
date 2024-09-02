package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func openSqlCon() *sql.DB {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("USR")
	pass := os.Getenv("PASS")
	port := os.Getenv("PORT")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/wordlydb", user, pass, port)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
