package handler

import (
	"gohappy/data"
	"gohappy/pb"
)

//PackNNFreeUser 打包进入百人玩家基础数据
func PackNNFreeUser(p *data.User) *pb.NNFreeUser {
	return &pb.NNFreeUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
	}
}

//PackNNRoomBets 打包百人下注信息
func PackNNRoomBets(bets map[uint32]int64) (msg []*pb.NNRoomBets) {
	for k, v := range bets {
		msg2 := &pb.NNRoomBets{
			Seat: k,
			Bets: v,
		}
		msg = append(msg, msg2)
	}
	return
}

//PackNNFreeRoom 打包百人房间信息
func PackNNFreeRoom(d *data.DeskData) *pb.NNFreeRoom {
	return &pb.NNFreeRoom{
		Roomid: d.Rid,   //牌局id
		Gtype:  d.Gtype, //game type
		Rtype:  d.Rtype, //room type
		Dtype:  d.Dtype, //desk type
		Rname:  d.Rname, //room name
		Count:  d.Count, //当前房间限制玩家数量
		Ante:   d.Ante,  //房间底分
	}
}

//NNLeaveMsg 离开消息
func NNLeaveMsg(userid string, seat uint32) *pb.SNNLeave {
	return &pb.SNNLeave{
		Seat:   seat,
		Userid: userid,
	}
}

//BeDealerMsg 上下庄消息
func BeDealerMsg(state int32, num, carry int64, dealer,
	userid, name, photo string) *pb.SNNFreeDealer {
	return &pb.SNNFreeDealer{
		State:    state,
		Coin:     uint32(num),
		Userid:   userid,
		Dealer:   dealer,
		Nickname: name,
		Photo:    photo,
		Carry:    uint32(carry),
	}
}

//PackNNCoinRoom 打包百人房间信息
func PackNNCoinRoom(d *data.DeskData) *pb.NNRoomData {
	return &pb.NNRoomData{
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

//PackNNCoinUser 打包进入百人玩家基础数据
func PackNNCoinUser(p *data.User) *pb.NNRoomUser {
	return &pb.NNRoomUser{
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

//PackNNCreateMsg 创建房间消息
func PackNNCreateMsg(d *data.DeskData) *pb.SNNCreateRoom {
	msg := new(pb.SNNCreateRoom)
	msg.Data = PackNNCoinRoom(d)
	return msg
}

//PackNNRoomList 房间列表数据
func PackNNRoomList(arg *pb.GetRoomList,
	desks map[string]*data.DeskBase) *pb.SNNRoomList {
	msg := new(pb.SNNRoomList)
	for _, v := range desks {
		switch arg.Rtype {
		case int32(pb.ROOM_TYPE1): //私人
			if v.DeskData.Cid != arg.Userid && !v.DeskData.Pub {
				continue
			}
		}
		if v.DeskData.Gtype == arg.Gtype &&
			v.DeskData.Rtype == arg.Rtype {
			msg2 := PackNNCoinRoom(v.DeskData)
			msg2.Number = v.Number
			msg.List = append(msg.List, msg2)
		}
	}
	return msg
}
