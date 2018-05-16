package config

import (
	"sync"

	"gohappy/data"
)

// 环境变量
var EnvMap *sync.Map

// 启动初始化
func InitEnv() {
	EnvMap = new(sync.Map)
	list := data.GetEnvList()
	for _, v := range list {
		SetEnv(v.Key, v.Value)
	}
}

// 启动初始化
func InitEnv2() {
	EnvMap = new(sync.Map)
}

// 设置元素,不操作数据库
func SetEnv(k interface{}, v interface{}) {
	EnvMap.Store(k, v)
}

// 删除元素
func DelEnv(k interface{}) {
	EnvMap.Delete(k)
}

// 全部元素
func GetEnvs() map[string]int32 {
	m := make(map[string]int32)
	EnvMap.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(int32)
		return true
	})
	return m
}

//获取变量,变量默认值设置
func GetEnv(k interface{}) int32 {
	key, ok := k.(string)
	if !ok {
		return 0
	}
	// 存在元素
	if v, ok := EnvMap.Load(k); ok {
		return v.(int32)
	}
	//默认值
	switch key {
	case data.ENV1:
		return 0 //注册赠送钻石
	case data.ENV2:
		return 0 //注册赠送金币
	case data.ENV3:
		return 0 //注册赠送筹码
	case data.ENV4:
		return 0 //注册赠送房卡
	case data.ENV5:
		return 0 //绑定赠送钻石
	}
	return 0
}
