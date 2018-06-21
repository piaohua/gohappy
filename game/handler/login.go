package handler

import (
	"gohappy/data"
	"gohappy/game/config"
	"gohappy/pb"
	"utils"

	"github.com/globalsign/mgo/bson"
)

//SetLoginPrize 连续登录处理
func SetLoginPrize(user *data.User) {
	//上次登录, TODO 全部领取完重置
	yesterDay := utils.Stamp2Time(utils.TimestampYesterday())
	if user.LoginTime.Before(yesterDay) {
		//隔天登录重置
		user.LoginTimes = (1 << 0)
		user.LoginPrize = 0
		return
	}
	//是否昨天登录过
	today := utils.TimestampTodayTime()
	if user.LoginTime.Before(today) {
		var i uint32
		for i = 0; i < 7; i++ {
			if (user.LoginTimes & (1 << i)) == 0 {
				user.LoginTimes |= (1 << i)
				break
			}
		}
	}
}

//GetLoginPrize 领取连续登录奖励
func GetLoginPrize(day uint32, user *data.User) (int64, int64, pb.ErrCode) {
	if (user.LoginPrize & (1 << day)) != 0 {
		return 0, 0, pb.AlreadyPrize
	}
	if (user.LoginTimes & (1 << day)) == 0 {
		return 0, 0, pb.AwardFaild
	}
	prize := config.GetLogin(day)
	if prize.Day != day {
		return 0, 0, pb.AwardFaild
	}
	user.LoginPrize |= (1 << day)
	return prize.Coin, prize.Diamond, pb.OK
}

//LoginPrizeInfo 获取连续登录信息
func LoginPrizeInfo(user *data.User) (msg []*pb.LoginPrize) {
	list := config.GetLogins()
	for _, v := range list {
		msg2 := new(pb.LoginPrize)
		msg2.Day = v.Day
		msg2.Coin = v.Coin
		msg2.Diamond = v.Diamond
		if (user.LoginPrize & (1 << v.Day)) != 0 {
			msg2.Status = pb.LoginPrizeGot
		} else if (user.LoginTimes & (1 << v.Day)) != 0 {
			msg2.Status = pb.LoginPrizeDone
		}
		msg = append(msg, msg2)
	}
	return
}

//LoginPrizeUpdateMsg 连续登录更新消息
func LoginPrizeUpdateMsg(user *data.User) (msg *pb.LoginPrizeUpdate) {
	msg = &pb.LoginPrizeUpdate{
		Userid:     user.GetUserid(),
		LoginTimes: user.LoginTimes,
		LoginPrize: user.LoginPrize,
		LoginTime:  utils.Time2Stamp(user.LoginTime),
		LoginIP:    user.LoginIp,
	}
	return
}

//SetLoginPrizeList 添加新任务
func SetLoginPrizeList() {
	var i uint32
	for i = 0; i < 7; i++ {
		t := data.LoginPrize{
			ID:    bson.NewObjectId().Hex(),
			Ctime: bson.Now(),
			//Diamond: 100 * int64((i + 1)),
			Coin: 2000 * int64((i + 1)),
			Day:  i,
		}
		config.SetLogin(t)
	}
}
