package notify

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/object"
	"reflect"
)

func NewNotify(configData object.INotifyData) (INotify, error) {
	switch reflect.TypeOf(configData).String() {
	case "*object.DingTalkRobotNotifyData":
		return &dingTalkRobot{
			configData: configData.(*object.DingTalkRobotNotifyData),
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的INotifyData类型，ID：%s", configData.GetNotifyId()))
	}
}
