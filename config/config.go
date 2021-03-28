package config

import (
	"github.com/BurntSushi/toml"
	"github.com/siddontang/go-mysql/canal"
	"io/ioutil"
	"math/rand"
	"time"
)

var (
	DatabaseConfig *canal.Config
	ElasticConfig *EsConfig
)

type EsConfig struct {
	Urls string `toml:"urls"`
}

func NewConfig()  {
	newElasticConfig()
	newDatabaseConfig()
}

func newElasticConfig() {
	data, err :=ioutil.ReadFile("config/autoload/elastic.toml")
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(data), &ElasticConfig)
	if err != nil {
		panic(err)
	}
}

/**
	获取数据库配置信息
 */
func newDatabaseConfig() {
	// 读取toml文件格式
	DatabaseConfig, _ = canal.NewConfigWithFile("config/autoload/database.toml")
	if (DatabaseConfig.ServerID == 0) {
		DatabaseConfig.ServerID = uint32(rand.New(rand.NewSource(time.Now().Unix())).Intn(1000)) + 1001
	}
}
