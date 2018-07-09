package handler

import (
	"math"
	"time"

	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	jsoniter "github.com/json-iterator/go"
)

//PackAgentProfitRankMsg 获取排行榜信息
func PackAgentProfitRankMsg(arg *pb.CAgentProfitRank) (msg *pb.SAgentProfitRank) {
	msg = new(pb.SAgentProfitRank)
	list, err := data.GetProfitRank()
	msg.Page = arg.Page
	msg.Count = uint32(len(list))
	if err != nil {
		glog.Errorf("PackAgentProfitRankMsg err %v", err)
	}
	glog.Debugf("rank list %#v", list)
	for _, v := range list {
		msg2 := new(pb.AgentProfit)
		if val, ok := v["profit"]; ok {
			msg2.Profit = val.(int64)
		}
		if val, ok := v["_id"]; ok {
			msg2.Userid = val.(string)
		}
		if val, ok := v["nickname"]; ok {
			msg2.Nickname = val.(string)
		}
		if val, ok := v["address"]; ok {
			msg2.Address = val.(string)
		}
		if msg2.Userid == "" {
			continue
		}
		msg.List = append(msg.List, msg2)
	}
	return
}

//PackAgentManageMsg 获取代理管理列表信息
func PackAgentManageMsg(arg *pb.CAgentManage) (msg *pb.SAgentManage) {
	msg = new(pb.SAgentManage)
	list, err := data.GetAgentManage(arg)
	msg.Page = arg.Page
	msg.Count = uint32(len(list))
	if err != nil {
		glog.Errorf("PackAgentManageMsg err %v", err)
	}
	glog.Debugf("PackAgentManageMsg list %#v", list)
	for _, v := range list {
		msg2 := new(pb.AgentManage)
		if val, ok := v["profit"]; ok {
			msg2.Profit = val.(int64)
		}
		if val, ok := v["_id"]; ok {
			msg2.Agentid = val.(string)
		}
		if val, ok := v["build"]; ok {
			msg2.Num = uint32(val.(int))
		}
		if val, ok := v["agent_level"]; ok {
			msg2.Level = uint32(val.(int))
		}
		if val, ok := v["address"]; ok {
			msg2.Address = val.(string)
		}
		if msg2.Agentid == "" {
			continue
		}
		msg.List = append(msg.List, msg2)
	}
	return
}

//PackPlayerManageMsg 获取玩家管理列表信息
func PackPlayerManageMsg(arg *pb.CAgentPlayerManage) (msg *pb.SAgentPlayerManage) {
	msg = new(pb.SAgentPlayerManage)
	msg.State = arg.State
	list, err := data.GetPlayerManage(arg)
	msg.Page = arg.Page
	msg.Count = uint32(len(list))
	if err != nil {
		glog.Errorf("PackPlayerManageMsg err %v", err)
	}
	glog.Debugf("PackPlayerManageMsg list %#v", list)
	for _, v := range list {
		msg2 := new(pb.AgentPlayerManage)
		if val, ok := v["coin"]; ok {
			msg2.Coin = val.(int64)
		}
		if val, ok := v["agent"]; ok {
			msg2.Agentid = val.(string)
		}
		if val, ok := v["_id"]; ok {
			msg2.Userid = val.(string)
		}
		if val, ok := v["agent_state"]; ok {
			msg2.State = uint32(val.(int))
		}
		if val, ok := v["agent_level"]; ok {
			msg2.Level = uint32(val.(int))
		}
		if val, ok := v["address"]; ok {
			msg2.Address = val.(string)
		}
		if val, ok := v["nickname"]; ok {
			msg2.Nickname = val.(string)
		}
		if val, ok := v["agent_name"]; ok {
			msg2.Agentname = val.(string)
		}
		if val, ok := v["agent_join_time"]; ok {
			msg2.Jointime = utils.Time2LocalStr(val.(time.Time))
		}
		if msg2.Userid == "" {
			continue
		}
		msg2.State = GetAgentState(msg2.State, msg2.Level)
		msg.List = append(msg.List, msg2)
	}
	return
}

