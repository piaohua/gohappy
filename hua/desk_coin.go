/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-04-30 18:01:29
 * Filename      : desk_coin.go
 * Description   : 金币场玩牌逻辑
 * *******************************************************/
package main

import (
	"math/rand"

	"gohappy/data"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
)

//'更新金币
func (t *Desk) sendCoin(userid string, num int64, ltype int32) {
	if num == 0 {
		return
	}
	//玩家在线
	if v, ok := t.roles[userid]; ok && v != nil {
		v.User.AddCoin(num)
		//在线
		if !v.Offline {
			//货币变更及时同步
			msg := &pb.ChangeCurrency{
				Userid: userid,
				Coin:   num,
				Type:   ltype,
			}
			v.Pid.Tell(msg)
			return
		}
	}
	//离线同步数据
	glog.Infof("sendCoin userid %s, num %d, ltype %d",
		userid, num, ltype)
	msg := &pb.OfflineCurrency{
		Userid: userid,
		Coin:   num,
		Type:   ltype,
	}
	//通过大厅通知其它节点
	t.rolePid.Tell(msg)
}

func (t *Desk) sendDiamond(userid string, num int64, ltype int32) {
	if num == 0 {
		return
	}
	//玩家在线
	if v, ok := t.roles[userid]; ok && v != nil {
		v.User.AddDiamond(num)
		//在线
		if !v.Offline {
			//货币变更及时同步
			msg := &pb.ChangeCurrency{
				Userid:  userid,
				Diamond: num,
				Type:    ltype,
			}
			v.Pid.Tell(msg)
			return
		}
	}
	//离线同步数据
	glog.Infof("sendDiamond userid %s, num %d, ltype %d",
		userid, num, ltype)
	msg := &pb.OfflineCurrency{
		Userid:  userid,
		Diamond: num,
		Type:    ltype,
	}
	//通过大厅通知其它节点
	t.rolePid.Tell(msg)
}

//.

//'进入房间响应消息
func (t *Desk) coinEnterMsg(userid string) *pb.SJHCoinEnterRoom {
	msg := new(pb.SJHCoinEnterRoom)
	//房间数据
	msg.Roominfo = handler.PackJHCoinRoom(t.DeskData)
	msg.Roominfo.State = t.state
	//坐下玩家信息
	msg.Userinfo = t.coinSeatBetsMsg(userid)
	//位置下注信息
	msg.Betsinfo = t.coinBetsMsg()
	return msg
}

//进入消息
func (t *Desk) coinCameinMsg(userid string) {
	msg := new(pb.SJHCamein)
	msg.Userinfo = t.coinRoleMsg(userid)
	t.broadcast(msg)
}

//位置上玩家数据
func (t *Desk) coinRoleMsg(userid string) (msg *pb.JHRoomUser) {
	if v, ok := t.roles[userid]; ok {
		if v.Seat == 0 {
			return //没有坐下不广播
		}
		msg = handler.PackJHCoinUser(v.User)
		msg.Seat = v.Seat
		msg.Offline = v.Offline
		if val, ok := t.seats[v.Seat]; ok {
			msg.Dealer = val.BeDealer
			msg.Bet = val.Bet
			msg.Num = val.DealerN
			msg.Niu = val.Niu
			msg.Ready = val.Ready
		}
		if t.DeskPriv != nil {
			msg.Score = t.DeskPriv.PrivScore[userid]
		}
	}
	return
}

//所有坐下玩家数据
func (t *Desk) coinSeatBetsMsg(userid string) (msg []*pb.JHRoomUser) {
	for k, v := range t.seats {
		msg2 := t.coinRoleMsg(v.Userid)
		if msg2 == nil {
			continue
		}
		//自己手牌
		if v.Userid == userid && t.isSee(k) {
			msg2.Cards = t.getHandCards(k)
		}
		msg = append(msg, msg2)
	}
	return
}

//玩家下注数据
func (t *Desk) coinBetsMsg() (msg []*pb.JHRoomBets) {
	for k, v := range t.seats {
		msg2 := &pb.JHRoomBets{
			Seat: k,
			Bets: v.Bet,
		}
		msg = append(msg, msg2)
	}
	return
}

//.

//'是否全部准备状态
func (t *Desk) allReady() bool {
	var num int = t.readyNum()
	if num != len(t.roles) {
		return false
	}
	//准备人数大于2
	if num < 2 {
		return false
	}
	//全部准备立即开始
	return true
}

//.

//' 超时操作

//准备超时,不等待全部准备
func (t *Desk) readyTimeout() {
	var num int = t.readyNum()
	if num >= 2 {
		t.gameStart() //开始牌局
		return
	}
	//房间人数为0时解散
	t.checkPubOver()
}

//准备人数
func (t *Desk) readyNum() (num int) {
	for _, v := range t.seats {
		if v.Ready {
			num++
		}
	}
	return
}

//操作超时放弃
func (t *Desk) betTimeout() {
	seat := t.DeskHua.ActSeat
	userid := t.getUserid(seat)
	t.coinFold(userid)
}

//.

//' 开始游戏
func (t *Desk) gameStart() {
	//未参与玩家离开位置
	t.startSitup()
	//抽水
	t.drawfee()
	//初始化
	t.gameStartInit()
	//打庄
	t.dealerHandler()
	//洗牌
	t.shuffle()
	//发牌
	t.deal()
	//初始化操作
	t.initAct()
}

//开始初始化
func (t *Desk) gameStartInit() {
	//设置房间抢庄状态
	t.timer = 0
	t.state = int32(pb.STATE_DEALER)
	t.gameInit()
}

//结束重置
func (t *Desk) gameOverInit() {
	for _, v := range t.seats {
		v.Ready = false
	}
	t.DeskGame.Cards = make([]uint32, 0)
	t.gameInit()
	//t.timer = 0
	//结算加长5秒
	t.timer = -5
	t.state = int32(pb.STATE_READY) //设置房间状态
	t.pushState()
}

//初始化
func (t *Desk) gameInit() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1): //私人
		t.DeskHua = new(data.DeskHua)
		t.DeskHua.ActSeats = make(map[uint32]*data.ActStatus)
		t.DeskHua.ActCallNum = 0
		t.DeskHua.ActRaiseNum = 0
		t.DeskHua.ActTimes = 0
		for _, v := range t.seats {
			//v.Ready = false
			v.BeDealer = 0
			v.DealerN = 0
			v.Bet = 0
			v.Cards = make([]uint32, 0)
			v.Power = 0
			v.Niu = false
		}
	case int32(pb.ROOM_TYPE2): //百人
	}
}

// 未参与玩家离开位置
func (t *Desk) startSitup() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1): //私人
		for _, v := range t.seats {
			if v.Ready {
				continue
			}
			//玩家站起
			t.roleSitUp(v.Userid)
		}
	}
}

//.

//' 打庄处理
func (t *Desk) dealerHandler() {
	//选择庄家
	t.dealer4()
	glog.Debugf("dealer -> %s", t.DeskGame.Dealer)
	glog.Debugf("dealer seat -> %d", t.DeskGame.DealerSeat)
	t.timer = 0
	t.state = int32(pb.STATE_BET) //切换状态
	t.pushDealer()
	t.pushState()
}

//状态消息
func (t *Desk) pushState() {
	msg := &pb.SJHPushState{
		State: t.state,
	}
	t.broadcast(msg)
}

//庄家消息
func (t *Desk) pushDealer() {
	msg := &pb.SJHPushDealer{
		DealerSeat: t.DeskGame.DealerSeat,
	}
	t.broadcast(msg)
}

//广播操作状态消息
func (t *Desk) pushActState() {
	msg := &pb.SJHPushActState{
		State:    t.DeskHua.ActState,
		Seat:     t.DeskHua.ActSeat,
		Pot:      t.DeskGame.BetNum,
		CallNum:  t.DeskHua.ActCallNum,
		RaiseNum: t.DeskHua.ActRaiseNum,
	}
	t.broadcast(msg)
}

