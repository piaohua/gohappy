/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-12-27 23:16:38
 * Filename      : response.go
 * Description   : 响应协议消息
 * *******************************************************/
package handler

/*
import (
	"gohappy/data"
	"gohappy/pb"
)

//坐下,起来离开消息
func SitMsg(userid string, seat uint32,
	state bool, p *data.User) interface{} {
	return &pb.SHuiYinSit{
		Seat:     seat,
		State:    state,
		Userid:   userid,
		Nickname: p.GetNickname(),
		Photo:    p.GetPhoto(),
		Chip:     p.GetChip(),
	}
}

//下注消息
func RoomBetMsg(seat, beseat, val uint32, chip,
	bets int64, userid string) interface{} {
	return &pb.SHuiYinRoomBet{
		Seat:   seat,
		Beseat: beseat,
		Value:  val,
		Chip:   chip,
		Bets:   bets,
		Userid: userid,
	}
}

//上下庄消息
func BeDealerMsg(state uint32, num int64, dealer,
	userid, name string, down bool) interface{} {
	return &pb.SHuiYinDealer{
		State:    state,
		Num:      num,
		Userid:   userid,
		Dealer:   dealer,
		Nickname: name,
		Down:     down,
	}
}

//文本聊天消息
func ChatMsg(seat uint32, userid string, msg []byte) interface{} {
	return &pb.SChatText{
		Seat:    seat,
		Userid:  userid,
		Content: msg,
	}
}

//语音聊天消息
func ChatMsg2(seat uint32, userid string, msg []byte) interface{} {
	return &pb.SChatVoice{
		Seat:    seat,
		Userid:  userid,
		Content: msg,
	}
}

//离开消息
func LeaveMsg(userid string, seat uint32) interface{} {
	return &pb.SHuiYinLeave{
		Seat:   seat,
		Userid: userid,
	}
}

//打包房间基础数据
func PackRoomInfo(d *data.DeskData) (r *pb.HuiYinRoomInfo) {
	r = new(pb.HuiYinRoomInfo)
	r.Info = PackGameInfo(d)
	return
}

// 游戏房间信息
func PackGameInfo(d *data.DeskData) (r *pb.HuiYinGameInfo) {
	return &pb.HuiYinGameInfo{
		Roomid: d.Rid,
		Gtype:  d.Gtype,
		Rtype:  d.Rtype,
		Ltype:  d.Ltype,
		Rname:  d.Rname,
		Count:  d.Count,
		Ante:   d.Ante,
		Cost:   d.Cost,
		Vip:    d.Vip,
		Chip:   d.Chip,
		Deal:   d.Deal,
		Carry:  d.Carry,
		Down:   d.Down,
		Top:    d.Top,
		Sit:    d.Sit,
	}
}
*/
