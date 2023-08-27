package main

import (
	"fmt"
	"log"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {

	//var user model.User
	//tx := driver.DB.Where(&model.User{Username: "刘德华14"}).Find(&user)
	//if tx.Error != nil {
	//	log.Fatalf("获取失败%s", tx.Error)
	//}
	//fmt.Println(user.Username)

	var users []model.User
	tx := driver.DB.Where(map[string]interface{}{
		"username": "刘德华14",
	}).Find(&users)

	if tx.Error != nil {
		log.Fatalf("查询失败%s", tx.Error)
	}
	fmt.Println(users[0].Username)
}
