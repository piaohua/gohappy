/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-04-30 18:01:29
 * Filename      : desk_free.go
 * Description   : 百人场玩牌逻辑
 * *******************************************************/
package main

import (
	"math/rand"
	"time"

	"gohappy/data"
	"gohappy/game/algo"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"
)

//'初始化
func (t *Desk) freeInit() {
	t.DeskGame.BetNum = 0
	t.DeskFree.Cards = make(map[uint32][]uint32) //手牌
	t.DeskFree.Power = make(map[uint32]uint32)   //牌力
	//userid:num, 玩家下注金额
	t.DeskFree.Bets = make(map[string]int64)
	//seat:num, 位置下注金额
	t.DeskFree.SeatBets = make(map[uint32]int64)
	//结果 seat:num,seat=(1,2,3,4,5),倍数
	t.DeskFree.Multiple = make(map[uint32]int64)
	//位置(1-5)输赢总量
	t.DeskFree.Score1 = make(map[uint32]int64)
	//每个闲家输赢总量
	t.DeskFree.Score2 = make(map[string]int64)
	//位置(1-5)上每个玩家输赢
	t.DeskFree.Score3 = make(map[uint32]map[string]int64)
	//位置下注详细
	t.DeskFree.SeatRoleBets = make(map[uint32]map[string]int64)
}

//.

//'进入房间响应消息
func (t *Desk) freeEnterMsg(userid string) *pb.SLHFreeEnterRoom {
	msg := new(pb.SLHFreeEnterRoom)
	//房间数据
	msg.Roominfo = handler.PackLHFreeRoom(t.DeskData)
	t.freeRoomDataMsg(msg.Roominfo)
	//坐下玩家信息
	msg.Userinfo = t.freeSeatBetsMsg()
	//位置下注信息
	msg.Betsinfo = handler.PackLHRoomBets(t.DeskFree.SeatBets)
	return msg
}

func (t *Desk) freeRoomDataMsg(msg *pb.LHFreeRoom) {
	msg.State = t.state
	msg.Dealer = t.DeskGame.DealerSeat
	msg.Userid = t.DeskGame.Dealer
	if t.DeskFree != nil && t.DeskFree.Carry > 0 {
		msg.Carry = uint32(t.DeskFree.Carry)
	}
	msg.LeftDealerNum = t.leftDealerTimes()
	msg.DealerNum = DealerTimes
	p := t.getPlayer(t.DeskGame.Dealer)
	if p != nil {
		msg.Photo = p.GetPhoto()
	}
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE2): //百人
		switch t.state {
		case int32(pb.STATE_READY):
			tt := ReadyTime - t.timer
			if tt < 0 {
				tt = 0
			}
			msg.Timer = uint32(tt)
		default:
			tt := FreeBetTime - t.timer
			if tt < 0 {
				tt = 0
			}
			msg.Timer = uint32(tt)
		}
	}
}

//进入消息
func (t *Desk) freeCameinMsg(userid string) {
	msg := new(pb.SLHFreeCamein)
	msg.Userinfo = t.freeSeatRoleMsg(userid)
	t.broadcast(msg)
}

//位置上玩家数据
func (t *Desk) freeSeatRoleMsg(userid string) (msg *pb.LHFreeUser) {
	if v, ok := t.roles[userid]; ok {
		if v.Seat == 0 {
			return //没有坐下不广播
		}
		msg = handler.PackLHFreeUser(v.User)
		msg.Seat = v.Seat
		if t.DeskFree == nil {
			return
		}
		msg.Bet = t.DeskFree.Bets[userid]
		msg.Bets = t.freeBetsMsg(userid)
	}
	return
}

//所有坐下玩家数据
func (t *Desk) freeSeatBetsMsg() (msg []*pb.LHFreeUser) {
	for _, v := range t.seats {
		msg2 := t.freeSeatRoleMsg(v.Userid)
		if msg2 == nil {
			continue
		}
		msg = append(msg, msg2)
	}
	return
}

//玩家下注数据
func (t *Desk) freeBetsMsg(userid string) (msg []*pb.LHRoomBets) {
	for i := pb.DESK_SEAT2; i <= pb.DESK_SEAT4; i++ {
		seat := uint32(i)
		bets := t.getFreeSeatBet(userid, seat)
		msg2 := &pb.LHRoomBets{
			Seat: seat,
			Bets: bets,
		}
		msg = append(msg, msg2)
	}
	return
}

