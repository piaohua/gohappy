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
func (t *Desk) freeEnterMsg(userid string) *pb.SSGFreeEnterRoom {
	msg := new(pb.SSGFreeEnterRoom)
	//房间数据
	msg.Roominfo = handler.PackSGFreeRoom(t.DeskData)
	t.freeRoomDataMsg(msg.Roominfo)
	//坐下玩家信息
	msg.Userinfo = t.freeSeatBetsMsg()
	//位置下注信息
	msg.Betsinfo = handler.PackSGRoomBets(t.DeskFree.SeatBets)
	return msg
}

func (t *Desk) freeRoomDataMsg(msg *pb.SGFreeRoom) {
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
	msg := new(pb.SSGFreeCamein)
	msg.Userinfo = t.freeSeatRoleMsg(userid)
	t.broadcast(msg)
}

//位置上玩家数据
func (t *Desk) freeSeatRoleMsg(userid string) (msg *pb.SGFreeUser) {
	if v, ok := t.roles[userid]; ok {
		if v.Seat == 0 {
			return //没有坐下不广播
		}
		msg = handler.PackSGFreeUser(v.User)
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
func (t *Desk) freeSeatBetsMsg() (msg []*pb.SGFreeUser) {
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
func (t *Desk) freeBetsMsg(userid string) (msg []*pb.SGRoomBets) {
	for i := pb.DESK_SEAT2; i <= pb.DESK_SEAT9; i++ {
		seat := uint32(i)
		bets := t.getFreeSeatBet(userid, seat)
		msg2 := &pb.SGRoomBets{
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
	//TODO 暂时全部带上庄
	num = uint32(user.GetCoin())
	switch st {
	case int32(pb.DEALER_DOWN):
		switch t.state {
		case int32(pb.STATE_BET): //下注中
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
func (t *Desk) dealerListMsg() (msg *pb.SSGFreeDealerList) {
	msg = new(pb.SSGFreeDealerList)
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
		msg2 := &pb.SGDealerList{
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
	//下注不能大于庄家携带1/4
	if t.DeskGame.Dealer != "" &&
		((t.DeskGame.BetNum + num) > (t.DeskFree.Carry / 4)) {
		return pb.BetTopLimit //下注限制
	}
	coin := user.GetCoin()          //剩余金额
	bets := t.DeskFree.Bets[userid] //已经下注额
	//本轮下注不能超过1/4
	if (num + bets) > ((coin + bets) / 4) {
		return pb.BetTopLimit //下注限制
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
	bets int64, userid string) *pb.SSGFreeBet {
	return &pb.SSGFreeBet{
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
func (t *Desk) freeTrends() *pb.SSGFreeTrend {
	msg := new(pb.SSGFreeTrend)
	for _, v := range t.DeskFree.Trends {
		msg2 := &pb.SGFreeTrend{
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
func (t *Desk) freeWiners() *pb.SSGFreeWiners {
	msg := new(pb.SSGFreeWiners)
	for _, v := range t.DeskFree.Winers {
		msg2 := &pb.SGFreeWiner{
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
func (t *Desk) freeRoles() *pb.SSGFreeRoles {
	msg := new(pb.SSGFreeRoles)
	for _, v := range t.roles {
		if v.User == nil {
			continue
		}
		msg2 := &pb.SGFreeRole{
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
func (t *Desk) freeSit(userid string, arg *pb.CSGSit) (msg *pb.SSGSit) {
	errcode := t.sitCheck(userid, arg)
	if errcode != pb.OK {
		msg = new(pb.SSGSit)
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
	msg = &pb.SSGSit{
		Type:     arg.Type,
		Seat:     arg.Seat,
		Userinfo: t.coinRoleMsg(userid),
	}
	return
}

func (t *Desk) sitCheck(userid string, arg *pb.CSGSit) pb.ErrCode {
	user := t.getPlayer(userid)
	if user == nil {
		glog.Errorf("userid %s not exist", userid)
		return pb.NotInRoom
	}
	seat := t.getSeat(userid)
	switch arg.Type {
	case pb.SitDown:
		if seat != 0 {
			return pb.SitDownFailed
		}
		if _, ok := t.seats[arg.Seat]; ok {
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
				return pb.NotEnoughCoin
			}
			if userid == t.DeskGame.Dealer { //庄家不能坐
				return pb.DealerSitFailed
			}
		}
	}
	return pb.OK
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
	t.freeStart()
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
	p := t.getPlayer(t.DeskGame.Dealer)
	if p != nil {
		photo = p.GetPhoto()
	}
	var left uint32 = t.leftDealerTimes()
	msg := resFreeStart(t.DeskGame.Dealer, photo, t.state,
		t.DeskFree.Carry, DealerTimes, left)
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
	var hand = 3
	for i := uint32(pb.DESK_SEAT1); i <= uint32(pb.DESK_SEAT5); i++ {
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
	cs3 := t.getHandCards(uint32(pb.DESK_SEAT3)) //闲家牌
	cs4 := t.getHandCards(uint32(pb.DESK_SEAT4)) //闲家牌
	cs5 := t.getHandCards(uint32(pb.DESK_SEAT5)) //闲家牌
	//
	var a1 uint32 = algo.San(cs1) //1位置庄家牌力
	var a2 uint32 = algo.San(cs2) //2位置闲家牌力
	var a3 uint32 = algo.San(cs3) //3位置闲家牌力
	var a4 uint32 = algo.San(cs4) //4位置闲家牌力
	var a5 uint32 = algo.San(cs5) //5位置闲家牌力
	t.DeskFree.Power[uint32(pb.DESK_SEAT1)] = a1
	t.DeskFree.Power[uint32(pb.DESK_SEAT2)] = a2
	t.DeskFree.Power[uint32(pb.DESK_SEAT3)] = a3
	t.DeskFree.Power[uint32(pb.DESK_SEAT4)] = a4
	t.DeskFree.Power[uint32(pb.DESK_SEAT5)] = a5
	//各位置和庄家对比的赔付倍数
	t.DeskFree.Multiple[uint32(pb.DESK_SEAT2)] = muliti(a1, a2, cs1, cs2)
	t.DeskFree.Multiple[uint32(pb.DESK_SEAT3)] = muliti(a1, a3, cs1, cs3)
	t.DeskFree.Multiple[uint32(pb.DESK_SEAT4)] = muliti(a1, a4, cs1, cs4)
	t.DeskFree.Multiple[uint32(pb.DESK_SEAT5)] = muliti(a1, a5, cs1, cs5)
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
	//task
	t.taskHandler()
	//重置状态
	t.freeOverInit()
	//检测不足做庄
	t.checkBeDealer()
	//踢除离线玩家
	t.kickOffline()
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
		case uint32(pb.DESK_SEAT5):
			trend.Seat5 = v < 0
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
			msg2 := handler.TaskUpdateMsg(1, pb.TASK_TYPE12, k)
			pid.Tell(msg2)
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
		if v >= 200000 {
			msg2 := handler.TaskUpdateMsg(1, pb.TASK_TYPE14, k)
			pid.Tell(msg2)
			msg3 := handler.TaskUpdateMsg(1, pb.TASK_TYPE16, k)
			pid.Tell(msg3)
		} else if v >= 100000 {
			msg2 := handler.TaskUpdateMsg(1, pb.TASK_TYPE13, k)
			pid.Tell(msg2)
			msg3 := handler.TaskUpdateMsg(1, pb.TASK_TYPE15, k)
			pid.Tell(msg3)
		}
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
	for k, v := range t.DeskFree.Multiple {
		var seat uint32
		if v < 0 { //表示庄家输
			//压负位置赔付
			seat = t.getFreeSeat(false, k)
		} else {
			//压正位置赔付
			seat = t.getFreeSeat(true, k)
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
	for k, v := range t.DeskFree.Multiple {
		var seat uint32 //赢家位置
		if v > 0 {      //表示庄家赢
			//赔付压负位置
			seat = t.getFreeSeat(false, k)
			num += t.DeskFree.SeatBets[seat] * v
		} else {
			//赔付压正位置
			seat = t.getFreeSeat(true, k)
			num += t.DeskFree.SeatBets[seat] * v * -1
		}
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
	for k, v := range t.DeskFree.Multiple {
		var seat uint32 //赢家位置
		if v > 0 {      //表示庄家赢
			//赔付压负位置
			seat = t.getFreeSeat(false, k)
		} else {
			//赔付压正位置
			seat = t.getFreeSeat(true, k)
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
			n := int64(num + betNum)
			val2 := n
			if val2 < 0 {
				val2 = 0
			} else {
				//	抽成
				val2 = t.drawcoin(userid, val2)
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
	for k, v := range t.DeskFree.Multiple {
		var seat uint32 //赢家位置
		if v > 0 {      //表示庄家赢
			//赔付压负位置
			seat = t.getFreeSeat(false, k)
			//当前位置的总金额
			num1 := t.DeskFree.SeatBets[seat] * v
			num2 := (num1 / num) * t.DeskFree.Carry
			m[seat] = num2 //位置分到金额
		} else {
			//赔付压正位置
			seat = t.getFreeSeat(true, k)
			//当前位置的总金额
			num1 := t.DeskFree.SeatBets[seat] * v * -1
			num2 := (num1 / num) * t.DeskFree.Carry
			m[seat] = num2 //位置分到金额
		}
	}
	for seat, val := range m {
		tmp := t.getFreeBets(seat)
		betsNum := t.DeskFree.SeatBets[seat]
		t.DeskFree.Score3[seat] = make(map[string]int64)
		for userid, betNum := range tmp {
			num2 := (betNum / betsNum) * val //分到金额
			num3 := num2 + betNum            //加上下注额
			val2 := int64(num3)
			if val2 < 0 {
				val2 = 0
			} else {
				//	抽成
				val2 = t.drawcoin(userid, val2)
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

//'获取相对输赢位置,win=压正负位置
func (t *Desk) getFreeSeat(win bool, k uint32) (seat uint32) {
	if win {
		//压正
		switch k {
		case uint32(pb.DESK_SEAT2):
			seat = uint32(pb.DESK_SEAT2)
		case uint32(pb.DESK_SEAT3):
			seat = uint32(pb.DESK_SEAT3)
		case uint32(pb.DESK_SEAT4):
			seat = uint32(pb.DESK_SEAT4)
		case uint32(pb.DESK_SEAT5):
			seat = uint32(pb.DESK_SEAT5)
		}
		return
	}
	//压负
	switch k {
	case uint32(pb.DESK_SEAT2):
		seat = uint32(pb.DESK_SEAT6)
	case uint32(pb.DESK_SEAT3):
		seat = uint32(pb.DESK_SEAT7)
	case uint32(pb.DESK_SEAT4):
		seat = uint32(pb.DESK_SEAT8)
	case uint32(pb.DESK_SEAT5):
		seat = uint32(pb.DESK_SEAT9)
	}
	return
}

//.

//'获取对应位置下注列表
func (t *Desk) getFreeBets(seat uint32) map[string]int64 {
	return t.SeatRoleBets[seat]
}

//.

//'返回庄家赢倍数,a1庄家牌力,an闲家牌力,庄家赢返回正数,输返回负数
func muliti(a1, an uint32, cs1, csn []uint32) int64 {
	switch {
	case a1 > an:
		return int64(algo.SanMultiple(a1))
	case a1 < an:
		return -1 * int64(algo.SanMultiple(an))
	case a1 == an:
		if algo.Compare(cs1, csn) {
			return int64(algo.SanMultiple(a1))
		}
		return -1 * int64(algo.SanMultiple(an))
	}
	return 1
}

//.

//' 发牌消息
func resDraw(seat uint32, state int32, cards []uint32) *pb.SSGDraw {
	return &pb.SSGDraw{
		Seat:  seat,
		State: state,
		Cards: cards,
	}
}

//.

//' 游戏开始消息
func resFreeStart(dealer, photo string, state int32, carry int64,
	dealerNum, left uint32) *pb.SSGFreeGamestart {
	return &pb.SSGFreeGamestart{
		Dealer:        dealer,
		Photo:         photo,
		State:         state,
		Coin:          carry,
		DealerNum:     dealerNum,
		LeftDealerNum: left,
	}
}

//.

//' 游戏结束消息
func (t *Desk) resOverFree() *pb.SSGFreeGameover {
	var left uint32
	if t.DeskGame.Dealer == "" {
		left = 1
	} else {
		left = t.leftDealerTimes()
	}
	msg := &pb.SSGFreeGameover{
		State:         t.state,
		Dealer:        t.DeskGame.Dealer,
		Coin:          t.DeskFree.Carry,
		DealerNum:     DealerTimes,
		LeftDealerNum: left,
	}
	for k, v := range t.DeskFree.Power {
		d := &pb.SGFreeRoomOver{
			Seat:  k,
			Value: v,
			Cards: t.getHandCards(k),
			Multi: t.DeskFree.Multiple[k],
		}
		msg.Data = append(msg.Data, d)
	}
	for k, v := range t.DeskFree.Score1 {
		d := &pb.SGFreeSeatOver{
			Seat:  k,
			Score: v,
			Total: t.DeskFree.SeatBets[k],
		}
		for k1, v1 := range t.DeskFree.Score3[k] {
			list := &pb.SGRoomScore{
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
		list := &pb.SGRoomScore{
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

//' 任务处理
func (t *Desk) taskHandler() {
	for k, v := range t.DeskFree.Power {
		var taskType pb.TaskType
		switch v {
		case algo.Gong1:
			taskType = pb.TASK_TYPE18
		case algo.Gong2:
			taskType = pb.TASK_TYPE19
		case algo.Gong3:
			taskType = pb.TASK_TYPE20
		}
		if int32(taskType) == 0 {
			continue
		}
		seat1 := t.getFreeSeat(false, k)
		seat2 := t.getFreeSeat(true, k)
		for k1 := range t.DeskFree.SeatRoleBets[seat1] {
			msg2 := handler.TaskUpdateMsg(1, taskType, k1)
			t.send3userid(k1, msg2)
		}
		for k2 := range t.DeskFree.SeatRoleBets[seat2] {
			msg2 := handler.TaskUpdateMsg(1, taskType, k2)
			t.send3userid(k2, msg2)
		}
	}
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
