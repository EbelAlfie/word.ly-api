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

func (repo *QuizRepositoryImpl) GetQuiz(quizType domain.QuizType) (*domain.QuizModel, error) {
	database := repo.db
	_, err := database.Query("SELECT * FROM quiz_table WHERE type is quizType")

	if err != nil {
		return nil, err
	}

	return &domain.QuizModel{}, nil
}

func (repo *QuizRepositoryImpl) UpdateQuiz() (*domain.QuizModel, error) {
	database := repo.db
	_, err := database.Query("UPDATE * FROM quiz_table")

	if err != nil {
		return nil, err
	}

	return &domain.QuizModel{}, nil
}

func (repo *QuizRepositoryImpl) InsertQuiz() error {
	database := repo.db
	_, err := database.Query("INSERT * FROM quiz_table")

	if err != nil {
		return err
	}

	return nil
}

func (repo *QuizRepositoryImpl) GetDetailByUserId(teacherId string) (*domain.QuizModel, error) {
	database := repo.db
	rows, err := database.Query("SELECT * FROM quiz_table WHERE TeacherId IS ?", teacherId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &domain.QuizModel{}, err
}
