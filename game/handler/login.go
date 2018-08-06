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
	//上次登录
	yesterDay := utils.Stamp2Time(utils.TimestampYesterday()).Local()
	if user.LoginTime.Before(yesterDay) {
		//隔天登录重置
		user.LoginTimes = (1 << 0)
		user.LoginPrize = 0
		return
	}
	//是否昨天登录过
	today := utils.TimestampTodayTime()
	if user.LoginTime.Before(today) {
		//全部领取完重置
		if user.LoginTimes == 127 && user.LoginPrize == 127 {
			user.LoginTimes = (1 << 0)
			user.LoginPrize = 0
			if user.LoginLoop == 3 {
				user.LoginLoop = 0
			} else {
				user.LoginLoop++
			}
			return
		}
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
	dayN := day + (user.LoginLoop * 7)
	prize := config.GetLogin(dayN)
	if prize.Day != dayN {
		return 0, 0, pb.AwardFaild
	}
	user.LoginPrize |= (1 << day)
	return prize.Coin, prize.Diamond, pb.OK
}

//LoginPrizeInfo 获取连续登录信息
func LoginPrizeInfo(user *data.User) (msg []*pb.LoginPrize) {
	list := config.GetLogins()
	for _, v := range list {
		day := (v.Day / 7)
		if day != user.LoginLoop {
			continue
		}
		msg2 := new(pb.LoginPrize)
		msg2.Day = day
		msg2.Coin = v.Coin
		msg2.Diamond = v.Diamond
		if (user.LoginPrize & (1 << day)) != 0 {
			msg2.Status = pb.LoginPrizeGot
		} else if (user.LoginTimes & (1 << day)) != 0 {
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
		LoginLoop:  user.LoginLoop,
		LoginTime:  utils.Time2Stamp(user.LoginTime),
		LoginIP:    user.LoginIP,
	}
	return
}

//SetLoginPrizeList 添加新任务
func SetLoginPrizeList() {
	var i uint32
	for i = 0; i < 28; i++ {
		num := (i % 7) * 50 + 300 //基本300,每天增加50
		if (i + 1) % 7 == 0 {
			num += 88 //第七天增加88
		}
		num += (i / 7) * 100 //每周增加100
		t := data.LoginPrize{
			//ID:      bson.NewObjectId().String(),
			ID:      data.ObjectIdString(bson.NewObjectId()),
			Ctime:   bson.Now(),
			//Diamond: 100 * int64((i + 1)),
			//Coin: 2000 * int64((i + 1)),
			Coin: int64(num),
			Day: i,
		}
		config.SetLogin(t)
		t.Save()
	}
}
