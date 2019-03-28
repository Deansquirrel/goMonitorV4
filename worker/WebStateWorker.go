package worker

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"net/http"
	"strings"
	"time"
)

type webStateWorker struct {
	configData *object.WebStateConfigData
}

func (wsw *webStateWorker) GetMsg() (string, object.IHisData) {
	comm := common{}
	if wsw.configData == nil {
		msg := comm.getMsg("", "配置内容为空")
		msg = wsw.formatMsg(msg)
		return msg, wsw.getHisData(-1, 0, msg)
	}
	useTime, httpCode, err := wsw.getWebStateTestData()
	if err != nil {
		return wsw.formatMsg(comm.getMsg(wsw.configData.FMsgTitle, err.Error())), wsw.getHisData(useTime, httpCode, err.Error())
	}
	var msg string
	if httpCode != 200 || useTime > 5*1000 {
		msg = comm.getMsg(wsw.configData.FMsgTitle, wsw.configData.FMsgContent)
		if msg != "" {
			msg = msg + "\n"
		}
		msg = msg + fmt.Sprintf("返回码：%d", httpCode)
		msg = msg + "\n"
		msg = msg + fmt.Sprintf("用时：%d", useTime)
	}
	return msg, wsw.getHisData(useTime, httpCode, msg)
}

//SaveSearchResult(data object.IHisData) error

func (wsw *webStateWorker) getHisRepository() (repository.IHisRepository, error) {
	return repository.NewHisRepository(global.HWebState)
}

func (wsw *webStateWorker) getWebStateTestData() (useTime, httpCode int, err error) {
	useTime = -1
	httpCode = 0
	begTime := time.Now()
	code, err := wsw.getHttpCode()
	endTime := time.Now()
	ns := endTime.Sub(begTime).Nanoseconds()
	ms := ns / 1000 / 1000
	if err != nil {
		log.Error(err.Error())
		return -1, 0, err
	}
	return int(ms), code, nil
}

func (wsw *webStateWorker) getHttpCode() (int, error) {
	req, err := http.NewRequest("GET", wsw.configData.FUrl, nil)
	if err != nil {
		return -1, err
	}
	client := &http.Client{
		Timeout: time.Second * global.HttpConnectTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return resp.StatusCode, nil
}

func (wsw *webStateWorker) getHisData(useTime, httpCode int, msg string) object.IHisData {
	return &object.WebStateHisData{
		FId:       strings.ToUpper(goToolCommon.Guid()),
		FConfigId: wsw.configData.FId,
		FUseTime:  useTime,
		FHttpCode: httpCode,
		FContent:  msg,
		FOprTime:  time.Now(),
	}
}

func (wsw *webStateWorker) formatMsg(msg string) string {
	if msg != "" {
		msg = goToolCommon.GetDateTimeStr(time.Now()) + "\n" + msg
		if wsw.configData.FUrl != "" {
			msg = msg + "\n" + wsw.configData.FUrl
		}
	}
	return msg
}
