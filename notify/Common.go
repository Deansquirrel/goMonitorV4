package notify

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/global"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/notify"
	"reflect"
)

import log "github.com/Deansquirrel/goToolLog"

var comm common

func init() {
	comm = common{}
}

func SendMsg(id string, msg string) error {
	list, err := comm.GetNotifyList(id)
	if err != nil {
		errMsg := fmt.Sprintf("获取通知列表时发生错误:%s，消息未发送：%s", err.Error(), msg)
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	if len(list) < 1 {
		log.Error("无消息发送对象")
		return nil
	}
	sendFlag := false
	for _, n := range list {
		err = n.SendMsg(msg)
		if err != nil {
			log.Error(fmt.Sprintf("发送消息时遇到错误:%s，消息未发送：%s", err.Error(), msg))
		} else {
			sendFlag = true
		}
	}
	if !sendFlag {
		errMsg := fmt.Sprintf("消息未发送：%s", msg)
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

type common struct {
}

func (c *common) NewNotify(configData object.INotifyData) (iNotify, error) {
	switch reflect.TypeOf(configData).String() {
	case "*object.DingTalkRobotNotifyData":
		return &dingTalkRobot{
			configData: configData.(*object.DingTalkRobotNotifyData),
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("未预知的INotifyData类型，ID：%s", configData.GetNotifyId()))
	}
}

func (c *common) GetNotifyList(id string) ([]iNotify, error) {
	nl := notify.NewNotifyList()
	d, err := nl.GetNotifyList(id)
	if err != nil {
		errMsg := fmt.Sprintf("获取通知配置时发生错误：%s", err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	result := make([]iNotify, 0)

	errMsg := ""

	//获取DingTalkRobot类型配置数据
	dingTalkRep, err := notify.NewNotifyRepository(global.NDingTalkRobot)
	if err != nil {
		return nil, err
	}
	for _, id := range d.DingTalkRobot {
		n, err := dingTalkRep.GetNotify(id)
		if err != nil {
			errMsg = c.updateNotifyErr(errMsg, err)
		} else {
			dt, err := c.NewNotify(n.(*object.DingTalkRobotNotifyData))
			if err != nil {
				errMsg = c.updateNotifyErr(errMsg, err)
			}
			result = append(result, dt)
		}
	}

	if errMsg != "" {
		err = errors.New(errMsg)
	} else {
		err = nil
	}

	return result, err
}

func (c *common) updateNotifyErr(old string, err error) string {
	errMsgFormat := "获取[%s]NotifyData时发生错误：%s；"
	m := fmt.Sprintf(errMsgFormat, "DingTalkRobot", err.Error())
	log.Error(m)
	return old + m
}
