package action

import (
	"errors"
	"github.com/Deansquirrel/goMonitorV4/object"
	"reflect"
)

type action struct {
	iConfig object.IActionData
	iAction IAction
}

func NewAction(iConfig object.IActionData) (*action, error) {
	ac, err := getAction(iConfig)
	if err != nil {
		return nil, err
	}
	return &action{
		iConfig: iConfig,
		iAction: ac,
	}, nil
}

func getAction(iConfig object.IActionData) (IAction, error) {
	switch reflect.TypeOf(iConfig).String() {
	case "*object.WindowsServiceActionData":
		return &windowsServiceAction{iConfig.(*object.WindowsServiceActionData)}, nil
	case "*object.IISAppPoolActionData":
		return &iisAppPoolAction{iConfig.(*object.IISAppPoolActionData)}, nil
	default:
		return nil, errors.New("未预知的配置类型:" + reflect.TypeOf(iConfig).String())
	}
}

func (ac *action) Do() error {
	b, err := ac.iAction.CheckTimes()
	if err != nil {
		return err
	}
	if b {
		return ac.iAction.Do()
	}
	return nil
}
