package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	pb "xiaozhi/grpc/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + request.GetName()}, nil
}

func main() {
	g := grpc.NewServer()
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
