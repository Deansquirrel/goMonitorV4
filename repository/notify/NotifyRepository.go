package notify

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/common"
	log "github.com/Deansquirrel/goToolLog"
)

type notifyRepository struct {
	iNotify INotify
}

func newNotifyRepository(iNotify INotify) *notifyRepository {
	return &notifyRepository{
		iNotify: iNotify,
	}
}

func (nr *notifyRepository) GetNotify(id string) (object.INotifyData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(nr.iNotify.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	list, err := nr.iNotify.getConfigListByRows(rows)
	if err != nil {
		return nil, err
	}
	if len(list) < 1 {
		return nil, nil
	}
	if len(list) > 1 {
		return nil, errors.New(fmt.Sprintf("Config列表数量异常，exp：1，act:%d", len(list)))
	}
	return list[0], nil
}
