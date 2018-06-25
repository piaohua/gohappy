package main

import (
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"gohappy/data"
	"utils"
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
			State:  arg.GetState(),
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

//代理反佣收益消息处理
func (a *RoleActor) agentProfitInfo(arg *pb.AgentProfitInfo, ctx actor.Context) {
	if v, ok := a.roles[arg.GetAgentid()]; ok && v != nil {
		v.Pid.Tell(arg)
		return
	}
	user := a.getUserById(arg.GetAgentid())
	if user == nil {
		glog.Errorf("get agentid %s fail", arg.GetAgentid())
		return
	}
	if user.AgentState != 1 {
		return
	}
	msg1, msg2, msg3, msg4 := handler.AddProfit(arg, user)
	//反给上级
	if msg1 != nil {
		rolePid.Tell(msg1)
	}
	//收益日志
	if msg2 != nil {
		loggerPid.Tell(msg2)
	}
	if msg3 != nil {
		user.UpdateAgentWeek()
	}
	if msg4 != nil {
		//暂时实时写入, TODO 异步数据更新
		user.UpdateAgentProfit()
	}
}

//代理申请提现
func (a *RoleActor) agentProfitApply(arg *pb.AgentProfitApply, ctx actor.Context) {
	rsp := new(pb.AgentProfitApplied)
	//rsp.Profit = arg.GetProfit()
	user := a.getUserById(arg.GetAgentid())
	if user == nil {
		glog.Errorf("get agentid %s fail", arg.GetAgentid())
		rsp.Error = pb.NotAgent
		ctx.Respond(rsp)
		return
	}
	record := &data.LogProfitOrder{
		Agentid: arg.GetAgentid(),
		Userid: arg.GetUserid(),
		Nickname: arg.GetNickname(),
		Profit: arg.GetProfit(),
	}
	if !record.Save() {
		rsp.Error = pb.Failed
	} else {
		rsp.Profit = arg.GetProfit()
		user.Profit -= arg.GetProfit()
	}
	ctx.Respond(rsp)
}

//代理提现受理
func (a *RoleActor) agentProfitReply(arg *pb.AgentProfitReply, ctx actor.Context) {
	rsp := new(pb.AgentProfitReplied)
	rsp.Orderid = arg.GetOrderid()
	rsp.State = arg.GetState()
	user := a.getUserById(arg.GetAgentid())
	if user == nil {
		glog.Errorf("get agentid %s fail", arg.GetAgentid())
		rsp.Error = pb.NotAgent
		ctx.Respond(rsp)
		return
	}
	record := new(data.LogProfitOrder)
	record.Get(arg.GetOrderid())
	//TODO 受理权限处理
	if record.Id != arg.GetOrderid() || record.Agentid != arg.GetAgentid() {
		glog.Errorf("profit orderid %s fail", arg.GetOrderid())
		rsp.Error = pb.ProfitOrderNotExist
		ctx.Respond(rsp)
		return
	}
	if record.State != 0 {
		glog.Errorf("profit orderid %s replied", arg.GetOrderid())
		rsp.Error = pb.ProfitOrderReplied
		ctx.Respond(rsp)
		return
	}
	if !record.Update(arg.GetState()) {
		rsp.Error = pb.Failed
	} else {
		rsp.Userid = record.Userid
		rsp.Profit = record.Profit
	}
	ctx.Respond(rsp)
}

//更新周时间
func (a *RoleActor) agentWeekUpdate(arg *pb.AgentWeekUpdate, ctx actor.Context) {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	user.WeekStart = utils.Str2Time(arg.GetStart())
	user.WeekEnd = utils.Str2Time(arg.GetEnd())
	user.UpdateAgentWeek()
}

//更新收益
func (a *RoleActor) agentProfitUpdate(arg *pb.AgentProfitUpdate, ctx actor.Context) {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	user.AddProfit(arg.GetIsagent(), arg.GetProfit())
	user.UpdateAgentProfit()
}

//提现受理消息
func (a *RoleActor) agentProfitReplyMsg(arg *pb.AgentProfitReplyMsg, ctx actor.Context) {
	if v, ok := a.roles[arg.GetUserid()]; ok && v != nil {
		v.Pid.Tell(arg)
		//return
	}
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	//TODO 优化
	if arg.GetBank() != 0 {
		a.syncBank(arg.GetBank(), int32(pb.LOG_TYPE49), arg.GetUserid())
	}
	if arg.GetProfit() != 0 {
		user.Profit += arg.GetProfit()
	}
}