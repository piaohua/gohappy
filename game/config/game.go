package config

import (
	"sync"

	"gohappy/data"
)

// 游戏列表
var GameMap *sync.Map

// 启动初始化
func InitGame() {
	GameMap = new(sync.Map)
	l := data.GetGameList()
	for _, v := range l {
		SetGame(v)
	}
}

// 启动初始化
func InitGame2() {
	GameMap = new(sync.Map)
}

// 同步时获取列表
func GetGames2() map[string]data.Game {
	m := make(map[string]data.Game)
	GameMap.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(data.Game)
		return true
	})
	return m
}

// 获取商品列表
func GetGames() []data.Game {
	list := make([]data.Game, 0)
	GameMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.Game); ok {
			if val.Del > 0 {
				return false
			}
			list = append(list, val)
		}
		return true
	})
	return list
}

// 删除元素
func DelGame(k interface{}) {
	GameMap.Delete(k)
}

// 添加新的公告
func SetGame(v data.Game) {
	if v.Del > 0 {
		GameMap.Delete(v.Id)
	} else {
		GameMap.Store(v.Id, v)
	}
}

// 获取
func GetGame(id string) data.Game {
	if v, ok := GameMap.Load(id); ok {
		return v.(data.Game)
	}
	return data.Game{}
}
