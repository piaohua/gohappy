package handler

import (
	"gohappy/data"
	"gohappy/pb"
	"utils"
	"gohappy/glog"
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
	p *data.User) (msg *pb.LogBank) {
	msg = &pb.LogBank{
		Userid: p.GetUserid(),
		Type:   int32(ltype),
		Num:    num,
		Rest:   p.GetBank(),
	}
	return
}

//LogRoomRecordMsg 打包创建房间记录日志消息
func LogRoomRecordMsg(d *data.DeskData) (msg *pb.RoomRecordInfo) {
	msg = &pb.RoomRecordInfo{
		Roomid:   d.Rid,
		Gtype:    d.Gtype,
		Rtype:    d.Rtype,
		Dtype:    d.Dtype,
		Rname:    d.Rname,
		Count:    d.Count,
		Ante:     d.Ante,
		Code:     d.Code,
		Round:    d.Round,
		Cid:      d.Cid,
		Ctime:    d.Ctime,
		Pub:      d.Pub,
		Minimum:  d.Minimum,
		Maximum:  d.Maximum,
		Mode:     d.Mode,
		Multiple: d.Multiple,
	}
	return
}

//Log2RoomRecord 创建房间记录日志
func Log2RoomRecord(msg *pb.RoomRecordInfo) {
	r := &data.RoomRecord{
		Roomid:   msg.Roomid,
		Gtype:    msg.Gtype,
		Rtype:    msg.Rtype,
		Dtype:    msg.Dtype,
		Rname:    msg.Rname,
		Count:    msg.Count,
		Ante:     msg.Ante,
		Code:     msg.Code,
		Round:    msg.Round,
		Cid:      msg.Cid,
		Ctime:    msg.Ctime,
		Pub:      msg.Pub,
		Minimum:  msg.Minimum,
		Maximum:  msg.Maximum,
		Mode:     msg.Mode,
		Multiple: msg.Multiple,
	}
	r.Save()
}

//Log2RoleRecord 个人房间结果记录
func Log2RoleRecord(msg *pb.RoleRecord) {
	r := &data.RoleRecord{
		Roomid:   msg.Roomid,
		Gtype:    msg.Gtype,
		Userid:   msg.Userid,
		Nickname: msg.Nickname,
		Photo:    msg.Photo,
		Score:    msg.Score,
		Rest:     msg.Rest,
		Joins:    msg.Joins,
	}
	r.Save()
}

//Log2RoundRecord 每局结算详情记录
func Log2RoundRecord(msg *pb.RoundRecord) {
	r := &data.RoundRecord{
		Roomid: msg.Roomid,
		Round:  msg.Round,
		Dealer: msg.Dealer,
	}
	for _, v := range msg.Roles {
		rs := data.RoundRoleRecord{
			Userid: v.Userid,
			Cards:  v.Cards,
			Value:  v.Value,
			Score:  v.Score,
			Rest:   v.Rest,
			Bets:   v.Bets,
		}
		r.Roles = append(r.Roles, rs)
	}
	r.Save()
}

//RoomRecordInfoMsg 房间记录信息
func RoomRecordInfoMsg(msg *data.RoomRecord) *pb.RoomRecordInfo {
	r := &pb.RoomRecordInfo{
		Roomid:   msg.Roomid,
		Gtype:    msg.Gtype,
		Rtype:    msg.Rtype,
		Dtype:    msg.Dtype,
		Rname:    msg.Rname,
		Count:    msg.Count,
		Ante:     msg.Ante,
		Code:     msg.Code,
		Round:    msg.Round,
		Cid:      msg.Cid,
		Ctime:    msg.Ctime,
		Pub:      msg.Pub,
		Minimum:  msg.Minimum,
		Maximum:  msg.Maximum,
		Mode:     msg.Mode,
		Multiple: msg.Multiple,
	}
	return r
}

//RoundRecordMsg 房间记录信息
func RoundRecordMsg(msg *data.RoundRecord) *pb.RoundRecord {
	r := &pb.RoundRecord{
		Roomid: msg.Roomid,
		Round:  msg.Round,
		Dealer: msg.Dealer,
		Ctime:  utils.String(utils.Time2Stamp(msg.Ctime)),
	}
	for _, v := range msg.Roles {
		rs := &pb.RoundRoleRecord{
			Userid: v.Userid,
			Cards:  v.Cards,
			Value:  v.Value,
			Score:  v.Score,
			Rest:   v.Rest,
			Bets:   v.Bets,
		}
		r.Roles = append(r.Roles, rs)
	}
	return r
}

//RoleRecordMsg 个人结果记录
func RoleRecordMsg(msg *data.RoleRecord) *pb.RoleRecord {
	r := &pb.RoleRecord{
		Roomid:   msg.Roomid,
		Gtype:    msg.Gtype,
		Userid:   msg.Userid,
		Nickname: msg.Nickname,
		Photo:    msg.Photo,
		Score:    msg.Score,
		Rest:     msg.Rest,
		Joins:    msg.Joins,
	}
	return r
}

//LogProfitMsg 打包收益日志消息
func LogProfitMsg(agentid, userid string, gtype int32, level, rate uint32,
	profit int64) (msg *pb.LogProfit) {
	msg = &pb.LogProfit{
		Agentid: agentid,
		Userid:  userid,
		Gtype:   gtype,
		Level:   level,
		Rate:    rate,
		Profit:  profit,
	}
	return
}

//LogSysProfitMsg 打包系统收益日志消息
func LogSysProfitMsg(agentid, userid string, gtype int32, level, rate uint32,
	profit, rest int64) (msg *pb.LogSysProfit) {
	msg = &pb.LogSysProfit{
		Agentid: agentid,
		Userid:  userid,
		Gtype:   gtype,
		Level:   level,
		Rate:    rate,
		Profit:  profit,
		Rest:    rest,
	}
	return
}

//PackBankLogMsg 获取银行操作记录信息
func PackBankLogMsg(arg *pb.CBankLog) (msg *pb.SBankLog) {
	msg = new(pb.SBankLog)
	list, err := data.GetBankLogs(arg)
	msg.Page = arg.Page
	msg.Count = uint32(len(list))
	if err != nil {
		glog.Errorf("GetBankLogs err %v", err)
	}
	glog.Debugf("PackBankLogMsg list %#v", list)
	for _, v := range list {
		msg2 := new(pb.BankLog)
		msg2.Ctime  = utils.Time2Str(v.Ctime) //时间
		msg2.Type   = v.Type //类型
		msg2.Num    = v.Num //数量
		msg2.Rest   = v.Rest //银行剩余数量
		msg.List = append(msg.List, msg2)
	}
	return
}