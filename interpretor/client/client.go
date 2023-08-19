package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
	"xiaozhi/interpretor/proto"
)

type customCredential struct {
}

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoke grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		md := metadata.New(map[string]string{
			"appId":  "101010",
			"appKey": "i am key",
		})

		ctx = metadata.NewOutgoingContext(context.Background(), md)
		err := invoke(ctx, method, req, reply, cc, opts...)
		fmt.Printf("用时:%s", time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	//opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("localhost:8084", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err1 := c.SayHello(context.Background(), &proto.HelloRequest{Name: "刘德华"})
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(r.Message)
}
