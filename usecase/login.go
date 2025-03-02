package usecase

import (
	"errors"
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
	"github.com/JscorpTech/jst-go/pkg/utils"
)

type LoginUsecase struct {
	UserRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) domain.LoginUsecase {
	return &LoginUsecase{
		UserRepository: userRepository,
	}
}

func (au *LoginUsecase) GetToken(user *models.UserModel) (*domain.Token, error) {
	accessToken, err := utils.GenerateJwt(&utils.Jwt{
		Sub:  int(user.ID),
		Type: "access",
	})
	if err != nil {
		return nil, err
	}
	refreshToken, err := utils.GenerateJwt(&utils.Jwt{
		Sub:  int(user.ID),
		Type: "refresh",
	})
	if err != nil {
		return nil, err
	}
	return &domain.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (au *LoginUsecase) Login(phone string, password string) (*models.UserModel, error) {
	user, err := au.UserRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if utils.CheckPasswordHash(password, user.Password) {
		return user, nil
	}
	return nil, errors.New("invalid credentials")
}
