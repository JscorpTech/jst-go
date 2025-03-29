package interfaces

import "github.com/JscorpTech/jst-go/internal/models"

type UserRepositoryPort interface {
	IsAlready(phone string) bool
	Create(firstName, LastName, phone, password string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
}
