package interfaces

import "gorm.io/gorm"

type BaseRepository[T any] interface {
	Create(entity *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(entity *T) error
	Delete(id uint) error
	Filter(model *T, query string, params ...any) *gorm.DB
}

type BaseUsecase[T any] interface {
	Create(entity *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(entity *T) error
	Delete(id uint) error
	Filter(model *T, query string, params ...any) *gorm.DB
}
