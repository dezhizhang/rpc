package main

import "fmt"

type User struct {
}

type System struct {
}

func main() {
	//fmt.Println(unsafe.Sizeof(int(0)))
	// fmt.Println(unsafe.Sizeof(uint(0)))

	//var user User
	//fmt.Println(unsafe.Sizeof(user))
	//var user User
	//fmt.Println(unsafe.Sizeof(user))
	//fmt.Printf("%p", &user)

	//var user User
	//fmt.Println(unsafe.Sizeof(user))
	//fmt.Printf("%p", &user)

	//var user User
	//fmt.Println(unsafe.Sizeof(user))
	//fmt.Printf("%p", &
	//m := map[string]struct{}{}
	//m["a"] = struct{}{}
	//m["b"] = struct{}{}
	//m["c"] = struct{}{}
	//fmt.Println(m)

	//fmt.Println(&System{})

	//fmt.Println(unsafe.Sizeof("刘德华"))
	//fmt.Println(unsafe.Sizeof("探迹科技张德志"))

	fmt.Println(len("探迹科技张德志"))
	fmt.Println(len("刘德华"))

}
