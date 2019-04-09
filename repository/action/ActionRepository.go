package action

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/common"
)

import log "github.com/Deansquirrel/goToolLog"

type actionRepository struct {
	iAction IAction
}

func newActionRepository(iAction IAction) *actionRepository {
	return &actionRepository{
		iAction: iAction,
	}
}

func (ar *actionRepository) GetAction(id string) (object.IActionData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(ar.iAction.GetSqlActionById(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return ar.iAction.GetActionListByRows(rows)
}
