/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-04-30 18:01:29
 * Filename      : desk_ebg.go
 * Description   : 私人场玩牌逻辑
 * *******************************************************/
package main

import (
	"gohappy/game/algo"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
	"gohappy/game/config"
)

//'进入房间响应消息
func (t *Desk) privEnterMsg(userid string) *pb.SEBEnterRoom {
	msg := new(pb.SEBEnterRoom)
	//房间数据
	msg.Roominfo = handler.PackEBCoinRoom(t.DeskData)
	t.roomDataMsg(msg.Roominfo)
	//坐下玩家信息
	msg.Userinfo = t.coinSeatBetsMsg(userid)
	//位置下注信息
	msg.Betsinfo = t.coinBetsMsg()
	//投票信息
	msg.Voteinfo = t.voteInfoMsg()
	return msg
}

//投票信息
func (t *Desk) voteInfoMsg() (msg *pb.EBRoomVote) {
	msg = new(pb.EBRoomVote)
	if t.DeskPriv != nil {
		msg.Seat = t.DeskPriv.VoteSeat
	}
	if msg.Seat == 0 {
		return
	}
	for k, v := range t.seats {
		if v.Vote == 1 {
			msg.Agree = append(msg.Agree, k)
		} else if v.Vote == 2 {
			msg.Disagree = append(msg.Disagree, k)
		}
	}
	return
}

//.

//'投票

func (t *Desk) checkVote() pb.ErrCode {
	if t.isFree() {
		return pb.OperateError
	}
	//if t.state > int32(pb.STATE_READY) {
	//	return pb.RunningNotVote
	//}
	if t.DeskPriv == nil {
		return pb.OperateError
	}
	return pb.OK
}

//发起投票
func (t *Desk) launchVote(userid string, vote uint32) (msg *pb.SEBLaunchVote) {
	msg = new(pb.SEBLaunchVote)
	errcode := t.checkVote()
	if errcode != pb.OK {
		msg.Error = errcode
		return
	}
	if t.DeskPriv.VoteSeat != 0 {
		msg.Error = pb.VotingCantLaunchVote
		return
	}
	seat := t.getSeat(userid)
	if v, ok := t.seats[seat]; ok {
		v.Vote = vote //投票
	}
	//发起投票者
	t.DeskPriv.VoteSeat = seat
	//超时设置(1分钟)
	glog.Debugf("VoteTime: %d, %d, %s", seat, vote, userid)
	t.DeskPriv.VoteTime = utils.Timestamp() + 60
	msg.Seat = seat
	t.broadcast(msg)
	t.pushVote(seat, vote)
	t.dismiss(false)
	return
}

func (t *Desk) voteTimeout() {
	errcode := t.checkVote()
	if errcode != pb.OK {
		return
	}
	if t.DeskPriv.VoteSeat == 0 {
		return
	}
	var now = utils.Timestamp()
	if now >= t.DeskPriv.VoteTime {
		t.dismiss(true)
	}
}

//投票
func (t *Desk) privVote(userid string, vote uint32) (msg *pb.SEBVote) {
	msg = new(pb.SEBVote)
	errcode := t.checkVote()
	if errcode != pb.OK {
		msg.Error = errcode
		return
	}
	if t.DeskPriv.VoteSeat == 0 {
		msg.Error = pb.NotVoteTime
		return
	}
	seat := t.getSeat(userid)
	if v, ok := t.seats[seat]; ok {
		v.Vote = vote //投票
	}
	t.pushVote(seat, vote)
	t.dismiss(false)
	return
}

//广播投票消息
func (t *Desk) pushVote(seat, vote uint32) {
	msg := &pb.SEBVote{
		Seat: seat,
		Vote: vote,
	}
	t.broadcast(msg)
}

//广播投票消息
func (t *Desk) pushVoteResult(vote uint32) {
	msg := &pb.SEBVoteResult{
		Vote: vote,
	}
	t.broadcast(msg)
}

