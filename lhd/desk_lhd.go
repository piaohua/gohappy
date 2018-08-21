/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-04-30 18:01:29
 * Filename      : desk_lhd.go
 * Description   : 私人场玩牌逻辑
 * *******************************************************/
package main

import (
	"gohappy/game/config"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
	"gohappy/game/handler"
)

//'踢出不足坐下玩家或超额玩家
func (t *Desk) limitOver() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0): //自由
	case int32(pb.ROOM_TYPE1): //私人
		if !t.DeskData.Pub {
			//return
		}
	case int32(pb.ROOM_TYPE2): //百人
		return
	}
	for k, v := range t.roles {
		coin := v.User.GetCoin()
		//if t.DeskData.Maximum == 0 {
		//	if coin >= t.DeskData.Minimum {
		//		continue
		//	}
		//} else {
		//	if coin >= t.DeskData.Minimum &&
		//		coin < t.DeskData.Maximum {
		//		continue
		//	}
		//}
		if coin >= t.DeskData.Minimum { //离场限制
			continue
		}
		errcode := t.leave(k)
		if errcode != pb.OK {
			continue
		}
		//离开状态消息
		t.userLeaveDesk(k)
	}
}

// 踢除离线玩家
func (t *Desk) kickOffline() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE0), //自由
		int32(pb.ROOM_TYPE1), //私人
		int32(pb.ROOM_TYPE2): //百人
		for k, v := range t.roles {
			if !v.Offline {
				continue
			}
			errcode := t.leave(k)
			if errcode != pb.OK {
				continue
			}
			//离开状态消息
			t.userLeaveDesk(k)
		}
	}
}

// pub房间人数为0时解散
func (t *Desk) checkPubOver() {
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE1): //私人
		if !t.DeskData.Pub {
			//return
		}
	default:
		//return
	}
	if len(t.roles) != 0 {
		return
	}
	g := config.GetGame(t.DeskData.Unique)
	if g.Id == t.DeskData.Unique {
		return //配置房间不关闭
	}
	//停止服务
	msg1 := new(pb.ServeStop)
	t.selfPid.Tell(msg1)
}

func (t *Desk) checkPubOver2() {
	switch t.state {
	case int32(pb.STATE_READY):
		t.closeTime++
		if t.closeTime == 60 {
			t.closeTime = 0
			t.checkPubOver()
		}
	default:
		t.closeTime = 0
	}
}

//.

//'结束游戏
//是否过期
func (t *Desk) checkExpire() bool {
	var now = utils.Timestamp()
	if now > int64(t.DeskData.Expire) {
		glog.Debugf("game stop expire -> %d, %d",
			t.DeskData.Expire, now)
		return true
	}
	return false
}

//是否结束游戏
func (t *Desk) checkOver() bool {
	if t.DeskData.Round == t.DeskGame.Round {
		glog.Debugf("game stop round -> %d, %d",
			t.DeskGame.Round, t.DeskData.Round)
		return true
	}
	return t.checkExpire()
}

//结束牌局
func (t *Desk) gameStop() {
	if !t.checkOver() {
		return
	}
	if t.DeskData.Pub { //大厅房间不解散
		//return
	}
	//停止服务
	msg1 := new(pb.ServeStop)
	t.selfPid.Tell(msg1)
}

//抽水
func (t *Desk) drawcoin(userid string, val int64) int64 {
	if val <= 0 {
		return val
	}
	var num int64 = handler.DrawCoin(t.DeskData.Rtype, t.DeskData.Mode, val)
	switch t.DeskData.Rtype {
	case int32(pb.ROOM_TYPE2): //百人
		//反佣和收益消息,抽成日志记录 val - num
		msg2 := handler.AgentProfitNumMsg(userid, t.DeskData.Gtype, num)
		t.send3userid(userid, msg2)
	}
	return val - num
}
//.

// vim: set foldmethod=marker foldmarker=//',//.:
