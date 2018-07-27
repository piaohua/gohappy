package main

import (
	"api/jtpay"
	"gohappy/data"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
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
	//充值数量
	var diamond, coin int64 = handler.GetGoods(trade)
	//充值消息提醒
	record, msg1 := handler.BuyNotice(coin, trade.Userid)
	if record != nil {
		loggerPid.Tell(record)
	}
	//订单状态更新
	handler.WxpaySendGoods(trade, user)
	//充值赠送,TODO 区分日志记录
	diamond2, coin2, msg2 := handler.FirstPay(trade.First, user)
	diamond += diamond2
	coin += coin2
	//更新有效代理绑定
	if msg2 != nil {
		rolePid.Tell(msg2)
	}
	//在线处理
	if v, ok := a.roles[trade.Userid]; ok {
		handler.WxpaySendGoodsOnline(trade, user)
		//在线,发送给玩家pid处理
		msg := new(pb.WxpayGoods)
		msg.Userid = trade.Userid
		msg.Orderid = trade.Id
		msg.Money = trade.Money
		//msg.Diamond = int64(trade.Diamond)
		msg.First = int32(trade.First)
		msg.Diamond = diamond
		msg.Coin = coin
		v.Pid.Tell(msg)
		if msg1 != nil {
			v.Pid.Tell(msg1)
		}
		return
	}
	//离线处理
	handler.WxpaySendGoodsOffline(trade, user)
	//货币变更及时同步
	msg3 := handler.ChangeCurrencyMsg(diamond, coin,
		0, 0, int32(pb.LOG_TYPE4), trade.Userid)
	//msg3.Money = int64(trade.Money)
	rolePid.Tell(msg3)
}

//货币变更及时同步
func (a *RoleActor) sendCurrency(userid string, diamond, coin int64, ltype int32) {
	msg := handler.ChangeCurrencyMsg(diamond, coin,0, 0, ltype, userid)
	//msg.Money = int64(trade.Money)
	if v, ok := a.roles[userid]; ok {
		v.Pid.Tell(msg)
		return
	}
	rolePid.Tell(msg)
}
