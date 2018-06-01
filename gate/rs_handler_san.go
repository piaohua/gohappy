package main

import (
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//san请求处理
func (rs *RoleActor) handlerSan(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CSGCoinEnterRoom:
		arg := msg.(*pb.CSGCoinEnterRoom)
		glog.Debugf("CSGCoinEnterRoom %#v", arg)
		rs.enterSGCoin(arg, ctx)
	case *pb.CSGFreeEnterRoom:
		arg := msg.(*pb.CSGFreeEnterRoom)
		glog.Debugf("CSGFreeEnterRoom %#v", arg)
		rs.enterSGFree(arg, ctx)
	case *pb.CSGFreeDealer:
		arg := msg.(*pb.CSGFreeDealer)
		glog.Debugf("CSGFreeDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGFreeDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGFreeDealerList:
		arg := msg.(*pb.CSGFreeDealerList)
		glog.Debugf("CSGFreeDealerList %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGFreeDealerList)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGSit:
		arg := msg.(*pb.CSGSit)
		glog.Debugf("CSGSit %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGSit)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGFreeBet:
		arg := msg.(*pb.CSGFreeBet)
		glog.Debugf("CSGFreeBet %#v", arg)
		rs.nnSGFreeBet(arg, ctx)
	case *pb.CSGFreeTrend:
		arg := msg.(*pb.CSGFreeTrend)
		glog.Debugf("CSGFreeTrend %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGFreeTrend)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGFreeWiners:
		arg := msg.(*pb.CSGFreeWiners)
		glog.Debugf("CSGFreeWiners %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGFreeWiners)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGFreeRoles:
		arg := msg.(*pb.CSGFreeRoles)
		glog.Debugf("CSGFreeRoles %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGFreeRoles)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGRoomList:
		arg := msg.(*pb.CSGRoomList)
		glog.Debugf("CSGRoomList %#v", arg)
		rs.getSGRoomList(arg, ctx)
	case *pb.CSGEnterRoom:
		arg := msg.(*pb.CSGEnterRoom)
		glog.Debugf("CSGEnterRoom %#v", arg)
		rs.enterSGPriv(arg, ctx)
	case *pb.CSGCreateRoom:
		arg := msg.(*pb.CSGCreateRoom)
		glog.Debugf("CSGCreateRoom %#v", arg)
		rs.createSGRoom(arg, ctx)
	case *pb.CSGLeave:
		arg := msg.(*pb.CSGLeave)
		glog.Debugf("CSGLeave %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGLeave)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGReady:
		arg := msg.(*pb.CSGReady)
		glog.Debugf("CSGReady %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGReady)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGDealer:
		arg := msg.(*pb.CSGDealer)
		glog.Debugf("CSGDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGBet:
		arg := msg.(*pb.CSGBet)
		glog.Debugf("CSGBet %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGBet)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGiu:
		arg := msg.(*pb.CSGiu)
		glog.Debugf("CSGiu %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGiu)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGGameRecord:
		arg := msg.(*pb.CSGGameRecord)
		glog.Debugf("CSGGameRecord %#v", arg)
		//TODO
	case *pb.CSGLaunchVote:
		arg := msg.(*pb.CSGLaunchVote)
		glog.Debugf("CSGLaunchVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGLaunchVote)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CSGVote:
		arg := msg.(*pb.CSGVote)
		glog.Debugf("CSGVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SSGVote)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	//case proto.Message:
	//	//响应消息
	//	rs.Send(msg)
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerHua(msg, ctx)
	}
}

//进入百人房间
func (rs *RoleActor) enterSGFree(arg *pb.CSGFreeEnterRoom, ctx actor.Context) {
	msg := rs.enterSGMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE2) //百人
		rs.selectDesk(msg, ctx)
	}
}

//进入私人房间
func (rs *RoleActor) enterSGPriv(arg *pb.CSGEnterRoom, ctx actor.Context) {
	msg := rs.enterSGMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE1) //私人
		msg.Code = arg.Code              //邀请码
		rs.selectDesk(msg, ctx)
	}
}

//进入自由房间
func (rs *RoleActor) enterSGCoin(arg *pb.CSGCoinEnterRoom, ctx actor.Context) {
	msg := rs.enterSGMatchDesk(ctx)
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
func (rs *RoleActor) enterSGMatchDesk(ctx actor.Context) *pb.MatchDesk {
	//已经在游戏中,直接加入
	if rs.enterGame(ctx) {
		return nil
	}
	//获取游戏服节点或者房间进程
	msg := new(pb.MatchDesk)
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.san").Name()
	msg.Gtype = int32(pb.SAN) //三公
	return msg
}

//进入或匹配桌子
func (rs *RoleActor) getSGRoomList(arg *pb.CSGRoomList, ctx actor.Context) {
	//获取游戏服节点或者房间进程
	msg := new(pb.GetRoomList)
	msg.Rtype = arg.Rtype
	msg.Userid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.san").Name()
	msg.Gtype = int32(pb.SAN) //三公
	rs.dbmsPid.Request(msg, ctx.Self())
}

//百人场下注
func (rs *RoleActor) nnSGFreeBet(arg *pb.CSGFreeBet, ctx actor.Context) {
	if rs.User.IsTourist() {
		rsp := new(pb.SSGFreeBet)
		rsp.Error = pb.TouristInoperable
		rs.Send(rsp)
		return
	}
	if rs.gamePid == nil {
		rsp := new(pb.SSGFreeBet)
		rsp.Error = pb.NotInRoom
		rs.Send(rsp)
		return
	}
	value := arg.GetValue()
	seat := arg.GetSeat()
	if !(seat >= uint32(pb.DESK_SEAT2) &&
		seat <= uint32(pb.DESK_SEAT9)) {
		rsp := new(pb.SSGFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if value <= 0 {
		rsp := new(pb.SSGFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if rs.User.GetCoin() < int64(value) {
		rsp := new(pb.SSGFreeBet)
		rsp.Error = pb.NotEnoughCoin
		rs.Send(rsp)
		return
	}
	rs.gamePid.Request(arg, ctx.Self())
}

//创建房间
func (rs *RoleActor) createSGRoom(arg *pb.CSGCreateRoom, ctx actor.Context) {
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
	msg.Name = cfg.Section("game.san").Name()
	msg.Gtype = int32(pb.SAN)        //三公
	msg.Rtype = int32(pb.ROOM_TYPE1) //私人
	msg.Cid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//节点中匹配
	rs.dbmsPid.Request(msg, ctx.Self())
}
