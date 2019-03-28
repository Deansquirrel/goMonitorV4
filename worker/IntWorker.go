package worker

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type intWorker struct {
	configData *object.IntConfigData
}

func (iw *intWorker) GetMsg() (string, object.IHisData) {
	comm := common{}
	if iw.configData == nil {
		msg := comm.getMsg("", "配置内容为空")
		msg = iw.formatMsg(msg)
		return msg, iw.getHisData(0, msg)
	}
	num, err := iw.getCheckNum()
	if err != nil {
		msg := comm.getMsg(iw.configData.FMsgTitle, "获取数量时遇到错误")
		msg = iw.formatMsg(msg)
		return msg, iw.getHisData(0, msg)
	}
	var msg string
	if num >= iw.configData.FCheckMax || num <= iw.configData.FCheckMin {
		msg = comm.getMsg(iw.configData.FMsgTitle,
			strings.Replace(iw.configData.FMsgContent, "title", strconv.Itoa(num), -1))
		dMsg := iw.getDMsg()
		if msg != "" {
			if dMsg != "" {
				msg = msg + "\n" + "--------------------" + "\n" + dMsg
			}
		} else {
			msg = dMsg
		}
		msg = iw.formatMsg(msg)
	}
	return msg, iw.getHisData(num, msg)
}

func (iw *intWorker) getHisData(num int, msg string) object.IHisData {
	return &object.IntHisData{
		FId:       strings.ToUpper(goToolCommon.Guid()),
		FConfigId: iw.configData.FId,
		FNum:      num,
		FContent:  msg,
		FOprTime:  time.Now(),
	}
}

func (iw *intWorker) formatMsg(msg string) string {
	if msg != "" {
		msg = goToolCommon.GetDateTimeStr(time.Now()) + "\n" + msg
	}
	return msg
}

func (iw *intWorker) getDMsg() string {
	rep, err := repository.NewConfigRepository(global.CIntD)
	if err != nil {
		log.Error(fmt.Sprintf(err.Error()))
		return ""
	}
	dConfig, err := rep.GetConfig(iw.configData.FId)
	if err != nil {
		log.Error(fmt.Sprintf("获取明细配置时遇到错误：%s，查询ID为：%s", err.Error(), iw.configData.FId))
		return ""
	}
	//无明细配置
	if dConfig == nil {
		return ""
	}
	switch reflect.TypeOf(dConfig).String() {
	case "*object.IntDConfigData":
		c, ok := dConfig.(*object.IntDConfigData)
		if ok {
			return iw.getSingleDMsg(c.FMsgSearch)
		} else {
			return "强制类型转换失败[IntConfigData]"
		}

	default:
		log.Error(fmt.Sprintf("获取的明细配置类型异常，expr：IntDConfigData"))
		return ""
	}
}

func (iw *intWorker) getSingleDMsg(search string) string {
	if search == "" {
		return ""
	}
	rows, err := iw.getRowsBySQL(search)
	if err != nil {
		return fmt.Sprintf("查询明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	defer func() {
		_ = rows.Close()
	}()
	titleList, err := rows.Columns()
	if err != nil {
		return fmt.Sprintf("获取明细内容表头时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	counter := len(titleList)
	values := make([]interface{}, counter)
	valuePointers := make([]interface{}, counter)
	for i := 0; i < counter; i++ {
		valuePointers[i] = &values[i]
	}

	var result string
	for rows.Next() {
		err = rows.Scan(valuePointers...)
		if err != nil {
			return fmt.Sprintf("读取明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
		}
		if result != "" {
			result = result + "\n" + "--------------------"
		}
		for i := 0; i < counter; i++ {
			if result != "" {
				result = result + "\n"
			}
			var v string
			if values[i] == nil {
				v = "Null"
			} else {
				v = goToolCommon.ConvertToString(values[i])
			}
			result = result + titleList[i] + " - " + v
		}
	}
	if rows.Err() != nil {
		return fmt.Sprintf("读取明细内容时遇到错误：%s，查询语句为：%s", err.Error(), search)
	}
	return result
}

func (iw *intWorker) getHisRepository() (repository.IHisRepository, error) {
	return repository.NewHisRepository(global.HInt)
}

//
//func (iw *intWorker) SaveSearchResult(data object.IHisData) error {
//	switch reflect.TypeOf(data).String() {
//	case "*object.IntHisData":
//		rep, err := repository.NewHisRepository(global.HInt)
//		if err != nil {
//			return err
//		}
//		iHisData, ok := data.(*object.IntHisData)
//		if ok {
//			err := rep.SetHis(iHisData)
//			if err != nil {
//				s, _ := goToolCommon.GetJsonStr(data)
//				errMsg := fmt.Sprintf("保存查询结果时遇到错误：%s，待保存内容：%s", err.Error(), s)
//				log.Error(errMsg)
//				return errors.New(errMsg)
//			}
//			return nil
//		} else {
//			s, _ := goToolCommon.GetJsonStr(data)
//			errMsg := fmt.Sprintf("强制类型转换失败[IntHisData]，待保存内容：%s", s)
//			log.Error(errMsg)
//			return errors.New(errMsg)
//		}
//	default:
//		s, _ := goToolCommon.GetJsonStr(data)
//		errMsg := fmt.Sprintf("历史数据类型异常，exp：IntHisData，待保存内容：%s", s)
//		log.Error(errMsg)
//		return errors.New(errMsg)
//	}
//}

//获取待检测值
func (iw *intWorker) getCheckNum() (int, error) {
	rows, err := iw.getRowsBySQL(iw.configData.FSearch)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rows.Close()
	}()
	list := make([]int, 0)
	var num int
	for rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			log.Error(err.Error())
			break
		} else {
			list = append(list, num)
		}
	}
	if err != nil {
		return 0, err
	}
	if len(list) != 1 {
		errMsg := fmt.Sprintf("SQL返回数量异常，exp:1,act:%d", len(list))
		log.Error(errMsg)
		return 0, errors.New(errMsg)
	}
	return list[0], nil
}

//查询数据
func (iw *intWorker) getRowsBySQL(sql string) (*sql.Rows, error) {
	conn, err := goToolMSSql.GetConn(iw.getDBConfig())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	rows, err := conn.Query(sql)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return rows, nil
}

//获取DB配置
func (iw *intWorker) getDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: iw.configData.FServer,
		Port:   iw.configData.FPort,
		DbName: iw.configData.FDbName,
		User:   iw.configData.FDbUser,
		Pwd:    iw.configData.FDbPwd,
	}
}
