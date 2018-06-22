package data

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
	"gohappy/pb"
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
	selector := make(bson.M, 4)
	selector["profit"] = true
	selector["build"] = true
	selector["agent_level"] = true
	selector["address"] = true
	selector["_id"] = true
	q := bson.M{"agent_level": bson.M{"$gt": 0},
		"agent": bson.M{"$eq": arg.Userid},
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

//GetPlayerManage 玩家管理列表信息查询
func GetPlayerManage(arg *pb.CAgentPlayerManage) ([]bson.M, error) {
	if arg.Page == 0 {
		arg.Page = 1
	}
	pageSize := 20 //取前20条
	skipNum, sortFieldR := parsePageAndSort(int(arg.Page), pageSize, "coin", false)
	var list []bson.M
	selector := make(bson.M, 4)
	selector["coin"] = true
	selector["agent"] = true
	selector["agent_level"] = true
	selector["address"] = true
	selector["nickname"] = true
	selector["agent_name"] = true
	selector["agent_state"] = true
	selector["agent_join_time"] = true
	selector["_id"] = true
	q := bson.M{"agent": bson.M{"$eq": arg.Selfid},
		"agent_state": bson.M{"$eq": uint32(arg.State)}}
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