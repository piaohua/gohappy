package main

import (
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//生成一个牌桌邀请码,全列表中唯一
func (a *RoomActor) genCode() (s string) {
	s = utils.RandStr(6)
	//是否已经存在
	if _, ok := a.codes[s]; ok {
		return a.genCode() //重复尝试,TODO:一定次数后放弃尝试
	}
	return
}

//生成房间ID
//func (a *RoomActor) genDesk(arg *pb.GenDesk, ctx actor.Context) {
//	glog.Debugf("genDesk Rtype: %d, Gtype: %d", arg.Rtype, arg.Gtype)
//	rsp := new(pb.GenedDesk)
//	rsp.Roomid = a.uniqueid.GenID()
//	//TODO 百人,私人
//	//rsp.Code = a.genCode()
//	//响应
//	ctx.Respond(rsp)
//}

//关闭房间
func (a *RoomActor) closeDesk(arg *pb.CloseDesk, ctx actor.Context) {
	//glog.Debugf("CloseDesk router %#v", a.router)
	//glog.Debugf("CloseDesk count %#v", a.count)
	//glog.Debugf("CloseDesk rules %#v", a.rules)
	delete(a.count, arg.Roomid)
	delete(a.codes, arg.Code)
	delete(a.rules, arg.Unique)
	delete(a.rooms, arg.Roomid)
	glog.Debugf("CloseDesk %d", len(a.rooms))
	//响应
	//rsp := new(pb.ClosedDesk)
	//ctx.Respond(rsp)
}

//离开房间
func (a *RoomActor) leaveDesk(arg *pb.LeaveDesk, ctx actor.Context) {
	//移除
	if _, ok := a.router[arg.Userid]; ok {
		delete(a.router, arg.Userid)
		if n, ok := a.count[arg.Roomid]; ok && n > 0 {
			a.count[arg.Roomid] = n - 1
		}
	}
	//响应
	//rsp := new(pb.LeftDesk)
	//ctx.Respond(rsp)
}

//加入房间
func (a *RoomActor) joinDesk(arg *pb.JoinDesk, ctx actor.Context) {
	//房间数据变更
	if _, ok := a.router[arg.Userid]; !ok {
		a.router[arg.Userid] = arg.Roomid
		a.count[arg.Roomid]++
	}
	//响应
	//rsp := new(pb.JoinedDesk)
	//ctx.Respond(rsp)
}

//添加房间
func (a *RoomActor) addDesk(arg *pb.AddDesk, ctx actor.Context) {
	glog.Debugf("addDesk Rtype: %d, Gtype: %d", arg.Rtype, arg.Gtype)
	rsp := new(pb.AddedDesk)
	//已经存在
	if _, ok := a.rules[arg.Unique]; ok && arg.Unique != "" {
		glog.Errorf("addDesk err Rtype: %d, Gtype: %d ",
			arg.Rtype, arg.Gtype)
		glog.Errorf("addDesk err Roomid: %s, Unique: %s",
			arg.Roomid, arg.Unique)
		rsp.Error = pb.Failed
		ctx.Respond(rsp)
		return
	}
	rsp.Roomid = a.uniqueid.GenID()
	switch arg.Rtype {
	case int32(pb.ROOM_TYPE1): //私人
		//邀请码
		rsp.Code = a.genCode()
		//私人房间
		a.codes[rsp.Code] = rsp.Roomid
	}
	//响应消息
	ctx.Respond(rsp)
	//添加房间
	a.rooms[rsp.Roomid] = arg.Desk
	a.rules[arg.Unique] = rsp.Roomid
}

//匹配房间
func (a *RoomActor) matchDesk(msg *pb.MatchDesk, ctx actor.Context) {
	glog.Debugf("matchDesk codes: %#v", a.codes)
	rsp := new(pb.MatchedDesk)
	rsp.Rtype = msg.Rtype
	rsp.Gtype = msg.Gtype
	rsp.Dtype = msg.Dtype
	rsp.Ltype = msg.Ltype
	if v, ok := a.codes[msg.Code]; ok &&
		msg.Code != "" {
		msg.Roomid = v
	}
	if v, ok := a.rooms[msg.Roomid]; ok &&
		msg.Roomid != "" {
		rsp.Desk = v
		ctx.Respond(rsp)
		return
	}
	rsp.Error = pb.MatchFail
	ctx.Respond(rsp)
}
