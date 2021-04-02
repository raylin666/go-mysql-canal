package database

import (
	"fmt"
	"go-mysql-canal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var database map[string]*gorm.DB

func InitDatabase() {
	var (
		err  error
		conn *gorm.DB
	)

	c := config.DatabaseConfig

	database = make(map[string]*gorm.DB, 1)

	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Addr,
		c.Dump.TableDB,
		c.Charset)

	conn, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "my_",				// 表前缀
				SingularTable: true,			// 全局禁用表名复数
			},
		})

	if err != nil {
		panic(err)
	}

	// 使用插件
	_ = conn.Use(&TracePlugin{})

	database[c.Dump.TableDB] = conn
}

func GetDB(connection string) *gorm.DB {
	return database[connection]
}

func Close(connection string) error {
	sqlDb, _ := database[connection].DB()
	return sqlDb.Close()
}

func CloseAll() error {
	for _, connection := range database {
		sqlDb, _ := connection.DB()
		if err := sqlDb.Close(); err != nil {
			return err
		}
	}

	return nil
}
