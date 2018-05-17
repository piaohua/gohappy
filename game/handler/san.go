package handler

import (
	"gohappy/data"
	"gohappy/pb"
)

//PackSGFreeUser 打包进入百人玩家基础数据
func PackSGFreeUser(p *data.User) *pb.SGFreeUser {
	return &pb.SGFreeUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
	}
}

//PackSGRoomBets 打包百人下注信息
func PackSGRoomBets(bets map[uint32]int64) (msg []*pb.SGRoomBets) {
	for k, v := range bets {
		msg2 := &pb.SGRoomBets{
			Seat: k,
			Bets: v,
		}
		msg = append(msg, msg2)
	}
	return
}

//PackSGFreeRoom 打包百人房间信息
func PackSGFreeRoom(d *data.DeskData) *pb.SGFreeRoom {
	return &pb.SGFreeRoom{
		Roomid: d.Rid,   //牌局id
		Gtype:  d.Gtype, //game type
		Rtype:  d.Rtype, //room type
		Dtype:  d.Dtype, //desk type
		Rname:  d.Rname, //room name
		Count:  d.Count, //当前房间限制玩家数量
		Ante:   d.Ante,  //房间底分
	}
}

//SGLeaveMsg 离开消息
func SGLeaveMsg(userid string, seat uint32) *pb.SSGLeave {
	return &pb.SSGLeave{
		Seat:   seat,
		Userid: userid,
	}
}

//BeSGDealerMsg 上下庄消息
func BeSGDealerMsg(state int32, num int64, dealer,
	userid, name string) *pb.SSGFreeDealer {
	return &pb.SSGFreeDealer{
		State:    state,
		Coin:     uint32(num),
		Userid:   userid,
		Dealer:   dealer,
		Nickname: name,
	}
}

//PackSGCoinRoom 打包百人房间信息
func PackSGCoinRoom(d *data.DeskData) *pb.SGRoomData {
	return &pb.SGRoomData{
		Roomid: d.Rid,    //牌局id
		Gtype:  d.Gtype,  //game type
		Rtype:  d.Rtype,  //room type
		Dtype:  d.Dtype,  //desk type
		Rname:  d.Rname,  //room name
		Count:  d.Count,  //当前房间限制玩家数量
		Ante:   d.Ante,   //房间底分
		Round:  d.Round,  //
		Userid: d.Cid,    //
		Expire: d.Expire, //
		Code:   d.Code,   //
	}
}

//PackSGCoinUser 打包进入百人玩家基础数据
func PackSGCoinUser(p *data.User) *pb.SGRoomUser {
	return &pb.SGRoomUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
	}
}

//PackSGCreateMsg 创建房间消息
func PackSGCreateMsg(d *data.DeskData) *pb.SSGCreateRoom {
	msg := new(pb.SSGCreateRoom)
	msg.Data = PackSGCoinRoom(d)
	return msg
}

//PackSGRoomList 房间列表数据
func PackSGRoomList(arg *pb.GetRoomList,
	desks map[string]*data.DeskBase) *pb.SSGRoomList {
	msg := new(pb.SSGRoomList)
	for _, v := range desks {
		switch arg.Rtype {
		case int32(pb.ROOM_TYPE1): //私人
			if v.DeskData.Cid != arg.Userid {
				continue
			}
		}
		if v.DeskData.Gtype == arg.Gtype &&
			v.DeskData.Rtype == arg.Rtype {
			msg2 := PackSGCoinRoom(v.DeskData)
			msg.List = append(msg.List, msg2)
		}
	}
	return msg
}
