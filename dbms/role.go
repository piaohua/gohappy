package main

import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

var (
	rolePid *actor.PID
)

//RoleActor 角色列表服务
type RoleActor struct {
	Name string
	//角色数据
	online map[string]*data.User
	//离线数据
	offline map[string]*data.User
	//角色userid - role,所在节点
	roles map[string]*data.Role
	//账号映射account-userid
	players map[string]string
	//缓存时间userid-timer
	caches map[string]int
	//更新状态userid-bool
	states map[string]bool
	//验证码有效期code - timer
	smstime map[string]int64
	//验证码code - phone
	smscode map[string]string
	//验证码phone - code
	smsphone map[string]string
	//游客注册ip - count
	tourist map[string]int
	//唯一id生成
	uniqueid *data.IDGen
	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int
}

// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
func (a *RoleActor) Receive(ctx actor.Context) {
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

func newRoleActor() actor.Actor {
	a := new(RoleActor)
	a.Name = cfg.Section("role").Name()
	glog.Debugf("a.Name %s", a.Name)
	a.online = make(map[string]*data.User)
	a.offline = make(map[string]*data.User)
	a.roles = make(map[string]*data.Role)
	a.players = make(map[string]string)
	a.caches = make(map[string]int)
	a.states = make(map[string]bool)
	a.smstime = make(map[string]int64)
	a.smscode = make(map[string]string)
	a.smsphone = make(map[string]string)
	a.tourist = make(map[string]int)
	//唯一id初始化
	a.uniqueid = data.InitIDGen(data.USERID_KEY)
	glog.Debugf("uniqueid %#v", a.uniqueid)
	a.stopCh = make(chan struct{})
	return a
}
