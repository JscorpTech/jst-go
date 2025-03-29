package usecase

import (
	"errors"

	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"github.com/JscorpTech/jst-go/pkg/utils"
)

type RegisterUsecase struct {
	UserRepository interfaces.UserRepositoryPort
}

func NewRegisterUsecase(userRepository interfaces.UserRepositoryPort) interfaces.RegisterUsecasePort {
	return &RegisterUsecase{
		UserRepository: userRepository,
	}
}

func (r *RegisterUsecase) CreateUser(phone, firstName, lastName, password string) (*models.User, error) {
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := r.UserRepository.Create(phone, firstName, lastName, passwordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *RegisterUsecase) CreateUserIfNotExist(phone, firstName, lastName, password string) (*models.User, error) {
	if r.UserRepository.IsAlready(phone) {
		return nil, errors.New("user already exists")
	}
	user, err := r.CreateUser(phone, firstName, lastName, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
