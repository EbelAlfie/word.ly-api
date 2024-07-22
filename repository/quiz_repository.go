package repository

import (
	"database/sql"

	domain "wordly/api/domain"
)

type QuizRepositoryImpl struct {
	//my sql
	db *sql.DB
}

func CreateQuizRepository() domain.QuizRepository {
	database := openSqlCon()
	return &QuizRepositoryImpl{
		db: database,
	}
}

func (repo *QuizRepositoryImpl) GetQuizes() {

}
