package main

import (
	"time"

	"gohappy/data"
	"gohappy/game/config"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家数据请求处理
func (rs *RoleActor) handlerUser(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CPing:
		arg := msg.(*pb.CPing)
		//glog.Debugf("CPing %#v", arg)
		rsp := handler.Ping(arg)
		rs.Send(rsp)
	case *pb.CNotice:
		arg := msg.(*pb.CNotice)
		glog.Debugf("CNotice %#v", arg)
		rsp := handler.GetNotice(data.NOTICE_TYPE1)
		rs.Send(rsp)
	case *pb.CGetCurrency:
		arg := msg.(*pb.CGetCurrency)
		glog.Debugf("CGetCurrency %#v", arg)
		//响应
		rsp := handler.GetCurrency(arg, rs.User)
		rs.Send(rsp)
	case *pb.CBuy:
		arg := msg.(*pb.CBuy)
		glog.Debugf("CBuy %#v", arg)
		//优化
		rsp, diamond, coin := handler.Buy(arg, rs.User)
		//同步兑换
		rs.addCurrency(diamond, coin, 0, 0, int32(pb.LOG_TYPE18))
		//响应
		rs.Send(rsp)
	case *pb.CShop:
		arg := msg.(*pb.CShop)
		glog.Debugf("CShop %#v", arg)
		//响应
		rsp := handler.Shop(arg, rs.User)
		rs.Send(rsp)
	case *pb.BankGive:
		arg := msg.(*pb.BankGive)
		glog.Debugf("BankGive %#v", arg)
		rs.addBank(arg.Coin, arg.Type)
	case *pb.CBank:
		arg := msg.(*pb.CBank)
		glog.Debugf("CBank %#v", arg)
		rs.bank(arg)
	case *pb.CRank:
		arg := msg.(*pb.CRank)
		glog.Debugf("CRank %#v", arg)
		rs.dbmsPid.Request(arg, ctx.Self())
	case *pb.TaskUpdate:
		arg := msg.(*pb.TaskUpdate)
		glog.Debugf("TaskUpdate %#v", arg)
		rs.taskUpdate(arg)
	case *pb.CTask:
		arg := msg.(*pb.CTask)
		glog.Debugf("CTask %#v", arg)
		rs.task()
	case *pb.CTaskPrize:
		arg := msg.(*pb.CTaskPrize)
		glog.Debugf("CTaskPrize %#v", arg)
		rs.taskPrize(arg.Type)
	case *pb.CUserData:
		arg := msg.(*pb.CUserData)
		glog.Debugf("CUserData %#v", arg)
		userid := arg.GetUserid()
		if userid == "" {
			userid = rs.User.GetUserid()
		}
		if userid != rs.User.GetUserid() {
			msg1 := new(pb.GetUserData)
			msg1.Userid = userid
			rs.rolePid.Request(msg1, ctx.Self())
		} else {
			//TODO 添加房间数据返回
			rsp := handler.GetUserDataMsg(arg, rs.User)
			rs.Send(rsp)
		}
	case *pb.GotUserData:
		arg := msg.(*pb.GotUserData)
		glog.Debugf("GotUserData %#v", arg)
		rsp := handler.UserDataMsg(arg)
		rs.Send(rsp)
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerPay(msg, ctx)
	}
}

/*
func (rs *RoleActor) addPrize(rtype, ltype, amount int32) {
	switch uint32(rtype) {
	case data.DIAMOND:
		rs.addCurrency(amount, 0, 0, 0, ltype)
	case data.COIN:
		rs.addCurrency(0, amount, 0, 0, ltype)
	case data.CARD:
		rs.addCurrency(0, 0, amount, 0, ltype)
	case data.CHIP:
		rs.addCurrency(0, 0, 0, amount, ltype)
	}
}

//消耗钻石
func (rs *RoleActor) expend(cost uint32, ltype int32) {
	diamond := -1 * int64(cost)
	rs.addCurrency(diamond, 0, 0, 0, ltype)
}
*/

