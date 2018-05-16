package main

import (
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

func (a *GateActor) handlerUser(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.SyncConfig:
		//同步配置
		arg := msg.(*pb.SyncConfig)
		glog.Debugf("SyncConfig %#v", arg)
		handler.SyncConfig(arg)
	case *pb.PayCurrency:
		//后台或充值同步到game房间
		arg := msg.(*pb.PayCurrency)
		glog.Debugf("PayCurrency %#v", arg)
		userid := arg.Userid
		if v, ok := a.online[userid]; ok {
			v.Tell(arg)
		} else if v, ok := a.offline[userid]; ok {
			v.Tell(arg)
		} else {
			//离线
			a.rolePid.Tell(arg)
		}
	case *pb.ChangeCurrency:
		arg := msg.(*pb.ChangeCurrency)
		userid := arg.Userid
		if v, ok := a.online[userid]; ok {
			glog.Infof("ChangeCurrency %#v", arg)
			v.Tell(msg)
		} else if v, ok := a.offline[userid]; ok {
			glog.Infof("ChangeCurrency %#v", arg)
			v.Tell(msg)
		} else {
			glog.Infof("ChangeCurrency %#v", arg)
			//离线
			a.rolePid.Tell(msg)
		}
	case *pb.WxpayCallback:
		arg := msg.(*pb.WxpayCallback)
		glog.Debugf("WxpayCallback %#v", arg)
		if !handler.WxpayVerify(arg) {
			return
		}
		a.rolePid.Tell(arg)
	case *pb.WxpayGoods:
		arg := msg.(*pb.WxpayGoods)
		glog.Debugf("WxpayGoods: %v", arg)
		userid := arg.Userid
		if v, ok := a.online[userid]; ok {
			v.Tell(arg)
		} else if v, ok := a.offline[userid]; ok {
			v.Tell(msg)
		} else {
			glog.Errorf("WxpayGoods: %v", arg)
		}
	case proto.Message:
		a.handlerLogin(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}
