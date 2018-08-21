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
	case *pb.CLHFreeEnterRoom:
		arg := msg.(*pb.CLHFreeEnterRoom)
		glog.Debugf("CLHFreeEnterRoom %#v", arg)
		userid := a.getRouter(ctx)
		msg1 := a.freeEnterMsg(userid)
		ctx.Respond(msg1)
		a.freeCameinMsg(userid)
		a.callRobot()
	case *pb.CLHFreeDealer:
		arg := msg.(*pb.CLHFreeDealer)
		glog.Debugf("CLHFreeDealer %#v", arg)
		userid := a.getRouter(ctx)
		var state int32 = arg.GetState()
		var num uint32 = arg.GetCoin()
		errcode := a.beDealer(userid, state, num)
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SLHFreeDealer)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CLHFreeDealerList:
		arg := msg.(*pb.CLHFreeDealerList)
		glog.Debugf("CLHFreeDealerList %#v", arg)
		rsp := a.dealerListMsg()
		ctx.Respond(rsp)
	case *pb.CLHSit:
		arg := msg.(*pb.CLHSit)
		glog.Debugf("CLHSit %#v", arg)
		userid := a.getRouter(ctx)
		rsp := a.freeSit(userid, arg)
		if rsp.Error == pb.OK {
			a.broadcast(rsp)
			return
		}
		ctx.Respond(rsp)
	case *pb.CLHFreeBet:
		arg := msg.(*pb.CLHFreeBet)
		glog.Debugf("CLHFreeBet %#v", arg)
		userid := a.getRouter(ctx)
		var seatBet uint32 = arg.GetSeat()
		var val uint32 = arg.GetValue()
		errcode := a.freeBet(userid, seatBet, int64(val))
		if errcode == pb.OK {
			return
		}
		//响应
		rsp := new(pb.SLHFreeBet)
		rsp.Error = errcode
		ctx.Respond(rsp)
	case *pb.CLHFreeTrend:
		arg := msg.(*pb.CLHFreeTrend)
		glog.Debugf("CLHFreeTrend %#v", arg)
		rsp := a.freeTrends()
		ctx.Respond(rsp)
	case *pb.CLHFreeWiners:
		arg := msg.(*pb.CLHFreeWiners)
		glog.Debugf("CLHFreeWiners %#v", arg)
		rsp := a.freeWiners()
		ctx.Respond(rsp)
	case *pb.CLHFreeRoles:
		arg := msg.(*pb.CLHFreeRoles)
		glog.Debugf("CLHFreeRoles %#v", arg)
		rsp := a.freeRoles()
		ctx.Respond(rsp)
	case *pb.CLHLeave:
		arg := msg.(*pb.CLHLeave)
		glog.Debugf("CLHLeave %#v", arg)
		userid := a.getRouter(ctx)
		a.nnLeave(userid, ctx)
	case *pb.BankGive:
		arg := msg.(*pb.BankGive)
		glog.Debugf("BankGive %#v", arg)
		if v, ok := a.roles[arg.GetUserid()]; ok && v != nil {
			v.User.AddCoin(arg.GetCoin())
		}
	default:
		glog.Errorf("unknown message %v", msg)
	}
}
