package data

import "time"

//vip
type Vip struct {
	Id     string    `bson:"_id" json:"id"`        //ID
	Name   string    `bson:"name" json:"name"`     //名字
	Level  int       `bson:"level" json:"level"`   //等级
	Number uint32    `bson:"number" json:"number"` //等级充值金额数量限制(分)
	Pay    uint32    `bson:"pay" json:"pay"`       //充值赠送百分比5=赠送充值的5%
	Del    int       `bson:"del" json:"del"`       //是否移除
	Etime  time.Time `bson:"etime" json:"etime"`   //过期时间
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
}

func GetVipList() []Vip {
	var list []Vip
	ListByQ(Vips, nil, &list)
	return list
}

//Save 写入数据库
func (t *Vip) Save() bool {
	//t.Ctime = bson.Now()
	return Insert(Vips, t)
}