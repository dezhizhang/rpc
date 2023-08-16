package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("连接失败")
	}

	var reply string
	//client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
