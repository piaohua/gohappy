package handler

import (
	"gohappy/data"
	"gohappy/pb"
)

//PackEBFreeUser 打包进入百人玩家基础数据
func PackEBFreeUser(p *data.User) *pb.EBFreeUser {
	return &pb.EBFreeUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
	}
}

//PackEBRoomBets 打包百人下注信息
func PackEBRoomBets(bets map[uint32]int64) (msg []*pb.EBRoomBets) {
	for k, v := range bets {
		msg2 := &pb.EBRoomBets{
			Seat: k,
			Bets: v,
		}
		msg = append(msg, msg2)
	}
	return
}

//PackEBFreeRoom 打包百人房间信息
func PackEBFreeRoom(d *data.DeskData) *pb.EBFreeRoom {
	return &pb.EBFreeRoom{
		Roomid: d.Rid,   //牌局id
		Gtype:  d.Gtype, //game type
		Rtype:  d.Rtype, //room type
		Dtype:  d.Dtype, //desk type
		Rname:  d.Rname, //room name
		Count:  d.Count, //当前房间限制玩家数量
		Ante:   d.Ante,  //房间底分
	}
}

//EBLeaveMsg 离开消息
func EBLeaveMsg(userid string, seat uint32) *pb.SEBLeave {
	return &pb.SEBLeave{
		Seat:   seat,
		Userid: userid,
	}
}

//EbgBeDealerMsg 上下庄消息
func EbgBeDealerMsg(state int32, num, carry int64, dealer,
	userid, name, photo string) *pb.SEBFreeDealer {
	return &pb.SEBFreeDealer{
		State:    state,
		Coin:     uint32(num),
		Userid:   userid,
		Dealer:   dealer,
		Nickname: name,
		Photo:    photo,
		Carry:    uint32(carry),
	}
}

//PackEBCoinRoom 打包百人房间信息
func PackEBCoinRoom(d *data.DeskData) *pb.EBRoomData {
	return &pb.EBRoomData{
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

//PackEBCoinUser 打包进入百人玩家基础数据
func PackEBCoinUser(p *data.User) *pb.EBRoomUser {
	return &pb.EBRoomUser{
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

//PackEBCreateMsg 创建房间消息
func PackEBCreateMsg(d *data.DeskData) *pb.SEBCreateRoom {
	msg := new(pb.SEBCreateRoom)
	msg.Data = PackEBCoinRoom(d)
	return msg
}

//PackEBRoomList 房间列表数据
func PackEBRoomList(arg *pb.GetRoomList,
	desks map[string]*data.DeskBase) *pb.SEBRoomList {
	msg := new(pb.SEBRoomList)
	for _, v := range desks {
		switch arg.Rtype {
		case int32(pb.ROOM_TYPE1): //私人
			if v.DeskData.Cid != arg.Userid && !v.DeskData.Pub {
				continue
			}
		}
		if v.DeskData.Gtype == arg.Gtype &&
			v.DeskData.Rtype == arg.Rtype {
			msg2 := PackEBCoinRoom(v.DeskData)
			msg2.Number = v.Number
			msg.List = append(msg.List, msg2)
		}
	}
	return msg
}
