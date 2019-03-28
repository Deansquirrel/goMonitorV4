package worker

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/notify"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository"
	log "github.com/Deansquirrel/goToolLog"
	"reflect"
)

type worker struct {
	iConfig object.IConfigData
	iWorker IWorker
}

func NewWorker(iConfig object.IConfigData) (*worker, error) {
	workerRunner, err := getWorker(iConfig)
	if err != nil {
		return nil, err
	}
	return &worker{
		iConfig: iConfig,
		iWorker: workerRunner,
	}, nil
}

func getWorker(iConfig object.IConfigData) (IWorker, error) {
	switch reflect.TypeOf(iConfig).String() {
	case "*object.IntConfigData":
		return &intWorker{iConfig.(*object.IntConfigData)}, nil
	case "*object.CrmDzXfTestConfigData":
		return &crmDzXfTestWorker{iConfig.(*object.CrmDzXfTestConfigData)}, nil
	case "*object.HealthConfigData":
		return &healthWorker{iConfig.(*object.HealthConfigData)}, nil
	case "*object.WebStateConfigData":
		return &webStateWorker{iConfig.(*object.WebStateConfigData)}, nil
	default:
		return nil, errors.New("未预知的配置类型:" + reflect.TypeOf(iConfig).String())
	}
}

func (w *worker) Run() {
	msg, hisData := w.iWorker.GetMsg()
	defer func() {
		if hisData != nil {
			rep, err := w.iWorker.getHisRepository()
			if err != nil {
				log.Error(err.Error())
				w.sendMsg(w.iConfig.GetConfigId(), err.Error())
			}
			err = rep.SetHis(hisData)
			if err != nil {
				log.Error(err.Error())
				w.sendMsg(w.iConfig.GetConfigId(), err.Error())
			}
		}
	}()
	if msg == "" {
		return
	}
	w.sendMsg(w.iConfig.GetConfigId(), msg)
}

func (w *worker) sendMsg(configId, msg string) {
	list, err := w.getNotifyList(w.iConfig.GetConfigId())
	if err != nil {
		log.Error(fmt.Sprintf("获取通知列表时发生错误:%s，消息未发送：%s", err.Error(), msg))
		return
	}
	sendFlag := false
	for _, n := range list {
		err = n.SendMsg(msg)
		if err != nil {
			log.Error(fmt.Sprintf("发送消息时遇到错误:%s，消息未发送：%s", err.Error(), msg))
		} else {
			sendFlag = true
		}
	}
	if !sendFlag {
		log.Warn(fmt.Sprintf("消息未发送：%s", msg))
	}
}

//func (w *worker) saveSearchResult(data object.IHisData) error {
//	rep,err := repository.NewHisRepository(global.HCrmDzXfTest)
//	if err != nil {
//		return err
//	}
//	return rep.SetHis(data)
//}

func (w *worker) getNotifyList(id string) ([]notify.INotify, error) {
	nl := repository.NotifyList{}
	d, err := nl.GetNotifyList(id)
	if err != nil {
		errMsg := fmt.Sprintf("获取通知配置时发生错误：%s", err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	result := make([]notify.INotify, 0)

	errMsg := ""

	//获取DingTalkRobot类型配置数据
	dingTalkRep, err := repository.NewNotifyRepository(global.NDingTalkRobot)
	if err != nil {
		return nil, err
	}
	for _, id := range d.DingTalkRobot {
		n, err := dingTalkRep.GetNotify(id)
		if err != nil {
			errMsg = w.updateNotifyErr(errMsg, err)
		} else {
			dt, err := notify.NewNotify(n.(*object.DingTalkRobotNotifyData))
			if err != nil {
				errMsg = w.updateNotifyErr(errMsg, err)
			}
			result = append(result, dt)
		}
	}

	if errMsg != "" {
		err = errors.New(errMsg)
	} else {
		err = nil
	}

	return result, err
}

func (w *worker) updateNotifyErr(old string, err error) string {
	errMsgFormat := "获取[%s]NotifyData时发生错误：%s；"
	m := fmt.Sprintf(errMsgFormat, "DingTalkRobot", err.Error())
	log.Error(m)
	return old + m
}
