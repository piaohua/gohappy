package main

import (
	"gohappy/data"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

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
	//if user.AgentLevel == 0 || user.AgentLevel >= 4 || !handler.IsAgent(user) {
	//	rsp.Error = pb.AgentLevelLow
	//	ctx.Respond(rsp)
	//	return
	//}
	if !handler.IsAgent(user) {
		glog.Errorf("get userid %s fail", arg.GetAgentid())
		rsp.Error = pb.AgentNotExist
		ctx.Respond(rsp)
		return
	}
	if !handler.IsVaild(user) {
		rsp.Error = pb.AgentJoinLimit
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
	if arg.GetAgentid() != "" {
		//更新下级代理绑定
		msg := handler.AgentBuildUpdateMsg(arg.GetAgentid(), user.GetUserid(), 0, 0, 1)
		ctx.Self().Tell(msg)
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
		glog.Errorf("get selfid %s, %s, %d fail", arg.GetSelfid(), user.GetAgent(), user.AgentLevel)
		rsp.Error = pb.NotAgent
		ctx.Respond(rsp)
		return
	}
	if handler.IsAgent(user) {
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
	if !handler.IsAgent(user) {
		return
	}
	msg1, msg2, msg3, msg4, msg5 := handler.AddProfit(arg, user)
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
	if msg5 != nil {
		rolePid.Tell(msg5)
	}
}

//代理区域收益消息处理
func (a *RoleActor) agentProfitMonthInfo(arg *pb.AgentProfitMonthInfo, ctx actor.Context) {
	if v, ok := a.roles[arg.GetAgentid()]; ok && v != nil {
		v.Pid.Tell(arg)
		return
	}
	user := a.getUserById(arg.GetAgentid())
	if user == nil {
		glog.Errorf("get agentid %s fail", arg.GetAgentid())
		return
	}
	if !handler.IsAgent(user) {
		return
	}
	msg1, msg2, msg3, msg4, msg5, msg6 := handler.AddProfitMonth(arg, user)
	//反给上级
	if msg1 != nil {
		rolePid.Tell(msg1)
	}
	//收益日志
	if msg2 != nil {
		loggerPid.Tell(msg2)
	}
	if msg3 != nil {
		loggerPid.Tell(msg3)
	}
	if msg4 != nil || msg5 != nil {
		//rolePid.Tell(msg4)
		user.UpdateAgentProfitMonth() //暂时实时写入, TODO 异步数据更新
	}
	if msg6 != nil {
		rolePid.Tell(msg6)
	}
}

//代理申请提现
func (a *RoleActor) agentProfitApply(arg *pb.AgentProfitApply, ctx actor.Context) {
	rsp := new(pb.AgentProfitApplied)
	//rsp.Profit = arg.GetProfit()
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		rsp.Error = pb.UserDataNotExist
		ctx.Respond(rsp)
		return
	}
	record := &data.LogProfitOrder{
		Agentid:  arg.GetAgentid(),
		Userid:   arg.GetUserid(),
		Nickname: arg.GetNickname(),
		Profit:   arg.GetProfit(),
		State:    1, //默认直接发放,不再需要审批
	}
	if !record.Save() {
		rsp.Error = pb.Failed
	} else {
		rsp.Profit = arg.GetProfit()
		//user.Profit -= arg.GetProfit()
		user.SubProfit(arg.GetProfit(), arg.GetProfitFirst(), arg.GetProfitSecond())
		//暂时实时写入, TODO 异步数据更新
		user.UpdateAgentProfit()
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
	user.UpdateAgentWeek() //暂时实时写入, TODO 异步数据更新
}

//更新收益
func (a *RoleActor) agentProfitUpdate(arg *pb.AgentProfitUpdate, ctx actor.Context) {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	user.AddProfit(arg.GetIsagent(), arg.GetLevel(), arg.GetProfit())
	user.UpdateAgentProfit() //暂时实时写入, TODO 异步数据更新
}

//更新区域收益
func (a *RoleActor) agentProfitMonthUpdate(arg *pb.AgentProfitMonthUpdate, ctx actor.Context) {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	user.AddProfitMonth(arg.GetProfit())
	user.Month = int(arg.GetMonth())
	user.UpdateAgentProfitMonth() //暂时实时写入, TODO 异步数据更新
}

//更新区域收益发放
func (a *RoleActor) agentProfitMonthSend(arg *pb.AgentProfitMonthSend, ctx actor.Context) {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	handler.AgentProfitMonthSend(arg, user)
	user.UpdateAgentProfitMonth() //暂时实时写入, TODO 异步数据更新
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
		a.syncBank(arg.GetBank(), int32(pb.LOG_TYPE49), arg.GetUserid(), "")
	}
	if arg.GetProfit() != 0 {
		user.Profit += arg.GetProfit()
		//暂时实时写入, TODO 异步数据更新
		user.UpdateAgentProfit()
	}
}

//代理确认
func (a *RoleActor) agentConfirm(arg *pb.AgentConfirm, ctx actor.Context) {
	rsp := new(pb.AgentConfirmed)
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		rsp.Error = pb.UserDataNotExist
		ctx.Respond(rsp)
		return
	}
	if !handler.IsAgent(user) {
		rsp.Error = pb.Failed
	}
	ctx.Respond(rsp)
}

