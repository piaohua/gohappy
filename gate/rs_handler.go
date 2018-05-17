package main

import (
	"time"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//Handler 消息处理
func (rs *RoleActor) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.OfflineStop:
		glog.Debugf("rs OfflineStop %s", ctx.Self().String())
		//断开连接
		rs.CloseWs()
		//停止,TODO 暂时使用离线方法
		rs.loginElse()
		//关闭
		rs.StopRs()
	case *pb.ServeClose:
		glog.Debugf("rs ServeClose %s", ctx.Self().String())
		arg := new(pb.SLoginOut)
		arg.Rtype = int32(pb.LOGOUT_TYPE2) //停服
		rs.Send(arg)
		//断开连接
		rs.CloseWs()
		//停止,TODO 暂时使用离线方法
		rs.loginElse()
		//关闭
		rs.StopRs()
	case *pb.ServeStop:
		glog.Debugf("rs ServeStop %s", ctx.Self().String())
		rs.stop(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStoped:
	case *pb.ServeStart:
		arg := msg.(*pb.ServeStart)
		glog.Debugf("rs ServeStart %s, %#v", ctx.Self().String(), arg)
		//启动时钟
		go rs.ticker(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStarted:
	case *pb.Tick:
		rs.ding(ctx)
	case proto.Message:
		//响应消息
		//rs.Send(msg)
		rs.handlerLogin(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//时钟
func (rs *RoleActor) ticker(ctx actor.Context) {
	tick := time.Tick(time.Second)
	msg := new(pb.Tick)
	for {
		select {
		case <-rs.stopCh:
			glog.Info("rs ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-rs.stopCh:
			glog.Info("rs ticker closed")
			return
		case <-tick:
			ctx.Self().Tell(msg)
		}
	}
}

//30秒同步一次
func (rs *RoleActor) ding(ctx actor.Context) {
	rs.timer++
	if rs.timer != 30 {
		return
	}
	rs.timer = 0
	if !rs.online {
		return
	}
	//同步数据
	rs.syncUser()
}

//断线
func (rs *RoleActor) stop(ctx actor.Context) {
	glog.Infof("rs stop: %v", ctx.Self().String())
	//已经断开,在别处登录
	if !rs.online {
		return
	}
	//关闭连接
	rs.CloseWs()
	//离开消息
	rs.leaveDesk()
	//回存数据
	rs.syncUser()
	//登出日志
	msg2 := &pb.LogLogout{
		Userid: rs.User.Userid,
		Event:  int32(pb.LOGOUT_TYPE1), //正常断开
	}
	rs.loggerPid.Tell(msg2)
	//断开处理
	msg := &pb.Logout{
		Sender: ctx.Self(),
		Userid: rs.User.Userid,
	}
	nodePid.Tell(msg)
	//表示已经断开
	rs.online = false
}
