package data

import (
	"errors"
	"time"

	"gohappy/pb"

	"github.com/globalsign/mgo/bson"
)

//Activity 活动配置
type Activity struct {
	Id      string    `bson:"_id" json:"id"`          //ID
	Type   int32     `bson:"type" json:"type"`     //类型0,1,2
	Del     int       `bson:"del" json:"del"`         //是否移除
	Title string    `bson:"title" json:"title"` //标题内容
	Content string    `bson:"content" json:"content"` //活动内容
	StartTime   time.Time `bson:"start_time" json:"start_time"`     //开始时间
	EndTime   time.Time `bson:"end_time" json:"end_time"`     //结束时间
	Ctime   time.Time `bson:"ctime" json:"ctime"`     //创建时间
}

//Save 保存消息记录
func (t *Activity) Save() bool {
	t.Id = ObjectIdString(bson.NewObjectId())
	t.Ctime = bson.Now()
	return Insert(Activitys, t)
}

//GetActivityList 获取列表
func GetActivityList() []Activity {
	var list []Activity
	q := bson.M{"del": 0, "end_time":  bson.M{"$gt": bson.Now()}}
	ListByQ(Activitys, q, &list)
	return list
}

//LogActivity 玩家参与活动记录
type LogActivity struct {
	//Id      string    `bson:"_id" json:"id"`          //ID
	Userid  string    `bson:"userid" json:"userid"`   //玩家
	Actid   string    `bson:"actid" json:"actid"` //activity id
	Prize   int64     `bson:"prize" json:"prize"` //奖励数量
	Num     int32     `bson:"num" json:"num"`     //完成次数
	Etime   time.Time `bson:"etime" json:"etime"`     //过期时间
	Jtime   time.Time `bson:"jtime" json:"jtime"`    //参与时间
	Utime   time.Time `bson:"utime" json:"utime"`   //update Time
	Ctime   time.Time `bson:"ctime" json:"ctime"`     //创建时间
}

//Save 保存消息记录
func (t *LogActivity) Save() bool {
	//t.Id = bson.NewObjectId().String()
	//t.Id = ObjectIdString(bson.NewObjectId())
	t.Ctime = bson.Now()
	//t.Etime = bson.Now().AddDate(0, 0, 7)
	return Insert(LogActivitys, t)
}

//Get 查询获取记录
func (t *LogActivity) Get() {
	GetByQ(LogActivitys, bson.M{"userid": t.Userid, "actid": t.Actid}, t)
}

//Has 记录是否存在
func (t *LogActivity) Has() bool {
	return Has(LogActivitys, bson.M{"userid": t.Userid, "actid": t.Actid})
}

//Update 更新记录
func (t *LogActivity) Update() bool {
	t.Utime = bson.Now()
	return Update(LogActivitys, bson.M{"userid": t.Userid, "actid": t.Actid},
		bson.M{"$set": bson.M{"utime": t.Utime}, "$inc": bson.M{"prize": t.Prize, "num": t.Num}})
}

//GetLogActivitys 获取玩家记录
func GetLogActivitys(userid string, page int, ids []string) ([]*LogActivity, error) {
	pageSize := 30
	skipNum, sortFieldR := parsePageAndSort(page, pageSize, "ctime", false)
	var list = make([]*LogActivity, 0)
	q := bson.M{"userid": userid}
	if len(ids) != 0 {
		q["actid"] = bson.M{"$in": ids}
	}
	err := LogActivitys.
		Find(q).
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

//getLogActivityList 获取参加活动玩家记录
func getLogActivityList(page int, q bson.M) ([]*LogActivity, error) {
	pageSize := 30
	skipNum, sortFieldR := parsePageAndSort(page, pageSize, "ctime", false)
	var list = make([]*LogActivity, 0)
	err := LogActivitys.
		Find(q).
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

//GetJoinActivityList 代理活动参加列表
func GetJoinActivityList(arg *pb.AgentActivity, Type int32) ([]*LogActivity, error) {
	q := bson.M{"actid": arg.GetActid()}
	switch Type {
	case int32(pb.ACT_TYPE0):
		q["prize"] = bson.M{"$lt": 50000000}
	case int32(pb.ACT_TYPE1):
		q["num"] = bson.M{"$eq": 0}
	case int32(pb.ACT_TYPE2):
		q["num"] = bson.M{"$eq": 0}
	default:
		return nil, errors.New("type error")
	}
	return getLogActivityList(int(arg.GetPage()), q)
}

//StatActivity 统计数据
func StatActivity(userid string, Type int32) (int64, error) {
	switch Type {
	case int32(pb.ACT_TYPE0):
		//q["prize"] = bson.M{"$lt": 50000000}
	case int32(pb.ACT_TYPE1):
		//q["num"] = bson.M{"$eq": 0}
	case int32(pb.ACT_TYPE2):
		//q["num"] = bson.M{"$eq": 0}
	default:
		return 0, errors.New("type error")
	}
	return 0, nil
}