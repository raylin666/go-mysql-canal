package model

import (
	"go-mysql-canal/constant"
	"go-mysql-canal/pkg/database"
	"gorm.io/gorm"
)

func GetMyServerDB() *gorm.DB {
	return database.GetDB("my_server")
}

func GetModelStruct(tableName string) interface{} {
	var (
		modelStruct interface{}
	)

	switch tableName {
	case constant.DbTableArticle:
		modelStruct = Article{}
	case constant.DbTableArticleCategory:
		modelStruct = ArticleCategory{}
	case constant.DbTableArticleExtend:
		modelStruct = ArticleExtend{}
	case constant.DbTableArticleCategoryRelation:
		modelStruct = ArticleCategoryRelation{}
	}

	return modelStruct
}
