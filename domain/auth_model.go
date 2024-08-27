package domain

type LoginRequest struct {
	Username string
	Password string
}

type RegisterRequest struct {
	Email    string
	Username string
	Password string
}

type AuthResponse struct {
	AuthToken string `json:"authToken"`
}

type UserData struct {
	Uid      string `json:"userId"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

func (data *UserData) ValidateField() bool {

}
