package testing

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"testing"
	pb "xiaozhi/proto"
)

func TestName(t *testing.T) {
	//fmt.Println("hello world")
	t.Errorf("测度失败")
}

func TestProto(t *testing.T) {
	user := pb.User{
		Username: "周大大",
		Age:      22,
	}
	marshal, err := proto.Marshal(&user)
	if err != nil {
		panic(err)
	}
	newUser := &pb.User{}
	err = proto.Unmarshal(marshal, newUser)
	fmt.Println(newUser)
}
