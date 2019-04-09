package notify

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
)

type INotifyRepository interface {
	GetNotify(id string) (object.INotifyData, error)
}

func NewNotifyRepository(notifyType global.NotifyType) (INotifyRepository, error) {
	switch notifyType {
	case global.NDingTalkRobot:
		return newNotifyRepository(&dingTalkRobotNotify{}), nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的NotifyType：%d", notifyType))
	}
}
