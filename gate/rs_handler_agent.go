package main

import (
	"time"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家数据请求处理
func (rs *RoleActor) handlerAgent(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CAgentJoin:
		arg := msg.(*pb.CAgentJoin)
		glog.Debugf("CAgentJoin %#v", arg)
		rs.agentJoin(arg, ctx)
	case *pb.CMyAgent:
		arg := msg.(*pb.CMyAgent)
		glog.Debugf("CMyAgent %#v", arg)
		rs.agentInfo()
	case *pb.CAgentManage:
		arg := msg.(*pb.CAgentManage)
		glog.Debugf("CAgentManage %#v", arg)
		rs.agentManage(arg, ctx)
	case *pb.CAgentProfit:
		arg := msg.(*pb.CAgentProfit)
		glog.Debugf("CAgentProfit: %v", arg)
		rs.Send(new(pb.SAgentProfit))
	case *pb.CAgentProfitOrder:
		arg := msg.(*pb.CAgentProfitOrder)
		glog.Debugf("CAgentProfitOrder %#v", arg)
		rs.Send(new(pb.SAgentProfitOrder))
	case *pb.CAgentProfitApply:
		arg := msg.(*pb.CAgentProfitApply)
		glog.Debugf("CAgentProfitApply %#v", arg)
		rs.Send(new(pb.SAgentProfitApply))
	case *pb.CAgentProfitRank:
		arg := msg.(*pb.CAgentProfitRank)
		glog.Debugf("CAgentProfitRank %#v", arg)
		rs.getProfitRank(arg, ctx)
	case *pb.CAgentPlayerManage:
		arg := msg.(*pb.CAgentPlayerManage)
		glog.Debugf("CAgentPlayerManage %#v", arg)
		rs.playerManage(arg, ctx)
	case *pb.CAgentPlayerApprove:
		arg := msg.(*pb.CAgentPlayerApprove)
		glog.Debugf("CAgentPlayerApprove %#v", arg)
		rs.agentApprove(arg, ctx)
	case *pb.AgentPlayerApprove:
		arg := msg.(*pb.AgentPlayerApprove)
		glog.Debugf("AgentPlayerApprove %#v", arg)
		errcode := handler.AgentApprove(arg.GetState(), arg.GetSelfid(), rs.User)
		glog.Debugf("AgentPlayerApprove errcode %v", errcode)
		//TODO 更新房间内玩家数据
	case *pb.AgentProfitInfo:
		arg := msg.(*pb.AgentProfitInfo)
		glog.Debugf("AgentProfitInfo %#v", arg)
		//TODO 收益日志
		//arg.Agentid = rs.User.GetAgent()
	//case proto.Message:
	//	//响应消息
	//	rs.Send(msg)
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerDesk(msg, ctx)
	}
}

//基础信息
func (rs *RoleActor) agentInfo() {
	rsp := new(pb.SMyAgent)
	rsp.Agentname = rs.User.AgentName
	rsp.Agentid = rs.User.Agent
	rsp.Address = rs.User.Address
	rsp.Profit = rs.User.Profit
	rsp.WeekProfit = rs.User.WeekProfit
	rsp.HistoryProfit = rs.User.HistoryProfit
	rsp.SubPlayerProfit = rs.User.SubPlayerProfit
	rsp.SubAgentProfit = rs.User.SubAgentProfit
	rs.Send(rsp)
}

//申请加入
func (rs *RoleActor) agentJoin(arg *pb.CAgentJoin, ctx actor.Context) {
	rsp := new(pb.SAgentJoin)
	if rs.User.BankPhone == "" {
		rsp.Error = pb.BankNotOpen
		rs.Send(rsp)
		return
	}
	//TODO rs.User.Agent 加入游戏时已经绑定
	if rs.User.AgentLevel != 0 && rs.User.AgentState == 0 {
		rsp.Error = pb.WaitForAudit
		rs.Send(rsp)
		return
	}
	if rs.User.AgentState == 1 || rs.User.AgentLevel != 0 {
		rsp.Error = pb.AlreadyBuild
		rs.Send(rsp)
		return
	}
	if arg.GetAgentid() == "" || arg.GetAgentname() == "" ||
		arg.GetRealname() == "" || arg.GetWeixin() == "" {
		rsp.Error = pb.ParamError
		rs.Send(rsp)
		return
	}
	res1 := rs.reqRole(arg, ctx)
	if response1, ok := res1.(*pb.SAgentJoin); ok {
		rsp.Error = response1.Error
		if response1.Error == pb.OK {
			rs.User.Agent = arg.GetAgentid()
			rs.User.AgentName = arg.GetAgentname()
			rs.User.Weixin = arg.GetWeixin()
			rs.User.RealName = arg.GetRealname()
			rs.User.AgentLevel = response1.Level
			rs.User.AgentJoinTime = time.Now()
			rsp.Level = response1.Level
			//更新数据库,等待审核,审核通过时修改状态, TODO 添加申请日志
			msg := handler.AgentJoinMsg(rs.User)
			rs.rolePid.Request(msg, ctx.Self())
		}
	}
	rs.Send(rsp)
}

//排行榜
func (rs *RoleActor) getProfitRank(arg *pb.CAgentProfitRank, ctx actor.Context) {
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp := new(pb.SAgentProfitRank)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	rs.dbmsPid.Request(arg, ctx.Self())
}

//代理管理列表
func (rs *RoleActor) agentManage(arg *pb.CAgentManage, ctx actor.Context) {
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp := new(pb.SAgentManage)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Userid = rs.User.GetUserid()
	rs.dbmsPid.Request(arg, ctx.Self())
}

//玩家管理列表
func (rs *RoleActor) playerManage(arg *pb.CAgentPlayerManage, ctx actor.Context) {
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp := new(pb.SAgentPlayerManage)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Selfid = rs.User.GetUserid()
	rs.dbmsPid.Request(arg, ctx.Self())
}

//审批
func (rs *RoleActor) agentApprove(arg *pb.CAgentPlayerApprove, ctx actor.Context) {
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp := new(pb.SAgentPlayerApprove)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Selfid = rs.User.GetUserid()
	rs.rolePid.Request(arg, ctx.Self())
}

func (rs *RoleActor) reqRole(msg interface{}, ctx actor.Context) interface{} {
	glog.Debugf("reqRole msg %#v", msg)
	if rs.rolePid == nil {
		glog.Errorf("reqRole err %#v", msg)
		return nil
	}
	timeout := 3 * time.Second
	res1, err1 := rs.rolePid.RequestFuture(msg, timeout).Result()
	if err1 != nil {
		glog.Errorf("reqRole err: %v, msg %#v", err1, msg)
		return nil
	}
	return res1
}
