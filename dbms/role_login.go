package main

import (
	"time"

	"gohappy/data"
	"gohappy/game/login"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//登录后获取数据
func (a *RoleActor) loginedGetUser(arg *pb.GetUser, ctx actor.Context) {
	rsp := new(pb.GotUser)
	user := a.getUserById(arg.Userid)
	if user == nil {
		glog.Errorf("get userid %s fail", arg.Userid)
		ctx.Respond(rsp)
		return
	}
	//打包
	glog.Debugf("loginedGetUser %#v", user)
	result, err1 := json.Marshal(user)
	if err1 != nil {
		glog.Errorf("userid %s Marshal err %v", arg.Userid, err1)
		ctx.Respond(rsp)
		return
	}
	rsp.Data = result
	ctx.Respond(rsp)
	//登录
	a.logined(arg, user, ctx)
}

//登录处理
func (a *RoleActor) logined(arg *pb.GetUser, user *data.User,
	ctx actor.Context) {
	//进程id映射, TODO 玩家rs pid
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		if v.Gate != arg.Gate && v.Pid != nil {
			//关闭旧进程
			glog.Errorf("loginElse arg %#v, v %v", arg, v)
			msg1 := new(pb.LoginElse)
			msg1.Userid = arg.Userid
			msg1.Gate = arg.Gate
			v.Pid.Request(msg1, ctx.Self())
		}
		//替换
		v.Pid = arg.RolePid
		v.Gate = arg.Gate
	} else {
		a.roles[arg.Userid] = &data.Role{
			Pid:  arg.RolePid,
			Gate: arg.Gate,
		}
	}
	//登录成功
	a.online[arg.Userid] = user
	//移除离线表
	delete(a.offline, arg.Userid)
}

//别处登录处理
func (a *RoleActor) loginElse(arg *pb.LoginElse, ctx actor.Context) {
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		if v.Gate != arg.Gate && v.Pid != nil {
			//关闭旧进程
			msg1 := new(pb.LoginElse)
			msg1.Userid = arg.Userid
			msg1.Gate = arg.Gate
			timeout := 3 * time.Second
			res1, err1 := v.Pid.RequestFuture(msg1, timeout).Result()
			if err1 != nil {
				glog.Errorf("loginElse res1 %#v, err1 %v", res1, err1)
			}
		}
	}
}

//登出处理
func (a *RoleActor) logouted(arg *pb.Logout, ctx actor.Context) {
	glog.Debugf("Logout userid: %s", arg.Userid)
	if v, ok := a.online[arg.Userid]; ok {
		//离线
		a.offline[arg.Userid] = v
		//移除
		delete(a.online, arg.Userid)
	}
	delete(a.roles, arg.Userid)
}

//重置密码
func (a *RoleActor) resetPwd(arg *pb.CResetPwd, ctx actor.Context) {
	rsp := new(pb.SResetPwd)
	var smscode string = arg.GetSmscode()
	var phone string = arg.GetPhone()
	var password string = arg.GetPassword()
	errcode := a.findSms(phone, smscode)
	if errcode != pb.OK {
		//TODO 暂时不验证
		//rsp.Error = errcode
		//ctx.Respond(rsp)
		//return
	}
	user := a.getUserByPhone(phone)
	if user == nil {
		rsp.Error = pb.PhoneNotRegist
		ctx.Respond(rsp)
		return
	}
	user.Password = utils.Md5(password + user.Auth)
	rsp.Userid = user.GetUserid()
	ctx.Respond(rsp)
	a.delCode(phone, smscode)
}

//注册处理
func (a *RoleActor) regist(arg *pb.RoleRegist, ctx actor.Context) {
	var smscode string = arg.GetSmscode()
	var phone string = arg.GetPhone()
	var safetycode string = arg.GetSafetycode()
	errcode := a.findSms(phone, smscode)
	if errcode != pb.OK {
		//TODO 暂时不验证
		//rsp := new(pb.RoleRegisted)
		//rsp.Error = errcode
		//ctx.Respond(rsp)
		//return
	}
	//安全码
	if !data.ExistAgency(safetycode) {
		rsp := new(pb.RoleRegisted)
		rsp.Error = pb.SafetycodeNotExist
		ctx.Respond(rsp)
		return
	}
	//在线表中查找,TODO 优化验证前被加载
	if a.getUserByPhone(phone) != nil {
		rsp := new(pb.RoleRegisted)
		rsp.Error = pb.PhoneRegisted
		ctx.Respond(rsp)
		return
	}
	//数据库中查找
	rsp, user := login.Regist(arg, a.uniqueid)
	if rsp.Error == pb.OK {
		a.loadingUser(user)
		//去掉验证码
		a.delCode(phone, smscode)
	}
	ctx.Respond(rsp)
}

//手机登录
func (a *RoleActor) loginByPhone(arg *pb.RoleLogin, ctx actor.Context) {
	var phone string = arg.GetPhone()
	//在线表中查找,TODO 优化验证前被加载
	user := a.getUserByPhone(phone)
	//数据库中查找
	rsp := login.Login(arg, user)
	glog.Debugf("RoleLogin rsp %#v", rsp)
	ctx.Respond(rsp)
}

//微信登录
func (a *RoleActor) loginByWx(arg *pb.WxLogin, ctx actor.Context) {
	var wxuid string = arg.GetWxuid()
	//在线表中查找,TODO 优化验证前被加载
	user := a.getUserByWx(wxuid)
	if user != nil {
		rsp := login.WxLogin(arg, user)
		ctx.Respond(rsp)
	} else {
		rsp, user2 := login.WxRegist(arg, a.uniqueid)
		if rsp.Error == pb.OK {
			a.loadingUser(user2)
		}
		ctx.Respond(rsp)
	}
}

//游客登录
func (a *RoleActor) loginByTourist(arg *pb.TouristLogin, ctx actor.Context) {
	var account string = arg.GetAccount()
	//在线表中查找,TODO 优化验证前被加载
	user := a.getUserByTourist(account)
	if user != nil {
		//数据库中查找
		rsp := login.TouristLogin(arg, user)
		glog.Debugf("TouristLogin rsp %#v", rsp)
		ctx.Respond(rsp)
		return
	}
	//注册ip限制
	if a.tourist[arg.Registip] > 5 {
		glog.Debugf("TouristLogin ip %d", arg.Registip)
		glog.Debugf("TouristLogin ip %d", a.tourist[arg.Registip])
		rsp := new(pb.TouristLogined)
		rsp.Error = pb.RegistError
		ctx.Respond(rsp)
		return
	}
	//数据库中查找
	rsp, user := login.TouristLoginRegist(arg, a.uniqueid)
	if rsp.Error == pb.OK {
		a.loadingUser(user)
		//ip限制
		a.tourist[arg.Registip]++
	}
	ctx.Respond(rsp)
}

//游客注册ip限制
func (a *RoleActor) touristIP() {
	for k, v := range a.tourist {
		if v == 0 {
			delete(a.tourist, k)
			continue
		}
		a.tourist[k] = v - 1
	}
}
