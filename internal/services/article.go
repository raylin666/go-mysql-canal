package services

import (
	"go-mysql-canal/constant"
	"go-mysql-canal/entity"
	"go-mysql-canal/model"
	"go-mysql-canal/pkg/elastic"
	"go-mysql-canal/pkg/logger"
	"go-mysql-canal/pkg/utils"
	"strconv"
)

// 操作文章
func ActionArticle(article model.Article) {
	var (
		isDeleteIndex = false
	)

	// 索引名称
	index := constant.ElasticIndexArticleService
	// 判断索引是否存在
	ok, _ := elastic.GetClient().IndexExists(index)
	if !ok {
		// 创建索引并设置配置项
		_, err := elastic.GetClient().CreateIndexToBodyString(index, entity.ArticleServiceIndexDefaultBody)
		if err != nil {
			isDeleteIndex = true
		}

		if isDeleteIndex {
			_, _ = elastic.GetClient().DeleteIndex(index)
			return
		}

		// 初始化索引内容
		rows, _ := model.GetArticleRows()
		defer rows.Close()

		for rows.Next() {
			articleFirst := model.GetArticleScanRows(rows)
			CreateArticleDocument(articleFirst)
		}

		logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
			"index": index,
		}.Fields()).Info("elasticsearch init success")

	} else {
		// 更新数据
		UpdateArticleDocument(article)
	}
}

// 新增文章文档
func CreateArticleDocument(article model.Article) {
	// 索引名称
	index := constant.ElasticIndexArticleService

	document := entity.ArticleService{
		Article: entity.Article{
			Id:                article.Id,
			Title:             article.Title,
			Author:            article.Author,
			Summary:           article.Summary,
			Cover:             article.Cover,
			Sort:              article.Sort,
			RecommendFlag:     utils.Int8ToBoolean(article.RecommendFlag),
			CommentFlag:       utils.Int8ToBoolean(article.CommentFlag),
			Status:            utils.Int8ToBoolean(article.Status),
			ViewCount:         article.ViewCount,
			CommentCount:      article.CommentCount,
			ShareCount:        article.ShareCount,
			PublisherUsername: article.PublisherUsername,
			UserId:            article.UserId,
			LastCommentTime:   article.LastCommentTime,
			CreatedAt:         article.CreatedAt,
			UpdatedAt:         article.UpdatedAt,
		},
	}

	_, err := elastic.GetClient().CreateDocument(index, strconv.Itoa(article.Id), document)
	if err != nil {
		logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
			"err":   err,
			"index": index,
			"id":    article.Id,
			"body":  document,
		}.Fields()).Error("elasticsearch add document err")
	}
}

// 更新文章文档
func UpdateArticleDocument(article model.Article) {
	// 索引名称
	index := constant.ElasticIndexArticleService

	switch article.Status {
	case 0:
		if _, err := elastic.GetClient().DeleteDocument(index, strconv.Itoa(article.Id)); err != nil {
			logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
				"err":   err,
				"index": index,
				"id":    article.Id,
			}.Fields()).Error("elasticsearch delete document err")
		}
	case 1:
		document := entity.ArticleService{
			Article: entity.Article{
				Id:                article.Id,
				Title:             article.Title,
				Author:            article.Author,
				Summary:           article.Summary,
				Cover:             article.Cover,
				Sort:              article.Sort,
				RecommendFlag:     utils.Int8ToBoolean(article.RecommendFlag),
				CommentFlag:       utils.Int8ToBoolean(article.CommentFlag),
				Status:            utils.Int8ToBoolean(article.Status),
				ViewCount:         article.ViewCount,
				CommentCount:      article.CommentCount,
				ShareCount:        article.ShareCount,
				PublisherUsername: article.PublisherUsername,
				UserId:            article.UserId,
				LastCommentTime:   article.LastCommentTime,
				CreatedAt:         article.CreatedAt,
				UpdatedAt:         article.UpdatedAt,
			},
		}

		if r, _ := elastic.GetClient().GetDocument(index, strconv.Itoa(article.Id)); r != nil {
			_, err := elastic.GetClient().UpdateDocument(index, strconv.Itoa(article.Id), document)
			if err != nil {
				logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
					"err":   err,
					"index": index,
					"id":    article.Id,
					"body":  document,
				}.Fields()).Error("elasticsearch update document err")
			}
		} else {
			_, err := elastic.GetClient().CreateDocument(index, strconv.Itoa(article.Id), document)
			if err != nil {
				logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
					"err":   err,
					"index": index,
					"id":    article.Id,
					"body":  document,
				}.Fields()).Error("elasticsearch add document err")
			}
		}
	}

	return
}
