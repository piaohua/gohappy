package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家数据请求处理
func (a *GateActor) handlerLogin(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.LoginElse:
		//在别的节点登录
		arg := msg.(*pb.LoginElse)
		glog.Debugf("LoginElse %#v", arg)
		a.loginElse(arg, ctx)
	case *pb.LoginedElse:
		//在别的节点登录
		arg := msg.(*pb.LoginedElse)
		glog.Debugf("LoginedElse %#v", arg)
	case *pb.SetLogin:
		arg := msg.(*pb.SetLogin)
		glog.Debugf("SetLogin %#v", arg)
		rsp := &pb.SetLogined{
			RolePid: a.rolePid,
		}
		glog.Infof("SetLogin %s", arg.Sender.String())
		ctx.Respond(rsp)
	case *pb.Logout:
		//登出成功
		arg := msg.(*pb.Logout)
		glog.Debugf("Logout %#v", arg)
		a.logout(arg, ctx)
	case *pb.SelectGate:
		//登录成功
		arg := msg.(*pb.SelectGate)
		glog.Debugf("SelectGate %#v", arg)
		a.selectRole(arg, ctx)
	//case *pb.Login2Gate:
	//	//登录成功
	//	arg := msg.(*pb.Login2Gate)
	//	glog.Debugf("Login2Gate %#v", arg)
	//	a.spawnRole(arg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//在别的节点登录
func (a *GateActor) loginElse(arg *pb.LoginElse, ctx actor.Context) {
	userid := arg.GetUserid()
	if p, ok := a.online[userid]; ok {
		p.Tell(arg)
		//直接关闭
		p.Tell(new(pb.OfflineStop))
	} else if p, ok := a.offline[userid]; ok {
		p.Tell(arg)
		//直接关闭
		p.Tell(new(pb.OfflineStop))
	}
	glog.Debugf("LoginElse userid: %s", userid)
	//移除
	delete(a.online, userid)
	delete(a.offline, userid)
	delete(a.offtime, userid)
}

//登出成功
func (a *GateActor) logout(arg *pb.Logout, ctx actor.Context) {
	userid := arg.GetUserid()
	glog.Debugf("Logout userid: %s", userid)
	//离线 arg.Sender == a.online[userid]
	a.offline[userid] = arg.Sender
	a.offtime[userid] = 10 //缓存5分钟
	//移除
	delete(a.online, userid)
	//不能离开,因为缓存了数据,离开会出现数据不同步
	//a.rolePid.Tell(arg)
	//a.roomPid.Tell(arg)
}

//下线离线玩家
func (a *GateActor) offlineStop(ctx actor.Context) {
	for k, v := range a.offline {
		if a.offtime[k] <= 0 {
			v.Tell(new(pb.OfflineStop))
			delete(a.offline, k)
			delete(a.offtime, k)
			//正式下线消息
			arg := new(pb.Logout)
			arg.Userid = k
			arg.Sender = v
			a.rolePid.Tell(arg)
			a.roomPid.Tell(arg)
			continue
		}
		a.offtime[k]--
	}
}

////断开其它节点连接
//func (a *GateActor) logoutOther(userid string, ctx actor.Context) {
//	//TODO 数据一致性,防止数据覆盖
//	msg1 := new(pb.LoginHall)
//	msg1.Userid = userid
//	msg1.NodeName = a.Name
//	a.rolePid.Tell(msg1)
//}

//登录成功查询
func (a *GateActor) selectRole(arg *pb.SelectGate, ctx actor.Context) {
	userid := arg.GetUserid()
	glog.Debugf("SelectGate userid: %s", userid)
	//断开其它节点连接
	//a.logoutOther(userid, ctx)
	//在线表查询
	if p, ok := a.online[userid]; ok {
		//断开当前节点旧连接
		msg1 := new(pb.LoginElse)
		msg1.Userid = userid
		msg1.Gate = a.Name
		p.Request(msg1, ctx.Self())
		glog.Debugf("SelectGate online userid: %s, %s", userid, p.String())
		//响应登录
		rsp := new(pb.SelectedGate)
		rsp.Role = p
		ctx.Respond(rsp)
		return
	}
	//离线表查找
	if p, ok := a.offline[userid]; ok {
		//切换到在线表
		a.online[userid] = p
		delete(a.offline, userid)
		delete(a.offtime, userid)
		glog.Debugf("SelectGate offline userid: %s, %s", userid, p.String())
		//响应登录
		rsp := new(pb.SelectedGate)
		rsp.Role = p
		ctx.Respond(rsp)
		return
	}
	//新玩家
	rolePid := a.spawnRole(userid)
	//添加
	a.online[userid] = rolePid
	//响应登录,不存在
	rsp := new(pb.SelectedGate)
	rsp.Role = rolePid
	ctx.Respond(rsp)
}

//新玩家
func (a *GateActor) spawnRole(userid string) *actor.PID {
	newRole := NewRole()
	newRole.dbmsPid = a.dbmsPid
	newRole.roomPid = a.roomPid
	newRole.rolePid = a.rolePid
	newRole.loggerPid = a.loggerPid
	rolePid := newRole.initRs()
	newRole.pid = rolePid
	msg1 := &pb.ServeStart{
		Message: userid,
	}
	rolePid.Tell(msg1)
	return rolePid
}
