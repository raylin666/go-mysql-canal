package model

import (
	"database/sql"
	"go-mysql-canal/constant"
	"time"
)

type ArticleCategory struct {
	Id        int       `json:"id" gorm:"primarykey;unique" mapstructure:"id"`
	Pid       int       `json:"pid" mapstructure:"pid"`
	Name      string    `json:"name" mapstructure:"name"`
	Sort      int       `json:"sort" mapstructure:"sort"`
	Status    bool      `json:"status" mapstructure:"status"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
}

type ArticleCategoryTable struct {}

func (t *ArticleCategoryTable) TableName() string {
	return constant.DbTableArticleCategory
}

func GetArticleCategoryRows() (*sql.Rows, error) {
	return GetMyServerDB().Model(ArticleCategory{}).Where("status = ?", 1).Where("deleted_at is null").Rows()
}

func GetArticleCategoryScanRows(row *sql.Rows) (article_category ArticleCategory) {
	_ = GetMyServerDB().ScanRows(row, &article_category)
	return
}

