package data

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Lucky 幸运星
type Lucky struct {
	ID      string    `bson:"_id" json:"id"`          //unique ID
	Luckyid int32     `bson:"luckyid" json:"luckyid"` //unique
	Gtype   int32     `bson:"gtype" json:"gtype"`     //游戏类型1 niu,2 san,3 jiu
	Name    string    `bson:"name" json:"name"`       //名称
	Count   uint32    `bson:"count" json:"count"`     //任务数值
	Diamond int64     `bson:"diamond" json:"diamond"` //钻石奖励
	Coin    int64     `bson:"coin" json:"coin"`       //金币奖励
	Del     int       `bson:"del" json:"del"`         //是否移除
	Ctime   time.Time `bson:"ctime" json:"ctime"`     //创建时间
}

//GetLuckyList 幸运星
func GetLuckyList() []Lucky {
	var list []Lucky
	ListByQ(Luckys, bson.M{"del": 0}, &list)
	return list
}

//LuckyInfo 幸运星信息
type LuckyInfo struct {
	Luckyid int32  `bson:"luckyid" json:"luckyid"` //unique
	Num     uint32 `bson:"num" json:"num"`         //完成数值
}

//Save 写入数据库
func (t *Lucky) Save() bool {
	//t.Ctime = bson.Now()
	return Insert(Luckys, t)
}
