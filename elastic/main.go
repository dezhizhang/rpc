package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Tel  string `json:"tel"`
}

// 新增
//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	user := &User{Name: "tom", Sex: "1", Tel: "15992478448"}
//	do, err1 := client.Index().Index("user").Type("_create").Id("1002").BodyJson(user).Do(context.Background())
//	if err1 != nil {
//		panic(err1)
//	}
//
//	fmt.Println(do)
//}

// 查询
//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	do, err := client.Get().Index("user").Id("1002").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	if do.Found {
//		fmt.Printf("Got document %s in version %d from index %s, type %s\n", do.Id, do.Version, do.Index, do.Type)
//		var user User
//		err := json.Unmarshal(do.Source, &user)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(user.Name, user.Sex, user.Tel)
//	}
//}

// 删除文档
//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//
//	if err != nil {
//		panic(err)
//	}
//	do, err := client.Delete().Index("user").Id("1002").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(do.Result)
//
//}

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	do, err := client.Delete().Index("user").Id("1002").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(do.Result)
//}

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//
//	user := User{
//		Name: "hello",
//		Sex:  "女",
//		Tel:  "159924784481",
//	}
//	do, err := client.Update().Index("user").Id("1002").Doc(&user).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(do.Result)
//
//}

func main() {
	host := "http://localhost:9200"
	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))

	user := User{
		Name: "tom",
		Sex:  "女",
		Tel:  "15083356190",
	}
	do, err := client.Update().Index("user").Id("1002").Doc(&user).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(do.Result)
}
