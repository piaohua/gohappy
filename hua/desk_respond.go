/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-04-30 18:02:06
 * Filename      : desk_respond.go
 * Description   : 响应请求
 * *******************************************************/
package main

import (
	"gohappy/data"
	"gohappy/game/algo"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// external function

//' 打印
func (t *Desk) printOver() {
	glog.Debugf("game over data -> %#v", t.DeskData)
	glog.Debugf("game over roles -> %d", len(t.roles))
	glog.Debugf("game over router -> %d", len(t.router))
	glog.Debugf("game over state -> %d", t.state)
	glog.Debugf("game over dealer -> %s", t.DeskGame.Dealer)
	glog.Debugf("game over dealer seat -> %d", t.DeskGame.DealerSeat)
	for k, v := range t.roles {
		glog.Debugf("game over userid %s -> %d", k, v.Seat)
	}
	for k, v := range t.seats {
		glog.Debugf("game over seat %d -> %s", k, v.Userid)
	}
	if t.isFree() {
		glog.Debugf("game over Carry -> %d", t.DeskFree.Carry)
		glog.Debugf("game over DealerNum -> %d", t.DeskFree.DealerNum)
		glog.Debugf("game over Cards -> %#v", t.DeskFree.Cards)
		glog.Debugf("game over Bets -> %#v", t.DeskFree.Bets)
		glog.Debugf("game over SeatBets -> %#v", t.DeskFree.SeatBets)
		glog.Debugf("game over Multiple -> %#v", t.DeskFree.Multiple)
		glog.Debugf("game over Score1 -> %#v", t.DeskFree.Score1)
		glog.Debugf("game over Score2 -> %#v", t.DeskFree.Score2)
		glog.Debugf("game over Score3 -> %#v", t.DeskFree.Score3)
	}
}

//.

//' 聊天
func (t *Desk) chatText(arg *pb.CChatText, ctx actor.Context) {
	userid := t.getRouter(ctx)
	seat := t.getSeat(userid)
	glog.Debugf("CChatText %s, %d", userid, seat)
	//房间消息广播,聊天
	t.broadcast(handler.ChatTextMsg(seat, userid, arg.Content))
}

func (t *Desk) chatVoice(arg *pb.CChatVoice, ctx actor.Context) {
	userid := t.getRouter(ctx)
	seat := t.getSeat(userid)
	glog.Debugf("CChatVoice %s, %d", userid, seat)
	//房间消息广播,聊天
	t.broadcast(handler.ChatVoiceMsg(seat, userid, arg.Content))
}

//.

//' enter 进入桌子
func (t *Desk) enter(user *data.User, pid *actor.PID) pb.ErrCode {
	errcode := t.enterCheck(user)
	if errcode != pb.OK {
		return errcode
	}
	//设置路由
	t.router[pid.String()] = user.GetUserid()
	//TODO 已经在房间防止数据覆盖
	if v, ok := t.roles[user.GetUserid()]; ok {
		v.Offline = false
		v.Pid = pid
		//在线消息
		t.offlineMsg(user.GetUserid())
		return pb.AlreadyInRoom //已经在房间内
	}
	//加入游戏
	p := new(data.DeskRole)
	p.User = user
	p.Pid = pid
	//私人,自由场分配位置
	t.enterSeat(p)
	t.roles[user.GetUserid()] = p
	return pb.OK
}

//进入桌子分配位置
func (t *Desk) enterSeat(p *data.DeskRole) {
	if p.Seat != 0 {
		return
	}
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE2): //百人
		return
	}
	var i uint32
	for i = 1; i <= t.DeskData.Count; i++ {
		if _, ok := t.seats[i]; !ok {
			p.Seat = i
			t.seats[i] = &data.DeskSeat{
				Userid: p.User.GetUserid(),
			}
			break
		}
	}
}

//TODO 不同类型房间进入限制
func (t *Desk) enterCheck(user *data.User) pb.ErrCode {
	//if user.GetVip() < t.DeskData.Vip {
	//	return pb.VipTooLow
	//}
	//if user.GetChip() < int64(t.DeskData.Chip) {
	//	return pb.ChipNotEnough
	//}
	if _, ok := t.roles[user.GetUserid()]; !ok {
		switch t.DeskData.Rtype {
		case int32(pb.ROOM_TYPE1): //私人
		//可以旁观
		default:
			if uint32(len(t.roles)) >= t.DeskData.Count {
				return pb.RoomFull //人数已满
			}
		}
	}
	//消耗
	if t.isAADesk() && user.GetUserid() != t.DeskData.Cid {
		if user.GetDiamond() < int64(t.DeskData.Cost) {
			return pb.NotEnoughDiamond
		}
	}
	return pb.OK
}

