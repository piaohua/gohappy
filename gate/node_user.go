package main

import (
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
	"gohappy/data"
	"utils"

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
		a.pushNotice(arg)
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

//节点广播消息
func (a *GateActor) broadcast(msg interface{}) {
	for _, v := range a.online {
			v.Tell(msg)
	}
}

//节点广播消息推送
func (a *GateActor) pushNotice(arg *pb.SyncConfig) {
	switch arg.Type {
	case pb.CONFIG_NOTICE: //公告
		b := make(map[string]data.Notice)
		err = json.Unmarshal(arg.Data, &b)
		if err != nil {
			glog.Errorf("syncConfig Unmarshal err %v, data %#v", err, arg.Data)
			return
		}
		glog.Debugf("pushNotice %#v", b)
		for _, v := range b {
			switch arg.Atype {
			case pb.CONFIG_DELETE:
			case pb.CONFIG_UPSERT:
				msg := new(pb.SPushNotice)
				msg.Info = &pb.Notice{
					Time:    utils.Time2LocalStr(v.Ctime),
					Rtype:   int32(v.Rtype),
					Acttype: int32(v.Acttype),
					Content: v.Content,
				}
				if v.Userid == "" {
					//广播消息通知玩家,只发送新消息
					a.broadcast(msg)
				} else {
					s := utils.Split(v.Userid, ",")
					for _, val := range s {
						//玩家个人消息单独通知
						if val, ok := a.online[val]; ok {
							val.Tell(msg)
						}
					}
				}
			}
		}
	}
}
