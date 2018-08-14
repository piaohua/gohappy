package handler

import (
	"errors"

	"gohappy/data"
	"gohappy/game/config"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
	"time"
	"github.com/globalsign/mgo/bson"
)

func packLogActivityMsg(v *data.LogActivity) (msg *pb.Activity) {
	msg = &pb.Activity{
		Id:        v.Actid,
		JoinTime:   utils.Time2LocalStr(v.Jtime),
	}
	switch v.Type {
	case int32(pb.ACT_TYPE0):
		if v.Num >= 500000 {
			msg.Over = true
		}
	case int32(pb.ACT_TYPE1):
		//TODO 当天是否完成
	case int32(pb.ACT_TYPE2):
		if v.Num >= 15 {
			msg.Over = true
		}
	}
	return
}

//PackUserActivity 打包玩家消息
func PackUserActivity(arg *pb.CActivity) (msg *pb.SActivity) {
	msg = new(pb.SActivity)
	list2 := config.GetActivitys()
	var ids []string
	for _, v := range list2 {
		ids = append(ids, v.Id)
	}
	list, err := data.GetLogActivitys(arg.Userid, int(arg.Page), ids)
	if err != nil {
		glog.Errorf("get notice err : %v, arg %#v", err, arg)
		return
	}
	for _, v := range list {
		body := packLogActivityMsg(v)
		msg.List = append(msg.List, body)
	}
	return
}

func packActivityMsg(v data.Activity) (msg *pb.Activity) {
	msg = &pb.Activity{
		Id:        v.Id,
		Title:     v.Title,
		Content:   v.Content,
		StartTime: utils.Time2LocalStr(v.StartTime),
		EndTime:   utils.Time2LocalStr(v.EndTime),
		Type:      v.Type,
	}
	return
}

//PackActivity 打包消息
func PackActivity(msg *pb.SActivity) {
	list := config.GetActivitys()
	msgList := make([]*pb.Activity, 0)
	for _, v := range list {
		body := packActivityMsg(v)
		for _, val := range msg.List {
			if v.Id != val.Id {
				continue
			}
			body.JoinTime = val.JoinTime
			body.Over = val.Over
		}
		msgList = append(msgList, body)
	}
	msg.List = msgList
}

//JoinActivity 参加消息
func JoinActivity(arg *pb.CJoinActivity) (msg *pb.SJoinActivity) {
	msg = new(pb.SJoinActivity)
	act := config.GetActivity(arg.GetActid())
	if act.Id != arg.GetActid() {
		msg.Error = pb.ActidError
		return
	}
	joinAct := &data.LogActivity{
		Actid: act.Id,
		Type:  act.Type,
		Etime: act.EndTime,
		Userid: arg.GetSelfid(),
	}
	if joinAct.Has() {
		msg.Error = pb.ActRepeatJoin
		return
	}
	joinAct.Jtime = utils.Stamp2Time(utils.TimestampTomorrow()).Local() //报名后第二天开始
	if !joinAct.Save() {
		glog.Errorf("joinAct err : %#v, arg %#v", joinAct, arg)
		msg.Error = pb.ActRepeatJoin
		return
	}
	msg.Actid = arg.GetActid()
	msg.JoinTime = utils.Time2LocalStr(joinAct.Jtime)
	return
}

//StatActivity 统计活动奖励发放给符合条件的玩家
func StatActivity(arg *pb.AgentActivity) (msg []*pb.AgentActivityProfit, err error) {
	act := config.GetActivity(arg.GetActid())
	if act.Id != arg.GetActid() {
		return nil, errors.New("actid error")
	}
	list, err := data.GetJoinActivityList(arg, act)
	if err != nil {
		glog.Errorf("GetJoinActivityList error %v, arg %#v", err, arg)
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("none data")
	}
	glog.Debugf("list %#v, arg %#v, act %#v", list, arg, act)
	for _, v := range list {
		switch act.Type {
		case int32(pb.ACT_TYPE0):
			num, err := data.StatActivity(v, act)
			glog.Debugf("type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
			if err != nil {
				glog.Errorf("StatActivity type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
				continue
			}
			if num >= 10000000 {
				msg2 := statActivityMsg(v.Userid, act.Id, act.Title, int32(pb.LOG_TYPE57), 50000, 0)
				msg = append(msg, msg2)
			}
		case int32(pb.ACT_TYPE1):
			num, err := data.StatActivity(v, act)
			glog.Debugf("type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
			if err != nil {
				glog.Errorf("StatActivity type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
				continue
			}
			if num >= 10000000 {
				msg2 := statActivityMsg(v.Userid, act.Id, act.Title, int32(pb.LOG_TYPE58), 50000, 0)
				msg = append(msg, msg2)
			}
			if num >= 20000000 {
				msg2 := statActivityMsg(v.Userid, act.Id, act.Title, int32(pb.LOG_TYPE58), 100000, 0)
				msg = append(msg, msg2)
			}
		case int32(pb.ACT_TYPE2):
			num, err := data.StatActivity(v, act)
			glog.Debugf("type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
			if err != nil {
				glog.Errorf("StatActivity type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
				continue
			}
			if num >= 5 && v.Num != 5 {
				msg2 := statActivityMsg(v.Userid, act.Id, act.Title, int32(pb.LOG_TYPE59), 100000, 5)
				msg = append(msg, msg2)
			}
			if num >= 10 && v.Num != 10 {
				msg2 := statActivityMsg(v.Userid, act.Id, act.Title, int32(pb.LOG_TYPE59), 200000, 10)
				msg = append(msg, msg2)
			}
		default:
			return nil, errors.New("type error")
		}
	}
	return
}

func statActivityMsg(userid, actid, title string, Type int32,
	profit int64, num uint32) (msg *pb.AgentActivityProfit) {
	msg = &pb.AgentActivityProfit{
		Userid: userid,
		Profit: profit,
		Type: Type,
		Actid: actid,
		Title: title,
		Num: num,
	}
	return
}

//StatActivityUpdate 更新活动
func StatActivityUpdate(arg *pb.AgentActivityProfit) {
	d := &data.LogActivity{
		Userid: arg.GetUserid(),
		Actid: arg.GetActid(),
		Prize: arg.GetProfit(),
		Num: arg.GetNum(),
	}
	if !d.Update() {
		glog.Errorf("StatActivityUpdate filed %#v", d)
	}
}


//SetActivityList 配置活动数据,测试数据
func SetActivityList() {
	var startTime, endTime time.Time
	startTime = utils.TimestampTodayTime()
	endTime = startTime.AddDate(0, 0, 30)
	NewActivity(int32(pb.ACT_TYPE0), "翻倍奖", "翻倍奖", startTime, endTime)
	endTime = startTime.AddDate(0, 0, 60)
	NewActivity(int32(pb.ACT_TYPE1), "增加奖", "增加奖", startTime, endTime)
	endTime = startTime.AddDate(0, 0, 90)
	NewActivity(int32(pb.ACT_TYPE2), "激活奖", "激活奖", startTime, endTime)
}

//NewActivity 添加新活动
func NewActivity(Type int32, title, content string, startTime, endTime time.Time) {
	t := data.Activity{
		//Id:      bson.NewObjectId().String(),
		Id:      data.ObjectIdString(bson.NewObjectId()),
		Ctime:   bson.Now(),
		Type:  Type,
		Title: title,
		Content: content,
		StartTime: startTime,
		EndTime: endTime,
	}
	config.SetActivity(t)
	t.Save()
}