//玩家位置下注数量
func (t *Desk) getFreeSeatBet(userid string, seat uint32) int64 {
	if t.DeskFree == nil {
		return 0
	}
	if m, ok := t.SeatRoleBets[seat]; ok {
		return m[userid]
	}
	return 0
}

//.

//'beDealer 没人上庄时都可以选择上庄,已经上庄的人可以补庄
//st:0下庄 1上庄 2补庄
func (t *Desk) beDealer(userid string, st int32, num uint32) pb.ErrCode {
	if !t.isFree() {
		return pb.NotDealerRoom
	}
	user := t.getPlayer(userid)
	if user == nil {
		glog.Errorf("userid %s not exist", userid)
		return pb.NotInRoom
	}
	//TODO 全部带上庄
	//num = uint32(user.GetCoin())
	switch st {
	case int32(pb.DEALER_DOWN):
		switch t.state {
		case int32(pb.STATE_BET): //下注中
			if userid == t.DeskGame.Dealer {
				t.DeskFree.DealerDown = true
				msg := handler.LhdBeDealerMsg(0, int64(num), t.DeskFree.CarryInit, t.DeskGame.Dealer,
					userid, user.GetNickname(), user.GetPhoto())
				msg.Down = true
				t.broadcast(msg)
				return pb.OK
			}
			return pb.DealerDownFail
		default:
			t.delBeDealer(userid, user)
			return pb.OK
		}
	case int32(pb.DEALER_UP):
		//已经上庄,暂时不能重复上
		if t.alreadyBeDealer(userid) {
			return pb.BeDealerAlready
		}
		//上庄限制
		if num < t.DeskData.Carry {
			return pb.BeDealerNotEnough
		}
		t.addBeDealer(userid, st, int64(num), user)
		return pb.OK
	case int32(pb.DEALER_BU):
		t.addBeDealer(userid, st, int64(num), user)
		return pb.OK
	}
	return pb.OperateError
}

//.

//'获取上庄列表消息
func (t *Desk) dealerListMsg() (msg *pb.SLHFreeDealerList) {
	msg = new(pb.SLHFreeDealerList)
	if !t.isFree() {
		msg.Error = pb.NotDealerRoom
		return
	}
	for k, v := range t.DeskFree.Dealers {
		user := t.getPlayer(k)
		if user == nil {
			glog.Errorf("userid %s not exist", k)
			continue
		}
		msg2 := &pb.LHDealerList{
			Userid:   k,
			Nickname: user.GetNickname(),
			Photo:    user.GetPhoto(),
			//Coin:     user.GetCoin(),
			Coin: v,
		}
		msg.List = append(msg.List, msg2)
	}
	return
}

//.

