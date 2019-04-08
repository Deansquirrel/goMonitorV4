package repositoryConfig

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
)

func NewConfigRepository(configType global.ConfigType) (IConfigRepository, error) {
	switch configType {
	case global.CM:
		return &configRepository{
			Config: &mConfig{},
		}, nil
	case global.CInt:
		return &configRepository{
			Config: &intConfig{},
		}, nil
	case global.CIntD:
		return &configRepository{
			Config: &intDConfig{},
		}, nil
	case global.CHealth:
		return &configRepository{
			Config: &healthConfig{},
		}, nil
	case global.CCrmDzXfTest:
		return &configRepository{
			Config: &crmDzXfTestConfig{},
		}, nil
	case global.CWebState:
		return &configRepository{
			Config: &webStateConfig{},
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的ConfigType：%d", configType))
	}
}
