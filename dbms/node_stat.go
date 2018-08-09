package main

import (
	"time"

	"gohappy/game/config"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//StatActor 统计服务
type StatActor struct {
	Name string
	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (state *StatActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		glog.Notice("Started, initialize actor here")
	case *actor.Stopping:
		glog.Notice("Stopping, actor is about shut down")
	case *actor.Stopped:
		glog.Notice("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		glog.Notice("Restarting, actor is about restart")
	case *pb.ServeStart:
		state.start(context)
	case *pb.ServeStop:
		state.stop(context)
	case *pb.Tick:
		state.ding(context)
	case *pb.AgentActivity:
		glog.Debugf("AgentActivity %#v", msg)
		list, err := handler.StatActivity(msg)
		if err != nil {
			glog.Errorf("err %v", err)
			return
		}
		for _, v := range list {
			rolePid.Tell(v) //send to role
		}
		msg.Page++
		context.Self().Tell(msg)
	case *pb.AgentActivityStat:
		glog.Debugf("AgentActivityStat %#v", msg)
		state.stat(msg.GetType(), context)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//启动服务
func (state *StatActor) start(context actor.Context) {
	glog.Infof("stat start: %v", context.Self().String())
	//启动
	go state.ticker(context)
}

func (state *StatActor) stop(context actor.Context) {
	glog.Infof("stat stop: %v", context.Self().String())
	//关闭
	state.closeTick()
}

//时钟
func (state *StatActor) ticker(context actor.Context) {
	tick := time.Tick(1 * time.Minute)
	msg := new(pb.Tick)
	for {
		select {
		case <-state.stopCh:
			glog.Info("dbms ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-state.stopCh:
			glog.Info("dbms ticker closed")
			return
		case <-tick:
			context.Self().Tell(msg)
		}
	}
}

//钟声
func (state *StatActor) ding(context actor.Context) {
	//glog.Debugf("ding: %v", context.Self().String())
	switch utils.Hour() {
	case 0:
	default:
		return
	}
	switch utils.Minute() {
	case 5:
		state.stat(int32(pb.ACT_TYPE0), context)
	case 10:
		state.stat(int32(pb.ACT_TYPE1), context)
	case 15:
		state.stat(int32(pb.ACT_TYPE2), context)
	}
}

func (state *StatActor) stat(Type int32, context actor.Context) {
	list := config.GetActivitys()
	for _, v := range list {
		if v.Type != Type {
			continue
		}
		msg := &pb.AgentActivity{
			Actid: v.Id,
			Page:  1,
		}
		context.Self().Tell(msg)

	}
}

//关闭时钟
func (state *StatActor) closeTick() {
	select {
	case <-state.stopCh:
		return
	default:
		//停止发送消息
		close(state.stopCh)
	}
}

func newStatActor() actor.Actor {
	state := new(StatActor)
	state.stopCh = make(chan struct{})
	return state
}

//NewStat 启动服务
func NewStat() *actor.PID {
	props := actor.FromProducer(newStatActor)
	pid := actor.Spawn(props)
	return pid
}

func statStart(pid *actor.PID) {
	msg := new(pb.ServeStart)
	pid.Tell(msg)
}

func statStop(pid *actor.PID) {
	msg := new(pb.ServeStop)
	pid.Tell(msg)
}
