package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"wordly/api/domain"
	"wordly/api/middleware"

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

func (repo *UserRepositoryImpl) Register(request domain.RegisterRequest) (*domain.AuthResponse, error) {

	errorEnv := godotenv.Load()
	if errorEnv != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		return nil, fmt.Errorf("error loading .env file")
	}

	db := repo.mysql
	rows, err := db.Query(
		"INSERT INTO user_data (email, username, password) VALUES (?, ?, ?)",
		request.Email, request.Username, request.Password,
	)

	if err != nil {
		return nil, err
	}

	// be careful deferring Queries if you are using transactions
	defer rows.Close()
	//create jwt token
	accessToken, atErr := middleware.CreateAccessToken(&domain.UserData{}, secret, 2)

	if atErr != nil {
		return nil, fmt.Errorf(atErr.Error())
	}

	return &domain.AuthResponse{
		AuthToken: accessToken,
	}, nil
}

func (repo *UserRepositoryImpl) Login() (*domain.AuthResponse, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		return nil, err
	}

	db := repo.mysql
	query, err := db.Query("SELECT * FROM test WHERE")

	if err != nil {
		return nil, err
	}

	var userData domain.UserData
	query.Scan(&userData)

	defer query.Close()

	accessToken, authErr := middleware.CreateAccessToken(&userData, secret, 2)

	if authErr != nil {
		return nil, authErr
	}

	return &domain.AuthResponse{
		AuthToken: accessToken,
	}, nil
}

func openSqlCon() *sql.DB {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	port := os.Getenv("PORT")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/wordlydb", user, pass, port)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
