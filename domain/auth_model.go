package domain

type LoginRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}

type RegisterRequest struct {
	email    string `json:"email"`
	username string `json:"username"`
	password string `json:"password"`
}

type UserData struct {
	uid      string `json:"userId"`
	email    string `json:"email"`
	username string `json:"username"`
	password string `json:"password"`
}
