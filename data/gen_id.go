package data

import (
	"utils"

	"github.com/globalsign/mgo/bson"
)

//在同一个collection中
const (
	ROOMID_KEY = "last_room_id" //房间唯一id
	USERID_KEY = "last_user_id" //玩家唯一id
)

type IDGen struct {
	Id   string `bson:"_id"`  //key
	Curr string `bson:"curr"` //当前id
	Step string `bson:"step"` //step
	Max  string `bson:"max"`  //max
}

func (this *IDGen) Save() bool {
	//glog.Debugf("GenID %#v", this)
	return Upsert(IDGens, bson.M{"_id": this.Id}, this)
}

func (this *IDGen) Get() {
	Get(IDGens, this.Id, this)
}

func (this *IDGen) GenID() string {
	//glog.Debugf("GenID %s, %s, %s, %s", this.Id, this.Curr, this.Step, this.Max)
	this.Curr = utils.StringAdd(this.Curr)
	if this.Curr >= this.Max {
		this.Max = utils.StringAdd2(this.Curr, this.Step)
		this.Save()
	}
	//暂时每次存储
	//this.Save()
	return this.Curr
}

func (this *IDGen) Index() string {
	return this.Curr
}

//初始化
func InitIDGen(key string) (r *IDGen) {
	r = new(IDGen)
	r.Id = key
	r.Get()
	if r.Curr == "" {
		switch key {
		case ROOMID_KEY:
			r.Curr = "1"
			r.Step = "100"
			r.Max = "101"
		case USERID_KEY:
			r.Curr = "100000"
			r.Step = "100"
			r.Max = "100100"
		}
	} else {
		r.Curr = r.Max
	}
	return
}
