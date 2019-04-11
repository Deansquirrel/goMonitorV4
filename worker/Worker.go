package worker

import (
	"errors"
	"github.com/Deansquirrel/goMonitorV4/action"
	"github.com/Deansquirrel/goMonitorV4/notify"
	"github.com/Deansquirrel/goMonitorV4/object"
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

	err := w.saveHis(hisData)
	if err != nil {
		log.Error(err.Error())
		w.sendMsg(w.iConfig.GetConfigId(), err.Error())
	}

	if msg == "" {
		return
	}
	w.sendMsg(w.iConfig.GetConfigId(), msg)

	err = action.CheckAction(w.iConfig.GetConfigId())
	if err != nil {
		log.Error(err.Error())
		w.sendMsg(w.iConfig.GetConfigId(), "Check Action Error: "+err.Error())
	}
}

//保存查询数据
func (w *worker) saveHis(hisData object.IHisData) error {
	if hisData != nil {
		rep, err := w.iWorker.getHisRepository()
		if err != nil {
			return err
		}
		err = rep.SetHis(hisData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *worker) sendMsg(id string, msg string) {
	err := notify.SendMsg(id, msg)
	if err != nil {
		log.Error(err.Error())
	}
}
