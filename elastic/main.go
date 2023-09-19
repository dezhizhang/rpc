package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	fmt.Println(client)

	q := elastic.NewMatchQuery("address", "street")
	do, err := client.Search().Index("user").Query(q).Do(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println(do.Hits.Hits)

}
