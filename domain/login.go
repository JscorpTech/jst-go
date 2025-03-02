package domain

import (
	"github.com/JscorpTech/jst-go/models"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	CheckPassword(username, password string) bool
	GetToken(user *models.UserModel) (*Token, error)
	Login(username, password string) (*models.UserModel, error)
}
