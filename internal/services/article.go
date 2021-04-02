package services

import "go-mysql-canal/model"

// 更新文章服务文档数据
func UpdateArticleServiceDocument(modelStruct interface{}, rows map[string]interface{}) {
	switch modelStruct.(type) {
	case model.Article:
	case model.ArticleExtend:
	case model.ArticleCategory:
	}
}
