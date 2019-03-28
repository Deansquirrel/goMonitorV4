package service

import (
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/taskService"
	"github.com/Deansquirrel/goMonitorV4/webService"
	"os"
	"os/signal"
	"syscall"
)
import log "github.com/Deansquirrel/goToolLog"

type cronService struct {
}

func NewCronService() *cronService {
	return &cronService{}
}

func (cs *cronService) Start() {
	log.Debug("CronService starting")
	defer log.Debug("CronService start complete")
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			os.Interrupt,
			syscall.SIGINT,
			os.Kill,
			syscall.SIGKILL,
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			defer global.Cancel()
		case <-global.Ctx.Done():
		}
	}()

	intTask, err := taskService.NewTask(global.CInt)
	if err != nil {
		log.Error(err.Error())
	} else {
		err = intTask.StartTask()
		if err != nil {
			log.Error(err.Error())
		}
	}

	crmDzXfTestTask, err := taskService.NewTask(global.CCrmDzXfTest)
	if err != nil {
		log.Error(err.Error())
	} else {
		err = crmDzXfTestTask.StartTask()
		if err != nil {
			log.Error(err.Error())
		}
	}

	healthTask, err := taskService.NewTask(global.CHealth)
	if err != nil {
		log.Error(err.Error())
	} else {
		err = healthTask.StartTask()
		if err != nil {
			log.Error(err.Error())
		}
	}

	webState, err := taskService.NewTask(global.CWebState)
	if err != nil {
		log.Error(err.Error())
	} else {
		err = webState.StartTask()
		if err != nil {
			log.Error(err.Error())
		}
	}

	ws := webService.NewWebServer(global.SysConfig.IrisConfig.Port)
	ws.StartWebService()
}
