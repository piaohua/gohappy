package handler

import (
	"api/jtpay"
	"gohappy/data"
	"gohappy/game/config"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

//Buy 购买
func Buy(ctos *pb.CBuy, p *data.User) (stoc *pb.SBuy,
	diamond, coin int64) {
	stoc = new(pb.SBuy)
	id := ctos.GetId()
	d := config.GetShop(utils.String(id))
	switch uint32(d.Payway) {
	case data.DIA:
		if p.GetDiamond() >= int64(d.Price) {
			stoc.Result = 0
			diamond = -1 * int64(d.Price)
			coin = int64(d.Number)
		} else {
			stoc.Result = 1
			stoc.Error = pb.NotEnoughDiamond
		}
	default:
		stoc.Error = pb.PayOrderFail
	}
	return
}

//Shop 商城列表
func Shop(ctos *pb.CShop, p *data.User) (stoc *pb.SShop) {
	stoc = new(pb.SShop)
	if p == nil {
		return
	}
	list := config.GetShops()
	for _, v := range list {
		id := utils.Uint64(v.Id)
		s := &pb.Shop{
			Id:     uint32(id),       //购买ID
			Status: uint32(v.Status), //物品状态,1=热卖
			Propid: uint32(v.Propid), //兑换的物品,1=钻石,2=金币
			Payway: uint32(v.Payway), //支付方式,1=RMB,2=钻石
			Number: v.Number,         //兑换的数量
			Price:  v.Price,          //支付价格
			Name:   v.Name,           //物品名字
			Info:   v.Info,           //物品信息
		}
		stoc.List = append(stoc.List, s)
	}
	return
}

//Order2Record 下单记录
func Order2Record(arg *pb.TradeOrder) (msg *data.TradeRecord) {
	msg = &data.TradeRecord{
		Id:       arg.Orderid,
		Userid:   arg.Userid,
		Amount:   arg.Amount,
		Itemid:   arg.Itemid,
		Diamond:  arg.Diamond,
		Money:    arg.Money,
		Result:   int(arg.Result),
		Clientip: arg.Clientip,
	}
	return
}

//JtpayTradeVerify 发货验证
func JtpayTradeVerify(t *jtpay.NotifyResult) *data.TradeRecord {
	//sign
	tradeRecord := &data.TradeRecord{
		Id: t.P2_ordernumber,
		//Transid: t.TransactionId,
	}
	//订单获取
	tradeRecord.Get()
	//glog.Infof("tradeRecord  %#v", tradeRecord)
	//glog.Infof("TradeResult  %#v", t)
	if tradeRecord.Userid == "" {
		//订单不存在或其它
		glog.Errorf("not exist orderid %#v", t)
		return nil
	}
	if tradeRecord.Result == 0 {
		//重复发货
		glog.Errorf("repeat resp %#v", t)
		return nil
	}
	//更新记录
	tradeRecord.Transid = t.P5_orderid
	tradeRecord.Transtime = utils.Time2Str(utils.LocalTime())
	//tradeRecord.Paytype = t.P6_productcode
	//money, err := strconv.Atoi(t.P13_zfmoney)
	//if err != nil {
	//	glog.Errorf("jtpay: %v, err: %#v", t, err)
	//}
	//if uint32(money*100) != tradeRecord.Money {
	//	glog.Errorf("jtpay money : %#v, err: %v", t, err)
	//}
	//tradeRecord.Money = uint32(money)      //转换为分
	tradeRecord.Result = data.TradeSuccess //交易成功
	//glog.Infof("tradeRecord  %#v", tradeRecord)
	return tradeRecord
}
