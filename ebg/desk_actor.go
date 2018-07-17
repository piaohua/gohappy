package main

import (
	"gohappy/glog"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (a *Desk) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		glog.Notice("Starting, initialize actor here")
	case *actor.Stopping:
		glog.Notice("Stopping, actor is about to shut down")
	case *actor.Stopped:
		glog.Notice("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		glog.Notice("Restarting, actor is about to restart")
	case proto.Message:
		a.Handler(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//初始化
func (a *Desk) newDesk() *actor.PID {
	props := actor.FromProducer(func() actor.Actor { return a }) //实例
	return actor.Spawn(props)                                    //启动一个进程
}
