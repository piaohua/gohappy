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
	case *pb.CAgentDayProfit:
		arg := msg.(*pb.CAgentDayProfit)
		glog.Debugf("CAgentDayProfit: %v", arg)
		rs.agentDayProfit(arg, ctx)
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
	case *pb.AgentProfitMonthInfo:
		arg := msg.(*pb.AgentProfitMonthInfo)
		glog.Debugf("AgentProfitMonthInfo %#v", arg)
		rs.agentProfitMonthInfo(arg)
	case *pb.AgentProfitNum:
		arg := msg.(*pb.AgentProfitNum)
		glog.Debugf("AgentProfitNum %#v", arg)
		rs.agentProfitNum(arg)
	case *pb.AgentProfitReplyMsg:
		arg := msg.(*pb.AgentProfitReplyMsg)
		glog.Debugf("AgentProfitReplyMsg %#v", arg)
		rs.agentProfitReplyMsg(arg)
	case *pb.AgentBuildUpdate:
		arg := msg.(*pb.AgentBuildUpdate)
		glog.Debugf("AgentBuildUpdate: %v", arg)
		handler.AgentBuildUpdate(arg, rs.User)
	case *pb.CSetAgentProfitRate:
		arg := msg.(*pb.CSetAgentProfitRate)
		glog.Debugf("CSetAgentProfitRate %#v", arg)
		rs.setAgentProfitRate(arg, ctx)
	case *pb.SetAgentProfitRate:
		arg := msg.(*pb.SetAgentProfitRate)
		glog.Debugf("SetAgentProfitRate %#v", arg)
		rs.agentProfitRate(arg, ctx)
	case *pb.SetAgentNote:
		arg := msg.(*pb.SetAgentNote)
		glog.Debugf("SetAgentNote %#v", arg)
		rs.setAgentNote(arg, ctx)
	case *pb.SetAgentBuild:
		arg := msg.(*pb.SetAgentBuild)
		glog.Debugf("SetAgentBuild %#v", arg)
		handler.SetAgentBuild(arg, rs.User)
	case *pb.SetAgentState:
		arg := msg.(*pb.SetAgentState)
		glog.Debugf("SetAgentState %#v", arg)
		handler.SetAgentState(arg, rs.User)
	case *pb.CGetAgent:
		arg := msg.(*pb.CGetAgent)
		glog.Debugf("CGetAgent %#v", arg)
		rs.getAgentInfo(arg, ctx)
	case *pb.CSetAgentNote:
		arg := msg.(*pb.CSetAgentNote)
		glog.Debugf("CSetAgentNote %#v", arg)
		rs.setAgentNotes(arg, ctx)
	case *pb.CAgentProfitManage:
		arg := msg.(*pb.CAgentProfitManage)
		glog.Debugf("CAgentProfitManage %#v", arg)
		rs.agentProfitManage(arg, ctx)
	case *pb.SAgentProfitManage:
		arg := msg.(*pb.SAgentProfitManage)
		glog.Debugf("SAgentProfitManage %#v", arg)
		rs.agentProfitManage2(arg, ctx)
	case *pb.AgentBringProfitNum:
		arg := msg.(*pb.AgentBringProfitNum)
		glog.Debugf("AgentBringProfitNum %#v", arg)
		rs.User.AddBringProfit(arg.GetProfit())
	case *pb.AgentActivityProfit:
		arg := msg.(*pb.AgentActivityProfit)
		glog.Debugf("AgentActivityProfit %#v", arg)
		rs.gentActivityProfit(arg, ctx)
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
	if !handler.IsAgent(rs.User) {
		return
	}
	msg1, msg2, msg3, msg4, msg5 := handler.AddProfit(arg, rs.User)
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
	if msg5 != nil {
		rs.rolePid.Tell(msg5)
	}
}

