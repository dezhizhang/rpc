package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"column:username;type:varchar(50)"`
	Age       int    `json:"age"`
	CompanyId int    `json:"companyId"`
	Company   Company
}

type Company struct {
	ID   int    `gorm:"colum:id;type:int" json:"ID"`
	Name string `gorm:"colum:name;type:varchar(50)"json:"Name"`
}
