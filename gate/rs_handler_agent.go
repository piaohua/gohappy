package main

import (
	"time"

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
		rs.Send(new(pb.SAgentManage))
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
		rs.Send(new(pb.SAgentPlayerManage))
	case *pb.CAgentPlayerApprove:
		arg := msg.(*pb.CAgentPlayerApprove)
		glog.Debugf("CAgentPlayerApprove %#v", arg)
		rs.Send(new(pb.SAgentPlayerApprove))
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
	if rs.User.GetAgent() != "" || rs.User.AgentLevel != 0 {
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
			//TODO 等待审核,审核通过时修改状态、等级和绑定时间
			rs.User.Weixin = arg.GetWeixin()
			rs.User.RealName = arg.GetRealname()
		}
	}
	rs.Send(rsp)
}

//排行榜
func (rs *RoleActor) getProfitRank(arg *pb.CAgentProfitRank, ctx actor.Context) {
	if rs.User.GetAgent() != "" || rs.User.AgentLevel != 0 {
		rsp := new(pb.SAgentProfitRank)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	rs.dbmsPid.Request(arg, ctx.Self())
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
