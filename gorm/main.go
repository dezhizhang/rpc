package main

import (
	"fmt"
	"log"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {

	company := model.Company{
		Name: "晓智科技",
	}

	user := model.User{
		Username: "晓智",
		Company:  company,
	}

	err := driver.DB.Create(&user).Error

	if err != nil {
		log.Printf("插入失败%s", err)
	}

	fmt.Println("插入成功")

}
