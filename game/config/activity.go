package config

import (
	"sync"

	"gohappy/data"
	"utils"
)

//ActivityMap 活动列表
var ActivityMap *sync.Map

//InitActivity 启动初始化
func InitActivity() {
	ActivityMap = new(sync.Map)
	l := data.GetActivityList()
	for _, v := range l {
		SetActivity(v)
	}
}

//InitActivity2 启动初始化
func InitActivity2() {
	ActivityMap = new(sync.Map)
}

//GetActivitys2 同步数据时获取
func GetActivitys2() map[string]data.Activity {
	m := make(map[string]data.Activity)
	ActivityMap.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(data.Activity)
		return true
	})
	return m
}

//GetActivitys 客户端获取消息列表
func GetActivitys() []data.Activity {
	list := make([]data.Activity, 0)
	ActivityMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.Activity); ok {
			if val.Del > 0 {
				return false
			}
			if val.EndTime.Before(utils.BsonNow()) {
				return false
			}
			list = append(list, val)
		}
		return true
	})
	return list
}

//GetActivitys1 客户端获取消息列表
func GetActivitys1() []data.Activity {
	list := make([]data.Activity, 0)
	ActivityMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.Activity); ok {
			if val.Del > 0 {
				return false
			}
			if val.EndTime.Before(utils.BsonNow().AddDate(0, 0, 1)) {
				return false //延长1天用于发货
			}
			list = append(list, val)
		}
		return true
	})
	return list
}

//DelActivity 删除元素
func DelActivity(k interface{}) {
	ActivityMap.Delete(k)
}

//SetActivity 添加新的活动
func SetActivity(v data.Activity) {
	if v.Del > 0 || v.EndTime.Before(utils.BsonNow()) {
		ActivityMap.Delete(v.Id)
	} else {
		ActivityMap.Store(v.Id, v)
	}
}

//GetActivity 获取指定活动
func GetActivity(id string) data.Activity {
	if v, ok := ActivityMap.Load(id); ok {
		//return v.(data.Activity)
		if val, ok := v.(data.Activity); ok {
			if val.Del > 0 {
				return data.Activity{}
			}
			if val.EndTime.Before(utils.BsonNow()) {
				return data.Activity{}
			}
			return val
		}
	}
	return data.Activity{}
}

//GetActivity1 获取指定活动
func GetActivity1(id string) data.Activity {
	if v, ok := ActivityMap.Load(id); ok {
		//return v.(data.Activity)
		if val, ok := v.(data.Activity); ok {
			if val.Del > 0 {
				return data.Activity{}
			}
			if val.EndTime.Before(utils.BsonNow().AddDate(0, 0, 1)) {
				return data.Activity{}
			}
			return val
		}
	}
	return data.Activity{}
}