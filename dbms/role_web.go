package main

import (
	"fmt"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//web请求处理
func (a *RoleActor) handlerWeb(arg *pb.WebRequest,
	rsp *pb.WebResponse, ctx actor.Context) {
	switch arg.Code {
	case pb.WebOnline:
		msg1 := make([]string, 0)
		err1 := json.Unmarshal(arg.Data, &msg1)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//响应
		resp := make(map[string]int)
		for _, v := range msg1 {
			if _, ok := a.online[v]; ok {
				resp[v] = 1
			} else {
				resp[v] = 0
			}
		}
		result, err2 := json.Marshal(resp)
		if err2 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err2)
			return
		}
		rsp.Result = result
	case pb.WebBuild:
		//后台设置绑定关系
		msg2 := new(pb.SetAgentBuild)
		err1 := msg2.Unmarshal(arg.Data)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		err2 := a.setBuild(msg2)
		if err2 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err2)
		}
	case pb.WebGive:
		//后台货币赠送同步到game房间
		msg2 := new(pb.PayCurrency)
		err1 := msg2.Unmarshal(arg.Data)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//消息
		a.msg2role(msg2)
	case pb.WebNumber:
		result, err2 := a.getNumber(ctx)
		if err2 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err2)
			return
		}
		rsp.Result = result
	case pb.WebRate:
		//后台设置区域奖励百分比
		msg2 := new(pb.SetAgentProfitRate)
		err1 := msg2.Unmarshal(arg.Data)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		err2 := a.setRate(msg2)
		if err2 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err2)
		}
	case pb.WebState:
		//后台设置代理
		msg2 := new(pb.SetAgentState)
		err1 := msg2.Unmarshal(arg.Data)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		err2 := a.setState(msg2)
		if err2 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err2)
		}
	case pb.WebVaild:
		//后台设置代理
		msg2 := new(pb.AgentBuildUpdate)
		err1 := msg2.Unmarshal(arg.Data)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		a.agentBuildUpdate(msg2)
	default:
		glog.Errorf("unknown message %v", arg)
	}
}

//消息通知到玩家
func (a *RoleActor) msg2role(arg *pb.PayCurrency) {
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		//存活在节点中
		v.Pid.Tell(arg)
		return
	}
	a.syncCurrency(arg.Diamond, arg.Coin, arg.Card,
		arg.Chip, arg.Type, arg.Userid)
}

//获取在线人数
func (a *RoleActor) getNumber(ctx actor.Context) ([]byte, error) {
	//响应1 机器人,2 玩家
	resp := make(map[int]int)
	for _, v := range a.online {
		if v.GetRobot() {
			resp[1]++
		} else {
			resp[2]++
		}
	}
	result, err2 := json.Marshal(resp)
	if err2 != nil {
		glog.Errorf("msg err: %v", err2)
		return nil, fmt.Errorf("msg err: %v", err2)
	}
	return result, nil
}

//设置区域奖励百分比
func (a *RoleActor) setRate(arg *pb.SetAgentProfitRate) error {
	agent := a.getUserById(arg.GetUserid())
	if agent == nil {
		return fmt.Errorf("userid %s not exist", arg.GetUserid())
	}
	if arg.GetRate() > 38 || arg.GetRate() == 0 {
		return fmt.Errorf("rate %d error", arg.GetRate())
	}
	if agent.ProfitRate != 0 {
		//return fmt.Errorf("AlreadySetRate")
	}
	if !handler.IsVaild(agent) {
		//return fmt.Errorf("ProfitLimit")
	}
	if !handler.IsAgent(agent) {
		return fmt.Errorf("NotAgent")
	}
	a.agentProfitRate(arg)
	return nil
}

//后台设置绑定关系
func (a *RoleActor) setBuild(arg *pb.SetAgentBuild) error {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		return fmt.Errorf("userid %s not exist", arg.GetUserid())
	}
	//TODO 限制条件,绑定数量统计和日志
	if arg.GetAgent() != "" {
		agent := a.getUserById(arg.GetAgent())
		if agent == nil {
			return fmt.Errorf("agent %s not exist", arg.GetAgent())
		}
		if !handler.IsAgent(agent) {
			return fmt.Errorf("NotAgent")
		}
	}
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		v.Pid.Tell(arg)
		//return
	}
	handler.SetAgentBuild(arg, user)
	user.UpdateAgent()
	return nil
}

//后台设置代理
func (a *RoleActor) setState(arg *pb.SetAgentState) error {
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		return fmt.Errorf("userid %s not exist", arg.GetUserid())
	}
	//TODO 限制条件
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		v.Pid.Tell(arg)
		//return
	}
	handler.SetAgentState(arg, user)
	user.UpdateAgentJoin()
	return nil
}
