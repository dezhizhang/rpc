package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string `json:"name"`
}
