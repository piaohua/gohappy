package main

import (
	"time"

	"gohappy/game/login"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//玩家数据请求处理
func (ws *WSConn) handlerLogin(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CRegist:
		//注册消息
		arg := msg.(*pb.CRegist)
		glog.Debugf("CRegist %#v", arg)
		ws.regist(arg, ctx)
	case *pb.CLogin:
		//登录消息
		arg := msg.(*pb.CLogin)
		glog.Debugf("CLogin %#v", arg)
		ws.login(arg, ctx)
	case *pb.CWxLogin:
		//登录消息
		arg := msg.(*pb.CWxLogin)
		glog.Debugf("CWxLogin %#v", arg)
		ws.wxlogin(arg, ctx)
	case *pb.CResetPwd:
		//重置密码消息
		arg := msg.(*pb.CResetPwd)
		glog.Debugf("CResetPwd %#v", arg)
		//TODO 暂时屏蔽使用
		//ws.resetPwd(arg, ctx)
	case *pb.CTourist:
		//登录消息
		arg := msg.(*pb.CTourist)
		glog.Debugf("CTourist %#v", arg)
		ws.touristLogin(arg, ctx)
	case proto.Message:
		//响应
		if ws.online {
			ws.Send(msg)
		}
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//游客登录
func (ws *WSConn) touristLogin(arg *pb.CTourist, ctx actor.Context) {
	//检测参数
	key := cfg.Section("gate").Key("tourist").Value()
	stoc := login.TouristLoginCheck(arg, key)
	if stoc.Error != pb.OK {
		ws.Send(stoc)
		return
	}
	//重复登录
	if ws.online {
		stoc.Error = pb.RepeatLogin
		ws.Send(stoc)
		return
	}
	msg1 := new(pb.TouristLogin)
	msg1.Account = arg.GetAccount()
	msg1.Password = arg.GetPassword()
	msg1.Registip = ws.GetIPAddr()
	//登录
	res1 := ws.reqRole(msg1, ctx)
	var response1 *pb.TouristLogined
	var ok bool
	if response1, ok = res1.(*pb.TouristLogined); ok {
		if response1.Error != pb.OK {
			glog.Errorf("CTourist fail %d", response1.Error)
			stoc.Error = response1.Error
			ws.Send(stoc)
			return
		}
	} else {
		glog.Error("CTourist fail")
		stoc.Error = pb.LoginError
		ws.Send(stoc)
		return
	}
	userid := response1.GetUserid()
	glog.Debugf("tourist login successfully %s", userid)
	if !ws.logining(userid, ctx) {
		glog.Debugf("login failed %s", ctx.Self())
		stoc.Error = pb.LoginError
		ws.Send(stoc)
		return
	}
	stoc.Userid = userid
	glog.Debugf("tourist login successfully %s", userid)
	ws.Send(stoc)
	//成功后处理
	ws.logined(userid, response1.IsRegist, ctx)
}

func (ws *WSConn) resetPwd(arg *pb.CResetPwd, ctx actor.Context) {
	stoc := login.RestPwdCheck(arg)
	if stoc.Error != pb.OK {
		ws.Send(stoc)
		return
	}
	//已经登录
	if ws.online {
		stoc.Error = pb.RepeatLogin
		ws.Send(stoc)
		return
	}
	//重置
	res1 := ws.reqRole(arg, ctx)
	var response1 *pb.SResetPwd
	var ok bool
	if response1, ok = res1.(*pb.SResetPwd); ok {
		if response1.Error != pb.OK {
			glog.Errorf("CRegist fail %d", response1.Error)
			stoc.Error = response1.Error
			ws.Send(stoc)
			return
		}
	} else {
		glog.Error("CResetPwd fail")
		stoc.Error = pb.ResetPwdFaild
		ws.Send(stoc)
		return
	}
	userid := response1.GetUserid()
	glog.Debugf("CResetPwd successfully %s", userid)
	if !ws.logining(userid, ctx) {
		glog.Debugf("CResetPwd failed %s", ctx.Self())
		stoc.Error = pb.ResetPwdFaild
		ws.Send(stoc)
		return
	}
	stoc.Userid = userid
	glog.Debugf("CResetPwd successfully %s", userid)
	ws.Send(stoc)
	//成功后处理
	ws.logined(userid, true, ctx)
}

//TODO 注册ip限制
func (ws *WSConn) regist(arg *pb.CRegist, ctx actor.Context) {
	stoc := login.RegistCheck(arg)
	if stoc.Error != pb.OK {
		ws.Send(stoc)
		return
	}
	//重复登录
	if ws.online {
		stoc.Error = pb.RepeatLogin
		ws.Send(stoc)
		return
	}
	msg1 := new(pb.RoleRegist)
	msg1.Phone = arg.GetPhone()
	msg1.Nickname = arg.GetNickname()
	msg1.Password = arg.GetPassword()
	msg1.Smscode = arg.GetSmscode()
	msg1.Safetycode = arg.GetSafetycode()
	//注册
	res1 := ws.reqRole(msg1, ctx)
	var response1 *pb.RoleRegisted
	var ok bool
	if response1, ok = res1.(*pb.RoleRegisted); ok {
		if response1.Error != pb.OK {
			glog.Errorf("CRegist fail %d", response1.Error)
			stoc.Error = response1.Error
			ws.Send(stoc)
			return
		}
	} else {
		glog.Error("CRegist fail")
		stoc.Error = pb.RegistError
		ws.Send(stoc)
		return
	}
	userid := response1.GetUserid()
	glog.Debugf("regist successfully %s", userid)
	if !ws.logining(userid, ctx) {
		glog.Debugf("regist failed %s", ctx.Self())
		stoc.Error = pb.RegistError
		ws.Send(stoc)
		return
	}
	stoc.Userid = userid
	glog.Debugf("regist successfully %s", userid)
	ws.Send(stoc)
	//成功后处理
	ws.logined(userid, true, ctx)
}

func (ws *WSConn) login(arg *pb.CLogin, ctx actor.Context) {
	//检测参数
	stoc := login.LoginCheck(arg)
	if stoc.Error != pb.OK {
		ws.Send(stoc)
		return
	}
	//重复登录
	if ws.online {
		stoc.Error = pb.RepeatLogin
		ws.Send(stoc)
		return
	}
	msg1 := new(pb.RoleLogin)
	msg1.Phone = arg.GetPhone()
	msg1.Password = arg.GetPassword()
	//登录
	res1 := ws.reqRole(msg1, ctx)
	var response1 *pb.RoleLogined
	var ok bool
	if response1, ok = res1.(*pb.RoleLogined); ok {
		if response1.Error != pb.OK {
			glog.Errorf("CLogin fail %d", response1.Error)
			stoc.Error = response1.Error
			ws.Send(stoc)
			return
		}
	} else {
		glog.Error("CLogin fail")
		stoc.Error = pb.LoginError
		ws.Send(stoc)
		return
	}
	userid := response1.GetUserid()
	glog.Debugf("login successfully %s", userid)
	if !ws.logining(userid, ctx) {
		glog.Debugf("login failed %s", ctx.Self())
		stoc.Error = pb.LoginError
		ws.Send(stoc)
		return
	}
	stoc.Userid = userid
	glog.Debugf("login successfully %s", userid)
	ws.Send(stoc)
	//成功后处理
	ws.logined(userid, false, ctx)
}

//微信
func (ws *WSConn) wxlogin(arg *pb.CWxLogin, ctx actor.Context) {
	stoc, wxdata := login.WxLoginCheck(arg)
	if stoc.Error != pb.OK {
		ws.Send(stoc)
		return
	}
	//重复登录
	if ws.online {
		stoc.Error = pb.RepeatLogin
		ws.Send(stoc)
		return
	}
	msg1 := new(pb.WxLogin)
	msg1.Wxuid = wxdata.OpenId
	msg1.Nickname = wxdata.Nickname
	msg1.Photo = wxdata.HeadImagUrl
	msg1.Sex = uint32(wxdata.Sex)
	//登录
	res1 := ws.reqRole(msg1, ctx)
	var response1 *pb.WxLogined
	var ok bool
	if response1, ok = res1.(*pb.WxLogined); ok {
		if response1.Error != pb.OK {
			glog.Errorf("CWxLogin fail %d", response1.Error)
			stoc.Error = response1.Error
			ws.Send(stoc)
			return
		}
	} else {
		glog.Error("CWxLogin fail")
		stoc.Error = pb.GetWechatUserInfoFail
		ws.Send(stoc)
		return
	}
	userid := response1.GetUserid()
	glog.Debugf("weixin login successfully %s", userid)
	if !ws.logining(userid, ctx) {
		stoc.Error = pb.GetWechatUserInfoFail
		ws.Send(stoc)
		return
	}
	stoc.Userid = userid
	glog.Debugf("weixin login successfully %s", userid)
	ws.Send(stoc)
	//成功后处理
	ws.logined(userid, response1.IsRegist, ctx)
}

//登录成功处理
func (ws *WSConn) logined(userid string, isRegist bool,
	ctx actor.Context) {
	//登录成功消息
	msg := new(pb.LoginSuccess)
	msg.IsRegist = isRegist
	msg.Ip = ws.GetIPAddr()
	msg.Userid = userid
	msg.WsPid = ctx.Self()
	//pid已经切换为rsPid
	ws.pid.Tell(msg)
	//登录成功
	ws.online = true
	//成功
	ctx.SetReceiveTimeout(0) //login Successfully, timeout off
}

//登录流程处理
func (ws *WSConn) logining(userid string, ctx actor.Context) bool {
	if userid == "" {
		glog.Debugf("logining failed %s", userid)
		return false
	}
	//当前节点中查询
	if !ws.selectGate(userid, ctx) {
		////节点中不存在新建一个
		//if !ws.loginUser(userid, ctx) {
		//	glog.Debugf("logining loginUser failed %s", userid)
		//	return false
		//}
		return false
	}
	return true
}

/*
//登录成功数据处理
func (ws *WSConn) loginUser(userid string, ctx actor.Context) bool {
	msg4 := new(pb.Login)
	msg4.Userid = userid
	//节点名称
	msg4.Gate = cfg.Section(nodeName).Name()
	msg4.RolePid = ws.pid
	//请求
	res4 := ws.reqRole(msg4, ctx)
	var response4 *pb.Logined
	var ok bool
	if response4, ok = res4.(*pb.Logined); !ok {
		glog.Debugf("loginUser failed %s", userid)
		return false
	}
	//glog.Debugf("response4: %#v", response4)
	msg1 := new(pb.Login2Gate)
	//msg1.WsPid = ws.pid
	msg1.Userid = userid
	msg1.Data = response4.Data //数据
	//在节点中spawn一个玩家进程
	if !ws.loginGate(msg1, ctx) {
		glog.Debugf("loginGate failed %s", userid)
		return false
	}
	return true
}

//登录节点
func (ws *WSConn) loginGate(msg2 *pb.Login2Gate, ctx actor.Context) bool {
	timeout := 3 * time.Second
	res2, err2 := nodePid.RequestFuture(msg2, timeout).Result()
	if err2 != nil {
		glog.Errorf("LoginGate err: %v", err2)
		return false
	}
	glog.Debugf("res2: %#v", res2)
	var response2 *pb.Logined2Gate
	var ok bool
	if response2, ok = res2.(*pb.Logined2Gate); !ok {
		return false
	}
	if response2.Error != pb.OK || response2.Role == nil {
		return false
	}
	glog.Debugf("ws loginGate %s", ws.pid.String())
	glog.Debugf("role loginGate %s", response2.Role.String())
	//登录成功,切换为玩家进程
	ws.pid = response2.Role
	return true
}
*/

//查询节点
func (ws *WSConn) selectGate(userid string, ctx actor.Context) bool {
	msg2 := new(pb.SelectGate)
	//msg2.WsPid = ws.pid
	msg2.Userid = userid
	timeout := 3 * time.Second
	res2, err2 := nodePid.RequestFuture(msg2, timeout).Result()
	if err2 != nil {
		glog.Errorf("selectGate err: %v", err2)
		return false
	}
	glog.Debugf("res2: %#v", res2)
	var response2 *pb.SelectedGate
	var ok bool
	if response2, ok = res2.(*pb.SelectedGate); !ok {
		return false
	}
	if response2.Error != pb.OK || response2.Role == nil {
		return false
	}
	glog.Debugf("ws selectGate %s", ws.pid.String())
	glog.Debugf("role selectGate %s", response2.Role.String())
	//登录成功,切换为玩家进程
	ws.pid = response2.Role
	return true
}

//登录成功数据处理
func (ws *WSConn) reqRole(msg interface{}, ctx actor.Context) interface{} {
	glog.Debugf("reqRole msg %#v", msg)
	if ws.rolePid == nil {
		glog.Errorf("reqRole err %#v", msg)
		ws.pid.Tell(new(pb.ServeStop))
		return nil
	}
	timeout := 3 * time.Second
	res1, err1 := ws.rolePid.RequestFuture(msg, timeout).Result()
	if err1 != nil {
		glog.Errorf("reqRole err: %v, msg %#v", err1, msg)
		return nil
	}
	return res1
}
