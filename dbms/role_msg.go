package main

import (
	"time"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//Handler 消息处理
func (a *RoleActor) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.Connected:
		arg := msg.(*pb.Connected)
		glog.Infof("Connected %s", arg.Name)
	case *pb.Disconnected:
		arg := msg.(*pb.Disconnected)
		glog.Infof("Disconnected %s", arg.Name)
	case *pb.ServeStop:
		//关闭服务
		a.handlerStop(ctx)
		//响应登录
		rsp := new(pb.ServeStoped)
		ctx.Respond(rsp)
	case *pb.ServeStart:
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
func (a *RoleActor) start(ctx actor.Context) {
	glog.Infof("role start: %v", ctx.Self().String())
	////注册更新机器人
	//phone := cfg.Section("robot").Key("phone").Value()
	//passwd := cfg.Section("robot").Key("passwd").Value()
	////data.RegistRobots(phone, passwd, a.uniqueid)
	////新机器人
	//head := cfg.Section("domain").Key("headimag").Value()
	////data.RegistRobots2(head, phone, passwd, a.uniqueid)
	//data.RegistRobots3(head, phone, passwd, a.uniqueid)
	//启动
	go a.ticker(ctx)
}

//时钟
func (a *RoleActor) ticker(ctx actor.Context) {
	tick := time.Tick(30 * time.Second)
	msg := new(pb.Tick)
	for {
		select {
		case <-a.stopCh:
			glog.Info("role ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-a.stopCh:
			glog.Info("role ticker closed")
			return
		case <-tick:
			ctx.Self().Tell(msg)
		}
	}
}

//钟声
func (a *RoleActor) ding(ctx actor.Context) {
	//glog.Debugf("ding: %v", ctx.Self().String())
	//glog.Debugf("timer: %d", a.timer)
	switch a.timer {
	case 4: //2分钟
		a.saveUser()
		a.timer = 0
		a.smsExpire()
		a.touristIP()
	case 2: //1分钟
		a.smsExpire()
		a.timer++
	default:
		a.timer++
	}
}

//关闭时钟
func (a *RoleActor) closeTick() {
	select {
	case <-a.stopCh:
		return
	default:
		//停止发送消息
		close(a.stopCh)
	}
}

func (a *RoleActor) handlerStop(ctx actor.Context) {
	glog.Debugf("handlerStop: %s", a.Name)
	//关闭
	a.closeTick()
	//回存数据
	if a.uniqueid != nil {
		a.uniqueid.Save()
	}
	for k, v := range a.offline {
		glog.Debugf("Stop offline: %s", k)
		v.Save()
	}
	for k, v := range a.online {
		glog.Debugf("Stop online: %s", k)
		v.Save()
	}
}
