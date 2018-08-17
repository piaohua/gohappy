/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-01-22 17:06:12
 * Filename      : recv.go
 * Description   : 机器人
 * *******************************************************/
package main

import (
	"gohappy/glog"
	"gohappy/pb"
)

//' receive 接收消息
func (r *Robot) receive(msg interface{}) {
	switch message := msg.(type) {
	//game
	case *pb.SRegist:
		r.recvRegist(message)
	case *pb.SLogin:
		r.recvLogin(message)
	case *pb.SUserData:
		r.recvdata(message)
	case *pb.SPushCurrency:
		r.recvPushCurrency(message)
	case *pb.SPing:
		r.recvPing(message)
	//niu
	case *pb.SNNFreeEnterRoom:
		r.recvNNFreeEnter(message)
	case *pb.SNNFreeCamein:
		r.recvNNFreeCamein(message)
	case *pb.SNNFreeGamestart:
		r.recvNNFreeStart(message)
	case *pb.SNNFreeGameover:
		r.recvNNFreeGameover(message)
	case *pb.SNNFreeBet:
		r.recvNNFreeBet(message)
	case *pb.SNNCoinEnterRoom:
		r.recvNNCoinEnter(message)
	case *pb.SNNCoinGameover:
		r.recvNNCoinGameover(message)
	case *pb.SNNGameover:
		r.recvNNGameover(message)
	case *pb.SNNEnterRoom:
		r.recvNNEnter(message)
	case *pb.SNNLeave:
		r.recvNNLeave(message)
	case *pb.SNNCamein:
		r.recvNNCamein(message)
	case *pb.SNNPushState:
		r.recvNNPushstate(message)
	case *pb.SNNDraw:
		r.recvNNDraw(message)
	case *pb.SNNDealer:
		r.recvNNDealer(message)
	case *pb.SNNPushDealer:
		r.recvNNPushDealer(message)
	case *pb.SNNReady:
		r.recvNNReady(message)
	case *pb.SNNPushDrawCoin:
	case *pb.SNNBet:
	case *pb.SNNiu:
	//ebg
	case *pb.SEBFreeEnterRoom:
		r.recvEBFreeEnter(message)
	case *pb.SEBFreeCamein:
		r.recvEBFreeCamein(message)
	case *pb.SEBFreeGamestart:
		r.recvEBFreeStart(message)
	case *pb.SEBFreeGameover:
		r.recvEBFreeGameover(message)
	case *pb.SEBFreeBet:
		r.recvEBFreeBet(message)
	case *pb.SEBCoinEnterRoom:
		r.recvEBCoinEnter(message)
	case *pb.SEBCamein:
		r.recvEBCamein(message)
	case *pb.SEBCoinGameover:
		r.recvEBCoinGameover(message)
	case *pb.SEBGameover:
		r.recvEBGameover(message)
	case *pb.SEBPushState:
		r.recvEBPushstate(message)
	case *pb.SEBDraw:
		r.recvEBDraw(message)
	case *pb.SEBEnterRoom:
		r.recvEBEnter(message)
	case *pb.SEBLeave:
		r.recvEBLeave(message)
	case *pb.SEBReady:
		r.recvEBReady(message)
	case *pb.SEBDealer:
		r.recvEBDealer(message)
	case *pb.SEBPushDealer:
		r.recvEBPushDealer(message)
	case *pb.SEBPushDrawCoin:
	case *pb.SEBBet:
	case *pb.SEBiu:
	default:
		glog.Errorf("unknow message: %#v", message)
	}
}

//.

//' 接收到服务器注册返回
func (r *Robot) recvRegist(s2c *pb.SRegist) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
		Logined(r.data.Phone, r.ltype) //登录成功
		r.regist = true                //注册成功
		r.data.Userid = s2c.GetUserid()
		glog.Infof("regist successful -> %s", r.data.Userid)
		r.sendUserData() // 获取玩家数据
		return
	case pb.PhoneRegisted:
		glog.Infof("phone registed -> %s", r.data.Phone)
		r.sendLogin() //尝试登录
		return
	default:
		glog.Infof("regist err -> %d", errcode)
	}
	//重新尝试登录
	//go ReLogined(r.roomid, r.data.Phone, r.code, r.rtype, r.envBet)
	r.Close()
}

//.

//' 接收到服务器登录返回
func (r *Robot) recvLogin(s2c *pb.SLogin) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
		Logined(r.data.Phone, r.ltype) //登录成功
		r.data.Userid = s2c.GetUserid()
		glog.Infof("login successful -> %s", r.data.Userid)
		r.sendUserData() // 获取玩家数据
		return
	default:
		glog.Infof("login err -> %d", errcode)
	}
	r.Close()
}

