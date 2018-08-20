package data

import (
	"errors"
	"time"

	"gohappy/pb"
	"gohappy/glog"
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
	skipNum, sortFieldR := parsePageAndSort(1, pageSize, "history_profit", false)
	var list []bson.M
	selector := make(bson.M, 4)
	selector["history_profit"] = true
	selector["nickname"] = true
	selector["address"] = true
	selector["_id"] = true
	q := bson.M{"history_profit": bson.M{"$gt": 0},
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
	selector := make(bson.M, 6)
	selector["bring_profit"] = true
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

//GetAgentProfitManage2 代理区域管理
func GetAgentProfitManage2(ids []string) ([]bson.M, error) {
	if len(ids) == 0 {
		return nil, errors.New("none record")
	}
	var list []bson.M
	selector := make(bson.M, 5)
	selector["agent_level"] = true
	selector["profit_rate"] = true
	selector["nickname"] = true
	selector["agent_note"] = true
	selector["_id"] = true
	q := bson.M{"_id": bson.M{"$in": ids}}
	err := PlayerUsers.
		Find(q).Select(selector).
		All(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("none record")
	}
	return list, nil
}

//GetAgentProfitManage3 代理区域管理
func GetAgentProfitManage3(arg *pb.CAgentProfitManage) ([]bson.M, error) {
	var list []bson.M
	q := bson.M{"agentid": arg.GetUserid()}
	if arg.Agentid != "" {
		q["userid"] = bson.M{"$eq": arg.Agentid}
	}
	if arg.Agentnote != "" {
		q["agent_note"] = bson.M{"$eq": arg.GetAgentnote()}
	}
	q = queryMonthProfit(arg.GetStartTime(), arg.GetEndTime(), q)
	list , err := agentDayProfitGroup3(q)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("none record")
	}
	return list, nil
}

//查询时间
func queryMonthProfit(startTime, endTime string, q bson.M) bson.M {
	if startTime == "" { //默认当月开始
		startTime = utils.Time2LocalStr(utils.Date(utils.Year(), int(utils.Month()), 1, 0, 0, 0, 0))
	}
	if startTime != "" && endTime != "" {
		q["day"] = bson.M{"$gte": utils.Time2DayDate(utils.Str2Time(startTime)),
			"$lt": utils.Time2DayDate(utils.Str2Time(endTime))}
		return q
	}
	if startTime != "" {
		q["day"] = bson.M{"$gte": utils.Time2DayDate(utils.Str2Time(startTime))}
		return q
	}
	if endTime != "" {
		q["day"] = bson.M{"$lt": utils.Time2DayDate(utils.Str2Time(endTime))}
		return q
	}
	return q
}

//分组统计, group by agentid userid
func agentDayProfitGroup3(match bson.M) (result []bson.M, err error) {
	m := bson.M{"$match": match}
	n := bson.M{
		"$group": bson.M{
			"_id": bson.M{"userid": "$userid", "agentid": "$agentid"},
			"profit_month": bson.M{
				"$sum": "$profit_month",
			},
		},
	}
	//统计
	operations := []bson.M{m, n}
	result = []bson.M{}
	pipe := LogDayProfits.Pipe(operations)
	err = pipe.All(&result)
	return
}

//GetPlayerManage 玩家管理列表信息查询
func GetPlayerManage(arg *pb.CAgentPlayerManage) ([]bson.M, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "coin", false)
	var list []bson.M
	selector := make(bson.M, 13)
	selector["ctime"] = true
	selector["coin"] = true
	selector["agent"] = true
	selector["agent_level"] = true
	selector["address"] = true
	selector["nickname"] = true
	selector["agent_name"] = true
	selector["agent_state"] = true
	selector["agent_join_time"] = true
	selector["login_time"] = true
	selector["agent_note"] = true
	selector["bring_profit"] = true
	selector["_id"] = true
	//q := bson.M{"agent": bson.M{"$eq": arg.Selfid},
	//	"agent_state": bson.M{"$eq": uint32(arg.State)}}
	q := bson.M{"agent": bson.M{"$eq": arg.Selfid}}
	if arg.Userid != "" {
		q["_id"] = bson.M{"$eq": arg.Userid}
	}
	if arg.GetAgentnote() != "" {
		q["agent_note"] = arg.GetAgentnote()
	}
	if arg.GetLevel() == 2 {
		q["agent_state"] = 1
		q["build_vaild"] = bson.M{"$gte": 3} //合格
	} else if arg.GetLevel() == 1 {
		q["agent_state"] = 0
	} else if arg.GetLevel() == 3 {
		q["agent_state"] = 1
		q["build_vaild"] = bson.M{"$lt": 3} //不合格
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

//GetAgentDayProfit 代理天收益明细
func GetAgentDayProfit(arg *pb.CAgentDayProfit) ([]LogDayProfit, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "ctime", false)
	var list []LogDayProfit
	q := getAgentDayProfitMatch(arg)
	//query by ctime
	if _, ok := q["day"]; ok {
		return getAgentDayProfitByTime(q)
	} else {
		return getAgentDayProfitByTime(q) //TODO 优化,暂时全部统计
	}
	err := LogDayProfits.
		Find(q).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&list)
	if err != nil {
	}
	return list, err
}

