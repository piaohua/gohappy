package data

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Task 任务信息
type Task struct {
	ID      string    `bson:"_id" json:"id"`          //unique ID
	Taskid  int32     `bson:"taskid" json:"taskid"`   //unique
	Type    int32     `bson:"type" json:"type"`       //类型
	Name    string    `bson:"name" json:"name"`       //名称
	Count   uint32    `bson:"count" json:"count"`     //任务数值
	Diamond int64     `bson:"diamond" json:"diamond"` //钻石奖励
	Coin    int64     `bson:"coin" json:"coin"`       //金币奖励
	Today   bool      `bson:"today" json:"today"`     //是否当日任务
	Del     int       `bson:"del" json:"del"`         //是否移除
	Nextid  int32     `bson:"nextid" json:"nextid"`   //下个任务
	Ctime   time.Time `bson:"ctime" json:"ctime"`     //创建时间
}

//GetTaskList 任务
func GetTaskList() []Task {
	var list []Task
	ListByQ(Tasks, bson.M{"del": 0}, &list)
	return list
}

//TaskInfo 玩家的任务信息
type TaskInfo struct {
	Taskid int32     `bson:"taskid" json:"taskid"` //unique
	Prize  bool      `bson:"prize" json:"prize"`   //是否已领奖
	Num    uint32    `bson:"num" json:"num"`       //完成数值
	Utime  time.Time `bson:"utime" json:"utime"`   //更新时间
}