//.

//' 离开房间处理

// nnLeave 玩家主动离开,TODO 下注也可以离开
func (t *Desk) nnLeave(userid string, ctx actor.Context) {
	//离线
	defer t.offlineDesk(userid)
	//seat := t.getSeat(userid)
	errcode := t.leave(userid)
	if errcode == pb.OK {
		//清除数据
		t.userLeaveDesk(userid)
		return
	}
	msg := new(pb.SJHLeave)
	//TODO 暂时不返回错误
	//msg.Error = errcode
	ctx.Respond(msg)
}

// leave 玩家离开,TODO 下注也可以离开
func (t *Desk) leave(userid string) pb.ErrCode {
	//离线处理
	defer t.userLeaveDeskMsg(userid)
	if _, ok := t.roles[userid]; !ok {
		return pb.NotInRoom
	}
	//离开房间条件
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		switch t.state {
		case int32(pb.STATE_READY):
		default:
			seat := t.getSeat(userid)
			if v, ok := t.seats[seat]; ok {
				if v.Ready {
					//TODO 主动放弃,离开房间
					return pb.GameStartedCannotLeave
				}
			}
		}
	case int32(pb.ROOM_TYPE1): //私人
		if t.DeskPriv.VoteSeat != 0 {
			return pb.VotingCantLaunchVote
		}
		switch t.state {
		case int32(pb.STATE_READY):
		default:
			seat := t.getSeat(userid)
			if v, ok := t.seats[seat]; ok {
				if v.Ready {
					return pb.GameStartedCannotLeave
				}
			}
		}
	case int32(pb.ROOM_TYPE2): //百人
		//清空上庄列表
		t.leaveBeDealer(userid)
		switch t.state {
		case int32(pb.STATE_BET): //下注中
			if _, ok := t.DeskFree.Bets[userid]; ok {
				return pb.GameStartedCannotLeave
			}
			if userid == t.DeskGame.Dealer {
				return pb.GameStartedCannotLeave
			}
		default:
			//庄家下庄
			user := t.getPlayer(userid)
			t.delBeDealer(userid, user)
		}
	}
	return pb.OK
}

//离开状态消息
func (t *Desk) userLeaveDeskMsg(userid string) {
	if v, ok := t.roles[userid]; ok {
		if !v.Offline {
			msg2 := new(pb.LeftDesk)
			v.Pid.Tell(msg2)
		}
		//
		//msg := new(pb.LeaveDesk)
		//msg.Userid = userid
		//msg.Roomid = t.DeskData.Rid
		//离开房间
		//nodePid.Tell(msg)
		//离开房间,TODO 玩家进程中操作,一致性
		//t.roomPid.Request(msg, t.selfPid)
	}
}

//离开状态消息
func (t *Desk) userLeaveDesk(userid string) {
	if v, ok := t.roles[userid]; ok {
		//离开状态消息
		msg1 := new(pb.SJHLeave)
		msg1.Seat = v.Seat
		msg1.Userid = userid
		t.broadcast(msg1)
		//msg2 := new(pb.LeftDesk)
		//v.Pid.Tell(msg2)
		//清除数据
		delete(t.seats, v.Seat)
		delete(t.roles, userid)
		//
		msg := new(pb.LeaveDesk)
		msg.Userid = userid
		msg.Roomid = t.DeskData.Rid
		//离开房间
		nodePid.Tell(msg)
		//离开房间,TODO 玩家进程中操作,一致性
		t.roomPid.Request(msg, t.selfPid)
	}
}

func (t *Desk) offlineMsg(userid string) {
	if v, ok := t.roles[userid]; ok && v != nil {
		msg := new(pb.SJHPushOffline)
		msg.Seat = v.Seat
		msg.Userid = userid
		msg.Offline = v.Offline
		t.broadcast2(userid, msg)
	}
}

//.

//' 庄家操作

//暂时同一个玩家上庄列表只能有一个
//上庄规则是最高携带优先

//上庄或补庄处理，st:0下庄 1上庄 2补庄
func (t *Desk) addBeDealer(userid string, st int32,
	num int64, user *data.User) {
	switch st {
	case int32(pb.DEALER_BU): //庄家补庄
	case int32(pb.DEALER_UP): //庄家上庄
	default:
		return
	}
	if userid == t.DeskGame.Dealer {
		t.DeskFree.Carry += num
	} else {
		t.DeskFree.Dealers[userid] += num
	}
	msg := handler.BeJHDealerMsg(st, num, t.DeskGame.Dealer,
		userid, user.GetNickname())
	t.broadcast(msg)
}

