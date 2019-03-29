package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goToolCommon"
	"strings"
)

const SqlGetActionList = "" +
	"SELECT WindowsService, IISAppPool" +
	" FROM ActionList" +
	" WHERE TaskId = ?"

type ActionList struct {
}

type ActionListData struct {
	WindowsService []string
	IISAppPool     []string
}

func (al *ActionList) GetActionList(id string) (*ActionListData, error) {
	comm := common{}
	rows, err := comm.getRowsBySQL(SqlGetActionList, id)
	if err != nil {
		return nil, err
	}
	return al.getActionListByRows(rows)
}

func (al *ActionList) getActionListByRows(rows *sql.Rows) (*ActionListData, error) {
	defer func() {
		_ = rows.Close()
	}()

	var windowsService, iisAppPool sql.NullString
	var windowsServiceList, iisAppPoolList []string
	windowsServiceList = make([]string, 0)
	iisAppPoolList = make([]string, 0)
	for rows.Next() {
		err := rows.Scan(&windowsService, &iisAppPool)
		if err != nil {
			return nil, err
		}
		if windowsService.Valid {
			list := strings.Split(windowsService.String, ",")
			for _, s := range list {
				windowsServiceList = append(windowsServiceList, s)
			}
		}
		if iisAppPool.Valid {
			list := strings.Split(iisAppPool.String, ",")
			for _, s := range list {
				iisAppPoolList = append(iisAppPoolList, s)
			}
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &ActionListData{
		WindowsService: goToolCommon.ClearBlock(windowsServiceList),
		IISAppPool:     goToolCommon.ClearBlock(iisAppPoolList),
	}, nil
}
