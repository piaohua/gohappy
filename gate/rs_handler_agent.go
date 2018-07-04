package main

import (
	"math"
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
		rs.agentProfit(arg, ctx)
	case *pb.CAgentProfitOrder:
		arg := msg.(*pb.CAgentProfitOrder)
		glog.Debugf("CAgentProfitOrder %#v", arg)
		rs.agentProfitOrder(arg, ctx)
	case *pb.CAgentProfitApply:
		arg := msg.(*pb.CAgentProfitApply)
		glog.Debugf("CAgentProfitApply %#v", arg)
		rs.agentProfitApply(arg, ctx)
	case *pb.CAgentProfitReply:
		arg := msg.(*pb.CAgentProfitReply)
		glog.Debugf("CAgentProfitReply %#v", arg)
		rs.agentProfitReply(arg, ctx)
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
	case *pb.AgentProfitInfo:
		arg := msg.(*pb.AgentProfitInfo)
		glog.Debugf("AgentProfitInfo %#v", arg)
		rs.agentProfitInfo(arg)
	case *pb.AgentProfitNum:
		arg := msg.(*pb.AgentProfitNum)
		glog.Debugf("AgentProfitNum %#v", arg)
		rs.agentProfitNum(arg)
	case *pb.AgentProfitReplyMsg:
		arg := msg.(*pb.AgentProfitReplyMsg)
		glog.Debugf("AgentProfitReplyMsg %#v", arg)
		rs.agentProfitReplyMsg(arg)
	//case proto.Message:
	//	//响应消息
	//	rs.Send(msg)
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerDesk(msg, ctx)
	}
}

//代理反佣收益消息处理
func (rs *RoleActor) agentProfitInfo(arg *pb.AgentProfitInfo) {
	if rs.User.AgentState != 1 {
		return
	}
	msg1, msg2, msg3, msg4 := handler.AddProfit(arg, rs.User)
	//反给上级
	if msg1 != nil {
		rs.rolePid.Tell(msg1)
	}
	//收益日志
	if msg2 != nil {
		rs.loggerPid.Tell(msg2)
	}
	//更新时间
	if msg3 != nil {
		rs.rolePid.Tell(msg3)
	}
	//更新数据
	if msg4 != nil {
		rs.rolePid.Tell(msg4)
	}
}

//玩家收益消息处理
func (rs *RoleActor) agentProfitNum(arg *pb.AgentProfitNum) {
	if rs.User.GetAgent() == "" {
		return
	}
	if arg.GetProfit() <= 0 {
		return
	}
	//系统抽成
	num := int64(math.Trunc(float64(arg.GetProfit()) * 0.1))
	rest := arg.GetProfit() - num
	if num > 0 {
		msg1 := handler.LogSysProfitMsg(rs.User.GetAgent(), arg.Userid,
			arg.Gtype, rs.User.AgentLevel, 10, num, rest)
		rs.loggerPid.Tell(msg1)
	}
	if rest <= 0 {
		return
	}
	//发送消息给代理
	msg2 := handler.AgentProfitInfoMsg(rs.User.GetUserid(), rs.User.GetAgent(),
		false, arg.Gtype, rs.User.AgentLevel, 100, rest)
	if rs.User.AgentState == 1 {
		msg2.Agent = true
	}
	rs.rolePid.Tell(msg2)
}

//基础信息
func (rs *RoleActor) agentInfo() {
	if handler.UpdateWeekProfit(0, rs.User) {
		msg := handler.UpdateWeekMsg(rs.User)
		rs.rolePid.Tell(msg)
	}
	rsp := new(pb.SMyAgent)
	rsp.Agentname = rs.User.AgentName
	rsp.Agentid = rs.User.Agent
	rsp.Address = rs.User.Address
	rsp.Profit = rs.User.Profit
	rsp.WeekProfit = rs.User.WeekProfit
	rsp.WeekProfit = rs.User.WeekPlayerProfit
	rsp.HistoryProfit = rs.User.HistoryProfit
	rsp.SubPlayerProfit = rs.User.SubPlayerProfit
	rsp.SubAgentProfit = rs.User.SubAgentProfit
	rsp.Level = rs.User.AgentLevel
	rsp.State = handler.GetAgentState(rs.User.AgentState, rs.User.AgentLevel)
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
		rsp.Error = pb.AlreadyAgent
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
			rs.User.AgentState = 1 //默认通过，不用审核
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
	//TODO 权限限制(有效玩家3个以上)
	arg.Selfid = rs.User.GetUserid()
	rs.rolePid.Request(arg, ctx.Self())
}

