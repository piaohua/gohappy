package handler

import (
	"gohappy/data"
	"gohappy/pb"
)

//PackJHFreeUser 打包进入百人玩家基础数据
func PackJHFreeUser(p *data.User) *pb.JHFreeUser {
	return &pb.JHFreeUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
	}
}

//PackJHRoomBets 打包百人下注信息
func PackJHRoomBets(bets map[uint32]int64) (msg []*pb.JHRoomBets) {
	for k, v := range bets {
		msg2 := &pb.JHRoomBets{
			Seat: k,
			Bets: v,
		}
		msg = append(msg, msg2)
	}
	return
}

//PackJHFreeRoom 打包百人房间信息
func PackJHFreeRoom(d *data.DeskData) *pb.JHFreeRoom {
	return &pb.JHFreeRoom{
		Roomid: d.Rid,   //牌局id
		Gtype:  d.Gtype, //game type
		Rtype:  d.Rtype, //room type
		Dtype:  d.Dtype, //desk type
		Rname:  d.Rname, //room name
		Count:  d.Count, //当前房间限制玩家数量
		Ante:   d.Ante,  //房间底分
	}
}

//JHLeaveMsg 离开消息
func JHLeaveMsg(userid string, seat uint32) *pb.SJHLeave {
	return &pb.SJHLeave{
		Seat:   seat,
		Userid: userid,
	}
}

//BeJHDealerMsg 上下庄消息
func BeJHDealerMsg(state int32, num int64, dealer,
	userid, name string) *pb.SJHFreeDealer {
	return &pb.SJHFreeDealer{
		State:    state,
		Coin:     uint32(num),
		Userid:   userid,
		Dealer:   dealer,
		Nickname: name,
	}
}

//PackJHCoinRoom 打包百人房间信息
func PackJHCoinRoom(d *data.DeskData) *pb.JHRoomData {
	return &pb.JHRoomData{
		Roomid:  d.Rid,     //牌局id
		Gtype:   d.Gtype,   //game type
		Rtype:   d.Rtype,   //room type
		Dtype:   d.Dtype,   //desk type
		Ltype:   d.Ltype,   //level type
		Rname:   d.Rname,   //room name
		Count:   d.Count,   //当前房间限制玩家数量
		Ante:    d.Ante,    //房间底分
		Round:   d.Round,   //
		Userid:  d.Cid,     //
		Expire:  d.Expire,  //
		Code:    d.Code,    //
		Minimum: d.Minimum, //
		Maximum: d.Maximum, //
	}
}

//PackJHCoinUser 打包进入百人玩家基础数据
func PackJHCoinUser(p *data.User) *pb.JHRoomUser {
	return &pb.JHRoomUser{
		Userid:   p.GetUserid(),
		Nickname: p.GetNickname(),
		Phone:    p.GetPhone(),
		Sex:      p.GetSex(),
		Photo:    p.GetPhoto(),
		Coin:     p.GetCoin(),
		Diamond:  p.GetDiamond(),
	}
}

//PackJHCreateMsg 创建房间消息
func PackJHCreateMsg(d *data.DeskData) *pb.SJHCreateRoom {
	msg := new(pb.SJHCreateRoom)
	msg.Data = PackJHCoinRoom(d)
	return msg
}

//PackJHRoomList 房间列表数据
func PackJHRoomList(arg *pb.GetRoomList,
	desks map[string]*data.DeskBase) *pb.SJHRoomList {
	msg := new(pb.SJHRoomList)
	for _, v := range desks {
		switch arg.Rtype {
		case int32(pb.ROOM_TYPE1): //私人
			if v.DeskData.Cid != arg.Userid {
				continue
			}
		}
		if v.DeskData.Gtype == arg.Gtype &&
			v.DeskData.Rtype == arg.Rtype {
			msg2 := PackJHCoinRoom(v.DeskData)
			msg2.Number = v.Number
			msg.List = append(msg.List, msg2)
		}
	}
	return msg
}
