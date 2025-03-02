package usecase

import (
	"github.com/JscorpTech/jst-go/models"
	"github.com/JscorpTech/jst-go/pkg/messages"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	DB *gorm.DB
}

func NewAuthUsecase(db *gorm.DB) *AuthUsecase {
	return &AuthUsecase{
		DB: db,
	}
}

func (au *AuthUsecase) IsAlready(phone string) bool {
	var count int64
	au.DB.Model(&models.UserModel{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func (au *AuthUsecase) CreateUser(phone string, firstName string, lastName string) (user *models.UserModel) {
	user = &models.UserModel{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
	}
	au.DB.Create(user)
	return
}

func (au *AuthUsecase) CreateUserIfNotExist(phone string, firstName string, lastName string) (*models.UserModel, map[string]string) {
	if au.IsAlready(phone) {
		return nil, map[string]string{
			"field":   "phone",
			"message": messages.UserAlreadyExist,
			"type":    "unique",
		}
	}
	return au.CreateUser(phone, firstName, lastName), nil
}
