package main

import (
	"go-mysql-canal/config"
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

	// 初始化创建\同步数据到 Elastic
	services.InitElasticDataSync()

	// canal 服务启动
	err := server.NewCanal(true);
	if err != nil {
		panic(err)
	}
}
