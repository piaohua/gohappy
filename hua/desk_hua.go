/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-04-30 18:01:29
 * Filename      : desk_hua.go
 * Description   : 私人场玩牌逻辑
 * *******************************************************/
package main

import (
	"gohappy/game/algo"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

//'进入房间响应消息
func (t *Desk) privEnterMsg(userid string) *pb.SJHEnterRoom {
	msg := new(pb.SJHEnterRoom)
	//房间数据
	msg.Roominfo = handler.PackJHCoinRoom(t.DeskData)
	msg.Roominfo.State = t.state
	//TODO 添加操作信息
	//坐下玩家信息
	msg.Userinfo = t.coinSeatBetsMsg(userid)
	//位置下注信息
	msg.Betsinfo = t.coinBetsMsg()
	//投票信息
	msg.Voteinfo = t.voteInfoMsg()
	return msg
}

//投票信息
func (t *Desk) voteInfoMsg() (msg *pb.JHRoomVote) {
	msg = new(pb.JHRoomVote)
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
func (t *Desk) launchVote(userid string, vote uint32) (msg *pb.SJHLaunchVote) {
	msg = new(pb.SJHLaunchVote)
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
func (t *Desk) privVote(userid string, vote uint32) (msg *pb.SJHVote) {
	msg = new(pb.SJHVote)
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
	msg := &pb.SJHVote{
		Seat: seat,
		Vote: vote,
	}
	t.broadcast(msg)
}

//广播投票消息
func (t *Desk) pushVoteResult(vote uint32) {
	msg := &pb.SJHVoteResult{
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
		default:
			t.timer = 0
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
		default:
			t.timer = 0
		}
	} else {
		t.timer++
	}
}

//.

//'结束游戏
func (t *Desk) gameOver() {
	winner := t.DeskHua.ActSeat
	for k, v := range t.DeskHua.ActSeats {
		if k == winner {
			continue
		}
		if !v.Alive {
			continue
		}
		cs1 := t.getHandCards(k)
		cs2 := t.getHandCards(winner)
		if algo.HuaCompare(cs1, cs2) {
			//cs1 赢
			if val, ok := t.DeskHua.ActSeats[winner]; ok {
				val.Alive = false
			}
			winner = k
		} else {
			v.Alive = false
		}
	}
	//当前局积分
	score := make(map[uint32]int64)
	score[winner] = t.DeskGame.BetNum
	for k, v := range t.DeskHua.ActSeats {
		if !v.Alive {
			score[k] = 0 - v.ActNum
		}
	}
	t.jiesuan2(int32(pb.LOG_TYPE45), score)
	//打印信息
	t.printOver()
	glog.Debugf("score %#v", score)
	//个人记录
	t.setRecord(score)
	//结算消息
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		//结算消息
		msg := t.resCoinOver(score)
		t.broadcast(msg)
		//TODO 记录
		//结束连庄处理
		t.dealerOver()
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
		t.broadcast(msg)
		//记录
		if !t.DeskData.Pub {
			t.saveRecord(score)
		}
		//结束连庄处理
		t.dealerOver()
		//重置状态
		t.gameOverInit()
		//踢出不足坐下玩家或超额玩家
		t.limitOver()
		//关闭房间
		t.gameStop()
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
		return
	}
	if len(t.roles) != 0 {
		return
	}
	//停止服务
	msg1 := new(pb.ServeStop)
	t.selfPid.Tell(msg1)
}

//.

//'结束连庄处理,赢家当庄
func (t *Desk) dealerOver() {
	t.DeskGame.DealerSeat = t.DeskHua.ActSeat
	if val, ok := t.seats[t.DeskGame.DealerSeat]; ok {
		t.DeskGame.Dealer = val.Userid
	}
}

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
	}
}

//.

//'结算
func (t *Desk) jiesuan2(ltype int32, score map[uint32]int64) {
	for k, v := range score {
		userid := t.getUserid(k)
		//抽水
		v = t.drawcoin(userid, v)
		switch t.DeskData.Rtype {
		case int32(pb.ROOM_TYPE0): //自由
			if v > 0 {
				t.sendCoin(userid, v, ltype)
			}
		case int32(pb.ROOM_TYPE1): //私人
			if v > 0 {
				t.sendCoin(userid, v, ltype)
			}
			t.DeskPriv.PrivScore[userid] += v
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
		msg := &pb.SJHPushDrawCoin{
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
