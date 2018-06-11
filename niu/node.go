package main

import (
	"time"

	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/gogo/protobuf/proto"
)

var (
	nodePid *actor.PID
)

//DeskActor 桌子服务
type DeskActor struct {
	Name string
	//数据中心
	dbmsPid *actor.PID
	//房间服务
	roomPid *actor.PID
	//角色服务
	rolePid *actor.PID
	//日志服务
	loggerPid *actor.PID

	//所有桌子roomid-deskPid
	desks map[string]*data.DeskBase
	//配置映射unique-roomid
	rules map[string]string

	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int64
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (a *DeskActor) Receive(ctx actor.Context) {
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

func newDeskActor() actor.Actor {
	a := new(DeskActor)
	a.Name = cfg.Section(nodeName).Name()
	a.desks = make(map[string]*data.DeskBase)
	a.rules = make(map[string]string)
	a.stopCh = make(chan struct{})
	return a
}

//NewRemote 启动服务
func NewRemote(bind, name string) {
	remote.Start(bind)
	props := actor.FromProducer(newDeskActor)
	remote.Register(name, props)
	nodePid, err = actor.SpawnNamed(props, name)
	if err != nil {
		glog.Fatalf("nodePid err %v", err)
	}
	glog.Infof("nodePid %s", nodePid.String())
	nodePid.Tell(new(pb.ServeStart))
}

//Stop 关闭服务
func Stop() {
	timeout := 10 * time.Second
	msg := new(pb.ServeStop)
	if nodePid != nil {
		res1, err1 := nodePid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			//TODO future: timeout
			glog.Errorf("nodePid Stop err: %v", err1)
		}
		if response1, ok := res1.(*pb.ServeStoped); ok {
			glog.Debugf("response1: %#v", response1)
		}
		nodePid.Stop()
	}
}
