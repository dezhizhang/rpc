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
}

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
