package action

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/configHis"
)

type windowsServiceAction struct {
	iConfig *object.WindowsServiceActionData
}

func (ws *windowsServiceAction) Do() (object.IHisData, error) {
	//TODO
	return nil, nil
}
func (ws *windowsServiceAction) GetHisData(msg string) object.IHisData {
	//TODO
	return nil
}
func (ws *windowsServiceAction) CheckAction() (bool, error) {
	//TODO
	return false, nil
}
func (ws *windowsServiceAction) GetHisRepository() (configHis.IHisRepository, error) {
	//TODO
	return nil, nil
}
