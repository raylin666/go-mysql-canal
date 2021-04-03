package model

type ArticleCategoryRelation struct {
	Id         int    `json:"id" gorm:"primarykey;unique" mapstructure:"id"`
	ArticleId  int    `json:"article_id" mapstructure:"article_id"`
	CategoryId int    `json:"category_id" mapstructure:"category_id"`
}

type ArticleCategoryRelationTable struct{}

func (t *ArticleCategoryRelationTable) TableName() string {
	return "my_article_category_relation"
}

type WithArticleCategory struct {
	ArticleCategoryRelation
	ArticleCategoryRelationTable
	Category struct{
		ArticleCategory
		ArticleCategoryTable
	}  `gorm:"foreignkey:category_id"`
}

func GetRelationToArticleIdRows(category_id int) (lists []int) {
	GetMyServerDB().Model(&ArticleCategoryRelation{}).
		Where("category_id = ?", category_id).
		Pluck("article_id", &lists)
	return
}

func GetRelationToCategoryRows(article_id int) (lists []WithArticleCategory) {
	GetMyServerDB().Model(&WithArticleCategory{}).
		Where("article_id = ?", article_id).
		Preload("Category").
		Find(&lists)
	return
}