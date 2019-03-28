package repository

import (
	"database/sql"
	"errors"
	"github.com/Deansquirrel/goMonitorV4/object"
	"reflect"
	"time"
)
import log "github.com/Deansquirrel/goToolLog"

const SqlGetWebStateTaskHis = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis"
const SqlGetWebStateTaskHisById = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis" +
	" WHERE FId = ?"
const SqlGetWebStateTaskHisByConfigId = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis" +
	" WHERE FConfigId = ?"
const SqlGetWebStateTaskHisByTime = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis" +
	" WHERE [FOprTime] >= ? AND [FOprTime] <= ?"

const SqlSetWebStateTaskHis = "" +
	"INSERT INTO WebStateTaskHis (FId, FConfigId, FUseTime, FHttpCode, FContent)" +
	" VALUES (?,?,?,?,?)"
const SqlDelWebStateTaskHis = "" +
	"DELETE FROM WebStateTaskHis" +
	" WHERE FOprTime < ?"

type webStateHis struct {
}

func (wsh *webStateHis) GetSqlHisList() string {
	return SqlGetWebStateTaskHis
}

func (wsh *webStateHis) GetSqlHisById() string {
	return SqlGetWebStateTaskHisById
}

func (wsh *webStateHis) GetSqlHisByConfigId() string {
	return SqlGetWebStateTaskHisByConfigId
}

func (wsh *webStateHis) GetSqlHisByTime() string {
	return SqlGetWebStateTaskHisByTime
}

func (wsh *webStateHis) GetSqlSetHis() string {
	return SqlSetWebStateTaskHis
}

func (wsh *webStateHis) GetSqlClearHis() string {
	return SqlDelWebStateTaskHis
}

func (wsh *webStateHis) getHisListByRows(rows *sql.Rows) ([]object.IHisData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fConfigId, fContent string
	var fUseTime, fHttpCode int
	var fOprTime time.Time
	resultList := make([]object.IHisData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fConfigId, &fUseTime, &fHttpCode, &fContent, &fOprTime)
		if err != nil {
			break
		}
		config := object.WebStateHisData{
			FId:       fId,
			FConfigId: fConfigId,
			FUseTime:  fUseTime,
			FHttpCode: fHttpCode,
			FContent:  fContent,
			FOprTime:  fOprTime,
		}
		resultList = append(resultList, &config)
	}
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if rows.Err() != nil {
		log.Error(rows.Err().Error())
		return nil, rows.Err()
	}
	return resultList, nil
}

func (wsh *webStateHis) getHisSetArgs(data object.IHisData) ([]interface{}, error) {
	switch reflect.TypeOf(data).String() {
	case "*object.WebStateHisData":
		iHisData, ok := data.(*object.WebStateHisData)
		if ok {
			result := make([]interface{}, 0)
			result = append(result, iHisData.FId)
			result = append(result, iHisData.FConfigId)
			result = append(result, iHisData.FUseTime)
			result = append(result, iHisData.FHttpCode)
			result = append(result, iHisData.FContent)
			return result, nil
		} else {
			return nil, errors.New("强制类型转换失败[WebStateHisData]")
		}
	default:
		return nil, errors.New("webStateHis getHisSetArgs 参数类型错误")
	}
}
