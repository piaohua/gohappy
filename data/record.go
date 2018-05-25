package data

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

//开奖结果记录
type Pk10Record struct {
	Expect        string    `bson:"_id"`
	Opencode      string    `bson:"opencode"`
	Opentime      string    `bson:"opentime"`
	Opentimestamp int64     `bson:"opentimestamp"`
	Code          string    `bson:"code"`
	Ctime         time.Time `bson:"ctime"`
}

func (this *Pk10Record) Save() bool {
	//this.Id = bson.NewObjectId().String()
	this.Ctime = bson.Now()
	return Insert(Pk10Records, this)
}

//开奖记录
func Pk10RecordLog(expect, opencode, opentime, code string,
	opentimestamp int64) {
	record := &Pk10Record{
		Expect:        expect,
		Opencode:      opencode,
		Opentime:      opentime,
		Opentimestamp: opentimestamp,
		Code:          code,
	}
	record.Save()
}

/*
//获取记录
func Pk10RecordLogs(page int, ctype uint32) ([]*Pk10Record, error) {
	pageSize := 10
	skipNum, sortFieldR := parsePageAndSort(page, pageSize, "_id", false)
	var list = make([]*Pk10Record, 0)
	var m bson.M
	switch ctype {
	case GAME_BJPK10:
		m = bson.M{"code": BJPK10}
	case GAME_MLAFT:
		m = bson.M{"code": MLAFT}
	}
	err := Pk10Records.
		Find(m).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
	}
	return list, err
}
*/

//房间单局记录,Roomid = GenOrderid()
//var roomid string = data.GenCporderid(roomid)
type GameRecord struct {
	Roomid      string         `bson:"_id"` //唯一
	Gametype    uint32         `bson:"gametype"`
	Roomtype    uint32         `bson:"roomtype"`
	Lotterytype uint32         `bson:"lotterytype"`
	Expect      string         `bson:"expect"`
	Opencode    string         `bson:"opencode"`
	Opentime    string         `bson:"opentime"`
	Num         uint32         `bson:"num"`        //参与人数
	RobotFee    int64          `bson:"robot_fee"`  //机器人抽佣数量
	PlayerFee   int64          `bson:"player_fee"` //玩家抽佣数量
	FeeNum      int64          `bson:"fee_num"`    //抽佣数量
	BetNum      int64          `bson:"bet_num"`    //下注总数量
	WinNum      int64          `bson:"win_num"`    //赢总数量
	LoseNum     int64          `bson:"lose_num"`   //输总数量
	RefundNum   int64          `bson:"refund_num"` //退款数量
	Trend       []TrendResult  `bson:"seats"`      //位置结果
	Result      []ResultRecord `bson:"result"`     //全部玩家输赢结果
	Record      []FeeResult    `bson:"record"`     //玩家抽佣明细
	Details     []FeeDetails   `bson:"details"`    //位置上玩家抽佣明细
	Ctime       time.Time      `bson:"ctime"`
}

//全部玩家信息
type ResultRecord struct {
	Userid string `bson:"userid"`
	Bets   int64  `bson:"bets"`   //下注总额
	Wins   int64  `bson:"wins"`   //输赢总额(不含本金)
	Refund int64  `bson:"refund"` //退款
}

//玩家抽佣明细
type FeeResult struct {
	Userid string `bson:"userid"`
	Fee    int64  `bson:"fee"` //抽佣数量
}

//位置抽佣明细
type FeeDetails struct {
	Seat   uint32      `bson:"seat"`   //位置
	Fee    int64       `bson:"fee"`    //位置抽佣数量
	Record []FeeResult `bson:"record"` //玩家抽佣明细
}

func (this *GameRecord) Save() bool {
	this.Ctime = bson.Now()
	return Insert(GameRecords, this)
}

func (this *GameRecord) Get() {
	Get(GameRecords, this.Roomid, this)
}

//获取记录
func GetGameRecords(list []*UserRecord) (map[string]*GameRecord, error) {
	var ls = make([]*GameRecord, 0)
	//查找
	in := make([]string, 0)
	for _, v := range list {
		in = append(in, v.Roomid)
	}
	ListByQ(GameRecords, bson.M{"_id": bson.M{"$in": in}}, &ls)
	if len(ls) == 0 {
		return nil, errors.New("query failed")
	}
	var rs = make(map[string]*GameRecord)
	for _, v := range ls {
		rs[v.Roomid] = v
	}
	return rs, nil
}

