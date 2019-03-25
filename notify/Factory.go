package notify

import "github.com/Deansquirrel/goMonitorV4/object"

func NewDingTalkRobot(configData *object.DingTalkRobotNotifyData) *dingTalkRobot {
	return &dingTalkRobot{
		configData: configData,
	}
}