//.

//' 随机选择庄
func (t *Desk) dealer4() {
	if t.DeskGame.DealerSeat != 0 {
		if v, ok := t.seats[t.DeskGame.DealerSeat]; ok {
			//庄家位置在游戏中
			if v.Ready {
				return
			}
		}
	}
	a := make([]uint32, 0)
	for k, v := range t.seats {
		//准备的人才是参与者
		if v.Ready {
			a = append(a, k)
		}
	}
	if len(a) == 0 {
		glog.Errorf("dealer4 err: %#v", t.seats)
		return
	}
	seat := a[rand.Intn(len(a))]
	if val, ok := t.seats[seat]; ok {
		t.DeskGame.Dealer = val.Userid
		t.DeskGame.DealerSeat = seat
	}
}

//.

//'发牌
func (t *Desk) deal() {
	var hand = 3
	for k, v := range t.seats {
		if !v.Ready {
			continue
		}
		v.Cards = make([]uint32, hand, hand)
		//准备的人才是参与者
		copy(v.Cards, t.DeskGame.Cards[:hand])
		t.DeskGame.Cards = t.DeskGame.Cards[hand:]
		//发牌消息
		//看不到牌值
		cards2 := make([]uint32, hand, hand)
		msg := resDraw(k, t.state, cards2)
		t.broadcast(msg)
	}
}

//.

//'自由场结算消息
func (t *Desk) resCoinOver(score map[uint32]int64) (msg *pb.SJHCoinGameover) {
	msg = &pb.SJHCoinGameover{
		Dealer: t.DeskGame.Dealer,
		State:  t.state,
	}
	for k, v := range score {
		d := &pb.JHCoinOver{
			Seat:  k,
			Score: v,
		}
		if val, ok := t.seats[k]; ok {
			d.Bets = val.Bet
			d.Value = val.Power
			d.Cards = val.Cards
			if p, ok2 := t.roles[val.Userid]; ok2 {
				d.Coin = p.User.GetCoin()
				d.Nickname = p.User.GetNickname()
				d.Photo = p.User.GetPhoto()
			}
		}
		msg.Data = append(msg.Data, d)
	}
	return
}

//.

//'私人局结算消息
func (t *Desk) resOver(score map[uint32]int64) (msg *pb.SJHGameover) {
	msg = &pb.SJHGameover{
		Dealer:     t.DeskGame.Dealer,
		DealerSeat: t.DeskGame.DealerSeat,
		Round:      t.DeskGame.Round,
		//LeftRound:  (t.DeskData.Round - t.DeskGame.Round),
	}
	if t.DeskData.Round > t.DeskGame.Round {
		msg.LeftRound = (t.DeskData.Round - t.DeskGame.Round)
	}
	for k, v := range score {
		d := &pb.JHRoomOver{
			Seat:  k,
			Score: v,
		}
		if val, ok := t.seats[k]; ok {
			d.Bets = val.Bet
			d.Value = val.Power
			d.Cards = val.Cards
			d.Total = t.DeskPriv.PrivScore[val.Userid]
			if p, ok2 := t.roles[val.Userid]; ok2 {
				d.Coin = p.User.GetCoin()
				d.Nickname = p.User.GetNickname()
				d.Photo = p.User.GetPhoto()
			}
		}
		msg.Data = append(msg.Data, d)
	}
	return
}

//.

//下个操作位置
func (t *Desk) getNextActSeat() uint32 {
	seat := t.DeskHua.ActSeat
	if seat == 0 {
		seat = t.DeskGame.DealerSeat
	}
	var i uint32 = seat
	for {
		var i uint32 = t.nextSeat(i)
		if i == seat {
			break
		}
		if t.qualified(i) {
			return i
		}
	}
	return 0
}

