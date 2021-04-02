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

// 创建默认索引 (不设置配置项)
func (c *Client) CreateDefaultIndex(index string) (*elastic.IndicesCreateResult, error) {
	return c.Elastic.CreateIndex(index).Do(ctx)
}

// 创建索引并设置配置项 (STRING 风格)
func (c *Client) CreateIndexToBodyString(index string, body string) (*elastic.IndicesCreateResult, error) {
	return c.Elastic.CreateIndex(index).BodyString(body).Do(ctx)
}

// 创建索引并设置配置项 (JSON 风格)
func (c *Client) CreateIndexToBodyJson(index string, body interface{}) (*elastic.IndicesCreateResult, error) {
	return c.Elastic.CreateIndex(index).BodyJson(body).Do(ctx)
}

// 删除索引
func (c *Client) DeleteIndex(index string) (*elastic.IndicesDeleteResponse, error) {
	return c.Elastic.DeleteIndex(index).Do(ctx)
}

// 获取索引设置
func (c *Client) GetSettings(index string) (map[string]*elastic.IndicesGetSettingsResponse, error) {
	return c.Elastic.IndexGetSettings(index).Do(ctx)
}

// 设定索引设置 (STRING 风格)
func (c *Client) PutSettingsToString(index string, setting string) (*elastic.IndicesPutSettingsResponse, error) {
	return c.Elastic.IndexPutSettings(index).BodyString(setting).Do(ctx)
}

// 设定索引设置 (JSON 风格)
func (c *Client) PutSettingsToJson(index string, setting interface{}) (*elastic.IndicesPutSettingsResponse, error) {
	return c.Elastic.IndexPutSettings(index).BodyJson(setting).Do(ctx)
}

// 获取索引映射
func (c *Client) GetMapping(index string) (map[string]interface{}, error) {
	return c.Elastic.GetMapping().Index(index).Do(ctx)
}

// 设定索引映射 (JSON 风格)
func (c *Client) PutMappingToJson(index string, mappings map[string]interface{}) (*elastic.PutMappingResponse, error) {
	res, err := c.Elastic.PutMapping().Index(index).BodyJson(mappings).Do(ctx)
	if err == nil {
		return res, nil
	}

	logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
		"index":    index,
		"mappings": mappings,
		"format":   "json",
	}.Fields()).Error()

	return nil, err
}

// 设定索引映射 (STRING 风格)
func (c *Client) PutMappingToString(index string, mappings string) (*elastic.PutMappingResponse, error) {
	res, err := c.Elastic.PutMapping().Index(index).BodyString(mappings).Do(ctx)
	if err == nil {
		return res, nil
	}

	logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
		"index":    index,
		"mappings": mappings,
		"format":   "string",
		"err":      err,
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

// 判断文档是否存在
func (c *Client) ExistsDocument(index string, id string) (bool, error) {
	return c.Elastic.Exists().Index(index).Id(id).Do(ctx)
}

// 创建文档
func (c *Client) CreateDocument(index string, id string, body interface{}) (*elastic.IndexResponse, error) {
	res, err := c.Elastic.Index().Index(index).Id(id).BodyJson(body).Refresh("wait_for").Do(ctx)
	if err == nil {
		return res, nil
	}

	logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
		"index":       index,
		"document id": id,
		"body":        body,
		"format":      "json",
		"err":         err,
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
