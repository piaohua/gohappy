package login

import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

//LoginCheck 登录验证
func LoginCheck(ctos *pb.CLogin) (stoc *pb.SLogin) {
	stoc = new(pb.SLogin)
	var phone string = ctos.GetPhone()
	var passwd string = ctos.GetPassword()
	if phone == "" {
		stoc.Error = pb.PhoneNumberEnpty
		return
	}
	if !utils.PhoneRegexp(phone) {
		glog.Errorf("PhoneRegexp error %s, %d", phone, len(phone))
		//stoc.Error = pb.PhoneNumberError
		//TODO 暂时不限制
		//return
	}
	if len(passwd) != 32 {
		stoc.Error = pb.PwdFormatError
	}
	return
}

//Login 登录
func Login(ctos *pb.RoleLogin, user *data.User) (stoc *pb.RoleLogined) {
	stoc = new(pb.RoleLogined)
	if user == nil {
		stoc.Error = pb.UsernameOrPwdError
		return
	}
	//var phone string = ctos.GetPhone()
	var passwd string = ctos.GetPassword()
	if !user.VerifyPwd(passwd) {
		glog.Errorf("Login error %s", user.GetUserid())
		stoc.Error = pb.UsernameOrPwdError
	}
	if user.Userid == "" {
		stoc.Error = pb.LoginError
	}
	if stoc.Error != pb.OK {
		return
	}
	stoc.Userid = user.Userid
	return
}
