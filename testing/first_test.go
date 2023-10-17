package testing

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"testing"
	p "xiaozhi/proto"
)

func TestName(t *testing.T) {
	//fmt.Println("hello world")
	t.Errorf("测度失败")
}

func TestProto(t *testing.T) {
	user := p.User{
		Username: "刘德华",
		Age:      44,
	}
	marshal, err := proto.Marshal(&user)
	if err != nil {
		panic(err)
	}

	newUser := &p.User{}
	err = proto.Unmarshal(marshal, newUser)

	fmt.Println(newUser)
}
