package worker

import "github.com/Deansquirrel/goMonitorV4/object"

func newIntWorker(intConfigData *object.IntConfigData) *intWorker {
	return &intWorker{
		intConfigData: intConfigData,
	}
}
