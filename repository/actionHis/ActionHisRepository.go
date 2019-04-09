package actionHis

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"time"
)

type actionHisRepository struct {
	ActionHis IActionHis
}

func (ah *actionHisRepository) GetTopHis(configId string) ([]object.IHisData, error) {
	//ah.ActionHis.GetSqlTopHisList()
	return nil, nil
}

func (ah *actionHisRepository) SetHis(data object.IHisData) error {
	return nil
}

func (ah *actionHisRepository) ClearHis(t time.Duration) error {
	return nil
}
