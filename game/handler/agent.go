package handler

import (
	"time"

	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
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
			msg2.Num = val.(uint32)
		}
		if val, ok := v["agent_level"]; ok {
			msg2.Level = val.(uint32)
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
			msg2.State = pb.AgentApproveState(val.(uint32))
		}
		if val, ok := v["agent_level"]; ok {
			msg2.Level = val.(uint32)
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
			msg2.Jointime = utils.Time2Str(val.(time.Time))
		}
		if msg2.Userid == "" {
			continue
		}
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
	return
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

//AgentProfitInfoMsg 收益消息
func AgentProfitInfoMsg(agentid string, agent bool, gtype int32,
	level, rate uint32, profit int64) (msg *pb.AgentProfitInfo) {
	msg = &pb.AgentProfitInfo{
		Agentid: agentid,
		Gtype:   gtype,
		Level:   level,
		Rate:    rate,
		Profit:  profit,
		Agent:   agent,
	}
	return
}
