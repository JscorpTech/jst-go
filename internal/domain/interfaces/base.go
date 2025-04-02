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

type BaseUsecase[T BaseRepository[U], U any] interface {
	Create(entity *U) error
	FindByID(id uint) (*U, error)
	FindAll() ([]U, error)
	Update(entity *U) error
	Delete(id uint) error
	Filter(model *U, query string, params ...any) *gorm.DB
}
