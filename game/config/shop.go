package config

import (
	"sort"
	"sync"

	"gohappy/data"
	"utils"
)

// 商城列表
var ShopMap *sync.Map

// 启动初始化
func InitShop() {
	ShopMap = new(sync.Map)
	l := data.GetShopList()
	for _, v := range l {
		SetShop(v)
	}
}

// 启动初始化
func InitShop2() {
	ShopMap = new(sync.Map)
}

// 同步数据获取
func GetShops2() map[string]data.Shop {
	m := make(map[string]data.Shop)
	ShopMap.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(data.Shop)
		return true
	})
	return m
}

// 客户端获取商品列表
func GetShops() []data.Shop {
	list := make([]data.Shop, 0)
	ShopMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.Shop); ok {
			if val.Del > 0 {
				return false
			}
			if val.Etime.Before(utils.BsonNow()) {
				return false
			}
			list = append(list, val)
		}
		return true
	})
	sort.Slice(list, func(i, j int) bool {
		return list[i].Price < list[j].Price
	})
	return list
}

// 删除元素
func DelShop(k interface{}) {
	ShopMap.Delete(k)
}

// 添加新的商品
func SetShop(v data.Shop) {
	if v.Del > 0 || v.Etime.Before(utils.BsonNow()) {
		ShopMap.Delete(v.Id)
	} else {
		ShopMap.Store(v.Id, v)
	}
}

// 获取商品
func GetShop(id string) data.Shop {
	if v, ok := ShopMap.Load(id); ok {
		return v.(data.Shop)
	}
	return data.Shop{}
}
