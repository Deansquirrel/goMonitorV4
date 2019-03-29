package repository

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

func NewHisRepository(hisType global.HisType) (IHisRepository, error) {
	switch hisType {
	case global.HInt:
		return &hisRepository{
			His: &intHis{},
		}, nil
	case global.HWebState:
		return &hisRepository{
			His: &webStateHis{},
		}, nil
	case global.HCrmDzXfTest:
		return &hisRepository{
			His: &crmDzXfTestHis{},
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的HisType：%d", hisType))
	}
}

func NewNotifyRepository(notifyType global.NotifyType) (INotifyRepository, error) {
	switch notifyType {
	case global.NDingTalkRobot:
		return &notifyRepository{
			Notify: &dingTalkRobotNotify{},
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的NotifyType：%d", notifyType))
	}
}

func NewActionRepository(actionType global.ActionType) (IActionRepository, error) {
	switch actionType {
	case global.AWindowsService:
		return &actionRepository{
			Action: &windowsServiceAction{},
		}, nil
	case global.AIISAppPool:
		return &actionRepository{
			Action: &iisAppPoolAction{},
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的actionType：%d", actionType))
	}
}
