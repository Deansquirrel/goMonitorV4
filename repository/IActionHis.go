package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IActionHis interface {
	GetSqlTopHisList() string
	GetSqlSetHis() string
	GetSqlClearHis() string

	getHisListByRows(rows *sql.Rows) ([]object.IHisData, error)
	getHisSetArgs(data object.IHisData) ([]interface{}, error)
}
