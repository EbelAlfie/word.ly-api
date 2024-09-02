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

func (repo *QuizRepositoryImpl) InsertQuiz(request domain.QuizRequest) error {
	database := repo.db
	rows, err := database.Query(
		"INSERT INTO choice_table (First, Second, Thrid, Fourth) VALUES (?, ?, ?, ?)",
		request.Jawaban[0], request.Jawaban[1], request.Jawaban[2], request.Jawaban[3],
	)

	ch := make(chan domain.ChoiceModel)
	for rows.Next() {
		var quizData domain.ChoiceModel

		if err := rows.Scan(&quizData.ChoiceId); err != nil {
			// do something with error
		} else {
			println(quizData.ChoiceId)
			ch <- quizData
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *QuizRepositoryImpl) GetQuizDetail(quizId string) (*domain.QuizModel, error) {
	database := repo.db
	rows, err := database.Query("SELECT * FROM quiz_table WHERE QuizId IS ?", quizId)

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
