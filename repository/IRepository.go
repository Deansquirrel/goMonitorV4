package repository

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"time"
)

type IConfigRepository interface {
	GetConfigList() ([]object.IConfigData, error)
	GetConfig(id string) (object.IConfigData, error)
}

type IHisRepository interface {
	GetHisList() ([]object.IHisData, error)

	GetHisById(id string) (object.IHisData, error)
	GetHisByConfigId(id string) ([]object.IHisData, error)
	GetHisByTime(begTime, endTime time.Time) ([]object.IHisData, error)
	SetHis(data object.IHisData) error
	ClearHis(t time.Duration) error
}

type INotifyRepository interface {
	GetNotify(id string) (object.INotifyData, error)
}

type IActionRepository interface {
	GetAction(id string) (object.IActionData, error)
}

type IActionHisRepository interface {
	GetTopHis(configId string) ([]object.IHisData, error)
	SetHis(data object.IHisData) error
	ClearHis(t time.Duration) error
}
