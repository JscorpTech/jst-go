package usecase

import (
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
)

type LoginUsecase struct {
	UserRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) domain.LoginUsecase {
	return &LoginUsecase{
		UserRepository: userRepository,
	}
}

func (au *LoginUsecase) CheckPassword(username, password string) bool {
	return false
}

func (au *LoginUsecase) GetToken(user *models.UserModel) (*domain.Token, error) {
	return &domain.Token{
		AccessToken:  "111",
		RefreshToken: "222",
	}, nil
}

func (au *LoginUsecase) Login(username string, password string) (*models.UserModel, error) {
	return &models.UserModel{}, nil
}
