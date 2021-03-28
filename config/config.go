package config

import (
	"github.com/siddontang/go-mysql/canal"
	"math/rand"
	"time"
)

/**
	获取数据库配置信息
 */
func NewDatabaseConfig() *canal.Config {
	// 读取toml文件格式
	cfg, _ := canal.NewConfigWithFile("config/autoload/database.toml")
	if (cfg.ServerID == 0) {
		cfg.ServerID = uint32(rand.New(rand.NewSource(time.Now().Unix())).Intn(1000)) + 1001
	}
	return cfg;
}
