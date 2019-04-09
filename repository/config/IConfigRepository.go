package config

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IConfigRepository interface {
	GetConfigList() ([]object.IConfigData, error)
	GetConfig(id string) (object.IConfigData, error)
}

func NewConfigRepository(configType global.ConfigType) (IConfigRepository, error) {
	switch configType {
	case global.CM:
		return newConfigRepository(&mConfig{}), nil
	case global.CInt:
		return newConfigRepository(&intConfig{}), nil
	case global.CIntD:
		return newConfigRepository(&intDConfig{}), nil
	case global.CHealth:
		return newConfigRepository(&healthConfig{}), nil
	case global.CCrmDzXfTest:
		return newConfigRepository(&crmDzXfTestConfig{}), nil
	case global.CWebState:
		return newConfigRepository(&webStateConfig{}), nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的ConfigType：%d", configType))
	}
}
