package notify

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goMonitorV4/repository/common"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"strings"
)

const SqlGetNotifyList = "" +
	"SELECT [DingTalkRobotId] " +
	"FROM NotifyList " +
	"WHERE [TaskId] = ? or TaskId = '-1'"

type notifyList struct {
}

func NewNotifyList() *notifyList {
	return &notifyList{}
}

func (nc *notifyList) GetNotifyList(id string) (*object.NotifyListData, error) {
	comm := common.Common{}
	rows, err := comm.GetRowsBySQL(SqlGetNotifyList, id)
	if err != nil {
		return nil, err
	}
	return nc.getNotifyListByRows(rows)
}

func (nc *notifyList) getNotifyListByRows(rows *sql.Rows) (*object.NotifyListData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var dingTalkRobot string
	dingTalkRobotList := make([]string, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&dingTalkRobot)
		if err != nil {
			break
		}
		list := strings.Split(dingTalkRobot, ",")
		list = goToolCommon.ClearBlock(list)
		for _, s := range list {
			dingTalkRobotList = append(dingTalkRobotList, s)
		}
	}
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if rows.Err() != nil {
		log.Error(rows.Err().Error())
		return nil, rows.Err()
	}
	dingTalkRobotList = goToolCommon.ClearRepeat(dingTalkRobotList)
	return &object.NotifyListData{
		DingTalkRobot: dingTalkRobotList,
	}, nil
}
