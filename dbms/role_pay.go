package main

import (
	"gohappy/game/handler"
	"gohappy/pb"
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
