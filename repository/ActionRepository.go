package repository

import "github.com/Deansquirrel/goMonitorV4/object"

import log "github.com/Deansquirrel/goToolLog"

type actionRepository struct {
	Action IAction
}

func (ar *actionRepository) GetAction(id string) (object.IActionData, error) {
	comm := common{}
	rows, err := comm.getRowsBySQL(ar.Action.GetSqlActionById(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return ar.Action.getActionListByRows(rows)
}