//query by time
func getAgentDayProfitByTime(match bson.M) ([]LogDayProfit, error) {
	glog.Debugf("match %#v", match)
	result, err2 := agentDayProfitGroup(match)
	glog.Debugf("result %#v, err2 %v", result, err2)
	if err2 != nil {
		return nil, err2
	}
	var list []LogDayProfit
	var ids []string
	for _, v := range result {
		l := LogDayProfit{}
		if val, ok := v["_id"]; ok {
			if m, ok := val.(bson.M); ok {
				if val2, ok := m["userid"]; ok {
					l.Userid = val2.(string)
				}
				if val2, ok := m["agentid"]; ok {
					l.Agentid = val2.(string)
				}
				if val2, ok := m["nickname"]; ok {
					l.Nickname = val2.(string)
				}
				if val2, ok := m["agent_note"]; ok {
					l.AgentNote = val2.(string)
				}
			}
		}
		if val, ok := v["profit"]; ok {
			l.Profit = utils.Int64(val)
		}
		if val, ok := v["profit_first"]; ok {
			l.ProfitFirst = utils.Int64(val)
		}
		if val, ok := v["profit_second"]; ok {
			l.ProfitSecond = utils.Int64(val)
		}
		if val, ok := v["profit_month"]; ok {
			l.ProfitMonth = utils.Int64(val)
		}
		list = append(list, l)
		ids = append(ids, l.Userid)
	}
	//添加昵称和备注
	list = queryAgentNote(ids, list)
	return list, nil
}

//查询条件
func getAgentDayProfitMatch(arg *pb.CAgentDayProfit) bson.M {
	q := bson.M{"agentid": arg.GetSelfid()}
	if arg.GetUserid() != "" {
		q["userid"] = arg.GetUserid()
	}
	if arg.GetAgentnote() != "" {
		q["agent_note"] = arg.GetAgentnote()
	}
	q = queryDay(arg.GetStartTime(), arg.GetEndTime(), q)
	return q
}

//分组统计, group by agentid userid
func agentDayProfitGroup(match bson.M) (result []bson.M, err error) {
	m := bson.M{"$match": match}
	n := bson.M{
		"$group": bson.M{
			//"_id": bson.M{"userid": "$userid", "agentid": "$agentid", "nickname": "$nickname", "agent_note": "$agent_note"},
			"_id": bson.M{"userid": "$userid", "agentid": "$agentid"},
			"profit": bson.M{
				"$sum": "$profit",
			},
			"profit_first": bson.M{
				"$sum": "$profit_first",
			},
			"profit_second": bson.M{
				"$sum": "$profit_second",
			},
			"profit_month": bson.M{
				"$sum": "$profit_month",
			},
		},
	}
	//统计
	operations := []bson.M{m, n}
	result = []bson.M{}
	pipe := LogDayProfits.Pipe(operations)
	err = pipe.All(&result)
	return
}

//AgentDayProfitGroup2 统计测试
func AgentDayProfitGroup2() (result []bson.M, err error) {
	return agentDayProfitGroup2(bson.M{})
}

//分组统计, group by agentid userid
func agentDayProfitGroup2(match bson.M) (result []bson.M, err error) {
	m := bson.M{"$match": match}
	n := bson.M{
		"$group": bson.M{
			"_id": bson.M{"agentid": "$agentid"},
			"profit": bson.M{
				"$sum": "$profit",
			},
			"profit_first": bson.M{
				"$sum": "$profit_first",
			},
			"profit_second": bson.M{
				"$sum": "$profit_second",
			},
		},
	}
	//统计
	operations := []bson.M{m, n}
	result = []bson.M{}
	pipe := LogDayProfits.Pipe(operations)
	err = pipe.All(&result)
	return
}

//GetAgentDayProfitCount 代理天收益明细
func GetAgentDayProfitCount(arg *pb.CAgentDayProfit) (int64, error) {
	q := getAgentDayProfitMatch(arg)
	return getAgentDayProfitCount2(q)
}

func getAgentDayProfitCount2(q bson.M) (int64, error) {
	result, err2 := agentDayProfitCount(q)
	glog.Debugf("result %#v, err2 %v", result, err2)
	if err2 != nil {
		return 0, err2
	}
	var num int64
	if val, ok := result["profit"]; ok {
		num += utils.Int64(val) //int64(val.(int))
	}
	if val, ok := result["profit_first"]; ok {
		num += utils.Int64(val) //int64(val.(int))
	}
	if val, ok := result["profit_second"]; ok {
		num += utils.Int64(val) //int64(val.(int))
	}
	//if val, ok := result["profit_month"]; ok {
	//	num += utils.Int64(val) //int64(val.(int))
	//}
	return num, nil
}

//分组统计
func agentDayProfitCount(match bson.M) (result bson.M, err error) {
	m := bson.M{"$match": match}
	n := bson.M{
		"$group": bson.M{
			"_id": 1,
			"profit": bson.M{
				"$sum": "$profit",
			},
			"profit_first": bson.M{
				"$sum": "$profit_first",
			},
			"profit_second": bson.M{
				"$sum": "$profit_second",
			},
			//"profit_month": bson.M{
			//	"$sum": "$profit_month",
			//},
		},
	}
	//统计
	operations := []bson.M{m, n}
	result = bson.M{}
	pipe := LogDayProfits.Pipe(operations)
	err = pipe.One(&result)
	return
}

//GetAgentDayProfitMonth 代理天区域收益明细
func GetAgentDayProfitMonth(arg *pb.CAgentProfitManage) (int64, error) {
	//if arg.Page == 0 {
	//	arg.Page = 1
	//}
	//pageSize := 20 //取前20条
	//skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "ctime", false)
	//var list []LogDayProfit
	q := bson.M{"agentid": arg.GetUserid()}
	if arg.GetAgentid() != "" {
		q["userid"] = arg.GetAgentid()
	}
	if arg.GetAgentnote() != "" {
		q["agent_note"] = arg.GetAgentnote()
	}
	q = queryDay(arg.GetStartTime(), arg.GetEndTime(), q)
	result, err2 := agentDayProfitMonthGroup(q)
	glog.Debugf("result %#v, err2 %v", result, err2)
	if err2 != nil {
		return 0, err2
	}
	if val, ok := result["profit_month"]; ok {
		return utils.Int64(val), nil //int64(val.(int)), nil
	}
	return 0, nil
}

//分组统计
func agentDayProfitMonthGroup(match bson.M) (result bson.M, err error) {
	m := bson.M{"$match": match}
	n := bson.M{
		"$group": bson.M{
			"_id": 1,
			"profit_month": bson.M{
				"$sum": "$profit_month",
			},
		},
	}
	//统计
	operations := []bson.M{m, n}
	result = bson.M{}
	pipe := LogDayProfits.Pipe(operations)
	err = pipe.One(&result)
	return
}

//查找昵称和备注
func queryAgentNote(ids []string, list []LogDayProfit) []LogDayProfit {
	glog.Debugf("ids %#v", ids)
	if len(ids) == 0 {
		return list
	}
	list2, err := queryAgentNote2(ids)
	glog.Debugf("list2 %#v, err %v", list2, err)
	if err != nil || len(list2) == 0 {
		return list
	}
	m := make(map[string]LogDayProfit)
	for _, v := range list2 {
		l := LogDayProfit{}
		if val, ok := v["_id"]; ok {
			l.Userid = val.(string)
		}
		if val, ok := v["nickname"]; ok {
			l.Nickname = val.(string)
		}
		if val, ok := v["agent_note"]; ok {
			l.AgentNote = val.(string)
		}
		if l.Userid == "" {
			continue
		}
		m[l.Userid] = l
	}
	for k, v := range list {
		if val, ok := m[v.Userid]; ok {
			v.Nickname = val.Nickname
			v.AgentNote = val.AgentNote
			list[k] = v
		}
	}
	return list
}

func queryAgentNote2(ids []string) (list []bson.M, err error) {
	if len(ids) == 0 {
		return
	}
	selector := make(bson.M, 3)
	selector["nickname"] = true
	selector["agent_note"] = true
	selector["_id"] = true
	q := bson.M{"_id": bson.M{"$in": ids}}
	err = PlayerUsers.
		Find(q).Select(selector).
		All(&list)
	if err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New("none record")
		return
	}
	return
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
	//t.Id = bson.NewObjectId().String()
	t.Id = ObjectIdString(bson.NewObjectId())
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
	q = queryTime(arg.GetStartTime(), arg.GetEndTime(), q)
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

//查询时间
func queryDay(startTime, endTime string, q bson.M) bson.M {
	if startTime != "" && endTime != "" {
		q["day"] = bson.M{"$gte": utils.Time2DayDate(utils.Str2Time(startTime)),
			"$lt": utils.Time2DayDate(utils.Str2Time(endTime))}
			return q
	}
	if startTime != "" {
		q["day"] = bson.M{"$gte": utils.Time2DayDate(utils.Str2Time(startTime))}
		return q
	}
	if endTime != "" {
		q["day"] = bson.M{"$lt": utils.Time2DayDate(utils.Str2Time(endTime))}
		return q
	}
	return q
}

func queryTime(startTime, endTime string, q bson.M) bson.M {
	if startTime != "" && endTime != "" {
		q["ctime"] = bson.M{"$gte": utils.Str2Time(startTime),
			"$lt": utils.Str2Time(endTime)}
		return q
	}
	if startTime != "" {
		q["ctime"] = bson.M{"$gte": utils.Str2Time(startTime)}
		return q
	}
	if endTime != "" {
		q["ctime"] = bson.M{"$lt": utils.Str2Time(endTime)}
		return q
	}
	return q
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
