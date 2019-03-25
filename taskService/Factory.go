package taskService

import "github.com/Deansquirrel/goMonitorV4/global"

func NewIntTask() *task {
	return newTask(&intTask{}, global.CInt, global.HInt)
}
