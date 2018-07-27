package data

import (
	"errors"
	"time"

	"gohappy/pb"
	"utils"

	"github.com/globalsign/mgo/bson"
)

//代理管理(代理ID为游戏内ID)
type Agency struct {
	Id         string    `bson:"_id" json:"id"`                  // AUTO_INCREMENT, PRIMARY KEY (`id`),
	UserName   string    `bson:"user_name" json:"user_name"`     // 用户名, UNIQUE KEY `user_name` (`user_name`)
	Password   string    `bson:"password" json:"password"`       // 密码
	Salt       string    `bson:"salt" json:"salt"`               // 密码盐
	Sex        int       `bson:"sex" json:"sex"`                 // 性别
	Email      string    `bson:"email" json:"email"`             // 邮箱
	LastLogin  time.Time `bson:"last_login" json:"last_login"`   // 最后登录时间
	LastIp     string    `bson:"last_ip" json:"last_ip"`         // 最后登录IP
	Status     int       `bson:"status" json:"status"`           // 状态，0正常 -1禁用
	CreateTime time.Time `bson:"create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `bson:"update_time" json:"update_time"` // 更新时间
	//RoleList   []Role    `bson:"role_list"`   // 角色列表
	//代理
	Phone    string    `bson:"phone" json:"phone"`         //绑定的手机号码(备用:非手机号注册时或多个手机时)
	Agent    string    `bson:"agent" json:"agent"`         //代理ID==Userid
	Level    int       `bson:"level" json:"level"`         //代理等级ID:1级,2级...
	Weixin   string    `bson:"weixin" json:"weixin"`       //微信ID
	Alipay   string    `bson:"alipay" json:"alipay"`       //支付宝ID
	QQ       string    `bson:"qq" json:"qq"`               //qq号码
	Address  string    `bson:"address" json:"address"`     //详细地址
	Number   uint32    `bson:"number" json:"number"`       //当前余额
	Expend   uint32    `bson:"expend" json:"expend"`       //总消耗
	Cash     float32   `bson:"cash" json:"cash"`           //当前可提取额
	Extract  float32   `bson:"extract" json:"extract"`     //已经提取额
	CashTime time.Time `bson:"cash_time" json:"cash_time"` //提取指定时间前所有
}

func (this *Agency) Get(userid string) {
	GetByQ(Agencys, bson.M{"agent": userid}, this)
}

func ExistAgency(safetycode string) bool {
	return Has(Agencys, bson.M{"agent": safetycode, "status": 0})
}

func ExistAgent(safetycode string) bool {
	return Has(PlayerUsers, bson.M{"_id": safetycode, "agent_state": 1})
}

//GetProfitRank 收益排行榜信息
func GetProfitRank() ([]bson.M, error) {
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(1, pageSize, "coin", false)
	var list []bson.M
	selector := make(bson.M, 4)
	selector["profit"] = true
	selector["nickname"] = true
	selector["address"] = true
	selector["_id"] = true
	q := bson.M{"profit": bson.M{"$gt": 0},
		"agent": bson.M{"$ne": ""}, "agent_state": bson.M{"$eq": 1}}
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

//GetAgentManage 代理管理列表信息查询
func GetAgentManage(arg *pb.CAgentManage) ([]bson.M, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "build", false)
	var list []bson.M
	selector := make(bson.M, 6)
	selector["profit_rate"] = true
	selector["profit"] = true
	selector["build"] = true
	selector["agent_level"] = true
	selector["address"] = true
	selector["_id"] = true
	q := bson.M{"agent_level": bson.M{"$gt": 0},
		"agent":       bson.M{"$eq": arg.Userid},
		"agent_state": bson.M{"$eq": 1}}
	if arg.Agentid != "" {
		q["_id"] = bson.M{"$eq": arg.Agentid}
	}
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

//GetAgentProfitManage 代理管理
func GetAgentProfitManage(arg *pb.CAgentProfitManage) ([]bson.M, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "build", false)
	var list []bson.M
	selector := make(bson.M, 5)
	selector["profit_rate"] = true
	selector["nickname"] = true
	selector["agent_level"] = true
	selector["agent_note"] = true
	selector["_id"] = true
	q := bson.M{"agent_level": bson.M{"$gt": 0},
		"agent":       bson.M{"$eq": arg.Userid},
		"agent_state": bson.M{"$eq": 1}}
	if arg.Agentid != "" {
		q["_id"] = bson.M{"$eq": arg.Agentid}
	}
	if arg.Agentnote != "" {
		q["agent_note"] = bson.M{"$eq": arg.GetAgentnote()}
	}
	if arg.GetRate() != 0 {
		q["profit_rate"] = bson.M{"$eq": arg.GetRate()}
	}
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

//GetPlayerManage 玩家管理列表信息查询
func GetPlayerManage(arg *pb.CAgentPlayerManage) ([]bson.M, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "coin", false)
	var list []bson.M
	selector := make(bson.M, 9)
	selector["coin"] = true
	selector["agent"] = true
	selector["agent_level"] = true
	selector["address"] = true
	selector["nickname"] = true
	selector["agent_name"] = true
	selector["agent_state"] = true
	selector["agent_join_time"] = true
	selector["_id"] = true
	//q := bson.M{"agent": bson.M{"$eq": arg.Selfid},
	//	"agent_state": bson.M{"$eq": uint32(arg.State)}}
	q := bson.M{"agent": bson.M{"$eq": arg.Selfid}}
	if arg.Userid != "" {
		q["_id"] = bson.M{"$eq": arg.Userid}
	}
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

//GetAgentProfit 代理收益明细
func GetAgentProfit(arg *pb.CAgentProfit) ([]LogProfit, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "ctime", false)
	var list []LogProfit
	q := bson.M{"agentid": arg.Agentid}
	if arg.GetLevel() != 0 {
		q["level"] = arg.GetLevel()
	}
	if arg.GetTime() != "" {
		q["ctime"] = bson.M{"$gte": utils.Str2Time(arg.GetTime()), "$lt": utils.Str2Time(arg.GetTime()).AddDate(0, 0, 1)}
	}
	err := LogProfits.
		Find(q).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
	}
	return list, err
}

//LogProfitOrder 提取收益订单
type LogProfitOrder struct {
	Id        string    `bson:"_id" json:"id"`                // AUTO_INCREMENT, PRIMARY KEY (`id`),
	Userid    string    `bson:"userid" json:"userid"`         // 玩家id(申请人)
	Agentid   string    `bson:"agentid" json:"agentid"`       // 代理id(受理人)
	Nickname  string    `bson:"nickname" json:"nickname"`     // 玩家昵称(申请人)
	Profit    int64     `bson:"profit" json:"profit"`         // 提取金额
	State     int32     `bson:"state" json:"state"`           // 状态,0等待处理,1成功,2失败
	ApplyTime time.Time `bson:"apply_time" json:"apply_time"` //提单时间
	ReplyTime time.Time `bson:"reply_time" json:"reply_time"` //响应时间
	Ctime     time.Time `bson:"ctime" json:"ctime"`           //记录生成时间
}

//Save 保存消息记录
func (t *LogProfitOrder) Save() bool {
	t.Id = bson.NewObjectId().String()
	t.Ctime = bson.Now()
	t.ApplyTime = bson.Now()
	return Insert(LogProfitsOrders, t)
}

//Update 更新记录
func (t *LogProfitOrder) Update(state int32) bool {
	t.ReplyTime = bson.Now()
	return Update(LogProfitsOrders, bson.M{"_id": t.Id},
		bson.M{"$set": bson.M{"state": t.State, "reply_time": t.ReplyTime}})
}

//Get 查询交易记录
func (t *LogProfitOrder) Get(orderid string) {
	Get(LogProfitsOrders, orderid, t)
}

//GetProfitOrder 收益提取订单
func GetProfitOrder(arg *pb.CAgentProfitOrder) ([]LogProfitOrder, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "ctime", false)
	var list []LogProfitOrder
	q := bson.M{"agentid": arg.Agentid}
	if arg.Type == 1 {
		q = bson.M{"userid": arg.Agentid}
	}
	err := LogProfitsOrders.
		Find(q).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
	}
	return list, err
}

// UserInfo store user information.
type UserInfo struct {
	Agentid       string    `bson:"agentid" json:"agentid,omitempty"`
	Subscribe     int       `bson:"subscribe" json:"subscribe,omitempty"`
	Language      string    `bson:"language" json:"language,omitempty"`
	OpenId        string    `bson:"openid" json:"openid,omitempty"` // nolint
	UnionId       string    `bson:"_id" json:"unionid,omitempty"`   // nolint
	Nickname      string    `bson:"nickname" json:"nickname,omitempty"`
	Sex           int       `bson:"sex" json:"sex,omitempty"`
	City          string    `bson:"city" json:"city,omitempty"`
	Country       string    `bson:"country" json:"country,omitempty"`
	Province      string    `bson:"province" json:"province,omitempty"`
	HeadImageUrl  string    `bson:"headimgurl" json:"headimgurl,omitempty"` // nolint
	SubscribeTime int64     `bson:"subscribe_time" json:"subscribe_time,omitempty"`
	Remark        string    `bson:"remark" json:"remark,omitempty"`
	GroupId       int       `bson:"groupid" json:"groupid,omitempty"` // nolint
	Ctime         time.Time `bson:"ctime" json:"ctime"`               //本条记录生成unix时间戳
}

//Has 判断记录是否存在
func (t *UserInfo) Has() bool {
	return Has(UserInfos, bson.M{"_id": t.UnionId})
}

//Get 获取一条记录
func (t *UserInfo) Get() {
	Get(UserInfos, t.UnionId, t)
}

//Update 更新一条记录
func (t *UserInfo) Update() bool {
	return Update(UserInfos, bson.M{"_id": t.UnionId}, t)
}

//GetAgentidByUnionid select agentid by unionid
func GetAgentidByUnionid(unionid string) string {
	var agentid string
	GetByQWithFields(UserInfos, bson.M{"_id": unionid}, []string{"agentid"}, &agentid)
	return agentid
}

//Save 写入新数据
func (t *UserInfo) Save() bool {
	if t.Has() {
		//return t.Update()
		return true
	}
	t.Ctime = bson.Now()
	return Insert(UserInfos, t)
}
