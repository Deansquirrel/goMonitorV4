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
		msg := comm.getMsg(wsw.configData.FMsgTitle, err.Error())
		msg = wsw.formatMsg(msg)
		return msg, wsw.getHisData(useTime, httpCode, msg)
	}
	var msg string
	if httpCode != 200 {
		msg = comm.getMsg(wsw.configData.FMsgTitle, wsw.configData.FMsgContent)
		if msg != "" {
			msg = msg + "\n"
		}
		msg = msg + fmt.Sprintf("返回码：%d", httpCode)
		msg = msg + "\n"
		msg = msg + fmt.Sprintf("用时：%d", useTime)
		msg = wsw.formatMsg(msg)
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
	}
	return int(ms), code, err
}

func (wsw *webStateWorker) getHttpCode() (int, error) {
	//req, err := http.NewRequest("GET", wsw.configData.FUrl, nil)
	//if err != nil {
	//	return -1, err
	//}
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Timeout:   time.Second * global.HttpConnectTimeout,
		Transport: &transport,
	}
	resp, err := client.Get(wsw.configData.FUrl)

	//resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	httpCode := resp.StatusCode
	//_, err = ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Error(err.Error())
	//}
	return httpCode, nil
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
