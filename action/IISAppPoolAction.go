package action

import "github.com/Deansquirrel/goMonitorV4/object"

type iisAppPoolAction struct {
	iConfig *object.IISAppPoolActionData
}

func (ia *iisAppPoolAction) CheckTimes() (bool, error) {
	//TODO
	return false, nil
}

func (ia *iisAppPoolAction) GetConfigData() object.IActionData {
	//TODO
	return nil
}

func (ia *iisAppPoolAction) Do() error {
	//TODO
	return nil
}