//申请数据更新
func (a *RoleActor) agentBuildUpdate(arg *pb.AgentBuildUpdate) {
	user := a.getUserById(arg.GetAgentid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetAgentid())
		return
	}
	if !handler.IsAgent(user) {
		glog.Errorf("agentBuildUpdate %#v", arg)
		//return
	}
	if v, ok := a.roles[user.GetUserid()]; ok {
		v.Pid.Tell(arg)
	}
	handler.AgentBuildUpdate2(arg, user) //暂时实时写入, TODO 异步数据更新
	//邀请每10人，奖励100豆子
	if arg.Build != 0 && (user.Build % 10 == 0) {
		a.sendCurrency(user.GetUserid(), 0, 100, int32(pb.LOG_TYPE55))
		//消息提醒
		record, msg2 := handler.BuildNotice(100, user.Build, arg.Userid)
		if record != nil {
			loggerPid.Tell(record)
		}
		if msg2 != nil {
			a.send2userid(user.GetUserid(), msg2)
		}
	}
}

//区域设置
func (a *RoleActor) setAgentProfitRate(arg *pb.CSetAgentProfitRate, ctx actor.Context) {
	rsp := new(pb.SSetAgentProfitRate)
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		rsp.Error = pb.UserDataNotExist
		ctx.Respond(rsp)
		return
	}
	if user.GetAgent() != arg.GetSelfid() {
		rsp.Error = pb.AgentSetLimit
	}
	if user.ProfitRate != 0 {
		rsp.Error = pb.AlreadySetRate
	}
	if !handler.IsVaild(user) {
		rsp.Error = pb.AgentSetLimit
	}
	if !handler.IsAgent(user) {
		rsp.Error = pb.NotAgent
	}
	ctx.Respond(rsp)
	if rsp.Error != pb.OK {
		return
	}
	agent := a.getUserById(arg.GetSelfid())
	if agent != nil {
		if agent.ProfitRate > arg.GetRate() {
			agent.ProfitRate -= arg.GetRate()
		} else {
			agent.ProfitRate = 1
		}
		agent.UpdateAgentProfitRate()
	}
}

//更新区域设置
func (a *RoleActor) agentProfitRate(arg *pb.SetAgentProfitRate) {
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		v.Pid.Tell(arg)
		//return
	}
	agent := a.getUserById(arg.GetUserid())
	if agent != nil {
		agent.ProfitRate += arg.GetRate()
		agent.UpdateAgentProfitRate()
	}
}

//查询代理信息
func (a *RoleActor) getAgentInfo(arg *pb.CGetAgent, ctx actor.Context) {
	user := a.getUserById(arg.GetAgentid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetAgentid())
		rsp := new(pb.SGetAgent)
		rsp.Error = pb.AgentNotExist
		ctx.Respond(rsp)
		return
	}
	rsp := handler.GetAgentMsg(user)
	ctx.Respond(rsp)
}

//设置备注
func (a *RoleActor) setAgentNote(arg *pb.CSetAgentNote, ctx actor.Context) {
	rsp := new(pb.SSetAgentNote)
	rsp.Userid = arg.GetUserid()
	rsp.Agentnote = arg.GetAgentnote()
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		rsp.Error = pb.AgentNotExist
		ctx.Respond(rsp)
		return
	}
	if user.GetAgent() != arg.GetSelfid() {
		rsp.Error = pb.AgentNotExist
		ctx.Respond(rsp)
		return
	}
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		msg := new(pb.SetAgentNote)
		msg.Userid = arg.GetUserid()
		msg.Agentnote = arg.GetAgentnote()
		v.Pid.Tell(msg)
		//return
	}
	user.AgentNote = arg.GetAgentnote()
	ctx.Respond(rsp)
	user.UpdateAgentNote()
}

//更新收益贡献值
func (a *RoleActor) bringProfit(arg *pb.AgentBringProfitNum) {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		return
	}
	if v, ok := a.roles[arg.GetUserid()]; ok && v != nil {
		v.Pid.Tell(arg)
		//return
	}
	user.AddBringProfit(arg.GetProfit())
	user.UpdateBringProfit()
}