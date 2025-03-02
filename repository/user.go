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
	user := &models.UserModel{
		Phone:     phone,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}
	if err := u.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) FindByPhone(phone string) (*models.UserModel, error) {
	var user models.UserModel
	if err := u.DB.First(&user, "phone = ?", phone).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindById(id int) (*models.UserModel, error) {
	var user models.UserModel
	if err := u.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) IsAlready(phone string) bool {
	var count int64
	u.DB.Model(&models.UserModel{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
