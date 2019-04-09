package configHis

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IHis interface {
	GetSqlHisList() string
	GetSqlHisById() string
	GetSqlHisByConfigId() string
	GetSqlHisByTime() string
	GetSqlSetHis() string
	GetSqlClearHis() string

	GetHisListByRows(rows *sql.Rows) ([]object.IHisData, error)
	GetHisSetArgs(data object.IHisData) ([]interface{}, error)
}
