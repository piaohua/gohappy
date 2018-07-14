package data

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

const (
	NOTICE_TYPE0 = 0 //购买消息
	NOTICE_TYPE1 = 1 //公告消息
	NOTICE_TYPE2 = 2 //广播消息
	NOTICE_TYPE3 = 3 //系统消息
)

const (
	NOTICE_ACT_TYPE0 = 0 //无操作消息
	NOTICE_ACT_TYPE1 = 1 //支付消息
	NOTICE_ACT_TYPE2 = 2 //活动消息
)

//公告
type Notice struct {
	Id      string    `bson:"_id" json:"id"`          //公告ID
	Userid  string    `bson:"userid" json:"userid"`   //玩家
	Rtype   int       `bson:"rtype" json:"rtype"`     //类型,1=公告消息,2=广播消息
	Atype   int32     `bson:"atype" json:"atype"`     //分包类型
	Acttype int       `bson:"acttype" json:"acttype"` //操作类型,0=无操作,1=支付,2=活动
	Top     int       `bson:"top" json:"top"`         //置顶
	Num     int       `bson:"num" json:"num"`         //广播次数
	Del     int       `bson:"del" json:"del"`         //是否移除
	Content string    `bson:"content" json:"content"` //广播内容
	Etime   time.Time `bson:"etime" json:"etime"`     //过期时间
	Ctime   time.Time `bson:"ctime" json:"ctime"`     //创建时间
}

//Save 保存消息记录
func (t *Notice) Save() bool {
	t.Id = bson.NewObjectId().Hex()
	t.Ctime = bson.Now()
	t.Etime = bson.Now().AddDate(0, 0, 7)
	return Insert(Notices, t)
}

//GetNoticeList 获取公共消息(系统消息)
func GetNoticeList(rtype int) []Notice {
	var list []Notice
	//q := bson.M{"del": 0, "rtype": rtype,
	//	"userid": "",
	//	"etime":  bson.M{"$gt": bson.Now()}}
	q := bson.M{"del": 0, "userid": "",
		"etime":  bson.M{"$gt": bson.Now()}}
	ListByQ(Notices, q, &list)
	return list
}

//GetLogNotices 获取玩家消息记录
func GetLogNotices(userid string, page int) ([]*Notice, error) {
	pageSize := 30 //TODO 优化数据量过大情况
	skipNum, sortFieldR := parsePageAndSort(page, pageSize, "ctime", false)
	var list = make([]*Notice, 0)
	err := Notices.
		Find(bson.M{"userid": userid}).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("none record")
	}
	return list, nil
}
