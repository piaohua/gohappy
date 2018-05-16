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

//LoginActor 登录服务
type LoginActor struct {
	Name string
	//中心服务
	dbmsPid *actor.PID
	//角色服务
	rolePid *actor.PID
	//角色上次登录节点
	roles map[string]string
	//验证码限制phone - times
	smsphone map[string]int
	//验证码限制ip - times (表示1分钟)
	smstimes map[string]int
	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (a *LoginActor) Receive(ctx actor.Context) {
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

func newLoginActor() actor.Actor {
	a := new(LoginActor)
	a.Name = cfg.Section("login").Name()
	//roles key=userid - "1"
	a.roles = make(map[string]string)
	a.smsphone = make(map[string]int)
	a.smstimes = make(map[string]int)
	a.stopCh = make(chan struct{})
	return a
}

//NewRemote 启动服务
func NewRemote(bind, kind string) {
	remote.Start(bind)
	loginProps := actor.FromProducer(newLoginActor)
	remote.Register(kind, loginProps)
	nodePid, err = actor.SpawnNamed(loginProps, kind)
	if err != nil {
		glog.Fatalf("nodePid err %v", err)
	}
	glog.Infof("nodePid %s", nodePid.String())
	nodePid.Tell(new(pb.ServeStart))
}

//Stop 关闭服务
func Stop() {
	timeout := 5 * time.Second
	msg := new(pb.ServeStop)
	if nodePid != nil {
		res1, err1 := nodePid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("nodePid Stop err: %v", err1)
		}
		response1 := res1.(*pb.ServeStoped)
		glog.Debugf("response1: %#v", response1)
		nodePid.Stop()
	}
}
