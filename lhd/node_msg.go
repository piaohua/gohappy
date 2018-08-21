package main

import (
	"time"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//Handler 消息处理
func (a *DeskActor) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.ServeStop:
		//关闭服务
		a.handlerStop(ctx)
		//响应登录
		rsp := new(pb.ServeStoped)
		ctx.Respond(rsp)
	case *pb.ServeStart:
		a.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.Tick:
		a.ding(ctx)
	default:
		//glog.Errorf("unknown message %v", msg)
		a.handlerMsg(msg, ctx)
	}
}

//启动服务
func (a *DeskActor) start(ctx actor.Context) {
	glog.Infof("desk start: %v", ctx.Self().String())
	//dbms
	bind := cfg.Section("dbms").Key("bind").Value()
	kind := cfg.Section("dbms").Key("kind").Value()
	room := cfg.Section("dbms").Key("room").Value()
	role := cfg.Section("dbms").Key("role").Value()
	logger := cfg.Section("dbms").Key("logger").Value()
	a.dbmsPid = actor.NewPID(bind, kind)
	a.roomPid = actor.NewPID(bind, room)
	a.rolePid = actor.NewPID(bind, role)
	a.loggerPid = actor.NewPID(bind, logger)
	glog.Infof("a.dbmsPid: %s", a.dbmsPid.String())
	glog.Infof("a.roomPid: %s", a.roomPid.String())
	glog.Infof("a.rolePid: %s", a.rolePid.String())
	glog.Infof("a.loggerPid: %s", a.loggerPid.String())
	//连接
	connect := &pb.Connect{
		Name: a.Name,
	}
	a.dbmsPid.Request(connect, ctx.Self())
	//主动同步配置
	//msg2 := new(pb.GetConfig)
	//a.dbmsPid.Request(msg2, ctx.Self())
	//启动
	//go a.ticker(ctx)
	//TODO test
	gameData := handler.NewFreeGameData(a.Name, int32(pb.LHD))
	if deskPid, ok := a.spawnDesk(gameData, ctx); ok {
		//deskPid.Tell(arg)
		glog.Infof("deskPid: %s", deskPid.String())
		return
	}
}

//时钟
func (a *DeskActor) ticker(ctx actor.Context) {
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
func (a *DeskActor) ding(ctx actor.Context) {
	//glog.Debugf("ding: %v", ctx.Self().String())
	//TODO 优化超时造成的延迟,单独routine处理
	//a.handing(ctx)
}

//关闭时钟
func (a *DeskActor) closeTick() {
	select {
	case <-a.stopCh:
		return
	default:
		//停止发送消息
		close(a.stopCh)
	}
}

func (a *DeskActor) handlerStop(ctx actor.Context) {
	glog.Debugf("handlerStop: %s", a.Name)
	//关闭
	a.closeTick()
	//关闭消息
	for k := range a.desks {
		glog.Infof("Stop desk: %s", k)
		a.stopDesk(k, ctx)
	}
	//延迟
	<-time.After(5 * time.Second)
	//断开处理
	msg := &pb.Disconnect{
		Name: a.Name,
	}
	if a.dbmsPid != nil {
		a.dbmsPid.Tell(msg)
	}
	//延迟
	<-time.After(3 * time.Second)
}
