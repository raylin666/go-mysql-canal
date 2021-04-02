package services

import (
	"go-mysql-canal/constant"
	"go-mysql-canal/entity"
	"go-mysql-canal/model"
	"go-mysql-canal/pkg/elastic"
	"go-mysql-canal/pkg/logger"
	"strconv"
)

// 创建文章分类服务文档数据
func CreateArticleCategoryServiceDocument(modelStruct interface{}, row interface{})  {
	var (
		index = constant.ElasticIndexArticleCategoryService
	)

	switch modelStruct.(type) {
	case model.ArticleCategory:
		// 模型类型映射
		rowModel := row.(model.ArticleCategory)
		switch rowModel.Status {
		case true:
			document := entity.ToArticleCategoryService{rowModel}.DocumentArticleCategoryService()
			// 创建文档
			_, err := elastic.GetClient().CreateDocument(index, strconv.Itoa(rowModel.Id), document)
			if err != nil {
				logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
					"err":    err,
					"index":  index,
					"id":     rowModel.Id,
					"body":   document,
					"action": "create",
				}.Fields()).Error("elasticsearch create document err")
			}

			return
		}

		return
	}
}

// 更新文章分类服务文档数据
func UpdateArticleCategoryServiceDocument(modelStruct interface{}, row interface{}) {
	var (
		index = constant.ElasticIndexArticleCategoryService
	)

	switch modelStruct.(type) {
	case model.ArticleCategory:
		// 模型类型映射
		rowModel := row.(model.ArticleCategory)
		if ok, _ := elastic.GetClient().ExistsDocument(index, strconv.Itoa(rowModel.Id)); !ok {
			switch rowModel.Status {
			case true:
				document := entity.ToArticleCategoryService{rowModel}.DocumentArticleCategoryService()
				// 创建文档
				_, err := elastic.GetClient().CreateDocument(index, strconv.Itoa(rowModel.Id), document)
				if err != nil {
					logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
						"err":    err,
						"index":  index,
						"id":     rowModel.Id,
						"body":   document,
						"action": "update",
					}.Fields()).Error("elasticsearch create document err")
				}

				return
			}
		} else {
			switch rowModel.Status {
			case true:
				document := entity.ToArticleCategoryService{rowModel}.DocumentArticleCategoryService()
				// 更新文档
				_, err := elastic.GetClient().UpdateDocument(index, strconv.Itoa(rowModel.Id), document)
				if err != nil {
					logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
						"err":    err,
						"index":  index,
						"id":     rowModel.Id,
						"body":   document,
						"action": "update",
					}.Fields()).Error("elasticsearch update document err")
				}

				return
			case false:
				// 删除文档
				_, err := elastic.GetClient().DeleteDocument(index, strconv.Itoa(rowModel.Id))
				if err != nil {
					logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
						"err":    err,
						"index":  index,
						"id":     rowModel.Id,
						"action": "update",
					}.Fields()).Error("elasticsearch delete document err")
				}

				return
			}
		}

		return
	}
}

// 删除文章分类服务文档数据
func DeleteArticleCategoryServiceDocument(modelStruct interface{}, row interface{})  {
	var (
		index = constant.ElasticIndexArticleCategoryService
	)

	switch modelStruct.(type) {
	case model.ArticleCategory:
		// 模型类型映射
		rowModel := row.(model.ArticleCategory)
		if ok, _ := elastic.GetClient().ExistsDocument(index, strconv.Itoa(rowModel.Id)); ok {
			// 删除文档
			_, err := elastic.GetClient().DeleteDocument(index, strconv.Itoa(rowModel.Id))
			if err != nil {
				logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
					"err":    err,
					"index":  index,
					"id":     rowModel.Id,
					"action": "delete",
				}.Fields()).Error("elasticsearch delete document err")
			}
		}

		return
	}
}