package handler

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"go-mysql-canal/constant"
	"go-mysql-canal/internal/services"
	"go-mysql-canal/model"
	"go-mysql-canal/pkg/logger"
	"go-mysql-canal/pkg/utils"
	"time"
)

const (
	TableArticle = "my_article"
)

/**
监听处理器
*/
type EventHandler struct {
	canal.DummyEventHandler
}

// 监听数据记录
func (h *EventHandler) OnRow(ev *canal.RowsEvent) error {
	// 库名 | 表名 | 行为 | 数据记录
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"db":     ev.Table.Schema,
		"table":  ev.Table.Name,
		"action": ev.Action,
		"rows":   ev.Rows,
	}.Fields()).Info()

	// 此处是参考 https://github.com/gitstliu/MysqlToAll 里面的获取字段和值的方法
	var (
		row         = ""
		valueMap    = map[string]interface{}{}
		valueStruct interface{}
		f           func()
	)

	for columnIndex, currColumn := range ev.Table.Columns {
		columnValue := ev.Rows[len(ev.Rows)-1][columnIndex]
		// 字段名 | 字段的索引顺序 | 字段对应的值
		row += fmt.Sprintf("field: %v value: %v postion: %v | ", currColumn.Name, columnValue, columnIndex)

		switch columnValue.(type) {
		// 解析text字段，转string
		case []uint8:
			columnValue = utils.Uint8ToString(columnValue.([]uint8)...)
		// 解析decimal字段
		case decimal.Decimal:
			v, _ := columnValue.(decimal.Decimal).Float64()
			columnValue = float32(v)
		case string:
			switch currColumn.Name {
			case "created_at", "updated_at":
				columnValue, _ = time.Parse("2006-01-02 15:04:05.000000", columnValue.(string))
			}
		}

		valueMap[currColumn.Name] = columnValue
	}

	switch ev.Action {
	case canal.InsertAction, canal.UpdateAction:
		switch ev.Table.Name {
		case TableArticle:
			valueStruct = model.Article{}
			f = func() {
				services.ActionArticle(valueStruct.(model.Article))
			}
		}

		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			WeaklyTypedInput: true,
			Result:           &valueStruct,
		})
		err := decoder.Decode(valueMap)
		if err != nil {
			logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
				"err":   err,
				"table": ev.Table.Name,
				"value": valueMap,
			}.Fields()).Error("mapstructure decode to struct err")
		}

		if f != nil {
			f()
		}
	case canal.DeleteAction:
	}

	logger.NewWrite(constant.LOG_MULTI_SQL).Println(row)
	return nil
}

// 创建、更改、重命名或删除表时触发，通常会需要清除与表相关的数据，如缓存。It will be called before OnDDL.
func (h *EventHandler) OnTableChanged(schema string, table string) error {
	// 库名 ｜ 表名
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"db":    schema,
		"table": table,
	}.Fields()).Info()
	return nil
}

// 监听 binlog 日志的变化文件与记录的位置
func (h *EventHandler) OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error {
	// 当 force 为 true，立即同步位置
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"binlog filename": pos.Name,
		"pos":             pos.Pos,
	}.Fields()).Info()
	return nil
}

// 当产生新的 binlog 日志后触发 (在达到内存的使用限制后（默认为 1GB），会开启另一个文件，每个新文件的名称后都会有一个增量)
func (h *EventHandler) OnRotate(r *replication.RotateEvent) error {
	// binlog 的记录位置，新 binlog 的文件名
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"next binlog filename": r.NextLogName,
		"pos":                  r.Position,
	}.Fields()).Info()
	return nil
}

// create alter drop truncate (删除当前表再新建一个一模一样的表结构)
func (h *EventHandler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	// binlog 日志的变化文件与记录的位置
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"next binlog filename": nextPos.Name,
		"next pos":             nextPos.Pos,
		"execution time":       queryEvent.ExecutionTime,
		"db":                   string(queryEvent.Schema),
		"event sql query":      string(queryEvent.StatusVars[:]),
		"slave proxy ID":       queryEvent.SlaveProxyID,
	}.Fields()).Info()
	return nil
}

func (h *EventHandler) String() string {
	return "EventHandler"
}
