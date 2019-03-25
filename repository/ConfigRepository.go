package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/object"
	log "github.com/Deansquirrel/goToolLog"
)

type configRepository struct {
	Config IConfig
}

func newConfigRepository(config IConfig) *configRepository {
	return &configRepository{
		Config: config,
	}
}

func (cr *configRepository) GetConfigList() ([]object.IConfigData, error) {
	rows, err := comm.getRowsBySQL(cr.Config.GetSqlGetConfigList())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return cr.Config.getConfigListByRows(rows)
}

func (cr *configRepository) GetConfig(id string) (object.IConfigData, error) {
	rows, err := comm.getRowsBySQL(cr.Config.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	list, err := cr.Config.getConfigListByRows(rows)
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
