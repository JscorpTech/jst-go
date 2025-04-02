package interfaces

import "gorm.io/gorm"

type BaseRepositoryPort[T any] interface {
	Create(entity *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(entity *T) error
	Delete(id uint) error
	Filter(model *T, query string, params ...any) *gorm.DB
}

type BaseUsecasePort[T BaseRepositoryPort[U], U any] interface {
	Create(payload any) (*U, error)
	FindByID(id uint) (*U, error)
	FindAll() ([]U, error)
	Update(entity *U) error
	Delete(id uint) error
	Filter(model *U, query string, params ...any) *gorm.DB
}
