/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-01-22 17:06:19
 * Filename      : sender.go
 * Description   : 机器人
 * *******************************************************/
package main

import (
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

//' 登录

//sendRegist 发送注册请求
func (c *Robot) sendRegist() {
	c2s := new(pb.CRegist)
	c2s.Phone = c.data.Phone
	c2s.Nickname = c.data.Nickname
	passwd := cfg.Section("robot").Key("passwd").Value()
	c2s.Password = utils.Md5(passwd)
	c.Sender(c2s)
}

//sendLogin 发送登录请求
func (c *Robot) sendLogin() {
	c2s := new(pb.CLogin)
	c2s.Phone = c.data.Phone
	passwd := cfg.Section("robot").Key("passwd").Value()
	c2s.Password = utils.Md5(passwd)
	c.SendDefer(c2s)
}

//sendUserData 获取玩家数据
func (c *Robot) sendUserData() {
	c2s := new(pb.CUserData)
	c2s.Userid = c.data.Userid
	c.Sender(c2s)
}

//SendPing 心跳
func (c *Robot) sendPing() {
	c2s := new(pb.CPing)
	c2s.Time = 1 //uint32(utils.Timestamp())
	c.Sender(c2s)
}

//addCurrency 添加货币
func (c *Robot) addCurrency() {
	msg4 := &pb.PayCurrency{
		Userid: c.data.Userid,
		Type:   int32(pb.LOG_TYPE44),
		Coin:   200000,
	}
	rolePid.Tell(msg4)
}

//SendDefer 延迟发送
func (c *Robot) SendDefer(msg interface{}) {
	utils.Sleep(utils.RandIntN(5) + 3) //随机
	c.Sender(msg)
}

//.

//' niu

//sendNNLeave 离开
func (c *Robot) sendNNLeave() {
	c2s := new(pb.CNNLeave)
	c.Sender(c2s)
}

//sendNNReady 准备
func (c *Robot) sendNNReady() {
	c2s := new(pb.CNNReady)
	c2s.Ready = true
	c.SendDefer(c2s)
}

//sendNNDealer 抢庄
func (c *Robot) sendNNDealer() {
	c2s := new(pb.CNNDealer)
	if utils.RandIntN(100) > 50 {
		c2s.Dealer = true
		c2s.Num = uint32(utils.RandIntN(2) + 1)
	}
	c.SendDefer(c2s)
}

//sendNNiu 提交
func (c *Robot) sendNNiu() {
	c2s := new(pb.CNNiu)
	c.SendDefer(c2s)
}

//sendNNStandup 玩家离坐
func (c *Robot) sendNNStandup() {
	c.sendNNLeave()
	utils.Sleep(2)
	c.Close() //下线
}

//sendNNEntryRoom 进入房间
func (c *Robot) sendNNEntryRoom() {
	switch c.rtype {
	case int32(pb.ROOM_TYPE0):
		glog.Debugf("enter roomid %s", c.roomid)
		c2s := new(pb.CNNCoinEnterRoom)
		c2s.Id = c.roomid
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE1):
		glog.Debugf("enter room code %s", c.code)
		c2s := new(pb.CNNEnterRoom)
		c2s.Code = c.code
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE2):
		c2s := new(pb.CNNFreeEnterRoom)
		c.SendDefer(c2s)
	}
}

//sendNNBet 玩家下注
func (c *Robot) sendNNBet() {
	switch c.rtype {
	case int32(pb.ROOM_TYPE0):
		c2s := new(pb.CNNBet)
		c2s.Seatbet = c.seat
		c2s.Value = uint32(utils.RandIntN(10))
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE1):
		c2s := new(pb.CNNBet)
		c2s.Seatbet = c.seat
		c2s.Value = 1
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE2):
		//随机下注次数
		c.bits = uint32(utils.RandInt32N(20) + 1)
		c.bitNum = uint32(utils.RandInt32N(7) * 5000)
		c.sendNNFreeBet() //下注
	}
}