//.

//' 接收到玩家数据
func (r *Robot) recvdata(s2c *pb.SUserData) {
	var errcode = s2c.GetError()
	if errcode != pb.OK {
		glog.Infof("get data err -> %d", errcode)
		r.Close() //断开
		return
	}
	userdata := s2c.GetData()
	// 设置数据
	r.data.Userid = userdata.GetUserid()     // 用户id
	r.data.Nickname = userdata.GetNickname() // 用户昵称
	r.data.Sex = userdata.GetSex()           //
	r.data.Coin = userdata.GetCoin()         // 金币
	r.data.Diamond = userdata.GetDiamond()   // 钻石
	r.data.Chip = userdata.GetChip()
	r.data.Card = userdata.GetCard()
	//chip 单位为分
	if r.data.Coin < 650000 {
		r.addCurrency() //充值
	}
	//进入房间
	switch r.gtype {
	case int32(pb.NIU):
		r.sendNNEntryRoom()
	case int32(pb.EBG):
		r.sendEBEntryRoom()
	default:
		r.Close()
	}
}

//更新金币
func (r *Robot) recvPushCurrency(s2c *pb.SPushCurrency) {
	currencyData := s2c.GetData()
	r.data.Coin += currencyData.GetCoin()
	r.data.Card += currencyData.GetCard()
	r.data.Chip += currencyData.GetChip()
	r.data.Diamond += currencyData.GetDiamond()
	if r.data.Coin < 650000 {
		//r.addCurrency()	//充值
		//r.sendNNStandup() //离开
	}
}

//游戏
func (r *Robot) recvPing(s2c *pb.SPing) {
	//glog.Debugf("ping %s", r.data.Userid)
}

//.

//' 离开房间
func (r *Robot) recvNNLeave(s2c *pb.SNNLeave) {
	if s2c.GetUserid() == r.data.Userid {
		r.Close() //下线
	}
}

func (r *Robot) recvEBLeave(s2c *pb.SEBLeave) {
	if s2c.GetUserid() == r.data.Userid {
		r.Close() //下线
	}
}

//.

//' 进入房间
func (r *Robot) recvNNFreeEnter(s2c *pb.SNNFreeEnterRoom) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
	default:
		glog.Infof("comein err -> %d", errcode)
		r.Close() //进入出错,关闭
		return
	}
	roominfo := s2c.GetRoominfo()
	r.gtype = roominfo.Gtype
	r.rtype = roominfo.Rtype
	r.dtype = roominfo.Dtype
	r.roomid = roominfo.Roomid
	userinfo := s2c.GetUserinfo()
	for _, v := range userinfo {
		//只返回坐下玩家
		if v.Userid == r.data.Userid {
			glog.Debugf("comein user info -> %s", v.Userid)
			r.seat = v.Seat
			break
		}
	}
	r.gameStart(s2c.Roominfo.State)
}

func (r *Robot) recvNNCoinEnter(s2c *pb.SNNCoinEnterRoom) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
	default:
		glog.Errorf("comein err -> %d", errcode)
		r.Close() //进入出错,关闭
		return
	}
	roominfo := s2c.GetRoominfo()
	r.gtype = roominfo.Gtype
	r.rtype = roominfo.Rtype
	r.dtype = roominfo.Dtype
	r.roomid = roominfo.Roomid
	userinfo := s2c.GetUserinfo()
	for _, v := range userinfo {
		//只返回坐下玩家
		if v.Userid == r.data.Userid {
			glog.Debugf("comein user info -> %s", v.Userid)
			r.seat = v.Seat
			break
		}
	}
	r.gameStart(s2c.Roominfo.State)
}

func (r *Robot) recvNNEnter(s2c *pb.SNNEnterRoom) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
	default:
		glog.Errorf("comein err -> %d", errcode)
		r.Close() //进入出错,关闭
		return
	}
	roominfo := s2c.GetRoominfo()
	r.gtype = roominfo.Gtype
	r.rtype = roominfo.Rtype
	r.dtype = roominfo.Dtype
	r.roomid = roominfo.Roomid
	userinfo := s2c.GetUserinfo()
	for _, v := range userinfo {
		//只返回坐下玩家
		if v.Userid == r.data.Userid {
			glog.Debugf("comein user info -> %s", v.Userid)
			r.seat = v.Seat
			break
		}
	}
	r.gameStart(s2c.Roominfo.State)
}

