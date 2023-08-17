package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
	"xiaozhi/streamrpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	// 客户端流模式
	c := proto.NewGreeterClient(conn)
	stream, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "刘德华"})
	for {
		a, err1 := stream.Recv()
		if err1 != nil {
			panic(err1)
		}
		fmt.Println(a)
	}
	// 客户端流模式
	putStream, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		putStream.Send(&proto.StreamReqData{Data: "晓智云"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}
