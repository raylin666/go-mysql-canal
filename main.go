package main

import (
	"go-mysql-canal/config"
	"go-mysql-canal/pkg/elastic"
	"go-mysql-canal/pkg/logger"
	"go-mysql-canal/server"
)

func init()  {
	config.NewConfig()
	logger.InitLogger()
}

func main() {
	elastic.NewClient()

	// canal 服务启动
	err := server.NewCanal(true);
	if err != nil {
		panic(err)
	}
}
