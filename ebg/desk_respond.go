/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-04-30 18:02:06
 * Filename      : desk_respond.go
 * Description   : 响应请求
 * *******************************************************/
package main

import (
	"strings"

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
		glog.Debugf("game over userid %s -> %d, %v", k, v.Seat, v.Offline)
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
	} else {
		for k, v := range t.seats {
			glog.Debugf("game over seat %d -> %#v", k, v)
		}
	}
	if t.DeskPriv != nil {
		glog.Debugf("game over priv vote %d", t.DeskPriv.VoteSeat)
		glog.Debugf("game over priv score %#v", t.DeskPriv.PrivScore)
		glog.Debugf("game over priv joins %#v", t.DeskPriv.Joins)
	}
}

//.

//' 聊天
func (t *Desk) chatText(arg *pb.CChatText, ctx actor.Context) {
	userid := t.getRouter(ctx)
	seat := t.getSeat(userid)
	glog.Debugf("CChatText %s, %d", userid, seat)
	//房间消息广播,聊天
	msg := handler.ChatTextMsg(seat, userid, arg.Content)
	msg.Error = t.chatEmoji(userid, arg.GetContent())
	if msg.Error != pb.OK {
		t.send2userid(userid, msg)
		return
	}
	t.broadcast(msg)
}

func (t *Desk) chatVoice(arg *pb.CChatVoice, ctx actor.Context) {
	userid := t.getRouter(ctx)
	seat := t.getSeat(userid)
	glog.Debugf("CChatVoice %s, %d", userid, seat)
	//房间消息广播,聊天
	t.broadcast(handler.ChatVoiceMsg(seat, userid, arg.Content))
}

// 表情包,TODO 严格验证或新加协议
func (t *Desk) chatEmoji(userid, context string) pb.ErrCode {
	if !strings.HasPrefix(context, "_p") {
		return pb.OK
	}
	s := strings.Split(context, "/")
	if len(s) != 3 {
		return pb.OK
	}
	var num int64 = 5 //TODO 消耗数量配置
	if s[1] == "-1" {
		n := len(t.seats) - 1
		if n > 0 {
			num *= int64(n)
		}
	}
	if v, ok := t.roles[userid]; ok && v != nil {
		if v.User.GetDiamond() < num {
			return pb.NotEnoughDiamond
		}
	}
	//TODO 货币变更消息广播
	t.sendDiamond(userid, -1*num, int32(pb.LOG_TYPE50))
	return pb.OK
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
	if user.GetCoin() < t.DeskData.Maximum {
		return pb.NotEnoughCoin
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
	glog.Debugf("nnLeave userid %s, err %v", userid, errcode)
	if errcode == pb.OK {
		//清除数据
		t.userLeaveDesk(userid)
		return
	}
	msg := new(pb.SEBLeave)
	msg.Seat = t.getSeat(userid)
	msg.Userid = userid
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
		msg1 := new(pb.SEBLeave)
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
		////离开房间,TODO 玩家进程中操作,一致性
		t.roomPid.Request(msg, t.selfPid)
	}
}