//'百人下注
func (t *Desk) freeBet(userid string, seatBet uint32,
	num int64) pb.ErrCode {
	if userid == t.DeskGame.Dealer { //庄家不用下注
		return pb.BetDealerFailed
	}
	if t.state != int32(pb.STATE_BET) {
		return pb.GameNotStart
	}
	if num <= 0 {
		return pb.OperateError
	}
	if !t.isFree() {
		return pb.NotDealerRoom
	}
	user := t.getPlayer(userid)
	if user == nil {
		glog.Errorf("userid %s not exist", userid)
		return pb.NotInRoom
	}
	if user.GetCoin() < num {
		return pb.NotEnoughCoin
	}
	//TODO 限制优化
	coin := user.GetCoin()          //剩余金额
	bets := t.DeskFree.Bets[userid] //已经下注额
	//本轮下注不能超过1/4
	if (num + bets) > ((coin + bets) / 4) {
		//TODO 暂时1赔1不限制
		//return pb.BetTopLimit //下注限制
	}
	switch seatBet {
	case uint32(pb.DESK_SEAT2):
		if (t.DeskFree.SeatBets[seatBet] + num) > (t.DeskFree.Carry + t.DeskFree.SeatBets[uint32(pb.DESK_SEAT3)]) {
			return pb.BetTopLimit //下注限制
		}
	case uint32(pb.DESK_SEAT3):
		if (t.DeskFree.SeatBets[seatBet] + num) > (t.DeskFree.Carry + t.DeskFree.SeatBets[uint32(pb.DESK_SEAT2)]) {
			return pb.BetTopLimit //下注限制
		}
	case uint32(pb.DESK_SEAT4):
		if (t.DeskFree.SeatBets[seatBet] + num) > (t.DeskFree.Carry / 6) {
			return pb.BetTopLimit //下注限制
		}
	default:
		return pb.OperateError
	}
	t.DeskFree.Bets[userid] += num      //个人总下注额
	t.DeskFree.SeatBets[seatBet] += num //当前位置总下注额
	t.DeskGame.BetNum += num            //当局总下注额
	//位置详细记录
	var betsNum int64 //玩家当前位置下注总数
	if m, ok := t.SeatRoleBets[seatBet]; ok {
		m[userid] += num
		t.SeatRoleBets[seatBet] = m
		betsNum = m[userid]
	} else {
		m := make(map[string]int64)
		m[userid] = num
		t.SeatRoleBets[seatBet] = m
		betsNum = m[userid]
	}
	t.sendCoin(userid, (-1 * num), int32(pb.LOG_TYPE5))
	seat := t.getSeat(userid)
	msg := resFreeBet(seat, seatBet, num,
		t.DeskFree.SeatBets[seatBet], betsNum, userid)
	t.broadcast(msg)
	return pb.OK
}

//下注消息
func resFreeBet(seat, beseat uint32, val, coin,
	bets int64, userid string) *pb.SLHFreeBet {
	return &pb.SLHFreeBet{
		Seat:   seat,
		Beseat: beseat,
		Value:  uint32(val),
		Coin:   coin,
		Bets:   bets,
		Userid: userid,
	}
}

//.

//' 百人输赢趋势
func (t *Desk) freeTrends() *pb.SLHFreeTrend {
	msg := new(pb.SLHFreeTrend)
	for _, v := range t.DeskFree.Trends {
		msg2 := &pb.LHFreeTrend{
			Seat2: v.Seat2,
			Seat3: v.Seat3,
			Seat4: v.Seat4,
			Seat5: v.Seat5,
		}
		msg.List = append(msg.List, msg2)
	}
	return msg
}

//.

//' 百人上局赢家列表
func (t *Desk) freeWiners() *pb.SLHFreeWiners {
	msg := new(pb.SLHFreeWiners)
	for _, v := range t.DeskFree.Winers {
		msg2 := &pb.LHFreeWiner{
			Userid:   v.Userid,
			Nickname: v.Nickname,
			Photo:    v.Photo,
			Coin:     v.Coin,
		}
		msg.List = append(msg.List, msg2)
	}
	return msg
}

//.

//' 百人玩家列表
func (t *Desk) freeRoles() *pb.SLHFreeRoles {
	msg := new(pb.SLHFreeRoles)
	for _, v := range t.roles {
		if v.User == nil {
			continue
		}
		msg2 := &pb.LHFreeRole{
			Userid:   v.User.GetUserid(),
			Nickname: v.User.GetNickname(),
			Photo:    v.User.GetPhoto(),
			Coin:     v.User.GetCoin(),
		}
		msg.List = append(msg.List, msg2)
	}
	return msg
}

//.
//' 百人坐下
func (t *Desk) freeSit(userid string, arg *pb.CLHSit) (msg *pb.SLHSit) {
	errcode := t.sitCheck(userid, arg)
	glog.Debugf("sit userid %s arg %#v", userid, arg)
	glog.Debugf("sit userid %s errcode %d", userid, errcode)
	if errcode != pb.OK {
		msg = new(pb.SLHSit)
		msg.Error = errcode
		return
	}
	switch arg.Type {
	case pb.SitDown:
		t.seats[arg.Seat] = &data.DeskSeat{
			Userid: userid,
		}
		if v, ok := t.roles[userid]; ok {
			v.Seat = arg.Seat
		}
	case pb.SitUp:
		delete(t.seats, arg.Seat)
		if v, ok := t.roles[userid]; ok {
			v.Seat = 0
		}
	}
	//广播消息
	msg = &pb.SLHSit{
		Type:     arg.Type,
		Seat:     arg.Seat,
		Userinfo: t.coinRoleMsg(userid),
	}
	return
}

