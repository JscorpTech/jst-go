package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string  `gorm:"not null"`
	LastName  string  `gorm:"type:varchar(255)"`
	Phone     string  `gorm:"type:varchar(255)"`
	Password  string  `gorm:"type:varchar(255)"`
	Tokens    []Token `gorm:"constraint:OnDelete:CASCADE"`
}

type Token struct {
	gorm.Model
	UserID uint
	User   User
}

func (t *Token) TableName() string {
	return "tokens"
}

func (u *User) TableName() string {
	return "users"
}
