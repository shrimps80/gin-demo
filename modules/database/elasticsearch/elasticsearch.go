package elasticsearch

import (
	"context"
	"encoding/json"
	"gin-demo/config"
	"time"

	"github.com/olivere/elastic"
)

type EsType struct {
	Conn *elastic.Client
}

var (
	Client    EsType
	indexName = config.GetEnv().EsIndexName
)

func init() {
	var err error

	Client.Conn, err = elastic.NewClient(
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(3),
		elastic.SetURL(config.GetEnv().EsServers),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
}

// 创建es的结构
func (client *EsType) IndexExists(mapping string) {
	ctx := context.Background()
	exists, err := client.Conn.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.Conn.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
	return
}

func (client *EsType) SetIndex(id string, data interface{}) *elastic.IndexResponse {
	ctx := context.Background()
	doc, err := client.Conn.Index().
		Index(indexName).
		Id(id).
		BodyJson(data).
		Refresh("wait_for").
		Do(ctx)

	if err != nil {
		panic(err)
	}

	return doc
}

func (client *EsType) GetIndex(id string) interface{} {
	ctx := context.Background()
	result, err := client.Conn.Get().
		Index(indexName).
		Id(id).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	if result.Found {
		var buf interface{}
		err := json.Unmarshal(result.Source, &buf)
		if err != nil {
			panic(err)
		}
		return buf
	}
	return nil
}

func (client *EsType) DelIndex(id string) bool {
	ctx := context.Background()
	res, err := client.Conn.Delete().
		Index(indexName).
		Id(id).
		Refresh("wait_for").
		Do(ctx)
	if err != nil {
		panic(err)
	}
	if res.Result == "deleted" {
		return true
	}
	return false
}
