package domain

import "github.com/JscorpTech/jst-go/models"

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}

type RegisterUsecase interface {
	CreateUser(phone, firstName, lastName, password string) (*models.UserModel, error)
	CreateUserIfNotExist(phone, firstName, lastName, password string) (*models.UserModel, error)
}