func (t *Desk) offlineMsg(userid string) {
	if v, ok := t.roles[userid]; ok && v != nil {
		msg := new(pb.SEBPushOffline)
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
	msg := handler.EbgBeDealerMsg(st, num, t.DeskGame.Dealer,
		userid, user.GetNickname(), user.GetPhoto())
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
	msg := handler.EbgBeDealerMsg(0, num, t.DeskGame.Dealer,
		userid, user.GetNickname(), user.GetPhoto())
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
		arg := &pb.CEBSit{
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
	msg := handler.EbgBeDealerMsg(1, num, t.DeskGame.Dealer,
		userid, user.GetNickname(), user.GetPhoto())
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
			//if val.GetCoin() > carry {
			//	userid = k
			//	carry = val.GetCoin()
			//}
			//选择金额上庄
			if v > carry && val.GetCoin() >= v {
				userid = k
				carry = v
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
func (t *Desk) readying(userid string, ready bool) (msg *pb.SEBReady) {
	msg = new(pb.SEBReady)
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

//' 抢庄 (看牌抢庄)
func (t *Desk) choiceDealer(userid string, isDealer bool,
	num uint32) (msg *pb.SEBDealer) {
	msg = new(pb.SEBDealer)
	if t.isFree() {
		msg.Error = pb.OperateError
		return
	}
	//不是抢庄状态
	if t.state != int32(pb.STATE_DEALER) {
		glog.Errorf("choiceDealer err:%s, %d", userid, num)
		msg.Error = pb.OperateError
		return
	}
	//看牌抢庄
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE3): //抢庄
	default:
		glog.Errorf("choiceDealer err:%s, %d", userid, num)
		msg.Error = pb.OperateError
		return
	}
	//状态
	seat := t.getSeat(userid)
	if v, ok := t.seats[seat]; ok {
		if !v.Ready {
			glog.Errorf("choiceDealer err:%s, %d", userid, num)
			msg.Error = pb.OperateError
			return
		}
		if isDealer {
			v.BeDealer = 1
			v.DealerN = num
		} else {
			v.BeDealer = 2
		}
	}
	msg.Dealer = isDealer
	msg.Seat = seat
	msg.Num = num
	t.broadcast(msg)
	//全部抢庄完成
	t.choiceDealerOver()
	return
}

//.

//' 下注
func (t *Desk) choiceBet(userid string, seatBet,
	num uint32) (msg *pb.SEBBet) {
	msg = new(pb.SEBBet)
	if t.isFree() {
		msg.Error = pb.OperateError
		return
	}
	//不是下注状态
	if t.state != int32(pb.STATE_BET) {
		glog.Errorf("choiceBet err:%s, %d, %d", userid, num, t.state)
		msg.Error = pb.OperateError
		return
	}
	if num == 0 { //下注不能下0
		glog.Errorf("choiceBet err:%s, %d", userid, num)
		msg.Error = pb.OperateError
		return
	}
	seat := t.getSeat(userid)
	//庄家不下注
	if seat == t.DeskGame.DealerSeat {
		msg.Error = pb.BetDealerFailed
		return
	}
	user := t.getPlayer(userid)
	if user == nil {
		msg.Error = pb.NotInRoom
		return
	}
	//参与者下注
	val := t.seats[seat]
	if val == nil {
		msg.Error = pb.OperateError
		return
	}
	if !val.Ready {
		msg.Error = pb.OperateError
		return
	}
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1): //私人
		coin := user.GetCoin() //剩余金额
		//本轮下注不能超过1/4
		if (int64(num) + val.Bet) > ((coin + val.Bet) / 4) {
			msg.Error = pb.BetTopLimit
			return
		}
		//扣除金币
		//t.sendCoin(userid, (-1 * int64(num)), int32(pb.LOG_TYPE5))
	}
	val.Bet = int64(num)
	msg.Seat = seat
	msg.Seatbet = seatBet
	msg.Value = num
	t.broadcast(msg)
	t.choiceBetOver()
	return
}

//.

//'提交组合
func (t *Desk) choiceNiu(userid string, num uint32,
	cs []uint32, ctx actor.Context) (msg *pb.SEBiu) {
	msg = new(pb.SEBiu)
	if t.isFree() {
		msg.Error = pb.OperateError
		return
	}
	//不在状态
	if t.state != int32(pb.STATE_NIU) {
		glog.Errorf("choiceNiu err:%#v, %d", cs, num)
		msg.Error = pb.GameNotStart
		return
	}
	seat := t.getSeat(userid)
	if v, ok := t.seats[seat]; ok {
		num = algo.Algo(t.DeskData.Mode, v.Cards)
		v.Power = num
		v.Niu = true
		msg.Seat = seat
		msg.Value = num
		msg.Cards = v.Cards
		t.broadcast(msg)
		t.choiceNiuOver(ctx)
	} else {
		msg.Error = pb.NotInRoom
		return
	}
	return
}

//结束提交
func (t *Desk) choiceNiuOver(ctx actor.Context) {
	for _, v := range t.seats {
		if v.Ready && !v.Niu {
			return
		}
	}
	//已经全部提交,结束游戏
	t.gameOver()
}

//.

//'换房间
func (t *Desk) changeDesk(ctx actor.Context) {
	userid := t.getRouter(ctx)
	errcode := t.changeDeskCheck(userid)
	if errcode != pb.OK {
		rsp := new(pb.SEBCoinChangeRoom)
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