//切换下个操作位置
func (t *Desk) setNextActSeat() {
	var last, curr uint32
	last = t.DeskHua.ActSeat
	curr = t.getNextActSeat()
	//结束,winner = last
	if curr == 0 {
		t.gameOver()
		return
	}
	//轮数计数
	if t.roundOver(last, curr) {
		t.DeskHua.ActTimes++
		//TODO 底池操作
	}
	if t.DeskHua.ActTimes == 20 {
		//TODO 结束限制
		t.gameOver()
		return
	}
	t.timer = 0
	t.DeskHua.ActSeat = curr
	//设置下家操作值
	t.setNextActState()
	//广播下家操作状态消息
	t.pushActState()
}

//是否结束一轮操作
func (t *Desk) roundOver(last, curr uint32) bool {
	//庄家操作完
	if last == t.DeskGame.DealerSeat {
		return true
	}
	//庄家已经输掉
	for {
		i := t.nextSeat(last)
		if i == curr {
			break
		}
		if i == t.DeskGame.DealerSeat {
			return true
		}
	}
	return false
}

//下家的位置
func (t *Desk) nextSeat(seat uint32) uint32 {
	if seat >= t.DeskData.Count {
		return 1
	}
	return seat + 1
}

//是否合格,是否已经放弃或者比牌输掉
func (t *Desk) qualified(seat uint32) bool {
	if v, ok := t.DeskHua.ActSeats[seat]; ok {
		if v.Alive {
			return true
		}
	}
	return false
}

//设置下个位置操作值
func (t *Desk) setNextActState() {
	var val int32
	val |= int32(pb.ACT_FOLD)
	//第一轮不能比
	if t.DeskHua.ActTimes != 0 {
		//TODO 比牌金币限制
		val |= int32(pb.ACT_BI)
	}
	if v, ok := t.DeskHua.ActSeats[t.DeskHua.ActSeat]; ok {
		if !v.See {
			val |= int32(pb.ACT_SEE)
		}
	}
	//TODO 加注和跟注上限
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		user := t.getUserBySeat(t.DeskHua.ActSeat)
		if user != nil {
			if user.GetCoin() >= t.DeskHua.ActRaiseNum {
				//为下注的指定倍数时才可以加注
				val |= int32(pb.ACT_RAISE)
			}
		}
	case int32(pb.ROOM_TYPE1): //私人
		val |= int32(pb.ACT_RAISE)
	}
	//上家下注大于当前位置下注
	if t.DeskHua.ActCallNum > 0 {
		val |= int32(pb.ACT_CALL)
	}
	t.DeskHua.ActState = val
}

//开始游戏初始化操作
func (t *Desk) initAct() {
	for k, v := range t.seats {
		if !v.Ready {
			continue
		}
		d := new(data.ActStatus)
		d.Alive = true
		t.DeskHua.ActSeats[k] = d
		//开始时下暗注
		t.setBet(k, v.Userid, int64(t.DeskData.Ante))
	}
	//初始化玩家操作
	//TODO 加注和跟注下限
	t.DeskHua.ActRaiseNum = int64(t.DeskData.Ante)
	t.DeskHua.ActCallNum = 0
	t.DeskHua.ActSeat = t.getNextActSeat()
	//设置下家操作值
	t.setNextActState()
	//广播下家操作状态消息
	t.pushActState()
}

//下注成功设置
func (t *Desk) setBet(seat uint32, userid string, num int64) {
	if v, ok := t.seats[seat]; ok {
		v.Bet += num
	}
	if v, ok := t.DeskHua.ActSeats[seat]; ok {
		v.ActNum += num
	}
	t.DeskGame.BetNum += num
	t.setBetMsg(seat, userid, num)
}

//开始时下暗注
func (t *Desk) setBetMsg(seat uint32, userid string, num int64) {
	msg := &pb.SJHCoinRaise{
		Seat:   seat,
		Userid: userid,
		Value:  num, //下底注
		Pot:    t.DeskGame.BetNum,
	}
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		t.sendCoin(userid, (-1 * num), int32(pb.LOG_TYPE5))
	case int32(pb.ROOM_TYPE1): //私人
		t.sendCoin(userid, (-1 * num), int32(pb.LOG_TYPE5))
	}
	t.broadcast(msg)
}

// vim: set foldmethod=marker foldmarker=//',//.:
