package main

import (
	"go-mysql-canal/pkg/logger"
	"go-mysql-canal/server"
)

func init()  {
	logger.InitLogger()
}

func main() {
	// canal 服务启动
	err := server.NewCanal(true);
	if err != nil {
		panic(err)
	}
}
