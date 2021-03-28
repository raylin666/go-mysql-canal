package handler

import (
	"fmt"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"go-mysql-canal/constant"
	"go-mysql-canal/pkg/logger"
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
		"库名":   ev.Table.Schema,
		"表名":   ev.Table.Name,
		"行为":   ev.Action,
		"数据记录": ev.Rows,
	}.Fields()).Info()

	// 此处是参考 https://github.com/gitstliu/MysqlToAll 里面的获取字段和值的方法
	var row = ""
	for columnIndex, currColumn := range ev.Table.Columns {
		// 字段名 | 字段的索引顺序 | 字段对应的值
		row += fmt.Sprintf("字段名: %v 字段值: %v 索引位置: %v | ", currColumn.Name, ev.Rows[len(ev.Rows)-1][columnIndex], columnIndex)
	}

	logger.NewWrite(constant.LOG_MULTI_SQL).Println(row)
	return nil
}

// 创建、更改、重命名或删除表时触发，通常会需要清除与表相关的数据，如缓存。It will be called before OnDDL.
func (h *EventHandler) OnTableChanged(schema string, table string) error {
	// 库名 ｜ 表名
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"库名": schema,
		"表名": table,
	}.Fields()).Info()
	return nil
}

// 监听 binlog 日志的变化文件与记录的位置
func (h *EventHandler) OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error {
	// 当 force 为 true，立即同步位置
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"binlog 文件名称": pos.Name,
		"pos 位置":      pos.Pos,
	}.Fields()).Info()
	return nil
}

// 当产生新的 binlog 日志后触发 (在达到内存的使用限制后（默认为 1GB），会开启另一个文件，每个新文件的名称后都会有一个增量)
func (h *EventHandler) OnRotate(r *replication.RotateEvent) error {
	// binlog 的记录位置，新 binlog 的文件名
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"下个 binlog 文件名称": r.NextLogName,
		"pos 位置":         r.Position,
	}.Fields()).Info()
	return nil
}

// create alter drop truncate (删除当前表再新建一个一模一样的表结构)
func (h *EventHandler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	// binlog 日志的变化文件与记录的位置
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"下个 binlog 文件名称": nextPos.Name,
		"下个 pos 位置":      nextPos.Pos,
		"执行时间":           queryEvent.ExecutionTime,
		"库名":             string(queryEvent.Schema),
		"变更的 sql 语句":     string(queryEvent.StatusVars[:]),
		"从库代理 ID":        queryEvent.SlaveProxyID,
	}.Fields()).Info()
	return nil
}

func (h *EventHandler) String() string {
	return "EventHandler"
}
