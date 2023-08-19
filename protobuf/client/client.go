package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"xiaozhi/protobuf/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	rsp, _ := client.SayHello(context.Background(), &proto.HelloRequest{
		Name:   "刘德华",
		Url:    "http://www.xiaozhi.shop",
		Gender: proto.Gender_FEMALE,
		Mp: map[string]string{
			"name": "周华建",
			"age":  "222",
		},
		CreateTime: timestamppb.New(time.Now()),
	})
	fmt.Println(rsp.Message)
}
