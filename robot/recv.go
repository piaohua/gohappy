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
	"utils"
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
	case *pb.SNNLeave:
		r.recvNNLeave(msg.(*pb.SNNLeave))
	case *pb.SNNFreeEnterRoom:
		r.recvNNComein(msg.(*pb.SNNFreeEnterRoom))
	case *pb.SNNCamein:
		r.recvNNCamein(msg.(*pb.SNNCamein))
	case *pb.SNNFreeCamein:
		r.recvNNFreeCamein(msg.(*pb.SNNFreeCamein))
	case *pb.SNNFreeSit:
		r.recvNNSitDown(msg.(*pb.SNNFreeSit))
	case *pb.SNNFreeBet:
		r.recvNNBet(msg.(*pb.SNNFreeBet))
	case *pb.SNNPushState:
		r.recvNNGamestate(msg.(*pb.SNNPushState))
	case *pb.SNNFreeGameover:
		r.recvNNGameover(msg.(*pb.SNNFreeGameover))
	case *pb.SNNDraw:
		r.recvNNDraw(msg.(*pb.SNNDraw))
	case *pb.SNNFreeGamestart:
		r.recvNNStart(msg.(*pb.SNNFreeGamestart))
	default:
		glog.Errorf("unknow message: %#v", msg)
	}
}

//.

//' 接收到服务器注册返回
func (r *Robot) recvRegist(stoc *pb.SRegist) {
	var errcode = stoc.GetError()
	switch errcode {
	case pb.OK:
		Logined(r.data.Phone, r.ltype) //登录成功
		r.regist = true                //注册成功
		r.data.Userid = stoc.GetUserid()
		glog.Infof("regist successful -> %s", r.data.Userid)
		r.SendUserData() // 获取玩家数据
		return
	case pb.PhoneRegisted:
		glog.Infof("phone registed -> %s", r.data.Phone)
		r.SendLogin() //尝试登录
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
func (r *Robot) recvLogin(stoc *pb.SLogin) {
	var errcode = stoc.GetError()
	switch errcode {
	case pb.OK:
		Logined(r.data.Phone, r.ltype) //登录成功
		r.data.Userid = stoc.GetUserid()
		glog.Infof("login successful -> %s", r.data.Userid)
		r.SendUserData() // 获取玩家数据
		return
	default:
		glog.Infof("login err -> %d", errcode)
	}
	r.Close()
}

//.

//' 接收到玩家数据
func (r *Robot) recvdata(stoc *pb.SUserData) {
	var errcode = stoc.GetError()
	if errcode != pb.OK {
		glog.Infof("get data err -> %d", errcode)
		r.Close() //断开
		return
	}
	userdata := stoc.GetData()
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
		//充值
		r.AddCurrency()
		//r.SendNNStandup()
		//return
	}
	//进入房间
	r.SendNNEntryRoom()
}

//更新金币
func (r *Robot) recvPushCurrency(stoc *pb.SPushCurrency) {
	currencyData := stoc.GetData()
	r.data.Coin += currencyData.GetCoin()
	r.data.Card += currencyData.GetCard()
	r.data.Chip += currencyData.GetChip()
	r.data.Diamond += currencyData.GetDiamond()
	if r.data.Coin < 650000 {
		//r.SendNNStandup()
	}
}

//游戏
func (r *Robot) recvPing(stoc *pb.SPing) {
	//glog.Debugf("ping %s", r.data.Userid)
	//TODO
}

//.

//' 离开房间
func (r *Robot) recvNNLeave(stoc *pb.SNNLeave) {
	if stoc.GetUserid() == r.data.Userid {
		r.Close() //下线
	}
}

//.

//' 进入房间
func (r *Robot) recvNNComein(stoc *pb.SNNFreeEnterRoom) {
	var errcode = stoc.GetError()
	switch errcode {
	case pb.OK:
		roominfo := stoc.GetRoominfo()
		r.gtype = roominfo.Gtype
		r.rtype = roominfo.Rtype
		r.dtype = roominfo.Dtype
		r.roomid = roominfo.Roomid
		userinfo := stoc.GetUserinfo()
		for _, v := range userinfo {
			//只返回坐下玩家
			if v.Userid == r.data.Userid {
				glog.Debugf("comein user info -> %s", v.Userid)
				r.seat = v.Seat
				break
			}
		}
		glog.Debugf("comein desk info -> %#v", roominfo)
		//坐下
		r.SendNNSitDown()
		//下注
		glog.Debugf("comein desk state -> %d", roominfo.State)
		switch roominfo.State {
		case int32(pb.STATE_BET):
			r.SendNNRoomBet()
		}
	default:
		glog.Infof("comein err -> %d", errcode)
		r.Close() //进入出错,关闭
	}
}

//进入房间
func (r *Robot) recvNNCamein(stoc *pb.SNNCamein) {
	if stoc.GetUserinfo().GetUserid() == r.data.Userid {
	}
}

func (r *Robot) recvNNFreeCamein(stoc *pb.SNNFreeCamein) {
	if stoc.GetUserinfo().GetUserid() == r.data.Userid {
	}
}

//.

//' 百人

//坐下
func (r *Robot) recvNNSitDown(stoc *pb.SNNFreeSit) {
	var errcode = stoc.GetError()
	var seat uint32 = stoc.GetSeat()
	var userid string = stoc.GetUserid()
	if userid != r.data.Userid {
		return
	}
	switch errcode {
	case pb.OK:
		r.seat = seat //坐下位置
	default:
		if r.sits > 4 { //尝试次数过多
			r.SendNNStandup()
		} else {
			r.SendNNSitDown()
		}
	}
}

//下注
func (r *Robot) recvNNBet(stoc *pb.SNNFreeBet) {
	var errcode = stoc.GetError()
	var userid string = stoc.GetUserid()
	glog.Debugf("bet userid %s, errcode %d", userid, errcode)
	switch errcode {
	case pb.OK:
		if userid == r.data.Userid {
			val := stoc.GetValue()
			if val > r.bitNum {
				r.bitNum = 0
			} else {
				r.bitNum -= val
			}
			r.bits--
		}
		if r.bits > 0 && r.bitNum > 0 {
			r.SendNNRoomBet()
		}
	default:
		r.SendNNStandup()
	}
}

//状态更新
func (r *Robot) recvNNGamestate(stoc *pb.SNNPushState) {
	var state int32 = stoc.GetState()
	r.gameStart(state)
}

//结束
func (r *Robot) recvNNGameover(stoc *pb.SNNFreeGameover) {
	r.bits = 0
	r.bitNum = 0
	r.round++
	if r.round >= 30 { //打10局下线
		r.SendNNStandup()
		return
	}
	//TODO
}

func (r *Robot) gameStart(state int32) {
	switch state {
	case int32(pb.STATE_READY):
	case int32(pb.STATE_BET):
		//随机下注次数
		r.bits = uint32(utils.RandInt32N(20) + 1)
		r.bitNum = uint32(utils.RandInt32N(7) * 5000)
		r.SendNNRoomBet() //下注
	case int32(pb.STATE_OVER):
	default:
		r.SendNNStandup()
	}
}

//开始
func (r *Robot) recvNNStart(stoc *pb.SNNFreeGamestart) {
	var state int32 = stoc.GetState()
	r.gameStart(state)
}

//开始
func (r *Robot) recvNNDraw(stoc *pb.SNNDraw) {
	//TODO
}

//.

// vim: set foldmethod=marker foldmarker=//',//.:
