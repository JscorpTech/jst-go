package usecase

import (
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"gorm.io/gorm"
)

type baseUsecase[T any] struct {
	repository interfaces.BaseRepository[T]
}

func NewBaseUsecase[T any](repository interfaces.BaseRepository[T]) interfaces.BaseUsecase[T] {
	return &baseUsecase[T]{
		repository: repository,
	}
}

func (u *baseUsecase[T]) Create(entity *T) error {
	return u.repository.Create(entity)
}

func (u *baseUsecase[T]) FindByID(id uint) (*T, error) {
	return u.repository.FindByID(id)
}

func (u *baseUsecase[T]) FindAll() ([]T, error) {
	return u.repository.FindAll()
}

func (u *baseUsecase[T]) Update(entity *T) error {
	return u.repository.Update(entity)
}

func (u *baseUsecase[T]) Delete(id uint) error {
	return u.repository.Delete(id)
}

func (u *baseUsecase[T]) Filter(model *T, query string, params ...any) *gorm.DB {
	return u.repository.Filter(model, query, params...)
}
