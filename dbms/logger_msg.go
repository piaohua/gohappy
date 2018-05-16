package main

import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//Handler 消息处理
func (a *LoggerActor) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.LogRegist:
		arg := msg.(*pb.LogRegist)
		data.RegistRecord(arg.Userid, arg.Nickname, arg.Ip, arg.Atype)
	case *pb.LogLogin:
		arg := msg.(*pb.LogLogin)
		data.LoginRecord(arg.Userid, arg.Ip, arg.Atype)
	case *pb.LogLogout:
		arg := msg.(*pb.LogLogout)
		data.LogoutRecord(arg.Userid, int(arg.Event))
	case *pb.LogDiamond:
		arg := msg.(*pb.LogDiamond)
		data.DiamondRecord(arg.Userid, arg.Type, arg.Rest, arg.Num)
	case *pb.LogCoin:
		arg := msg.(*pb.LogCoin)
		data.CoinRecord(arg.Userid, arg.Type, arg.Rest, arg.Num)
	case *pb.LogCard:
		arg := msg.(*pb.LogCard)
		data.CardRecord(arg.Userid, arg.Type, arg.Rest, arg.Num)
	case *pb.LogChip:
		arg := msg.(*pb.LogChip)
		data.ChipRecord(arg.Userid, arg.Type, arg.Rest, arg.Num)
	case *pb.ServeStop:
		//关闭服务
		a.handlerStop(ctx)
		//响应登录
		rsp := new(pb.ServeStoped)
		ctx.Respond(rsp)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

func (a *LoggerActor) handlerStop(ctx actor.Context) {
	glog.Debugf("handlerStop: %s", a.Name)
	//TODO clean mailbox
}
