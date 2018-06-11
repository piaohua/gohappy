package main

import (
	"gohappy/data"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

func (a *Desk) handlerLogic(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CloseDesk:
		arg := msg.(*pb.CloseDesk)
		glog.Debugf("CloseDesk %#v", arg)
		a.closeDesk(arg, ctx)
	case *pb.LeaveDesk:
		arg := msg.(*pb.LeaveDesk)
		glog.Debugf("LeaveDesk %#v", arg)
		a.leaveDesk(arg, ctx)
	case *pb.SyncConfig:
		//更新配置
		arg := msg.(*pb.SyncConfig)
		glog.Debugf("SyncConfig %#v", arg)
		a.syncConfig(arg, ctx)
	case *pb.PrintDesk:
		//打印牌局状态信息,test
		a.printOver()
	case *pb.EnterDesk:
		arg := msg.(*pb.EnterDesk)
		glog.Debugf("EnterDesk %#v", arg)
		a.enterDesk(arg, ctx)
	case *pb.OfflineDesk:
		arg := msg.(*pb.OfflineDesk)
		glog.Debugf("OfflineDesk %#v", arg)
		//离线消息
		a.offlineDesk(arg.Userid)
	case *pb.ChangeCurrency:
		arg := msg.(*pb.ChangeCurrency)
		//充值或购买同步
		a.changeCurrency(arg)
	case proto.Message:
		//请求消息
		a.handlerRequest(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//关闭桌子
func (a *Desk) closeDesk(arg *pb.CloseDesk, ctx actor.Context) {
	//TODO
	//响应
	//rsp := new(pb.ClosedDesk)
	//ctx.Respond(rsp)
}

//'离开房间

//掉线
func (a *Desk) offlineDesk(userid string) {
	a.setOffline(userid, true)
	pid := a.getPid(userid)
	if pid != nil {
		delete(a.router, pid.String())
	}
	//离线消息
	a.offlineMsg(userid)
}

//玩家掉线离开桌子
func (a *Desk) leaveDesk(arg *pb.LeaveDesk, ctx actor.Context) {
	//TODO 同一个房间?
	//if arg.Roomid != a.DeskData.Rid {
	//}
	//离线
	defer a.offlineDesk(arg.Userid)
	//响应消息
	rsp := new(pb.LeftDesk)
	if _, ok := a.roles[arg.Userid]; ok {
		errcode := a.leave(arg.Userid)
		rsp.Error = errcode
	}
	ctx.Respond(rsp)
	//成功离开后移除
	if rsp.Error == pb.OK {
		//TODO 私人房间掉线是否移除
		a.userLeaveDesk(arg.Userid)
	}
}

//.

//'进入房间
func (a *Desk) enterDesk(arg *pb.EnterDesk, ctx actor.Context) {
	rsp := new(pb.EnteredDesk)
	user := new(data.User)
	err2 := json.Unmarshal(arg.Data, user)
	if err2 != nil {
		glog.Errorf("user Unmarshal err %v", err2)
		rsp.Error = pb.RoomNotExist
		//ctx.Respond(rsp)
		arg.Sender.Tell(rsp)
		return
	}
	//加入桌子
	errcode := a.enter(user, arg.Sender)
	if errcode != pb.OK && errcode != pb.AlreadyInRoom {
		glog.Errorf("entry Desk err: %d", errcode)
		rsp.Error = errcode
		//ctx.Respond(rsp)
		arg.Sender.Tell(rsp)
		return
	}
	if errcode == pb.OK {
		//加入房间消耗
		a.enterDeskCost(user.GetUserid())
	}
	//响应消息
	rsp.Roomid = a.DeskData.Rid
	rsp.Rtype = a.DeskData.Rtype
	rsp.Gtype = a.DeskData.Gtype
	rsp.Userid = user.GetUserid()
	rsp.Desk = ctx.Self()
	//ctx.Respond(rsp)
	arg.Sender.Tell(rsp)
	//TODO 优化
	if errcode == pb.AlreadyInRoom {
		return
	}
	//进入消息
	msg3 := new(pb.JoinDesk)
	msg3.Roomid = a.DeskData.Rid
	msg3.Rtype = a.DeskData.Rtype
	msg3.Gtype = a.DeskData.Gtype
	msg3.Userid = user.Userid
	msg3.Sender = arg.Sender
	nodePid.Request(msg3, ctx.Self())
	a.roomPid.Request(msg3, ctx.Self())
}

//AA消耗房间
func (a *Desk) isAADesk() bool {
	switch a.DeskData.Rtype {
	case int32(pb.ROOM_TYPE1): //私人
		return a.DeskData.Payment == 1 //AA
	}
	return false
}

//加入房间消耗
func (a *Desk) enterDeskCost(userid string) {
	if !a.isAADesk() || userid == a.DeskData.Cid {
		return
	}
	msg2 := handler.ChangeCurrencyMsg(int64(a.DeskData.Cost),
		0, 0, 0, int32(pb.LOG_TYPE2), userid)
	a.send2userid(userid, msg2)
}

//.

//'同步配置
func (a *Desk) syncConfig(arg *pb.SyncConfig, ctx actor.Context) {
	b := make(map[string]data.Game)
	err = json.Unmarshal(arg.Data, &b)
	if err != nil {
		glog.Errorf("syncConfig Unmarshal err %v", err)
		return
	}
	for _, v := range b {
		if a.DeskData.Unique == v.Id {
			//TODO 只更新可变内容
			//a.DeskData.Ante = v.Ante
			//a.DeskData.Chip = v.Chip
			deskData := handler.NewDeskData(&v)
			deskData.Rid = a.DeskData.Rid
			a.DeskData = deskData
			return
		}
	}
}

//.

//更新货币
func (a *Desk) changeCurrency(arg *pb.ChangeCurrency) {
	user := a.getPlayer(arg.Userid)
	if user == nil {
		glog.Debugf("changeCurrency err %s", arg.Userid)
		return
	}
	user.AddCurrency(arg.Diamond, arg.Coin, arg.Card, arg.Chip)
}

// vim: set foldmethod=marker foldmarker=//',//.:
