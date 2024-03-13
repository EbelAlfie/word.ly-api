package domain

type UserRepository interface {
	Register()
	Login()
}

type UserController interface {
	Register()
	Login()
}
