package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

func main() {
	host := "http://localhost:9200"
	logger := log.New(os.Stdout, "xiaozhi", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	do, i, _ := client.Ping(host).Do(context.Background())

	version, _ := client.ElasticsearchVersion(host)
	fmt.Println(do, i)
	fmt.Println("-----------")
	fmt.Println(version)

}
