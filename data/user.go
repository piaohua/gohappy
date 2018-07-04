package data

import (
	"time"

	"utils"

	"github.com/globalsign/mgo/bson"
)

type User struct {
	Userid   string `bson:"_id" json:"userid"`          // 用户id
	Nickname string `bson:"nickname" json:"nickname"`   // 用户昵称
	Photo    string `bson:"photo" json:"photo"`         // 头像
	Wxuid    string `bson:"wxuid" json:"wxuid"`         // 微信uid
	Sex      uint32 `bson:"sex" json:"sex"`             // 用户性别,男1 女2 非男非女3
	Phone    string `bson:"phone" json:"phone"`         // 绑定的手机号码
	Tourist  string `bson:"tourist" json:"tourist"`     // 游客
	Auth     string `bson:"auth" json:"auth"`           // 密码验证码
	Password string `bson:"password" json:"password"`   // MD5密码
	RegistIp string `bson:"regist_ip" json:"regist_ip"` // 注册账户时的IP地址
	LoginIp  string `bson:"login_ip" json:"login_ip"`   // 登录账户时的IP地址
	Diamond  int64  `bson:"diamond" json:"diamond"`     // 钻石
	Coin     int64  `bson:"coin" json:"coin"`           // 金币
	Chip     int64  `bson:"chip" json:"chip"`           // 筹码
	Card     int64  `bson:"card" json:"card"`           // 房卡
	Vip      uint32 `bson:"vip" json:"vip"`             // vip
	Status   uint32 `bson:"status" json:"status"`       // 正常1  锁定2  黑名单3
	Robot    bool   `bson:"robot" json:"robot"`         // 是否是机器人
	//战绩
	Win  uint32 `bson:"win" json:"win"`   // 赢
	Lost uint32 `bson:"lost" json:"lost"` // 输
	Ping uint32 `bson:"ping" json:"ping"` // 平
	//最高充值
	Money uint32 `bson:"money" json:"money"` // 充值总金额(分)
	//最高
	TopDiamonds int64 `bson:"top_diamonds" json:"top_diamonds"` // 最高拥有钻石总金额
	TopCoins    int64 `bson:"top_coins" json:"top_coins"`       // 最高拥有金币总金额
	TopChips    int64 `bson:"top_chips" json:"top_chips"`       // 最高拥有筹码总金额
	TopCards    int64 `bson:"top_cards" json:"top_cards"`       // 最高拥有房卡总数
	//单局
	TopWinDiamond int64 `bson:"top_win_diamond" json:"top_win_diamond"` // 单局赢最高钻石金额
	TopWinCoin    int64 `bson:"top_win_coin" json:"top_win_coin"`       // 单局赢最高金币金额
	TopWinChip    int64 `bson:"top_win_chip" json:"top_win_chip"`       // 单局赢最高筹码金额
	//代理, TODO 分佣设置，上级和上级分佣
	Agent            string    `bson:"agent" json:"agent"`                           // 绑定的代理ID
	Atime            time.Time `bson:"atime" json:"atime"`                           // 绑定代理时间
	AgentJoinTime    time.Time `bson:"agent_join_time" json:"agent_join_time"`       // 申请成为代理时间
	AgentState       uint32    `bson:"agent_state" json:"agent_state"`               // 是否是代理状态1通过
	AgentLevel       uint32    `bson:"agent_level" json:"agent_level"`               // 代理等级,1，2，3，4
	Build            uint32    `bson:"build" json:"build"`                           // 下属绑定数量
	AgentName        string    `bson:"agent_name" json:"agent_name"`                 // 代理名字
	RealName         string    `bson:"real_name" json:"real_name"`                   // 真实姓名
	Weixin           string    `bson:"weixin" json:"weixin"`                         // 微信
	ProfitRate       uint32    `bson:"profit_rate" json:"profit_rate"`               // 分佣比例
	Profit           int64     `bson:"profit" json:"profit"`                         // 收益
	WeekProfit       int64     `bson:"week_profit" json:"week_profit"`               // 本周收益
	WeekPlayerProfit int64     `bson:"week_player_profit" json:"week_player_profit"` // 本周玩家收益
	WeekStart        time.Time `bson:"week_start" json:"week_start"`                 // 每周日重置
	WeekEnd          time.Time `bson:"week_end" json:"week_end"`                     // 每周日重置
	HistoryProfit    int64     `bson:"history_profit" json:"history_profit"`         // 历史收益
	SubPlayerProfit  int64     `bson:"sub_player_profit" json:"sub_player_profit"`   // 下属玩家业绩收益
	SubAgentProfit   int64     `bson:"sub_agent_profit" json:"sub_agent_profit"`     // 下属代理业绩收益
	//时间
	Ctime     time.Time `bson:"ctime" json:"ctime"`           // 注册时间
	LoginTime time.Time `bson:"login_time" json:"login_time"` // 最后登录时间
	//银行
	Bank         int64  `bson:"bank" json:"bank"`                   // 个人银行
	BankPhone    string `bson:"bank_phone" json:"bank_phone"`       // 个人银行
	BankPassword string `bson:"bank_password" json:"bank_password"` // 个人银行
	//登录
	LoginTimes uint32 `bson:"login_times" json:"login_times"` //连续登录次数
	LoginPrize uint32 `bson:"login_prize" json:"login_prize"` //连续登录奖励
	Sign       string `bson:"sign" json:"sign"`               //个性签名
	//位置
	Lat     string `bson:"lat" json:"lat"`         //Latitude
	Lng     string `bson:"lng" json:"lng"`         //Longitude
	Address string `bson:"address" json:"address"` //Address
	//任务
	Task map[string]TaskInfo `bson:"task" json:"task"` // 已经完成或者还在继续的任务
}

