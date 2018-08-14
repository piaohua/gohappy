package main

import (
	"time"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//Handler 消息处理
func (a *Desk) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.ServeStop:
		//关闭服务
		glog.Infof("ServeStop: %#v", msg)
		a.handlerStop(ctx)
		//响应登录
		//rsp := new(pb.ServeStoped)
		//ctx.Respond(rsp)
	case *pb.ServeStart:
		a.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.Tick:
		a.ding(ctx)
	default:
		//glog.Errorf("unknown message %v", msg)
		a.handlerLogic(msg, ctx)
	}
}

//启动服务
func (a *Desk) start(ctx actor.Context) {
	glog.Infof("desk start: %v", ctx.Self().String())
	//初始化
	a.InitDesk()
	//启动
	go a.ticker(ctx)
	//日志记录
	switch a.DeskData.Rtype {
	case int32(pb.ROOM_TYPE1): //私人
		msg := handler.LogRoomRecordMsg(a.DeskData)
		a.loggerPid.Tell(msg)
	}
}

//时钟
func (a *Desk) ticker(ctx actor.Context) {
	tick := time.Tick(time.Second)
	msg := new(pb.Tick)
	for {
		select {
		case <-a.stopCh:
			glog.Info("desk ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-a.stopCh:
			glog.Info("desk ticker closed")
			return
		case <-tick:
			ctx.Self().Tell(msg)
		}
	}
}

//钟声
func (a *Desk) ding(ctx actor.Context) {
	//glog.Debugf("ding: %v", ctx.Self().String())
	//逻辑处理
	switch a.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		a.coinTimeout()
	case int32(pb.ROOM_TYPE1): //私人
		a.privTimeout()
	case int32(pb.ROOM_TYPE2): //百人
		a.freeTimeout()
	}
}

//关闭时钟
func (a *Desk) closeTick() {
	select {
	case <-a.stopCh:
		return
	default:
		//停止发送消息
		close(a.stopCh)
	}
}

func (a *Desk) handlerStop(ctx actor.Context) {
	glog.Infof("handlerStop: %s", a.Name)
	//关闭
	a.closeTick()
	//玩家结算退出
	switch a.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		//a.gameOver()
		switch a.state {
		case int32(pb.STATE_NIU):
			a.niuTimeout()
		}
	case int32(pb.ROOM_TYPE1): //私人
		//解散结算消息
		if !a.checkOver() {
			//TODO 结算后就不会退款了
			//a.gameOver()
			switch a.state {
			case int32(pb.STATE_NIU):
				a.niuTimeout()
			}
		}
	case int32(pb.ROOM_TYPE2): //百人
		switch a.state {
		case int32(pb.STATE_BET):
			a.freeGameOver()
		}
	}
	//玩家退出
	for k := range a.roles {
		errcode := a.leave(k)
		glog.Debugf("stop userid %s, err %v", k, errcode)
		if errcode != pb.OK {
			//continue
		}
		//离开状态消息
		a.userLeaveDesk(k)
	}
	//关闭房间消息
	msg := new(pb.CloseDesk)
	msg.Roomid = a.DeskData.Rid
	msg.Code = a.DeskData.Code
	msg.Rtype = a.DeskData.Rtype
	msg.Gtype = a.DeskData.Gtype
	msg.Unique = a.DeskData.Unique
	nodePid.Tell(msg)
	//TODO 优化
	if a.roomPid != nil {
		a.roomPid.Request(msg, ctx.Self())
	}
	//停掉服务
	ctx.Self().Stop()
}
