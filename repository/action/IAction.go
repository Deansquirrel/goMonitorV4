package action

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IAction interface {
	GetSqlActionById() string

	GetActionListByRows(rows *sql.Rows) (object.IActionData, error)
}
