package data

import (
	"time"

	"utils"

	"github.com/globalsign/mgo/bson"
)

//TODO 数据统计 玩家7日，30日，总赢亏

//1注册赠送,2开房消耗,3房间解散返还,
//4充值购买,5下注,7上庄，8下庄
//8下庄, 9后台操作,11破产补助
//18商城购买,19绑定赠送,20首充赠送
//23进入房间消耗
//24通比牛牛,25看牌抢庄,26牛牛抢庄
//27代理发放,28vip赠送
//38退款,39本金返还,40输赢,41坐庄输赢
//42异常退款,43反佣
//44机器人破产补助
//45庄家抽佣
const (
	LogType1  int32 = 1
	LogType2  int32 = 2
	LogType3  int32 = 3
	LogType4  int32 = 4
	LogType5  int32 = 5
	LogType7  int32 = 7
	LogType8  int32 = 8
	LogType9  int32 = 9
	LogType11 int32 = 11
	LogType18 int32 = 18
	LogType19 int32 = 19
	LogType20 int32 = 20
	LogType23 int32 = 23
	LogType24 int32 = 24
	LogType25 int32 = 25
	LogType26 int32 = 26
	LogType27 int32 = 27
	LogType28 int32 = 28
	LogType38 int32 = 38
	LogType39 int32 = 39
	LogType40 int32 = 40
	LogType41 int32 = 41
	LogType42 int32 = 42
	LogType43 int32 = 43
	LogType44 int32 = 44
	LogType45 int32 = 45
)

//注册日志
type LogRegist struct {
	//Id       string    `bson:"_id"`
	Userid   string    `bson:"userid"`    //账户ID
	Nickname string    `bson:"nickname"`  //账户名称
	Ip       string    `bson:"ip"`        //注册IP
	DayStamp time.Time `bson:"day_stamp"` //regist Time Today
	DayDate  int       `bson:"day_date"`  //regist day date
	Ctime    time.Time `bson:"ctime"`     //create Time
	Atype    uint32    `bson:"atype"`     //regist type
}

func (this *LogRegist) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.DayStamp = utils.TimestampTodayTime()
	this.DayDate = utils.DayDate()
	this.Ctime = bson.Now()
	return Insert(LogRegists, this)
}

//注册记录
func RegistRecord(userid, nickname, ip string, atype uint32) {
	record := &LogRegist{
		Userid:   userid,
		Nickname: nickname,
		Ip:       ip,
		Atype:    atype,
	}
	record.Save()
}

//登录日志
type LogLogin struct {
	//Id         string `bson:"_id"`
	Userid     string    `bson:"userid"`      //账户ID
	Event      int       `bson:"event"`       //事件：0=登录,1=正常退出,2＝系统关闭时被迫退出,3＝被动退出,4＝其它情况导致的退出
	Ip         string    `bson:"ip"`          //登录IP
	DayStamp   time.Time `bson:"day_stamp"`   //login Time Today
	LoginTime  time.Time `bson:"login_time"`  //login Time
	LogoutTime time.Time `bson:"logout_time"` //logout Time
	Atype      uint32    `bson:"atype"`       //login type
}

func (this *LogLogin) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.DayStamp = utils.TimestampTodayTime()
	this.LoginTime = bson.Now()
	return Insert(LogLogins, this)
}

func (this *LogLogin) Update(event int) bool {
	this.LogoutTime = bson.Now()
	return Update(LogLogins, bson.M{"userid": this.Userid, "event": 0},
		bson.M{"$set": bson.M{"event": event, "logout_time": this.LogoutTime}})
}

//登录记录
func LoginRecord(userid, ip string, atype uint32) {
	record := &LogLogin{
		Userid: userid,
		Event:  0,
		Ip:     ip,
		Atype:  atype,
	}
	record.Save()
}

//登录记录
func LogoutRecord(userid string, event int) {
	record := &LogLogin{
		Userid: userid,
	}
	record.Update(event)
}

//钻石日志
type LogDiamond struct {
	//Id     string `bson:"_id"`
	Userid string    `bson:"userid"` //账户ID
	Type   int32     `bson:"type"`   //类型
	Num    int64     `bson:"num"`    //数量
	Rest   int64     `bson:"rest"`   //剩余数量
	Ctime  time.Time `bson:"ctime"`  //create Time
}

func (this *LogDiamond) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.Ctime = bson.Now()
	return Insert(LogDiamonds, this)
}

//钻石记录
func DiamondRecord(userid string, rtype int32, rest, num int64) {
	record := &LogDiamond{
		Userid: userid,
		Type:   rtype,
		Num:    num,
		Rest:   rest,
	}
	record.Save()
}