//投票解散,agree >= unagree
func (t *Desk) dismiss(force bool) {
	var agree = 0
	var unagree = 0
	var voted = 0
	for _, v := range t.seats {
		if v.Vote == 1 {
			agree++
		} else {
			unagree++
		}
		if v.Vote != 0 {
			voted++
		}
	}
	glog.Debugf("dismiss force %t, vote %d, len %d",
		force, t.DeskPriv.VoteSeat, len(t.seats))
	glog.Debugf("dismiss voted %d,agree %d, unagree %d",
		voted, agree, unagree)
	//一半以上通过即可
	if agree >= unagree {
		//0解散,1不解散
		t.pushVoteResult(0)
		//停止服务
		msg1 := new(pb.ServeStop)
		t.selfPid.Tell(msg1)
	} else if force || voted == len(t.seats) {
		//结束投票
		t.pushVoteResult(1)
		//重置
		for _, v := range t.seats {
			v.Vote = 0
		}
		//发起投票者
		t.DeskPriv.VoteSeat = 0
		t.DeskPriv.VoteTime = 0
	}
}

//.

//' 超时处理
func (t *Desk) coinTimeout() {
	switch t.state {
	case int32(pb.STATE_READY):
		if t.timer == ReadyTime {
			//准备超时,不等待全部准备
			t.readyTimeout()
			t.timer = 0
		} else {
			t.timer++
		}
		return
	}
	if t.timer == BetTime {
		switch t.state {
		case int32(pb.STATE_DEALER):
			//抢庄超时,打庄
			t.dealerHandler()
		case int32(pb.STATE_BET):
			//下注超时
			t.betTimeout()
		case int32(pb.STATE_NIU):
			//提交组合超时,结束
			t.niuTimeout()
		}
	} else {
		t.timer++
	}
}

//.

//' 超时处理
func (t *Desk) privTimeout() {
	t.voteTimeout()
	switch t.state {
	case int32(pb.STATE_READY):
		//过期关闭
		if t.checkExpire() {
			//关闭房间
			t.gameStop()
			return
		}
		if t.timer == ReadyTime {
			//准备超时,不等待全部准备
			t.readyTimeout()
			t.timer = 0
		} else {
			t.timer++
		}
		return
	}
	if t.timer == BetTime {
		switch t.state {
		case int32(pb.STATE_DEALER):
			//抢庄超时,打庄
			t.dealerHandler()
		case int32(pb.STATE_BET):
			//下注超时
			t.betTimeout()
		case int32(pb.STATE_NIU):
			//提交组合超时,结束
			t.niuTimeout()
		}
	} else {
		t.timer++
	}
}

//.

//'结束游戏
func (t *Desk) gameOver() {
	//当前局积分
	score := make(map[uint32]int64)
	switch t.DeskData.Dtype {
	//case int32(pb.DESK_TYPE1): //通比牛牛
	//	var seat uint32
	//	seat, score = t.pailiOver3(score)
	//	//赢家seat收钱
	//	score = t.jiesuan3(seat, score)
	case int32(pb.DESK_TYPE0), //看牌抢庄
		int32(pb.DESK_TYPE1), //通比牛牛
		int32(pb.DESK_TYPE3), //抢庄
		int32(pb.DESK_TYPE2): //抢庄看牌
		score = t.pailiOver1(score)
		//庄家先收钱，再赔付
		score = t.jiesuan1(score)
	}
	//打印信息
	t.printOver()
	glog.Debugf("score %#v", score)
	//个人记录
	t.setRecord(score)
	//牌型奖励
	t.powerAward()
	//结算消息
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		//结算消息
		msg := t.resCoinOver(score)
		t.broadcast(msg)
		//TODO 记录
		//结束连庄处理
		//t.dealerOver()
		//重置状态
		t.gameOverInit()
		//踢出不足坐下玩家或超额玩家
		t.limitOver()
		//踢除离线玩家
		t.kickOffline()
	case int32(pb.ROOM_TYPE1): //私人
		//牌局数累加一次
		if !t.DeskData.Pub {
			t.DeskGame.Round++
		}
		//结算消息
		msg := t.resOver(score)
		glog.Debugf("SEBGameover %#v", msg)
		for k, v := range t.roles {
			glog.Debugf("game over userid %s -> %d, %v", k, v.Seat, v.Offline)
		}
		t.broadcast(msg)
		//记录
		if !t.DeskData.Pub {
			//t.saveRecord(score)
		}
		t.saveRecord(score)
		//结束连庄处理
		//t.dealerOver()
		//重置状态
		t.gameOverInit()
		//踢出不足坐下玩家或超额玩家
		t.limitOver()
		//关闭房间
		//t.gameStop()
	case int32(pb.ROOM_TYPE2): //百人
	}
}

