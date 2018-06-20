package data

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//TODO 添加房间id,人数同步后台显示

//Game 游戏金币房间配置
type Game struct {
	Id     string    `bson:"_id" json:"id"`        //unique ID
	Gtype  int32     `bson:"gtype" json:"gtype"`   //游戏类型1 niu,2 san,3 jiu
	Rtype  int32     `bson:"rtype" json:"rtype"`   //房间类型0免佣,1抽佣
	Dtype  int32     `bson:"dtype" json:"dtype"`   //桌子类型
	Ltype  int32     `bson:"ltype" json:"ltype"`   //彩票类型1bjpk10,1mlaft
	Name   string    `bson:"name" json:"name"`     //房间名称
	Status uint32    `bson:"status" json:"status"` //房间状态1打开,2关闭,3隐藏
	Count  uint32    `bson:"count" json:"count"`   //房间限制人数
	Ante   uint32    `bson:"ante" json:"ante"`     //房间底分
	Cost   uint32    `bson:"cost" json:"cost"`     //房间抽佣百分比
	Vip    uint32    `bson:"vip" json:"vip"`       //房间vip限制
	Chip   uint32    `bson:"chip" json:"chip"`     //房间进入筹码限制
	Deal   bool      `bson:"deal" json:"deal"`     //房间是否可以上庄
	Carry  uint32    `bson:"carry" json:"carry"`   //房间上庄最小携带筹码限制
	Down   uint32    `bson:"down" json:"down"`     //房间下庄最小携带筹码限制
	Top    uint32    `bson:"top" json:"Top"`       //房间下庄最大携带筹码限制
	Sit    uint32    `bson:"sit" json:"sit"`       //房间内坐下限制
	Del    int       `bson:"del" json:"del"`       //是否移除
	Node   string    `bson:"node" json:"node"`     //所在节点(game.huiyin1|game.huiyin2)
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
	//Num   uint32    `bson:"num" json:"num"`      //启动房间数量
	Minimum int64 `bson:"minimum" json:"minimum"` //房间最低限制
	Maximum int64 `bson:"maximum" json:"maximum"` //房间最高限制
	Pub     bool  `bson:"pub" json:"pub"`         //公开展示
	Mode     uint32 `bson:"mode" json:"mode"`      //模式，0普通，1疯狂
	Multiple uint32 `bson:"multiple" json:"multiple"` //倍数，0低，1中，2高
}

func GetGameList() []Game {
	var list []Game
	ListByQ(Games, bson.M{"del": 0}, &list)
	return list
}
