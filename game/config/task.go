package config

import (
	"sort"
	"sync"

	"gohappy/data"
)

//TaskMap 游戏列表
var TaskMap *sync.Map

//TaskList 游戏列表
var TaskList []data.Task

//InitTask 启动初始化
func InitTask() {
	TaskMap = new(sync.Map)
	l := data.GetTaskList()
	for _, v := range l {
		SetTask(v)
	}
}

//InitTask2 启动初始化
func InitTask2() {
	TaskMap = new(sync.Map)
}

//GetTasks2 同步时获取列表
func GetTasks2() map[int32]data.Task {
	m := make(map[int32]data.Task)
	TaskMap.Range(func(k, v interface{}) bool {
		m[k.(int32)] = v.(data.Task)
		return true
	})
	return m
}

//GetTasks 获取任务列表
func GetTasks() []data.Task {
	list := make([]data.Task, 0)
	TaskMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.Task); ok {
			if val.Del > 0 {
				return false
			}
			list = append(list, val)
		}
		return true
	})
	return list
}

//GetOrderTasks 获取有序任务列表
func GetOrderTasks() []data.Task {
	return TaskList
}

//DelTask 删除元素
func DelTask(k interface{}) {
	TaskMap.Delete(k)
	sortTask()
}

//SetTask 添加或更新任务,类型做唯一key
func SetTask(v data.Task) {
	if v.Del > 0 {
		TaskMap.Delete(v.Taskid)
	} else {
		TaskMap.Store(v.Taskid, v)
	}
	sortTask()
}

//GetTask 获取指定任务
func GetTask(taskid int32) data.Task {
	if v, ok := TaskMap.Load(taskid); ok {
		return v.(data.Task)
	}
	return data.Task{}
}

func sortTask() {
	list := GetTasks()
	sort.Slice(list, func(i, j int) bool {
		return list[i].Taskid < list[j].Taskid
	})
	TaskList = list
}