//PackAgentProfitMsg 获取代理收益明细列表
func PackAgentProfitMsg(arg *pb.CAgentProfit) (msg *pb.SAgentProfit) {
	msg = new(pb.SAgentProfit)
	list, err := data.GetAgentProfit(arg)
	msg.Page = arg.Page
	msg.Count = uint32(len(list))
	if err != nil {
		glog.Errorf("PackAgentProfitMsg err %v", err)
	}
	glog.Debugf("PackAgentProfitMsg list %#v", list)
	for _, v := range list {
		msg2 := new(pb.AgentProfitDetail)
		msg2.Userid = v.Userid //代理id
		msg2.Profit = v.Profit //收益
		msg2.Level = v.Level   //收益等级
		msg2.Gtype = v.Gtype   //game type
		msg2.Rate = v.Rate     //收益比例
		msg.List = append(msg.List, msg2)
	}
	return
}

//PackAgentProfitOrderMsg 获取收益订单列表
func PackAgentProfitOrderMsg(arg *pb.CAgentProfitOrder) (msg *pb.SAgentProfitOrder) {
	msg = new(pb.SAgentProfitOrder)
	list, err := data.GetProfitOrder(arg)
	msg.Page = arg.Page
	msg.Type = arg.Type
	msg.Count = uint32(len(list))
	if err != nil {
		glog.Errorf("PackAgentProfitOrderMsg err %v", err)
	}
	glog.Debugf("PackAgentProfitOrderMsg list %#v", list)
	for _, v := range list {
		msg2 := new(pb.AgentProfitOrder)
		msg2.Orderid = v.Id                               //代理id
		msg2.Userid = v.Userid                            //提单人id
		msg2.Nickname = v.Nickname                        //代理id
		msg2.Profit = v.Profit                            //收益
		msg2.Applytime = utils.Time2LocalStr(v.ApplyTime) //提单时间
		msg2.Replytime = utils.Time2LocalStr(v.ReplyTime) //响应时间
		msg2.State = v.State                              //状态,0等待处理,1成功,2失败
		msg.List = append(msg.List, msg2)
	}
	return
}

//AgentJoinMsg 申请消息
func AgentJoinMsg(user *data.User) (msg *pb.AgentJoin) {
	msg = new(pb.AgentJoin)
	msg.Agentname = user.AgentName
	msg.Agentid = user.Agent
	msg.Realname = user.RealName
	msg.Weixin = user.Weixin
	msg.Level = user.AgentLevel
	msg.Time = utils.Time2Str(user.AgentJoinTime)
	msg.Userid = user.GetUserid()
	return
}

//AgentJoin2User 更新玩家申请数据
func AgentJoin2User(msg *pb.AgentJoin, user *data.User) {
	user.AgentName = msg.Agentname
	user.Agent = msg.Agentid
	user.RealName = msg.Realname
	user.Weixin = msg.Weixin
	user.AgentLevel = msg.Level
	user.AgentJoinTime = utils.Str2Time(msg.Time)
	user.AgentState = 1 //默认通过，不用审核
	SetAgentProfitRate(user)
	return
}

//SetAgentProfitRate 默认抽成设置，1级作为大代理不再分成
func SetAgentProfitRate(user *data.User) {
	switch user.AgentLevel {
	case 2:
		user.ProfitRate = 10 //TODO 优化可配置
	case 3:
		user.ProfitRate = 20
	case 4:
		user.ProfitRate = 50
	}
}

//AgentApprove 审批,修改状态
func AgentApprove(state pb.AgentApproveState, selfid string, user *data.User) pb.ErrCode {
	switch state {
	case pb.AgentAgreed:
		user.AgentState = 1 //通过
	case pb.AgentRefused:
		user.AgentLevel = 0 //清除
		user.AgentState = 0 //拒绝
		user.AgentName = ""
		user.RealName = ""
		user.Weixin = ""
	default:
		return pb.Failed
	}
	return pb.OK
}

//GetAgentState 返回代理状态0不是代理，1已经是代理，2审核中
func GetAgentState(state, level uint32) uint32 {
	if state == 1 {
		return 1
	} else if state == 0 && level != 0 {
		return 2 //审核中
	}
	return 0
}

//AgentProfitInfoMsg 代理收益消息
func AgentProfitInfoMsg(userid, agentid string, agent bool, gtype int32,
	level, rate uint32, profit int64) (msg *pb.AgentProfitInfo) {
	msg = &pb.AgentProfitInfo{
		Userid:  userid,
		Agentid: agentid,
		Gtype:   gtype,
		Level:   level,
		Rate:    rate,
		Profit:  profit,
		Agent:   agent,
	}
	return
}

//AgentProfitNumMsg 收益消息
func AgentProfitNumMsg(userid string, gtype int32, profit int64) (msg *pb.AgentProfitNum) {
	msg = &pb.AgentProfitNum{
		Gtype:  gtype,
		Profit: profit,
		Userid: userid,
	}
	return
}

