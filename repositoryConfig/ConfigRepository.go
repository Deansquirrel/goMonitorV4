package repositoryConfig

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
	"github.com/Deansquirrel/goMonitorV4/repositoryCommon"
	log "github.com/Deansquirrel/goToolLog"
)

type configRepository struct {
	Config IConfig
}

func (cr *configRepository) GetConfigList() ([]object.IConfigData, error) {
	comm := repositoryCommon.Common{}
	rows, err := comm.GetRowsBySQL(cr.Config.GetSqlGetConfigList())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return cr.Config.getConfigListByRows(rows)
}

func (cr *configRepository) GetConfig(id string) (object.IConfigData, error) {
	comm := common{}
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
