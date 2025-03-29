package usecase

import (
	"errors"

	"github.com/JscorpTech/jst-go/internal/domain/dto"
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"github.com/JscorpTech/jst-go/pkg/utils"
)

type LoginUsecase struct {
	UserRepository interfaces.UserRepositoryPort
}

func NewLoginUsecase(userRepository interfaces.UserRepositoryPort) interfaces.LoginUsecasePort {
	return &LoginUsecase{
		UserRepository: userRepository,
	}
}

func (au *LoginUsecase) GetToken(user *models.User) (*dto.Token, error) {
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
	return &dto.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (au *LoginUsecase) Login(phone string, password string) (*models.User, error) {
	user, err := au.UserRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if utils.CheckPasswordHash(password, user.Password) {
		return user, nil
	}
	return nil, errors.New("invalid credentials")
}