func (t *Desk) sitCheck(userid string, arg *pb.CLHSit) pb.ErrCode {
	user := t.getPlayer(userid)
	if user == nil {
		glog.Errorf("userid %s not exist", userid)
		return pb.NotInRoom
	}
	seat := t.getSeat(userid)
	switch arg.Type {
	case pb.SitDown:
		if seat != 0 {
			glog.Errorf("sit faild %s %d", userid, seat)
			return pb.AlreadySitDown
		}
		if v, ok := t.seats[arg.Seat]; ok {
			glog.Errorf("sit faild %s %d", userid, seat)
			glog.Errorf("sit faild %s v %#v", userid, v)
			return pb.SitDownFailed
		}
	case pb.SitUp:
		if seat == 0 {
			return pb.StandUpFailed
		}
		arg.Seat = seat //站起位置
	}
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1): //私人
		if arg.Seat > t.DeskData.Count || arg.Seat <= 0 {
			return pb.OperateError
		}
		switch t.state {
		case int32(pb.STATE_READY):
		default:
			if v, ok := t.seats[seat]; ok && v.Ready {
				return pb.GameStartedCannotLeave
			}
		}
	case int32(pb.ROOM_TYPE2): //百人
		if !(arg.Seat >= 1 && arg.Seat <= 8) {
			return pb.OperateError
		}
		switch arg.Type {
		case pb.SitDown:
			if user.GetCoin() < int64(t.DeskData.Sit) {
				return pb.SitNotEnough
			}
			if userid == t.DeskGame.Dealer { //庄家不能坐
				return pb.DealerSitFailed
			}
		}
	}
	return pb.OK
}

//玩家站起
func (t *Desk) roleSitUp(userid string) {
	seat := t.getSeat(userid)
	if seat == 0 {
		return
	}
	arg := &pb.CLHSit{
		Type: pb.SitUp,
		Seat: seat,
	}
	rsp := t.freeSit(userid, arg)
	if rsp.Error == pb.OK {
		t.broadcast(rsp)
	} else {
		glog.Errorf("roleSitUp sit up err %s, %d", userid, seat)
	}
}

//.

//' 超时处理
func (t *Desk) freeTimeout() {
	switch t.state {
	case int32(pb.STATE_READY):
		t.freeStart()
	case int32(pb.STATE_BET):
		if t.timer == FreeBetTime {
			//出牌超时处理
			t.timer = 0
			t.freeGameOver()
		} else {
			t.timer++
		}
	case int32(pb.STATE_OVER):
		if t.timer == RestTime {
			//出牌超时处理
			t.timer = 0
			t.gameStartBet()
		} else {
			t.timer++
		}
	default:
		t.timer++
	}
}

//.

//'游戏状态

//结束重置
func (t *Desk) freeOverInit() {
	t.freeInit()
	t.state = int32(pb.STATE_OVER) //休息停顿
}

//开始状态初始化
func (t *Desk) freeStart() {
	if len(t.roles) > 0 {
		t.state = int32(pb.STATE_OVER) //休息停顿
	} else {
		t.state = int32(pb.STATE_READY) //准备
		return
	}
	t.freeStartMsg()
}

//状态变更消息
func (t *Desk) freeStartMsg() {
	var photo string
	var nickname string
	p := t.getPlayer(t.DeskGame.Dealer)
	if p != nil {
		photo = p.GetPhoto()
		nickname = p.GetNickname()
	}
	var left uint32 = t.leftDealerTimes()
	msg := resFreeStart(t.DeskGame.Dealer, photo, nickname,
		t.state, t.DeskFree.Carry, DealerTimes, left)
	t.broadcast(msg)
}

//开始下注
func (t *Desk) gameStartBet() {
	//下注状态
	t.state = int32(pb.STATE_BET)
	t.freeStartMsg()
}

//剩余坐庄次数
func (t *Desk) leftDealerTimes() uint32 {
	if DealerTimes >= t.DeskFree.DealerNum {
		return DealerTimes - t.DeskFree.DealerNum
	}
	glog.Errorf("Dealer %s, DealerNum %d", t.DeskGame.Dealer, t.DeskFree.DealerNum)
	return 0
}

