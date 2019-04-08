package action

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
	"github.com/Deansquirrel/goToolCommon"
	"strings"
	"time"
)

type iisAppPoolAction struct {
	iConfig *object.IISAppPoolActionData
}

func (ia *iisAppPoolAction) Do() (object.IHisData, error) {
	//TODO
	return nil, nil
}
func (ia *iisAppPoolAction) GetHisData(msg string) object.IHisData {
	return object.IISAppPoolActionHisData{
		FId:            strings.ToUpper(goToolCommon.Guid()),
		FConfigId:      ia.iConfig.FId,
		FAgentServerId: ia.iConfig.FAgentServerId,
		FName:          ia.iConfig.FName,
		FCheckTimes:    ia.iConfig.FCheckTimes,
		FContent:       msg,
		FOprTime:       time.Now(),
	}
}
func (ia *iisAppPoolAction) CheckAction() (bool, error) {
	//TODO
	return false, nil
}
func (ia *iisAppPoolAction) GetHisRepository() (repository.IHisRepository, error) {
	//TODO
	return nil, nil
}
