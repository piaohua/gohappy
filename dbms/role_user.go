package main

import (
	"gohappy/data"
	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家请求处理
func (a *RoleActor) handlerUser(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.GetUser:
		arg := msg.(*pb.GetUser)
		a.loginedGetUser(arg, ctx)
	case *pb.SyncUser:
		arg := msg.(*pb.SyncUser)
		a.syncUser(arg, ctx)
	case *pb.ChangeCurrency:
		arg := msg.(*pb.ChangeCurrency)
		//glog.Debugf("ChangeCurrency %#v", arg)
		//更新货币
		a.syncCurrency(arg.Diamond, arg.Coin, arg.Card,
			arg.Chip, arg.Type, arg.Userid)
	case *pb.OfflineCurrency:
		arg := msg.(*pb.OfflineCurrency)
		glog.Debugf("OfflineCurrency %#v", arg)
		a.offlineCurrency(arg)
	case *pb.PayCurrency:
		arg := msg.(*pb.PayCurrency)
		glog.Debugf("PayCurrency %#v", arg)
		//后台或充值同步到game房间
		a.payCurrency(arg)
	case *pb.LoginElse:
		arg := msg.(*pb.LoginElse)
		a.loginElse(arg, ctx) //别处登录
		//响应登录
		rsp := new(pb.LoginedElse)
		rsp.Userid = arg.Userid
		rsp.Gate = arg.Gate
		ctx.Respond(rsp)
	//case *pb.Login:
	//	//登录成功
	//	arg := msg.(*pb.Login)
	//	glog.Debugf("login : %#v", arg)
	//	a.logined(arg, ctx)
	case *pb.Logout:
		//登出成功
		arg := msg.(*pb.Logout)
		a.logouted(arg, ctx)
	case *pb.RoleRegist:
		arg := msg.(*pb.RoleRegist)
		glog.Debugf("RoleRegist %#v", arg)
		a.regist(arg, ctx)
	case *pb.RoleLogin:
		arg := msg.(*pb.RoleLogin)
		glog.Debugf("RoleLogin %#v", arg)
		a.loginByPhone(arg, ctx)
	case *pb.TouristLogin:
		arg := msg.(*pb.TouristLogin)
		glog.Debugf("TouristLogin %#v", arg)
		a.loginByTourist(arg, ctx)
	case *pb.WxLogin:
		arg := msg.(*pb.WxLogin)
		glog.Debugf("WxLogin %#v", arg)
		a.loginByWx(arg, ctx)
	case *pb.GetUserData:
		arg := msg.(*pb.GetUserData)
		user := a.getUserById(arg.Userid)
		rsp := handler.GetUserData(user)
		ctx.Respond(rsp)
	case *pb.ApplePay:
		arg := msg.(*pb.ApplePay)
		rsp := handler.AppleVerify(arg)
		ctx.Respond(rsp)
	case *pb.WxpayCallback:
		arg := msg.(*pb.WxpayCallback)
		a.payHandler(arg)
	case *pb.SmscodeRegist:
		arg := msg.(*pb.SmscodeRegist)
		glog.Debugf("SmscodeRegist %#v", arg)
		a.smsbao(arg, ctx)
	case *pb.CResetPwd:
		//重置密码消息
		arg := msg.(*pb.CResetPwd)
		glog.Debugf("CResetPwd %#v", arg)
		a.resetPwd(arg, ctx)
	//case *pb.GetNumber:
	//	//后台请求
	//	arg := msg.(*pb.GetNumber)
	//	glog.Debugf("GetNumber %#v", arg)
	//	rsp := new(pb.GotNumber)
	//	for k, v := range a.online {
	//		if v.GetRobot() {
	//			rsp.Robot = append(rsp.Robot, k)
	//		} else {
	//			rsp.Role = append(rsp.Role, k)
	//		}
	//	}
	//	ctx.Respond(rsp)
	case *pb.WebRequest:
		arg := msg.(*pb.WebRequest)
		glog.Debugf("WebRequest %#v", arg)
		rsp := new(pb.WebResponse)
		rsp.Code = arg.Code
		a.handlerWeb(arg, rsp, ctx)
		ctx.Respond(rsp)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//定期离线数据清理,移除,存储
func (a *RoleActor) saveUser() {
	glog.Debugf("saveUser caches %#v", a.caches)
	glog.Debugf("saveUser %d, %d", len(a.offline), len(a.online))
	//离线表
	for k, v := range a.offline {
		//TODO 优化缓存策略
		if a.states[k] {
			v.Save()
			delete(a.states, k)
		}
		glog.Debugf("saveUser offline %s, %d", k, v.GetChip())
		if a.caches[k] <= 0 {
			if v.Save() {
				a.delUserMap(v)
				delete(a.caches, k)
				//移除离线表
				delete(a.offline, k)
			} else {
				glog.Errorf("saveUser offline failed %s", k)
			}
		} else {
			a.caches[k]--
		}
	}
	//在线表
	for k, v := range a.online {
		//TODO 优化缓存策略
		if a.states[k] {
			glog.Debugf("saveUser online %s, %d", k, v.GetChip())
			v.Save()
			delete(a.states, k)
		}
	}
}

//立即更新数据库
func (a *RoleActor) saveUserQuickly(userid string) {
	user := a.getUser(userid)
	if user == nil {
		glog.Errorf("saveUser Quickly failed %s", userid)
		return
	}
	user.Save()
}

//在线表中查找,不存在时离线表中获取
func (a *RoleActor) getUser(userid string) *data.User {
	if user, ok := a.online[userid]; ok {
		return user
	}
	if user, ok := a.offline[userid]; ok {
		return user
	}
	return nil
}

//在线表中查找,不存在时离线表中获取,不在离线表从数据库中加载
func (a *RoleActor) getUserById(userid string) *data.User {
	user := a.getUser(userid)
	if user != nil {
		return user
	}
	newUser := new(data.User)
	newUser.GetById(userid) //数据库中取
	if newUser.Userid == "" {
		glog.Debugf("getUserById failed %s", userid)
		return nil
	}
	a.loadingUser(newUser)
	return newUser
}

//在线表中查找
func (a *RoleActor) getUserByTourist(account string) *data.User {
	if v, ok := a.players[account]; ok {
		return a.getUserById(v)
	}
	user := new(data.User)
	user.Tourist = account
	user.GetByTourist() //数据库中取
	if user.Userid == "" {
		glog.Debugf("getUserByTourist failed %s", account)
		return nil
	}
	a.loadingUser(user)
	return user
}

//在线表中查找
func (a *RoleActor) getUserByPhone(account string) *data.User {
	if v, ok := a.players[account]; ok {
		return a.getUserById(v)
	}
	user := new(data.User)
	user.Phone = account
	user.GetByPhone() //数据库中取
	if user.Userid == "" {
		glog.Debugf("getUserByPhone failed %s", account)
		return nil
	}
	a.loadingUser(user)
	return user
}

//在线表中查找
func (a *RoleActor) getUserByWx(account string) *data.User {
	if v, ok := a.players[account]; ok {
		return a.getUserById(v)
	}
	user := new(data.User)
	user.Wxuid = account
	user.GetByWechat() //数据库中取
	if user.GetUserid() == "" {
		glog.Debugf("getUserByWx failed %s", account)
		return nil
	}
	a.loadingUser(user)
	return user
}

//加载
func (a *RoleActor) loadingUser(user *data.User) {
	a.offline[user.GetUserid()] = user
	//映射
	a.setUserMap(user)
	a.caches[user.GetUserid()] = 2 //缓存4分钟
}

//添加映射
func (a *RoleActor) setUserMap(user *data.User) {
	if user.GetWxuid() != "" {
		a.players[user.GetWxuid()] = user.GetUserid()
		glog.Debugf("setUserMap %s = %s", user.GetWxuid(), user.GetUserid())
	} else if user.GetPhone() != "" {
		a.players[user.GetPhone()] = user.GetUserid()
		glog.Debugf("setUserMap %s = %s", user.GetPhone(), user.GetUserid())
	} else if user.GetTourist() != "" {
		a.players[user.GetTourist()] = user.GetUserid()
		glog.Debugf("setUserMap %s = %s", user.GetTourist(), user.GetUserid())
	} else {
		glog.Errorf("user mapping err %s", user.GetUserid())
	}
}

//移除映射
func (a *RoleActor) delUserMap(user *data.User) {
	if user.GetWxuid() != "" {
		delete(a.players, user.GetWxuid())
		glog.Debugf("delUserMap %s", user.GetWxuid())
	} else if user.GetPhone() != "" {
		delete(a.players, user.GetPhone())
		glog.Debugf("delUserMap %s", user.GetPhone())
	} else if user.GetTourist() != "" {
		delete(a.players, user.GetTourist())
		glog.Debugf("delUserMap %s", user.GetTourist())
	} else {
		glog.Errorf("user mapping err %s", user.GetUserid())
	}
}

//离线同步数据
func (a *RoleActor) offlineCurrency(arg *pb.OfflineCurrency) {
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		msg := handler.Offline2Change(arg)
		v.Pid.Tell(msg)
		return
	}
	a.syncCurrency(arg.Diamond, arg.Coin, arg.Card,
		arg.Chip, arg.Type, arg.Userid)
}

//充值同步数据
func (a *RoleActor) payCurrency(arg *pb.PayCurrency) {
	if v, ok := a.roles[arg.Userid]; ok && v != nil {
		v.Pid.Tell(arg)
		return
	}
	a.syncCurrency(arg.Diamond, arg.Coin, arg.Card,
		arg.Chip, arg.Type, arg.Userid)
}

//在线同步数据
func (a *RoleActor) syncUser(arg *pb.SyncUser, ctx actor.Context) {
	glog.Debugf("SyncUser %#v", arg.Userid)
	user := a.getUserById(arg.Userid)
	if user == nil {
		glog.Errorf("syncUser user err %s", arg.Userid)
		return
	}
	err := json.Unmarshal(arg.Data, user)
	if err != nil {
		glog.Errorf("userid %s Unmarshal err %v", arg.Userid, err)
		return
	}
	glog.Debugf("sync user successful %s", arg.Userid)
	a.states[arg.Userid] = true
}

//货币变更
func (a *RoleActor) syncCurrency(diamond, coin, card, chip int64,
	ltype int32, userid string) {
	//日志记录
	user := a.getUserById(userid)
	if user == nil {
		glog.Errorf("syncCurrency err userid %s, type %d, chip %d",
			userid, ltype, chip)
		return
	}
	if chip < 0 && ((chip + user.GetChip()) < 0) {
		chip = 0 - user.GetChip()
	}
	if diamond < 0 && ((diamond + user.GetDiamond()) < 0) {
		diamond = 0 - user.GetDiamond()
	}
	if coin < 0 && ((coin + user.GetCoin()) < 0) {
		coin = 0 - user.GetCoin()
	}
	if card < 0 && ((card + user.GetCard()) < 0) {
		card = 0 - user.GetCard()
	}
	//更新操作
	user.AddCurrency(diamond, coin, card, chip)
	//更新状态
	//a.states[userid] = true
	//暂时实时写入, TODO 异步数据更新
	user.UpdateCurrency()
	//TODO 机器人不写日志
	//if user.GetRobot() {
	//	return
	//}
	//日志记录
	if diamond != 0 {
		msg1 := handler.LogDiamondMsg(diamond, ltype, user)
		loggerPid.Tell(msg1)
	}
	if coin != 0 {
		msg1 := handler.LogCoinMsg(coin, ltype, user)
		loggerPid.Tell(msg1)
	}
	if card != 0 {
		msg1 := handler.LogCardMsg(card, ltype, user)
		loggerPid.Tell(msg1)
	}
	if chip != 0 {
		msg1 := handler.LogChipMsg(chip, ltype, user)
		loggerPid.Tell(msg1)
	}
}
