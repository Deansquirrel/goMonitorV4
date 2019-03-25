package global

import (
	"context"
	"github.com/Deansquirrel/goMonitorV4/config"
	"github.com/Deansquirrel/goToolMSSql"
	"time"
)

const (
	//PreVersion = "0.0.3 Build20190315"
	//TestVersion = "0.0.0 Build20190101"
	Version = "0.0.0 Build20190101"
)

const (
	HttpConnectTimeout = 30
)

type ConfigType int

const (
	CMConfig ConfigType = iota
	CInt
)

type NotifyType int

const (
	NDingTalkRobot NotifyType = iota
)

type WorkerType int

const (
	WInt WorkerType = iota
)

type HisType int

const (
	HInt HisType = iota
)

var SysConfig *config.SysConfig
var Ctx context.Context
var Cancel func()

func init() {
	goToolMSSql.SetMaxIdleConn(15)
	goToolMSSql.SetMaxOpenConn(15)
	goToolMSSql.SetMaxLifetime(time.Second * 60)
}
