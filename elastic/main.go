package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	host := "http://localhost:9200"
	logger := log.New(os.Stdout, "xiaozhi", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}
	q := elastic.NewMatchQuery("address", "street")
	do, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}

	for _, value := range do.Hits.Hits {
		var user User
		err := json.Unmarshal(value.Source, &user)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	}

	user := User{Name: "张德地", Age: 30}
	put, err := client.Index().Index("user").Id("3").BodyJson(&user).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(put)

}
