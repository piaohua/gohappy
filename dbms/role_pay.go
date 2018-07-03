package main

import (
	"gohappy/game/handler"
	"gohappy/pb"
	"github.com/AsynkronIT/protoactor-go/actor"
	"gohappy/glog"
	"utils"
	"api/jtpay"
)

//支付处理
func (a *RoleActor) payHandler(arg *pb.WxpayCallback) {
	//数据解析
	result := handler.WxpayCallback(arg)
	if result == nil {
		return
	}
	//订单验证
	trade := handler.WxpayTradeVerify(result)
	if trade == nil {
		return
	}
	userid := trade.Userid
	//发货,TODO 优化
	if user, ok := a.online[userid]; ok {
		handler.WxpaySendGoods(true, trade, user)
		//在线,发送给gate处理
		msg2 := new(pb.WxpayGoods)
		msg2.Userid = user.GetUserid()
		msg2.Orderid = trade.Id
		msg2.Money = trade.Money
		msg2.Diamond = trade.Diamond
		msg2.First = int32(trade.First)
		//a.hallPid.Tell(msg2)
		nodePid.Tell(msg2)
		return
	}
	//离线,直接处理
	user := a.getUserById(userid)
	handler.WxpaySendGoods(false, trade, user)
}

//交易下单
func (a *RoleActor) tradeOrder(arg *pb.TradeOrder, ctx actor.Context) {
	rsp := new(pb.TradedOrder)
	user := a.getUserById(arg.GetUserid())
	if user == nil {
		glog.Errorf("get userid %s fail", arg.GetUserid())
		ctx.Respond(rsp)
		return
	}
	order := handler.Order2Record(arg)
	order.DayStamp = utils.TimestampTodayTime()
	order.Agent = user.GetAgent()
	if order.Save() {
		rsp.Result = true
		glog.Errorf("tradeOrder success %#v", order)
	} else {
		glog.Errorf("tradeOrder failed %#v", order)
	}
	ctx.Respond(rsp)
}

//支付结果通知处理
func (a *RoleActor) jtpayHandler(arg *pb.JtpayCallback) bool {
	//数据解析
	notifyResult := new(jtpay.NotifyResult)
	err1 := json.Unmarshal(arg.Result, notifyResult)
	if err1 != nil {
		return false
	}
	//订单验证
	trade := handler.JtpayTradeVerify(notifyResult)
	if trade == nil {
		return false
	}
	userid := trade.Userid
	user := a.getUserById(userid)
	//发货,TODO 消息通知优化
	if v, ok := a.roles[userid]; ok {
		handler.WxpaySendGoods(true, trade, user)
		//在线,发送给gate处理
		msg2 := new(pb.WxpayGoods)
		msg2.Userid = user.GetUserid()
		msg2.Orderid = trade.Id
		msg2.Money = trade.Money
		msg2.Diamond = trade.Diamond
		msg2.First = int32(trade.First)
		v.Pid.Tell(msg2)
		return true
	}
	//离线,直接处理
	handler.WxpaySendGoods(false, trade, user)
	return true
}