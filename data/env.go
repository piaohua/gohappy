package data

import "github.com/globalsign/mgo/bson"

const (
	BANKRUPT      int64   = 0         //破产补助限制和金额
	BANKRUPT_TIME uint32  = 3         //每天破产补助次数
	DRAW_MONEY    int64   = 5000      //赠送最低金额限制
	DRAW_MONEY_LOW int64   = 1      //提现最低金额限制
	GIVE_PERCENT  float64 = 0.1       //赠送抽成
	GIVE_LIMIT    int64   = 100000000 //赠送上限
	TAX_NUMBER    int64   = 100       //小于这个数抽成为1
)

//设置变量
//key             value
const (
	ENV1  = "Regist_diamond"  //注册赠送钻石
	ENV2  = "Regist_coin"     //注册赠送金币
	ENV3  = "Regist_chip"     //注册赠送筹码
	ENV4  = "Regist_card"     //注册赠送房卡
	ENV5  = "Build"           //绑定赠送
	ENV6  = "First_pay_multi" //首充送n倍
	ENV7  = "First_pay_coin"  //首充送金币
	ENV8  = "Relieve"         //救济金次数
	ENV9  = "Prizedraw"       //转盘抽奖次数
	ENV10 = "Bankrupt_coin"   //破产金额
	ENV11 = "Relieve_coin"    //救济金额
	ENV12 = "Robot_num"       //虚假人数
	ENV13 = "Robot_allot1"    //机器人分配规则1
	ENV14 = "Robot_allot2"    //机器人分配规则2
	ENV15 = "Robot_bet"       //机器人下注AI
)

type Env struct {
	Key   string `bson:"_id" json:"key"`     //key
	Value int32  `bson:"value" json:"value"` //value
}

func GetEnvList() []Env {
	var list []Env
	ListByQ(Envs, nil, &list)
	return list
}

func (this *Env) DelEnv() bool {
	return Delete(Envs, this)
}

func (this *Env) SetEnv() bool {
	return Upsert(Envs, bson.M{"_id": this.Key}, this)
}
