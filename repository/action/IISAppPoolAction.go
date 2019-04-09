package action

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type iisAppPoolAction struct {
}

func (ia *iisAppPoolAction) GetSqlActionById() string {
	//TODO
	return ""
}

func (ia *iisAppPoolAction) GetActionListByRows(rows *sql.Rows) (object.IActionData, error) {
	//TODO
	return nil, nil
}
