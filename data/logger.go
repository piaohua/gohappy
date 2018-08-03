package data

import (
	"time"

	"gohappy/pb"
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

//LogTask 任务日志
type LogTask struct {
	//Id       string    `bson:"_id"`
	Userid string    `bson:"userid"` //账户ID
	Taskid int32     `bson:"taskid"` //taskid
	Type   int32     `bson:"type"`   //task type
	Ctime  time.Time `bson:"ctime"`  //create Time
}

//Save 保存消息记录
func (t *LogTask) Save() bool {
	//t.Id = bson.NewObjectId().String()
	t.Ctime = bson.Now()
	return Insert(LogTasks, t)
}

//TaskRecord 任务记录
func TaskRecord(userid string, taskid, ttype int32) {
	record := &LogTask{
		Userid: userid,
		Taskid: taskid,
		Type:   ttype,
	}
	record.Save()
}

//LogProfit 代理收益日志
type LogProfit struct {
	//Id       string    `bson:"_id"`
	Agentid string    `bson:"agentid"` //代理ID,to
	Userid  string    `bson:"userid"`  //玩家ID,from
	Gtype   int32     `bson:"gtype"`   //game type
	Level   uint32    `bson:"level"`   //level type, 表示相对agentid的等级
	Rate    uint32    `bson:"rate"`    //rate
	Profit  int64     `bson:"profit"`  //Profit
	Type    int32     `bson:"type"`   //类型
	Ctime   time.Time `bson:"ctime"`   //create Time
}

//Save 保存消息记录
func (t *LogProfit) Save() bool {
	//t.Id = bson.NewObjectId().String()
	t.Ctime = bson.Now()
	return Insert(LogProfits, t)
}

//ProfitRecord 代理收益记录
func ProfitRecord(arg *pb.LogProfit) {
	record := &LogProfit{
		Userid:  arg.GetUserid(),
		Agentid: arg.GetAgentid(),
		Gtype:   arg.GetGtype(),
		Level:   arg.GetLevel(),
		Rate:    arg.GetRate(),
		Profit:  arg.GetProfit(),
		Type:    arg.GetType(),
	}
	record.Save()
	//添加天统计
	DayProfitRecord(arg)
}

//LogDayProfit 代理收益统计日志
type LogDayProfit struct {
	//Id       string    `bson:"_id"`
	Agentid string    `bson:"agentid"` //代理ID,to
	Userid  string    `bson:"userid"`  //玩家ID,from
	Nickname  string  `bson:"nickname"`  //昵称,from
	AgentNote string  `bson:"agent_note"` // 代理备注,from
	Day     int       `bson:"day"`     //day
	Profit  int64     `bson:"profit"`  //Profit
	ProfitFirst  int64     `bson:"profit_first"`  //ProfitFirst
	ProfitSecond  int64     `bson:"profit_second"`  //ProfitSecond
	ProfitMonth  int64     `bson:"profit_month"`  //ProfitMonth
	Utime   time.Time `bson:"utime"`   //update Time
	Ctime   time.Time `bson:"ctime"`   //create Time
}

//Save 保存消息记录
func (t *LogDayProfit) Save() bool {
	//t.Id = bson.NewObjectId().String()
	t.Ctime = bson.Now()
	return Insert(LogDayProfits, t)
}

func (t *LogDayProfit) Has() bool {
	return Has(LogDayProfits, bson.M{"userid": t.Userid, "agentid": t.Agentid, "day": t.Day})
}

func (t *LogDayProfit) Update(field string, profit int64) bool {
	if field == "" {
		return false
	}
	return Upsert(LogDayProfits, bson.M{"userid": t.Userid, "agentid": t.Agentid, "day": t.Day},
		bson.M{"$set": bson.M{"utime": t.Utime, "agent_note": t.AgentNote, "nickname": t.Nickname, "ctime": t.Utime},
		"$inc": bson.M{field: profit}})
}

//DayProfitRecord 代理收益统计记录
func DayProfitRecord(arg *pb.LogProfit) {
	record := &LogDayProfit{
		Userid: arg.GetUserid(),
		Agentid: arg.GetAgentid(),
		//Profit: arg.GetProfit(),
		Nickname: arg.GetNickname(),
		AgentNote: arg.GetAgentnote(),
	}
	record.Utime = bson.Now()
	record.Day = utils.Time2DayDate(record.Utime)
	var field string
	switch arg.GetType() {
	case int32(pb.LOG_TYPE53),
		int32(pb.LOG_TYPE54):
		field = "profit_month"
		record.ProfitMonth = arg.GetProfit()
	case int32(pb.LOG_TYPE52):
		switch arg.GetLevel() {
		case 1:
			field = "profit"
			record.Profit = arg.GetProfit()
		case 2:
			field = "profit_first"
			record.ProfitFirst = arg.GetProfit()
		case 3:
			field = "profit_second"
			record.ProfitSecond = arg.GetProfit()
		}
	}
	record.Update(field, arg.GetProfit())
	//if record.Has() {
	//	record.Update(field)
	//} else {
	//	record.Save()
	//}
}

//LogSysProfit 系统收益日志
type LogSysProfit struct {
	//Id       string    `bson:"_id"`
	Agentid string    `bson:"agentid"` //代理ID
	Userid  string    `bson:"userid"`  //玩家ID
	Gtype   int32     `bson:"gtype"`   //game type
	Level   uint32    `bson:"level"`   //level type, 表示userid等级
	Rate    uint32    `bson:"rate"`    //rate
	Profit  int64     `bson:"profit"`  //Profit
	Rest    int64     `bson:"rest"`    //Rest
	Ctime   time.Time `bson:"ctime"`   //create Time
}

//Save 保存消息记录
func (t *LogSysProfit) Save() bool {
	//t.Id = bson.NewObjectId().String()
	t.Ctime = bson.Now()
	return Insert(LogSysProfits, t)
}

//SysProfitRecord 系统收益记录
func SysProfitRecord(agentid, userid string, gtype int32, level, rate uint32, profit, rest int64) {
	record := &LogSysProfit{
		Userid:  userid,
		Agentid: agentid,
		Gtype:   gtype,
		Level:   level,
		Rate:    rate,
		Profit:  profit,
		Rest:    rest,
	}
	record.Save()
}

//LogBank 银行日志
type LogBank struct {
	//Id     string `bson:"_id"`
	Userid string    `bson:"userid"` //账户ID
	Type   int32     `bson:"type"`   //类型
	Num    int64     `bson:"num"`    //数量
	Rest   int64     `bson:"rest"`   //剩余数量
	From   string    `bson:"from"`   //赠送者
	Ctime  time.Time `bson:"ctime"`  //create Time
}

func (t *LogBank) Save() bool {
	//t.Id = bson.NewObjectId().String()
	t.Ctime = bson.Now()
	return Insert(LogBanks, t)
}

//银行记录
func BankRecord(userid, from string, rtype int32, rest, num int64) {
	record := &LogBank{
		Userid: userid,
		Type:   rtype,
		Num:    num,
		Rest:   rest,
		From:   from,
	}
	record.Save()
}

//GetBankLogs 获取银行操作记录
func GetBankLogs(arg *pb.CBankLog) ([]LogBank, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "ctime", false)
	var list []LogBank
	q := bson.M{"userid": arg.Userid}
	err := LogBanks.
		Find(q).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
	}
	return list, err
}