//代理区域收益消息处理
func (rs *RoleActor) agentProfitMonthInfo(arg *pb.AgentProfitMonthInfo) {
	if !handler.IsAgent(rs.User) {
		return
	}
	msg1, msg2, msg3, msg4, msg5, msg6 := handler.AddProfitMonth(arg, rs.User)
	//反给上级
	if msg1 != nil {
		rs.rolePid.Tell(msg1)
	}
	//收益日志
	if msg2 != nil {
		rs.loggerPid.Tell(msg2)
	}
	if msg3 != nil {
		rs.loggerPid.Tell(msg3)
	}
	//更新数据
	if msg4 != nil {
		rs.rolePid.Tell(msg4)
	}
	if msg5 != nil {
		rs.rolePid.Tell(msg5)
	}
	if msg6 != nil {
		rs.rolePid.Tell(msg6)
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
	msg2 := handler.AgentProfitInfoMsg(rs.User.GetUserid(), rs.User.GetNickname(), rs.User.GetAgentNote(),
		rs.User.GetAgent(), false, arg.Gtype, 1, 0, rest) //level表示相对当前代理的等级,不是rs.User.AgentLevel
	msg3 := handler.AgentProfitMonthInfoMsg(rs.User.GetUserid(), rs.User.GetNickname(), rs.User.GetAgentNote(),
		rs.User.GetAgent(), false, arg.Gtype, 1, 0, rest) //level表示相对当前代理的等级,不是rs.User.AgentLevel
	if handler.IsAgent(rs.User) {
		msg2.Agent = true
		msg3.Agent = true
	}
	rs.rolePid.Tell(msg2)
	rs.rolePid.Tell(msg3)
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
	rsp.Build = rs.User.Build
	rsp.BuildVaild = rs.User.BuildVaild
	rsp.AgentChild = rs.User.AgentChild
	//rsp.ProfitRate = rs.User.ProfitRate
	rsp.ProfitRate = rs.User.ProfitRateSum
	rsp.ProfitMonth = rs.User.ProfitMonth
	rsp.AgentTitle = handler.GetAgentTitle(rs.User)
	rsp.ProfitFirst = rs.User.ProfitFirst
	rsp.ProfitSecond = rs.User.ProfitSecond
	rsp.ProfitLastMonth = rs.User.ProfitLastMonth
	if rsp.AgentTitle == 1 { //合伙人
		//rsp.ProfitRate = 23 //固定值展示
	}
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
	if rs.User.AgentLevel != 0 && !handler.IsAgent(rs.User) {
		rsp.Error = pb.WaitForAudit
		rs.Send(rsp)
		return
	}
	if rs.User.AgentLevel != 0 || handler.IsAgent(rs.User) {
		rsp.Error = pb.AlreadyAgent
		rs.Send(rsp)
		return
	}
	if len(arg.GetAgentname()) > 50 || len(arg.GetRealname()) > 50 ||
		len(arg.GetWeixin()) > 50 {
		rsp.Error = pb.NameTooLong
		rs.Send(rsp)
		return
	}
	//rs.User.Agent 加入游戏时已经绑定
	if rs.User.GetAgent() != "" && rs.User.GetAgent() != arg.GetAgentid() {
		arg.Agentid = rs.User.GetAgent()
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
			//rs.User.Agent = arg.GetAgentid()
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
			//handler.SetAgentProfitRate(rs.User) //TODO 优化为消息同步
		}
	}
	rs.Send(rsp)
}

