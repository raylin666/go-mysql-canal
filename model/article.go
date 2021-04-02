package model

import (
	"go-mysql-canal/constant"
	"time"
)

type Article struct {
	Id                int       `json:"id" gorm:"primarykey;unique" mapstructure:"id"`
	Title             string    `json:"title" mapstructure:"title"`
	Author            string    `json:"author" mapstructure:"author"`
	Summary           string    `json:"summary" mapstructure:"summary"`
	Cover             string    `json:"cover" mapstructure:"cover"`
	Sort              int       `json:"sort" mapstructure:"sort"`
	RecommendFlag     int8      `json:"recommend_flag" mapstructure:"recommend_flag"`
	CommentFlag       int8      `json:"comment_flag" mapstructure:"comment_flag"`
	Status            int8      `json:"status" mapstructure:"status"`
	ViewCount         int       `json:"view_count" mapstructure:"view_count"`
	CommentCount      int       `json:"comment_count" mapstructure:"comment_count"`
	ShareCount        int       `json:"share_count" mapstructure:"share_count"`
	PublisherUsername string    `json:"publisher_username" mapstructure:"publisher_username"`
	UserId            int       `json:"user_id" mapstructure:"user_id"`
	LastCommentTime   int       `json:"last_comment_time" mapstructure:"last_comment_time"`
	CreatedAt         time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" mapstructure:"updated_at"`
}

type ArticleTable struct {}

func (a *ArticleTable) TableName() string {
	return constant.DbTableArticle
}

type WithArticle struct {
	Article
	ArticleTable
	ArticleExtend struct{
		ArticleExtend
		ArticleExtendTable
	} `gorm:"foreignkey:article_id"`
	ArticleCategoryRelation []struct{
		ArticleCategoryRelation
		ArticleCategoryRelationTable
		Category ArticleCategory `gorm:"foreignkey:id;references:category_id"`
	} `gorm:"foreignkey:article_id"`
}

func GetWithArticleRows() (lists []WithArticle) {
	GetMyServerDB().Model(WithArticle{}).
		Preload("ArticleExtend").
		Preload("ArticleCategoryRelation").
		Preload("ArticleCategoryRelation.Category").
		Where("status = ?", 1).
		Where("deleted_at is null").
		Find(&lists)
	return
}

