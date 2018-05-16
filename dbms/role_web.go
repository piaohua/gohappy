package main

import (
	"fmt"

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
		//TODO
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
