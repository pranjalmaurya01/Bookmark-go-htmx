package models

import "gorm.io/gorm"

type Website struct {
	gorm.Model
	Name        string
	Description string `gorm:"type:text"`
	Img         string `gorm:"type:text"`
	IsFav       bool   `gorm:"type:bool;default:false"`
	CreatedAt   int64  `gorm:"autoCreateTime"`
	Url         string `gorm:"type:text"`
	UserId      int64  `gorm:"foreignkey:ID;"`
}
