package action

import "github.com/Deansquirrel/goMonitorV4/object"

type IAction interface {
	CheckTimes() (bool, error)
	GetConfigData() object.IActionData
	Do() error
}
