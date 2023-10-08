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

func main() {
	host := "http://localhost:9200"
	logger := log.New(os.Stdout, "xiaozhi", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	user := &User{Name: "周华建", Sex: "男", Tel: "15083356190"}
	do, _ := client.Index().Index("info").Type("user").Id("1001").BodyJson(user).Do(context.Background())

	fmt.Println("-----", do)

	//do, i, _ := client.Ping(host).Do(context.Background())

	//version, _ := client.ElasticsearchVersion(host)
	//fmt.Println(do, i)
	//fmt.Println("-----------")
	//fmt.Println(version)

}

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "xiaozhi", log.LstdFlags)
//	cient, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//
//}
