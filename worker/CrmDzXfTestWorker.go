package worker

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type crmDzXfTestWorker struct {
	configData *object.CrmDzXfTestConfigData
}

func (cw *crmDzXfTestWorker) GetMsg() (string, object.IHisData) {
	comm := common{}
	if cw.configData == nil {
		msg := comm.getMsg(cw.configData.FMsgTitle, "配置内容为空")
		msg = cw.formatMsg(msg)
		return msg, cw.getHisData(-1, 0, msg)
	}
	useTime, httpCode, err := cw.getCrmDzXfTestTestData()
	if err != nil {
		return comm.getMsg(cw.configData.FMsgTitle, err.Error()), cw.getHisData(useTime, httpCode, err.Error())
	}

	var msg string
	if httpCode != 200 || useTime > 5*1000 {
		msg = comm.getMsg(cw.configData.FMsgTitle, cw.configData.FMsgContent)
		if msg != "" {
			msg = msg + "\n"
		}
		msg = msg + fmt.Sprintf("返回码：%d", httpCode)
		msg = msg + "\n"
		msg = msg + fmt.Sprintf("用时：%d", useTime)
	}
	return msg, cw.getHisData(useTime, httpCode, msg)
}

func (cw *crmDzXfTestWorker) getCrmDzXfTestTestData() (useTime, httpCode int, err error) {
	useTime = -1
	httpCode = 0
	begTime := time.Now()
	code, err := cw.getHttpCode()
	endTime := time.Now()
	ns := endTime.Sub(begTime).Nanoseconds()
	ms := ns / 1000 / 1000
	if err != nil {
		log.Error(err.Error())
		return -1, 0, err
	}
	return int(ms), code, nil
}

func (cw *crmDzXfTestWorker) getHttpCode() (int, error) {
	testData, err := goToolCommon.GetJsonStr(cw.getTestData())
	if err != nil {
		return 0, errors.New("构造测试数据时发生错误：" + err.Error())
	}
	req, err := http.NewRequest("POST", cw.configData.FAddress, bytes.NewBuffer([]byte(testData)))
	if err != nil {
		return 0, errors.New("构造http请求数据时发生错误：" + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("passporttype", strconv.Itoa(cw.configData.FPassportType))
	req.Header.Set("passport", cw.configData.FPassport)

	client := &http.Client{
		Timeout: time.Second * global.HttpConnectTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.New("发送http请求时错误：" + err.Error())
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return resp.StatusCode, nil
}

func (cw *crmDzXfTestWorker) getHisData(useTime, httpCode int, msg string) object.IHisData {
	return &object.CrmDzXfTestHisData{
		FId:       strings.ToUpper(goToolCommon.Guid()),
		FConfigId: cw.configData.FId,
		FUseTime:  useTime,
		FHttpCode: httpCode,
		FContent:  msg,
		FOprTime:  time.Now(),
	}
}

func (cw *crmDzXfTestWorker) getHisRepository() (repository.IHisRepository, error) {
	return repository.NewHisRepository(global.HCrmDzXfTest)
}

func (cw *crmDzXfTestWorker) formatMsg(msg string) string {
	if msg != "" {
		msg = goToolCommon.GetDateTimeStr(time.Now()) + "\n" + msg
	}
	return msg
}

func (cw *crmDzXfTestWorker) getTestData() *crmDzXfRequestData {
	ywCore := crmDzXfRequestDataYwCore{
		Oprtime:   goToolCommon.GetDateTimeStr(time.Now()),
		Oprbrid:   10001,
		Oprbrname: "测试请求",
		Oprxfje:   1000000,
	}
	ywInfo := crmDzXfRequestDataYwInfo{
		Oprywsno:    "YW" + goToolCommon.GetDateStr(time.Now()) + "01",
		Oprppid:     10001,
		Oprppname:   "",
		Oprid:       182,
		Oprname:     "管理员",
		Oprywdate:   goToolCommon.GetDateStr(time.Now()) + " 00:00:00",
		Oprskbrid:   0,
		Oprskbrname: "",
		Oprskppid:   0,
		Oprskppname: "",
		Oprywwindow: "测试请求",
		Oprywbno:    "",
		Oprsummary:  "",
	}
	return &crmDzXfRequestData{
		YwCore: ywCore,
		YwInfo: ywInfo,
	}
}

type crmDzXfRequestData struct {
	YwCore crmDzXfRequestDataYwCore
	YwInfo crmDzXfRequestDataYwInfo
}

type crmDzXfRequestDataYwCore struct {
	Oprtime   string
	Oprbrid   int
	Oprbrname string
	Oprxfje   int
}

type crmDzXfRequestDataYwInfo struct {
	Oprywsno    string
	Oprppid     int
	Oprppname   string
	Oprid       int
	Oprname     string
	Oprywdate   string
	Oprskbrid   int
	Oprskbrname string
	Oprskppid   int
	Oprskppname string
	Oprywwindow string
	Oprywbno    string
	Oprsummary  string
}
