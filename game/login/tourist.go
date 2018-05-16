package login

import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

//TouristLoginCheck 登录验证
func TouristLoginCheck(ctos *pb.CTourist, key string) (stoc *pb.STourist) {
	stoc = new(pb.STourist)
	var account string = ctos.GetAccount()
	var passwd string = ctos.GetPassword()
	if len(passwd) != 32 {
		stoc.Error = pb.PwdFormatError
		return
	}
	//5秒失效
	n := utils.Timestamp() - (ctos.GetTime() / 1000)
	if n > 5 {
		glog.Debugf("tourist n %d, %d", n, ctos.GetTime())
		stoc.Error = pb.LoginError
		return
	}
	//解密验证
	account, err1 := TouristAccount(account, key)
	glog.Debugf("account %s, err1 %v", account, err1)
	if err1 != nil {
		stoc.Error = pb.LoginError
		return
	}
	if account == "" {
		stoc.Error = pb.LoginError
		return
	}
	//账号简单规则验证,TODO 优化
	str := utils.Split(account, "_")
	if len(str) != 2 {
		stoc.Error = pb.LoginError
		return
	}
	ran := utils.Split(str[1], ".")
	if len(ran) != 2 {
		stoc.Error = pb.LoginError
		return
	}
	//密码验证
	if passwd != utils.Md5(account) {
		stoc.Error = pb.LoginError
		return
	}
	ctos.Account = account
	return
}

//TouristLogin 游客登录
func TouristLogin(ctos *pb.TouristLogin, user *data.User) (stoc *pb.TouristLogined) {
	stoc = new(pb.TouristLogined)
	if user == nil {
		stoc.Error = pb.UsernameOrPwdError
		return
	}
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

//TouristLoginRegist 游客注册
func TouristLoginRegist(arg *pb.TouristLogin, genid *data.IDGen) (stoc *pb.TouristLogined,
	user *data.User) {
	var account string = arg.GetAccount()
	var passwd string = arg.GetPassword()
	stoc = new(pb.TouristLogined)
	user = new(data.User)
	user.Tourist = account
	//if user.ExistsTourist() {
	//	stoc.Error = pb.PhoneRegisted
	//	user = nil
	//	return
	//}
	userid := genid.GenID()
	nickname := "游客" + userid
	glog.Debugf("TouristLogin userid %s", userid)
	auth := string(utils.GetAuth())
	user = &data.User{
		Userid:   userid,
		Nickname: nickname,
		Auth:     auth,
		Password: utils.Md5(passwd + auth),
		Tourist:  account,
		Ctime:    utils.BsonNow(),
	}
	if !user.Save() {
		glog.Errorf("Regist save error %s", userid)
		stoc.Error = pb.RegistError
		return
	}
	stoc.Userid = user.Userid
	stoc.IsRegist = true
	return
}
