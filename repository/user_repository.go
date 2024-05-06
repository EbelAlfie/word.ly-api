package repository

import (
	"database/sql"
	"wordly/api/domain"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	return db
}
