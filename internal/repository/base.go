package repository

import "gorm.io/gorm"

type BaseRepository[T any] interface {
	Create(entity *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(entity *T) error
	Delete(id uint) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

// Yangi repository yaratish
func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

// Create
func (r *baseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

// FindByID
func (r *baseRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll
func (r *baseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Update
func (r *baseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// Delete
func (r *baseRepository[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}
