package config

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/common"
	log "github.com/Deansquirrel/goToolLog"
)

type configRepository struct {
	iConfig IConfig
}

func newConfigRepository(iConfig IConfig) *configRepository {
	return &configRepository{
		iConfig: iConfig,
	}
}

func (cr *configRepository) GetConfigList() ([]object.IConfigData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(cr.iConfig.GetSqlGetConfigList())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return cr.iConfig.GetConfigListByRows(rows)
}

func (cr *configRepository) GetConfig(id string) (object.IConfigData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(cr.iConfig.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	list, err := cr.iConfig.GetConfigListByRows(rows)
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
