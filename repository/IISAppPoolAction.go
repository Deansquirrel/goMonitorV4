package repository

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

func (ia *iisAppPoolAction) getActionListByRows(rows *sql.Rows) (object.IActionData, error) {
	//TODO
	return nil, nil
}
