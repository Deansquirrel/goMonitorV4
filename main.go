package main

import (
	"context"
	"github.com/Deansquirrel/goMonitorV4/common"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/taskService"
	log "github.com/Deansquirrel/goToolLog"
)

func main() {
	//==================================================================================================================
	log.Warn("程序启动")
	defer log.Warn("程序退出")
	//==================================================================================================================
	config, err := common.GetSysConfig("config.toml")
	if err != nil {
		log.Error("加载配置文件时遇到错误：" + err.Error())
		return
	}
	config.FormatConfig()
	global.SysConfig = config
	err = common.RefreshSysConfig(*global.SysConfig)
	if err != nil {
		log.Error("刷新配置时遇到错误：" + err.Error())
		return
	}
	global.Ctx, global.Cancel = context.WithCancel(context.Background())

	//intTask, err := taskService.NewTask(global.CInt)
	//if err != nil {
	//	log.Debug(err.Error())
	//} else {
	//	err = intTask.StartTask()
	//	if err != nil {
	//		log.Debug(err.Error())
	//	}
	//}

	//crmDzXfTestTask, err := taskService.NewTask(global.CCrmDzXfTest)
	//if err != nil {
	//	log.Debug(err.Error())
	//} else {
	//	err = crmDzXfTestTask.StartTask()
	//	if err != nil {
	//		log.Debug(err.Error())
	//	}
	//}

	//healthTask, err := taskService.NewTask(global.CHealth)
	//if err != nil {
	//	log.Debug(err.Error())
	//} else {
	//	err = healthTask.StartTask()
	//	if err != nil {
	//		log.Debug(err.Error())
	//	}
	//}

	webState, err := taskService.NewTask(global.CWebState)
	if err != nil {
		log.Debug(err.Error())
	} else {
		err = webState.StartTask()
		if err != nil {
			log.Debug(err.Error())
		}
	}

	select {
	case <-global.Ctx.Done():
	}
	//==================================================================================================================
}
