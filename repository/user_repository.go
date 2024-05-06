package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"wordly/api/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type UserRepositoryImpl struct {
	//db apalah
	mysql *sql.DB
}

func CreateUserRepo() domain.UserRepository {
	db := openSqlCon()
	return &UserRepositoryImpl{
		mysql: db,
	}
}

func (repo *UserRepositoryImpl) Register(request domain.RegisterRequest) error {
	// perform a db.Query insert
	db := repo.mysql
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return nil
}

func (repo *UserRepositoryImpl) Login() (*domain.UserData, error) {
	db := repo.mysql
	query, err := db.Query("SELECT * FROM test WHERE")

	if err != nil {
		return nil, err
	}

	defer query.Close()
	return &domain.UserData{}, nil
}

func openSqlCon() *sql.DB {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	port := os.Getenv("PORT")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/user_data", user, pass, port)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
