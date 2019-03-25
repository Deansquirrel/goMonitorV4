package worker

import (
	"github.com/Deansquirrel/goMonitorV4/object"
)

type IWorker interface {
	GetMsg() (string, object.IHisData)
	SaveSearchResult(data object.IHisData) error

	formatMsg(msg string) string
}
