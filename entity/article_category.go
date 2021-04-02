package entity

import (
	"go-mysql-canal/model"
	"time"
)

// 文章分类服务文档结构体定义
type ArticleCategoryService struct {
	ArticleCategory
}

// 文章分类模型关系映射
type ToArticleCategoryService struct {
	model.ArticleCategory
}

type ArticleCategory struct {
	Id        int       `json:"id" mapstructure:"id"`
	Pid       int       `json:"pid" mapstructure:"pid"`
	Name      string    `json:"name" mapstructure:"name"`
	Sort      int       `json:"sort" mapstructure:"sort"`
	Status    bool      `json:"status" mapstructure:"status"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
}

// 文章分类服务索引配置项
var ArticleCategoryServiceIndexBody = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 1
	},
	"mappings":{
		"properties":{
			"id":{
				"type":"long"
			},
			"pid":{
				"type":"long"
			},
			"name":{
				"type":"text"
			},
			"sort":{
				"type":"integer"
			},
			"status":{
				"type":"boolean"
			},
			"created_at":{
				"type":"date",
				"format": "strict_date_optional_time||epoch_millis"
			},
			"updated_at":{
				"type":"date",
				"format": "strict_date_optional_time||epoch_millis"
			}
		}
	}
}`

// 组织生成文章分类服务文档数据
func (row ToArticleCategoryService) DocumentArticleCategoryService() ArticleCategoryService {
	document := ArticleCategoryService{
		ArticleCategory: ArticleCategory{
			Id:        row.Id,
			Pid:       row.Pid,
			Name:      row.Name,
			Sort:      row.Sort,
			Status:    row.Status,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		},
	}

	return document
}