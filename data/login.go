package data

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//LoginPrize 连续登录奖励配置
type LoginPrize struct {
	ID      string    `bson:"_id" json:"id"`          //unique ID
	Day     uint32    `bson:"day" json:"day"`         //unique
	Del     int       `bson:"del" json:"del"`         //是否移除
	Coin    int64     `bson:"coin" json:"coin"`       //金币奖励
	Diamond int64     `bson:"diamond" json:"diamond"` //钻石奖励
	Ctime   time.Time `bson:"ctime" json:"ctime"`     //创建时间
}

//GetLoginPrizeList 获取连续登录奖励配置
func GetLoginPrizeList() []LoginPrize {
	var list []LoginPrize
	ListByQ(LoginPrizes, bson.M{"del": 0}, &list)
	return list
}
