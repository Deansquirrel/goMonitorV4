package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/object"
	log "github.com/Deansquirrel/goToolLog"
)

type notifyRepository struct {
	Notify INotify
}

func newNotifyRepository(notify INotify) *notifyRepository {
	return &notifyRepository{
		Notify: notify,
	}
}

func (nr *notifyRepository) GetNotify(id string) (object.INotifyData, error) {
	rows, err := comm.getRowsBySQL(nr.Notify.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	list, err := nr.Notify.getConfigListByRows(rows)
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
