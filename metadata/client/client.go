package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"xiaozhi/metadata/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	//md := metadata.Pairs("timestamp", time.Now().Format("timestampFormat"))
	md := metadata.New(map[string]string{
		"name":     "bobby",
		"password": "imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err1 := c.SayHello(ctx, &proto.HelloRequest{Name: "bobby"})
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(r.Message)
}
