package worker

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
)

type IWorker interface {
	GetMsg() (string, object.IHisData)
	//SaveSearchResult(data object.IHisData) error

	getHisRepository() (repository.IHisRepository, error)
	formatMsg(msg string) string
}
