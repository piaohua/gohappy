package handler

import (
	"gohappy/data"
	"gohappy/game/config"
	"gohappy/pb"

	"github.com/globalsign/mgo/bson"
)

//TaskUpdateMsg 任务变动消息
func TaskUpdateMsg(num uint32, taskType pb.TaskType,
	userid string) (msg *pb.TaskUpdate) {
	msg = new(pb.TaskUpdate)
	msg.Userid = userid
	msg.Num = num
	msg.Type = taskType
	return
}

//SetTaskList 配置任务数据,测试数据
func SetTaskList() {
	NewTask(0, 0, 0, "完成新手指引", 1, 100, 1000)
	NewTask(1, 1, 2, "连续登录3天", 1, 100, 3000)
	NewTask(2, 1, 0, "连续登录7天", 1, 100, 7000)

	NewTask(3, 2, 4, "百人牛牛赢10场", 10, 100, 10000)
	NewTask(4, 2, 0, "百人牛牛赢20场", 20, 100, 20000)
	NewTask(5, 3, 0, "百人牛牛单局赢100000", 1, 100, 100000)
	NewTask(6, 4, 0, "百人牛牛单局赢200000", 1, 100, 200000)
	NewTask(7, 5, 0, "百人牛牛10次单局赢100000", 10, 100, 100)
	NewTask(8, 6, 0, "百人牛牛20次单局赢200000", 20, 100, 100)
	NewTask(9, 7, 10, "牛牛金币场赢10场", 10, 100, 100)
	NewTask(10, 7, 11, "牛牛金币场赢20场", 20, 100, 100)
	NewTask(11, 7, 0, "牛牛金币场赢30场", 30, 100, 100)
	NewTask(12, 8, 0, "百人牛牛押中牛牛", 1, 100, 100)
	NewTask(13, 9, 0, "百人牛牛押中五花", 1, 100, 100)
	NewTask(14, 10, 0, "百人牛牛押中炸弹", 1, 100, 100)
	NewTask(15, 11, 0, "百人牛牛押中五小", 1, 100, 100)

	NewTask(16, 12, 17, "百人三公赢10场", 10, 100, 100)
	NewTask(17, 12, 0, "百人三公赢20场", 20, 100, 100)
	NewTask(18, 13, 0, "百人三公单局赢100000", 1, 100, 100)
	NewTask(19, 14, 0, "百人三公单局赢200000", 1, 100, 100)
	NewTask(20, 15, 0, "百人三公10次单局赢100000", 10, 100, 100)
	NewTask(21, 16, 0, "百人三公20次单局赢200000", 20, 100, 100)
	NewTask(22, 17, 23, "三公金币场赢10场", 10, 100, 100)
	NewTask(23, 17, 24, "三公金币场赢20场", 20, 100, 100)
	NewTask(24, 17, 0, "三公金币场赢30场", 30, 100, 100)
	NewTask(25, 18, 0, "百人三公押中混三公", 1, 100, 100)
	NewTask(26, 19, 0, "百人三公押中小三公", 1, 100, 100)
	NewTask(27, 20, 0, "百人三公押中大三公", 1, 100, 100)
}

//NewTask 添加新任务
func NewTask(taskid, taskType, nextid int32, name string, count uint32,
	diamond, coin int64) {
	t := data.Task{
		ID:      bson.NewObjectId().Hex(),
		Ctime:   bson.Now(),
		Taskid:  taskid,
		Nextid:  nextid,
		Type:    taskType,
		Name:    name,
		Count:   count,
		//Diamond: diamond,
		Coin:    coin,
	}
	config.SetTask(t)
}
