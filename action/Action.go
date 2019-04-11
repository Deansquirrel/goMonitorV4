package action

import (
	"errors"
	"github.com/Deansquirrel/goMonitorV4/object"
	"reflect"
)

import log "github.com/Deansquirrel/goToolLog"

type action struct {
	iConfig object.IActionData
	iAction IAction
}

func newAction(iConfig object.IActionData) (*action, error) {
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
	b, err := ac.iAction.CheckAction()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if !b {
		return ac.setHIsData("check action false")
	}
	hisData, err := ac.iAction.Do()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = ac.setHIsData(hisData)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (ac *action) setHIsData(data object.IHisData) error {
	rep, err := ac.iAction.GetHisRepository()
	if err != nil {
		return err
	}
	return rep.SetHis(data)
}
