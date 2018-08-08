package handler

import (
	"gohappy/data"
	"gohappy/game/config"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

func packLogActivityMsg(v *data.LogActivity) (msg *pb.Activity) {
	msg = &pb.Activity{
		Id:        v.Actid,
		JoinTime:   utils.Time2LocalStr(v.Jtime),
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
	}
	return
}

//PackActivity 打包消息
func PackActivity(msg *pb.SActivity) {
	list := config.GetActivitys()
	for _, v := range list {
		body := packActivityMsg(v)
		for _, val := range msg.List {
			if v.Id != val.Id {
				continue
			}
			body.JoinTime = val.JoinTime
		}
		msg.List = append(msg.List, body)
	}
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
		Etime: act.EndTime,
		Userid: arg.GetSelfid(),
	}
	if joinAct.Has() {
		msg.Error = pb.ActRepeatJoin
		return
	}
	joinAct.Jtime = utils.Stamp2Time(utils.TimestampTomorrow()).Local()
	if !joinAct.Save() {
		glog.Errorf("joinAct err : %#v, arg %#v", joinAct, arg)
		msg.Error = pb.ActRepeatJoin
		return
	}
	msg.Actid = arg.GetActid()
	msg.JoinTime = utils.Time2LocalStr(joinAct.Jtime)
	return
}

//StatActivity 统计符合条件的玩家
func StatActivity(arg *pb.AgentActivity) (msg []*pb.AgentActivityProfit) {
	act := config.GetActivity(arg.GetActid())
	if act.Id != arg.GetActid() {
		return
	}
	list, err := data.GetJoinActivityList(arg, act.Type)
	if err != nil {
		glog.Errorf("GetJoinActivityList error %v, arg %#v", err, arg)
		return
	}
	if len(list) == 0 {
		return
	}
	for _, v := range list {
		switch act.Type {
		case int32(pb.ACT_TYPE0):
			num, err := data.StatActivity(v.Userid, act.Type)
			glog.Debugf("type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
		case int32(pb.ACT_TYPE1):
			num, err := data.StatActivity(v.Userid, act.Type)
			glog.Debugf("type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
		case int32(pb.ACT_TYPE2):
			num, err := data.StatActivity(v.Userid, act.Type)
			glog.Debugf("type %d, userid %s, num %d, err %v", act.Type, v.Userid, num, err)
		default:
			return
		}
	}
	return
}