//sendNNBet 玩家下注
func (c *Robot) sendNNFreeBet() {
	//不同游戏位置不同
	var seats = []uint32{2, 3, 4, 5, 6, 7, 8, 9}
	var bets = []uint32{100, 500, 1000, 5000, 10000}
	var coin uint32 = uint32(c.data.Coin) / 4
	var max int
	for i := 4; i >= 0; i-- {
		if coin >= bets[i] {
			max = i
			break
		}
	}
	var v int
	switch max {
	case 0:
		v = max
	default:
		v = utils.RandIntN(max) //随机
	}
	var k = utils.RandIntN(len(seats)) //随机
	c2s := &pb.CNNFreeBet{
		Value: bets[v],
		Seat:  seats[k],
	}
	c.SendDefer(c2s)
}

//.

//' ebg

//sendEBLeave 离开
func (c *Robot) sendEBLeave() {
	c2s := new(pb.CEBLeave)
	c.Sender(c2s)
}

//sendEBReady 准备
func (c *Robot) sendEBReady() {
	c2s := new(pb.CEBReady)
	c2s.Ready = true
	c.SendDefer(c2s)
}

//sendEBDealer 抢庄
func (c *Robot) sendEBDealer() {
	c2s := new(pb.CEBDealer)
	if utils.RandIntN(100) > 50 {
		c2s.Dealer = true
		c2s.Num = uint32(utils.RandIntN(2) + 1)
	}
	c.SendDefer(c2s)
}

//sendEBiu 提交
func (c *Robot) sendEBiu() {
	c2s := new(pb.CEBiu)
	c.SendDefer(c2s)
}

//sendEBStandup 玩家离坐
func (c *Robot) sendEBStandup() {
	c.sendEBLeave()
	utils.Sleep(2)
	c.Close() //下线
}

//sendEBEntryRoom 进入房间
func (c *Robot) sendEBEntryRoom() {
	switch c.rtype {
	case int32(pb.ROOM_TYPE0):
		glog.Debugf("enter roomid %s", c.roomid)
		c2s := new(pb.CEBCoinEnterRoom)
		c2s.Id = c.roomid
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE1):
		glog.Debugf("enter room code %s", c.code)
		c2s := new(pb.CEBEnterRoom)
		c2s.Code = c.code
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE2):
		c2s := new(pb.CEBFreeEnterRoom)
		c.SendDefer(c2s)
	}
}

//sendEBBet 玩家下注
func (c *Robot) sendEBBet() {
	switch c.rtype {
	case int32(pb.ROOM_TYPE0):
		c2s := new(pb.CEBBet)
		c2s.Seatbet = c.seat
		c2s.Value = uint32(utils.RandIntN(10))
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE1):
		c2s := new(pb.CEBBet)
		c2s.Seatbet = c.seat
		c2s.Value = 1
		c.SendDefer(c2s)
	case int32(pb.ROOM_TYPE2):
		//随机下注次数
		c.bits = uint32(utils.RandInt32N(20) + 1)
		c.bitNum = uint32(utils.RandInt32N(7) * 5000)
		c.sendEBFreeBet() //下注
	}
}

//sendEBFreeBet 玩家下注
func (c *Robot) sendEBFreeBet() {
	//不同游戏位置不同
	var seats = []uint32{2, 3, 4, 5, 6, 7, 8, 9}
	var bets = []uint32{100, 500, 1000, 5000, 10000}
	var coin uint32 = uint32(c.data.Coin) / 4
	var max int
	for i := 4; i >= 0; i-- {
		if coin >= bets[i] {
			max = i
			break
		}
	}
	var v int
	switch max {
	case 0:
		v = max
	default:
		v = utils.RandIntN(max) //随机
	}
	var k = utils.RandIntN(len(seats)) //随机
	c2s := &pb.CEBFreeBet{
		Value: bets[v],
		Seat:  seats[k],
	}
	c.SendDefer(c2s)
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
