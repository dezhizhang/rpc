package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
	"xiaozhi/timeout/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "刘德华"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