//.

//'踢出不足坐下玩家或超额玩家
func (t *Desk) limitOver() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
	case int32(pb.ROOM_TYPE1): //私人
		if !t.DeskData.Pub {
			//return
		}
	case int32(pb.ROOM_TYPE2): //百人
		return
	}
	for k, v := range t.roles {
		coin := v.User.GetCoin()
		//if t.DeskData.Maximum == 0 {
		//	if coin >= t.DeskData.Minimum {
		//		continue
		//	}
		//} else {
		//	if coin >= t.DeskData.Minimum &&
		//		coin < t.DeskData.Maximum {
		//		continue
		//	}
		//}
		if coin >= t.DeskData.Minimum { //离场限制
			continue
		}
		errcode := t.leave(k)
		if errcode != pb.OK {
			continue
		}
		//离开状态消息
		t.userLeaveDesk(k)
	}
}

// 踢除离线玩家
func (t *Desk) kickOffline() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE2): //百人
		for k, v := range t.roles {
			if !v.Offline {
				continue
			}
			errcode := t.leave(k)
			if errcode != pb.OK {
				continue
			}
			//离开状态消息
			t.userLeaveDesk(k)
		}
	}
}

// pub房间人数为0时解散
func (t *Desk) checkPubOver() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE1): //私人
		if !t.DeskData.Pub {
			//return
		}
	default:
		//return
	}
	if len(t.roles) != 0 {
		return
	}
	g := config.GetGame(t.DeskData.Unique)
	if g.Id == t.DeskData.Unique {
		return //配置房间不关闭
	}
	//停止服务
	msg1 := new(pb.ServeStop)
	t.selfPid.Tell(msg1)
}

//.

//'结束连庄处理
/*
func (t *Desk) dealerOver() {
	//牛牛抢庄 房间不重置
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE2): //抢庄看牌
	default:
		t.DeskGame.Dealer = ""
		t.DeskGame.DealerSeat = 0
		return
	}
	a := make([]uint32, 0)
	for k, v := range t.seats {
		switch v.Power {
		case algo.NiuNiu:
			a = append(a, k)
		}
	}
	if len(a) > 0 {
		t.DeskGame.DealerSeat = a[0]
		a = a[1:]
		for {
			if len(a) <= 0 {
				break
			}
			if !t.pailiCompare(t.DeskGame.DealerSeat, a[0]) {
				t.DeskGame.DealerSeat = a[0]
			}
			a = a[1:]
		}
	}
	if val, ok := t.seats[t.DeskGame.DealerSeat]; ok {
		t.DeskGame.Dealer = val.Userid
	}
}
*/
//.

//'结束游戏
//是否过期
func (t *Desk) checkExpire() bool {
	var now = utils.Timestamp()
	if now > int64(t.DeskData.Expire) {
		glog.Debugf("game stop expire -> %d, %d",
			t.DeskData.Expire, now)
		return true
	}
	return false
}

//是否结束游戏
func (t *Desk) checkOver() bool {
	if t.DeskData.Round == t.DeskGame.Round {
		glog.Debugf("game stop round -> %d, %d",
			t.DeskGame.Round, t.DeskData.Round)
		return true
	}
	return t.checkExpire()
}

//结束牌局
func (t *Desk) gameStop() {
	if !t.checkOver() {
		return
	}
	if t.DeskData.Pub { //大厅房间不解散
		return
	}
	//返回未开局钻石
	t.backCost()
	//停止服务
	msg1 := new(pb.ServeStop)
	t.selfPid.Tell(msg1)
}