// 数据库操作

func (this *User) Save() bool {
	return Upsert(PlayerUsers, bson.M{"_id": this.Userid}, this)
}

func (this *User) UpdateCurrency() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"diamond": this.Diamond, "coin": this.Coin,
			"chip": this.Chip, "card": this.Card}})
}

func (this *User) UpdateBank() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"bank": this.Bank}})
}

func (this *User) UpdateTask() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"task": this.Task}})
}

func (this *User) UpdateLogin() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"login_times": this.LoginTimes,
			"login_prize": this.LoginPrize, "login_time": this.LoginTime,
			"login_ip": this.LoginIp}})
}

func (this *User) UpdateSign() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"sign": this.Sign}})
}

func (this *User) UpdateAgentJoin() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"agent_join_time": this.AgentJoinTime,
			"agent_name": this.AgentName, "real_name": this.RealName,
			"weixin": this.Weixin, "agent_level": this.AgentLevel,
			"agent": this.Agent, "agent_state": this.AgentState}})
}

func (this *User) UpdateAgentProfit() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"week_profit": this.WeekProfit,
			"week_player_profit": this.WeekPlayerProfit,
			"history_profit":     this.HistoryProfit,
			"sub_player_profit":  this.SubPlayerProfit,
			"sub_agent_profit":   this.SubAgentProfit}})
}

func (this *User) UpdateAgentWeek() bool {
	return Update(PlayerUsers, bson.M{"_id": this.Userid},
		bson.M{"$set": bson.M{"week_start": this.WeekStart,
			"week_end": this.WeekEnd}})
}

func (this *User) Get() {
	Get(PlayerUsers, this.Userid, this)
}

func (this *User) GetById(userid string) {
	GetByQ(PlayerUsers, bson.M{"_id": userid}, this)
}

func (this *User) GetByPhone() {
	GetByQ(PlayerUsers, bson.M{"phone": this.Phone}, this)
}

func (this *User) GetByTourist() {
	GetByQ(PlayerUsers, bson.M{"tourist": this.Tourist}, this)
}

func (this *User) GetByWechat() {
	GetByQ(PlayerUsers, bson.M{"wxuid": this.Wxuid}, this)
}

//密码验证
func (this *User) VerifyPwd(pwd string) bool {
	return utils.Md5(pwd+this.Auth) == this.Password
}

// 非数据库操作

func (this *User) GetDiamond() int64 {
	return this.Diamond
}

