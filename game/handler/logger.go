package handler

import (
	"gohappy/data"
	"gohappy/pb"
)

//LogDiamondMsg 打包钻石日志消息
func LogDiamondMsg(num int64, ltype int32,
	p *data.User) (msg *pb.LogDiamond) {
	msg = &pb.LogDiamond{
		Userid: p.GetUserid(),
		Type:   int32(ltype),
		Num:    num,
		Rest:   p.GetDiamond(),
	}
	return
}

//LogCoinMsg 打包金币日志消息
func LogCoinMsg(num int64, ltype int32,
	p *data.User) (msg *pb.LogCoin) {
	msg = &pb.LogCoin{
		Userid: p.GetUserid(),
		Type:   int32(ltype),
		Num:    num,
		Rest:   p.GetCoin(),
	}
	return
}

//LogCardMsg 打包卡片日志消息
func LogCardMsg(num int64, ltype int32,
	p *data.User) (msg *pb.LogCard) {
	msg = &pb.LogCard{
		Userid: p.GetUserid(),
		Type:   int32(ltype),
		Num:    num,
		Rest:   p.GetCard(),
	}
	return
}

//LogChipMsg 打包筹码日志消息
func LogChipMsg(num int64, ltype int32,
	p *data.User) (msg *pb.LogChip) {
	msg = &pb.LogChip{
		Userid: p.GetUserid(),
		Type:   int32(ltype),
		Num:    num,
		Rest:   p.GetChip(),
	}
	return
}

//LogBankMsg 打包银行日志消息
func LogBankMsg(num int64, ltype int32,
	p *data.User) (msg *pb.LogCoin) {
	msg = &pb.LogCoin{
		Userid: p.GetUserid(),
		Type:   int32(ltype),
		Num:    num,
		Rest:   p.GetBank(),
	}
	return
}
