package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	pb "xiaozhi/interpretor/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + request.GetName()}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return rsp, status.Error(codes.Unauthenticated, "无token认证信息")
		}
		var appId string
		var appKey string
		if val, ok := md["appId"]; ok {
			appId = val[0]
		}
		if val, ok := md["appKey"]; ok {
			appKey = val[0]
		}

		fmt.Println(appId, appKey)

		if appId != "101010" || appKey != "i am key" {
			return rsp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		//fmt.Println("接收到一个请求")
		i, err := handler(ctx, req)
		return i, err
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	pb.RegisterGreeterServer(g, &Server{})
	listen, err := net.Listen("tcp", ":8084")
	if err != nil {
		panic(err)
	}
	err = g.Serve(listen)
	if err != nil {
		panic(err)
	}
}
