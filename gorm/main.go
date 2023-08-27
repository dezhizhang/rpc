package main

import (
	"log"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {

	user := model.User{
		Username: "周华建",
		Company: model.Company{
			Name: "晓智科技",
			ID:   1,
		},
	}

	err := driver.DB.Create(&user).Error
	if err != nil {
		log.Fatalf("设置失败%s", err)
	}

}