func (this *User) AddDiamond(num int64) {
	this.Diamond += num
	if this.Diamond < 0 {
		this.Diamond = 0
	}
	if this.Diamond > this.TopDiamonds {
		this.TopDiamonds = this.Diamond
	}
	if num > this.TopWinDiamond {
		this.TopWinDiamond = num
	}
}

func (this *User) GetCoin() int64 {
	return this.Coin
}

func (this *User) AddCoin(num int64) {
	this.Coin += num
	if this.Coin < 0 {
		this.Coin = 0
	}
	if this.Coin > this.TopCoins {
		this.TopCoins = this.Coin
	}
	if num > this.TopWinCoin {
		this.TopWinCoin = num
	}
}

func (this *User) GetCard() int64 {
	return this.Card
}

func (this *User) AddCard(num int64) {
	this.Card += num
	if this.Card < 0 {
		this.Card = 0
	}
	if this.Card > this.TopCards {
		this.TopCards = this.Card
	}
}

func (this *User) GetChip() int64 {
	return this.Chip
}

func (this *User) AddChip(num int64) {
	this.Chip += num
	if this.Chip < 0 {
		this.Chip = 0
	}
	if this.Chip > this.TopChips {
		this.TopChips = this.Chip
	}
	if num > this.TopWinChip {
		this.TopWinChip = num
	}
}

func (this *User) GetVip() uint32 {
	return this.Vip
}

func (this *User) SetVip(num uint32) {
	this.Vip = num
}

func (this *User) AddMoney(num uint32) {
	this.Money += num
}

func (this *User) GetMoney() uint32 {
	return this.Money
}

func (this *User) GetUserid() string {
	return this.Userid
}

func (this *User) GetNickname() string {
	return this.Nickname
}

func (this *User) GetSex() uint32 {
	return this.Sex
}

func (this *User) GetPhoto() string {
	return this.Photo
}

func (this *User) GetWxuid() string {
	return this.Wxuid
}

func (this *User) GetPhone() string {
	return this.Phone
}

func (this *User) GetTourist() string {
	return this.Tourist
}

func (this *User) GetRobot() bool {
	return this.Robot
}

func (this *User) IsTourist() bool {
	if this.GetWxuid() != "" {
		return false
	}
	if this.GetPhone() != "" {
		return false
	}
	if this.GetTourist() != "" {
		return true
	}
	return false
}

func (this *User) GetTopDiamonds() int64 {
	return this.TopDiamonds
}

func (this *User) GetTopCoins() int64 {
	return this.TopCoins
}

func (this *User) GetTopChips() int64 {
	return this.TopChips
}

func (this *User) GetTopCards() int64 {
	return this.TopCards
}

func (this *User) GetTopWinDiamond() int64 {
	return this.TopWinDiamond
}

func (this *User) GetTopWinCoin() int64 {
	return this.TopWinCoin
}

func (this *User) GetTopWinChip() int64 {
	return this.TopWinChip
}

func (this *User) GetRegistTime() time.Time {
	return this.Ctime
}

func (this *User) GetLoginTime() time.Time {
	return this.LoginTime
}

func (this *User) SetAgent(agent string) {
	this.Agent = agent
	this.Atime = utils.LocalTime()
}

func (this *User) GetAgent() string {
	return this.Agent
}

func (this *User) SetRecord(value int32) {
	if value > 0 {
		this.Win++
	} else if value < 0 {
		this.Lost++
	} else {
		this.Ping++
	}
}

func (this *User) AddCurrency(diamond, coin, card, chip int64) {
	this.AddDiamond(diamond)
	this.AddCoin(coin)
	this.AddCard(card)
	this.AddChip(chip)
}

func (this *User) GetBank() int64 {
	return this.Bank
}

func (this *User) AddBank(num int64) {
	this.Bank += num
	if this.Bank < 0 {
		this.Bank = 0
	}
}

func (this *User) SetSign(content string) {
	this.Sign = content
}

func (this *User) GetSign() string {
	return this.Sign
}

func (this *User) AddProfit(isagent bool, num int64) {
	if num <= 0 {
		return
	}
	this.Profit += num
	this.HistoryProfit += num
	this.WeekProfit += num
	this.WeekPlayerProfit += num
	if isagent {
		this.SubAgentProfit += num
	} else {
		this.SubPlayerProfit += num
	}
}
