package entity

type ArticleService struct {
	Article
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
	CreatedAt         string 	`json:"created_at" mapstructure:"created_at"`
	UpdatedAt         string 	`json:"updated_at" mapstructure:"updated_at"`
}

var ArticleServiceIndexDefaultBody = `
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
			}
		}
	}
}`