//AgentProfitNumMsg 收益
func AddProfit(arg *pb.AgentProfitInfo, user *data.User) (msg *pb.AgentProfitInfo,
	msg2 *pb.LogProfit, msg3 *pb.AgentWeekUpdate, msg4 *pb.AgentProfitUpdate) {
	num := int64(math.Trunc(float64(user.ProfitRate)/100) * float64(arg.Profit))
	profit := arg.Profit - num
	user.AddProfit(arg.Agent, profit)
	if UpdateWeekProfit(profit, user) {
		//更新时间消息
		msg3 = UpdateWeekMsg(user)
	}
	//日志消息
	msg2 = LogProfitMsg(arg.Agentid, arg.Userid, arg.Gtype, arg.Level, arg.Rate, profit)
	//更新消息
	msg4 = &pb.AgentProfitUpdate{
		Userid:  user.GetUserid(),
		Profit:  profit,
		Isagent: arg.Agent,
	}
	if num <= 0 {
		return
	}
	//反给上级消息
	msg = AgentProfitInfoMsg(user.GetUserid(), user.GetAgent(), false,
		arg.Gtype, user.AgentLevel, user.ProfitRate, num)
	if user.AgentState == 1 {
		msg.Agent = true
	}
	return
}

//UpdateWeekProfit 更新周收益统计
func UpdateWeekProfit(num int64, user *data.User) bool {
	if user.AgentState != 1 {
		return false
	}
	now := utils.LocalTime()
	if user.WeekStart.Before(now) && user.WeekEnd.After(now) {
		return false
	}
	user.WeekPlayerProfit = num
	user.WeekProfit = num
	user.WeekStart, user.WeekEnd = utils.ThisWeek()
	return true
}

//UpdateWeekMsg 更新周统计时间消息
func UpdateWeekMsg(user *data.User) (msg *pb.AgentWeekUpdate) {
	msg = &pb.AgentWeekUpdate{
		Userid: user.GetUserid(),
		Start:  utils.Time2Str(user.WeekStart),
		End:    utils.Time2Str(user.WeekEnd),
	}
	return
}

//AgentOauth2Confirm 代理授权关系确认
func AgentOauth2Confirm(arg *pb.AgentOauth2Confirm) (msg *pb.AgentOauth2Confirmed) {
	msg = new(pb.AgentOauth2Confirmed)
	userInfo := new(data.UserInfo)
	err := jsoniter.Unmarshal(arg.Userinfo, userInfo)
	if err != nil {
		glog.Errorf("AgentOauth2Confirm err %v, arg %#v", err, arg)
		msg.Error = pb.Failed
		return
	}
	userInfo.Agentid = arg.GetAgentid()
	if userInfo.Agentid == "" {
		msg.Error = pb.Failed
		return
	}
	glog.Debugf("userInfo %#v", userInfo)
	if !userInfo.Save() {
		glog.Errorf("userInfo save failed %#v", userInfo)
		msg.Error = pb.Failed
	}
	return
}

//AgentBuildUpdateMsg 绑定数量消息
func AgentBuildUpdateMsg(agentid, userid string, build, vaild, child uint32) (msg *pb.AgentBuildUpdate) {
	msg = &pb.AgentBuildUpdate{
		AgentChild: child,
		BuildVaild: vaild,
		Build:      build,
		Agentid:    agentid,
		Userid:     userid,
	}
	return
}

//AgentBuildUpdate 绑定数量更新
func AgentBuildUpdate(msg *pb.AgentBuildUpdate, user *data.User) {
	if msg.Build != 0 {
		user.Build += msg.Build
	}
	if msg.BuildVaild != 0 {
		user.BuildVaild += msg.BuildVaild
	}
	if msg.AgentChild != 0 {
		user.AgentChild += msg.AgentChild
	}
}

//AgentBuildUpdate2 绑定数量实时更新写入
func AgentBuildUpdate2(msg *pb.AgentBuildUpdate, user *data.User) {
	if msg.Build != 0 {
		user.Build += msg.Build
		user.UpdateBuild()
	}
	if msg.BuildVaild != 0 {
		user.BuildVaild += msg.BuildVaild
		user.UpdateBuildVaild()
	}
	if msg.AgentChild != 0 {
		user.AgentChild += msg.AgentChild
		user.UpdateAgentChild()
	}
}