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
	case *pb.CEBCoinEnterRoom:
		arg := msg.(*pb.CEBCoinEnterRoom)
		glog.Debugf("CEBCoinEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.coinEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CEBFreeEnterRoom:
		arg := msg.(*pb.CEBFreeEnterRoom)
		glog.Debugf("CEBFreeEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.freeEnterMsg(userid)
		ctx.Respond(msg1)
		a.freeCameinMsg(userid)
	case *pb.CEBFreeDealer:
		arg := msg.(*pb.CEBFreeDealer)
		glog.Debugf("CEBFreeDealer %#v", arg)
		userid := a.getRouter(ctx)
		var state int32 = arg.GetState()
		var num uint32 = arg.GetCoin()
		errcode := a.beDealer(userid, state, num)
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SEBFreeDealer)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CEBFreeDealerList:
		arg := msg.(*pb.CEBFreeDealerList)
		glog.Debugf("CEBFreeDealerList %#v", arg)
		rsp := a.dealerListMsg()
		ctx.Respond(rsp)
	case *pb.CEBSit:
		arg := msg.(*pb.CEBSit)
		glog.Debugf("CEBSit %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.freeSit(userid, arg)
		if rsp.Error == pb.OK {
			a.broadcast(rsp)
			return
		}
		ctx.Respond(rsp)
	case *pb.CEBFreeBet:
		arg := msg.(*pb.CEBFreeBet)
		glog.Debugf("CEBFreeBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeat()
		var val uint32 = arg.GetValue()
		errcode := a.freeBet(userid, seatBet, int64(val))
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SEBFreeBet)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CEBFreeTrend:
		arg := msg.(*pb.CEBFreeTrend)
		glog.Debugf("CEBFreeTrend %#v", arg)
		rsp := a.freeTrends()
		ctx.Respond(rsp)
	case *pb.CEBFreeWiners:
		arg := msg.(*pb.CEBFreeWiners)
		glog.Debugf("CEBFreeWiners %#v", arg)
		rsp := a.freeWiners()
		ctx.Respond(rsp)
	case *pb.CEBFreeRoles:
		arg := msg.(*pb.CEBFreeRoles)
		glog.Debugf("CEBFreeRoles %#v", arg)
		rsp := a.freeRoles()
		ctx.Respond(rsp)
	//case *pb.CEBRoomList:
	//	arg := msg.(*pb.CEBRoomList)
	//	glog.Debugf("CEBRoomList %#v", arg)
	//	//TODO
	//case *pb.CEBCreateRoom:
	//	arg := msg.(*pb.CEBCreateRoom)
	//	glog.Debugf("CEBCreateRoom %#v", arg)
	//	//TODO
	case *pb.CEBEnterRoom:
		arg := msg.(*pb.CEBEnterRoom)
		glog.Debugf("CEBEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.privEnterMsg(userid)
		ctx.Respond(msg1)
		a.coinCameinMsg(userid)
	case *pb.CEBLeave:
		arg := msg.(*pb.CEBLeave)
		glog.Debugf("CEBLeave %#v", arg)
		userid := a.getRouter(ctx)
		a.nnLeave(userid, ctx)
	case *pb.CEBReady:
		arg := msg.(*pb.CEBReady)
		glog.Debugf("CEBReady %#v", arg)
		userid := a.getRouter(ctx)
		var ready bool = arg.GetReady()
		rsp := a.readying(userid, ready)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CEBDealer:
		arg := msg.(*pb.CEBDealer)
		glog.Debugf("CEBDealer %#v", arg)
		userid := a.getRouter(ctx)
		var isDealer bool = arg.GetDealer()
		var num uint32 = arg.GetNum()
		rsp := a.choiceDealer(userid, isDealer, num)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CEBBet:
		arg := msg.(*pb.CEBBet)
		glog.Debugf("CEBBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeatbet()
		var value uint32 = arg.GetValue()
		rsp := a.choiceBet(userid, seatBet, value)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CEBiu:
		arg := msg.(*pb.CEBiu)
		glog.Debugf("CEBiu %#v", arg)
		userid := a.getRouter(ctx)
		var cards []uint32 = arg.GetCards()
		var value uint32 = arg.GetValue()
		rsp := a.choiceNiu(userid, value, cards, ctx)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CEBGameRecord:
		arg := msg.(*pb.CEBGameRecord)
		glog.Debugf("CEBGameRecord %#v", arg)
		//TODO
	case *pb.CEBLaunchVote:
		arg := msg.(*pb.CEBLaunchVote)
		glog.Debugf("CEBLaunchVote %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.launchVote(userid, 1)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CEBVote:
		arg := msg.(*pb.CEBVote)
		glog.Debugf("CEBVote %#v", arg)
		userid := a.getRouter(ctx)
		var vote uint32 = arg.GetVote()
		rsp := a.privVote(userid, vote)
		if rsp.Error == pb.OK {
			return
		}
		ctx.Respond(rsp)
	case *pb.CEBCoinChangeRoom:
		arg := msg.(*pb.CEBCoinChangeRoom)
		glog.Debugf("CEBCoinChangeRoom %#v", arg)
		a.changeDesk(ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}
