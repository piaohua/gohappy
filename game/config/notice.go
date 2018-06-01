package config

import (
	"sort"
	"sync"

	"gohappy/data"
	"utils"
)

// 公告列表
var NoticeMap *sync.Map

// 启动初始化
func InitNotice() {
	NoticeMap = new(sync.Map)
	l := data.GetNoticeList(data.NOTICE_TYPE1)
	for _, v := range l {
		SetNotice(v)
	}
}

// 启动初始化
func InitNotice2() {
	NoticeMap = new(sync.Map)
}

// 同步数据时获取
func GetNotices2() map[string]data.Notice {
	m := make(map[string]data.Notice)
	NoticeMap.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(data.Notice)
		return true
	})
	return m
}

// 客户端获取消息列表
func GetNotices(atype int32) []data.Notice {
	tops := make([]data.Notice, 0)
	list := make([]data.Notice, 0)
	NoticeMap.Range(func(k, v interface{}) bool {
		if val, ok := v.(data.Notice); ok {
			if val.Del > 0 {
				return false
			}
			if val.Etime.Before(utils.BsonNow()) {
				return false
			}
			if val.Top > 0 {
				tops = append(tops, val)
				return true
			}
			list = append(list, val)
		}
		return true
	})
	sort.Slice(tops, func(i, j int) bool {
		return tops[i].Ctime.After(tops[j].Ctime)
	})
	sort.Slice(list, func(i, j int) bool {
		return list[i].Ctime.After(list[j].Ctime)
	})
	return append(tops, list...)
}

// 删除元素
func DelNotice(k interface{}) {
	NoticeMap.Delete(k)
}

// 添加新的公告
func SetNotice(v data.Notice) {
	if v.Del > 0 || v.Etime.Before(utils.BsonNow()) {
		NoticeMap.Delete(v.Id)
	} else {
		NoticeMap.Store(v.Id, v)
	}
}
