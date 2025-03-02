package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(255)"`
	Password  string `gorm:"type:varchar(255)"`
}

func (u *UserModel) TableName() string {
	return "users"
}
