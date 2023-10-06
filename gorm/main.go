package main

import (
	"fmt"
	"xiaozhi/gorm/driver"
	"xiaozhi/gorm/model"
)

func main() {
	var user model.User
	driver.DB.Preload("Languages").Find(&user)

	for _, language := range user.Languages {
		fmt.Println(language.Name)
	}
}
