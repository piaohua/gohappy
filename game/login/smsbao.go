package login

import (
	"fmt"
	"time"

	"api/sms253"
	"api/smsbao"
	"gohappy/glog"
)

//SendSms 发送验证码
func SendSms(phone, code, username, password string) bool {
	return sendSms2(phone, code, username, password, 0)
}

func sendSms2(phone, code, username, password string, n int) bool {
	content := fmt.Sprintf("【欢乐】你的验证码%s，请勿泄露。", code)
	err := smsbao.SendSmsbao(phone, content, username, password)
	if err == nil {
		glog.Debugf("send sms successfully phone %s, code %s", phone, code)
		return true
	}
	glog.Errorf("send sms failed phone %s, code %s, err %v", phone, code, err)
	if n >= 3 { //失败尝试次数3次
		return false
	}
	<-time.After(3 * time.Second)
	return sendSms2(phone, code, username, password, n+1)
}

//SendSms253 发送验证码
func SendSms253(targetURL, account, password, phone, code string) bool {
	return sendSms3(targetURL, account, password, phone, code, 0)
}

func sendSms3(targetURL, account, password, phone, code string, n int) bool {
	msg := fmt.Sprintf("【欢斗娱乐】你的验证码%s", code) //平台报备签名才能成功,否则不添加签名
	err := sms253.SendSms(targetURL, account, password, phone, msg)
	if err == nil {
		glog.Debugf("send sms successfully phone %s, code %s", phone, code)
		return true
	}
	glog.Errorf("send sms failed phone %s, code %s, err %v", phone, code, err)
	if n >= 3 { //失败尝试次数3次
		return false
	}
	<-time.After(3 * time.Second)
	return sendSms3(targetURL, account, password, phone, code, n+1)
}
