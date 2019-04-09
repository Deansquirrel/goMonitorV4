package configHis

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/common"
	"github.com/Deansquirrel/goToolCommon"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

type hisRepository struct {
	iHis IHis
}

func newHisRepository(iHis IHis) *hisRepository {
	return &hisRepository{
		iHis: iHis,
	}
}

func (hr *hisRepository) GetHisList() ([]object.IHisData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(hr.iHis.GetSqlHisList())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.iHis.GetHisListByRows(rows)
}

func (hr *hisRepository) GetHisById(id string) (object.IHisData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(hr.iHis.GetSqlHisById(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.iHis.GetHisListByRows(rows)
}

func (hr *hisRepository) GetHisByConfigId(id string) ([]object.IHisData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(hr.iHis.GetSqlHisByConfigId(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.iHis.GetHisListByRows(rows)
}

func (hr *hisRepository) GetHisByTime(begTime, endTime time.Time) ([]object.IHisData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(hr.iHis.GetSqlHisByTime(), begTime, endTime)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.iHis.GetHisListByRows(rows)
}

func (hr *hisRepository) SetHis(data object.IHisData) error {
	args, err := hr.iHis.GetHisSetArgs(data)
	if err != nil {
		return err
	}
	comm := common.Common{}
	err = comm.SetRowsBySQL(hr.iHis.GetSqlSetHis(), args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (hr *hisRepository) ClearHis(t time.Duration) error {
	dateP := goToolCommon.GetDateTimeStr(time.Now().Add(-t))
	comm := common.Common{}
	err := comm.SetRowsBySQL(hr.iHis.GetSqlClearHis(), dateP)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
