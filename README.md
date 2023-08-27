# rpc

### 生成proto
```go
 protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative hello.proto
```
## 运程调用服务端

```go
### 服务端
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
## protobuf不同proto的引用
### 通过import来引用
```go
syntax = "proto3";
import "base.proto";
import  "google/protobuf/empty.proto";
option go_package=".;proto";

service Greeter {
  rpc SayHello(HelloRequest) returns(HelloReply);
  rpc Ping(google.protobuf.Empty) returns(Pong);
}

message HelloRequest {
  string url = 1;
  string name = 2;
}

message HelloReply {
  string message = 1;
}
```
### protobuf枚举类型
```go
enum Gender {
  MALE = 0;
  FEMALE = 1;
}

message HelloRequest {
  string url = 1;
  string name = 2;
  Gender  gender = 3;
}
```
### 枚举类型的引用
```go
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
	})
	fmt.Println(rsp.Message)
}

```
### protobuf中map的定义
```go
message HelloRequest {
  string url = 1; //网站
  string name = 2; // 姓名
  Gender  gender = 3; // 姓别
  map<string,string> mp = 4;
}
```
### map的使用
```go
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
	})
	fmt.Println(rsp.Message)
}
```
### 使用protobuf内置的timestamp类型
```go
message HelloRequest {
    string url = 1; //网站
    string name = 2; // 姓名
    Gender  gender = 3; // 姓别
    map<string,string> mp = 4;
    google.protobuf.Timestamp createTime = 5;
}
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
```
## grpc拦截器
### 服务端
```go
type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + request.GetName()}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp interface{}, err error) {
		fmt.Println("接收到一个新的请求")
		return nil, nil
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

```
### 客户端
```go
func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoke grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoke(ctx, method, req, reply, cc, opts...)
		fmt.Printf("用时:%s", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure(), opt)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err1 := c.SayHello(context.Background(), &proto.HelloRequest{Name: "刘德华"})
	if err1 != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}

```


## grpc验证
### 服务端
```go
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

```
### 客户端
```go
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

```
### gorm的批量插入
```go
func main() {
	var users []model.User
	for i := 0; i < 100; i++ {
		users = append(users, model.User{Username: fmt.Sprintf("刘德华%d", i), Age: i})
	}

	tx := driver.DB.Model(&model.User{}).CreateInBatches(&users, 100)
	fmt.Println(tx.RowsAffected)

}
```
### gorm询操作
```go
func main() {

	var user model.User
	tx := driver.DB.Model(&user).Where("username = ?", "刘德华20").Find(&user)
	if tx.Error != nil {
		log.Fatalf("获取失败%s", tx.Error)
	}
	fmt.Println(user)

}

```
### 更新操作
```go
func main() {

	var user model.User
	driver.DB.Where(&model.User{Username: "刘德华0"}).Find(&user)

	user.Username = "周华建"
	err := driver.DB.Save(&user).Error
	if err != nil {
		log.Fatalf("更新失败%s", err)
	}
	fmt.Println("更新成功")

}
```
### 更新操作
```go
err := driver.DB.Where("id = ?", 205).Updates(&model.User{Username: "哈哈", Age: 44}).Error
    if err != nil {
		log.Printf("更新失败%s", err)
	}
	
	fmt.Println("更新成功")
```
### 删除操作
```go
func main() {
	
	err := driver.DB.Where("id = ?", 206).Unscoped().Delete(&model.User{}).Error
	if err != nil {
		log.Printf("删除失败%s", err)
	}

	fmt.Println("删除成功")

}
```