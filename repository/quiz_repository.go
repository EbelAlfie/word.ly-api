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

func (repo *QuizRepositoryImpl) GetCerpen() (*domain.QuizModel, error) {
	database := repo.db
	_, err := database.Query("SELECT * FROM test WHERE type is quizType")

	if err != nil {
		return nil, err
	}

	return &domain.QuizModel{}, nil
}

func (repo *QuizRepositoryImpl) GetKalimatEfektif() (*domain.QuizModel, error) {
	database := repo.db
	_, err := db.Query("SELECT * FROM quiz_table")

	if err != nil {
		return nil, err
	}

}
