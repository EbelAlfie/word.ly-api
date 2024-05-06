package repository

import (
	domain "wordly/api/domain"
)

type QuizRepositoryImpl struct (
	//my sql
)

func CreateQuizRepository() domain.QuizRepository {
	return &QuizRepositoryImpl{}
}

func (repo *domain.QuizRepository) GetQuizes() {

}