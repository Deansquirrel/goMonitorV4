package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IConfig interface {
	GetSqlGetConfigList() string
	GetSqlGetConfig() string

	getConfigListByRows(rows *sql.Rows) ([]object.IConfigData, error)
}
