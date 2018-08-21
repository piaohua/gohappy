package handler

import (
	"gohappy/data"
	"gohappy/pb"
)

//PackLHFreeUser 打包进入百人玩家基础数据
func PackLHFreeUser(p *data.User) *pb.LHFreeUser {
	return &pb.LHFreeUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
	}
}

//PackLHRoomBets 打包百人下注信息
func PackLHRoomBets(bets map[uint32]int64) (msg []*pb.LHRoomBets) {
	for k, v := range bets {
		msg2 := &pb.LHRoomBets{
			Seat: k,
			Bets: v,
		}
		msg = append(msg, msg2)
	}
	return
}

//PackLHFreeRoom 打包百人房间信息
func PackLHFreeRoom(d *data.DeskData) *pb.LHFreeRoom {
	return &pb.LHFreeRoom{
		Roomid: d.Rid,   //牌局id
		Gtype:  d.Gtype, //game type
		Rtype:  d.Rtype, //room type
		Dtype:  d.Dtype, //desk type
		Rname:  d.Rname, //room name
		Count:  d.Count, //当前房间限制玩家数量
		Ante:   d.Ante,  //房间底分
	}
}

//LHLeaveMsg 离开消息
func LHLeaveMsg(userid string, seat uint32) *pb.SLHLeave {
	return &pb.SLHLeave{
		Seat:   seat,
		Userid: userid,
	}
}

//LhdBeDealerMsg 上下庄消息
func LhdBeDealerMsg(state int32, num, carry int64, dealer,
	userid, name, photo string) *pb.SLHFreeDealer {
	return &pb.SLHFreeDealer{
		State:    state,
		Coin:     uint32(num),
		Userid:   userid,
		Dealer:   dealer,
		Nickname: name,
		Photo:    photo,
		Carry:    uint32(carry),
	}
}

//PackLHCoinRoom 打包百人房间信息
func PackLHCoinRoom(d *data.DeskData) *pb.LHRoomData {
	return &pb.LHRoomData{
		Roomid:   d.Rid,     //牌局id
		Gtype:    d.Gtype,   //game type
		Rtype:    d.Rtype,   //room type
		Dtype:    d.Dtype,   //desk type
		Ltype:    d.Ltype,   //level type
		Rname:    d.Rname,   //room name
		Count:    d.Count,   //当前房间限制玩家数量
		Ante:     d.Ante,    //房间底分
		Round:    d.Round,   //
		Userid:   d.Cid,     //
		Expire:   d.Expire,  //
		Code:     d.Code,    //
		Minimum:  d.Minimum, //
		Maximum:  d.Maximum, //
		Pub:      d.Pub,
		Mode:     d.Mode,
		Multiple: d.Multiple,
	}
}

//PackLHCoinUser 打包进入百人玩家基础数据
func PackLHCoinUser(p *data.User) *pb.LHRoomUser {
	return &pb.LHRoomUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
		Lat:      p.Lat,
		Lng:      p.Lng,
		Address:  p.Address,
		Sign:     p.Sign,
	}
}