//下庄处理
func (t *Desk) delBeDealer(userid string, user *data.User) {
	if !t.isFreeDealer() {
		return
	}
	if t.DeskGame.Dealer != userid {
		return
	}
	var num int64
	num = t.DeskFree.Carry
	t.sendCoin(userid, num, int32(pb.LOG_TYPE8))
	t.DeskGame.Dealer = ""
	t.DeskGame.DealerSeat = 0
	t.DeskFree.Carry = 0
	t.DeskFree.DealerNum = 0
	msg := handler.BeJHDealerMsg(0, num, t.DeskGame.Dealer,
		userid, user.GetNickname())
	t.broadcast(msg)
}

//离开房间清空上庄列表
func (t *Desk) leaveBeDealer(userid string) {
	if !t.isFree() {
		return
	}
	delete(t.DeskFree.Dealers, userid)
}

//开始游戏时选择玩家成为庄家
func (t *Desk) beComeDealer() {
	//if !t.isFreeDealer() {
	//	return
	//}
	userid, num := t.findBeDealer()
	if userid == "" {
		//glog.Errorf("beComeDealer failed %s, %d", userid, num)
		return
	}
	seat := t.getSeat(userid)
	if seat != 0 {
		arg := &pb.CJHSit{
			Type: pb.SitUp,
			Seat: seat,
		}
		rsp := t.freeSit(userid, arg)
		if rsp.Error == pb.OK {
			t.broadcast(rsp)
		} else {
			glog.Errorf("free sit up err %s, %d", userid, seat)
		}
	}
	//上庄成功扣除
	t.sendCoin(userid, (-1 * num), int32(pb.LOG_TYPE7))
	//成为庄家
	t.DeskGame.Dealer = userid
	t.DeskFree.Carry = num
	t.DeskFree.DealerNum = 0
	delete(t.DeskFree.Dealers, userid)
	//消息
	user := t.getPlayer(userid)
	if user == nil {
		return
	}
	msg := handler.BeJHDealerMsg(1, num, t.DeskGame.Dealer,
		userid, user.GetNickname())
	t.broadcast(msg)
}

//携带最大的优先做庄
func (t *Desk) findBeDealer() (userid string, carry int64) {
	if !t.isFree() {
		return
	}
	for k, v := range t.DeskFree.Dealers {
		if !t.isOnline(k) || v < int64(t.DeskData.Carry) {
			delete(t.DeskFree.Dealers, k)
			continue
		}
		if val, ok := t.roles[k]; ok {
			//自动下庄金额不足玩家
			if val.GetCoin() < int64(t.DeskData.Carry) {
				delete(t.DeskFree.Dealers, k)
				continue
			}
			//全部资金上庄
			if val.GetCoin() > carry {
				userid = k
				carry = val.GetCoin()
			}
		}
	}
	return
}

//是否百人上庄
func (t *Desk) isFree() bool {
	if !t.DeskData.Deal {
		return false
	}
	if t.DeskFree == nil {
		return false
	}
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE2): //百人
	default:
		return false
	}
	return true
}

//是否百人有庄
func (t *Desk) isFreeDealer() bool {
	if !t.isFree() {
		return false
	}
	if t.DeskGame.Dealer == "" {
		return false
	}
	return true
}

//结束时检测不足做庄
func (t *Desk) checkBeDealer() {
	if !t.isFree() {
		return
	}
	//做庄次数
	t.DeskFree.DealerNum++
	if v, ok := t.roles[t.DeskGame.Dealer]; ok {
		//离线或者不足
		if v.Offline || t.leftDealerTimes() == 0 ||
			t.DeskFree.Carry <= int64(t.DeskData.Down) ||
			t.DeskFree.Carry >= int64(t.DeskData.Top) {
			t.delBeDealer(t.DeskGame.Dealer, v.User)
		} else {
			return
		}
	}
	//选择成为庄家
	t.beComeDealer()
	//无人坐庄
	if t.DeskGame.Dealer == "" {
		//庄家每次都补庄
		if t.DeskFree.Carry < SysCarry {
			t.DeskFree.Carry = SysCarry
		}
		//重置次数
		if t.leftDealerTimes() == 0 {
			t.DeskFree.DealerNum = 0
		}
	}
}

//是否已经是庄家或者已经申请上庄
func (t *Desk) alreadyBeDealer(userid string) bool {
	if !t.isFree() {
		return false
	}
	//已经是庄
	if t.DeskGame.Dealer == userid {
		return true
	}
	//已经申请
	if _, ok := t.DeskFree.Dealers[userid]; ok {
		return true
	}
	return false
}