//奖励发放
func (rs *RoleActor) addCurrency(diamond, coin, card, chip int64, ltype int32) {
	if rs.User == nil {
		glog.Errorf("add currency user err: %d", ltype)
		return
	}
	//日志记录
	if diamond < 0 && ((rs.User.GetDiamond() + diamond) < 0) {
		diamond = 0 - rs.User.GetDiamond()
	}
	if chip < 0 && ((rs.User.GetChip() + chip) < 0) {
		chip = 0 - rs.User.GetChip()
	}
	if coin < 0 && ((rs.User.GetCoin() + coin) < 0) {
		coin = 0 - rs.User.GetCoin()
	}
	if card < 0 && ((rs.User.GetCard() + card) < 0) {
		card = 0 - rs.User.GetCard()
	}
	rs.User.AddCurrency(diamond, coin, card, chip)
	//货币变更及时同步
	msg2 := handler.ChangeCurrencyMsg(diamond, coin,
		card, chip, ltype, rs.User.GetUserid())
	rs.rolePid.Tell(msg2)
	//消息
	msg := handler.PushCurrencyMsg(diamond, coin,
		card, chip, ltype)
	rs.Send(msg)
	//TODO 机器人不写日志
	//if rs.User.GetRobot() {
	//	return
	//}
	//rs.status = true
	//日志
	//TODO 日志放在dbms中统一写入
	//if diamond != 0 {
	//	msg1 := handler.LogDiamondMsg(diamond, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
	//if coin != 0 {
	//	msg1 := handler.LogCoinMsg(coin, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
	//if card != 0 {
	//	msg1 := handler.LogCardMsg(card, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
	//if chip != 0 {
	//	msg1 := handler.LogChipMsg(chip, ltype, rs.User)
	//	rs.loggerPid.Tell(msg1)
	//}
}

//同步数据
func (rs *RoleActor) syncUser() {
	if rs.User == nil {
		return
	}
	if rs.rolePid == nil {
		return
	}
	if !rs.status { //有变更才同步
		return
	}
	rs.status = false
	msg := new(pb.SyncUser)
	msg.Userid = rs.User.GetUserid()
	result, err := json.Marshal(rs.User)
	if err != nil {
		glog.Errorf("user %s Marshal err %v", rs.User.GetUserid(), err)
		return
	}
	msg.Data = result
	rs.rolePid.Tell(msg)
}

//银行发放
func (rs *RoleActor) addBank(coin int64, ltype int32) {
	if rs.User == nil {
		glog.Errorf("add addBank user err: %d", ltype)
		return
	}
	//日志记录
	if coin < 0 && ((rs.User.GetBank() + coin) < 0) {
		coin = 0 - rs.User.GetBank()
	}
	rs.User.AddBank(coin)
	//银行变动及时同步
	msg2 := handler.BankChangeMsg(coin,
		ltype, rs.User.GetUserid())
	rs.rolePid.Tell(msg2)
}

//1存入,2取出,3赠送
func (rs *RoleActor) bank(arg *pb.CBank) {
	msg := new(pb.SBank)
	rtype := arg.GetRtype()
	amount := int64(arg.GetAmount())
	userid := arg.GetUserid()
	coin := rs.User.GetCoin()
	switch rtype {
	case pb.BankDeposit: //存入
		if (coin - amount) < data.BANKRUPT {
			msg.Error = pb.NotEnoughCoin
		} else if amount <= 0 {
			msg.Error = pb.DepositNumberError
		} else {
			rs.addCurrency(0, -1*amount, 0, 0, int32(pb.LOG_TYPE12))
			rs.addBank(amount, int32(pb.LOG_TYPE12))
		}
	case pb.BankDraw: //取出
		if amount > rs.User.GetBank() {
			msg.Error = pb.NotEnoughCoin
		} else if amount < data.DRAW_MONEY {
			msg.Error = pb.DrawMoneyNumberError
		} else {
			rs.addCurrency(0, amount, 0, 0, int32(pb.LOG_TYPE13))
			rs.addBank(-1*amount, int32(pb.LOG_TYPE13))
		}
	case pb.BankGift: //赠送
		if amount > rs.User.GetBank() {
			msg.Error = pb.NotEnoughCoin
		} else if amount < data.DRAW_MONEY {
			msg.Error = pb.GiveNumberError
		} else if userid == "" {
			msg.Error = pb.GiveUseridError
		} else {
			msg1 := handler.GiveBankMsg(amount, int32(pb.LOG_TYPE15), userid)
			if rs.bank2give(msg1) {
				rs.addBank(-1*amount, int32(pb.LOG_TYPE15))
			} else {
				msg.Error = pb.GiveUseridError
			}
		}
	case pb.BankSelect: //查询
	}
	msg.Rtype = rtype
	msg.Amount = arg.GetAmount()
	msg.Userid = userid
	msg.Balance = rs.User.GetBank()
	rs.Send(msg)
}

