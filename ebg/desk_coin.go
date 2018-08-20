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

	"gohappy/game/algo"
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
	//TODO 检测是否在其它房间内,如果在则通过房间同步,否则正常同步
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
func (t *Desk) coinEnterMsg(userid string) *pb.SEBCoinEnterRoom {
	msg := new(pb.SEBCoinEnterRoom)
	//房间数据
	msg.Roominfo = handler.PackEBCoinRoom(t.DeskData)
	t.roomDataMsg(msg.Roominfo)
	//坐下玩家信息
	msg.Userinfo = t.coinSeatBetsMsg(userid)
	//位置下注信息
	msg.Betsinfo = t.coinBetsMsg()
	return msg
}

func (t *Desk) roomDataMsg(msg *pb.EBRoomData) {
	msg.State = t.state
	msg.Number = uint32(len(t.roles))
	msg.Dealer = t.DeskGame.DealerSeat
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
		switch t.state {
		case int32(pb.STATE_READY):
			tt := ReadyTime - t.timer
			if tt < 0 {
				tt = 0
			}
			msg.Timer = uint32(tt)
		default:
			tt := BetTime - t.timer
			if tt < 0 {
				tt = 0
			}
			msg.Timer = uint32(tt)
		}
	case int32(pb.ROOM_TYPE1): //私人
		msg.Rest = t.DeskData.Round - t.DeskGame.Round
		switch t.state {
		case int32(pb.STATE_READY):
		default:
			tt := BetTime - t.timer
			if tt < 0 {
				tt = 0
			}
			msg.Timer = uint32(tt)
		}
	}
}

//进入消息
func (t *Desk) coinCameinMsg(userid string) {
	msg := new(pb.SEBCamein)
	msg.Userinfo = t.coinRoleMsg(userid)
	t.broadcast(msg)
}

//召唤机器人
func (t *Desk) loadRobot() {
	t.robotTime++
	if t.robotTime < 10 {
		return
	}
	t.robotTime = 0
	t.callRobot()
}

func (t *Desk) callRobot() {
	if len(t.roles) >= 5 {
		return
	}
	r, n := t.roleCountNum()
	if r == 0 {
		return
	}
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0):
		if r >= 3 || n >= 2 {
			return
		}
	case int32(pb.ROOM_TYPE1):
		if !t.DeskData.Pub {
			return
		}
		if r >= 2 || n >= 2 {
			return
		}
	case int32(pb.ROOM_TYPE2):
	}
	msg := new(pb.RobotMsg)
	msg.Roomid = t.DeskData.Rid
	msg.Code = t.DeskData.Code
	msg.Rtype = t.DeskData.Rtype
	msg.Ltype = t.DeskData.Ltype
	msg.Gtype = t.DeskData.Gtype
	msg.EnvBet = int32(t.DeskData.Multiple)
	msg.Num = 1
	t.dbmsPid.Tell(msg)
}

func (t *Desk) roleCountNum() (r, n int) {
	for _, v := range t.roles {
		if v.User.GetRobot() {
			n++
		} else {
			r++
		}
	}
	return
}

//位置上玩家数据
func (t *Desk) coinRoleMsg(userid string) (msg *pb.EBRoomUser) {
	if v, ok := t.roles[userid]; ok {
		if v.Seat == 0 {
			return //没有坐下不广播
		}
		msg = handler.PackEBCoinUser(v.User)
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
func (t *Desk) coinSeatBetsMsg(userid string) (msg []*pb.EBRoomUser) {
	for k, v := range t.seats {
		glog.Debugf("coinSeatBetsMsg %#v, %d", v, k)
		msg2 := t.coinRoleMsg(v.Userid)
		if msg2 == nil {
			continue
		}
		//自己手牌
		if v.Userid == userid {
			t.getCardsMsg(k, msg2)
		}
		msg = append(msg, msg2)
	}
	return
}

func (t *Desk) getCardsMsg(k uint32, msg2 *pb.EBRoomUser) {
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE0),
		int32(pb.DESK_TYPE1),
		int32(pb.DESK_TYPE2): //普通
		switch t.state {
		case int32(pb.STATE_DEALER):
		case int32(pb.STATE_BET), int32(pb.STATE_NIU):
			msg2.Cards = t.getHandCards(k)
		}
	case int32(pb.DESK_TYPE3): //疯狂
		msg2.Cards = t.getHandCards(k)
	}
}

