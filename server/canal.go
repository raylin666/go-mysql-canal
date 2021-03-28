package server

import (
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"go-mysql-canal/config"
	"go-mysql-canal/handler"
)

func getCanal() *canal.Canal {
	c, err := canal.NewCanal(config.DatabaseConfig)
	if err != nil {
		panic(err)
	}

	c.SetEventHandler(&handler.EventHandler{})
	return c;
}

/**
	创建 Canal , 从递增位置开始监听
*/
func NewCanal(isPos bool) error {
	c := getCanal()

	if isPos == false {
		// 从头开始监听
		return c.Run()
	}

	// 获取最新位置
	masterPos, _ := c.GetMasterPos()

	// 递增监听位置
	startPos := mysql.Position{Name:masterPos.Name, Pos:masterPos.Pos}

	// 根据位置监听
	return c.RunFrom(startPos)
}

