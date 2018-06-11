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
	case *pb.CSGCoinEnterRoom:
		arg := msg.(*pb.CSGCoinEnterRoom)
		glog.Debugf("CSGCoinEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.coinEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CSGFreeEnterRoom:
		arg := msg.(*pb.CSGFreeEnterRoom)
		glog.Debugf("CSGFreeEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.freeEnterMsg(userid)
		ctx.Respond(msg1)
		a.freeCameinMsg(userid)
	case *pb.CSGFreeDealer:
		arg := msg.(*pb.CSGFreeDealer)
		glog.Debugf("CSGFreeDealer %#v", arg)
		userid := a.getRouter(ctx)
		var state int32 = arg.GetState()
		var num uint32 = arg.GetCoin()
		errcode := a.beDealer(userid, state, num)
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SSGFreeDealer)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CSGFreeDealerList:
		arg := msg.(*pb.CSGFreeDealerList)
		glog.Debugf("CSGFreeDealerList %#v", arg)
		rsp := a.dealerListMsg()
		ctx.Respond(rsp)
	case *pb.CSGSit:
		arg := msg.(*pb.CSGSit)
		glog.Debugf("CSGSit %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.freeSit(userid, arg)
		if rsp.Error == pb.OK {
			a.broadcast(rsp)
			return
		}
		ctx.Respond(rsp)
	case *pb.CSGFreeBet:
		arg := msg.(*pb.CSGFreeBet)
		glog.Debugf("CSGFreeBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeat()
		var val uint32 = arg.GetValue()
		errcode := a.freeBet(userid, seatBet, int64(val))
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SSGFreeBet)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CSGFreeTrend:
		arg := msg.(*pb.CSGFreeTrend)
		glog.Debugf("CSGFreeTrend %#v", arg)
		rsp := a.freeTrends()
		ctx.Respond(rsp)
	case *pb.CSGFreeWiners:
		arg := msg.(*pb.CSGFreeWiners)
		glog.Debugf("CSGFreeWiners %#v", arg)
		rsp := a.freeWiners()
		ctx.Respond(rsp)
	case *pb.CSGFreeRoles:
		arg := msg.(*pb.CSGFreeRoles)
		glog.Debugf("CSGFreeRoles %#v", arg)
		rsp := a.freeRoles()
		ctx.Respond(rsp)
	//case *pb.CSGRoomList:
	//	arg := msg.(*pb.CSGRoomList)
	//	glog.Debugf("CSGRoomList %#v", arg)
	//	//TODO
	//case *pb.CSGCreateRoom:
	//	arg := msg.(*pb.CSGCreateRoom)
	//	glog.Debugf("CSGCreateRoom %#v", arg)
	//	//TODO
	case *pb.CSGEnterRoom:
		arg := msg.(*pb.CSGEnterRoom)
		glog.Debugf("CSGEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.privEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CSGLeave:
		arg := msg.(*pb.CSGLeave)
		glog.Debugf("CSGLeave %#v", arg)
		userid := a.getRouter(ctx)
		a.nnLeave(userid, ctx)
	case *pb.CSGReady:
		arg := msg.(*pb.CSGReady)
		glog.Debugf("CSGReady %#v", arg)
		userid := a.getRouter(ctx)
		var ready bool = arg.GetReady()
		rsp := a.readying(userid, ready)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CSGDealer:
		arg := msg.(*pb.CSGDealer)
		glog.Debugf("CSGDealer %#v", arg)
		userid := a.getRouter(ctx)
		var isDealer bool = arg.GetDealer()
		var num uint32 = arg.GetNum()
		rsp := a.choiceDealer(userid, isDealer, num)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CSGBet:
		arg := msg.(*pb.CSGBet)
		glog.Debugf("CSGBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeatbet()
		var value uint32 = arg.GetValue()
		rsp := a.choiceBet(userid, seatBet, value)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CSGiu:
		arg := msg.(*pb.CSGiu)
		glog.Debugf("CSGiu %#v", arg)
		userid := a.getRouter(ctx)
		var cards []uint32 = arg.GetCards()
		var value uint32 = arg.GetValue()
		rsp := a.choiceNiu(userid, value, cards, ctx)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CSGGameRecord:
		arg := msg.(*pb.CSGGameRecord)
		glog.Debugf("CSGGameRecord %#v", arg)
		//TODO
	case *pb.CSGLaunchVote:
		arg := msg.(*pb.CSGLaunchVote)
		glog.Debugf("CSGLaunchVote %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.launchVote(userid, 1)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CSGVote:
		arg := msg.(*pb.CSGVote)
		glog.Debugf("CSGVote %#v", arg)
		userid := a.getRouter(ctx)
		var vote uint32 = arg.GetVote()
		rsp := a.privVote(userid, vote)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}
