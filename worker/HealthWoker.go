package worker

import (
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolEnvironment"
	log "github.com/Deansquirrel/goToolLog"
	"time"
)

type healthWorker struct {
	configData *object.HealthConfigData
}

func (hw *healthWorker) GetMsg() (string, object.IHisData) {
	comm := common{}
	if hw.configData == nil {
		msg := comm.getMsg("", "配置内容为空")
		msg = hw.formatMsg(msg)
		return msg, nil
	}
	msg := comm.getMsg(hw.configData.FMsgTitle, hw.configData.FMsgContent)
	return hw.formatMsg(msg), nil
}

func (hw *healthWorker) getHisRepository() (repository.IHisRepository, error) {
	return nil, nil
}

func (hw *healthWorker) formatMsg(msg string) string {
	var iAddr string
	iAddr, err := goToolEnvironment.GetInternetAddr()
	if err != nil {
		log.Error("获取外网地址时遇到错误：" + err.Error())
	}
	if msg != "" {
		if iAddr != "" {
			msg = iAddr + "\n" + msg
		}
		msg = goToolCommon.GetDateTimeStr(time.Now()) + "\n" + msg
	}
	return msg
}
