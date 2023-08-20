package main

import (
	"fmt"
	"log"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {

	var user model.User
	tx := driver.DB.Model(&user).Where("username = ?", "刘德华20").Find(&user)
	if tx.Error != nil {
		log.Fatalf("获取失败%s", tx.Error)
	}
	fmt.Println(user)

}
