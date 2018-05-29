package config

import (
	"sync"

	"gohappy/data"
)

//LoginMap 游戏列表
var LoginMap *sync.Map

//InitLogin 启动初始化
func InitLogin() {
	LoginMap = new(sync.Map)
	l := data.GetLoginPrizeList()
	for _, v := range l {
		SetLogin(v)
	}
}

//InitLogin2 启动初始化
func InitLogin2() {
	LoginMap = new(sync.Map)
}

//GetLogins2 同步时获取列表
func GetLogins2() map[uint32]data.LoginPrize {
	m := make(map[uint32]data.LoginPrize)
	LoginMap.Range(func(k, v interface{}) bool {
		m[k.(uint32)] = v.(data.LoginPrize)
		return true
	})
	return m
}

//GetLogins 获取任务列表
func GetLogins() []data.LoginPrize {
	list := make([]data.LoginPrize, 0)
	LoginMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.LoginPrize); ok {
			if val.Del > 0 {
				return false
			}
			list = append(list, val)
		}
		return true
	})
	return list
}

//DelLogin 删除元素
func DelLogin(k interface{}) {
	LoginMap.Delete(k)
}

//SetLogin 添加或更新任务,类型做唯一key
func SetLogin(v data.LoginPrize) {
	if v.Del > 0 {
		LoginMap.Delete(v.Day)
	} else {
		LoginMap.Store(v.Day, v)
	}
}

//GetLogin 获取指定任务
func GetLogin(day uint32) data.LoginPrize {
	if v, ok := LoginMap.Load(day); ok {
		return v.(data.LoginPrize)
	}
	return data.LoginPrize{}
}
