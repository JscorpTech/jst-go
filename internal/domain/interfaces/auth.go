package interfaces

import (
	"github.com/JscorpTech/jst-go/internal/domain/dto"
	"github.com/JscorpTech/jst-go/internal/models"
)

type LoginUsecasePort interface {
	GetToken(user *models.User) (*dto.Token, error)
	Login(phone, password string) (*models.User, error)
}

type RegisterUsecasePort interface {
	CreateUser(phone, firstName, lastName, password string) (*models.User, error)
	CreateUserIfNotExist(phone, firstName, lastName, password string) (*models.User, error)
}