//排行榜
func (rs *RoleActor) getProfitRank(arg *pb.CAgentProfitRank, ctx actor.Context) {
	if handler.IsNotAgent(rs.User) {
		rsp := new(pb.SAgentProfitRank)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	rs.dbmsPid.Request(arg, ctx.Self())
}

//代理管理列表
func (rs *RoleActor) agentManage(arg *pb.CAgentManage, ctx actor.Context) {
	if handler.IsNotAgent(rs.User) {
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
	if handler.IsNotAgent(rs.User) {
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
	if handler.IsNotAgent(rs.User) {
		rsp := new(pb.SAgentPlayerApprove)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	//权限限制(有效玩家3个以上)
	if !handler.IsVaild(rs.User) {
		rsp := new(pb.SAgentPlayerApprove)
		rsp.Error = pb.ProfitLimit
		rs.Send(rsp)
		return
	}
	arg.Selfid = rs.User.GetUserid()
	rs.rolePid.Request(arg, ctx.Self())
}

//代理收益明细列表
func (rs *RoleActor) agentProfit(arg *pb.CAgentProfit, ctx actor.Context) {
	if handler.IsNotAgent(rs.User) {
		rsp := new(pb.SAgentProfit)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Agentid = rs.User.GetUserid()
	rs.dbmsPid.Request(arg, ctx.Self())
}

//代理天收益明细列表
func (rs *RoleActor) agentDayProfit(arg *pb.CAgentDayProfit, ctx actor.Context) {
	if handler.IsNotAgent(rs.User) {
		rsp := new(pb.SAgentDayProfit)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Selfid = rs.User.GetUserid()
	rs.dbmsPid.Request(arg, ctx.Self())
}

//代理收益订单列表
func (rs *RoleActor) agentProfitOrder(arg *pb.CAgentProfitOrder, ctx actor.Context) {
	if handler.IsNotAgent(rs.User) {
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
	profit := arg.GetProfit() * 100
	if handler.IsNotAgent(rs.User) {
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	profit = uint32(rs.User.Profit + rs.User.ProfitFirst + rs.User.ProfitSecond) //默认全部提取
	//if rs.User.Profit < int64(profit) || !handler.IsProfitApply(rs.User) {
	if profit <= 0 || !handler.IsProfitApply(rs.User) {
		rsp.Error = pb.ProfitNotEnough
		rs.Send(rsp)
		return
	}
	//权限限制(有效玩家3个以上,绑定10个以上),合伙人无限制
	if !handler.IsVaild(rs.User) && handler.GetAgentTitle(rs.User) != 1 {
		rsp.Error = pb.ProfitLimit
		rs.Send(rsp)
		return
	}
	msg := &pb.AgentProfitApply{
		Agentid:  rs.User.GetAgent(),    //受理人userid
		Userid:   rs.User.GetUserid(),   //申请人玩家id
		Nickname: rs.User.GetNickname(), //玩家昵称
		//Profit:   int64(profit),         //提取金额
		Profit:       rs.User.GetProfit(),
		ProfitFirst:  rs.User.GetProfitFirst(),
		ProfitSecond: rs.User.GetProfitSecond(),
	}
	res1 := rs.reqRole(msg, ctx)
	if response1, ok := res1.(*pb.AgentProfitApplied); ok {
		rsp.Error = response1.Error
		if response1.Error == pb.OK {
			//扣除收益
			//rs.User.Profit -= response1.Profit
			rs.User.SubProfit(response1.GetProfit(), response1.GetProfitFirst(), response1.GetProfitSecond())
			//默认直接发放,不再需要审批,TODO 保留小数位
			//rs.addBank(response1.Profit/100, int32(pb.LOG_TYPE49), "")
			rs.addCurrency(0, int64(profit/100), 0, 0, int32(pb.LOG_TYPE49))
			//消息提醒
			record2, msg2 := handler.ProfitNotice(int64(profit/100), rs.User.GetUserid())
			if record2 != nil {
				rs.loggerPid.Tell(record2)
			}
			if msg2 != nil {
				rs.Send(msg2)
			}
		}
	}
	rsp.Profit = profit
	rs.Send(rsp)
}

//设置区域收益
func (rs *RoleActor) setAgentProfitRate(arg *pb.CSetAgentProfitRate, ctx actor.Context) {
	rsp := new(pb.SSetAgentProfitRate)
	rsp.Userid = arg.GetUserid()
	rsp.Rate = arg.GetRate()
	if handler.IsNotAgent(rs.User) {
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	if arg.GetPassword() != rs.User.BankPassword {
		rsp.Error = pb.PwdError
		rs.Send(rsp)
		return
	}
	if !handler.IsSetProfitRate(rs.User) {
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	if arg.GetRate() == 0 {
		rsp.Error = pb.ParamError
		rs.Send(rsp)
		return
	}
	profitRate := handler.GetChildProfitRate(arg.GetUserid(), rs.User)
	title := handler.GetAgentTitle(rs.User)
	if !((profitRate >= (arg.GetRate() + 5) && title == 2) || (profitRate >= (arg.GetRate() + 3) && title == 1)) { //保留5%
		rsp.Error = pb.ProfitRateNotEnough
		rs.Send(rsp)
		return
	}
	if !handler.IsVaild(rs.User) {
		rsp.Error = pb.AgentSetLimit
		rs.Send(rsp)
		return
	}
	arg.Selfid = rs.User.GetUserid()
	res1 := rs.reqRole(arg, ctx)
	if response1, ok := res1.(*pb.SSetAgentProfitRate); ok {
		rsp.Error = response1.Error
		if response1.Error == pb.OK {
			//更新同步数据
			msg1 := &pb.SetAgentProfitRate{
				Userid: arg.GetUserid(),
				Rate:   arg.GetRate(),
			}
			rs.rolePid.Tell(msg1)
			//rs.User.ProfitRate -= arg.GetRate() //TODO 优化为消息同步
			if profitRate > arg.GetRate() {
				profitRate -= arg.GetRate()
				handler.SetChildProfitRate(arg.GetUserid(), profitRate, rs.User)
			} else {
				glog.Errorf("CSetAgentProfitRate filed %#v, rate %#v", arg, rs.User.ProfitRate)
				handler.SetChildProfitRate(arg.GetUserid(), 1, rs.User)
			}
		}
	}
	rs.Send(rsp)
}

//设置区域收益
func (rs *RoleActor) agentProfitRate(arg *pb.SetAgentProfitRate, ctx actor.Context) {
	handler.SetAgentProfitRate(arg.GetRate(), rs.User)
}

//奖励发放更新收益
func (rs *RoleActor) gentActivityProfit(arg *pb.AgentActivityProfit, ctx actor.Context) {
	//奖励发放
	rs.addCurrency(0, arg.GetProfit(), 0, 0, arg.GetType())
	//消息提醒
	//record, msg2 := handler.ActNotice(arg)
	//if record != nil {
	//	rs.loggerPid.Tell(record)
	//}
	//if msg2 != nil {
	//	rs.Send(msg2)
	//}
}

//设置区域备注
func (rs *RoleActor) setAgentNote(arg *pb.SetAgentNote, ctx actor.Context) {
	rs.User.AgentNote = arg.GetAgentnote()
}

//收益提现受理
func (rs *RoleActor) agentProfitReply(arg *pb.CAgentProfitReply, ctx actor.Context) {
	rsp := new(pb.SAgentProfitReply)
	rsp.Orderid = arg.GetOrderid()
	rsp.State = arg.GetState()
	if handler.IsNotAgent(rs.User) {
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

//查询代理信息
func (rs *RoleActor) getAgentInfo(arg *pb.CGetAgent, ctx actor.Context) {
	if rs.User.GetAgent() == "" {
		rsp := new(pb.SGetAgent)
		rsp.Error = pb.AgentNotExist
		return
	}
	arg.Agentid = rs.User.GetAgent()
	rs.rolePid.Request(arg, ctx.Self())
}

//设置代理备注
func (rs *RoleActor) setAgentNotes(arg *pb.CSetAgentNote, ctx actor.Context) {
	if len(arg.GetAgentnote()) > 50 {
		rsp := new(pb.SSetAgentNote)
		rsp.Error = pb.NameTooLong
		return
	}
	arg.Selfid = rs.User.GetUserid()
	rs.rolePid.Request(arg, ctx.Self())
}

//代理管理列表
func (rs *RoleActor) agentProfitManage(arg *pb.CAgentProfitManage, ctx actor.Context) {
	if handler.IsNotAgent(rs.User) {
		rsp := new(pb.SAgentProfitManage)
		rsp.Error = pb.NotAgent
		rs.Send(rsp)
		return
	}
	arg.Userid = rs.User.GetUserid()
	rs.dbmsPid.Request(arg, ctx.Self())
}

//代理管理列表, 玩家数据整合为一条
func (rs *RoleActor) agentProfitManage2(arg *pb.SAgentProfitManage, ctx actor.Context) {
	list := make([]*pb.AgentProfitManage, 0)
	msg2 := &pb.AgentProfitManage{
		AgentTitle: handler.GetAgentTitle(rs.User),
		Agentid:    rs.User.GetUserid(),
		Nickname:   rs.User.GetNickname(),
		Agentnote:  rs.User.AgentNote,
		//Rate:       rs.User.ProfitRate,
		Rate:       rs.User.ProfitRateSum,
	}
	msg3 := new(pb.AgentProfitManage)
	*msg3 = *msg2
	for _, v := range arg.List {
		if v.GetAgentTitle() == 4 || v.GetAgentid() == rs.User.GetUserid() {
			msg2.BringProfit += v.GetBringProfit()
			msg2.Count++
			continue
		} else if v.GetAgentTitle() == 3 && v.GetVaild() >= 3 {
			msg3.BringProfit += v.GetBringProfit()
			msg3.Count++
			continue
		}
		list = append(list, v)
	}
	if msg2.BringProfit != 0 {
		list = append(list, msg2)
	}
	if msg3.BringProfit != 0 {
		list = append(list, msg3)
	}
	arg.List = list
	rs.Send(arg)
}
