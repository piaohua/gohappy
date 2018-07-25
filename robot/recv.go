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
	switch msg.(type) {
	//game
	case *pb.SRegist:
		r.recvRegist(msg.(*pb.SRegist))
	case *pb.SLogin:
		r.recvLogin(msg.(*pb.SLogin))
	case *pb.SUserData:
		r.recvdata(msg.(*pb.SUserData))
	case *pb.SPushCurrency:
		r.recvPushCurrency(msg.(*pb.SPushCurrency))
	case *pb.SPing:
		r.recvPing(msg.(*pb.SPing))
	//niu
	case *pb.SNNFreeEnterRoom:
		r.recvNNFreeEnter(msg.(*pb.SNNFreeEnterRoom))
	case *pb.SNNFreeCamein:
		r.recvNNFreeCamein(msg.(*pb.SNNFreeCamein))
	case *pb.SNNFreeGamestart:
		r.recvNNFreeStart(msg.(*pb.SNNFreeGamestart))
	case *pb.SNNFreeGameover:
		r.recvNNFreeGameover(msg.(*pb.SNNFreeGameover))
	case *pb.SNNFreeBet:
		r.recvNNFreeBet(msg.(*pb.SNNFreeBet))
	case *pb.SNNCoinEnterRoom:
		r.recvNNCoinEnter(msg.(*pb.SNNCoinEnterRoom))
	case *pb.SNNCoinGameover:
		r.recvNNCoinGameover(msg.(*pb.SNNCoinGameover))
	case *pb.SNNGameover:
		r.recvNNGameover(msg.(*pb.SNNGameover))
	case *pb.SNNEnterRoom:
		r.recvNNEnter(msg.(*pb.SNNEnterRoom))
	case *pb.SNNLeave:
		r.recvNNLeave(msg.(*pb.SNNLeave))
	case *pb.SNNCamein:
		r.recvNNCamein(msg.(*pb.SNNCamein))
	case *pb.SNNPushState:
		r.recvNNPushstate(msg.(*pb.SNNPushState))
	case *pb.SNNDraw:
		r.recvNNDraw(msg.(*pb.SNNDraw))
	case *pb.SNNDealer:
		r.recvNNDealer(msg.(*pb.SNNDealer))
	case *pb.SNNPushDealer:
		r.recvNNPushDealer(msg.(*pb.SNNPushDealer))
	case *pb.SNNPushDrawCoin:
	case *pb.SNNBet:
	case *pb.SNNiu:
	default:
		glog.Errorf("unknow message: %#v", msg)
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
	r.sendNNEntryRoom()
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
}

func (r *Robot) gameStart(state int32) {
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

// vim: set foldmethod=marker foldmarker=//',//.:
