package worker

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/configHis"
)

type IWorker interface {
	GetMsg() (string, object.IHisData)
	//SaveSearchResult(data object.IHisData) error

	getHisRepository() (configHis.IHisRepository, error)
	formatMsg(msg string) string
}
