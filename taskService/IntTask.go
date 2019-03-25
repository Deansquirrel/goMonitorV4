package taskService

import (
	"reflect"
)

type intTask struct {
}

func (it *intTask) getCacheIdList() []string {
	tc := TaskCache{}
	list := make([]string, 0)
	for _, c := range tc.GetTaskList() {
		switch reflect.TypeOf(c.Config).String() {
		case "*object.IntConfigData":
			list = append(list, c.Config.GetConfigId())
		}
	}
	return list
}
