package main

import (
	"fmt"
	"unsafe"
)

type User struct {
}

type

func main() {
	//fmt.Println(unsafe.Sizeof(int(0)))
	// fmt.Println(unsafe.Sizeof(uint(0)))

	//var user User
	//fmt.Println(unsafe.Sizeof(user))
	//var user User
	//fmt.Println(unsafe.Sizeof(user))
	//fmt.Printf("%p", &user)

	var user User
	fmt.Println(unsafe.Sizeof(user))
	fmt.Printf("%p", &user)
}
