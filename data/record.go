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

//RoomRecord 创建房间记录
type RoomRecord struct {
	//Id     string    `bson:"_id"`
	Roomid string `bson:"roomid"` //唯一
	Gtype  int32  `bson:"gtype"`
	Rtype  int32  `bson:"rtype"`
	Dtype  int32  `bson:"dtype"`
	Rname  string `bson:"rname"`
	Count  uint32 `bson:"count"`
	Ante   uint32 `bson:"ante"`
	Code   string `bson:"code"`
	Round  uint32 `bson:"round"`
	Cid    string `bson:"cid"`
	Ctime  uint32 `bson:"ctime"`
}

//Save 保存记录
func (r *RoomRecord) Save() bool {
	return Insert(RoomRecords, r)
}

//RoleRecord 个人房间结果记录
type RoleRecord struct {
	//ID       string    `bson:"_id"`
	//Key      string    `bson:"key"`      //唯一
	Roomid   string    `bson:"roomid"` //房间id
	Gtype    int32     `bson:"gtype"`
	Userid   string    `bson:"userid"`   //玩家id
	Nickname string    `bson:"nickname"` //
	Photo    string    `bson:"photo"`    //
	Score    int64     `bson:"score"`    //输赢数量
	Rest     int64     `bson:"rest"`     //剩余
	Joins    uint32    `bson:"joins"`    //参与局数
	Ctime    time.Time `bson:"ctime"`
}

//Save 更新记录
func (r *RoleRecord) Save() bool {
	r.Ctime = bson.Now()
	q := bson.M{"roomid": r.Roomid,
		"userid": r.Userid}
	return Upsert(RoleRecords, q, r)
}

//RoundRecord 每局结算详情记录
//key=(私人场为房间id,匹配场生成唯一id)
//TODO 添加匹配场记录
type RoundRecord struct {
	//ID     string    `bson:"_id"`
	//Key    string    `bson:"key"`    //唯一
	Roomid string            `bson:"roomid"` //房间id
	Round  uint32            `bson:"round"`  //局数
	Dealer string            `bson:"dealer"` //庄家
	Roles  []RoundRoleRecord `bson:"roles"`  //局数
	Ctime  time.Time         `bson:"ctime"`
}

//RoundRoleRecord 每局详情记录
type RoundRoleRecord struct {
	Userid string   `bson:"userid"` //玩家id
	Cards  []uint32 `bson:"cards"`  //手牌
	Value  uint32   `bson:"value"`  //牌力
	Score  int64    `bson:"score"`  //输赢数量
	Bets   int64    `bson:"bets"`   //下注数量倍数
}

//Save 保存记录
func (r *RoundRecord) Save() bool {
	r.Ctime = bson.Now()
	return Insert(RoundRecords, r)
}

//getRoleRecords 获取个人房间结果记录
func getRoleRecords(userid string, gtype int32,
	page int) ([]string, error) {
	pageSize := 10 //TODO 优化数据量过大情况
	skipNum, sortFieldR := parsePageAndSort(page, pageSize, "ctime", false)
	var list []bson.M
	selector := bson.M{"roomid": true}
	err := RoleRecords.
		Find(bson.M{"userid": userid, "gtype": gtype}).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		Select(selector).
		All(&list)
	if err != nil {
		return nil, err
	}
	in := make([]string, 0)
	for _, v := range list {
		if val, ok := v["roomid"]; ok {
			in = append(in, val.(string))
		}
	}
	if len(in) == 0 {
		return nil, errors.New("none record")
	}
	return in, nil
}

//GetRoomRecords 获取私人房间记录
func GetRoomRecords(userid string, gtype int32, page int) ([]*RoomRecord,
	[]*RoundRecord, []*RoleRecord, error) {
	in, err := getRoleRecords(userid, gtype, page)
	if err != nil {
		return nil, nil, nil, err
	}
	//
	var ls = make([]*RoomRecord, 0)
	ListByQ(RoomRecords, bson.M{"roomid": bson.M{"$in": in}}, &ls)
	if len(ls) == 0 {
		return nil, nil, nil, errors.New("query failed")
	}
	//
	var ds = make([]*RoundRecord, 0)
	ListByQ(RoundRecords, bson.M{"roomid": bson.M{"$in": in}}, &ds)
	if len(ds) == 0 {
		return nil, nil, nil, errors.New("query failed")
	}
	//
	var us = make([]*RoleRecord, 0)
	ListByQ(RoleRecords, bson.M{"roomid": bson.M{"$in": in}}, &us)
	if len(us) == 0 {
		return nil, nil, nil, errors.New("query failed")
	}
	return ls, ds, us, nil
}
