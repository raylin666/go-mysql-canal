package entity

import (
	"encoding/json"
	"go-mysql-canal/model"
	"go-mysql-canal/pkg/utils"
	"time"
)

// 文章服务文档结构体定义
type ArticleService struct {
	Article
	ArticleExtend
	Category []ArticleCategory
}

// 文章模型关系映射
type WithArticleService struct {
	model.WithArticle
}

// 文章分类服务文档结构体定义
type ArticleCategoryService struct {
	ArticleCategory
}

// 文章分类模型关系映射
type ToArticleCategoryService struct {
	model.ArticleCategory
}

type Article struct {
	Id                int       `json:"id" mapstructure:"id"`
	Title             string    `json:"title" mapstructure:"title"`
	Author            string    `json:"author" mapstructure:"author"`
	Summary           string    `json:"summary" mapstructure:"summary"`
	Cover             string    `json:"cover" mapstructure:"cover"`
	Sort              int       `json:"sort" mapstructure:"sort"`
	RecommendFlag     bool      `json:"recommend_flag" mapstructure:"recommend_flag"`
	CommentFlag       bool      `json:"comment_flag" mapstructure:"comment_flag"`
	Status            bool      `json:"status" mapstructure:"status"`
	ViewCount         int       `json:"view_count" mapstructure:"view_count"`
	CommentCount      int       `json:"comment_count" mapstructure:"comment_count"`
	ShareCount        int       `json:"share_count" mapstructure:"share_count"`
	PublisherUsername string    `json:"publisher_username" mapstructure:"publisher_username"`
	UserId            int       `json:"user_id" mapstructure:"user_id"`
	LastCommentTime   int       `json:"last_comment_time" mapstructure:"last_comment_time"`
	CreatedAt         time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" mapstructure:"updated_at"`
}

type ArticleExtend struct {
	ArticleId      int         `json:"article_id" mapstructure:"article_id"`
	Source         string      `json:"source" mapstructure:"source"`
	SourceUrl      string      `json:"source_url" mapstructure:"source_url"`
	Content        string      `json:"content" mapstructure:"content"`
	Keyword        string      `json:"keyword" mapstructure:"keyword"`
	AttachmentPath interface{} `json:"attachment_path" mapstructure:"attachment_path"`
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

// 文章服务索引配置项
var ArticleServiceIndexBody = `
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
			"title":{
				"type":"text"
			},
			"summary":{
				"type":"text"
			},
			"author":{
				"type":"text",
				"fields":{
					"keyword":{
						"type":"keyword",
						"ignore_above":30
					}
				}
			},
			"cover":{
				"type":"text",
				"fields":{
					"keyword":{
						"type":"keyword",
						"ignore_above":200
					}
				}
			},
			"sort":{
				"type":"integer"
			},
			"recommend_flag":{
				"type":"boolean"
			},
			"comment_flag":{
				"type":"boolean"
			},
			"status":{
				"type":"boolean"
			},
			"view_count":{
				"type":"integer"
			},
			"comment_count":{
				"type":"integer"
			},
			"share_count":{
				"type":"integer"
			},
			"publisher_username":{
				"type":"text",
				"fields":{
					"keyword":{
						"type":"keyword",
						"ignore_above":30
					}
				}
			},
			"user_id":{
				"type":"long"
			},
			"last_comment_time":{
				"type":"integer"
			},
			"created_at":{
				"type":"date",
				"format": "strict_date_optional_time||epoch_millis"
			},
			"updated_at":{
				"type":"date",
				"format": "strict_date_optional_time||epoch_millis"
			},
			"article_id":{
				"type":"long"
			},
			"source":{
				"type":"text",
				"fields":{
					"keyword":{
						"type":"keyword",
						"ignore_above":32
					}
				}
			},
			"source_url":{
				"type":"text",
				"fields":{
					"keyword":{
						"type":"keyword",
						"ignore_above":200
					}
				}
			},
			"content":{
				"type":"text"
			},
			"keyword":{
				"type":"text"
			},
			"attachment_path":{
				"type":"text",
				"index":"false"
			},
			"category": {
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
		}
	}
}`

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

// 生成文章服务文档
func (row WithArticleService) DocumentArticleService() ArticleService {
	document := ArticleService{
		Article: Article{
			Id:                row.Id,
			Title:             row.Title,
			Author:            row.Author,
			Summary:           row.Summary,
			Cover:             row.Cover,
			Sort:              row.Sort,
			RecommendFlag:     utils.Int8ToBoolean(row.RecommendFlag),
			CommentFlag:       utils.Int8ToBoolean(row.CommentFlag),
			Status:            utils.Int8ToBoolean(row.Status),
			ViewCount:         row.ViewCount,
			CommentCount:      row.CommentCount,
			ShareCount:        row.ShareCount,
			PublisherUsername: row.PublisherUsername,
			UserId:            row.UserId,
			LastCommentTime:   row.LastCommentTime,
			CreatedAt:         row.CreatedAt,
			UpdatedAt:         row.UpdatedAt,
		},
		ArticleExtend: ArticleExtend{
			ArticleId:      row.ArticleExtend.ArticleId,
			Source:         row.ArticleExtend.Source,
			SourceUrl:      row.ArticleExtend.SourceUrl,
			Content:        row.ArticleExtend.Content,
			Keyword:        row.ArticleExtend.Keyword,
			AttachmentPath: row.BuilderArticleServiceAttachmentPath(),
		},
		Category: row.BuilderArticleServiceCategory(),
	}

	return document
}

// 构建文章服务附件数据
func (row WithArticleService) BuilderArticleServiceAttachmentPath() interface{} {
	var attachmentPath interface{}
	_ = json.Unmarshal([]byte(row.ArticleExtend.AttachmentPath), &attachmentPath)
	return attachmentPath
}

// 构建文章服务分类数据
func (row WithArticleService) BuilderArticleServiceCategory() []ArticleCategory {
	var category = make([]ArticleCategory, len(row.ArticleCategoryRelation))
	for k, relation := range row.ArticleCategoryRelation {
		category[k] = ArticleCategory(model.ArticleCategory{
			Id:        relation.Category.Id,
			Pid:       relation.Category.Pid,
			Name:      relation.Category.Name,
			Sort:      relation.Category.Sort,
			Status:    relation.Category.Status,
			CreatedAt: relation.Category.CreatedAt,
			UpdatedAt: relation.Category.UpdatedAt,
		})
	}
	return category
}

// 生成文章分类服务文档
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