//玩家下注数据
func (t *Desk) coinBetsMsg() (msg []*pb.EBRoomBets) {
	for k, v := range t.seats {
		msg2 := &pb.EBRoomBets{
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

//下注超时
func (t *Desk) betTimeout() {
	for k, v := range t.seats {
		if !v.Ready {
			continue
		}
		//庄家不下注
		if v.Userid == t.DeskGame.Dealer {
			continue
		}
		//玩家已经下注
		if v.Bet != 0 {
			continue
		}
		//默认下1注
		v.Bet = 1
		msg := new(pb.SEBBet)
		msg.Seat = k
		msg.Seatbet = k
		msg.Value = 1
		t.broadcast(msg)
	}
	//发最后一张
	t.deal()
	//等待提交组合
	t.timer = 0
	t.state = int32(pb.STATE_NIU)
	t.pushState()
	//直接结束
	//t.niuTimeout()
}

//提交组合超时,结束
func (t *Desk) niuTimeout() {
	for k, v := range t.seats {
		if !v.Ready {
			continue
		}
		//已经提交
		if v.Niu {
			continue
		}
		num := algo.Ebg(t.DeskData.Mode, v.Cards)
		v.Power = num
		v.Niu = true
		msg := new(pb.SEBiu)
		msg.Seat = k
		msg.Value = num
		msg.Cards = v.Cards
		t.broadcast(msg)
	}
	t.gameOver() //结束游戏
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
	//洗牌
	t.shuffle()
	//发牌
	t.deal()
	//等待玩家操作
	//打庄
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE3): //抢庄
		//等待抢庄
		t.pushState()
	default:
		t.dealerHandler()
	}
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
	t.DeskGame.BetNum = 0
	t.gameInit()
	//t.timer = 0
	//结算
	t.timer = 0
	t.state = int32(pb.STATE_READY) //设置房间状态
	t.pushState()
}

//初始化
func (t *Desk) gameInit() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1): //私人
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

//'全部抢庄完成
func (t *Desk) choiceDealerOver() {
	for _, v := range t.seats {
		if v.Ready && v.BeDealer == 0 {
			return
		}
	}
	//打庄
	t.dealerHandler()
}

//.

//' 打庄处理
func (t *Desk) dealerHandler() {
	//选择庄家
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE0): //随机
		t.dealer0()
	case int32(pb.DESK_TYPE1): //固定庄
		t.dealer1()
	case int32(pb.DESK_TYPE2): //轮流庄
		t.dealer2()
	case int32(pb.DESK_TYPE3): //抢庄
		t.dealer3()
	}
	glog.Debugf("dealer -> %s", t.DeskGame.Dealer)
	glog.Debugf("dealer seat -> %d", t.DeskGame.DealerSeat)
	t.timer = 0
	t.state = int32(pb.STATE_BET) //切换状态
	t.pushDealer()
	t.pushState()
}

//状态消息
func (t *Desk) pushState() {
	msg := &pb.SEBPushState{
		State: t.state,
	}
	t.broadcast(msg)
}

//庄家消息
func (t *Desk) pushDealer() {
	msg := &pb.SEBPushDealer{
		DealerSeat: t.DeskGame.DealerSeat,
	}
	t.broadcast(msg)
}

//.

//'看牌抢庄
func (t *Desk) dealer3() []uint32 {
	a := make([]uint32, 0) //抢庄位置
	b := make([]uint32, 0) //抢庄最大倍数位置
	var max uint32         //最大倍数
	for k, v := range t.seats {
		if !v.Ready {
			continue
		}
		if v.BeDealer == 2 {
			continue
		}
		a = append(a, k)
		if v.DealerN > max {
			b = make([]uint32, 0)
			b = append(b, k)
			max = v.DealerN
		} else if v.DealerN == max {
			b = append(b, k)
		}
	}
	if len(b) != 0 {
		seat := b[rand.Intn(len(b))]
		if val, ok := t.seats[seat]; ok {
			t.DeskGame.Dealer = val.Userid
			t.DeskGame.DealerSeat = seat
		}
		return a
	}
	if len(a) == 0 { //没有一人提交抢
		for k, v := range t.seats {
			//准备的人才是参与者
			if v.Ready {
				a = append(a, k)
			}
		}
	}
	seat := a[rand.Intn(len(a))]
	if val, ok := t.seats[seat]; ok {
		t.DeskGame.Dealer = val.Userid
		t.DeskGame.DealerSeat = seat
	}
	return a
}

//.

//'牛牛坐庄
func (t *Desk) dealer1() {
	if t.DeskGame.DealerSeat != 0 {
		if v, ok := t.seats[t.DeskGame.DealerSeat]; ok {
			//庄家位置在游戏中
			if v.Ready {
				return
			}
		}
	}
	t.dealer0()
}

