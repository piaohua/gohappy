/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-01-22 17:06:19
 * Filename      : sender.go
 * Description   : 机器人
 * *******************************************************/
package main

import (
	"crypto/md5"
	"encoding/hex"

	"gohappy/pb"
	"utils"
)

//' 登录

//SendRegist 发送注册请求
func (c *Robot) SendRegist() {
	ctos := new(pb.CRegist)
	ctos.Phone = c.data.Phone
	ctos.Nickname = c.data.Nickname
	h := md5.New()
	passwd := cfg.Section("robot").Key("passwd").Value()
	h.Write([]byte(passwd)) // 需要加密的字符串为
	pwd := hex.EncodeToString(h.Sum(nil))
	ctos.Password = pwd
	c.Sender(ctos)
}

//SendLogin 发送登录请求
func (c *Robot) SendLogin() {
	ctos := new(pb.CLogin)
	ctos.Phone = c.data.Phone
	h := md5.New()
	passwd := cfg.Section("robot").Key("passwd").Value()
	h.Write([]byte(passwd)) // 需要加密的字符串为
	pwd := hex.EncodeToString(h.Sum(nil))
	ctos.Password = pwd
	//glog.Infof("ctos -> %#v", ctos)
	utils.Sleep(2)
	c.Sender(ctos)
}

//SendUserData 获取玩家数据
func (c *Robot) SendUserData() {
	ctos := new(pb.CUserData)
	ctos.Userid = c.data.Userid
	c.Sender(ctos)
}

//SendPing 心跳
func (c *Robot) SendPing() {
	ctos := new(pb.CPing)
	//ctos.Time := uint32(utils.Timestamp())
	ctos.Time = 1
	//glog.Debugf("ping : %#v", ctos)
	c.Sender(ctos)
}

//AddCurrency 添加货币
func (c *Robot) AddCurrency() {
	msg4 := &pb.PayCurrency{
		Userid: c.data.Userid,
		Type:   int32(pb.LOG_TYPE44),
		Coin:   200000,
	}
	rolePid.Tell(msg4)
}

//.

//' 百人

//SendNNLeave 离开
func (c *Robot) SendNNLeave() {
	ctos := new(pb.CNNLeave)
	c.Sender(ctos)
}

//SendNNEntryRoom 进入房间
func (c *Robot) SendNNEntryRoom() {
	ctos := new(pb.CNNFreeEnterRoom)
	//glog.Debugf("enter roomid %s", roomid)
	utils.Sleep(2)
	c.Sender(ctos)
}

//SendNNSitDown 玩家入坐
func (c *Robot) SendNNSitDown() {
	seat := uint32(utils.RandInt32N(4) + 1) //随机
	ctos := &pb.CNNFreeSit{
		State: true,
		Seat:  seat,
	}
	c.sits++ //尝试次数
	utils.Sleep(2)
	c.Sender(ctos)
}

//SendNNStandup 玩家离坐
func (c *Robot) SendNNStandup() {
	ctos := &pb.CNNFreeSit{
		State: false,
		Seat:  c.seat,
	}
	utils.Sleep(2)
	c.Sender(ctos)
	utils.Sleep(2)
	c.SendNNLeave()
	utils.Sleep(2)
	c.Close() //下线
}

//SendNNRoomBet 玩家下注
func (c *Robot) SendNNRoomBet() {
	//不同游戏位置不同
	var seats = []uint32{2, 3, 4, 5, 6, 7, 8, 9}
	var bets = []uint32{100, 1000, 10000, 50000, 100000, 200000}
	var coin uint32 = uint32(c.data.Coin) / 4
	var max int
	for i := 5; i >= 0; i-- {
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
	ctos := &pb.CNNFreeBet{
		Value: bets[v],
		Seat:  seats[k],
	}
	var t1 int = utils.RandIntN(4) + 1 //随机
	utils.Sleep(t1)
	c.Sender(ctos)
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
