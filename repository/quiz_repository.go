package repository

import (
	"database/sql"
	"fmt"

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

func (repo *QuizRepositoryImpl) InsertQuiz(teacherId string, request domain.QuizRequest) error {
	database := repo.db
	transaction, transactionError := database.Begin()

	if transactionError != nil {
		return transactionError
	}

	rows, choiceErr := transaction.Query(
		"INSERT INTO choice_table (First, Second, Third, Fourth) VALUES (?, ?, ?, ?)",
		request.Choices[0], request.Choices[1], request.Choices[2], request.Choices[3],
	)

	if choiceErr != nil {
		_ = transaction.Rollback()
		return choiceErr
	}

	var choiceData domain.ChoiceModel
	for rows.Next() {
		if rowErr := rows.Scan(&choiceData.ChoiceId); rowErr != nil {
			_ = transaction.Rollback()
			return rowErr
		}
		fmt.Printf("%s wefwefwf", choiceData.ChoiceId)
	}

	_, quizErr := transaction.Query(
		"INSERT INTO quiz_table (TeacherId, ChoiceId, Soal, CorrectAnswer, Hint, Score, Type) VALUES (?, ?, ?, ?, ?, ?, ?)",
		teacherId, choiceData.ChoiceId, request.Question, request.CorrectAnswer, request.Hint, request.Score, request.Type,
	)

	if quizErr != nil {
		_ = transaction.Rollback()
		return quizErr
	}

	if trxErr := transaction.Commit(); trxErr != nil {
		return trxErr
	}

	return nil
}

func (repo *QuizRepositoryImpl) GetQuizDetail(quizId string) (*domain.QuizModel, error) {
	database := repo.db
	rows, err := database.Query(
		"SELECT * FROM quiz_table WHERE QuizId = ? ",
		quizId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &domain.QuizModel{}, err
}

func (repo *QuizRepositoryImpl) GetQuizesByUserId(teacherId string) (*domain.QuizModel, error) {
	database := repo.db
	rows, err := database.Query("SELECT * FROM quiz_table WHERE TeacherId IS ?", teacherId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &domain.QuizModel{}, err
}
