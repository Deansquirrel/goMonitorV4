package action

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
)

import repAction "github.com/Deansquirrel/goMonitorV4/repository/action"

import log "github.com/Deansquirrel/goToolLog"

func CheckAction(id string) error {
	actionListRepository := repAction.NewActionList()
	actionList, err := actionListRepository.GetActionList(id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var errMsg string
	for _, s := range actionList.WindowsService {
		err = checkActionWorker(s, global.AWindowsService)
		if err != nil {
			errMsg = errMsg + err.Error() + ";"
		}
	}

	for _, s := range actionList.IISAppPool {
		err = checkActionWorker(s, global.AIISAppPool)
		if err != nil {
			errMsg = errMsg + err.Error() + ";"
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func checkActionWorker(id string, t global.ActionType) error {
	log.Debug(fmt.Sprintf("Action ID: %s，Action Type: %d", id, t))
	rep, err := repAction.NewActionRepository(t)
	if err != nil {
		return err
	}

	actionData, err := rep.GetAction(id)
	if err != nil {
		return err
	}
	ac, err := newAction(actionData)
	if err != nil {
		return err
	}
	return ac.Do()
}
