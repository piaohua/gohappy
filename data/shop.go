package data

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

const (
	//物品类型
	DIAMOND uint32 = 1 //钻石
	COIN    uint32 = 2 //金币
	CARD    uint32 = 3 //房卡
	CHIP    uint32 = 4 //筹码
	VIP     uint32 = 5 //VIP
	//支付方式
	RMB uint32 = 1 //人民币
	DIA uint32 = 2 //钻石
)

//商城
type Shop struct {
	Id     string    `bson:"_id" json:"id"`        //购买ID
	Status int       `bson:"status" json:"status"` //物品状态,1=热卖
	Propid int       `bson:"propid" json:"propid"` //兑换的物品,1=钻石,2=金币
	Payway int       `bson:"payway" json:"payway"` //支付方式,1=RMB,,2=钻石
	Number uint32    `bson:"number" json:"number"` //兑换的数量
	Price  uint32    `bson:"price" json:"price"`   //支付价格(单位元)
	Name   string    `bson:"name" json:"name"`     //物品名字
	Info   string    `bson:"info" json:"info"`     //物品信息
	Del    int       `bson:"del" json:"del"`       //是否移除
	Etime  time.Time `bson:"etime" json:"etime"`   //过期时间
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
}

func GetShopList() []Shop {
	var list []Shop
	ListByQ(Shops, bson.M{"del": 0, "etime": bson.M{"$gt": bson.Now()}}, &list)
	return list
}

//Save 写入数据库
func (t *Shop) Save() bool {
	//t.Ctime = bson.Now()
	return Insert(Shops, t)
}