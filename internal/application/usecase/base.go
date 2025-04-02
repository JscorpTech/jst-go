package usecase

import (
	"errors"

	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"gorm.io/gorm"
)

type baseUsecase[T interfaces.BaseRepositoryPort[U], U any] struct {
	repository T
}

func NewBaseUsecase[T interfaces.BaseRepositoryPort[U], U any](repository T) interfaces.BaseUsecasePort[T, U] {
	return &baseUsecase[T, U]{
		repository: repository,
	}
}

func (u *baseUsecase[T, U]) Create(payload any) (*U, error) {
	return nil, errors.New("not implement")
}

func (u *baseUsecase[T, U]) FindByID(id uint) (*U, error) {
	return u.repository.FindByID(id)
}

func (u *baseUsecase[T, U]) FindAll() ([]U, error) {
	return u.repository.FindAll()
}

func (u *baseUsecase[T, U]) Update(entity *U) error {
	return u.repository.Update(entity)
}

func (u *baseUsecase[T, U]) Delete(id uint) error {
	return u.repository.Delete(id)
}

func (u *baseUsecase[T, U]) Filter(model *U, query string, params ...any) *gorm.DB {
	return u.repository.Filter(model, query, params...)
}
