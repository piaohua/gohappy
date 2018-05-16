package main

import (
	"gohappy/data"
	"gohappy/glog"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//RoleActor 玩家角色进程
type RoleActor struct {
	stopCh chan struct{} // 关闭通道

	wsPid     *actor.PID // ws进程ID
	pid       *actor.PID // rs进程ID
	dbmsPid   *actor.PID // 数据中心
	roomPid   *actor.PID // 房间列表
	rolePid   *actor.PID // 角色服务
	loggerPid *actor.PID // 角色服务
	gamePid   *actor.PID // 游戏逻辑

	*data.User //玩家在线数据

	online bool //在线状态
	status bool //更新状态
	timer  int  //计时
}

//NewRole 启动一个新玩家角色
func NewRole() *RoleActor {
	return &RoleActor{
		stopCh: make(chan struct{}),
	}
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (rs *RoleActor) Receive(ctx actor.Context) {
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
	case proto.Message:
		rs.Handler(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//初始化
func (rs *RoleActor) initRs() *actor.PID {
	props := actor.FromProducer(func() actor.Actor { return rs }) //实例
	return actor.Spawn(props)                                     //启动一个进程
}
