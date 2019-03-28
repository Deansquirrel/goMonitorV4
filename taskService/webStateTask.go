package taskService

import "reflect"

type webStateTask struct {
}

func (wst *webStateTask) getCacheIdList() []string {
	tc := TaskCache{}
	list := make([]string, 0)
	for _, c := range tc.GetTaskList() {
		switch reflect.TypeOf(c.Config).String() {
		case "*object.WebStateConfigData":
			list = append(list, c.Config.GetConfigId())
		}
	}
	return list
}
