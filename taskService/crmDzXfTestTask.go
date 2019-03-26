package taskService

import "reflect"

type crmDzXfTestTask struct {
}

func (ct *crmDzXfTestTask) getCacheIdList() []string {
	tc := TaskCache{}
	list := make([]string, 0)
	for _, c := range tc.GetTaskList() {
		switch reflect.TypeOf(c.Config).String() {
		case "*object.CrmDzXfTestConfigData":
			list = append(list, c.Config.GetConfigId())
		}
	}
	return list
}
