package main

import (
	"api/jtpay"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
	"gohappy/data"
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
	a.tradeHandler(trade) //发货
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
		glog.Errorf("jtpayHandler err1 %v", err1)
		return false
	}
	//订单验证
	trade := handler.JtpayTradeVerify(notifyResult)
	if trade == nil {
		glog.Errorf("jtpayHandler verify err %#v", notifyResult)
		return false
	}
	a.tradeHandler(trade) //发货
	return true
}

//订单发货处理
func (a *RoleActor) tradeHandler(trade *data.TradeRecord) {
	user := a.getUserById(trade.Userid)
	var diamond, coin int64 = handler.GetGoods(trade)
	//充值消息提醒
	record, msg2 := handler.BuyNotice(coin, user.GetUserid())
	if record != nil {
		loggerPid.Tell(record)
	}
	//在线
	if v, ok := a.roles[trade.Userid]; ok {
		handler.WxpaySendGoods(true, trade, user)
		//在线,发送给玩家pid处理
		msg := new(pb.WxpayGoods)
		msg.Userid = user.GetUserid()
		msg.Orderid = trade.Id
		msg.Money = trade.Money
		//msg.Diamond = int64(trade.Diamond)
		msg.First = int32(trade.First)
		msg.Diamond = diamond
		msg.Coin = coin
		v.Pid.Tell(msg)
		if msg2 != nil {
			v.Pid.Tell(msg2)
		}
		return
	}
	//离线,直接处理
	handler.WxpaySendGoods(false, trade, user)
	if trade.First == 1 {
		//更新有效代理绑定
		msg := handler.AgentBuildUpdateMsg(user.GetAgent(), 0, 1, 0)
		rolePid.Tell(msg)
	}
}