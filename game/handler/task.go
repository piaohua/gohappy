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
	NewTask(28, 21, 29, "大厅房100局", 100, 5888, 0)
	NewTask(29, 21, 30, "大厅房200局", 200, 8888, 0)
	NewTask(30, 21, 31, "大厅房500局", 500, 18888, 0)
	NewTask(31, 21, 32, "大厅房1000局", 1000, 28888, 0)
	NewTask(32, 21, 33, "大厅房5000局", 5000, 58888, 0)
	NewTask(33, 21, 0, "大厅房10000局", 10000, 88888, 0)
	NewTask(34, 22, 0, "百人二八杠押中二八杠", 1, 100, 100)
	NewTask(35, 23, 0, "百人二八杠押中对子", 1, 100, 100)
	NewTask(36, 24, 0, "百人二八杠押中对十", 1, 100, 100)
}

//SetTaskList2 配置任务数据,测试数据
func SetTaskList2() {
	NewTask(37, 25, 38, "斗十匹配房玩牌200局", 200, 0, 8888)
	NewTask(38, 25, 39, "斗十匹配房玩牌500局", 500, 0, 18888)
	NewTask(39, 25, 40, "斗十匹配房玩牌1000局", 1000, 0, 28888)
	NewTask(40, 25, 0, "斗十匹配房玩牌5000局", 5000, 0, 58888)
}

//SetTaskList3 配置任务数据,测试数据
func SetTaskList3() {
	NewTask(41, 26, 42, "斗十大厅房玩牌200局", 200, 0, 8888)
	NewTask(42, 26, 43, "斗十大厅房玩牌500局", 500, 0, 18888)
	NewTask(43, 26, 44, "斗十大厅房玩牌1000局", 1000, 0, 28888)
	NewTask(44, 26, 0, "斗十大厅房玩牌5000局", 5000, 0, 58888)
	NewTask(45, 27, 46, "斗十百人场玩牌200局", 200, 0, 8888)
	NewTask(46, 27, 47, "斗十百人场玩牌500局", 500, 0, 18888)
	NewTask(47, 27, 48, "斗十百人场玩牌1000局", 1000, 0, 28888)
	NewTask(48, 27, 0, "斗十百人场玩牌5000局", 5000, 0, 58888)
}

//SetTaskList4 配置任务数据,测试数据
func SetTaskList4() {
	NewTask(49, 28, 50, "斗二八大厅房玩牌200局", 200, 0, 8888)
	NewTask(50, 28, 51, "斗二八大厅房玩牌500局", 500, 0, 18888)
	NewTask(51, 28, 52, "斗二八大厅房玩牌1000局", 1000, 0, 28888)
	NewTask(52, 28, 0, "斗二八大厅房玩牌5000局", 5000, 0, 58888)
	NewTask(53, 29, 54, "斗二八百人场玩牌200局", 200, 0, 8888)
	NewTask(54, 29, 55, "斗二八百人场玩牌500局", 500, 0, 18888)
	NewTask(55, 29, 56, "斗二八百人场玩牌1000局", 1000, 0, 28888)
	NewTask(56, 29, 0, "斗二八百人场玩牌5000局", 5000, 0, 58888)
	NewTask(57, 30, 58, "斗二八百人场玩牌200局", 200, 0, 8888)
	NewTask(58, 30, 59, "斗二八百人场玩牌500局", 500, 0, 18888)
	NewTask(59, 30, 60, "斗二八百人场玩牌1000局", 1000, 0, 28888)
	NewTask(60, 30, 0, "斗二八百人场玩牌5000局", 5000, 0, 58888)

}

//SetTaskList5 配置任务数据,测试数据
func SetTaskList5() {
	NewTask(61, 31, 62, "玩牌200局", 200, 0, 8888)
	NewTask(62, 31, 63, "玩牌500局", 500, 0, 18888)
	NewTask(63, 31, 64, "玩牌1000局", 1000, 0, 28888)
	NewTask(64, 31, 0, "玩牌5000局", 5000, 0, 58888)
}

//NewTask 添加新任务
func NewTask(taskid, taskType, nextid int32, name string, count uint32,
	diamond, coin int64) {
	t := data.Task{
		//ID:      bson.NewObjectId().String(),
		ID:      data.ObjectIdString(bson.NewObjectId()),
		Ctime:   bson.Now(),
		Taskid:  taskid,
		Nextid:  nextid,
		Type:    taskType,
		Name:    name,
		Count:   count,
		Diamond: diamond,
		Coin:    coin,
	}
	config.SetTask(t)
	t.Save()
}
