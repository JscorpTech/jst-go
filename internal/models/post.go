package models

import "gorm.io/gorm"

type PostModel struct {
	gorm.Model
	Title string
	Desc  string
	Image string
}

func (p *PostModel) TableName() string {
	return "post"
}