//返还钻石
func (t *Desk) backCost() {
	//已经打过的房间不返还
	if t.DeskGame.Round != 0 {
		return
	}
	//已经开始游戏不返还
	if t.state != int32(pb.STATE_READY) {
		return
	}
	//A
	if t.DeskData.Payment != 1 {
		t.sendDiamond(t.DeskData.Cid,
			int64(t.DeskData.Cost), int32(pb.LOG_TYPE3))
		return
	}
	//AA
	for k := range t.roles {
		t.sendDiamond(k, int64(t.DeskData.Cost), int32(pb.LOG_TYPE3))
	}
}

//.

//'个人记录
func (t *Desk) setRecord(score map[uint32]int64) {
	for k, v := range score {
		user := t.getUserBySeat(k)
		if user == nil {
			continue
		}
		pid := t.getPid(user.GetUserid())
		if pid == nil {
			continue
		}
		msg := new(pb.SetRecord)
		if v > 0 {
			msg.Rtype = 1
		} else if v < 0 {
			msg.Rtype = -1
		} else {
			msg.Rtype = 0
		}
		//更新游戏内数据
		user.SetRecord(msg.Rtype)
		//更新节点数据
		pid.Tell(msg)
		switch t.DeskData.Rtype {
		case int32(pb.ROOM_TYPE0): //自由
			if v > 0 {
				msg2 := handler.TaskUpdateMsg(1, pb.TASK_TYPE7, user.GetUserid())
				pid.Tell(msg2)
			}
		case int32(pb.ROOM_TYPE1): //私人
			if t.DeskData.Pub {
				msg2 := handler.TaskUpdateMsg(1, pb.TASK_TYPE21, user.GetUserid())
				pid.Tell(msg2)
			}
		}
	}
}

//.

//' 结算

//牌力比较(看牌抢庄 | 牛牛抢庄)
func (t *Desk) pailiOver1(score map[uint32]int64) map[uint32]int64 {
	//庄家牌力
	dealerSeat := t.DeskGame.DealerSeat
	var a uint32 = t.getPower(dealerSeat)
	ante := t.getAnteOver(dealerSeat)
	for k, v := range t.seats {
		if k == dealerSeat || !v.Ready {
			continue
		}
		switch {
		case v.Power > a:
			val := int64(v.Bet * int64(algo.EbgMultiple(t.DeskData.Mode, v.Power)) * ante)
			score = t.over3(k, dealerSeat, val, score)
		case v.Power < a:
			val := int64(v.Bet * int64(algo.EbgMultiple(t.DeskData.Mode, a)) * ante)
			score = t.over3(dealerSeat, k, val, score)
		case v.Power == a:
			//麻将牌：庄家跟闲家为0点时闲家赢，否则庄家赢
			if a != algo.EBG0 {
			//if t.pailiCompare(dealerSeat, k) {
				val := int64(v.Bet * int64(algo.EbgMultiple(t.DeskData.Mode, a)) * ante)
				score = t.over3(dealerSeat, k, val, score)
			} else {
				val := int64(v.Bet * int64(algo.EbgMultiple(t.DeskData.Mode, v.Power)) * ante)
				score = t.over3(k, dealerSeat, val, score)
			}
		}
	}
	return score
}

//牌力比较(通比牛牛)
func (t *Desk) pailiOver3(score map[uint32]int64) (uint32, map[uint32]int64) {
	//所有提交牛位置
	a := make([]uint32, 0)
	for k, v := range t.seats {
		if v.Niu && v.Ready {
			a = append(a, k)
		}
	}
	if len(a) == 0 {
		t.printOver()
		for k, v := range t.seats {
			glog.Errorf("pailiOver3 seat %d -> %#v", k, v)
		}
	}
	//赢家位置和牌力
	var seat = a[0]
	var val uint32 = t.getPower(a[0])
	a = a[1:]
	for {
		if len(a) <= 0 {
			break
		}
		v := t.getPower(a[0])
		switch {
		case v > val:
			seat = a[0]
			val = v
		case v == val:
			//麻将牌：庄家跟闲家为0点时闲家赢，否则庄家赢
			if val == algo.EBG0 {
			//if !t.pailiCompare(seat, a[0]) {
				seat = a[0]
				val = v
			}
		}
		a = a[1:]
	}
	//赢家
	var a1 uint32 = t.getPower(seat)
	bet1 := t.getBets(seat)
	ante := t.getAnteOver(seat)
	for k, v := range t.seats {
		if k == seat || !v.Ready {
			continue
		}
		val := int64(v.Bet * bet1 * int64(algo.EbgMultiple(t.DeskData.Mode, a1)) * ante)
		score = t.over3(seat, k, val, score)
	}
	return seat, score
}

