package model

import "go-mysql-canal/constant"

type ArticleExtend struct {
	ArticleId      int    `json:"article_id" mapstructure:"article_id"`
	Source         string `json:"source" mapstructure:"source"`
	SourceUrl      string `json:"source_url" mapstructure:"source_url"`
	Content        string `json:"content" mapstructure:"content"`
	Keyword        string `json:"keyword" mapstructure:"keyword"`
	AttachmentPath string `json:"attachment_path" mapstructure:"attachment_path"`
}

type ArticleExtendTable struct {}

func (t *ArticleExtendTable) TableName() string {
	return constant.DbTableArticleExtend
}

