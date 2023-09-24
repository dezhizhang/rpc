package main

import "fmt"

type User[T int | bool] struct {
	Name string `json:"name"`
	Age  T      `json:"age"`
}

func main() {
	user := User[bool]{
		Name: "刘德华",
		Age:  false,
	}

	fmt.Println(user)
}
