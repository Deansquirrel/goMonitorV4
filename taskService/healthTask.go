package taskService

import "reflect"

type healthTask struct {
}

func (ht *healthTask) getCacheIdList() []string {
	tc := TaskCache{}
	list := make([]string, 0)
	for _, c := range tc.GetTaskList() {
		switch reflect.TypeOf(c.Config).String() {
		case "*object.HealthConfigData":
			list = append(list, c.Config.GetConfigId())
		}
	}
	return list
}
