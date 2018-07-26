package login

import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

//RestPwdCheck 重置密码验证
func RestPwdCheck(ctos *pb.CResetPwd) (stoc *pb.SResetPwd) {
	stoc = new(pb.SResetPwd)
	var phone string = ctos.GetPhone()
	var passwd string = ctos.GetPassword()
	var smscode string = ctos.GetSmscode()
	if len(smscode) != 6 {
		stoc.Error = pb.SmsCodeEmpty
		return
	}
	if !utils.PhoneValidate(phone) {
		glog.Errorf("PhoneValidate error %s", phone)
		stoc.Error = pb.PhoneNumberError
		return
	}
	if len(passwd) != 32 {
		stoc.Error = pb.PwdFormatError
	}
	return
}

//RegistCheck 注册验证
func RegistCheck(ctos *pb.CRegist) (stoc *pb.SRegist) {
	stoc = new(pb.SRegist)
	var nickname string = ctos.GetNickname()
	var phone string = ctos.GetPhone()
	var passwd string = ctos.GetPassword()
	var smscode string = ctos.GetSmscode()
	var safetycode string = ctos.GetSafetycode()
	if nickname == "" {
		stoc.Error = pb.UsernameEmpty
		return
	}
	if phone == "" {
		stoc.Error = pb.PhoneNumberEnpty
		return
	}
	if len(safetycode) == 0 {
		stoc.Error = pb.SafetycodeEmpty
		return
	}
	if len(smscode) != 6 {
		//stoc.Error = pb.SmsCodeEmpty
		//TODO 暂时不限制
		//return
	}
	if !utils.LegalName(nickname, 7) {
		glog.Errorf("LegalName error %s", nickname)
		stoc.Error = pb.NameTooLong
		return
	}
	if !utils.PhoneValidate(phone) {
		glog.Errorf("PhoneValidate error %s", phone)
		//stoc.Error = pb.PhoneNumberError
		//TODO 暂时不限制
		//return
	}
	if len(passwd) != 32 {
		stoc.Error = pb.PwdFormatError
	}
	return
}

//Regist 注册处理
func Regist(arg *pb.RoleRegist, genid *data.IDGen) (stoc *pb.RoleRegisted,
	user *data.User) {
	var nickname string = arg.GetNickname()
	var phone string = arg.GetPhone()
	var passwd string = arg.GetPassword()
	var safetycode string = arg.GetSafetycode()
	stoc = new(pb.RoleRegisted)
	user = new(data.User)
	user.Phone = phone
	//if user.ExistsPhone() {
	//	stoc.Error = pb.PhoneRegisted
	//	user = nil
	//	return
	//}
	userid := genid.GenID()
	glog.Debugf("RoleRegist userid %s", userid)
	auth := string(utils.GetAuth())
	user = &data.User{
		Userid:   userid,
		Nickname: nickname,
		Auth:     auth,
		Agent:    safetycode,
		Atime:    utils.BsonNow(),
		Password: utils.Md5(passwd + auth),
		Phone:    phone,
		Ctime:    utils.BsonNow(),
	}
	if !user.Save() {
		glog.Errorf("Regist save error %s", userid)
		stoc.Error = pb.RegistError
		return
	}
	stoc.Userid = user.Userid
	return
}

//RobotRegist Robot注册处理
func RobotRegist(arg *pb.RobotRegist, genid *data.IDGen) {
	user := new(data.User)
	user.Phone = arg.GetPhone()
	user.GetByPhone() //数据库中取
	if user.Userid != "" {
		glog.Debugf("account %s exist", arg.GetPhone())
		return
	}
	userid := genid.GenID()
	glog.Debugf("RobotRegist userid %s", userid)
	user.Userid = userid
	user.Nickname = arg.GetNickname()
	user.Phone = arg.GetPhone()
	user.Photo = arg.GetPhoto()
	user.Password = utils.Md5(arg.GetPassword() + arg.GetAuth())
	user.Auth = arg.GetAuth()
	user.Sex = arg.GetSex()
	user.Vip = arg.GetVip()
	user.Coin = arg.GetCoin()
	user.Diamond = arg.GetDiamond()
	user.Ctime = utils.BsonNow()
	user.Robot = true
	if !user.Save() {
		glog.Errorf("RobotRegist save failed %#v, userid %s", arg, userid)
		return
	}
	glog.Debugf("RobotRegist successfully userid %s, phone %s", userid, arg.GetPhone())
}