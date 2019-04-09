package notify

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type INotify interface {
	GetSqlGetConfig() string
	getConfigListByRows(rows *sql.Rows) ([]object.INotifyData, error)
}
