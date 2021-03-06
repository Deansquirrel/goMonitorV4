package main

import (
	"context"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/common"
	"github.com/Deansquirrel/goMonitorV4/global"
	myService "github.com/Deansquirrel/goMonitorV4/service"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kardianos/service"
	"os"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	//服务所执行的代码
	cs := myService.NewCronService()
	cs.Start()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	log.StdOut = true
	//==================================================================================================================
	log.Warn("程序启动")
	log.Warn(fmt.Sprintf("Version %s", global.Version))
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
	//==================================================================================================================
	svcConfig := &service.Config{
		Name:        global.SysConfig.ServiceConfig.Name,
		DisplayName: global.SysConfig.ServiceConfig.DisplayName,
		Description: global.SysConfig.ServiceConfig.Description,
	}
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Error("定义服务配置时遇到错误：" + err.Error())
		return
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err = s.Install()
			msg := ""
			if err != nil {
				msg = err.Error()
			} else {
				msg = "服务安装成功"
			}
			log.Warn(msg)
			return
		case "uninstall":
			err = s.Uninstall()
			msg := ""
			if err != nil {
				msg = err.Error()
			} else {
				msg = "服务卸载成功"
			}
			log.Warn(msg)
			return
		default:
			log.Error("未识别的参数名称\n安装服务:install\n卸载服务:uninstall")
			return
		}
	} else {
		err = s.Run()
		if err != nil {
			log.Error(err.Error())
		}
	}
	//==================================================================================================================
}
