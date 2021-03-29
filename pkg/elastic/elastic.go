package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-mysql-canal/config"
	"go-mysql-canal/constant"
	"go-mysql-canal/pkg/logger"
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

// 获取所有索引
func (c *Client) Indices() (elastic.CatIndicesResponse, error) {
	return c.Elastic.CatIndices().Do(ctx)
}

// 索引状态
func (c *Client) IndexStats(indices ...string) (*elastic.IndicesStatsResponse, error) {
	return c.Elastic.IndexStats(indices...).Do(ctx)
}

// 索引是否存在
func (c *Client) IndexExists(index string) (bool, error) {
	return c.Elastic.IndexExists(index).Do(ctx)
}

// 创建索引
func (c *Client) CreateIndex(name string) (*elastic.IndicesCreateResult, error) {
	return c.Elastic.CreateIndex(name).Do(ctx)
}

// 删除索引
func (c *Client) DeleteIndex(name string) (*elastic.IndicesDeleteResponse, error) {
	return c.Elastic.DeleteIndex(name).Do(ctx)
}

// 获取索引映射
func (c *Client) GetMapping(name string) (map[string]interface{}, error) {
	return c.Elastic.GetMapping().Index(name).Do(ctx)
}

// 设定索引映射 (JSON 风格)
func (c *Client) PutMappingToJson(name string, json map[string]interface{}) (*elastic.PutMappingResponse, error) {
	builder := c.Elastic.PutMapping().Index(name).BodyJson(json)
	err := builder.Validate()
	if err == nil {
		return builder.Do(ctx)
	}

	logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
		"索引名称": name,
		"索引映射": json,
		"映射格式": "json",
	}.Fields()).Error()

	return nil, err
}

// 设定索引映射 (STRING 风格)
func (c *Client) PutMappingToString(name string, string string) (*elastic.PutMappingResponse, error) {
	builder := c.Elastic.PutMapping().Index(name).BodyString(string)
	err := builder.Validate()
	if err == nil {
		return builder.Do(ctx)
	}

	logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
		"索引名称": name,
		"索引映射": string,
		"映射格式": "string",
		"err": err,
	}.Fields()).Error()

	return nil, err
}

// 搜索文档
func (c *Client) SearchDocument(index string, query elastic.Query, fields ...string) (*elastic.SearchResult, error) {
	return c.Elastic.Search().Index(index).Query(query).DocvalueFields(fields...).Do(ctx)
}

// 获取文档
func (c *Client) GetDocument(index string, id string) (*elastic.GetResult, error) {
	return c.Elastic.Get().Index(index).Id(id).Do(ctx)
}

// 添加文档
func (c *Client) AddDocument(index string, id string, body interface{}) (*elastic.IndexResponse, error) {
	builder := c.Elastic.Index().Index(index).Id(id).BodyJson(body)
	err := builder.Validate()
	if err == nil {
		return builder.Refresh("wait_for").Do(ctx)
	}

	logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
		"索引名称": index,
		"文档ID": id,
		"文档内容": body,
		"文档格式": "json",
		"err": err,
	}.Fields()).Error()

	return nil, err
}

// 更新文档
func (c *Client) UpdateDocument(index string, id string, body interface{}) (*elastic.UpdateResponse, error) {
	return c.Elastic.Update().Index(index).Id(id).Doc(body).Refresh("wait_for").Do(ctx)
}

// 删除文档
func (c *Client) DeleteDocument(index string, id string) (*elastic.DeleteResponse, error) {
	return c.Elastic.Delete().Index(index).Id(id).Do(ctx)
}