package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
	pb "xiaozhi/streamrpc/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) AllStream(allServer pb.Greeter_AllStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetStream(req *pb.StreamReqData, res pb.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&pb.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *Server) PutStream(clientStream pb.Greeter_PutStreamServer) error {

	for {
		//if a, err := clientStream.Recv(); err != nil {
		//	fmt.Println(err)
		//	break
		//} else {
		//	fmt.Println(a)
		//}
		recv, err := clientStream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(recv)

	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	err1 := s.Serve(listen)
	if err1 != nil {
		panic(err1)
	}

}
