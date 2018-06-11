/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-01-13 10:56:33
 * Filename      : desk_logic.go
 * Description   : 内部方法
 * *******************************************************/
package main

import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// internal function

//NewDesk 新建一张牌桌
func NewDesk(deskData *data.DeskData) *Desk {
	desk := &Desk{
		roles:  make(map[string]*data.DeskRole),
		seats:  make(map[uint32]*data.DeskSeat),
		router: make(map[string]string),
		stopCh: make(chan struct{}),
	}
	desk.DeskData = deskData
	return desk
}

//InitDesk 初始化
func (t *Desk) InitDesk() {
	t.DeskGame = new(data.DeskGame)
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
	case int32(pb.ROOM_TYPE1): //私人
		t.DeskPriv = new(data.DeskPriv)
		t.DeskPriv.PrivScore = make(map[string]int64)
		t.DeskPriv.Joins = make(map[string]uint32)
	case int32(pb.ROOM_TYPE2): //百人
		t.DeskFree = new(data.DeskFree)
		t.DeskFree.Dealers = make(map[string]int64) //上庄列表,userid: carry
		t.freeInit()
	}
}

//房间消息广播
func (t *Desk) broadcast(msg interface{}) {
	for _, v := range t.roles {
		if v == nil {
			continue
		}
		if v.Pid == nil {
			continue
		}
		if v.Offline {
			continue
		}
		v.Pid.Tell(msg)
	}
}

//房间消息广播(除userid外)
func (t *Desk) broadcast2(userid string, msg interface{}) {
	for k, v := range t.roles {
		if v == nil {
			continue
		}
		if v.Pid == nil {
			continue
		}
		if v.Offline {
			continue
		}
		if k != userid {
			v.Pid.Tell(msg)
		}
	}
}

//房间消息广播(除seat外)
func (t *Desk) broadcast3(seat uint32, msg interface{}) {
	if v, ok := t.seats[seat]; ok && v != nil {
		t.broadcast2(v.Userid, msg)
	}
}

//给玩家发送消息
func (t *Desk) send2userid(userid string, msg interface{}) {
	if v, ok := t.roles[userid]; ok && v != nil {
		v.Pid.Tell(msg)
	}
}

//给位置发送消息
func (t *Desk) send2seat(seat uint32, msg interface{}) {
	if v, ok := t.seats[seat]; ok && v != nil {
		t.send2userid(v.Userid, msg)
	}
}

//给玩家发送消息,离线也发送
func (t *Desk) send3userid(userid string, msg interface{}) {
	if v, ok := t.roles[userid]; ok && v != nil {
		if v.Offline {
			return
		}
		v.Pid.Tell(msg)
	}
}

//给位置发送消息,离线也发送
func (t *Desk) send3seat(seat uint32, msg interface{}) {
	if v, ok := t.seats[seat]; ok && v != nil {
		t.send3userid(v.Userid, msg)
	}
}

//获取路由
func (t *Desk) getRouter(ctx actor.Context) string {
	glog.Debugf("getRouter %s", ctx.Sender().String())
	return t.router[ctx.Sender().String()]
}

//获取进程pid
func (t *Desk) getPid(userid string) *actor.PID {
	if v, ok := t.roles[userid]; ok && v != nil {
		return v.Pid
	}
	return nil
}

//获取玩家数据
func (t *Desk) getPlayer(userid string) *data.User {
	if v, ok := t.roles[userid]; ok && v != nil {
		return v.User
	}
	return nil
}

//获取玩家数据
func (t *Desk) getUserBySeat(seat uint32) *data.User {
	if v, ok := t.seats[seat]; ok && v != nil {
		return t.getPlayer(v.Userid)
	}
	return nil
}

//获取位置
func (t *Desk) getSeat(userid string) uint32 {
	if v, ok := t.roles[userid]; ok && v != nil {
		return v.Seat
	}
	return 0
}

//获取位置
func (t *Desk) getUserid(seat uint32) string {
	if v, ok := t.seats[seat]; ok && v != nil {
		return v.Userid
	}
	return ""
}

//玩家是否在线
func (t *Desk) isOnline(userid string) bool {
	if v, ok := t.roles[userid]; ok && v != nil {
		return !v.Offline
	}
	return false
}

//设置玩家是否离线
func (t *Desk) setOffline(userid string, offline bool) {
	if v, ok := t.roles[userid]; ok && v != nil {
		v.Offline = offline
	}
}

//获取手牌
func (t *Desk) getHandCards(seat uint32) []uint32 {
	//房间类型 百人场
	if t.DeskData.Rtype == int32(pb.ROOM_TYPE2) &&
		t.DeskFree != nil {
		return t.DeskFree.Cards[seat]
	}
	//房间类型 非百人场
	if v, ok := t.seats[seat]; ok && v != nil {
		return v.Cards
	}
	glog.Errorf("getHandCards %d", seat)
	t.printOver()
	return []uint32{}
	//panic(fmt.Sprintf("getHandCards error:%d", seat))
}

//玩家牌力
func (t *Desk) getPower(seat uint32) uint32 {
	if v, ok := t.seats[seat]; ok && v != nil {
		return v.Power
	}
	return 0
}

//位置下注
func (t *Desk) getBets(seat uint32) int64 {
	if v, ok := t.seats[seat]; ok && v != nil {
		return v.Bet
	}
	return 0
}

//是否翻牌
func (t *Desk) isSee(seat uint32) bool {
	if t.DeskHua == nil {
		return false
	}
	if v, ok := t.DeskHua.ActSeats[seat]; ok {
		return v.See
	}
	return false
}

// vim: set foldmethod=marker foldmarker=//',//.:
