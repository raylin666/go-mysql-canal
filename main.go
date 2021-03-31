package main

import (
	"go-mysql-canal/config"
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
	// elastic.GetClient().DeleteIndex("my_server@article_service");

	// canal 服务启动
	err := server.NewCanal(true);
	if err != nil {
		panic(err)
	}
}