//积分结算
func (t *Desk) over3(win, lose uint32, v int64,
	score map[uint32]int64) map[uint32]int64 {
	user := t.getUserBySeat(lose)
	if user != nil {
		if user.GetCoin() < v {
			v = user.GetCoin() //不足赔付
		}
	}
	score[win] += v
	score[lose] -= v
	return score
}

//牌力比较
func (t *Desk) pailiCompare(s1, s2 uint32) bool {
	cs1 := t.getHandCards(s1)
	cs2 := t.getHandCards(s2)
	return algo.Compare(cs1, cs2)
}

//结算倍数
func (t *Desk) getAnteOver(seat uint32) (ante int64) {
	if v, ok := t.seats[seat]; ok {
		//抢庄倍数
		if v.DealerN > 0 {
			ante = int64(t.DeskData.Ante * v.DealerN)
			return
		}
	}
	//底分
	ante = int64(t.DeskData.Ante)
	return
}

//结算
func (t *Desk) jiesuan(seat uint32, score map[uint32]int64) (map[uint32]int64, int64, int64) {
	var num1, num2 int64
	for k, v := range score {
		if k == seat {
			continue
		}
		if v >= 0 {
			num2 += v //闲家赢,正数
			continue
		}
		user := t.getUserBySeat(k)
		if user == nil {
			continue
		}
		coin := user.GetCoin()
		val := (-1 * int64(coin))
		if val > v { //不足
			v = val
		}
		score[k] = v
		num1 += v //闲家输,负数
	}
	return score, num1, num2
}

//结算
func (t *Desk) jiesuan1(score map[uint32]int64) map[uint32]int64 {
	score, num1, num2 := t.jiesuan(t.DeskGame.DealerSeat, score)
	user := t.getPlayer(t.DeskGame.Dealer)
	if user == nil {
		glog.Errorf("jiesuan1 err %#v, %d, %d", score, num1, num2)
		glog.Errorf("jiesuan1 err data -> %#v", t.DeskData)
		glog.Errorf("jiesuan1 err dealer -> %s", t.DeskGame.Dealer)
		glog.Errorf("jiesuan1 err dealer seat -> %d", t.DeskGame.DealerSeat)
		return score
	}
	coin := user.GetCoin() //庄家
	if int64(coin) >= (num2 + num1) {
		//足够赔付
	} else {
		//不足赔付, 庄家先收钱，再赔付
		num := int64(coin) + (-1 * num1) //总金额
		winer := make([]uint32, 0)
		for k, v := range score {
			if v > 0 {
				winer = append(winer, k)
			}
		}
		val := num
		//按比例分到每个人
		for {
			if val <= 0 {
				break
			}
			if len(winer) == 0 {
				break
			}
			seat := winer[0]
			n := score[seat]
			if len(winer) == 1 {
				score[seat] = val
			} else {
				score[seat] = (n / num2) * num
				val -= score[seat]
			}
			winer = winer[1:]
		}
		score[t.DeskGame.DealerSeat] = -1 * int64(coin)
	}
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE0): //看牌抢庄
		t.jiesuan2(int32(pb.LOG_TYPE25), score)
	case int32(pb.DESK_TYPE1): //通比牛牛
	case int32(pb.DESK_TYPE2): //抢庄看牌
		t.jiesuan2(int32(pb.LOG_TYPE26), score)
	}
	return score
}

//结算
func (t *Desk) jiesuan3(seat uint32, score map[uint32]int64) map[uint32]int64 {
	score, num1, _ := t.jiesuan(seat, score)
	score[seat] = (-1 * num1)
	t.jiesuan2(int32(pb.LOG_TYPE24), score)
	return score
}

