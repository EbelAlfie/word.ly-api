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

func (repo *QuizRepositoryImpl) GetQuiz(quizType domain.QuizType) ([]domain.QuizModel, error) {
	database := repo.db

	rows, quizError := database.Query(
		`SELECT QuizId, Soal, CorrectAnswer, Hint, Score, Type, First, Second, Third, Fourth FROM quiz_table 
		INNER JOIN choice_table ON quiz_table.ChoiceId = choice_table.ChoiceId WHERE Type = ?
		ORDER BY RAND() LIMIT 6`,
		quizType,
	)

	if quizError != nil {
		return nil, quizError
	}

	var quizes []domain.QuizModel
	for rows.Next() {
		var quiz domain.QuizModel
		var choices domain.ChoiceModel
		if scanErr := rows.Scan(&quiz.Id, &quiz.Question, &quiz.CorrectAnswer, &quiz.Hint, &quiz.Score, &quiz.Type, &choices.ChoiceA, &choices.ChoiceB, &choices.ChoiceC, &choices.ChoiceD); scanErr != nil {
			return nil, scanErr
		}

		quiz.Choices = []string{choices.ChoiceA, choices.ChoiceB, choices.ChoiceC, choices.ChoiceD}
		quizes = append(quizes, quiz)
	}

	return quizes, nil
}

func (repo *QuizRepositoryImpl) UpdateQuiz(request domain.QuizRequest) error {
	database := repo.db
	transaction, trxErr := database.Begin()

	if trxErr != nil {
		return trxErr
	}

	if _, quizUpdateErr := transaction.Exec(
		`UPDATE quiz_table 
		SET Soal = ?, CorrectAnswer = ?, Hint = ? 
		WHERE QuizId = ?`,
		request.Question, request.CorrectAnswer, request.Hint, request.QuizId,
	); quizUpdateErr != nil {
		transaction.Rollback()
		return quizUpdateErr
	}

	if _, choiceUpdateErr := transaction.Exec(
		`UPDATE choice_table
		SET First = ?, Second = ?, Third = ?, Fourth = ?
		WHERE ChoiceId = ?`,
		request.Choices[0], request.Choices[1], request.Choices[2], request.Choices[3], request.ChoiceId,
	); choiceUpdateErr != nil {
		transaction.Rollback()
		return choiceUpdateErr
	}

	if trxErr := transaction.Commit(); trxErr != nil {
		transaction.Rollback()
		return trxErr
	}

	return nil
}

func (repo *QuizRepositoryImpl) InsertQuiz(teacherId string, request domain.QuizRequest) error {
	database := repo.db
	transaction, transactionError := database.Begin()

	if transactionError != nil {
		return transactionError
	}

	result, choiceErr := transaction.Exec(
		"INSERT INTO choice_table (First, Second, Third, Fourth) VALUES (?, ?, ?, ?)",
		request.Choices[0], request.Choices[1], request.Choices[2], request.Choices[3],
	)

	if choiceErr != nil {
		_ = transaction.Rollback()
		return choiceErr
	}

	id, idErr := result.LastInsertId()

	if idErr != nil {
		_ = transaction.Rollback()
		return idErr
	}

	if _, quizErr := transaction.Exec(
		"INSERT INTO quiz_table (TeacherId, ChoiceId, Soal, CorrectAnswer, Hint, Score, Type) VALUES (?, ?, ?, ?, ?, ?, ?)",
		teacherId, id, request.Question, request.CorrectAnswer, request.Hint, request.Score, request.Type,
	); quizErr != nil {
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
	transaction, trxErr := database.Begin()

	if trxErr != nil {
		return nil, trxErr
	}

	row := transaction.QueryRow(
		"SELECT * FROM quiz_table WHERE QuizId = ? ",
		quizId,
	)

	var quiz domain.QuizModel

	var choiceId string
	if scanErr := row.Scan(&quiz.Id, &quiz.TId, &choiceId, &quiz.Question, &quiz.CorrectAnswer, &quiz.Hint, &quiz.Score, &quiz.Type); scanErr != nil {
		transaction.Rollback()
		return nil, scanErr
	}

	choice := transaction.QueryRow(
		"SELECT * FROM choice_table WHERE ChoiceId = ?",
		choiceId,
	)

	var choices domain.ChoiceModel
	if scanErr := choice.Scan(&choices.ChoiceId, &choices.ChoiceA, &choices.ChoiceB, &choices.ChoiceC, &choices.ChoiceD); scanErr != nil {
		transaction.Rollback()
		return nil, scanErr
	}

	return &domain.QuizModel{
		Id:            quiz.Id,
		TId:           quiz.TId,
		Question:      quiz.Question,
		ChoiceId:      choices.ChoiceId,
		Choices:       []string{choices.ChoiceA, choices.ChoiceB, choices.ChoiceC, choices.ChoiceD},
		CorrectAnswer: quiz.CorrectAnswer,
		Hint:          quiz.Hint,
		Score:         quiz.Score,
		Type:          quiz.Type,
	}, nil
}

func (repo *QuizRepositoryImpl) GetQuizesByUserId(teacherId string) ([]domain.QuizModel, error) {
	database := repo.db
	rows, err := database.Query("SELECT * FROM quiz_table WHERE TeacherId = ?", teacherId)

	if err != nil {
		return nil, err
	}

	var quizList []domain.QuizModel

	for rows.Next() {
		var quiz domain.QuizModel
		var trash string
		quizErr := rows.Scan(&quiz.Id, &quiz.TId, &trash, &quiz.Question, &quiz.CorrectAnswer, &quiz.Hint, &quiz.Score, &quiz.Type)
		if quizErr != nil {
			return nil, quizErr
		} else {
			quizList = append(quizList, quiz)
		}
	}

	defer rows.Close()

	return quizList, err
}