//代理收益明细列表
func (rs *RoleActor) agentProfit(arg *pb.CAgentProfit, ctx actor.Context) {
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp := new(pb.SAgentProfit)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Agentid = rs.User.GetUserid()
	rs.dbmsPid.Request(arg, ctx.Self())
}

//代理收益订单列表
func (rs *RoleActor) agentProfitOrder(arg *pb.CAgentProfitOrder, ctx actor.Context) {
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp := new(pb.SAgentProfitOrder)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Type = 1 //固定只能查自己订单
	arg.Agentid = rs.User.GetUserid()
	rs.dbmsPid.Request(arg, ctx.Self())
}

//收益提现申请
func (rs *RoleActor) agentProfitApply(arg *pb.CAgentProfitApply, ctx actor.Context) {
	rsp := new(pb.SAgentProfitApply)
	rsp.Profit = arg.GetProfit()
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	if rs.User.Profit < int64(arg.GetProfit()) {
		rsp.Error = pb.ProfitNotEnough
		rs.Send(rsp)
		return
	}
	//TODO 权限限制(有效玩家3个以上)
	msg := &pb.AgentProfitApply{
		Agentid:  rs.User.GetAgent(),     //受理人userid
		Userid:   rs.User.GetUserid(),    //申请人玩家id
		Nickname: rs.User.GetNickname(),  //玩家昵称
		Profit:   int64(arg.GetProfit()), //提取金额
	}
	res1 := rs.reqRole(msg, ctx)
	if response1, ok := res1.(*pb.AgentProfitApplied); ok {
		rsp.Error = response1.Error
		if response1.Error == pb.OK {
			//扣除收益
			rs.User.Profit -= response1.Profit
			//默认直接发放,不再需要审批
			rs.addBank(response1.Profit, int32(pb.LOG_TYPE49))
		}
	}
	rs.Send(rsp)
}

//收益提现受理
func (rs *RoleActor) agentProfitReply(arg *pb.CAgentProfitReply, ctx actor.Context) {
	rsp := new(pb.SAgentProfitReply)
	rsp.Orderid = arg.GetOrderid()
	rsp.State = arg.GetState()
	if rs.User.AgentLevel == 0 || rs.User.AgentState != 1 {
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	if arg.GetState() != 1 && arg.GetState() != 2 {
		rsp.Error = pb.Failed
		rs.Send(rsp)
		return
	}
	msg := &pb.AgentProfitReply{
		Orderid: arg.GetOrderid(),    //order id
		Agentid: rs.User.GetUserid(), //受理人userid
		State:   arg.GetState(),      //状态,1同意,2拒绝
	}
	res1 := rs.reqRole(msg, ctx)
	if response1, ok := res1.(*pb.AgentProfitReplied); ok {
		rsp.Error = response1.Error
		if response1.Error == pb.OK {
			rsp.Profit = response1.Profit
			if arg.GetState() == 1 {
				//银行账户增加收入
				msg1 := &pb.AgentProfitReplyMsg{
					Userid: response1.Userid,
					Bank:   response1.Profit,
				}
				rs.rolePid.Tell(msg1)
			} else if arg.GetState() == 2 {
				//返还收益
				msg1 := &pb.AgentProfitReplyMsg{
					Userid: response1.Userid,
					Profit: response1.Profit,
				}
				rs.rolePid.Tell(msg1)
			}
		}
	}
	rs.Send(rsp)
}

//提现受理消息
func (rs *RoleActor) agentProfitReplyMsg(arg *pb.AgentProfitReplyMsg) {
	if arg.GetBank() != 0 {
		rs.User.AddBank(arg.GetBank())
	}
	if arg.GetProfit() != 0 {
		rs.User.Profit += arg.GetProfit()
	}
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
