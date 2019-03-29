package action

import "github.com/Deansquirrel/goMonitorV4/object"

type windowsServiceAction struct {
	iConfig *object.WindowsServiceActionData
}

func (ws *windowsServiceAction) CheckTimes() (bool, error) {
	//TODO
	return false, nil
}

func (ws *windowsServiceAction) GetConfigData() object.IActionData {
	//TODO
	return nil
}

func (ws *windowsServiceAction) Do() error {
	//TODO
	return nil
}
