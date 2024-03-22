package repository

import {
	"wordly/api/domain"
}

type QuizRepoImpl struct {
	//db
}

func CreateQuizRepo() QuizRepoImpl {
	return QuizRepoImpl{}
}

func (domain.QuizRepository) CreateOrGetGameRoom() {
	
}
