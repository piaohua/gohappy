package main

import (
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//牛牛请求处理
func (rs *RoleActor) handlerNiu(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CChatText:
		arg := msg.(*pb.CChatText)
		glog.Debugf("CChatText %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SChatText)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CChatVoice:
		arg := msg.(*pb.CChatVoice)
		glog.Debugf("CChatVoice %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SChatVoice)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNCoinEnterRoom:
		arg := msg.(*pb.CNNCoinEnterRoom)
		glog.Debugf("CNNCoinEnterRoom %#v", arg)
		rs.enterNNCoin(arg, ctx)
	case *pb.CNNFreeEnterRoom:
		arg := msg.(*pb.CNNFreeEnterRoom)
		glog.Debugf("CNNFreeEnterRoom %#v", arg)
		rs.enterNNFree(arg, ctx)
	case *pb.CNNFreeDealer:
		arg := msg.(*pb.CNNFreeDealer)
		glog.Debugf("CNNFreeDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNFreeDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNFreeDealerList:
		arg := msg.(*pb.CNNFreeDealerList)
		glog.Debugf("CNNFreeDealerList %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNFreeDealerList)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNSit:
		arg := msg.(*pb.CNNSit)
		glog.Debugf("CNNSit %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNSit)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNFreeBet:
		arg := msg.(*pb.CNNFreeBet)
		glog.Debugf("CNNFreeBet %#v", arg)
		rs.nnFreeBet(arg, ctx)
	case *pb.CNNFreeTrend:
		arg := msg.(*pb.CNNFreeTrend)
		glog.Debugf("CNNFreeTrend %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNFreeTrend)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNFreeWiners:
		arg := msg.(*pb.CNNFreeWiners)
		glog.Debugf("CNNFreeWiners %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNFreeWiners)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNFreeRoles:
		arg := msg.(*pb.CNNFreeRoles)
		glog.Debugf("CNNFreeRoles %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNFreeRoles)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNRoomList:
		arg := msg.(*pb.CNNRoomList)
		glog.Debugf("CNNRoomList %#v", arg)
		rs.getRoomList(arg, ctx)
	case *pb.CNNEnterRoom:
		arg := msg.(*pb.CNNEnterRoom)
		glog.Debugf("CNNEnterRoom %#v", arg)
		rs.enterNNPriv(arg, ctx)
	case *pb.CNNCreateRoom:
		arg := msg.(*pb.CNNCreateRoom)
		glog.Debugf("CNNCreateRoom %#v", arg)
		rs.createRoom(arg, ctx)
	case *pb.CNNLeave:
		arg := msg.(*pb.CNNLeave)
		glog.Debugf("CNNLeave %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNLeave)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNReady:
		arg := msg.(*pb.CNNReady)
		glog.Debugf("CNNReady %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNReady)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNDealer:
		arg := msg.(*pb.CNNDealer)
		glog.Debugf("CNNDealer %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNDealer)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNBet:
		arg := msg.(*pb.CNNBet)
		glog.Debugf("CNNBet %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNBet)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNiu:
		arg := msg.(*pb.CNNiu)
		glog.Debugf("CNNiu %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNiu)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNGameRecord:
		arg := msg.(*pb.CNNGameRecord)
		glog.Debugf("CNNGameRecord %#v", arg)
		//TODO
	case *pb.CNNLaunchVote:
		arg := msg.(*pb.CNNLaunchVote)
		glog.Debugf("CNNLaunchVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNLaunchVote)
			rsp.Error = pb.NotInRoom
			rs.Send(rsp)
			return
		}
		rs.gamePid.Request(arg, ctx.Self())
	case *pb.CNNVote:
		arg := msg.(*pb.CNNVote)
		glog.Debugf("CNNVote %#v", arg)
		if rs.gamePid == nil {
			rsp := new(pb.SNNVote)
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
func (rs *RoleActor) enterNNFree(arg *pb.CNNFreeEnterRoom, ctx actor.Context) {
	msg := rs.enterMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE2) //百人
		rs.selectDesk(msg, ctx)
	}
}

//进入私人房间
func (rs *RoleActor) enterNNPriv(arg *pb.CNNEnterRoom, ctx actor.Context) {
	msg := rs.enterMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE1) //私人
		msg.Code = arg.Code              //邀请码
		rs.selectDesk(msg, ctx)
	}
}

//进入自由房间
func (rs *RoleActor) enterNNCoin(arg *pb.CNNCoinEnterRoom, ctx actor.Context) {
	msg := rs.enterMatchDesk(ctx)
	if msg != nil {
		msg.Rtype = int32(pb.ROOM_TYPE0) //自由
		msg.Roomid = arg.Id              //房间id
		msg.Dtype = int32(pb.DESK_TYPE0) //玩法类型
		//计算出匹配房间等级,算法一致
		//msg.Ltype = int32(pb.ROOM_LEVEL1) //等级
		msg.Ltype = handler.MatchLevel(rs.User.GetCoin())
		rs.selectDesk(msg, ctx)
	}
}

//进入或匹配桌子
func (rs *RoleActor) enterMatchDesk(ctx actor.Context) *pb.MatchDesk {
	//已经在游戏中,直接加入
	if rs.enterGame(ctx) {
		return nil
	}
	//获取游戏服节点或者房间进程
	msg := new(pb.MatchDesk)
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.niu").Name()
	msg.Gtype = int32(pb.NIU) //牛牛
	return msg
}

//进入或匹配桌子
func (rs *RoleActor) getRoomList(arg *pb.CNNRoomList, ctx actor.Context) {
	//获取游戏服节点或者房间进程
	msg := new(pb.GetRoomList)
	msg.Rtype = arg.Rtype
	msg.Userid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//TODO 优化查找规则
	msg.Name = cfg.Section("game.niu").Name()
	msg.Gtype = int32(pb.NIU) //牛牛
	rs.dbmsPid.Request(msg, ctx.Self())
}

//百人场下注
func (rs *RoleActor) nnFreeBet(arg *pb.CNNFreeBet, ctx actor.Context) {
	if rs.User.IsTourist() {
		rsp := new(pb.SNNFreeBet)
		rsp.Error = pb.TouristInoperable
		rs.Send(rsp)
		return
	}
	if rs.gamePid == nil {
		rsp := new(pb.SNNFreeBet)
		rsp.Error = pb.NotInRoom
		rs.Send(rsp)
		return
	}
	value := arg.GetValue()
	seat := arg.GetSeat()
	if !(seat >= uint32(pb.DESK_SEAT2) &&
		seat <= uint32(pb.DESK_SEAT9)) {
		rsp := new(pb.SNNFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if value <= 0 {
		rsp := new(pb.SNNFreeBet)
		rsp.Error = pb.OperateError
		rs.Send(rsp)
		return
	}
	if rs.User.GetCoin() < int64(value) {
		rsp := new(pb.SNNFreeBet)
		rsp.Error = pb.NotEnoughCoin
		rs.Send(rsp)
		return
	}
	rs.gamePid.Request(arg, ctx.Self())
}

//创建房间
func (rs *RoleActor) createRoom(arg *pb.CNNCreateRoom, ctx actor.Context) {
	//TODO 验证
	msg := &pb.CreateDesk{
		Rname:   arg.Rname,
		Dtype:   arg.Dtype,
		Ante:    arg.Ante,
		Round:   arg.Round,
		Payment: arg.Payment,
		Count:   arg.Count,
		Pub:     arg.Pub,
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
	msg.Name = cfg.Section("game.niu").Name()
	msg.Gtype = int32(pb.NIU)        //牛牛
	msg.Rtype = int32(pb.ROOM_TYPE1) //私人
	msg.Cid = rs.User.GetUserid()
	msg.Sender = ctx.Self()
	//节点中匹配
	rs.dbmsPid.Request(msg, ctx.Self())
}
