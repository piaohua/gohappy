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
func (t *Desk) coinEnterMsg(userid string) *pb.SSGCoinEnterRoom {
	msg := new(pb.SSGCoinEnterRoom)
	//房间数据
	msg.Roominfo = handler.PackSGCoinRoom(t.DeskData)
	t.roomDataMsg(msg.Roominfo)
	//坐下玩家信息
	msg.Userinfo = t.coinSeatBetsMsg(userid)
	//位置下注信息
	msg.Betsinfo = t.coinBetsMsg()
	return msg
}

func (t *Desk) roomDataMsg(msg *pb.SGRoomData) {
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
	msg := new(pb.SSGCamein)
	msg.Userinfo = t.coinRoleMsg(userid)
	t.broadcast(msg)
}

//位置上玩家数据
func (t *Desk) coinRoleMsg(userid string) (msg *pb.SGRoomUser) {
	if v, ok := t.roles[userid]; ok {
		if v.Seat == 0 {
			return //没有坐下不广播
		}
		msg = handler.PackSGCoinUser(v.User)
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
func (t *Desk) coinSeatBetsMsg(userid string) (msg []*pb.SGRoomUser) {
	for k, v := range t.seats {
		msg2 := t.coinRoleMsg(v.Userid)
		if msg2 == nil {
			continue
		}
		//自己手牌
		if v.Userid == userid {
			switch t.DeskData.Dtype {
			case int32(pb.DESK_TYPE0): //看牌抢庄
				msg2.Cards = t.getHandCards(k)
			default:
				cs := t.getHandCards(k)
				if len(cs) == 3 {
					msg2.Cards = cs
				}
			}
		}
		msg = append(msg, msg2)
	}
	return
}

//玩家下注数据
func (t *Desk) coinBetsMsg() (msg []*pb.SGRoomBets) {
	for k, v := range t.seats {
		msg2 := &pb.SGRoomBets{
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
	var num int
	for _, v := range t.seats {
		if v.Ready {
			num++
		}
	}
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
	var num int
	for _, v := range t.seats {
		if v.Ready {
			num++
		}
	}
	if num >= 2 {
		t.gameStart() //开始牌局
		return
	}
	//房间人数为0时解散
	t.checkPubOver()
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
		if v.Bet != 0 {
			continue
		}
		//默认下1注
		v.Bet = 1
		msg := new(pb.SSGBet)
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
		num := algo.San(v.Cards)
		v.Power = num
		v.Niu = true
		msg := new(pb.SSGiu)
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
	//抽水
	t.drawfee()
	//初始化
	t.gameStartInit()
	//洗牌
	t.shuffle()
	//发牌
	t.deal()
	//等待玩家操作
	//打庄, (通比牛牛 | 牛牛坐庄)
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE1): //通比牛牛
		t.dealerHandler()
		switch t.DeskData.Mode {
		case 0://普通
			t.betTimeout() //去掉下注流程
		default:
		}
	default:
		t.pushState()
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
	case int32(pb.DESK_TYPE0): //看牌抢庄
		t.dealer1()
	case int32(pb.DESK_TYPE1): //通比牛牛
	case int32(pb.DESK_TYPE2): //抢庄看牌
		//t.dealer4()
		t.dealer1()
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
	msg := &pb.SSGPushState{
		State: t.state,
	}
	t.broadcast(msg)
}

//庄家消息
func (t *Desk) pushDealer() {
	msg := &pb.SSGPushDealer{
		DealerSeat: t.DeskGame.DealerSeat,
	}
	t.broadcast(msg)
}

//.

//'看牌抢庄
func (t *Desk) dealer1() []uint32 {
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

//'全部下注完开始下步操作
func (t *Desk) choiceBetOver() {
	//准备的人才是参与者, 通比所有人下注才开始
	switch t.DeskData.Dtype {
	case int32(pb.DESK_TYPE0), //看牌抢庄
		int32(pb.DESK_TYPE2): //抢庄看牌
		for _, v := range t.seats {
			if v.Ready && v.Bet == 0 &&
				v.Userid != t.DeskGame.Dealer {
				return
			}
		}
	case int32(pb.DESK_TYPE1): //通比牛牛
		for _, v := range t.seats {
			if v.Ready && v.Bet == 0 {
				return
			}
		}
	}
	t.deal() //发最后一张
	//等待提交组合
	t.timer = 0
	t.state = int32(pb.STATE_NIU) //切换状态
	t.pushState()
}

//.

//'发牌
func (t *Desk) deal() {
	var hand int
	switch t.state {
	case int32(pb.STATE_DEALER):
		hand = 2
	case int32(pb.STATE_BET):
		hand = 1
	}
	if hand == 0 {
		glog.Errorf("deal err: %d", t.state)
		return
	}
	for k, v := range t.seats {
		if !v.Ready {
			continue
		}
		cards := make([]uint32, hand, hand)
		//准备的人才是参与者
		copy(cards, t.DeskGame.Cards[:hand])
		t.DeskGame.Cards = t.DeskGame.Cards[hand:]
		v.Cards = append(v.Cards, cards...)
		//发牌消息
		//(通比牛牛 | 牛牛坐庄)前2张牌不广播
		switch t.DeskData.Dtype {
		case int32(pb.DESK_TYPE2), //抢庄看牌
			int32(pb.DESK_TYPE1): //通比牛牛
			if hand == 2 {
				//看不到牌值
				cards2 := make([]uint32, hand+1, hand+1)
				msg := resDraw(k, t.state, cards2)
				t.broadcast(msg)
			} else {
				//cards2 := make([]uint32, hand, hand)
				//msg2 := resDraw(k, t.state, cards2)
				//t.broadcast3(k, msg2)
				//自己可看到全部牌值
				msg := resDraw(k, t.state, v.Cards)
				t.send2seat(k, msg)
			}
		case int32(pb.DESK_TYPE0): //看牌抢庄
			//看不到别人的牌值
			cards2 := make([]uint32, hand, hand)
			msg2 := resDraw(k, t.state, cards2)
			t.broadcast3(k, msg2)
			//自己可看到牌值
			msg := resDraw(k, t.state, cards)
			t.send2seat(k, msg)
		}
	}
}

//.

//'自由场结算消息
func (t *Desk) resCoinOver(score map[uint32]int64) (msg *pb.SSGCoinGameover) {
	msg = &pb.SSGCoinGameover{
		Dealer: t.DeskGame.Dealer,
		State:  t.state,
	}
	for k, v := range score {
		d := &pb.SGCoinOver{
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
func (t *Desk) resOver(score map[uint32]int64) (msg *pb.SSGGameover) {
	msg = &pb.SSGGameover{
		Dealer:     t.DeskGame.Dealer,
		DealerSeat: t.DeskGame.DealerSeat,
		Round:      t.DeskGame.Round,
		//LeftRound:  (t.DeskData.Round - t.DeskGame.Round),
	}
	if t.DeskData.Round > t.DeskGame.Round {
		msg.LeftRound = (t.DeskData.Round - t.DeskGame.Round)
	}
	for k, v := range score {
		d := &pb.SGRoomOver{
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
