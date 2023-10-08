package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"
)

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
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

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//
//	user := User{
//		Name: "tom",
//		Sex:  "女",
//		Tel:  "15083356190",
//	}
//	do, err := client.Update().Index("user").Id("1002").Doc(&user).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(do.Result)
//}

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	user := User{
//		Name: "刘德华",
//		Sex:  "男",
//		Tel:  "1541609448@qq.com",
//	}
//	do, err := client.Update().Index("user").Id("1002").Doc(&user).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(do.Result)
//}

//func main() {
//	//host := "http://localhost:9200"
//	//logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	//
//	//client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//do, err := client.Get().Index("user").Id("1002").Do(context.Background())
//	//if do.Found {
//	//	fmt.Printf("got document %s in version %d form index%s", do.Id, do.Version, do.Index)
//	//	var user User
//	//	err := json.Unmarshal(do.Source, &user)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//	fmt.Println(user.Name, user.Sex, user.Tel)
//	//}
//
//}

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	do, err := client.Get().Index("user").Id("1002").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	if do.Found {
//		var user User
//		err := json.Unmarshal(do.Source, &user)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(user.Name, user.Sex, user.Tel)
//
//	}
//}

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	user := User{
//		Name: "周华呢",
//		Sex:  "男",
//		Tel:  "15992478448",
//	}
//	do, err := client.Update().Index("user").Id("1002").Doc(&user).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(do.Result)
//
//}

//func main() {
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//
//	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	user := User{
//		Name: "周大大",
//		Sex:  "男",
//		Tel:  "1599278448",
//	}
//	do, err := client.Update().Index("user").Id("1002").Doc(&user).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(do.Result)
//
//}

//func main() {
//	var err error
//	var user User
//	var client *elastic.Client
//	var res *elastic.SearchResult
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//
//	res, err = client.Search("user").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//
//	for _, item := range res.Each(reflect.TypeOf(user)) {
//		t := item.(User)
//		fmt.Printf("%#v\t", t)
//	}
//
//}

//func main() {
//	var err error
//	var user User
//	var client *elastic.Client
//	var res *elastic.SearchResult
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//
//	client, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	res, err = client.Search("user").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	for _, item := range res.Each(reflect.TypeOf(user)) {
//		t := item.(User)
//		fmt.Println(t)
//	}
//
//}

// ### 范围查询
//func main() {
//	var err error
//	var client *elastic.Client
//	var res *elastic.SearchResult
//
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//
//	boolq := elastic.NewBoolQuery()
//	boolq.Must(elastic.NewMatchQuery("name", "刘德华"))
//	boolq.Filter(elastic.NewRangeQuery("age").Gt(30))
//	res, err = client.Search("user").Query(boolq).Do(context.Background())
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(res)
//}

//func main() {
//	var err error
//	var client *elastic.Client
//	var res *elastic.SearchResult
//
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//
//	boolq := elastic.NewBoolQuery()
//	boolq.Must(elastic.NewMatchQuery("name", "刘德华"))
//	boolq.Filter(elastic.NewRangeQuery("age").Gt(30))
//	res, err = client.Search("user").Query(boolq).Do(context.Background())
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("----", res.Hits)
//}

//func main() {
//	var err error
//	var client *elastic.Client
//	var res *elastic.SearchResult
//
//	host := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//	client, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	matchPhase := elastic.NewMatchPhraseQuery("about", "华")
//	res, err = client.Search("user").Query(matchPhase).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(res)
//}

//func main() {
//	var err error
//	var client *elastic.Client
//	var res *elastic.SearchResult
//
//	hot := "http://localhost:9200"
//	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
//
//	client, err = elastic.NewClient(elastic.SetURL(hot), elastic.SetSniff(false), elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//	matchPhase := elastic.NewMatchPhraseQuery("about", "刘")
//	res, err = client.Search("user").Query(matchPhase).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(res)
//
//}

func main() {
	// Tweet is a structure used for serializing/deserializing data in Elasticsearch.

	const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"change":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`
	host := "http://localhost:9200"
	logger := log.New(os.Stdout, "elastic", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}
	createIndex, err := client.CreateIndex("change").BodyString(mapping).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}
}
