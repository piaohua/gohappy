/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-01-21 18:31:54
 * Filename      : record.go
 * Description   : 记录
 * *******************************************************/
package handler

/*
import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"
)

//玩家获取官网开奖记录
func GetPk10Record(arg *pb.CPk10Record) interface{} {
	msg2 := new(pb.SPk10Record)
	list, err := data.Pk10RecordLogs(int(arg.Page), arg.Type)
	if err != nil {
		glog.Errorf("GetPk10Record err %v, %d", err, arg.Page)
		return msg2
	}
	for _, v := range list {
		msg3 := &pb.Pk10Record{
			Expect:   v.Expect,
			Opencode: v.Opencode,
			Opentime: v.Opentime,
		}
		msg2.List = append(msg2.List, msg3)
	}
	return msg2
}

//玩家获取战绩
func GetHuiYinRecords(arg *pb.CHuiYinRecords) interface{} {
	msg2 := new(pb.SHuiYinRecords)
	userid := arg.Userid
	page := arg.Page
	list, us, ms, err2 := data.GetRecords(userid, int(page))
	if err2 != nil {
		glog.Errorf("GetHuiYinRecords err %v, %s, %d", err2, userid, page)
		return msg2
	}
	for _, v := range list {
		d := &pb.HuiYinRecords{
			Gtype: v.Gametype,
			Rtype: v.Roomtype,
			Rest:  v.Rest,
		}
		//个人结果
		for _, val2 := range v.Details {
			d2 := &pb.HuiYinSeatRecords{
				Seat: val2.Seat,
				Bets: val2.Bets,
				Wins: val2.Wins,
			}
			d.Selfinfo = append(d.Selfinfo, d2)
		}
		//单个房间结果
		if val, ok := ms[v.Roomid]; ok {
			d.Expect = val.Expect
			d.Opencode = val.Opencode
			d.Opentime = val.Opentime
			d.Num = val.Num
			//开奖结果
			for _, v2 := range val.Trend {
				d2 := &pb.HuiYinRoomCards{
					Seat:  v2.Seat,
					Cards: v2.Cards,
					Rank:  v2.Rank,
					Point: v2.Point,
				}
				d.Cards = append(d.Cards, d2)
			}
			//全部结果
			for _, v2 := range val.Result {
				d2 := &pb.HuiYinWinRecords{
					Userid: v2.Userid,
					Bets:   v2.Bets,
					Wins:   v2.Wins,
				}
				d.Result = append(d.Result, d2)
			}
		}
		msg2.List = append(msg2.List, d)
	}
	//参与玩家信息
	for _, v := range us {
		d := &pb.HuiYinUserRecords{
			Userid:   v.Userid,
			Nickname: v.Nickname,
			Photo:    v.Photo,
		}
		msg2.Userlist = append(msg2.Userlist, d)
	}
	return msg2
}

//房间内获取输赢趋势
func GetHuiYinOpenResult(arg []*data.Trend) interface{} {
	msg2 := new(pb.SGetOpenResult)
	for _, v := range arg {
		d := &pb.RoomOpenResult{
			Expect:   v.Expect,
			Opencode: v.Opencode,
			Opentime: v.Opentime,
		}
		for _, v2 := range v.Result {
			d2 := &pb.OpenResult{
				Rank:  v2.Rank,
				Seat:  v2.Seat,
				Point: v2.Point,
				Cards: v2.Cards,
			}
			d.Info = append(d.Info, d2)
		}
		msg2.List = append(msg2.List, d)
	}
	return msg2
}

//房间内获取输赢趋势
func GetHuiYinTrends(arg []*data.Trend) interface{} {
	msg2 := new(pb.SGetTrend)
	for _, v := range arg {
		d := new(pb.HuiYinTrend)
		for _, v2 := range v.Result {
			d2 := &pb.TrendInfo{
				Rank: v2.Rank,
				Seat: v2.Seat,
			}
			d.Info = append(d.Info, d2)
		}
		msg2.List = append(msg2.List, d)
	}
	return msg2
}

//房间内获取上局赢家
func GetHuiYinWiners(arg []*data.Winer) interface{} {
	msg2 := new(pb.SGetLastWins)
	for _, v := range arg {
		d := &pb.LastWins{
			Userid:   v.Userid,
			Nickname: v.Nickname,
			Photo:    v.Photo,
			Chip:     v.Chip,
			Dealer:   v.Dealer,
		}
		msg2.List = append(msg2.List, d)
	}
	return msg2
}

//日志记录写入

//输赢趋势
func Pk10TrendLog(arg *pb.Pk10TrendLog) {
	record := &data.Trend{
		Expect:   arg.Expect,
		Opencode: arg.Opencode,
		Opentime: arg.Opentime,
	}
	for _, v := range arg.Result {
		d := data.TrendResult{
			Rank:  v.Rank,
			Seat:  v.Seat,
			Point: v.Point,
			Cards: v.Cards,
		}
		record.Result = append(record.Result, d)
	}
	record.Save()
}

//个人单局记录
func Pk10UseridLog(arg *pb.Pk10UseridLog) {
	record := &data.UserRecord{
		Roomid:      arg.Roomid,
		Gametype:    arg.Gametype,
		Roomtype:    arg.Roomtype,
		Lotterytype: arg.Lotterytype,
		Expect:      arg.Expect,
		Userid:      arg.Userid,
		Robot:       arg.Robot,
		Rest:        arg.Rest,
		Profits:     arg.Profits,
		Fee:         arg.Fee,
		Bets:        arg.Bets,
	}
	for _, v := range arg.Details {
		d := data.UseridDetails{
			Seat:   v.Seat,
			Bets:   v.Bets,
			Wins:   v.Wins,
			Refund: v.Refund,
		}
		record.Details = append(record.Details, d)
	}
	record.Save()
}

//房间单局记录日志
func Pk10GameLog(arg *pb.Pk10GameLog) {
	record := &data.GameRecord{
		Roomid:      arg.Roomid,
		Gametype:    arg.Gametype,
		Roomtype:    arg.Roomtype,
		Lotterytype: arg.Lotterytype,
		Expect:      arg.Expect,
		Opencode:    arg.Opencode,
		Opentime:    arg.Opentime,
		Num:         arg.Num,
		RobotFee:    arg.RobotFee,
		PlayerFee:   arg.PlayerFee,
		FeeNum:      arg.FeeNum,
		BetNum:      arg.BetNum,
		WinNum:      arg.WinNum,
		LoseNum:     arg.LoseNum,
		RefundNum:   arg.RefundNum,
	}
	for _, v := range arg.Trend {
		d := data.TrendResult{
			Rank:  v.Rank,
			Seat:  v.Seat,
			Point: v.Point,
			Cards: v.Cards,
		}
		record.Trend = append(record.Trend, d)
	}
	for _, v := range arg.Result {
		d := data.ResultRecord{
			Userid: v.Userid,
			Bets:   v.Bets,
			Wins:   v.Wins,
			Refund: v.Refund,
		}
		record.Result = append(record.Result, d)
	}
	for _, v := range arg.Record {
		d := data.FeeResult{
			Userid: v.Userid,
			Fee:    v.Fee,
		}
		record.Record = append(record.Record, d)
	}
	for _, v := range arg.Details {
		d := data.FeeDetails{
			Seat: v.Seat,
			Fee:  v.Fee,
		}
		for _, v2 := range v.Record {
			d2 := data.FeeResult{
				Userid: v2.Userid,
				Fee:    v2.Fee,
			}
			d.Record = append(d.Record, d2)
		}
		record.Details = append(record.Details, d)
	}
	record.Save()
}

//玩家获取盈亏统计
func GetHuiYinProfit(arg *pb.CHuiYinProfit) interface{} {
	stat := new(data.UserProfitStat)
	stat.Userid = arg.Userid
	stat.Get()
	msg2 := new(pb.SHuiYinProfit)
	msg2.Seven = stat.Seven
	msg2.Thirty = stat.Thirty
	msg2.All = stat.All
	return msg2
}
*/