//.

//'玩家准备
func (t *Desk) readying(userid string, ready bool) (msg *pb.SJHReady) {
	msg = new(pb.SJHReady)
	if t.isFree() {
		msg.Error = pb.OperateError
		return
	}
	//投票中不能准备
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE1): //私人
		if t.DeskPriv.VoteSeat != 0 {
			msg.Error = pb.VotingCantLaunchVote
			return
		}
	}
	//已经开始不能准备
	if t.state != int32(pb.STATE_READY) {
		msg.Error = pb.GameStarted
		return
	}
	user := t.getPlayer(userid)
	if user == nil {
		msg.Error = pb.NotInRoom
		return
	}
	if user.GetCoin() < int64(t.DeskData.Ante) {
		msg.Error = pb.NotEnoughCoin
		return
	}
	//设置状态
	seat := t.getSeat(userid)
	if v, ok := t.seats[seat]; ok {
		v.Ready = ready
	}
	msg.Seat = seat
	msg.Ready = ready
	t.broadcast(msg) //广播消息
	if t.allReady() {
		t.gameStart() //开始牌局
	}
	return
}

//.

//'操作验证
func (t *Desk) coinCheck(seat uint32, val int32) pb.ErrCode {
	if t.isFree() {
		return pb.OperateError
	}
	if t.state == int32(pb.STATE_READY) {
		return pb.GameNotStart
	}
	if t.DeskHua.ActSeat != seat {
		glog.Errorf("coinCheck err %d, %d, %d",
			t.DeskHua.ActSeat, seat, val)
		return pb.NotYourTurn
	}
	if (t.DeskHua.ActState & val) != val {
		glog.Errorf("coinCheck err %d, %d, %d",
			t.DeskHua.ActState, seat, val)
		return pb.ErrorOperateValue
	}
	if v, ok := t.DeskHua.ActSeats[seat]; ok {
		if !v.Alive {
			return pb.AlreadyFold
		}
	} else {
		return pb.NotReady
	}
	return pb.OK
}

//.

//'看牌
func (t *Desk) coinSee(userid string) {
	msg := new(pb.SJHCoinSee)
	seat := t.getSeat(userid)
	msg.Error = t.coinCheck(seat, int32(pb.ACT_SEE))
	if msg.Error != pb.OK {
		t.send2userid(userid, msg)
		return
	}
	msg.Seat = seat
	msg.Userid = userid
	if v, ok := t.DeskHua.ActSeats[seat]; ok {
		v.See = true
		if val, ok := t.seats[seat]; ok {
			msg.Cards = val.Cards
		}
		//操作值中去掉看牌
		t.DeskHua.ActState ^= int32(pb.ACT_SEE)
		t.send2userid(userid, msg)
	}
	msg2 := new(pb.SJHCoinSee)
	msg.Seat = seat
	msg.Userid = userid
	t.broadcast2(userid, msg2)
}

//.

//'弃牌
func (t *Desk) coinFold(userid string) {
	msg := new(pb.SJHCoinFold)
	seat := t.getSeat(userid)
	msg.Error = t.coinCheck(seat, int32(pb.ACT_FOLD))
	if msg.Error != pb.OK {
		t.send2userid(userid, msg)
		return
	}
	msg.Seat = seat
	msg.Userid = userid
	if v, ok := t.DeskHua.ActSeats[seat]; ok {
		v.Alive = false
	}
	//操作值中去掉操作
	t.DeskHua.ActState ^= int32(pb.ACT_FOLD)
	t.broadcast(msg)
	//结束或者继续
	t.setNextActSeat()
}

//.

//'跟注
func (t *Desk) coinCall(userid string, num int64) {
	msg := new(pb.SJHCoinCall)
	seat := t.getSeat(userid)
	msg.Error = t.coinCheck(seat, int32(pb.ACT_CALL))
	if msg.Error != pb.OK {
		t.send2userid(userid, msg)
		return
	}
	user := t.getPlayer(userid)
	if user == nil {
		msg.Error = pb.NotInRoom
		t.send2userid(userid, msg)
		return
	}
	if user.GetCoin() < num {
		msg.Error = pb.NotEnoughCoin
		t.send2userid(userid, msg)
		return
	}
	msg.Seat = seat
	msg.Userid = userid
	//验证跟注数量
	var double int64 = 1
	if t.isSee(seat) {
		double = 2
	}
	glog.Debugf("userid %s, double %d, num %d", userid, double, num)
	glog.Debugf("userid %s, raise %d, call %d",
		userid, t.DeskHua.ActRaiseNum, t.DeskHua.ActCallNum)
	if (double * t.DeskHua.ActCallNum) != num {
		msg.Error = pb.CallError
		t.send2userid(userid, msg)
		return
	}
	//广播消息
	t.setBet(seat, userid, num)
	//结束或者继续
	t.setNextActSeat()
}

