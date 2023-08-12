package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//type HelloService struct {
//}
//
//func (s *HelloService) Hello(request string, reply *string) error {
//	*reply = "hello" + request
//	return nil
//}
//
//func main() {
//	listener, _ := net.Listen("tcp", ":8080")
//	_ = rpc.RegisterName("HelloService", &HelloService{})
//
//	for {
//		conn, _ := listener.Accept()
//		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
//	}
//
//}

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello\t" + request
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		log.Printf("注册失败%s", err)
	}
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
