package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestStrings(t *testing.T) {
	//s := "刘德华 hello"
	//sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//fmt.Println(sh.Len)
	s := "刘德华 hello"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sh.Len)
}

func TestStringPrint(t *testing.T) {
	//s := "刘德华 hello"
	//
	//for _, value := range s {
	//	fmt.Printf("%c\n", value)
	//}
	s := "刘德华 hello world"
	for _, value := range s {
		fmt.Printf("%c\n", value)
	}
}

func TestSlicePrint(t *testing.T) {
	//slice := []int{1, 2, 3}
	//fmt.Println(slice)
	//slice := make([]int, 10)
	//
	//fmt.Println(slice)

	slice := make([]int, 10)
	fmt.Println(slice)
}
