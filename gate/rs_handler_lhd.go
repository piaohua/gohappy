package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//lhd请求处理
func (rs *RoleActor) handlerLhd(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CLHFreeEnterRoom:
		arg := msg.(*pb.CLHFreeEnterRoom)
		glog.Debugf("CLHFreeEnterRoom %#v", arg)
		rs.enterLHFree(arg, ctx)
	case *pb.CLHFreeDealer:
		arg := msg.(*pb.CLHFreeDealer)
		glog.Debugf("CLHFreeDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SLHFreeDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CLHFreeDealerList:
		arg := msg.(*pb.CLHFreeDealerList)
		glog.Debugf("CLHFreeDealerList %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SLHFreeDealerList)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CLHSit:
		arg := msg.(*pb.CLHSit)
		glog.Debugf("CLHSit %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SLHSit)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CLHFreeBet:
		arg := msg.(*pb.CLHFreeBet)
		glog.Debugf("CLHFreeBet %#v", arg)
		rs.lhdFreeBet(arg, ctx)
	case *pb.CLHFreeTrend:
		arg := msg.(*pb.CLHFreeTrend)
		glog.Debugf("CLHFreeTrend %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SLHFreeTrend)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CLHFreeWiners:
		arg := msg.(*pb.CLHFreeWiners)
		glog.Debugf("CLHFreeWiners %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SLHFreeWiners)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CLHFreeRoles:
		arg := msg.(*pb.CLHFreeRoles)
		glog.Debugf("CLHFreeRoles %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SLHFreeRoles)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CLHRoomList:
		arg := msg.(*pb.CLHRoomList)
		glog.Debugf("CLHRoomList %#v", arg)
		rs.getLhdRoomList(arg, ctx)
	case *pb.CLHLeave:
		arg := msg.(*pb.CLHLeave)
		glog.Debugf("CLHLeave %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SLHLeave)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerHua(msg, ctx)
	}
}

//进入百人房间
func (rs *RoleActor) enterLHFree(arg *pb.CLHFreeEnterRoom, ctx actor.Context) {
	msg := rs.enterLhdMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE2) //百人
		rs.selectDesk(msg, ctx)
	}
}

//进入或匹配桌子
func (rs *RoleActor) enterLhdMatchDesk(ctx actor.Context) *pb.MatchDesk {
	//已经在游戏中,直接加入
	if rs.enterGame(ctx) {
		return nil
	}
	//获取游戏服节点或者房间进程
	msg := new(pb.MatchDesk)
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.lhd").Name()
	msg.Gtype = int32(pb.LHD) //lhd
	return msg
}

//进入或匹配桌子
func (rs *RoleActor) getLhdRoomList(arg *pb.CLHRoomList, ctx actor.Context) {
	//获取游戏服节点或者房间进程
	msg := new(pb.GetRoomList)
	msg.Rtype = arg.Rtype
	msg.Userid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.lhd").Name()
	msg.Gtype = int32(pb.LHD) //lhd
	rs.dbmsPid.Request(msg, ctx.Self())
}

//百人场下注
func (rs *RoleActor) lhdFreeBet(arg *pb.CLHFreeBet, ctx actor.Context) {
	if rs.User.IsTourist() {
		rsp := new(pb.SLHFreeBet)
		rsp.Error = pb.TouristInoperable
		rs.Send(rsp)
		return
	}
	if rs.gamePid == nil {
		rsp := new(pb.SLHFreeBet)
		rsp.Error = pb.NotInRoom
		rs.Send(rsp)
		return
	}
	value := arg.GetValue()
	seat := arg.GetSeat()
	if !(seat >= uint32(pb.DESK_SEAT2) &&
		seat <= uint32(pb.DESK_SEAT4)) {
		rsp := new(pb.SLHFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if value <= 0 {
		rsp := new(pb.SLHFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if rs.User.GetCoin() < int64(value) {
		rsp := new(pb.SLHFreeBet)
		rsp.Error = pb.NotEnoughCoin
		rs.Send(rsp)
		return
	}
	rs.gamePid.Request(arg, ctx.Self())
}
