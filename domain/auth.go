package domain

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthUsecase interface {
	CheckPassword(username, password string) bool
	CreateUser(username, password string) error
	GetToken(username, password string) (string, error)
}
