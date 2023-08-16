package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"xiaozhi/grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err1 := c.SayHello(context.Background(), &proto.HelloRequest{Name: "刘德华"})
	if err1 != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}
