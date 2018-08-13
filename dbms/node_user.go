package main

import (
	"strings"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家请求处理
func (a *DBMSActor) handlerUser(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.GetConfig:
		arg := msg.(*pb.GetConfig)
		glog.Debugf("GetConfig %#v", arg)
		//ctx.Respond(handler.GetSyncConfig(arg.Type))
		//同步配置
		a.syncConfig2(ctx.Sender())
	case *pb.SyncConfig:
		//同步配置
		arg := msg.(*pb.SyncConfig)
		glog.Debugf("SyncConfig %#v", arg)
		handler.SyncConfig(arg)
	case *pb.WebRequest:
		arg := msg.(*pb.WebRequest)
		glog.Debugf("WebRequest %#v", arg)
		rsp := new(pb.WebResponse)
		rsp.Code = arg.Code
		a.handlerWeb(arg, rsp, ctx)
		ctx.Respond(rsp)
	case *pb.MatchDesk:
		arg := msg.(*pb.MatchDesk)
		glog.Debugf("MatchDesk %#v", arg)
		a.matchDesk(arg, ctx)
	case *pb.CreateDesk:
		arg := msg.(*pb.CreateDesk)
		glog.Debugf("CreateDesk %#v", arg)
		a.createDesk(arg, ctx)
	case *pb.GetRoomList:
		arg := msg.(*pb.GetRoomList)
		glog.Debugf("GetRoomList %#v", arg)
		a.getRoomList(arg, ctx)
	case *pb.CRank:
		arg := msg.(*pb.CRank)
		glog.Debugf("CRank %#v", arg)
		//TODO 缓存
		rsp := handler.PackRankMsg()
		ctx.Respond(rsp)
	case *pb.CBankLog:
		arg := msg.(*pb.CBankLog)
		glog.Debugf("CBankLog %#v", arg)
		rsp := handler.PackBankLogMsg(arg)
		ctx.Respond(rsp)
	case *pb.CNotice:
		arg := msg.(*pb.CNotice)
		glog.Debugf("CNotice %#v", arg)
		//TODO 缓存
		rsp := handler.PackUserNotice(arg)
		ctx.Respond(rsp)
	case *pb.CActivity:
		arg := msg.(*pb.CActivity)
		glog.Debugf("CActivity %#v", arg)
		//TODO 缓存
		rsp := handler.PackUserActivity(arg)
		ctx.Respond(rsp)
	case *pb.CJoinActivity:
		arg := msg.(*pb.CJoinActivity)
		glog.Debugf("CJoinActivity %#v", arg)
		rsp := handler.JoinActivity(arg)
		ctx.Respond(rsp)
	case *pb.CAgentProfitRank:
		arg := msg.(*pb.CAgentProfitRank)
		glog.Debugf("CAgentProfitRank %#v", arg)
		//TODO 缓存
		rsp := handler.PackAgentProfitRankMsg(arg)
		ctx.Respond(rsp)
	case *pb.CAgentManage:
		arg := msg.(*pb.CAgentManage)
		glog.Debugf("CAgentManage %#v", arg)
		rsp := handler.PackAgentManageMsg(arg)
		ctx.Respond(rsp)
	case *pb.CAgentProfitManage:
		arg := msg.(*pb.CAgentProfitManage)
		glog.Debugf("CAgentProfitManage %#v", arg)
		//rsp := handler.PackAgentProfitManageMsg(arg)
		rsp := handler.PackAgentProfitManageMsg2(arg)
		ctx.Respond(rsp)
	case *pb.CAgentPlayerManage:
		arg := msg.(*pb.CAgentPlayerManage)
		glog.Debugf("CAgentPlayerManage %#v", arg)
		rsp := handler.PackPlayerManageMsg(arg)
		ctx.Respond(rsp)
	case *pb.CAgentProfit:
		arg := msg.(*pb.CAgentProfit)
		glog.Debugf("CAgentProfit: %v", arg)
		rsp := handler.PackAgentProfitMsg(arg)
		ctx.Respond(rsp)
	case *pb.CAgentDayProfit:
		arg := msg.(*pb.CAgentDayProfit)
		glog.Debugf("CAgentDayProfit: %v", arg)
		rsp := handler.PackAgentDayProfitMsg(arg)
		ctx.Respond(rsp)
	case *pb.CAgentProfitOrder:
		arg := msg.(*pb.CAgentProfitOrder)
		glog.Debugf("CAgentProfitOrder %#v", arg)
		rsp := handler.PackAgentProfitOrderMsg(arg)
		ctx.Respond(rsp)
	case *pb.AgentOauth2Confirm:
		arg := msg.(*pb.AgentOauth2Confirm)
		glog.Debugf("AgentOauth2Confirm: %v", arg)
		rsp, msg2 := handler.AgentOauth2Confirm(arg)
		ctx.Respond(rsp)
		if msg2 != nil {
			rolePid.Tell(msg2)
		}
	case *pb.GetRoomRecord:
		arg := msg.(*pb.GetRoomRecord)
		glog.Debugf("GetRoomRecord %#v", arg)
		rsp := handler.PackRecordMsg(arg)
		ctx.Respond(rsp)
	case *pb.RobotMsg:
		arg := msg.(*pb.RobotMsg)
		glog.Debugf("RobotMsg %#v", arg)
		robotName := cfg.Section("robot").Name()
		if v, ok := a.serve[robotName]; ok {
			v.Tell(arg)
		}
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//匹配节点
func (a *DBMSActor) matchDesk(msg *pb.MatchDesk, ctx actor.Context) {
	rsp := new(pb.MatchedDesk)
	rsp.Rtype = msg.Rtype
	rsp.Gtype = msg.Gtype
	rsp.Dtype = msg.Dtype
	rsp.Ltype = msg.Ltype
	if msg.Name == "" {
		rsp.Error = pb.Failed
		ctx.Respond(rsp)
		return
	}
	for k, v := range a.serve {
		if strings.Contains(k, msg.Name) {
			rsp.Desk = v
			ctx.Respond(rsp)
			return
		}
	}
	glog.Errorf("match desk filed serve %#v, msg %#v", a.serve, msg)
	rsp.Error = pb.Failed
	ctx.Respond(rsp)
}

//创建房间
func (a *DBMSActor) createDesk(msg *pb.CreateDesk, ctx actor.Context) {
	for k, v := range a.serve {
		if strings.Contains(k, msg.Name) {
			v.Tell(msg)
			return
		}
	}
	glog.Errorf("create desk filed serve %#v, msg %#v", a.serve, msg)
	rsp := new(pb.CreatedDesk)
	rsp.Rtype = msg.Rtype
	rsp.Gtype = msg.Gtype
	rsp.Error = pb.Failed
	ctx.Respond(rsp)
}

//获取房间列表
func (a *DBMSActor) getRoomList(msg *pb.GetRoomList, ctx actor.Context) {
	for k, v := range a.serve {
		if strings.Contains(k, msg.Name) {
			v.Tell(msg)
			return
		}
	}
	glog.Errorf("get room filed serve %#v, msg %#v", a.serve, msg)
	rsp := new(pb.GotRoomList)
	rsp.Rtype = msg.Rtype
	rsp.Gtype = msg.Gtype
	rsp.Error = pb.Failed
	ctx.Respond(rsp)
}
