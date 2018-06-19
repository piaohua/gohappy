package main

import (
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
	case *pb.CMyAgent:
		arg := msg.(*pb.CMyAgent)
		glog.Debugf("CMyAgent %#v", arg)
	case *pb.CAgentManage:
		arg := msg.(*pb.CAgentManage)
		glog.Debugf("CAgentManage %#v", arg)
	case *pb.CAgentProfit:
		arg := msg.(*pb.CAgentProfit)
		glog.Debugf("CAgentProfit: %v", arg)
	case *pb.CAgentProfitOrder:
		arg := msg.(*pb.CAgentProfitOrder)
		glog.Debugf("CAgentProfitOrder %#v", arg)
	case *pb.CAgentProfitApply:
		arg := msg.(*pb.CAgentProfitApply)
		glog.Debugf("CAgentProfitApply %#v", arg)
	case *pb.CAgentProfitRank:
		arg := msg.(*pb.CAgentProfitRank)
		glog.Debugf("CAgentProfitRank %#v", arg)
	case *pb.CAgentPlayerManage:
		arg := msg.(*pb.CAgentPlayerManage)
		glog.Debugf("CAgentPlayerManage %#v", arg)
	case *pb.CAgentPlayerApprove:
		arg := msg.(*pb.CAgentPlayerApprove)
		glog.Debugf("CAgentPlayerApprove %#v", arg)
	//case proto.Message:
	//	//响应消息
	//	rs.Send(msg)
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerDesk(msg, ctx)
	}
}