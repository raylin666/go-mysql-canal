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
