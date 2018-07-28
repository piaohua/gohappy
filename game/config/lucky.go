package config

import (
	"sort"
	"sync"

	"gohappy/data"
)

//LuckyMap 幸运星列表
var LuckyMap *sync.Map

//LuckyList 幸运星列表
var LuckyList []data.Lucky

//InitLucky 启动初始化
func InitLucky() {
	LuckyMap = new(sync.Map)
	l := data.GetLuckyList()
	for _, v := range l {
		SetLucky(v)
	}
}

//InitLucky2 启动初始化
func InitLucky2() {
	LuckyMap = new(sync.Map)
}

//GetLuckys2 同步时获取列表
func GetLuckys2() map[int32]data.Lucky {
	m := make(map[int32]data.Lucky)
	LuckyMap.Range(func(k, v interface{}) bool {
		m[k.(int32)] = v.(data.Lucky)
		return true
	})
	return m
}

//GetLuckys 获取幸运星列表
func GetLuckys() []data.Lucky {
	list := make([]data.Lucky, 0)
	LuckyMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.Lucky); ok {
			if val.Del > 0 {
				return false
			}
			list = append(list, val)
		}
		return true
	})
	return list
}

//GetOrderLuckys 获取有序幸运星列表
func GetOrderLuckys() []data.Lucky {
	return LuckyList
}

//DelLucky 删除元素
func DelLucky(k interface{}) {
	LuckyMap.Delete(k)
	sortLucky()
}

//SetLucky 添加或更新幸运星,类型做唯一key
func SetLucky(v data.Lucky) {
	if v.Del > 0 {
		LuckyMap.Delete(v.Luckyid)
	} else {
		LuckyMap.Store(v.Luckyid, v)
	}
	sortLucky()
}

//GetLucky 获取指定幸运星
func GetLucky(luckyid int32) data.Lucky {
	if v, ok := LuckyMap.Load(luckyid); ok {
		return v.(data.Lucky)
	}
	return data.Lucky{}
}

func sortLucky() {
	list := GetLuckys()
	sort.Slice(list, func(i, j int) bool {
		return list[i].Luckyid < list[j].Luckyid
	})
	LuckyList = list
}
