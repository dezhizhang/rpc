package main

import (
	"fmt"
	"log"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {

	user := model.User{
		Name: "周华建",
		Age:  22,
	}
	err := driver.DB.Model(&model.User{}).Where("id = ?", 2).Updates(&user).Error
	if err != nil {
		log.Fatalf("更新失败%s", err)
	}
	fmt.Println("更新成功")

}
