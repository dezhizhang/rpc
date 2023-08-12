package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//func main() {
//	conn, err := net.Dial("tcp", "localhost:8080")
//	if err != nil {
//		panic("连接失败")
//	}
//	var reply string
//	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
//	err = client.Call("HelloService.Hello", "bobby", &reply)
//	if err != nil {
//		panic("调用失败")
//	}
//	fmt.Println(reply)
//}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("连接失败")
	}
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
