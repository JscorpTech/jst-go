package repository

import (
	"github.com/JscorpTech/jst-go/internal/domain/interfaces"
	"github.com/JscorpTech/jst-go/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
	BaseRepository[models.User]
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryPort {
	return &UserRepository{
		DB:             db,
		BaseRepository: NewBaseRepository[models.User](db),
	}
}

func (u *UserRepository) Create(phone, firstName, lastName, password string) (*models.User, error) {
	user := &models.User{
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

func (u *UserRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	if err := u.DB.First(&user, "phone = ?", phone).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) IsAlready(phone string) bool {
	var count int64
	u.DB.Model(&models.User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