//.

//'加注
func (t *Desk) coinRaise(userid string, num int64) {
	msg := new(pb.SJHCoinRaise)
	seat := t.getSeat(userid)
	msg.Error = t.coinCheck(seat, int32(pb.ACT_RAISE))
	if msg.Error != pb.OK {
		t.send2userid(userid, msg)
		return
	}
	user := t.getPlayer(userid)
	if user == nil {
		msg.Error = pb.NotInRoom
		t.send2userid(userid, msg)
		return
	}
	if user.GetCoin() < num {
		msg.Error = pb.NotEnoughCoin
		t.send2userid(userid, msg)
		return
	}
	msg.Seat = seat
	msg.Userid = userid
	//验证加注数量
	var double int64 = 1
	if t.isSee(seat) {
		double = 2
	}
	glog.Debugf("userid %s, double %d, num %d", userid, double, num)
	glog.Debugf("userid %s, raise %d, call %d",
		userid, t.DeskHua.ActRaiseNum, t.DeskHua.ActCallNum)
	n := double * t.DeskHua.ActRaiseNum
	i := num / n
	if ((i % double) != 0) || ((num % n) != 0) {
		msg.Error = pb.RaiseError
		t.send2userid(userid, msg)
		return
	}
	//广播消息
	t.setBet(seat, userid, num)
	//加注底池
	if double == 1 {
		t.DeskHua.ActRaiseNum += num
		t.DeskHua.ActCallNum = num
	} else {
		t.DeskHua.ActRaiseNum += (num / 2)
		t.DeskHua.ActCallNum = (num / 2)
	}
	//结束或者继续
	t.setNextActSeat()
}

//.

//'比牌
func (t *Desk) coinBi(userid string, seat uint32) {
	msg := new(pb.SJHCoinBi)
	biseat := t.getSeat(userid)
	msg.Error = t.coinCheck(biseat, int32(pb.ACT_BI))
	if msg.Error != pb.OK {
		t.send2userid(userid, msg)
		return
	}
	msg.Seat = seat
	msg.Biseat = biseat
	cs1 := t.getHandCards(seat)
	cs2 := t.getHandCards(biseat)
	//TODO 扣除金币,比牌金额不足怎么处理
	if algo.HuaCompare(cs1, cs2) {
		//cs1 赢
		msg.Winseat = seat
		msg.Loseseat = biseat
	} else {
		//cs2 赢
		msg.Winseat = biseat
		msg.Loseseat = seat
	}
	if v, ok := t.DeskHua.ActSeats[msg.Loseseat]; ok {
		v.Alive = false
	}
	//操作值中去掉操作
	t.DeskHua.ActState ^= int32(pb.ACT_BI)
	t.broadcast(msg)
	//结束或者继续
	curr := t.getNextActSeat()
	//结束,winner = last
	if curr == 0 {
		t.gameOver()
		return
	}
	//不切换位置
	t.timer = 0
	//广播操作状态消息
	t.pushActState()
}

//.

//'换房间
func (t *Desk) changeDesk(ctx actor.Context) {
	userid := t.getRouter(ctx)
	errcode := t.changeDeskCheck(userid)
	if errcode != pb.OK {
		rsp := new(pb.SJHCoinChangeRoom)
		rsp.Error = errcode
		ctx.Respond(rsp)
		return
	}
	t.nnLeave(userid, ctx)
	//匹配房间消息
	msg := new(pb.ChangeDesk)
	msg.Roomid = t.DeskData.Rid
	msg.Rtype = t.DeskData.Rtype
	msg.Gtype = t.DeskData.Gtype
	msg.Ltype = t.DeskData.Ltype
	msg.Dtype = t.DeskData.Dtype
	msg.Userid = userid
	msg.Sender = ctx.Sender() //玩家进程
	nodePid.Request(msg, ctx.Self())
}

func (t *Desk) changeDeskCheck(userid string) pb.ErrCode {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		switch t.state {
		case int32(pb.STATE_READY):
		default:
			seat := t.getSeat(userid)
			if v, ok := t.seats[seat]; ok {
				if v.Ready {
					return pb.GameStartedCannotLeave
				}
			}
		}
	default:
		return pb.OperateError
	}
	return pb.OK
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
