package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IAction interface {
	GetSqlActionById() string

	getActionListByRows(rows *sql.Rows) (object.IActionData, error)
}
