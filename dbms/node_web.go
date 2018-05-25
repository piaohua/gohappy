package main

import (
	"fmt"
	"strings"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//web请求处理
func (a *DBMSActor) handlerWeb(arg *pb.WebRequest,
	rsp *pb.WebResponse, ctx actor.Context) {
	switch arg.Code {
	case pb.WebShop:
		//更新配置
		msg2 := handler.SyncConfig2(pb.CONFIG_SHOP, arg.Atype, arg.Data)
		err1 := handler.SyncConfig(msg2)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//广播所有节点,主动通知同步配置,只同步修改数据
		a.broadcast(msg2, ctx)
	case pb.WebEnv:
		//更新配置
		msg2 := handler.SyncConfig2(pb.CONFIG_ENV, arg.Atype, arg.Data)
		err1 := handler.SyncConfig(msg2)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//广播所有节点,主动通知同步配置,只同步修改数据
		a.broadcast(msg2, ctx)
	case pb.WebNotice:
		//更新配置
		msg2 := handler.SyncConfig2(pb.CONFIG_NOTICE, arg.Atype, arg.Data)
		err1 := handler.SyncConfig(msg2)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//广播所有节点,主动通知同步配置,只同步修改数据
		a.broadcast(msg2, ctx)
	case pb.WebGame:
		//更新配置
		msg2 := handler.SyncConfig2(pb.CONFIG_GAMES, arg.Atype, arg.Data)
		err1 := handler.SyncConfig(msg2)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//广播所有节点,主动通知同步配置,只同步修改数据
		a.broadcast(msg2, ctx)
	case pb.WebVip:
		//更新配置
		msg2 := handler.SyncConfig2(pb.CONFIG_VIP, arg.Atype, arg.Data)
		err1 := handler.SyncConfig(msg2)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//广播所有节点,主动通知同步配置,只同步修改数据
		a.broadcast(msg2, ctx)
	case pb.WebTask:
		//更新配置
		msg2 := handler.SyncConfig2(pb.CONFIG_TASK, arg.Atype, arg.Data)
		err1 := handler.SyncConfig(msg2)
		if err1 != nil {
			rsp.ErrMsg = fmt.Sprintf("msg err: %v", err1)
			return
		}
		//广播所有节点,主动通知同步配置,只同步修改数据
		a.broadcast(msg2, ctx)
	default:
		glog.Errorf("unknown message %v", arg)
	}
}

//广播所有节点,游戏逻辑服,dbms
func (a *DBMSActor) broadcast(msg interface{}, ctx actor.Context) {
	for k, v := range a.serve {
		if strings.Contains(k, "gate.") ||
			strings.Contains(k, "game.") {
			v.Tell(msg)
		}
	}
}