//进入房间
func (r *Robot) recvNNCamein(s2c *pb.SNNCamein) {
	if s2c.GetUserinfo().GetUserid() == r.data.Userid {
	}
}

func (r *Robot) recvNNFreeCamein(s2c *pb.SNNFreeCamein) {
	if s2c.GetUserinfo().GetUserid() == r.data.Userid {
	}
}

//.

//' 百人

//下注
func (r *Robot) recvNNFreeBet(s2c *pb.SNNFreeBet) {
	var errcode = s2c.GetError()
	var userid string = s2c.GetUserid()
	glog.Debugf("bet userid %s, errcode %d", userid, errcode)
	switch errcode {
	case pb.OK:
	default:
		r.sendNNStandup() //离开
		return
	}
	if userid == r.data.Userid {
		val := s2c.GetValue()
		if val > r.bitNum {
			r.bitNum = 0
		} else {
			r.bitNum -= val
		}
		r.bits--
	}
	if r.bits > 0 && r.bitNum > 0 {
		r.sendNNFreeBet() //下注
	}
}

//状态更新
func (r *Robot) recvNNPushstate(s2c *pb.SNNPushState) {
	var state int32 = s2c.GetState()
	r.gameStart(state)
}

//结束
func (r *Robot) recvNNFreeGameover(s2c *pb.SNNFreeGameover) {
	r.bits = 0
	r.bitNum = 0
	r.round++
	if r.round >= 30 { //打10局下线
		r.sendNNStandup() //离开
		return
	}
}

//结束
func (r *Robot) recvNNCoinGameover(s2c *pb.SNNCoinGameover) {
	r.bits = 0
	r.bitNum = 0
	r.round++
	if r.round >= 30 { //打10局下线
		r.sendNNStandup() //离开
		return
	}
}

//结束
func (r *Robot) recvNNGameover(s2c *pb.SNNGameover) {
	if s2c.LeftRound == 0 {
		r.sendNNStandup() //离开
	}
	r.ready = false
}

func (r *Robot) gameStart(state int32) {
	switch r.gtype {
	case int32(pb.EBG):
		r.ebgGameStart(state)
	case int32(pb.NIU):
		r.niuGameStart(state)
	default:
	}
}

func (r *Robot) niuGameStart(state int32) {
	switch state {
	case int32(pb.STATE_READY):
		r.sendNNReady()
	case int32(pb.STATE_DEALER):
		r.sendNNDealer()
	case int32(pb.STATE_NIU):
		r.sendNNiu()
	case int32(pb.STATE_BET):
		r.sendNNBet() //下注
	case int32(pb.STATE_OVER):
	default:
		r.sendNNStandup() //离开
	}
}

//开始
func (r *Robot) recvNNFreeStart(s2c *pb.SNNFreeGamestart) {
	r.gameStart(s2c.GetState())
}

//准备
func (r *Robot) recvNNReady(s2c *pb.SNNReady) {
	if s2c.GetSeat() == r.seat {
		r.ready = s2c.GetReady()
	} else 	if !r.ready {
		r.sendNNReady()
	}
}

//发牌
func (r *Robot) recvNNDraw(s2c *pb.SNNDraw) {
	//TODO 计算牌力大小
	if s2c.GetSeat() == r.seat {
		glog.Debugf("draw cards %#v", s2c)
	}
}

//抢庄结果
func (r *Robot) recvNNDealer(s2c *pb.SNNDealer) {
}

//庄家
func (r *Robot) recvNNPushDealer(s2c *pb.SNNPushDealer) {
}

//.

//' ebg

//下注
func (r *Robot) recvEBFreeBet(s2c *pb.SEBFreeBet) {
	var errcode = s2c.GetError()
	var userid string = s2c.GetUserid()
	glog.Debugf("bet userid %s, errcode %d", userid, errcode)
	switch errcode {
	case pb.OK:
	default:
		r.sendEBStandup() //离开
		return
	}
	if userid == r.data.Userid {
		val := s2c.GetValue()
		if val > r.bitNum {
			r.bitNum = 0
		} else {
			r.bitNum -= val
		}
		r.bits--
	}
	if r.bits > 0 && r.bitNum > 0 {
		r.sendEBFreeBet() //下注
	}
}

//状态更新
func (r *Robot) recvEBPushstate(s2c *pb.SEBPushState) {
	var state int32 = s2c.GetState()
	r.gameStart(state)
}

//结束
func (r *Robot) recvEBFreeGameover(s2c *pb.SEBFreeGameover) {
	r.bits = 0
	r.bitNum = 0
	r.round++
	if r.round >= 30 { //打10局下线
		r.sendEBStandup() //离开
		return
	}
}

