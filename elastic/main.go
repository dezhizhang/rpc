package main

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

//func createIndex(client *elastic.Client, indexName string) {
//	ctx := context.Background()
//	exists, err := client.IndexExists(indexName).Do(ctx)
//
//	if err != nil {
//		log.Fatalf("Error checking if index [%s] exists: %v", indexName, err)
//	}
//
//	if !exists {
//		createIndex, err := client.CreateIndex(indexName).Do(ctx)
//
//		if err != nil {
//			log.Fatalf("Error creating index: %v", err)
//		}
//		if createIndex.Acknowledged {
//			fmt.Printf("Index [%s] created\n", indexName)
//		}
//	}
//}

func main() {
	//host := "http://localhost:9200"
	//logger := log.New(os.Stdout, "xiaozhi", log.LstdFlags)
	//client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	//if err != nil {
	//	panic(err)
	//}

	//createIndex(client, "article")

	//// 创建索引
	//do, err := client.CreateIndex("goods").BodyString(mapping).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(do)

}
