package worker

import (
	"errors"
	"github.com/Deansquirrel/goMonitorV4/action"
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

	err := w.saveHis(hisData)
	if err != nil {
		log.Error(err.Error())
		w.sendMsg(w.iConfig.GetConfigId(), err.Error())
	}

	if msg == "" {
		return
	}
	w.sendMsg(w.iConfig.GetConfigId(), msg)

	err = w.checkAction(w.iConfig.GetConfigId())
	if err != nil {
		log.Error(err.Error())
		w.sendMsg(w.iConfig.GetConfigId(), err.Error())
	}
}

//检查并执行相关操作
func (w *worker) checkAction(id string) error {
	actionListRepository := repository.ActionList{}
	actionList, err := actionListRepository.GetActionList(id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var errMsg string
	for _, s := range actionList.WindowsService {
		err = w.checkActionWorker(s, global.AWindowsService)
		if err != nil {
			errMsg = errMsg + err.Error() + ";"
		}
	}

	for _, s := range actionList.IISAppPool {
		err = w.checkActionWorker(s, global.AIISAppPool)
		if err != nil {
			errMsg = errMsg + err.Error() + ";"
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func (w *worker) checkActionWorker(id string, t global.ActionType) error {
	rep, err := repository.NewActionRepository(t)
	if err != nil {
		return err
	}
	actionData, err := rep.GetAction(id)
	if err != nil {
		return err
	}
	ac, err := action.NewAction(actionData)
	if err != nil {
		return err
	}
	return ac.Do()
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
