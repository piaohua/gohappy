package main

import (
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//ebg请求处理
func (rs *RoleActor) handlerEbg(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CEBCoinEnterRoom:
		arg := msg.(*pb.CEBCoinEnterRoom)
		glog.Debugf("CEBCoinEnterRoom %#v", arg)
		rs.enterEbgCoin(arg, ctx)
	case *pb.CEBFreeEnterRoom:
		arg := msg.(*pb.CEBFreeEnterRoom)
		glog.Debugf("CEBFreeEnterRoom %#v", arg)
		rs.enterEbgFree(arg, ctx)
	case *pb.CEBFreeDealer:
		arg := msg.(*pb.CEBFreeDealer)
		glog.Debugf("CEBFreeDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBFreeDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBFreeDealerList:
		arg := msg.(*pb.CEBFreeDealerList)
		glog.Debugf("CEBFreeDealerList %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBFreeDealerList)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBSit:
		arg := msg.(*pb.CEBSit)
		glog.Debugf("CEBSit %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBSit)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBFreeBet:
		arg := msg.(*pb.CEBFreeBet)
		glog.Debugf("CEBFreeBet %#v", arg)
		rs.ebgFreeBet(arg, ctx)
	case *pb.CEBFreeTrend:
		arg := msg.(*pb.CEBFreeTrend)
		glog.Debugf("CEBFreeTrend %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBFreeTrend)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBFreeWiners:
		arg := msg.(*pb.CEBFreeWiners)
		glog.Debugf("CEBFreeWiners %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBFreeWiners)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBFreeRoles:
		arg := msg.(*pb.CEBFreeRoles)
		glog.Debugf("CEBFreeRoles %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBFreeRoles)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBRoomList:
		arg := msg.(*pb.CEBRoomList)
		glog.Debugf("CEBRoomList %#v", arg)
		rs.getEbgRoomList(arg, ctx)
	case *pb.CEBEnterRoom:
		arg := msg.(*pb.CEBEnterRoom)
		glog.Debugf("CEBEnterRoom %#v", arg)
		rs.enterEbgPriv(arg, ctx)
	case *pb.CEBCreateRoom:
		arg := msg.(*pb.CEBCreateRoom)
		glog.Debugf("CEBCreateRoom %#v", arg)
		rs.createEbgRoom(arg, ctx)
	case *pb.CEBLeave:
		arg := msg.(*pb.CEBLeave)
		glog.Debugf("CEBLeave %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBLeave)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBReady:
		arg := msg.(*pb.CEBReady)
		glog.Debugf("CEBReady %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBReady)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBDealer:
		arg := msg.(*pb.CEBDealer)
		glog.Debugf("CEBDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBBet:
		arg := msg.(*pb.CEBBet)
		glog.Debugf("CEBBet %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBBet)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBiu:
		arg := msg.(*pb.CEBiu)
		glog.Debugf("CEBiu %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBiu)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBGameRecord:
		arg := msg.(*pb.CEBGameRecord)
		glog.Debugf("CEBGameRecord %#v", arg)
		//TODO
	case *pb.CEBLaunchVote:
		arg := msg.(*pb.CEBLaunchVote)
		glog.Debugf("CEBLaunchVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBLaunchVote)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBVote:
		arg := msg.(*pb.CEBVote)
		glog.Debugf("CEBVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBVote)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CEBCoinChangeRoom:
		arg := msg.(*pb.CEBCoinChangeRoom)
		glog.Debugf("CEBCoinChangeRoom %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SEBCoinChangeRoom)
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
		rs.handlerSan(msg, ctx)
	}
}

//进入百人房间
func (rs *RoleActor) enterEbgFree(arg *pb.CEBFreeEnterRoom, ctx actor.Context) {
	msg := rs.enterEbgMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE2) //百人
		rs.selectDesk(msg, ctx)
	}
}

//进入私人房间
func (rs *RoleActor) enterEbgPriv(arg *pb.CEBEnterRoom, ctx actor.Context) {
	msg := rs.enterEbgMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE1) //私人
		msg.Code = arg.Code              //邀请码
		rs.selectDesk(msg, ctx)
	}
}

//进入自由房间
func (rs *RoleActor) enterEbgCoin(arg *pb.CEBCoinEnterRoom, ctx actor.Context) {
	msg := rs.enterEbgMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE0) //自由
		msg.Roomid = arg.Id              //房间id
		msg.Dtype = int32(pb.DESK_TYPE0) //玩法类型
		//计算出匹配房间等级,算法一致
		//msg.Ltype = int32(pb.ROOM_LEVEL1) //等级
		msg.Ltype = handler.MatchLevel(rs.User.GetCoin())
		if msg.Ltype < 0 {
			rsp := new(pb.SEBCoinEnterRoom)
			rsp.Error = pb.NotEnoughCoin
			rs.Send(rsp)
			return
		}
		rs.selectDesk(msg, ctx)
	}
}

//进入或匹配桌子
func (rs *RoleActor) enterEbgMatchDesk(ctx actor.Context) *pb.MatchDesk {
	//已经在游戏中,直接加入
	if rs.enterGame(ctx) {
		return nil
	}
	//获取游戏服节点或者房间进程
	msg := new(pb.MatchDesk)
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.ebg").Name()
	msg.Gtype = int32(pb.EBG) //ebg
	return msg
}

//进入或匹配桌子
func (rs *RoleActor) getEbgRoomList(arg *pb.CEBRoomList, ctx actor.Context) {
	//获取游戏服节点或者房间进程
	msg := new(pb.GetRoomList)
	msg.Rtype = arg.Rtype
	msg.Userid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.ebg").Name()
	msg.Gtype = int32(pb.EBG) //ebg
	rs.dbmsPid.Request(msg, ctx.Self())
}

//百人场下注
func (rs *RoleActor) ebgFreeBet(arg *pb.CEBFreeBet, ctx actor.Context) {
	if rs.User.IsTourist() {
		rsp := new(pb.SEBFreeBet)
		rsp.Error = pb.TouristInoperable
		rs.Send(rsp)
		return
	}
	if rs.gamePid == nil {
		rsp := new(pb.SEBFreeBet)
		rsp.Error = pb.NotInRoom
		rs.Send(rsp)
		return
	}
	value := arg.GetValue()
	seat := arg.GetSeat()
	if !(seat >= uint32(pb.DESK_SEAT2) &&
		seat <= uint32(pb.DESK_SEAT9)) {
		rsp := new(pb.SEBFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if value <= 0 {
		rsp := new(pb.SEBFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if rs.User.GetCoin() < int64(value) {
		rsp := new(pb.SEBFreeBet)
		rsp.Error = pb.NotEnoughCoin
		rs.Send(rsp)
		return
	}
	rs.gamePid.Request(arg, ctx.Self())
}

//创建房间
func (rs *RoleActor) createEbgRoom(arg *pb.CEBCreateRoom, ctx actor.Context) {
	//TODO 验证
	msg := &pb.CreateDesk{
		Rname:    arg.Rname,
		Dtype:    arg.Dtype,
		Ante:     arg.Ante,
		Round:    arg.Round,
		Payment:  arg.Payment,
		Count:    arg.Count,
		Pub:      arg.Pub,
		Minimum:  int64(arg.Minimum),
		Maximum:  int64(arg.Maximum),
		Mode:     arg.Mode,
		Multiple: arg.Multiple,
		//TODO 消耗
		Cost: 100,
	}
	switch msg.Dtype {
	case int32(pb.DESK_TYPE0):
	case int32(pb.DESK_TYPE1):
	case int32(pb.DESK_TYPE2):
	case int32(pb.DESK_TYPE3):
	default:
		msg.Dtype = int32(pb.DESK_TYPE0)
	}
	msg.Name = cfg.Section("game.ebg").Name()
	msg.Gtype = int32(pb.EBG)        //ebg
	msg.Rtype = int32(pb.ROOM_TYPE1) //私人
	msg.Cid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//节点中匹配
	rs.dbmsPid.Request(msg, ctx.Self())
}