//TODO 添加索引优化查询
//err := collection.EnsureIndex(index)
//金币日志
type LogCoin struct {
	//Id     string `bson:"_id"`
	Userid string    `bson:"userid"` //账户ID
	Type   int32     `bson:"type"`   //类型
	Num    int64     `bson:"num"`    //数量
	Rest   int64     `bson:"rest"`   //剩余数量
	Ctime  time.Time `bson:"ctime"`  //create Time
}

func (this *LogCoin) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.Ctime = bson.Now()
	return Insert(LogCoins, this)
}

//金币记录
func CoinRecord(userid string, rtype int32, rest, num int64) {
	record := &LogCoin{
		Userid: userid,
		Type:   rtype,
		Num:    num,
		Rest:   rest,
	}
	record.Save()
}

//房卡日志
type LogCard struct {
	//Id     string `bson:"_id"`
	Userid string    `bson:"userid"` //账户ID
	Type   int32     `bson:"type"`   //类型
	Num    int64     `bson:"num"`    //数量
	Rest   int64     `bson:"rest"`   //剩余数量
	Ctime  time.Time `bson:"ctime"`  //create Time
}

func (this *LogCard) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.Ctime = bson.Now()
	return Insert(LogCards, this)
}

//房卡记录
func CardRecord(userid string, rtype int32, rest, num int64) {
	record := &LogCard{
		Userid: userid,
		Type:   rtype,
		Num:    num,
		Rest:   rest,
	}
	record.Save()
}

//筹码日志
type LogChip struct {
	//Id     string `bson:"_id"`
	Userid string    `bson:"userid"` //账户ID
	Type   int32     `bson:"type"`   //类型
	Num    int64     `bson:"num"`    //数量
	Rest   int64     `bson:"rest"`   //剩余数量
	Ctime  time.Time `bson:"ctime"`  //create Time
}

func (this *LogChip) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.Ctime = bson.Now()
	return Insert(LogChips, this)
}

//筹码记录
func ChipRecord(userid string, rtype int32, rest, num int64) {
	record := &LogChip{
		Userid: userid,
		Type:   rtype,
		Num:    num,
		Rest:   rest,
	}
	record.Save()
}

//绑定日志
type LogBuildAgency struct {
	//Id       string `bson:"_id"`
	Userid    string    `bson:"userid"`     //账户ID
	Agent     string    `bson:"agent"`      //绑定ID
	DayStamp  time.Time `bson:"day_stamp"`  //regist Time Today
	Day       int       `bson:"day"`        //regist day
	Month     int       `bson:"month"`      //regist month
	Ctime     time.Time `bson:"ctime"`      //create Time
	DayDate   int       `bson:"day_date"`   //regist day date
	MonthDate int       `bson:"month_date"` //regist month date
}

func (this *LogBuildAgency) Save() bool {
	//this.Id = bson.NewObjectId().Hex()
	this.DayStamp = utils.TimestampTodayTime()
	this.DayDate = utils.DayDate()
	this.MonthDate = utils.MonthDate()
	this.Day = utils.Day()
	this.Month = int(utils.Month())
	this.Ctime = bson.Now()
	return Insert(LogBuildAgencys, this)
}

//绑定记录
func BuildRecord(userid, agent string) {
	record := &LogBuildAgency{
		Userid: userid,
		Agent:  agent,
	}
	record.Save()
}

//在线日志
type LogOnline struct {
	//Id       string `bson:"_id"`
	Num      int       `bson:"num"`       //online count
	DayStamp time.Time `bson:"day_stamp"` //Time Today
	Ctime    time.Time `bson:"ctime"`     //create Time
}

func (this *LogOnline) Save() bool {
	//this.Id = bson.NewObjectId().Hex()
	this.DayStamp = utils.TimestampTodayTime()
	this.Ctime = bson.Now()
	return Insert(LogOnlines, this)
}

//在线记录
func OnlineRecord(num int) {
	record := &LogOnline{
		Num: num,
	}
	record.Save()
}

//期号日志
type LogExpect struct {
	//Id     string `bson:"_id"`
	Expect    string    `bson:"expect"`     //期号ID
	Codes     string    `bson:"codes"`      //开奖号码
	OpenTimer int64     `bson:"open_timer"` //开奖时间
	Ctime     time.Time `bson:"ctime"`      //create Time
}

func (this *LogExpect) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.Ctime = bson.Now()
	return Insert(LogExpects, this)
}

//期号记录
func ExpectRecord(expect, codes string, openTimer int64) {
	record := &LogExpect{
		Expect:    expect,
		Codes:     codes,
		OpenTimer: openTimer,
	}
	record.Save()
}
