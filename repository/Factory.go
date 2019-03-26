package repository

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/kataras/iris/core/errors"
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
	default:
		return nil, errors.New(fmt.Sprintf("未预知的ConfigType：%d", configType))
	}
}

func NewHisRepository(hisType global.HisType) (IHisRepository, error) {
	switch hisType {
	case global.HInt:
		return &hisRepository{
			His: &intHis{},
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("hisType：%d", hisType))
	}
}

func NewNotifyRepository(notifyType global.NotifyType) (INotifyRepository, error) {
	switch notifyType {
	case global.NDingTalkRobot:
		return &notifyRepository{
			Notify: &dingTalkRobotNotify{},
		}, nil
	}
	return nil, nil
}
