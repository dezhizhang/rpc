package main

import (
	"google.golang.org/grpc"
	"net"
	pb "xiaozhi/streamrpc/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) AllStream(streamServer pb.Greeter_AllStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedGreeterServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetStream(req *pb.StreamReqData, res pb.Greeter_GetStreamServer) error {
	return nil
}

func (s *Server) PutStream(clientStream pb.Greeter_PutStreamServer) error {
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
}
