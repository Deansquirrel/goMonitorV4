package action

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IActionRepository interface {
	GetAction(id string) (object.IActionData, error)
}

func NewActionRepository(actionType global.ActionType) (IActionRepository, error) {
	switch actionType {
	case global.AWindowsService:
		return newActionRepository(&windowsServiceAction{}), nil
	case global.AIISAppPool:
		return newActionRepository(&iisAppPoolAction{}), nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的actionType：%d", actionType))
	}
}
