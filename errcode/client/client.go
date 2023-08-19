package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"xiaozhi/errcode/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "刘德华"})
	if err != nil {
		//panic(err)
		st, ok := status.FromError(err)
		if !ok {
			panic("角析error失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())

	}

	fmt.Println(r.Message)
}