//银行赠送
func (rs *RoleActor) bank2give(msg1 interface{}) bool {
	timeout := 3 * time.Second
	res1, err1 := rs.rolePid.RequestFuture(msg1, timeout).Result()
	if err1 != nil {
		glog.Errorf("logined GetUser failed: %v", err1)
		return false
	}
	if response1, ok := res1.(*pb.BankGiven); ok {
		if response1.Error == pb.OK {
			return true
		}
		glog.Errorf("BankGiven err %#v", response1)
		return false
	}
	return false
}

//任务信息,TODO next任务不显示和重置当日任务
func (rs *RoleActor) task() {
	msg := new(pb.STask)
	list := config.GetOrderTasks()
	m := make(map[int32]bool)
	for _, v := range list {
		if val, ok := rs.User.Task[v.Type]; ok {
			if val.Prize {
				continue
			}
			if val.Taskid != v.Taskid {
				continue
			}
		}
		if _, ok := m[v.Type]; ok {
			continue
		}
		msg2 := &pb.Task{
			Taskid:  v.Taskid,
			Type:    v.Type,
			Name:    v.Name,
			Count:   v.Count,
			Coin:    v.Coin,
			Diamond: v.Diamond,
		}
		if val, ok := rs.User.Task[v.Type]; ok {
			msg2.Num = val.Num
		}
		m[v.Type] = true
		msg.List = append(msg.List, msg2)
	}
	rs.Send(msg)
}

//任务奖励领取
func (rs *RoleActor) taskPrize(taskType int32) {
	msg := new(pb.STaskPrize)
	if val, ok := rs.User.Task[taskType]; ok {
		task := config.GetTask(val.Taskid)
		if val.Num < task.Count || task.Taskid != val.Taskid {
			msg.Error = pb.AwardFaild
			rs.Send(msg)
			return
		}
		//奖励发放
		rs.addCurrency(task.Diamond, task.Coin,
			0, 0, int32(pb.LOG_TYPE46))
		val.Prize = true
		//响应消息
		msg.Type = taskType
		msg.Coin = task.Coin
		msg.Diamond = task.Diamond
		//添加新任务
		rs.nextTask(taskType, task.Nextid, msg)
	} else {
		msg.Error = pb.AwardFaild
	}
	rs.Send(msg)
}

func (rs *RoleActor) nextTask(taskType, nextid int32, msg *pb.STaskPrize) {
	if nextid == 0 {
		return
	}
	//存在下个任务
	delete(rs.User.Task, taskType) //移除
	//TODO 任务完成日志
	msg2 := handler.TaskUpdateMsg(0, pb.TaskType(taskType),
		rs.User.GetUserid())
	msg2.Prize = true //移除标识
	rs.rolePid.Tell(msg2)
	//查找
	task := config.GetTask(nextid)
	if task.Taskid != nextid {
		return
	}
	msg.Next = &pb.Task{
		Taskid:  task.Taskid,
		Type:    task.Type,
		Name:    task.Name,
		Count:   task.Count,
		Coin:    task.Coin,
		Diamond: task.Diamond,
	}
	//添加新任务
	taskInfo := data.TaskInfo{
		Taskid: task.Taskid,
		Utime:  time.Now(),
	}
	rs.User.Task[int32(task.Type)] = taskInfo
	msg3 := handler.TaskUpdateMsg(0, pb.TaskType(task.Type),
		rs.User.GetUserid())
	msg3.Taskid = task.Taskid
	rs.rolePid.Tell(msg3)
}

//更新任务数据
func (rs *RoleActor) taskUpdate(arg *pb.TaskUpdate) {
	if val, ok := rs.User.Task[int32(arg.Type)]; ok {
		if val.Prize {
			return
		}
		//TODO 数值超出不再更新
		//task := config.GetTask(val.Taskid)
		//if val.Num >= task.Count {
		//	return
		//}
		val.Num += arg.Num
		val.Utime = time.Now()
		rs.User.Task[int32(arg.Type)] = val
		rs.rolePid.Tell(arg)
	} else {
		list := config.GetOrderTasks()
		for _, v := range list {
			if v.Type != int32(arg.Type) {
				continue
			}
			taskInfo := data.TaskInfo{
				Taskid: v.Taskid,
				Num:    arg.Num,
				Utime:  time.Now(),
			}
			rs.User.Task[int32(arg.Type)] = taskInfo
			rs.rolePid.Tell(arg)
			break
		}
	}
}
