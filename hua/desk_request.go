package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家请求
func (a *Desk) handlerRequest(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CChatText:
		arg := msg.(*pb.CChatText)
		glog.Debugf("CChatText %#v", arg)
		a.chatText(arg, ctx)
	case *pb.CChatVoice:
		arg := msg.(*pb.CChatVoice)
		glog.Debugf("CChatVoice %#v", arg)
		a.chatVoice(arg, ctx)
	case *pb.CJHCoinEnterRoom:
		arg := msg.(*pb.CJHCoinEnterRoom)
		glog.Debugf("CJHCoinEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.coinEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CJHFreeEnterRoom:
		arg := msg.(*pb.CJHFreeEnterRoom)
		glog.Debugf("CJHFreeEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.freeEnterMsg(userid)
		ctx.Respond(msg1)
		a.freeCameinMsg(userid)
	case *pb.CJHFreeDealer:
		arg := msg.(*pb.CJHFreeDealer)
		glog.Debugf("CJHFreeDealer %#v", arg)
		userid := a.getRouter(ctx)
		var state int32 = arg.GetState()
		var num uint32 = arg.GetCoin()
		errcode := a.beDealer(userid, state, num)
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SJHFreeDealer)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CJHFreeDealerList:
		arg := msg.(*pb.CJHFreeDealerList)
		glog.Debugf("CJHFreeDealerList %#v", arg)
		rsp := a.dealerListMsg()
		ctx.Respond(rsp)
	case *pb.CJHSit:
		arg := msg.(*pb.CJHSit)
		glog.Debugf("CJHSit %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.freeSit(userid, arg)
		if rsp.Error == pb.OK {
			a.broadcast(rsp)
			return
		}
		ctx.Respond(rsp)
	case *pb.CJHFreeBet:
		arg := msg.(*pb.CJHFreeBet)
		glog.Debugf("CJHFreeBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeat()
		var val uint32 = arg.GetValue()
		errcode := a.freeBet(userid, seatBet, int64(val))
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SJHFreeBet)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CJHFreeTrend:
		arg := msg.(*pb.CJHFreeTrend)
		glog.Debugf("CJHFreeTrend %#v", arg)
		rsp := a.freeTrends()
		ctx.Respond(rsp)
	case *pb.CJHFreeWiners:
		arg := msg.(*pb.CJHFreeWiners)
		glog.Debugf("CJHFreeWiners %#v", arg)
		rsp := a.freeWiners()
		ctx.Respond(rsp)
	case *pb.CJHFreeRoles:
		arg := msg.(*pb.CJHFreeRoles)
		glog.Debugf("CJHFreeRoles %#v", arg)
		rsp := a.freeRoles()
		ctx.Respond(rsp)
	//case *pb.CJHRoomList:
	//	arg := msg.(*pb.CJHRoomList)
	//	glog.Debugf("CJHRoomList %#v", arg)
	//	//TODO
	//case *pb.CJHCreateRoom:
	//	arg := msg.(*pb.CJHCreateRoom)
	//	glog.Debugf("CJHCreateRoom %#v", arg)
	//	//TODO
	case *pb.CJHEnterRoom:
		arg := msg.(*pb.CJHEnterRoom)
		glog.Debugf("CJHEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.privEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CJHLeave:
		arg := msg.(*pb.CJHLeave)
		glog.Debugf("CJHLeave %#v", arg)
		userid := a.getRouter(ctx)
		a.nnLeave(userid, ctx)
	case *pb.CJHReady:
		arg := msg.(*pb.CJHReady)
		glog.Debugf("CJHReady %#v", arg)
		userid := a.getRouter(ctx)
		var ready bool = arg.GetReady()
		rsp := a.readying(userid, ready)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CJHGameRecord:
		arg := msg.(*pb.CJHGameRecord)
		glog.Debugf("CJHGameRecord %#v", arg)
		//TODO
	case *pb.CJHLaunchVote:
		arg := msg.(*pb.CJHLaunchVote)
		glog.Debugf("CJHLaunchVote %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.launchVote(userid, 1)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CJHVote:
		arg := msg.(*pb.CJHVote)
		glog.Debugf("CJHVote %#v", arg)
		userid := a.getRouter(ctx)
		var vote uint32 = arg.GetVote()
		rsp := a.privVote(userid, vote)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CJHCoinSee:
		arg := msg.(*pb.CJHCoinSee)
		glog.Debugf("CJHCoinSee %#v", arg)
		userid := a.getRouter(ctx)
		a.coinSee(userid)
	case *pb.CJHCoinCall:
		arg := msg.(*pb.CJHCoinCall)
		glog.Debugf("CJHCoinCall %#v", arg)
		userid := a.getRouter(ctx)
		a.coinCall(userid, int64(arg.Value))
	case *pb.CJHCoinRaise:
		arg := msg.(*pb.CJHCoinRaise)
		glog.Debugf("CJHCoinRaise %#v", arg)
		userid := a.getRouter(ctx)
		a.coinRaise(userid, int64(arg.Value))
	case *pb.CJHCoinFold:
		arg := msg.(*pb.CJHCoinFold)
		glog.Debugf("CJHCoinFold %#v", arg)
		userid := a.getRouter(ctx)
		a.coinFold(userid)
	case *pb.CJHCoinBi:
		arg := msg.(*pb.CJHCoinBi)
		glog.Debugf("CJHCoinBi %#v", arg)
		userid := a.getRouter(ctx)
		a.coinBi(userid, arg.Seat)
	case *pb.CJHCoinChangeRoom:
		arg := msg.(*pb.CJHCoinChangeRoom)
		glog.Debugf("CJHCoinChangeRoom %#v", arg)
		a.changeDesk(ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}
