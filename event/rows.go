package event

import (
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/siddontang/go-mysql/canal"
	"go-mysql-canal/constant"
	"go-mysql-canal/internal/services"
	"go-mysql-canal/model"
	"go-mysql-canal/pkg/logger"
	"go-mysql-canal/pkg/utils"
	"time"
)

// 解析数据行值并组合为待操作数据
func GetParseValue(canalRowsEvent *canal.RowsEvent) (*canal.RowsEvent, map[string]interface{}) {
	var (
		rows = map[string]interface{}{}
	)

	for columnIndex, currColumn := range canalRowsEvent.Table.Columns {
		columnValue := canalRowsEvent.Rows[len(canalRowsEvent.Rows)-1][columnIndex]

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
			case "created_at", "updated_at", "deleted_at":
				if columnValue.(string) != "" {
					columnValue, _ = time.Parse("2006-01-02 15:04:05.000000", columnValue.(string))
				}
			}
		}

		rows[currColumn.Name] = columnValue
	}

	return canalRowsEvent, rows
}

// 相关表数据变更事件分发
func TableEventDispatcher(canalRowsEvent *canal.RowsEvent, rows map[string]interface{}) {
	// 行数据校验
	modelStruct := model.GetModelStruct(canalRowsEvent.Table.Name)
	ok := MapstructureRows(canalRowsEvent, modelStruct, rows)
	if !ok {
		return
	}

	// 数据变更行为操作
	switch canalRowsEvent.Action {
	case canal.InsertAction:
	case canal.UpdateAction:
		switch canalRowsEvent.Table.Name {
		case constant.DbTableArticleCategory:
			services.UpdateArticleServiceDocument(modelStruct, rows)
		case constant.DbTableArticle, constant.DbTableArticleExtend:
			services.UpdateArticleServiceDocument(modelStruct, rows)
		}
	case canal.DeleteAction:
	}
}

// Mapstructure 行数据校验
func MapstructureRows(canalRowsEvent *canal.RowsEvent, model interface{}, rows map[string]interface{}) bool {
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &model,
	})
	err := decoder.Decode(rows)
	if err != nil {
		logger.NewWrite(constant.LOG_MULTI_ELASTIC).WithFields(logger.Fields{
			"err":   err,
			"table": canalRowsEvent.Table.Name,
			"value": rows,
		}.Fields()).Error("mapstructure decode to struct err")
		return false
	}

	return true
}