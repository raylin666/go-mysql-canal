package handler

import (
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"go-mysql-canal/constant"
	"go-mysql-canal/event"
	"go-mysql-canal/pkg/logger"
	"go-mysql-canal/pkg/utils"
)

/**
监听处理器
*/
type EventHandler struct {
	canal.DummyEventHandler
}

// 监听数据记录
func (h *EventHandler) OnRow(canalRowsEvent *canal.RowsEvent) error {
	// 库名 | 表名 | 行为 | 数据记录
	logger.NewWrite(constant.LOG_MULTI_SQL).WithFields(logger.Fields{
		"db":     canalRowsEvent.Table.Schema,
		"table":  canalRowsEvent.Table.Name,
		"action": canalRowsEvent.Action,
		"rows":   canalRowsEvent.Rows,
	}.Fields()).Info()

	// 解析数据行值并组合为待操作数据
	_, rows := event.GetParseValue(canalRowsEvent)

	// 相关表的操作将触发, 否则不处理
	if _, ok := utils.InArray(canalRowsEvent.Table.Name, constant.EventDbTables); ok {
		event.TableEventDispatcher(canalRowsEvent, rows)
	}

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