//个人单局记录
type UserRecord struct {
	//Id     string    `bson:"_id"`
	Roomid      string          `bson:"roomid"` //唯一
	Gametype    uint32          `bson:"gametype"`
	Roomtype    uint32          `bson:"roomtype"`
	Lotterytype uint32          `bson:"lotterytype"`
	Expect      string          `bson:"expect"`
	Userid      string          `bson:"userid"`
	Robot       bool            `bson:"robot"`   // 是否是机器人
	Rest        int64           `bson:"rest"`    //剩余
	Bets        int64           `bson:"bets"`    //下注额
	Profits     int64           `bson:"profits"` //输赢
	Fee         int64           `bson:"fee"`     //抽佣
	Details     []UseridDetails `bson:"details"`
	Ctime       time.Time       `bson:"ctime"`
}

//个人详细结果
type UseridDetails struct {
	Seat   uint32 `bson:"seat"`   //位置
	Bets   int64  `bson:"bets"`   //位置个人下注总额
	Wins   int64  `bson:"wins"`   //位置个人输赢总额(不含本金)
	Refund int64  `bson:"refund"` //退款
}

func (this *UserRecord) Save() bool {
	this.Ctime = bson.Now()
	return Insert(UserRecords, this)
}

func (this *UserRecord) Get() {
	Get(UserRecords, this.Roomid, this)
}

//获取记录
func GetUserRecords(userid string, page int) ([]*UserRecord, error) {
	pageSize := 20 //TODO 优化数据量过大情况
	skipNum, sortFieldR := parsePageAndSort(page, pageSize, "ctime", false)
	var list = make([]*UserRecord, 0)
	err := UserRecords.
		Find(bson.M{"userid": userid}).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("none record")
	}
	return list, nil
}

//全部参与玩家信息
type UserInfoRecords struct {
	Userid   string `bson:"_id" json:"userid"`        // 用户id
	Nickname string `bson:"nickname" json:"nickname"` // 用户昵称
	Photo    string `bson:"photo" json:"photo"`       // 头像
}

func GetUserInfoRecords(ms map[string]*GameRecord) ([]*UserInfoRecords, error) {
	in := make([]string, 0)
	us := make(map[string]bool)
	for _, v := range ms {
		for _, v2 := range v.Result {
			if _, ok := us[v2.Userid]; !ok {
				in = append(in, v2.Userid)
				us[v2.Userid] = true
			}
		}
	}
	if len(in) == 0 {
		return nil, errors.New("none players")
	}
	rs := make([]*UserInfoRecords, 0)
	fields := []string{"_id", "nickname", "photo"}
	ListByQWithFields(PlayerUsers, bson.M{"_id": bson.M{"$in": in}}, fields, &rs)
	return rs, nil
}

//获取记录
func GetRecords(userid string, page int) ([]*UserRecord,
	[]*UserInfoRecords, map[string]*GameRecord, error) {
	//玩家个人记录
	list, err := GetUserRecords(userid, page)
	if err != nil {
		return nil, nil, nil, err
	}
	//房间单条记录
	ms, err := GetGameRecords(list)
	if err != nil {
		return nil, nil, nil, err
	}
	//玩家显示基础数据
	us, err := GetUserInfoRecords(ms)
	if err != nil {
		return nil, nil, nil, err
	}
	return list, us, ms, nil
}

//输赢趋势
type Trend struct {
	//Id     string    `bson:"_id"`
	Expect   string        `bson:"expect"`   //期号
	Opencode string        `bson:"opencode"` //号码
	Opentime string        `bson:"opentime"` //开奖时间
	Result   []TrendResult `bson:"result"`   //结果
	Ctime    time.Time     `bson:"ctime"`    //
}

//开牌结果
type TrendResult struct {
	Rank  uint32   `bson:"rank"`  //排名(大小排行1->5)
	Seat  uint32   `bson:"seat"`  //位置(门内,第n门)
	Point uint32   `bson:"point"` //点数
	Cards []uint32 `bson:"cards"` //牌
}

func (this *Trend) Save() bool {
	this.Ctime = bson.Now()
	return Insert(Trends, this)
}

func (this *Trend) Get(id string) {
	Get(Trends, id, this)
}

//获取记录
func GetTrends(roomid string, page int) ([]*Trend, error) {
	pageSize := 10
	skipNum, sortFieldR := parsePageAndSort(page, pageSize, "ctime", false)
	var list = make([]*Trend, 0)
	err := Trends.
		Find(bson.M{"roomid": roomid}).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

//赢家
type Winer struct {
	Userid   string
	Nickname string
	Photo    string
	Chip     int64 //赢利数量
	Dealer   bool  //是否是庄家
}

//GetRank 排行榜信息
func GetRank() ([]bson.M, error) {
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(1, pageSize, "coin", false)
	var list []bson.M
	selector := make(bson.M, 4)
	selector["coin"] = true
	selector["nickname"] = true
	selector["photo"] = true
	selector["_id"] = true
	q := bson.M{"coin": bson.M{"$gt": 0}}
	err := PlayerUsers.
		Find(q).Select(selector).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("none record")
	}
	return list, nil
}
