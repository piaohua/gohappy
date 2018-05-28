package main

import (
	"time"

	"gohappy/data"
	"gohappy/game/config"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//玩家数据请求处理
func (rs *RoleActor) handlerLogin(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CRegist:
		//注册消息
		arg := msg.(*pb.CRegist)
		glog.Debugf("CRegist %#v", arg)
		//重复登录
		stoc := new(pb.SRegist)
		stoc.Error = pb.RepeatLogin
		rs.Send(stoc)
	case *pb.CLogin:
		//登录消息
		arg := msg.(*pb.CLogin)
		glog.Debugf("CLogin %#v", arg)
		//重复登录
		stoc := new(pb.SLogin)
		stoc.Error = pb.RepeatLogin
		rs.Send(stoc)
	case *pb.CWxLogin:
		//登录消息
		arg := msg.(*pb.CWxLogin)
		glog.Debugf("CWxLogin %#v", arg)
		//重复登录
		stoc := new(pb.SWxLogin)
		stoc.Error = pb.RepeatLogin
		rs.Send(stoc)
	case *pb.CResetPwd:
		//重置密码消息
		arg := msg.(*pb.CResetPwd)
		glog.Debugf("CResetPwd %#v", arg)
		//重复登录
		stoc := new(pb.SResetPwd)
		stoc.Error = pb.RepeatLogin
		rs.Send(stoc)
	case *pb.CTourist:
		//登录消息
		arg := msg.(*pb.CTourist)
		glog.Debugf("CTourist %#v", arg)
		//重复登录
		stoc := new(pb.STourist)
		stoc.Error = pb.RepeatLogin
		rs.Send(stoc)
	case *pb.LoginElse:
		arg := msg.(*pb.LoginElse)
		rs.loginElse() //别处登录
		//响应登录
		rsp := new(pb.LoginedElse)
		rsp.Userid = arg.Userid
		rsp.Gate = arg.Gate
		ctx.Respond(rsp)
	case *pb.LoginSuccess:
		//登录成功处理
		arg := msg.(*pb.LoginSuccess)
		glog.Debugf("LoginSuccess %#v", arg)
		rs.logined(arg, ctx)
	case proto.Message:
		if rs.User == nil {
			glog.Errorf("user empty message %v", msg)
			return
		}
		rs.handlerUser(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//别处登录
func (rs *RoleActor) loginElse() {
	arg := new(pb.SLoginOut)
	glog.Debugf("SLoginOut %s", rs.User.Userid)
	arg.Rtype = int32(pb.LOGOUT_TYPE4) //别处登录
	rs.Send(arg)
	//已经断开
	if !rs.online {
		return
	}
	//断开连接
	rs.CloseWs()
	//离开游戏消息
	rs.leaveDesk()
	//同步数据
	rs.syncUser()
	//登出日志
	msg3 := &pb.LogLogout{
		Userid: rs.User.Userid,
		Event:  int32(pb.LOGOUT_TYPE4), //别处登录
	}
	rs.loggerPid.Tell(msg3)
	//表示已经断开
	rs.online = false
}

//离开游戏处理
func (rs *RoleActor) leaveDesk() {
	if rs.gamePid == nil {
		return
	}
	//离线
	msg2 := new(pb.OfflineDesk)
	if rs.User != nil {
		msg2.Userid = rs.User.GetUserid()
	}
	rs.gamePid.Tell(msg2)
	//下线
	msg3 := new(pb.LeaveDesk)
	if rs.User != nil {
		msg3.Userid = rs.User.GetUserid()
	}
	//rs.gamePid.Tell(msg3)
	timeout := 3 * time.Second
	res1, err1 := rs.gamePid.RequestFuture(msg3, timeout).Result()
	if err1 != nil {
		glog.Errorf("leave desk failed: %v", err1)
	}
	if response1, ok := res1.(*pb.LeftDesk); ok {
		glog.Debugf("left desk response1: %#v", response1)
		if response1.Error != pb.OK {
			glog.Errorf("leave desk failed: %#v", res1)
		} else {
			//不同游戏可能规则不同，由游戏控制是否离开
			rs.gamePid = nil
		}
	} else {
		glog.Errorf("leave desk failed: %#v", res1)
	}
}

//登录成功处理
func (rs *RoleActor) logined(arg *pb.LoginSuccess, ctx actor.Context) {
	if !rs.loginedGetUser(arg.Userid, ctx) {
		glog.Errorf("logined fail %s, %s", ctx.Self().String(), arg.Userid)
		rs.loginFailed(ctx)
		return
	}
	//断开旧连接
	rs.CloseWs()
	//新连接
	rs.wsPid = arg.WsPid
	//头像
	rs.setHeadImag(arg.IsRegist, ctx)
	//日志
	rs.loginedLog(arg)
	//登录成功
	rs.online = true
}

//登录后获取数据
func (rs *RoleActor) loginedGetUser(userid string, ctx actor.Context) bool {
	//关闭旧进程
	rs.loginElseGate(userid)
	//登录
	msg1 := &pb.GetUser{
		Userid:  userid,
		Gate:    cfg.Section(nodeName).Name(),
		RolePid: ctx.Self(),
	}
	timeout := 3 * time.Second
	res1, err1 := rs.rolePid.RequestFuture(msg1, timeout).Result()
	if err1 != nil {
		glog.Errorf("logined GetUser failed: %v", err1)
		return false
	}
	if response1, ok := res1.(*pb.GotUser); ok {
		user := new(data.User)
		err2 := json.Unmarshal(response1.Data, user)
		if err2 != nil {
			glog.Errorf("user Unmarshal err %v", err2)
			return false
		}
		if user.GetUserid() == "" {
			return false
		}
		rs.User = user
		return true
	}
	return false
}

//别处登录处理
func (rs *RoleActor) loginElseGate(userid string) {
	msg1 := new(pb.LoginElse)
	msg1.Userid = userid
	msg1.Gate = cfg.Section(nodeName).Name()
	timeout := 3 * time.Second
	res1, err1 := rs.rolePid.RequestFuture(msg1, timeout).Result()
	if err1 != nil {
		glog.Errorf("loginElseGate res1 %#v, err1 %v", res1, err1)
	}
}

//登录失败 TODO 优化
func (rs *RoleActor) loginFailed(ctx actor.Context) {
	ctx.Self().Tell(new(pb.OfflineStop))
}

//默认头像
func (rs *RoleActor) setHeadImag(isRegist bool, ctx actor.Context) {
	if !isRegist {
		return
	}
	if rs.User == nil {
		return
	}
	if rs.User.GetPhoto() != "" {
		return
	}
	if len(HeadImagList) == 0 {
		return
	}
	head := cfg.Section("domain").Key("headimag").Value()
	if head == "" {
		return
	}
	i := utils.RandIntN(len(HeadImagList))
	rs.User.Photo = head + "/" + HeadImagList[i].Photo
}

//登录成功日志处理
func (rs *RoleActor) loginedLog(arg *pb.LoginSuccess) {
	rs.User.LoginIp = arg.Ip
	//连续登录
	rs.loginPrizeInit()
	if arg.IsRegist {
		//注册ip
		rs.User.RegistIp = arg.Ip
		if !rs.User.IsTourist() {
			//注册奖励发放
			var diamond int64 = int64(config.GetEnv(data.ENV1))
			var coin int64 = int64(config.GetEnv(data.ENV2))
			var chip int64 = int64(config.GetEnv(data.ENV3))
			var card int64 = int64(config.GetEnv(data.ENV4))
			rs.addCurrency(diamond, coin, card, chip, int32(pb.LOG_TYPE1))
			//注册日志
			msg1 := &pb.LogRegist{
				Userid:   rs.User.Userid,
				Nickname: rs.User.Nickname,
				Ip:       arg.Ip,
			}
			rs.loggerPid.Tell(msg1)
		}
	}
	//登录日志
	msg2 := &pb.LogLogin{
		Userid: rs.User.Userid,
		Ip:     arg.Ip,
	}
	rs.loggerPid.Tell(msg2)
	//TODO test
	rs.loginedLog2()
}

//test
func (rs *RoleActor) loginedLog2() {
	var diamond, coin int64
	if rs.User.GetDiamond() < 10000 {
		diamond = 10000
	}
	if rs.User.GetCoin() < 5000000 {
		coin = 5000000
	}
	glog.Debugf("loginedLog2 userid %s, diamond %d, coin %d",
		rs.User.GetUserid(), diamond, coin)
	if diamond == 0 && coin == 0 {
		return
	}
	rs.addCurrency(diamond, coin, 0, 0, int32(pb.LOG_TYPE11))
}

//Send 发送消息
func (rs *RoleActor) Send(msg interface{}) {
	//glog.Debugf("Send %#v", msg)
	if rs.stopCh == nil {
		glog.Errorf("rs msg channel closed %v", msg)
		return
	}
	if rs.wsPid == nil {
		glog.Errorf("ws pid stoped %v", msg)
		return
	}
	//glog.Debugf("send message %s", rs.wsPid.String())
	select {
	case <-rs.stopCh:
		return
	default:
	}
	select {
	case <-rs.stopCh:
		return
	default:
		//glog.Debugf("send message %#v", msg)
		rs.wsPid.Tell(msg)
	}
}

//StopRs 关闭
func (rs *RoleActor) StopRs() {
	select {
	case <-rs.stopCh:
		return
	default:
		//停止发送消息
		close(rs.stopCh)
	}
	//停止
	rs.pid.Stop()
}

//CloseWs 关闭连接
func (rs *RoleActor) CloseWs() {
	if rs.wsPid == nil {
		return
	}
	glog.Debugf("CloseWs userid: %s", rs.wsPid.String())
	msg1 := new(pb.ServeStop)
	//关闭连接
	rs.wsPid.Tell(msg1)
	//断开
	rs.wsPid = nil
}
