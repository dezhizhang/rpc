# rpc
## 运程调用服务端
### 服务端
```go
func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("解析参数错误:%s", err)
		}
		fmt.Println("path", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])

		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		w.Write(jData)
	})
	http.ListenAndServe(":8080", nil)
}
```
### 客户端
```go
type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()

	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8080/add?a=%v&b=%v", a, b))
	body, _ := res.Body()

	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)

	fmt.Println(string(body))
	return rspData.Data
}

func main() {
	rsp := Add(1, 2)
	fmt.Println(rsp)
	//http.NewRequest()
}

```
## rcp调用
### 服务端
```go
type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
func main() {
	//1 实例化一个server
	listen, _ := net.Listen("tcp", "localhost:8084")

	//2 注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 启动服务
	conn, _ := listen.Accept()

	rpc.ServeConn(conn)
}
```
### 客户端
```go
func main() {
	client, err := rpc.Dial("tcp", "localhost:8084")
	if err != nil {
		log.Printf("链接失败%s", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)

}
```
## jsonrpc调用
### 服务端
```go
type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	_ = rpc.RegisterName("HelloService", &HelloService{})

	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}

```
### 客户端
```go
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("连接失败")
	}
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}

```
## go下的grpc调用
### proto文件
```go
syntax = "proto3";

option go_package = ".;proto";


service Greeter {
  rpc SayHello(HelloRequest) returns(HelloReply){}
}


message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```
### 服务端
```go
type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) mustEmbedUnimplementedGreeterServer() {
	//TODO implement me
	fmt.Println("hello")

}

// SayHello rpc调用
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

```
### 客户端
```go
func main() {
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	r, err1 := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby"})
	if err1 != nil {
		fmt.Println("err", err1.Error())
	}

	fmt.Println(r.Message)

}
```