//.

//'洗牌
func (t *Desk) shuffle() {
	rand.Seed(time.Now().UnixNano())
	d := make([]uint32, algo.NumCard, algo.NumCard)
	copy(d, algo.NiuCARDS)
	//测试暂时去掉洗牌
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	t.DeskGame.Cards = d
}

//.

//'发牌,直接发5张,百人场发固定位置
func (t *Desk) freeDeal() {
	var hand = 1
	for i := uint32(pb.DESK_SEAT1); i <= uint32(pb.DESK_SEAT2); i++ {
		cards := make([]uint32, hand, hand)
		tmp := t.DeskGame.Cards[:hand]
		copy(cards, tmp)
		t.DeskFree.Cards[i] = cards
		t.DeskGame.Cards = t.DeskGame.Cards[hand:]
		msg := resDraw(i, t.state, cards)
		t.broadcast(msg)
	}
}

//.

//'结束游戏
func (t *Desk) freeGameOver() {
	t.shuffle()  //洗牌
	t.freeDeal() //发牌
	// 结算
	cs1 := t.getHandCards(uint32(pb.DESK_SEAT1)) //庄家牌
	cs2 := t.getHandCards(uint32(pb.DESK_SEAT2)) //闲家牌
	//
	var muliti2, muliti3, muliti4 int64 //庄家赢返回正数,输返回负数
	var a1 uint32 = algo.Lhd(cs1, cs2) //1位置庄家牌力,2位置闲家牌力
	switch a1 {
	case 0://和
		muliti2, muliti3, muliti4 = 1, 1, -8
	case 1://庄
		muliti2, muliti3, muliti4 = -1, 1, 1
	case 2://闲
		muliti2, muliti3, muliti4 = 1, -1, 1
	}
	t.DeskFree.Power[uint32(pb.DESK_SEAT1)] = algo.LhdRank(cs1)
	t.DeskFree.Power[uint32(pb.DESK_SEAT2)] = algo.LhdRank(cs2)
	//各位置和庄家对比的赔付倍数
	t.DeskFree.Multiple[uint32(pb.DESK_SEAT2)] = muliti2
	t.DeskFree.Multiple[uint32(pb.DESK_SEAT3)] = muliti3
	t.DeskFree.Multiple[uint32(pb.DESK_SEAT4)] = muliti4
	//牌局数累加一次
	//t.DeskGame.Round++
	t.xianjiaJiesuan() //结算,闲家赔付
	//庄家收钱,TODO 奖池抽成
	t.dealerWin()
	//庄家赔付,闲家收钱,奖池抽成
	t.dealerJiesuan()
	//打印信息
	t.printOver()
	//结束消息
	msg := t.resOverFree()
	//广播
	t.broadcast(msg)
	//记录房间趋势
	t.saveTrend()
	//记录房间上局赢家
	t.saveWiners()
	//个人记录
	t.setFreeRecord()
	//重置状态
	t.freeOverInit()
	//检测不足做庄
	t.checkBeDealer()
	//踢除离线玩家
	t.kickOffline()
	//消息广播
	t.freeStart()
}

//.

//'记录房间趋势
func (t *Desk) saveTrend() {
	trend := new(data.FreeTrend)
	for k, v := range t.DeskFree.Multiple {
		switch k {
		case uint32(pb.DESK_SEAT2):
			trend.Seat2 = v < 0
		case uint32(pb.DESK_SEAT3):
			trend.Seat3 = v < 0
		case uint32(pb.DESK_SEAT4):
			trend.Seat4 = v < 0
		}
	}
	t.DeskFree.Trends = append(t.DeskFree.Trends, trend)
	//保留20条
	if len(t.DeskFree.Trends) >= 20 {
		t.DeskFree.Trends = t.DeskFree.Trends[1:]
	}
}

//.

