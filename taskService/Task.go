package taskService

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/config"
	"github.com/Deansquirrel/goMonitorV4/repository/configHis"
	"github.com/Deansquirrel/goMonitorV4/worker"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/robfig/cron"
	"time"
)

type task struct {
	iTask ITask
	cType global.ConfigType
	hType global.HisType
}

func NewTask(configType global.ConfigType) (*task, error) {
	switch configType {
	case global.CM, global.CIntD:
		return nil, errors.New(fmt.Sprintf("无效的ConfigType：%d", configType))
	case global.CInt:
		return &task{
			iTask: &intTask{},
			cType: global.CInt,
			hType: global.HInt,
		}, nil
	case global.CHealth:
		return &task{
			iTask: &healthTask{},
			cType: global.CHealth,
		}, nil
	case global.CCrmDzXfTest:
		return &task{
			iTask: &crmDzXfTestTask{},
			cType: global.CCrmDzXfTest,
			hType: global.HCrmDzXfTest,
		}, nil
	case global.CWebState:
		return &task{
			iTask: &webStateTask{},
			cType: global.CWebState,
			hType: global.HWebState,
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的ConfigType：%d", configType))
	}
}

func (t *task) StartTask() error {
	//==================================================
	t.clearCache()
	t.startRegularRefresh()
	//==================================================
	var rep config.IConfigRepository
	rep, err := config.NewConfigRepository(t.cType)
	if err != nil {
		return err
	}
	//获取配置列表
	list, err := rep.GetConfigList()
	if err != nil {
		return err
	}

	errMsg := ""
	errMsgFormat := "添加任务[%s]报错：%s；"
	for _, taskConfig := range list {
		err = t.addJob(taskConfig)
		if err != nil {
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, taskConfig.GetConfigId(), err.Error())
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func (t *task) addJob(iConfig object.IConfigData) error {
	tc := TaskCache{}
	w, err := worker.NewWorker(iConfig)
	if err != nil {
		tc.AddTask(iConfig.GetConfigId(), &cache{
			Config:    iConfig,
			Cron:      nil,
			IsRunning: false,
		})
		return err
	}
	c := cron.New()
	tc.AddTask(iConfig.GetConfigId(), &cache{
		Config:    iConfig,
		Cron:      c,
		IsRunning: false,
	})
	err = c.AddJob(iConfig.GetSpec(), w)
	if err != nil {
		return nil
	} else {
		c.Start()
		tc.UpdateTaskState(iConfig.GetConfigId(), true)
		return nil
	}
}

func (t *task) delJob(id string) {
	tc := TaskCache{}
	tc.DelTask(id)
}

func (t *task) startJob(id string) error {
	tc := TaskCache{}
	list := tc.GetTaskList()
	cache, ok := list[id]
	if ok {
		if cache.Cron == nil {
			return errors.New(fmt.Sprintf("配置无效，ID：%s", id))
		}
		if cache.IsRunning {
			return nil
		}
		cache.Cron.Start()
		tc.UpdateTaskState(id, true)
		return nil
	} else {
		return errors.New(fmt.Sprintf("无效的ID：%s", id))
	}
}

func (t *task) stopJob(id string) error {
	cache := TaskCache{}
	list := cache.GetTaskList()
	tc, ok := list[id]
	if ok {
		if tc.Cron == nil {
			return errors.New(fmt.Sprintf("配置无效，ID：%s", id))
		}
		if !tc.IsRunning {
			return nil
		}
		tc.Cron.Stop()
		cache.UpdateTaskState(id, false)
		return nil
	} else {
		return errors.New(fmt.Sprintf("无效的ID：%s", id))
	}
}

//清除缓存配置
func (t *task) clearCache() {
	for _, id := range t.iTask.getCacheIdList() {
		t.delJob(id)
	}
}

func (t *task) startRegularRefresh() {
	c := cron.New()
	var err error
	err = c.AddFunc("0 0/1 * * * ?", t.refreshConfig)
	if err != nil {
		log.Error("添加配置刷新任务时遇到错误：" + err.Error())
	} else {
		log.Info("添加配置刷新任务完成")
	}
	err = c.AddFunc("0 0 0 * * ?", t.delHisData)
	if err != nil {
		log.Error("添加删除历史数据任务时遇到错误：" + err.Error())
	} else {
		log.Info("添加删除历史数据任务完成")
	}
	c.Start()
}

//刷新任务配置
func (t *task) refreshConfig() {
	err := t.refreshConfigWorker()
	if err != nil {
		log.Error("刷新配置时遇到错误：" + err.Error())
	}
}

//删除历史数据
func (t *task) delHisData() {
	d := time.Duration(1000 * 1000 * 1000 * 60 * 60 * 24 * global.SysConfig.TaskConfig.KeepDays)
	rep, err := configHis.NewHisRepository(t.hType)
	if err != nil {
		log.Error(err.Error())
		return
	}
	err = rep.ClearHis(d)
	if err != nil {
		log.Error("删除历史数据时遇到错误：" + err.Error())
	}
}

func (t *task) refreshConfigWorker() error {
	var rep config.IConfigRepository
	rep, err := config.NewConfigRepository(t.cType)
	if err != nil {
		return err
	}
	//获取配置列表
	list, err := rep.GetConfigList()
	if err != nil {
		return err
	}
	idList := make([]string, 0)
	idMap := make(map[string]object.IConfigData, 0)
	for _, iConfig := range list {
		taskConfig, ok := iConfig.(object.IConfigData)
		if !ok {
			return errors.New("不是有效的IConfigData")
		}
		idList = append(idList, taskConfig.GetConfigId())
		idMap[taskConfig.GetConfigId()] = taskConfig
	}

	addList, delList, checkList := goToolCommon.CheckDiff(idList, t.iTask.getCacheIdList())

	errMsg := ""
	errMsgFormat := "刷新任务[%s]报错：%s；"

	for _, id := range addList {
		err = t.addJob(idMap[id])
		if err != nil {
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
		}
	}
	for _, id := range delList {
		t.delJob(id)
	}
	tc := TaskCache{}
	for _, id := range checkList {
		tcList := tc.GetTaskList()
		tc := tcList[id]
		if !tc.Config.IsEqual(idMap[id]) {
			t.delJob(id)
			err = t.addJob(idMap[id])
			if err != nil {
				errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
			}
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}
