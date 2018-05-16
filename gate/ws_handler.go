package main

import (
	"time"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//Handler 消息处理
func (ws *WSConn) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.SetLogined:
		//设置连接,还未登录
		arg := msg.(*pb.SetLogined)
		glog.Debugf("ws SetLogined %#v", arg)
		ws.rolePid = arg.RolePid
	case *pb.ServeClose:
		glog.Debugf("ws ServeClose %#v", msg)
		//断开连接
		ws.stop(ctx)
	case *pb.ServeStop:
		glog.Debugf("ws ServeStop %#v", msg)
		//断开连接
		ws.stop(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStoped:
	case *pb.ServeStart:
		glog.Debugf("ws ServeStart %#v", msg)
		ws.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.ServeStarted:
	case proto.Message:
		//响应消息
		//ws.Send(msg)
		ws.handlerLogin(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

func (ws *WSConn) start(ctx actor.Context) {
	glog.Infof("ws start: %v", ctx.Self().String())
	ctx.SetReceiveTimeout(waitForLogin) //login timeout set
	msg2 := &pb.SetLogin{
		Sender: ctx.Self(),
	}
	//nodePid.Tell(msg2)
	timeout := 3 * time.Second
	res1, err1 := nodePid.RequestFuture(msg2, timeout).Result()
	if err1 != nil {
		glog.Errorf("reqRole err: %v", err1)
		return
	}
	if response1, ok := res1.(*pb.SetLogined); ok {
		ws.rolePid = response1.RolePid
	}
}

func (ws *WSConn) stop(ctx actor.Context) {
	glog.Infof("ws stop: %v", ctx.Self().String())
	//断开连接
	ws.Close()
	//表示已经断开
	ws.online = false
	ctx.Self().Stop()
}
