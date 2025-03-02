package repository

import (
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(phone, firstName, lastName, password string) (*models.UserModel, error) {
	return &models.UserModel{}, nil
}

func (u *UserRepository) IsAlready(phone string) bool {
	var count int64
	u.DB.Model(&models.UserModel{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
