package main

import (
	"time"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家数据请求处理
func (rs *RoleActor) handlerPay(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CApplePay:
		arg := msg.(*pb.CApplePay)
		glog.Debugf("CApplePay %#v", arg)
		rs.applePay(arg)
	case *pb.CWxpayOrder:
		arg := msg.(*pb.CWxpayOrder)
		glog.Debugf("CWxpayOrder %#v", arg)
		rs.wxPay(arg)
	case *pb.CWxpayQuery:
		arg := msg.(*pb.CWxpayQuery)
		glog.Debugf("CWxpayQuery %#v", arg)
		rsp := handler.WxQuery(arg)
		rs.Send(rsp)
	case *pb.WxpayGoods:
		arg := msg.(*pb.WxpayGoods)
		//发货
		glog.Debugf("WxpayGoods: %v", arg)
		//userid := arg.Userid
		msg2 := new(pb.SWxpayQuery)
		msg2.Orderid = arg.Orderid
		rs.Send(msg2)
		rs.sendGoods(arg.Diamond, arg.Coin, arg.Money, int(arg.First))
	case *pb.PayCurrency:
		arg := msg.(*pb.PayCurrency)
		glog.Debugf("PayCurrency %#v", arg)
		//后台或充值同步到game房间
		if rs.gamePid != nil {
			msg2 := handler.Pay2ChangeCurr(arg)
			rs.gamePid.Tell(msg2)
		}
		diamond := arg.Diamond
		coin := arg.Coin
		chip := arg.Chip
		card := arg.Card
		ltype := arg.Type
		rs.addCurrency(diamond, coin, card, chip, ltype)
	case *pb.ChangeCurrency:
		//货币变更
		arg := msg.(*pb.ChangeCurrency)
		diamond := arg.Diamond
		coin := arg.Coin
		chip := arg.Chip
		card := arg.Card
		ltype := arg.Type
		rs.addCurrency(diamond, coin, card, chip, ltype)
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerAgent(msg, ctx)
	}
}

func (rs *RoleActor) applePay(arg *pb.CApplePay) {
	rsp, record, trade := handler.AppleOrder(arg, rs.User)
	if rsp.Error != pb.OK {
		rs.Send(rsp)
		return
	}
	//验证
	msg1 := new(pb.ApplePay)
	msg1.Trade = trade
	timeout := 3 * time.Second
	res1, err1 := rs.rolePid.RequestFuture(msg1, timeout).Result()
	if err1 != nil {
		glog.Errorf("ApplePay err: %v", err1)
		rsp.Error = pb.AppleOrderFail
		rs.Send(rsp)
		return
	}
	if response1, ok := res1.(*pb.ApplePaid); ok {
		if !response1.Result {
			glog.Error("ApplePay fail")
			rsp.Error = pb.AppleOrderFail
			rs.Send(rsp)
			return
		}
	} else {
		glog.Error("ApplePay fail")
		rsp.Error = pb.AppleOrderFail
		rs.Send(rsp)
		return
	}
	rs.sendGoods(int64(record.Diamond), 0, record.Money, record.First)
	rs.Send(rsp)
}

func (rs *RoleActor) wxPay(arg *pb.CWxpayOrder) {
	var ip string = rs.User.LoginIP
	rsp, trade := handler.WxOrder(arg, rs.User, ip)
	if rsp.Error != pb.OK {
		rs.Send(rsp)
		return
	}
	//验证
	msg1 := new(pb.ApplePay)
	msg1.Trade = trade
	timeout := 3 * time.Second
	res1, err1 := rs.rolePid.RequestFuture(msg1, timeout).Result()
	if err1 != nil {
		glog.Errorf("wxPay err: %v", err1)
		rsp.Error = pb.PayOrderFail
		rs.Send(rsp)
		return
	}
	if response1, ok := res1.(*pb.ApplePaid); ok {
		if !response1.Result {
			glog.Error("wxPay fail")
			rsp.Error = pb.PayOrderFail
			rs.Send(rsp)
			return
		}
	} else {
		glog.Error("wxPay fail")
		rsp.Error = pb.PayOrderFail
		rs.Send(rsp)
		return
	}
	//下单成功
	rs.Send(rsp)
	//主动查询发货
	go rs.wxPayQuery(rsp.Orderid)
}

//主动查询发货
func (rs *RoleActor) wxPayQuery(orderid string) {
	//查询
	result := handler.ActWxpayQuery(orderid) //查询
	if result == "" {
		return
	}
	if rs.rolePid == nil {
		return
	}
	//发货
	msg2 := new(pb.WxpayCallback)
	msg2.Result = result
	rs.rolePid.Tell(msg2)
}

//发货
func (rs *RoleActor) sendGoods(diamond, coin int64, money uint32, first int) {
	rs.User.AddMoney(money) //TODO 同步money到数据库
	//消息
	rs.addCurrency(diamond, coin, 0, 0, int32(pb.LOG_TYPE4))
	//消息
	stoc := new(pb.SGetCurrency)
	stoc.Data = handler.PackCurrency(rs.User)
	rs.Send(stoc)
}
