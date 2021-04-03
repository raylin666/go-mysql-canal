package constant

const (
	/** ---------------------------------------------------
		Elastic Index 索引列表
	  ---------------------------------------------------*/
	// 文章索引
	ElasticIndexArticleService = "my_server@article_service"
	// 文章分类索引
	ElasticIndexArticleCategoryService = "my_server@article_category_service"

	/** ---------------------------------------------------
		DB Table 事件变动表
	  ---------------------------------------------------*/
	DbTableArticle                 = "my_article"
	DbTableArticleExtend           = "my_article_extend"
	DbTableArticleCategory         = "my_article_category"
	DbTableArticleCategoryRelation = "my_article_category_relation"
)

var (
	// DB Table 事件变动表集合
	EventDbTables = []string{
		DbTableArticle,
		DbTableArticleExtend,
		DbTableArticleCategory,
		DbTableArticleCategoryRelation,
	}
)
