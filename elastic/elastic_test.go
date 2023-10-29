package elastic

import (
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestInsert(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}
	t.Log(client)
}
