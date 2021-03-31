package model

import (
	"database/sql"
	"go-mysql-canal/pkg/database"
	"gorm.io/gorm"
)

type Article struct {
	Id                int    `json:"id" gorm:"primarykey;unique" mapstructure:"id"`
	Title             string `json:"title" mapstructure:"title"`
	Author            string `json:"author" mapstructure:"author"`
	Summary           string `json:"summary" mapstructure:"summary"`
	Cover             string `json:"cover" mapstructure:"cover"`
	Sort              int    `json:"sort" mapstructure:"sort"`
	RecommendFlag     int8   `json:"recommend_flag" mapstructure:"recommend_flag"`
	CommentFlag       int8   `json:"comment_flag" mapstructure:"comment_flag"`
	Status            int8   `json:"status" mapstructure:"status"`
	ViewCount         int    `json:"view_count" mapstructure:"view_count"`
	CommentCount      int    `json:"comment_count" mapstructure:"comment_count"`
	ShareCount        int    `json:"share_count" mapstructure:"share_count"`
	PublisherUsername string `json:"publisher_username" mapstructure:"publisher_username"`
	UserId            int    `json:"user_id" mapstructure:"user_id"`
	LastCommentTime   int    `json:"last_comment_time" mapstructure:"last_comment_time"`
	CreatedAt         string `json:"created_at" mapstructure:"created_at"`
	UpdatedAt         string `json:"updated_at" mapstructure:"updated_at"`
}

func getDB() *gorm.DB {
	return database.GetDB("default")
}

func GetArticleRows() (*sql.Rows, error) {
	return getDB().Model(Article{}).Where("status = ?", 1).Where("deleted_at is null").Rows()
}

func GetArticleScanRows(rows *sql.Rows) (article Article) {
	_ = getDB().ScanRows(rows, &article)
	return
}

func CloseArticleStatus(id int) *gorm.DB {
	return getDB().Model(Article{}).Where("id = ?", id).Update("status", 0)
}

func GetArticleById(id int) *gorm.DB {
	return getDB().Where("id = ?", id).First(&Article{})
}