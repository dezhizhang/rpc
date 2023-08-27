package main

import (
	"fmt"
	"log"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {

	var user model.User
	driver.DB.Where(&model.User{Username: "刘德华0"}).Find(&user)

	user.Username = "周华建"
	err := driver.DB.Save(&user).Error
	if err != nil {
		log.Fatalf("更新失败%s", err)
	}
	fmt.Println("更新成功")

}
