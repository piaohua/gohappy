package main

import (
	"time"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//Handler 消息处理
func (a *GateActor) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.Connected:
		//连接成功
		arg := msg.(*pb.Connected)
		glog.Infof("Connected %s", arg.Name)
	case *pb.Disconnected:
		//成功断开
		arg := msg.(*pb.Disconnected)
		glog.Infof("Disconnected %s", arg.Name)
	case *pb.ServeStop:
		//关闭服务
		a.handlerStop(ctx)
		//响应登录
		rsp := new(pb.ServeStoped)
		ctx.Respond(rsp)
	case *pb.ServeStart:
		//初始化建立连接
		a.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.Tick:
		a.ding(ctx)
	case proto.Message:
		a.handlerUser(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//启动服务
func (a *GateActor) start(ctx actor.Context) {
	glog.Infof("gate start: %v", ctx.Self().String())
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
	connect := &pb.Connect{
		Name: a.Name,
	}
	a.dbmsPid.Request(connect, ctx.Self())
	glog.Infof("a.dbmsPid: %s", a.dbmsPid.String())
	glog.Infof("a.roomPid: %s", a.roomPid.String())
	glog.Infof("a.rolePid: %s", a.rolePid.String())
	glog.Infof("a.loggerPid: %s", a.loggerPid.String())
	//启动
	go a.ticker(ctx)
}

//时钟
func (a *GateActor) ticker(ctx actor.Context) {
	tick := time.Tick(30 * time.Second)
	msg := new(pb.Tick)
	for {
		select {
		case <-a.stopCh:
			glog.Info("gate ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-a.stopCh:
			glog.Info("gate ticker closed")
			return
		case <-tick:
			ctx.Self().Tell(msg)
		}
	}
}

//钟声
func (a *GateActor) ding(ctx actor.Context) {
	//glog.Debugf("ding: %v", ctx.Self().String())
	//下线离线玩家
	a.offlineStop(ctx)
}

//关闭时钟
func (a *GateActor) closeTick() {
	select {
	case <-a.stopCh:
		return
	default:
		//停止发送消息
		close(a.stopCh)
	}
}

func (a *GateActor) handlerStop(ctx actor.Context) {
	glog.Debugf("handlerStop: %s", a.Name)
	//关闭
	a.closeTick()
	//关闭消息
	msg1 := new(pb.OfflineStop)
	for k, v := range a.offline {
		glog.Debugf("Stop offline role: %s", k)
		v.Tell(msg1)
	}
	//关闭消息
	msg2 := new(pb.ServeClose)
	for k, v := range a.online {
		glog.Debugf("Stop role: %s", k)
		v.Tell(msg1)
		v.Tell(msg2)
	}
	//延迟
	<-time.After(3 * time.Second)
	//断开处理
	msg := &pb.Disconnect{
		Name: a.Name,
	}
	if a.dbmsPid != nil {
		a.dbmsPid.Request(msg, ctx.Self())
	}
	//延迟
	<-time.After(2 * time.Second)
}
