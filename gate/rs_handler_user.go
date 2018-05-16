package main

import (
	"gohappy/data"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家数据请求处理
func (rs *RoleActor) handlerUser(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CPing:
		arg := msg.(*pb.CPing)
		//glog.Debugf("CPing %#v", arg)
		rsp := handler.Ping(arg)
		rs.Send(rsp)
	case *pb.CNotice:
		arg := msg.(*pb.CNotice)
		glog.Debugf("CNotice %#v", arg)
		rsp := handler.GetNotice(data.NOTICE_TYPE1)
		rs.Send(rsp)
	case *pb.CGetCurrency:
		arg := msg.(*pb.CGetCurrency)
		glog.Debugf("CGetCurrency %#v", arg)
		//响应
		rsp := handler.GetCurrency(arg, rs.User)
		rs.Send(rsp)
	case *pb.CBuy:
		arg := msg.(*pb.CBuy)
		glog.Debugf("CBuy %#v", arg)
		//优化
		rsp, diamond, coin := handler.Buy(arg, rs.User)
		//同步兑换
		rs.addCurrency(diamond, coin, 0, 0, int32(pb.LOG_TYPE18))
		//响应
		rs.Send(rsp)
	case *pb.CShop:
		arg := msg.(*pb.CShop)
		glog.Debugf("CShop %#v", arg)
		//响应
		rsp := handler.Shop(arg, rs.User)
		rs.Send(rsp)
	case *pb.CUserData:
		arg := msg.(*pb.CUserData)
		glog.Debugf("CUserData %#v", arg)
		userid := arg.GetUserid()
		if userid == "" {
			userid = rs.User.GetUserid()
		}
		if userid != rs.User.GetUserid() {
			msg1 := new(pb.GetUserData)
			msg1.Userid = userid
			rs.rolePid.Request(msg1, ctx.Self())
		} else {
			//TODO 添加房间数据返回
			rsp := handler.GetUserDataMsg(arg, rs.User)
			rs.Send(rsp)
		}
	case *pb.GotUserData:
		arg := msg.(*pb.GotUserData)
		glog.Debugf("GotUserData %#v", arg)
		rsp := handler.UserDataMsg(arg)
		rs.Send(rsp)
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerPay(msg, ctx)
	}
}

/*
func (rs *RoleActor) addPrize(rtype, ltype, amount int32) {
	switch uint32(rtype) {
	case data.DIAMOND:
		rs.addCurrency(amount, 0, 0, 0, ltype)
	case data.COIN:
		rs.addCurrency(0, amount, 0, 0, ltype)
	case data.CARD:
		rs.addCurrency(0, 0, amount, 0, ltype)
	case data.CHIP:
		rs.addCurrency(0, 0, 0, amount, ltype)
	}
}

//消耗钻石
func (rs *RoleActor) expend(cost uint32, ltype int32) {
	diamond := -1 * int64(cost)
	rs.addCurrency(diamond, 0, 0, 0, ltype)
}
*/

//奖励发放
func (rs *RoleActor) addCurrency(diamond, coin, card, chip int64, ltype int32) {
	if rs.User == nil {
		glog.Errorf("add currency user err: %d", ltype)
		return
	}
	//日志记录
	if diamond < 0 && ((rs.User.GetDiamond() + diamond) < 0) {
		diamond = 0 - rs.User.GetDiamond()
	}
	if chip < 0 && ((rs.User.GetChip() + chip) < 0) {
		chip = 0 - rs.User.GetChip()
	}
	if coin < 0 && ((rs.User.GetCoin() + coin) < 0) {
		coin = 0 - rs.User.GetCoin()
	}
	if card < 0 && ((rs.User.GetCard() + card) < 0) {
		card = 0 - rs.User.GetCard()
	}
	rs.User.AddCurrency(diamond, coin, card, chip)
	//货币变更及时同步
	msg2 := handler.ChangeCurrencyMsg(diamond, coin,
		card, chip, ltype, rs.User.GetUserid())
	rs.rolePid.Tell(msg2)
	//消息
	msg := handler.PushCurrencyMsg(diamond, coin,
		card, chip, ltype)
	rs.Send(msg)
	//TODO 机器人不写日志
	//if rs.User.GetRobot() {
	//	return
	//}
	//rs.status = true
	//日志
	//TODO 日志放在dbms中统一写入
	//if diamond != 0 {
	//	msg1 := handler.LogDiamondMsg(diamond, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
	//if coin != 0 {
	//	msg1 := handler.LogCoinMsg(coin, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
	//if card != 0 {
	//	msg1 := handler.LogCardMsg(card, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
	//if chip != 0 {
	//	msg1 := handler.LogChipMsg(chip, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
}

//同步数据
func (rs *RoleActor) syncUser() {
	if rs.User == nil {
		return
	}
	if rs.rolePid == nil {
		return
	}
	if !rs.status { //有变更才同步
		return
	}
	rs.status = false
	msg := new(pb.SyncUser)
	msg.Userid = rs.User.GetUserid()
	result, err := json.Marshal(rs.User)
	if err != nil {
		glog.Errorf("user %s Marshal err %v", rs.User.GetUserid(), err)
		return
	}
	msg.Data = result
	rs.rolePid.Tell(msg)
}