//结束
func (r *Robot) recvEBCoinGameover(s2c *pb.SEBCoinGameover) {
	r.bits = 0
	r.bitNum = 0
	r.round++
	if r.round >= 30 { //打10局下线
		r.sendEBStandup() //离开
		return
	}
}

//结束
func (r *Robot) recvEBGameover(s2c *pb.SEBGameover) {
	if s2c.LeftRound == 0 {
		r.sendEBStandup() //离开
	}
	r.ready = false
}

func (r *Robot) ebgGameStart(state int32) {
	switch state {
	case int32(pb.STATE_READY):
		r.sendEBReady()
	case int32(pb.STATE_DEALER):
		r.sendEBDealer()
	case int32(pb.STATE_NIU):
		r.sendEBiu()
	case int32(pb.STATE_BET):
		r.sendEBBet() //下注
	case int32(pb.STATE_OVER):
	default:
		r.sendEBStandup() //离开
	}
}

//开始
func (r *Robot) recvEBFreeStart(s2c *pb.SEBFreeGamestart) {
	r.gameStart(s2c.GetState())
}

//准备
func (r *Robot) recvEBReady(s2c *pb.SEBReady) {
	if s2c.GetSeat() == r.seat {
		r.ready = s2c.GetReady()
	} else 	if !r.ready {
		r.sendEBReady()
	}
}

//发牌
func (r *Robot) recvEBDraw(s2c *pb.SEBDraw) {
	//TODO 计算牌力大小
	if s2c.GetSeat() == r.seat {
		glog.Debugf("draw cards %#v", s2c)
	}
}

//抢庄结果
func (r *Robot) recvEBDealer(s2c *pb.SEBDealer) {
}

//庄家
func (r *Robot) recvEBPushDealer(s2c *pb.SEBPushDealer) {
}
//.


//' 进入房间
func (r *Robot) recvEBFreeEnter(s2c *pb.SEBFreeEnterRoom) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
	default:
		glog.Infof("comein err -> %d", errcode)
		r.Close() //进入出错,关闭
		return
	}
	roominfo := s2c.GetRoominfo()
	r.gtype = roominfo.Gtype
	r.rtype = roominfo.Rtype
	r.dtype = roominfo.Dtype
	r.roomid = roominfo.Roomid
	userinfo := s2c.GetUserinfo()
	for _, v := range userinfo {
		//只返回坐下玩家
		if v.Userid == r.data.Userid {
			glog.Debugf("comein user info -> %s", v.Userid)
			r.seat = v.Seat
			break
		}
	}
	r.gameStart(s2c.Roominfo.State)
}

func (r *Robot) recvEBCoinEnter(s2c *pb.SEBCoinEnterRoom) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
	default:
		glog.Errorf("comein err -> %d", errcode)
		r.Close() //进入出错,关闭
		return
	}
	roominfo := s2c.GetRoominfo()
	r.gtype = roominfo.Gtype
	r.rtype = roominfo.Rtype
	r.dtype = roominfo.Dtype
	r.roomid = roominfo.Roomid
	userinfo := s2c.GetUserinfo()
	for _, v := range userinfo {
		//只返回坐下玩家
		if v.Userid == r.data.Userid {
			glog.Debugf("comein user info -> %s", v.Userid)
			r.seat = v.Seat
			break
		}
	}
	r.gameStart(s2c.Roominfo.State)
}

func (r *Robot) recvEBEnter(s2c *pb.SEBEnterRoom) {
	var errcode = s2c.GetError()
	switch errcode {
	case pb.OK:
	default:
		glog.Errorf("comein err -> %d", errcode)
		r.Close() //进入出错,关闭
		return
	}
	roominfo := s2c.GetRoominfo()
	r.gtype = roominfo.Gtype
	r.rtype = roominfo.Rtype
	r.dtype = roominfo.Dtype
	r.roomid = roominfo.Roomid
	userinfo := s2c.GetUserinfo()
	for _, v := range userinfo {
		//只返回坐下玩家
		if v.Userid == r.data.Userid {
			glog.Debugf("comein user info -> %s", v.Userid)
			r.seat = v.Seat
			break
		}
	}
	r.gameStart(s2c.Roominfo.State)
}

//进入房间
func (r *Robot) recvEBCamein(s2c *pb.SEBCamein) {
	if s2c.GetUserinfo().GetUserid() == r.data.Userid {
	}
}

func (r *Robot) recvEBFreeCamein(s2c *pb.SEBFreeCamein) {
	if s2c.GetUserinfo().GetUserid() == r.data.Userid {
	}
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