//'记录房间上局赢家
func (t *Desk) saveWiners() {
	t.DeskFree.Winers = make([]*data.FreeWiner, 0)
	for k, v := range t.DeskFree.Score2 {
		if v < 0 {
			continue
		}
		if (v - t.DeskFree.Bets[k]) < 0 {
			continue
		}
		user := t.getPlayer(k)
		if user == nil {
			glog.Errorf("userid %s not exist", k)
			continue
		}
		winer := &data.FreeWiner{
			Userid:   k,
			Nickname: user.GetNickname(),
			Photo:    user.GetPhoto(),
			Coin:     v - t.DeskFree.Bets[k],
		}
		t.DeskFree.Winers = append(t.DeskFree.Winers, winer)
	}
}

//.

//'个人记录
func (t *Desk) setFreeRecord() {
	for k, v := range t.DeskFree.Score2 {
		pid := t.getPid(k)
		if pid == nil {
			continue
		}
		user := t.getPlayer(k)
		if user == nil {
			glog.Errorf("userid %s not exist", k)
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

//'庄家收钱,TODO 奖池抽成
func (t *Desk) dealerWin() {
	//庄家收钱总额
	var val int64
	for seat, m := range t.DeskFree.Score3 {
		for _, v := range m {
			t.DeskFree.Score1[seat] += v
			val += v
		}
	}
	//庄家收钱转为正数
	if val < 0 {
		val *= -1
	}
	if val > 0 {
		//	抽成
		val = t.drawcoin(t.DeskGame.Dealer, val)
	}
	if val != 0 {
		t.DeskFree.Score1[uint32(pb.DESK_SEAT1)] = val
	}
	t.DeskFree.Carry += val //更新庄家收入携带
}

//.

//'闲家赔付
func (t *Desk) xianjiaJiesuan() {
	for seat, v := range t.DeskFree.Multiple {
		if v < 0 {      //表示庄家输
			continue
		}
		tmp := t.getFreeBets(seat)
		t.DeskFree.Score3[seat] = make(map[string]int64)
		//赔付倍数
		var val int64
		if v > 1 {
			val = v - 1
		} else if v < -1 {
			val = (v * -1) - 1
		}
		if val != 0 { //表示庄家赢,且大于1倍从玩家身上扣赔付倍数
			for userid, betNum := range tmp {
				p := t.getPlayer(userid)
				if p == nil {
					glog.Errorf("userid %s not exist", userid)
					continue
				}
				coin := p.GetCoin()
				num := val * betNum
				if num > coin {
					num = coin
				}
				//扣除位置数
				t.sendCoin(userid, (-1 * int64(num)), int32(pb.LOG_TYPE6))
				t.DeskFree.Score3[seat][userid] = -1 * int64((betNum + num))
				t.DeskFree.Score2[userid] += -1 * int64((betNum + num))
			}
		} else {
			for userid, betNum := range tmp {
				t.DeskFree.Score3[seat][userid] = -1 * int64(betNum)
				t.DeskFree.Score2[userid] += -1 * int64(betNum)
			}
		}
	}
}

//.

//'庄家赔付
func (t *Desk) dealerJiesuan() {
	var num int64 //庄家赔付金额
	for seat, v := range t.DeskFree.Multiple {
		if v > 0 {      //表示庄家赢
			continue
		}
		num += t.DeskFree.SeatBets[seat] * v * -1
	}
	if t.DeskFree.Carry >= num { //足够赔付
		t.dealerJiesuan1()
	} else { //不足赔付
		t.dealerJiesuan2(num)
	}
}

//.

//'足够赔付
func (t *Desk) dealerJiesuan1() {
	for seat, v := range t.DeskFree.Multiple {
		if v > 0 {      //表示庄家赢
			continue
		}
		tmp := t.getFreeBets(seat)
		t.DeskFree.Score3[seat] = make(map[string]int64)
		var val int64 = v
		if v < 0 { //表示庄家输
			val = v * -1
		}
		for userid, betNum := range tmp {
			num := val * betNum
			if num > t.DeskFree.Carry {
				num = t.DeskFree.Carry
			}
			t.DeskFree.Carry -= num
			//	抽成, 赢利中抽取
			num2 := t.drawcoin(userid, num)
			val2 := int64(num2 + betNum)
			if val2 < 0 {
				val2 = 0
			}
			//扣除位置数
			t.sendCoin(userid, val2, int32(pb.LOG_TYPE6))
			t.DeskFree.Score3[seat][userid] = val2
			t.DeskFree.Score1[seat] += val2
			t.DeskFree.Score2[userid] += val2
			if num != 0 {
				t.DeskFree.Score1[uint32(pb.DESK_SEAT1)] -= int64(num)
			}
		}
	}
}

//.

//'不足赔付
func (t *Desk) dealerJiesuan2(num int64) {
	m := make(map[uint32]int64)
	for seat, v := range t.DeskFree.Multiple {
		if v > 0 { //表示庄家赢
			continue
		}
		//当前位置的总金额
		num1 := t.DeskFree.SeatBets[seat] * v * -1
		num2 := (num1 / num) * t.DeskFree.Carry
		m[seat] = num2 //位置分到金额
	}
	for seat, val := range m {
		tmp := t.getFreeBets(seat)
		betsNum := t.DeskFree.SeatBets[seat]
		t.DeskFree.Score3[seat] = make(map[string]int64)
		for userid, betNum := range tmp {
			num2 := (betNum / betsNum) * val //分到金额
			//	抽成, 赢利中抽取
			num3 := t.drawcoin(userid, num2)
			val2 := num3 + betNum //加上下注额
			if val2 < 0 {
				val2 = 0
			}
			//扣除位置数
			t.sendCoin(userid, val2, int32(pb.LOG_TYPE6))
			t.DeskFree.Score3[seat][userid] = val2
			t.DeskFree.Score1[seat] += val2
			t.DeskFree.Score2[userid] += val2
			if num2 != 0 {
				t.DeskFree.Score1[uint32(pb.DESK_SEAT1)] -= int64(num2)
			}
		}
	}
}

//.

//'获取对应位置下注列表
func (t *Desk) getFreeBets(seat uint32) map[string]int64 {
	return t.SeatRoleBets[seat]
}

//.

//' 发牌消息
func resDraw(seat uint32, state int32, cards []uint32) *pb.SLHDraw {
	return &pb.SLHDraw{
		Seat:  seat,
		State: state,
		Cards: cards,
	}
}

//.

//' 游戏开始消息
func resFreeStart(dealer, photo, nickname string, state int32,
	carry int64, dealerNum, left uint32) *pb.SLHFreeGamestart {
	return &pb.SLHFreeGamestart{
		Dealer:        dealer,
		Photo:         photo,
		State:         state,
		Coin:          carry,
		DealerNum:     dealerNum,
		LeftDealerNum: left,
		Nickname:      nickname,
	}
}

//.

//' 游戏结束消息
func (t *Desk) resOverFree() *pb.SLHFreeGameover {
	var left uint32
	if t.DeskGame.Dealer == "" {
		left = 1
	} else {
		left = t.leftDealerTimes()
	}
	msg := &pb.SLHFreeGameover{
		State:         t.state,
		Dealer:        t.DeskGame.Dealer,
		Coin:          t.DeskFree.Carry,
		DealerNum:     DealerTimes,
		LeftDealerNum: left,
	}
	for k, v := range t.DeskFree.Power {
		d := &pb.LHFreeRoomOver{
			Seat:  k,
			Value: v,
			Cards: t.getHandCards(k),
			Multi: t.DeskFree.Multiple[k],
		}
		msg.Data = append(msg.Data, d)
	}
	for k, v := range t.DeskFree.Score1 {
		d := &pb.LHFreeSeatOver{
			Seat:  k,
			Score: v,
			Total: t.DeskFree.SeatBets[k],
		}
		for k1, v1 := range t.DeskFree.Score3[k] {
			list := &pb.LHRoomScore{
				Seat:   t.getSeat(k1),
				Userid: k1,
				Score:  v1,
			}
			d.List = append(d.List, list)
		}
		msg.Info = append(msg.Info, d)
	}
	for k, v := range t.DeskFree.Score2 {
		p := t.getPlayer(k)
		if p == nil {
			glog.Errorf("userid %s not exist", k)
			continue
		}
		list := &pb.LHRoomScore{
			Seat:   t.getSeat(k),
			Userid: k,
			Score:  v,
			Coin:   p.GetCoin(),
		}
		msg.List = append(msg.List, list)
	}
	return msg
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
