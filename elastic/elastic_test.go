package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}
	t.Log(client)
}

type User struct {
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	CreateAt time.Time `json:"create_at"`
	Desc     string    `json:"desc"`
}

// 创建索引
func TestIndexMapping(t *testing.T) {
	mapping := `
		{
			"settings" : {
				"number_of_shards" : 1,
				"number_of_replicas" : 1
			},
			"mappings":{
				"properties":{
					"nickname":{
						"type":"text"
 					},
					"username":{
						"type":"keyword"
					},
					"id":{
						"type":"integer"
					},
					"create_at":{
						"type":"date",
						"format":"[YYYY-MM-dd HH:mm:ss]"
					}
                 }
			}
		}
	`
	//t.Log(mapping)
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}

	do, err1 := client.CreateIndex("user").BodyString(mapping).Do(context.Background())
	if err1 != nil {
		t.Logf("创建mapping失败%s", err1.Error())
		return
	}
	t.Log("创建mapping成功", do)

}

// 判断索引是否存在
func TestIndexExists(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}
	do, err1 := client.IndexExists("user").Do(context.Background())
	if err1 != nil {
		panic(err1)
	}
	t.Log(do)
}

// 删除索引
func TestDeleteIndex(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err)
	}

	do, err1 := client.DeleteIndex("user").Do(context.Background())
	if err1 != nil {
		panic(err1)
	}
	t.Log(do)

}

// 添加文档
func TestCreateDoc(t *testing.T) {
	user := &User{
		ID:       1,
		Username: "张德志",
		Nickname: "晓智",
		CreateAt: time.Now(),
		Desc:     "It从业人员",
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)

	if err != nil {
		panic(err)
	}

	do, err1 := client.Index().Index("user").BodyJson(user).Do(context.Background())
	if err1 != nil {
		panic(err1)
	}
	t.Log(do)

}

// 根据id删除
func TestDeleteDocById(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}
	deleteId := "7VKzh4sBP_UzlGBnhzrD"
	do, err1 := client.Delete().Index("user").Id(deleteId).Refresh("true").Do(context.Background())
	if err1 != nil {
		panic(err1)
	}

	t.Log(do)

}

// 批量删除
func TestBatchDeleteDocById(t *testing.T) {
	list := []string{"91LDh4sBP_UzlGBnwzpC", "9VLDh4sBP_UzlGBntjr8"}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}
	bulk := client.Bulk().Index("user").Refresh("true")
	for _, s := range list {
		req := elastic.NewBulkDeleteRequest().Id(s)
		bulk.Add(req)
	}

	do, err1 := bulk.Do(context.Background())
	if err1 != nil {
		panic(err1)
	}
	t.Log(do.Succeeded())
}

// 批量添加
func TestBatchCreate(t *testing.T) {
	list := []User{
		{
			ID:       12,
			Username: "张三",
			Nickname: "三",
			CreateAt: time.Now(),
		},
		{
			ID:       13,
			Username: "李四",
			Nickname: "阿四",
			CreateAt: time.Now(),
		},
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}
	bulk := client.Bulk().Index("user").Refresh("true")
	for _, user := range list {
		doc := elastic.NewBulkCreateRequest().Doc(&user)
		bulk.Add(doc)
	}

	do, err1 := bulk.Do(context.Background())
	if err1 != nil {
		panic(err1)
	}

	t.Log(do.Created())
}
