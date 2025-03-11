package domain

import (
	"github.com/JscorpTech/jst-go/models"
)

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,len=12"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginUsecase interface {
	GetToken(user *models.User) (*Token, error)
	Login(phone, password string) (*models.User, error)
}
