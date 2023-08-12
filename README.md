# rpc
### 运程调用服务端
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
