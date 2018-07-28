package handler

import (
	"gohappy/data"
	"gohappy/game/config"
	"gohappy/pb"

	"github.com/globalsign/mgo/bson"
)

//LuckyUpdateMsg 任务变动消息
func LuckyUpdateMsg(num uint32, gtype int32,
	userid string) (msg *pb.LuckyUpdate) {
	msg = new(pb.LuckyUpdate)
	msg.Userid = userid
	msg.Num = num
	msg.Gtype = gtype
	return
}

//SetLuckyList 配置任务数据,测试数据
func SetLuckyList() {
	//五小牛，同花顺，炸弹牛，五花牛，同花牛，葫芦牛，顺子牛
	NewLucky(110, "顺子牛", 20, 100, 100)
	NewLucky(120, "葫芦牛", 20, 200, 200)
	NewLucky(130, "同花牛", 20, 300, 300)
	NewLucky(140, "五花牛", 20, 400, 400)
	NewLucky(150, "炸弹牛", 20, 500, 500)
	NewLucky(160, "同花顺", 20, 600, 600)
	NewLucky(170, "五小牛", 20, 700, 700)
}

//NewLucky 添加新任务
func NewLucky(luckyid int32, name string, count uint32,
	diamond, coin int64) {
	t := data.Lucky{
		//ID:      bson.NewObjectId().String(),
		ID:      data.ObjectIdString(bson.NewObjectId()),
		Ctime:   bson.Now(),
		Luckyid:  luckyid,
		Name:    name,
		Count:   count,
		Diamond: diamond,
		//Coin: coin,
	}
	config.SetLucky(t)
	t.Save()
}
