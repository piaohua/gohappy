package main

import (
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//hua请求处理
func (rs *RoleActor) handlerHua(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CJHCoinEnterRoom:
		arg := msg.(*pb.CJHCoinEnterRoom)
		glog.Debugf("CJHCoinEnterRoom %#v", arg)
		rs.enterJHCoin(arg, ctx)
	case *pb.CJHFreeEnterRoom:
		arg := msg.(*pb.CJHFreeEnterRoom)
		glog.Debugf("CJHFreeEnterRoom %#v", arg)
		rs.enterJHFree(arg, ctx)
	case *pb.CJHFreeDealer:
		arg := msg.(*pb.CJHFreeDealer)
		glog.Debugf("CJHFreeDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHFreeDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHFreeDealerList:
		arg := msg.(*pb.CJHFreeDealerList)
		glog.Debugf("CJHFreeDealerList %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHFreeDealerList)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHSit:
		arg := msg.(*pb.CJHSit)
		glog.Debugf("CJHSit %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHSit)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHFreeBet:
		arg := msg.(*pb.CJHFreeBet)
		glog.Debugf("CJHFreeBet %#v", arg)
		rs.nnJHFreeBet(arg, ctx)
	case *pb.CJHFreeTrend:
		arg := msg.(*pb.CJHFreeTrend)
		glog.Debugf("CJHFreeTrend %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHFreeTrend)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHFreeWiners:
		arg := msg.(*pb.CJHFreeWiners)
		glog.Debugf("CJHFreeWiners %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHFreeWiners)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHFreeRoles:
		arg := msg.(*pb.CJHFreeRoles)
		glog.Debugf("CJHFreeRoles %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHFreeRoles)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHRoomList:
		arg := msg.(*pb.CJHRoomList)
		glog.Debugf("CJHRoomList %#v", arg)
		rs.getJHRoomList(arg, ctx)
	case *pb.CJHEnterRoom:
		arg := msg.(*pb.CJHEnterRoom)
		glog.Debugf("CJHEnterRoom %#v", arg)
		rs.enterJHPriv(arg, ctx)
	case *pb.CJHCreateRoom:
		arg := msg.(*pb.CJHCreateRoom)
		glog.Debugf("CJHCreateRoom %#v", arg)
		rs.createJHRoom(arg, ctx)
	case *pb.CJHLeave:
		arg := msg.(*pb.CJHLeave)
		glog.Debugf("CJHLeave %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHLeave)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHReady:
		arg := msg.(*pb.CJHReady)
		glog.Debugf("CJHReady %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHReady)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHGameRecord:
		arg := msg.(*pb.CJHGameRecord)
		glog.Debugf("CJHGameRecord %#v", arg)
		//TODO
	case *pb.CJHLaunchVote:
		arg := msg.(*pb.CJHLaunchVote)
		glog.Debugf("CJHLaunchVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHLaunchVote)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHVote:
		arg := msg.(*pb.CJHVote)
		glog.Debugf("CJHVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHVote)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHCoinSee:
		arg := msg.(*pb.CJHCoinSee)
		glog.Debugf("CJHCoinSee %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHCoinSee)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHCoinCall:
		arg := msg.(*pb.CJHCoinCall)
		glog.Debugf("CJHCoinCall %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHCoinCall)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHCoinRaise:
		arg := msg.(*pb.CJHCoinRaise)
		glog.Debugf("CJHCoinRaise %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHCoinRaise)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHCoinFold:
		arg := msg.(*pb.CJHCoinFold)
		glog.Debugf("CJHCoinFold %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHCoinFold)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CJHCoinBi:
		arg := msg.(*pb.CJHCoinBi)
		glog.Debugf("CJHCoinBi %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SJHCoinBi)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case proto.Message:
		//响应消息
		rs.Send(msg)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//进入百人房间
func (rs *RoleActor) enterJHFree(arg *pb.CJHFreeEnterRoom, ctx actor.Context) {
	msg := rs.enterJHMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE2) //百人
		rs.selectDesk(msg, ctx)
	}
}

//进入私人房间
func (rs *RoleActor) enterJHPriv(arg *pb.CJHEnterRoom, ctx actor.Context) {
	msg := rs.enterJHMatchDesk(ctx)
	if msg != nil {
		//msg.Rtype = int32(pb.ROOM_TYPE1) //私人
		msg.Rtype = int32(pb.ROOM_TYPE0) //自由
		msg.Code = arg.Code              //邀请码
		rs.selectDesk(msg, ctx)
	}
}

//进入自由房间
func (rs *RoleActor) enterJHCoin(arg *pb.CJHCoinEnterRoom, ctx actor.Context) {
	msg := rs.enterJHMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE0) //自由
		msg.Roomid = arg.Id              //房间id
		msg.Dtype = int32(pb.DESK_TYPE1) //玩法类型
		//计算出匹配房间等级,算法一致
		//msg.Ltype = int32(pb.ROOM_LEVEL1) //等级
		msg.Ltype = handler.MatchLevel(rs.User.GetCoin())
		rs.selectDesk(msg, ctx)
	}
}

//进入或匹配桌子
func (rs *RoleActor) enterJHMatchDesk(ctx actor.Context) *pb.MatchDesk {
	//已经在游戏中,直接加入
	if rs.enterGame(ctx) {
		return nil
	}
	//获取游戏服节点或者房间进程
	msg := new(pb.MatchDesk)
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.hua").Name()
	msg.Gtype = int32(pb.HUA) //金花
	return msg
}

//进入或匹配桌子
func (rs *RoleActor) getJHRoomList(arg *pb.CJHRoomList, ctx actor.Context) {
	//获取游戏服节点或者房间进程
	msg := new(pb.GetRoomList)
	msg.Rtype = arg.Rtype
	msg.Userid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.hua").Name()
	msg.Gtype = int32(pb.HUA) //金花
	rs.dbmsPid.Request(msg, ctx.Self())
}

//百人场下注
func (rs *RoleActor) nnJHFreeBet(arg *pb.CJHFreeBet, ctx actor.Context) {
	if rs.User.IsTourist() {
		rsp := new(pb.SJHFreeBet)
		rsp.Error = pb.TouristInoperable
		rs.Send(rsp)
		return
	}
	if rs.gamePid == nil {
		rsp := new(pb.SJHFreeBet)
		rsp.Error = pb.NotInRoom
		rs.Send(rsp)
		return
	}
	value := arg.GetValue()
	seat := arg.GetSeat()
	if !(seat >= uint32(pb.DESK_SEAT2) &&
		seat <= uint32(pb.DESK_SEAT9)) {
		rsp := new(pb.SJHFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if value <= 0 {
		rsp := new(pb.SJHFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if rs.User.GetCoin() < int64(value) {
		rsp := new(pb.SJHFreeBet)
		rsp.Error = pb.NotEnoughCoin
		rs.Send(rsp)
		return
	}
	rs.gamePid.Request(arg, ctx.Self())
}

//创建房间
func (rs *RoleActor) createJHRoom(arg *pb.CJHCreateRoom, ctx actor.Context) {
	//TODO 验证
	msg := &pb.CreateDesk{
		Rname:   arg.Rname,
		Dtype:   arg.Dtype,
		Ante:    arg.Ante,
		Round:   arg.Round,
		Payment: arg.Payment,
		Count:   arg.Count,
		//TODO 消耗
		Cost: 100,
	}
	switch msg.Dtype {
	case int32(pb.DESK_TYPE0):
	case int32(pb.DESK_TYPE1):
	case int32(pb.DESK_TYPE2):
	default:
		msg.Dtype = int32(pb.DESK_TYPE0)
	}
	msg.Name = cfg.Section("game.hua").Name()
	msg.Gtype = int32(pb.HUA)        //金花
	msg.Rtype = int32(pb.ROOM_TYPE1) //私人
	msg.Cid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//节点中匹配
	rs.dbmsPid.Request(msg, ctx.Self())
}
