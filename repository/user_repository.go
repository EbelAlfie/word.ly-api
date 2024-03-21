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

func (repo *UserRepositoryImpl) Register(request domain.RegisterRequest) {
	//name not taken return ok
}

func (repo *UserRepositoryImpl) Login(request domain.LoginRequest) {
	//name yes return ok
}
