package handler

import (
	"gohappy/data"
	"gohappy/game/config"
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
