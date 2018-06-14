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
	case *pb.CNNCoinEnterRoom:
		arg := msg.(*pb.CNNCoinEnterRoom)
		glog.Debugf("CNNCoinEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.coinEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CNNFreeEnterRoom:
		arg := msg.(*pb.CNNFreeEnterRoom)
		glog.Debugf("CNNFreeEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.freeEnterMsg(userid)
		ctx.Respond(msg1)
		a.freeCameinMsg(userid)
	case *pb.CNNFreeDealer:
		arg := msg.(*pb.CNNFreeDealer)
		glog.Debugf("CNNFreeDealer %#v", arg)
		userid := a.getRouter(ctx)
		var state int32 = arg.GetState()
		var num uint32 = arg.GetCoin()
		errcode := a.beDealer(userid, state, num)
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SNNFreeDealer)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CNNFreeDealerList:
		arg := msg.(*pb.CNNFreeDealerList)
		glog.Debugf("CNNFreeDealerList %#v", arg)
		rsp := a.dealerListMsg()
		ctx.Respond(rsp)
	case *pb.CNNSit:
		arg := msg.(*pb.CNNSit)
		glog.Debugf("CNNSit %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.freeSit(userid, arg)
		if rsp.Error == pb.OK {
			a.broadcast(rsp)
			return
		}
		ctx.Respond(rsp)
	case *pb.CNNFreeBet:
		arg := msg.(*pb.CNNFreeBet)
		glog.Debugf("CNNFreeBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeat()
		var val uint32 = arg.GetValue()
		errcode := a.freeBet(userid, seatBet, int64(val))
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SNNFreeBet)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CNNFreeTrend:
		arg := msg.(*pb.CNNFreeTrend)
		glog.Debugf("CNNFreeTrend %#v", arg)
		rsp := a.freeTrends()
		ctx.Respond(rsp)
	case *pb.CNNFreeWiners:
		arg := msg.(*pb.CNNFreeWiners)
		glog.Debugf("CNNFreeWiners %#v", arg)
		rsp := a.freeWiners()
		ctx.Respond(rsp)
	case *pb.CNNFreeRoles:
		arg := msg.(*pb.CNNFreeRoles)
		glog.Debugf("CNNFreeRoles %#v", arg)
		rsp := a.freeRoles()
		ctx.Respond(rsp)
	//case *pb.CNNRoomList:
	//	arg := msg.(*pb.CNNRoomList)
	//	glog.Debugf("CNNRoomList %#v", arg)
	//	//TODO
	//case *pb.CNNCreateRoom:
	//	arg := msg.(*pb.CNNCreateRoom)
	//	glog.Debugf("CNNCreateRoom %#v", arg)
	//	//TODO
	case *pb.CNNEnterRoom:
		arg := msg.(*pb.CNNEnterRoom)
		glog.Debugf("CNNEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.privEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CNNLeave:
		arg := msg.(*pb.CNNLeave)
		glog.Debugf("CNNLeave %#v", arg)
		userid := a.getRouter(ctx)
		a.nnLeave(userid, ctx)
	case *pb.CNNReady:
		arg := msg.(*pb.CNNReady)
		glog.Debugf("CNNReady %#v", arg)
		userid := a.getRouter(ctx)
		var ready bool = arg.GetReady()
		rsp := a.readying(userid, ready)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CNNDealer:
		arg := msg.(*pb.CNNDealer)
		glog.Debugf("CNNDealer %#v", arg)
		userid := a.getRouter(ctx)
		var isDealer bool = arg.GetDealer()
		var num uint32 = arg.GetNum()
		rsp := a.choiceDealer(userid, isDealer, num)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CNNBet:
		arg := msg.(*pb.CNNBet)
		glog.Debugf("CNNBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeatbet()
		var value uint32 = arg.GetValue()
		rsp := a.choiceBet(userid, seatBet, value)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CNNiu:
		arg := msg.(*pb.CNNiu)
		glog.Debugf("CNNiu %#v", arg)
		userid := a.getRouter(ctx)
		var cards []uint32 = arg.GetCards()
		var value uint32 = arg.GetValue()
		rsp := a.choiceNiu(userid, value, cards, ctx)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CNNGameRecord:
		arg := msg.(*pb.CNNGameRecord)
		glog.Debugf("CNNGameRecord %#v", arg)
		//TODO
	case *pb.CNNLaunchVote:
		arg := msg.(*pb.CNNLaunchVote)
		glog.Debugf("CNNLaunchVote %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.launchVote(userid, 1)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CNNVote:
		arg := msg.(*pb.CNNVote)
		glog.Debugf("CNNVote %#v", arg)
		userid := a.getRouter(ctx)
		var vote uint32 = arg.GetVote()
		rsp := a.privVote(userid, vote)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CNNCoinChangeRoom:
		arg := msg.(*pb.CNNCoinChangeRoom)
		glog.Debugf("CNNCoinChangeRoom %#v", arg)
		a.changeDesk(ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}
