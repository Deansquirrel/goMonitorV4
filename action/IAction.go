package action

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/configHis"
)

type IAction interface {
	Do() (object.IHisData, error)
	GetHisData(msg string) object.IHisData
	CheckAction() (bool, error)
	GetHisRepository() (configHis.IHisRepository, error)
}
