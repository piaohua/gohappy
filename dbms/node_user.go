package main

import (
	"strings"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家请求处理
func (a *DBMSActor) handlerUser(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.GetConfig:
		arg := msg.(*pb.GetConfig)
		glog.Debugf("GetConfig %#v", arg)
		//ctx.Respond(handler.GetSyncConfig(arg.Type))
		//同步配置
		a.syncConfig2(ctx.Sender())
	case *pb.SyncConfig:
		//同步配置
		arg := msg.(*pb.SyncConfig)
		glog.Debugf("SyncConfig %#v", arg)
		handler.SyncConfig(arg)
	case *pb.WebRequest:
		arg := msg.(*pb.WebRequest)
		glog.Debugf("WebRequest %#v", arg)
		rsp := new(pb.WebResponse)
		rsp.Code = arg.Code
		a.handlerWeb(arg, rsp, ctx)
		ctx.Respond(rsp)
	case *pb.MatchDesk:
		arg := msg.(*pb.MatchDesk)
		glog.Debugf("MatchDesk %#v", arg)
		a.matchDesk(arg, ctx)
	case *pb.CreateDesk:
		arg := msg.(*pb.CreateDesk)
		glog.Debugf("CreateDesk %#v", arg)
		a.createDesk(arg, ctx)
	case *pb.GetRoomList:
		arg := msg.(*pb.GetRoomList)
		glog.Debugf("GetRoomList %#v", arg)
		a.getRoomList(arg, ctx)
	case *pb.CRank:
		arg := msg.(*pb.CRank)
		glog.Debugf("CRank %#v", arg)
		//TODO 缓存
		rsp := handler.PackRankMsg()
		ctx.Respond(rsp)
	case *pb.GetRoomRecord:
		arg := msg.(*pb.GetRoomRecord)
		glog.Debugf("GetRoomRecord %#v", arg)
		rsp := handler.PackRecordMsg(arg)
		ctx.Respond(rsp)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//匹配节点
func (a *DBMSActor) matchDesk(msg *pb.MatchDesk, ctx actor.Context) {
	rsp := new(pb.MatchedDesk)
	rsp.Rtype = msg.Rtype
	rsp.Gtype = msg.Gtype
	rsp.Dtype = msg.Dtype
	rsp.Ltype = msg.Ltype
	if msg.Name == "" {
		rsp.Error = pb.Failed
		ctx.Respond(rsp)
		return
	}
	for k, v := range a.serve {
		if strings.Contains(k, msg.Name) {
			rsp.Desk = v
			ctx.Respond(rsp)
			return
		}
	}
	rsp.Error = pb.Failed
	ctx.Respond(rsp)
}

//创建房间
func (a *DBMSActor) createDesk(msg *pb.CreateDesk, ctx actor.Context) {
	for k, v := range a.serve {
		if strings.Contains(k, msg.Name) {
			v.Tell(msg)
			return
		}
	}
	rsp := new(pb.CreatedDesk)
	rsp.Rtype = msg.Rtype
	rsp.Gtype = msg.Gtype
	rsp.Error = pb.Failed
	ctx.Respond(rsp)
}

//获取房间列表
func (a *DBMSActor) getRoomList(msg *pb.GetRoomList, ctx actor.Context) {
	for k, v := range a.serve {
		if strings.Contains(k, msg.Name) {
			v.Tell(msg)
			return
		}
	}
	rsp := new(pb.GotRoomList)
	rsp.Rtype = msg.Rtype
	rsp.Gtype = msg.Gtype
	rsp.Error = pb.Failed
	ctx.Respond(rsp)
}
