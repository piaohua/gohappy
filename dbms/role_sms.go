package main

import (
	"gohappy/game/login"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//验证验证码
func (a *RoleActor) findSms(phone, smscode string) pb.ErrCode {
	//TODO 暂时不限制
	//return pb.OK

	if phone != a.smscode[smscode] {
		return pb.SmsCodeWrong
	}
	//验证码过期
	if a.smstime[smscode] <= utils.Timestamp() {
		return pb.SmsCodeExpired
	}
	return pb.OK
}

//去掉验证码
func (a *RoleActor) delCode(phone, code string) {
	if v, ok := a.smscode[code]; ok {
		delete(a.smsphone, v)
	}
	delete(a.smstime, code)
	delete(a.smscode, code)
	delete(a.smsphone, phone)
}

//短信验证码
func (a *RoleActor) smsbao(arg *pb.SmscodeRegist, ctx actor.Context) {
	switch arg.Type {
	case 1: //生成
		if v, ok := a.smsphone[arg.Phone]; ok {
			glog.Errorf("phone %s code %s already exist", arg.Phone, v)
		} else {
			code := a.genCode()
			a.smscode[code] = arg.Phone
			a.smstime[code] = utils.Timestamp() + (60 * 3)
			a.smsphone[arg.Phone] = code
			glog.Debugf("phone %s, code %s", arg.Phone, code)
			//if cfg.Section("smsbao").Key("status").MustBool(false) {
			//	go login.SendSms(arg.Phone, code, smsusername, smspassword)
			//}
			//sms253
			if cfg.Section("sms253").Key("status").MustBool(false) {
				go login.SendSms253(sms253URL, sms253account, sms253password, arg.Phone, code)
			}
		}
	case 2: //删除
		a.delCode(arg.Phone, arg.Smscode)
	case 3: //查询
		if _, ok := a.smsphone[arg.Phone]; ok {
		}
	default:
	}
}

//过期检测
func (a *RoleActor) smsExpire() {
	now := utils.Timestamp()
	for k, v := range a.smsphone {
		if a.smstime[v] > now {
			continue
		}
		a.delCode(k, v)
	}
}

//生成一个验证码,唯一
func (a *RoleActor) genCode() (s string) {
	s = utils.RandStr(6)
	//是否已经存在
	if _, ok := a.smscode[s]; ok {
		return a.genCode() //重复尝试,TODO:一定次数后放弃尝试
	}
	return
}
