package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//请求处理
func (a *RoomActor) handlerDesk(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	//case *pb.GenDesk:
	//	arg := msg.(*pb.GenDesk)
	//	glog.Debugf("GenDesk: %v", arg)
	//	a.genDesk(arg, ctx)
	case *pb.AddDesk:
		arg := msg.(*pb.AddDesk)
		glog.Debugf("AddDesk: %v", arg)
		a.addDesk(arg, ctx)
	case *pb.JoinDesk:
		arg := msg.(*pb.JoinDesk)
		glog.Debugf("JoinDesk %#v", arg)
		a.joinDesk(arg, ctx)
	case *pb.LeaveDesk:
		arg := msg.(*pb.LeaveDesk)
		glog.Debugf("LeaveDesk %#v", arg)
		a.leaveDesk(arg, ctx)
	case *pb.Logout:
		arg := msg.(*pb.Logout)
		glog.Debugf("Logout %#v", arg)
		//TODO 暂时不处理
	case *pb.CloseDesk:
		arg := msg.(*pb.CloseDesk)
		glog.Debugf("CloseDesk %#v", arg)
		a.closeDesk(arg, ctx)
	case *pb.MatchDesk:
		arg := msg.(*pb.MatchDesk)
		glog.Debugf("MatchDesk %#v", arg)
		a.matchDesk(arg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}
