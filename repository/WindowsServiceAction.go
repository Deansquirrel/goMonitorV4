package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type windowsServiceAction struct {
}

func (wsa *windowsServiceAction) GetSqlActionById() string {
	//TODO
	return ""
}

func (wsa *windowsServiceAction) getActionListByRows(rows *sql.Rows) (object.IActionData, error) {
	//TODO
	return nil, nil
}
