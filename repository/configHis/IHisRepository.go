package configHis

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
	"time"
)

type IHisRepository interface {
	GetHisList() ([]object.IHisData, error)

	GetHisById(id string) (object.IHisData, error)
	GetHisByConfigId(id string) ([]object.IHisData, error)
	GetHisByTime(begTime, endTime time.Time) ([]object.IHisData, error)
	SetHis(data object.IHisData) error
	ClearHis(t time.Duration) error
}

func NewHisRepository(hisType global.HisType) (IHisRepository, error) {
	switch hisType {
	case global.HInt:
		return newHisRepository(&intHis{}), nil
	case global.HWebState:
		return newHisRepository(&webStateHis{}), nil
	case global.HCrmDzXfTest:
		return newHisRepository(&crmDzXfTestHis{}), nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的HisType：%d", hisType))
	}
}
