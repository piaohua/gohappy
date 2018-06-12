package data

import "github.com/AsynkronIT/protoactor-go/actor"

//DeskData 房间基础数据
type DeskData struct {
	Rid    string `json:"rid"`    //房间ID
	Unique string `json:"unique"` //配置表唯一ID
	Gtype  int32  `json:"gtype"`  //游戏类型
	Rtype  int32  `json:"rtype"`  //房间类型
	Dtype  int32  `json:"dtype"`  //桌子类型
	Ltype  int32  `json:"ltype"`  //彩票类型/房间等级类型
	Rname  string `json:"rname"`  //房间名字
	Count  uint32 `json:"count"`  //牌局人数限制
	Ante   uint32 `json:"ante"`   //底分
	Cost   uint32 `json:"cost"`   //抽佣百分比/创建消耗
	Vip    uint32 `json:"vip"`    //vip限制
	Chip   uint32 `json:"chip"`   //chip限制/进入房间限制
	//
	Deal  bool   `json:"deal"`  //房间是否可以上庄
	Carry uint32 `json:"carry"` //上庄携带限制
	Down  uint32 `json:"down"`  //下庄携带限制
	Top   uint32 `json:"top"`   //下庄最高携带限制
	//
	Sit   uint32 `json:"sit"`   //房间内坐下限制
	Ctime uint32 `json:"ctime"` //创建时间
	//
	Code    string `json:"code"`    //房间邀请码
	Cid     string `json:"cid"`     //房间创建人
	Expire  int64  `json:"expire"`  //牌局设定的过期时间
	Round   uint32 `json:"round"`   //牌局数
	Payment uint32 `json:"payment"` //付费方式1=AA or 0=房主支付
	//
	Minimum int64 `json:"minimum"` //房间最低限制
	Maximum int64 `json:"maximum"` //房间最高限制
	//
	Pub bool `json:"pub"` //公开显示
}

//DeskRole 牌桌玩家数据
type DeskRole struct {
	//数据
	*User
	//进程ID
	Pid *actor.PID
	//离线状态
	Offline bool
	//位置
	Seat uint32
}

//DeskSeat 牌桌位置数据
type DeskSeat struct {
	Userid   string   //玩家id
	Ready    bool     //是否准备
	BeDealer uint32   //1抢庄,2不抢,0等待抢
	DealerN  uint32   //抢庄倍数
	Bet      int64    //下注数量
	Cards    []uint32 //手牌
	Power    uint32   //牌力
	Vote     uint32   //投票解散1同意,2反对
	Niu      bool     //提交操作
}

//DeskGame 私人局牌桌当局数据
type DeskGame struct {
	Round      uint32   //私人房间打牌局数
	Cards      []uint32 //没摸起的海底牌
	BetNum     int64    //当前局下注总数/总底池
	Dealer     string   //庄家
	DealerSeat uint32   //庄家的座位
}

//DeskHua 金花数据, TODO 按轮记数
type DeskHua struct {
	//当前操作位置
	ActSeat uint32
	//当前操作值
	ActState int32
	//当前跟注额
	ActCallNum int64
	//当前最小加注额
	ActRaiseNum int64
	//第几轮
	ActTimes int32
	////当前操作的底池
	//ActPot int
	////底池,可多个池
	//ActPots map[int]*DeskPot
	//位置操作
	ActSeats map[uint32]*ActStatus
}

//ActStatus 位置操作状态
type ActStatus struct {
	See   bool //是否看牌
	Alive bool //是否还在
	//Allin bool //all in
	////本轮下注额
	ActNum int64
}

////DeskPot 底池
//type DeskPot struct {
//	//参与底池位置
//	PotSeats []uint32
//	//底池金额
//	PotNum int64
//}

//DeskPriv 私人局牌桌当局数据
type DeskPriv struct {
	VoteSeat uint32 //投票发起者座位号
	VoteTime int64  //结束投票时间
	//
	PrivScore map[string]int64  //私人局用户战绩积分
	Joins     map[string]uint32 //私人局用户参与次数
	//TODO Record 当局记录
}

//DeskFree 百人场牌桌当局数据
type DeskFree struct {
	Carry     int64               //庄家的携带
	DealerNum uint32              //做庄次数
	Dealers   map[string]int64    //上庄列表,userid: carry
	Cards     map[uint32][]uint32 //手牌
	Power     map[uint32]uint32   //牌力
	Bets      map[string]int64    //userid:num, 玩家下注金额
	SeatBets  map[uint32]int64    //seat:num, 位置总下注金额
	//位置下注详细
	SeatRoleBets map[uint32]map[string]int64
	//结果 seat:num,seat=(1,2,3,4,5),倍数
	Multiple map[uint32]int64
	Score1   map[uint32]int64 //位置(1-5)输赢总量
	Score2   map[string]int64 //每个闲家输赢总量
	//位置(1-5)上每个玩家输赢
	Score3 map[uint32]map[string]int64
	//Trend 输赢趋势
	Trends []*FreeTrend
	//上局赢家
	Winers []*FreeWiner
}

//FreeTrend 输赢趋势, //true 赢 false 输
type FreeTrend struct {
	Seat2 bool //天
	Seat3 bool //地
	Seat4 bool //玄
	Seat5 bool //黄
}

//FreeWiner 上局赢家
type FreeWiner struct {
	Userid   string
	Nickname string
	Photo    string
	Coin     int64 //赢利数量
}

//DeskBase 房间牌桌数据
type DeskBase struct {
	//基础数据
	*DeskData
	//进程ID
	Pid *actor.PID
	//房间人数
	Number uint32
}
