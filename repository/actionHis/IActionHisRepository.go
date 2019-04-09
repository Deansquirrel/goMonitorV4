package actionHis

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"time"
)

type IActionHisRepository interface {
	GetTopHis(configId string) ([]object.IHisData, error)
	SetHis(data object.IHisData) error
	ClearHis(t time.Duration) error
}
