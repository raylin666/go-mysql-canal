package main

import (
	"go-mysql-canal/config"
	"go-mysql-canal/constant"
	"go-mysql-canal/internal/services"
	"go-mysql-canal/pkg/database"
	"go-mysql-canal/pkg/elastic"
	"go-mysql-canal/pkg/logger"
	"go-mysql-canal/server"
)

func init()  {
	config.NewConfig()
	logger.InitLogger()
	database.InitDatabase()
}

func main() {
	elastic.NewClient()

	// 用来测试时实时删除索引, 上线将删除
	_, _ = elastic.GetClient().DeleteIndex(constant.ElasticIndexArticleService)
	_, _ = elastic.GetClient().DeleteIndex(constant.ElasticIndexArticleCategoryService)

	// 初始化创建\同步数据到 Elastic
	services.InitElasticDataSync()

	// canal 服务启动
	err := server.NewCanal(true);
	if err != nil {
		panic(err)
	}
}
