package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"gohappy/game/handler"
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
	rsp.Level = user.AgentLevel + 1
	ctx.Respond(rsp)
}

//申请数据更新
func (a *RoleActor) syncAgentJoin(arg *pb.AgentJoin, ctx actor.Context) {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	handler.AgentJoin2User(arg, user)
	//暂时实时写入, TODO 异步数据更新
	user.UpdateAgentJoin()
}

//审批
func (a *RoleActor) agentApprove(arg *pb.CAgentPlayerApprove, ctx actor.Context) {
	rsp := new(pb.SAgentPlayerApprove)
	rsp.Userid = arg.GetUserid()
	rsp.State = arg.GetState()
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		rsp.Error = pb.UserDataNotExist
		ctx.Respond(rsp)
		return
	}
	if user.GetAgent() != arg.GetSelfid() || user.AgentLevel == 0 {
		glog.Errorf("get selfid %s fail", arg.GetSelfid())
		rsp.Error = pb.NotAgent
		ctx.Respond(rsp)
		return
	}
	if user.AgentState == 1 {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		rsp.Error = pb.AlreadyAgent
		ctx.Respond(rsp)
		return
	}
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		msg := &pb.AgentPlayerApprove{
			State: arg.GetState(),
			Userid: arg.GetUserid(),
			Selfid: arg.GetSelfid(),
		}
		v.Pid.Tell(msg)
	}
	errcode := handler.AgentApprove(arg.GetState(), arg.GetSelfid(), user)
	rsp.Error = errcode
	ctx.Respond(rsp)
	//暂时实时写入, TODO 异步数据更新
	user.UpdateAgentJoin()
}