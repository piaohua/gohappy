package handler

import (
	"gohappy/data"
	"gohappy/game/config"
	"gohappy/pb"

	"github.com/globalsign/mgo/bson"
)

//LuckyUpdateMsg 任务变动消息
func LuckyUpdateMsg(num uint32, gtype, luckyid int32,
	userid string) (msg *pb.LuckyUpdate) {
	msg = new(pb.LuckyUpdate)
	msg.Userid = userid
	msg.Num = num
	msg.Gtype = gtype
	msg.Luckyid = luckyid
	return
}

//SetLuckyList 配置任务数据,测试数据
func SetLuckyList() {
	//五小牛，同花顺，炸弹牛，五花牛，同花牛，葫芦牛，顺子牛
	NewLucky(110, int32(pb.NIU), "顺子牛", 50, 100, 100)
	NewLucky(120, int32(pb.NIU), "葫芦牛", 50, 200, 200)
	NewLucky(130, int32(pb.NIU), "同花牛", 50, 300, 300)
	NewLucky(140, int32(pb.NIU), "五花牛", 50, 400, 400)
	NewLucky(150, int32(pb.NIU), "炸弹牛", 50, 500, 500)
	NewLucky(160, int32(pb.NIU), "同花顺", 50, 600, 600)
	NewLucky(170, int32(pb.NIU), "五小牛", 50, 700, 700)
}

//NewLucky 添加新任务
func NewLucky(luckyid, gtype int32, name string, count uint32,
	diamond, coin int64) {
	t := data.Lucky{
		//ID:      bson.NewObjectId().String(),
		ID:      data.ObjectIdString(bson.NewObjectId()),
		Ctime:   bson.Now(),
		Luckyid: luckyid,
		Name:    name,
		Gtype:   gtype,
		Count:   count,
		//Diamond: diamond,
		Coin: coin,
	}
	config.SetLucky(t)
	t.Save()
}
