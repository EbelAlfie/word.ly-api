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

type AuthResponse struct {
	AuthToken string `json:"authToken"`
}

type UserData struct {
	Uid      string `json:"userId"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	password string `json:"password"`
}
