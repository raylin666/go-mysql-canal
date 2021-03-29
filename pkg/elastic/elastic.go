package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-mysql-canal/config"
)

var (
	client *Client
	ctx    = context.Background()
)

type Client struct {
	Elastic *elastic.Client
}

func NewClient() *Client {
	var (
		err error
		c   Client
	)
	c.Elastic, err = elastic.NewClient(
		elastic.SetURL(config.ElasticConfig.Urls),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	client = &c
	defer c.Elastic.Stop()
	return client
}

func GetClient() *Client {
	return client
}

// 索引是否存在
func (c *Client) IndexExists(index string) (bool, error) {
	return c.Elastic.IndexExists(index).Do(ctx)
}
