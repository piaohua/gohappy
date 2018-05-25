package data

import "time"

//statistics

//每日赢亏统计
type ProfitStat struct {
	//Id     string    `bson:"_id"`
	Userid    string    `bson:"userid"`
	Robot     bool      `bson:"robot"` // 是否是机器人
	Day       int       `bson:"day"`   //20180205
	Month     int       `bson:"month"` //201802
	Yesterday int64     `bson:"yesterday"`
	Seven     int64     `bson:"seven"`
	Thirty    int64     `bson:"thirty"`
	DayStamp  time.Time `bson:"day_stamp"` //Time Today
	Ctime     time.Time `bson:"ctime"`
}

//赢亏统计,按玩家单个统计
type UserProfitStat struct {
	Userid    string    `bson:"_id"`
	Robot     bool      `bson:"robot"` // 是否是机器人
	Yesterday int64     `bson:"yesterday"`
	Seven     int64     `bson:"seven"`
	Thirty    int64     `bson:"thirty"`
	All       int64     `bson:"all"`
	DayStamp  time.Time `bson:"day_stamp"` //Time Today
	Utime     time.Time `bson:"utime"`     //update time
	Ctime     time.Time `bson:"ctime"`
}

func (this *UserProfitStat) Get() {
	Get(UserStatRecords, this.Userid, this)
}
