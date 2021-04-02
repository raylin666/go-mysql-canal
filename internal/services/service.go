package services

import (
	"go-mysql-canal/constant"
	"go-mysql-canal/entity"
	"go-mysql-canal/model"
	"go-mysql-canal/pkg/elastic"
	"go-mysql-canal/pkg/logger"
	"strconv"
	"sync"
)

// 初始化创建\同步数据到 Elastic
func InitElasticDataSync() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		// 初始化文章数据
		initArticleService()
		wg.Done()
	}()

	go func() {
		// 初始化文章分类数据
		initArticleCategoryService()
		wg.Done()
	}()

	wg.Wait()
}

// 初始化文章数据
func initArticleService() {
	var (
		// 服务索引
		index = constant.ElasticIndexArticleService
		// 文章服务配置项
		indexBody = entity.ArticleServiceIndexBody
		// 初始化数据函数
		callback func()
	)

	callback = func() {
		// 初始化数据
		rows := model.GetWithArticleRows()
		for _, row := range rows {
			// 文档生成
			document := entity.WithArticleService{row}.DocumentArticleService()
			// 创建文档
			_, err := elastic.GetClient().CreateDocument(index, strconv.Itoa(row.Id), document)
			if err != nil {
				logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
					"err":   err,
					"index": index,
					"id":    row.Id,
					"body":  document,
				}.Fields()).Error("elasticsearch add document err")
			}
		}
	}

	createIndexToInitDocument(index, indexBody, callback)
}

// 初始化文章分类数据
func initArticleCategoryService() {
	var (
		// 服务索引
		index = constant.ElasticIndexArticleCategoryService
		// 文章服务配置项
		indexBody = entity.ArticleCategoryServiceIndexBody
		// 初始化数据函数
		callback func()
	)

	callback = func() {
		// 初始化数据
		rows, _ := model.GetArticleCategoryRows()
		defer rows.Close()

		for rows.Next() {
			row := model.GetArticleCategoryScanRows(rows)
			// 文档生成
			document := entity.ToArticleCategoryService{row}.DocumentArticleCategoryService()
			// 创建文档
			_, err := elastic.GetClient().CreateDocument(index, strconv.Itoa(row.Id), document)
			if err != nil {
				logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
					"err":   err,
					"index": index,
					"id":    row.Id,
					"body":  document,
				}.Fields()).Error("elasticsearch add document err")
			}
		}
	}

	createIndexToInitDocument(index, indexBody, callback)
}

// 封装创建到初始化索引数据方法
func createIndexToInitDocument(index string, indexBody string, callback func()) {
	// 判断索引是否存在
	ok, _ := elastic.GetClient().IndexExists(index)
	if !ok {
		// 创建索引并设置配置项
		_, err := elastic.GetClient().CreateIndexToBodyString(index, indexBody)
		if err != nil {
			logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
				"index":            index,
				"indexDefaultBody": indexBody,
				"err":              err,
			}.Fields()).Error("elasticsearch init settings error")
			_, _ = elastic.GetClient().DeleteIndex(index)
			return
		}

		logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
			"index":            index,
			"indexDefaultBody": indexBody,
		}.Fields()).Info("elasticsearch init settings success")

		callback()
	}
}

