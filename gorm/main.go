package main

import (
	"fmt"
	"log"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {
	company := model.Company{}

	user := model.User{
		Company: company,
	}

	err := driver.DB.Preload("Company").Model(&model.User{}).Where("id=?", 4).Find(&user).Error

	if err != nil {
		log.Printf("查询失败%s", err.Error())
	}

	fmt.Println(user.Company.Name)
}
