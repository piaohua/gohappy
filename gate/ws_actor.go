package main

import (
	"gohappy/glog"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//test
type testWs struct{ Who string }

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (ws *WSConn) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
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
		//登录超时
		ws.Close()
	case *testWs:
		glog.Infof("self %s\n", ctx.Self().String())
		glog.Infof("msg.Who %v\n", msg.Who)
	case proto.Message:
		ws.Handler(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//初始化
func (ws *WSConn) initWs() *actor.PID {
	props := actor.FromProducer(func() actor.Actor { return ws }) //实例
	return actor.Spawn(props)                                     //启动一个进程
}
