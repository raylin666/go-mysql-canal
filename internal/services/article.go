package services

import (
	"encoding/json"
	"go-mysql-canal/constant"
	"go-mysql-canal/entity"
	"go-mysql-canal/model"
	"go-mysql-canal/pkg/elastic"
	"strconv"
)

// 更新文章服务文档数据
func UpdateArticleServiceDocument(modelStruct interface{}, row interface{}) {
	switch modelStruct.(type) {
	case model.Article:
	case model.ArticleExtend:
	case model.ArticleCategoryRelation:
		// 模型类型映射
		rowModel := row.(model.ArticleCategoryRelation)
		// 查询分类对应的文章ID并更新文章文档对应的分类
		if exist, _ := elastic.GetClient().ExistsDocument(constant.ElasticIndexArticleService, strconv.Itoa(rowModel.ArticleId)); !exist {
			r := model.GetWithArticleRow(rowModel.ArticleId)
			// 文档生成
			document := entity.WithArticleService{r}.DocumentArticleService()
			// 创建文档
			_, err := elastic.GetClient().CreateDocument(constant.ElasticIndexArticleService, strconv.Itoa(r.Id), document)
			if err != nil {
				elasticLoggerWrite("updateCreate", constant.ElasticIndexArticleService, strconv.Itoa(r.Id), err, document)
			}
		} else {
			var entityModel entity.ArticleService
			entityDoc, _ := elastic.GetClient().GetDocument(constant.ElasticIndexArticleService, strconv.Itoa(rowModel.ArticleId))
			doc := entityModel
			// 解析出最终的索引文档内容
			docSource, _ := entityDoc.Source.MarshalJSON()
			_ = json.Unmarshal(docSource, &doc)
			if doc.Id > 0 {
				lists := model.GetRelationToCategoryRows(rowModel.ArticleId)
				var category = make([]entity.ArticleCategory, len(lists))
				for r := range lists {
					category[r] = entity.ArticleCategory{
						Id:        lists[r].Category.Id,
						Pid:       lists[r].Category.Pid,
						Name:      lists[r].Category.Name,
						Sort:      lists[r].Category.Sort,
						Status:    lists[r].Category.Status,
						CreatedAt: lists[r].Category.CreatedAt,
						UpdatedAt: lists[r].Category.UpdatedAt,
					}
				}
				doc.Category = category
				_, err := elastic.GetClient().UpdateDocument(constant.ElasticIndexArticleService, strconv.Itoa(rowModel.ArticleId), doc)
				elasticLoggerWrite("update", constant.ElasticIndexArticleService, strconv.Itoa(rowModel.ArticleId), err, doc)
			}
		}

		return
	}
}