//结算
func (t *Desk) jiesuan2(ltype int32, score map[uint32]int64) {
	for k, v := range score {
		userid := t.getUserid(k)
		//抽水
		//v = t.drawcoin(userid, v)
		switch t.DeskData.Rtype {
		case int32(pb.ROOM_TYPE0): //自由
			t.sendCoin(userid, v, ltype)
		case int32(pb.ROOM_TYPE1): //私人
			t.sendCoin(userid, v, ltype)
			t.DeskPriv.PrivScore[userid] += v
			t.DeskPriv.Joins[userid]++
		}
	}
}

//抽水
func (t *Desk) drawcoin(userid string, val int64) int64 {
	if val <= 0 {
		return val
	}
	var num int64 = handler.DrawCoin(t.DeskData.Rtype, t.DeskData.Mode, val)
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1): //私人
		//反佣和收益消息,抽成日志记录 val - num
		msg2 := handler.AgentProfitNumMsg(userid, t.DeskData.Gtype, num)
		t.send3userid(userid, msg2)
	case int32(pb.ROOM_TYPE2): //百人
		//反佣和收益消息,抽成日志记录 val - num
		msg2 := handler.AgentProfitNumMsg(userid, t.DeskData.Gtype, num)
		t.send3userid(userid, msg2)
	}
	return val - num
}

//开始前扣除抽水
func (t *Desk) drawfee() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1): //私人
	case int32(pb.ROOM_TYPE2): //百人
		return
	}
	if t.state != int32(pb.STATE_READY) {
		return
	}
	//计算反佣和收益
	var num int64 = handler.DrawFee(t.DeskData.Mode, t.DeskData.Ante)
	for k, v := range t.seats {
		if !v.Ready {
			continue
		}
		if num <= 0 {
			continue
		}
		t.sendCoin(v.Userid, (-1 * num), int32(pb.LOG_TYPE48))
		//抽水消息广播
		msg := &pb.SEBPushDrawCoin{
			Rtype:  uint32(pb.LOG_TYPE48),
			Userid: v.Userid,
			Seat:   k,
			Coin:   (-1 * num),
		}
		t.broadcast(msg)
		//反佣和收益消息
		msg2 := handler.AgentProfitNumMsg(v.Userid, t.DeskData.Gtype, num)
		t.send3userid(v.Userid, msg2)
	}
}

//牌型奖励
func (t *Desk) powerAward() {
	for _, v := range t.seats {
		if !v.Ready {
			continue
		}
		var num = t.DeskData.Ante
		switch v.Power {
		case algo.EBG10, algo.EBGDui,	algo.BAIBAN:
			num += algo.EbgMultiple(t.DeskData.Mode, v.Power)
		default:
			continue
		}
		t.sendDiamond(v.Userid, int64(num), int32(pb.LOG_TYPE51))
	}
}

//.

//'日志记录
func (t *Desk) saveRecord(score map[uint32]int64) {
	msg := new(pb.RoundRecord)
	msg.Roomid = t.DeskData.Rid
	msg.Round = t.DeskData.Round
	msg.Dealer = t.DeskGame.Dealer
	for k, v := range score {
		if val, ok := t.seats[k]; ok {
			msg2 := &pb.RoundRoleRecord{
				Userid: val.Userid,
				Cards:  val.Cards,
				Value:  val.Power,
				Bets:   val.Bet,
				Score:  v,
			}
			if val2, ok2 := t.roles[val.Userid]; ok2 {
				msg2.Rest = val2.User.GetCoin()
			}
			msg.Roles = append(msg.Roles, msg2)
		}
	}
	t.loggerPid.Tell(msg)
	for k := range score {
		user := t.getUserBySeat(k)
		if user == nil {
			continue
		}
		msg1 := new(pb.RoleRecord)
		msg1.Roomid = t.DeskData.Rid
		msg1.Gtype = t.DeskData.Gtype
		msg1.Userid = user.GetUserid()
		msg1.Nickname = user.GetNickname()
		msg1.Photo = user.GetPhoto()
		msg1.Rest = user.GetCoin()
		msg1.Score = t.DeskPriv.PrivScore[user.GetUserid()]
		msg1.Joins = t.DeskPriv.Joins[user.GetUserid()]
		t.loggerPid.Tell(msg1)
	}
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
