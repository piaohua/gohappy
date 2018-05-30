package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/mailbox"
	"github.com/AsynkronIT/protoactor-go/router"
	"github.com/gogo/protobuf/proto"
)

var (
	loggerPid *actor.PID
)

const maxConcurrency = 5

//LoggerActor 日志记录服务
type LoggerActor struct {
	Name string
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (a *LoggerActor) Receive(ctx actor.Context) {
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

func newLoggerActor() actor.Actor {
	a := new(LoggerActor)
	//name
	a.Name = cfg.Section("logger").Name()
	return a
}

//func NewLogger() *actor.PID {
//	props := actor.FromProducer(newLoggerActor)
//	pid := actor.Spawn(props)
//	return pid
//}

//NewLoggerProps 启动
func NewLoggerProps() *actor.Props {
	return router.NewRoundRobinPool(maxConcurrency).
		WithProducer(newLoggerActor).
		WithMailbox(mailbox.Unbounded())
}
