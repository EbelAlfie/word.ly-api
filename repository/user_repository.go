package repository

import (
	"wordly/api/domain"
)

type UserRepositoryImpl struct {
	//db apalah
}

func CreateUserRepo() domain.UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Register() {

}

func (repo *UserRepositoryImpl) Login() {

}