func (t *Desk) dealer0() {
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

func (t *Desk) dealer2() {
	if t.DeskGame.DealerSeat != 0 {
		var n uint32
		for {
			seat := t.nextSeat(t.DeskGame.DealerSeat)
			if val, ok := t.seats[seat]; ok {
				//庄家位置在游戏中
				if val.Ready {
					t.DeskGame.Dealer = val.Userid
					t.DeskGame.DealerSeat = seat
					return
				}
			}
			n++
			if n >= t.DeskData.Count {
				break
			}
		}
	}
	t.dealer0()
}

func (t *Desk) nextSeat(seat uint32) uint32 {
	if seat == t.DeskData.Count {
		return 1
	}
	return seat + 1
}

//.

//'全部下注完开始下步操作
func (t *Desk) choiceBetOver() {
	//准备的人才是参与者, 通比所有人下注才开始
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE0), //看牌抢庄
		int32(pb.DESK_TYPE1), //
		int32(pb.DESK_TYPE3), //抢庄
		int32(pb.DESK_TYPE2): //抢庄看牌
		for _, v := range t.seats {
			if v.Ready && v.Bet == 0 &&
				v.Userid != t.DeskGame.Dealer {
				return
			}
		}
	}
	t.deal() //发最后一张
	//等待提交组合
	t.timer = 0
	t.state = int32(pb.STATE_NIU) //切换状态
	t.pushState()
	//直接结束
	//t.niuTimeout()
}

//.

//'发牌
func (t *Desk) deal() {
	var hand int
	for k, v := range t.seats {
		//准备的人才是参与者
		if !v.Ready {
			continue
		}
		//(通比牛牛 | 牛牛坐庄)前4张牌不广播
		switch t.DeskData.Dtype {
		case int32(pb.DESK_TYPE3): //疯狂,抢庄
			hand = 1
			//发牌消息
			cards := make([]uint32, hand, hand)
			copy(cards, t.DeskGame.Cards[:hand])
			t.DeskGame.Cards = t.DeskGame.Cards[hand:]
			v.Cards = append(v.Cards, cards...)
			switch t.state {
			case int32(pb.STATE_DEALER):
				//一明一暗
				cards2 := make([]uint32, hand+1, hand+1)
				msg2 := resDraw(k, t.state, cards2)
				t.broadcast3(k, msg2)
				//可看到自己牌值
				cards3 := append(v.Cards, 0)
				msg := resDraw(k, t.state, cards3)
				t.send2seat(k, msg)
			case int32(pb.STATE_BET):
				//可看到自己牌值
				msg := resDraw(k, t.state, v.Cards)
				t.send2seat(k, msg)
			}
		case int32(pb.DESK_TYPE2), //
			int32(pb.DESK_TYPE1), //
			int32(pb.DESK_TYPE0): //普通,随机庄、固定庄、轮流庄
			hand = 2
			switch t.state {
			case int32(pb.STATE_DEALER):
				//二暗,看不到牌值,发牌消息
				cards2 := make([]uint32, hand, hand)
				msg2 := resDraw(k, t.state, cards2)
				t.broadcast(msg2)
			case int32(pb.STATE_BET):
				//发牌
				cards := make([]uint32, hand, hand)
				copy(cards, t.DeskGame.Cards[:hand])
				t.DeskGame.Cards = t.DeskGame.Cards[hand:]
				v.Cards = append(v.Cards, cards...)
				//自己可看到牌值
				msg := resDraw(k, t.state, cards)
				t.send2seat(k, msg)
			}
		}
	}
}

//.

//'自由场结算消息
func (t *Desk) resCoinOver(score map[uint32]int64) (msg *pb.SEBCoinGameover) {
	msg = &pb.SEBCoinGameover{
		Dealer: t.DeskGame.Dealer,
		State:  t.state,
	}
	for k, v := range score {
		d := &pb.EBCoinOver{
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
func (t *Desk) resOver(score map[uint32]int64) (msg *pb.SEBGameover) {
	msg = &pb.SEBGameover{
		Dealer:     t.DeskGame.Dealer,
		DealerSeat: t.DeskGame.DealerSeat,
		Round:      t.DeskGame.Round,
		//LeftRound:  (t.DeskData.Round - t.DeskGame.Round),
	}
	if t.DeskData.Round > t.DeskGame.Round {
		msg.LeftRound = (t.DeskData.Round - t.DeskGame.Round)
	}
	for k, v := range score {
		d := &pb.EBRoomOver{
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

// vim: set foldmethod=marker foldmarker=//',//.:
