package main

import (
	"time"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/gogo/protobuf/proto"
)

var (
	nodePid *actor.PID
)

//DBMSActor 数据库操作服务
type DBMSActor struct {
	Name string
	//服务注册
	serve map[string]*actor.PID
	//统计服务
	statPid *actor.PID
	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (a *DBMSActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *pb.Request:
		ctx.Respond(&pb.Response{})
	case *actor.Started:
		glog.Notice("Starting, initialize actor here")
	case *actor.Stopping:
		glog.Notice("Stopping, actor is about to shut down")
	case *actor.Stopped:
		glog.Notice("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		glog.Notice("Restarting, actor is about to restart")
	case *actor.ReceiveTimeout:
		glog.Infof("ReceiveTimeout: %v", ctx.Self().String())
	case proto.Message:
		a.Handler(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

func newDBMSActor() actor.Actor {
	a := new(DBMSActor)
	a.Name = cfg.Section("dbms").Name()
	glog.Debugf("a.Name %s", a.Name)
	a.serve = make(map[string]*actor.PID)
	a.stopCh = make(chan struct{})
	return a
}

//NewRemote 启动服务
func NewRemote(bind, kind, room, role, logger string) {
	remote.Start(bind)
	//
	msg := new(pb.ServeStart)
	//remote.Register(kind, actor.FromProducer(newDBMSActor))
	dbmsProps := actor.FromProducer(newDBMSActor)
	remote.Register(kind, dbmsProps)
	nodePid, err = actor.SpawnNamed(dbmsProps, kind)
	if err != nil {
		glog.Fatalf("nodePid err %v", err)
	}
	glog.Infof("nodePid %s", nodePid.String())
	nodePid.Tell(msg)
	//
	//remote.Register(room, actor.FromProducer(newRoomActor))
	roomProps := actor.FromProducer(newRoomActor)
	remote.Register(room, roomProps)
	roomPid, err = actor.SpawnNamed(roomProps, room)
	if err != nil {
		glog.Fatalf("roomPid err %v", err)
	}
	glog.Infof("roomPid %s", roomPid.String())
	roomPid.Tell(msg)
	//
	//remote.Register(role, actor.FromProducer(newRoleActor))
	roleProps := actor.FromProducer(newRoleActor)
	remote.Register(role, roleProps)
	rolePid, err = actor.SpawnNamed(roleProps, role)
	if err != nil {
		glog.Fatalf("rolePid err %v", err)
	}
	glog.Infof("rolePid %s", rolePid.String())
	rolePid.Tell(msg)
	//
	//remote.Register(logger, actor.FromProducer(newLoggerActor))
	loggerProps := NewLoggerProps()
	remote.Register(logger, loggerProps)
	loggerPid, err = actor.SpawnNamed(loggerProps, logger)
	if err != nil {
		glog.Fatalf("loggerPid err %v", err)
	}
	glog.Infof("loggerPid %s", loggerPid.String())
}

//Stop 关闭服务
func Stop() {
	timeout := 5 * time.Second
	msg := new(pb.ServeStop)
	if rolePid != nil {
		res1, err1 := rolePid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("rolePid Stop err: %v", err1)
		}
		if response1, ok := res1.(*pb.ServeStoped); ok {
			glog.Debugf("response1: %#v", response1)
		}
		rolePid.Stop()
	}
	if roomPid != nil {
		res1, err1 := roomPid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("roomPid Stop err: %v", err1)
		}
		if response1, ok := res1.(*pb.ServeStoped); ok {
			glog.Debugf("response1: %#v", response1)
		}
		roomPid.Stop()
	}
	if loggerPid != nil {
		res1, err1 := loggerPid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("loggerPid Stop err: %v", err1)
		}
		if response1, ok := res1.(*pb.ServeStoped); ok {
			glog.Debugf("response1: %#v", response1)
		}
		loggerPid.Stop()
	}
	if nodePid != nil {
		res1, err1 := nodePid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("nodePid Stop err: %v", err1)
		}
		if response1, ok := res1.(*pb.ServeStoped); ok {
			glog.Debugf("response1: %#v", response1)
		}
		nodePid.Stop()
	}
}
