package main

import (
	"gohappy/data"
	"gohappy/game/handler"
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
	case *pb.LogTask:
		arg := msg.(*pb.LogTask)
		data.TaskRecord(arg.Userid, arg.Taskid, arg.Type)
	case *pb.LogNotice:
		arg := msg.(*pb.LogNotice)
		handler.SaveNotice(arg)
	case *pb.LogBank:
		arg := msg.(*pb.LogBank)
		data.BankRecord(arg.Userid, arg.Type, arg.Rest, arg.Num)
	case *pb.RoomRecordInfo:
		arg := msg.(*pb.RoomRecordInfo)
		handler.Log2RoomRecord(arg)
	case *pb.RoleRecord:
		arg := msg.(*pb.RoleRecord)
		handler.Log2RoleRecord(arg)
	case *pb.RoundRecord:
		arg := msg.(*pb.RoundRecord)
		handler.Log2RoundRecord(arg)
	case *pb.LogProfit:
		arg := msg.(*pb.LogProfit)
		data.ProfitRecord(arg.Agentid, arg.Userid, arg.Gtype,
			arg.Level, arg.Rate, arg.Profit)
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
