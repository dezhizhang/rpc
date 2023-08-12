package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	//实例化
	listen, err := net.Listen("tcp", "localhost:8084")
	if err != nil {
		log.Printf("实例化失败%s", err)
	}
	//2 注册处理逻辑handler
	err2 := rpc.RegisterName("HelloService", &HelloService{})
	if err2 != nil {
		log.Printf("注册失败%s", err)
	}
	// 启动服务
	conn, err1 := listen.Accept()
	if err1 != nil {
		log.Printf("启动服务失败%s", err1)
	}
	rpc.ServeConn(conn)
}
