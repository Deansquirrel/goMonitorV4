package taskService

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV4/object"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/robfig/cron"
	"sync"
)

var lock sync.Mutex

var cacheList map[string]*cache

func init() {
	cacheList = make(map[string]*cache)
}

type cache struct {
	Config    object.IConfigData
	Cron      *cron.Cron
	IsRunning bool
}

type TaskCache struct {
}

func (tc *TaskCache) GetTaskList() map[string]*cache {
	lock.Lock()
	defer lock.Unlock()
	return cacheList
}

func (tc *TaskCache) AddTask(id string, cache *cache) {
	lock.Lock()
	defer lock.Unlock()
	cacheList[id] = cache
	//------------------------------------------------------------------------------------------------------------------
	configStr, err := goToolCommon.GetJsonStr(cache.Config)
	if err != nil {
		log.Warn(fmt.Sprintf("Add Task，转换配置内容时遇到错误:%s", configStr))
	} else {
		log.Warn(fmt.Sprintf("Add Task:%s", configStr))
	}
	//------------------------------------------------------------------------------------------------------------------
}

func (tc *TaskCache) DelTask(id string) {
	lock.Lock()
	defer lock.Unlock()
	t, ok := cacheList[id]
	if ok {
		//------------------------------------------------------------------------------------------------------------------
		configStr, err := goToolCommon.GetJsonStr(t.Config)
		if err != nil {
			log.Warn(fmt.Sprintf("Del Task，转换配置内容时遇到错误:%s", configStr))
		} else {
			log.Warn(fmt.Sprintf("Del Task:%s", configStr))
		}
		//------------------------------------------------------------------------------------------------------------------
		t.IsRunning = false
		t.Cron.Stop()
		delete(cacheList, id)
	}
}

func (tc *TaskCache) UpdateTaskState(id string, isRunning bool) {
	lock.Lock()
	defer lock.Unlock()
	cache, ok := cacheList[id]
	if ok {
		cache.IsRunning = isRunning
	}
}
