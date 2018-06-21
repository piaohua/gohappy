package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//申请
func (a *RoleActor) agentJoin(arg *pb.CAgentJoin, ctx actor.Context) {
	rsp := new(pb.SAgentJoin)
	user := a.getUserById(arg.GetAgentid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetAgentid())
		rsp.Error = pb.AgentNotExist
		ctx.Respond(rsp)
		return
	}
	//等级1，2，3，4
	if user.AgentLevel == 0 || user.AgentLevel >= 4 || user.AgentState != 1 {
		rsp.Error = pb.AgentLevelLow
		ctx.Respond(rsp)
		return
	}
	ctx.Respond(rsp)
}
