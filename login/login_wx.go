package main

import (
	"net/http"

	"gohappy/glog"

	"github.com/wizjin/weixin"
)

// 文本消息的处理函数
func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content // 获取用户发送的消息
	glog.Debugf("echo txt %s", txt)
	w.ReplyText(txt)          // 回复一条文本消息
	w.PostText("Post:" + txt) // 发送一条文本消息
	//w.ReplyOK()
}

// 关注事件的处理函数
func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	glog.Debugf("Subscribe txt %s", r.Content)
	w.ReplyText("欢迎关注") // 有新人关注，返回欢迎消息
}

// 授权的处理函数
func Download(w weixin.ResponseWriter, r *weixin.Request) {
	glog.Debugf("Content %s", r.Content)
	glog.Debugf("url %s", r.Url)
	wx := w.GetWeixin()
	code := ""
	userAccessToken, err := wx.GetUserAccessToken(code)
	glog.Debugf("userAccessToken %#v, err %v", userAccessToken, err)
	userInfo, err := wx.GetUserInfo(userAccessToken.OpenId)
	glog.Debugf("userInfo %#v, err %v", userInfo, err)
	w.ReplyOK()
}

var mux *weixin.Weixin

func Wxmp() {
	token := cfg.Section("wxmp").Key("token").Value()
	appid := cfg.Section("wxmp").Key("appid").Value()
	secret := cfg.Section("wxmp").Key("appsecret").Value()
	key := cfg.Section("wxmp").Key("encoding-AES-key").Value()
	pattern := cfg.Section("wxmp").Key("pattern").Value()
	// my-token 验证微信公众平台的Token
	// app-id, app-secret用于高级API调用。
	// 如果仅使用接收/回复消息，则可以不填写，使用下面语句
	// mux := weixin.New("my-token", "", "")
	mux = weixin.New(token, appid, secret)
	// 设置AES密钥，如果不使用AES可以省略这行代码
	mux.SetEncodingAESKey(key)
	// 注册文本消息的处理函数
	mux.HandleFunc(weixin.MsgTypeText, Echo)
	// 注册关注事件的处理函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, Subscribe)
	glog.Debug("wxmp start...")
	//mux.RefreshAccessToken()
	//urlStr string, scope string, state string
	//urlStr := "https://bbbbvb.cn/wxmp/oauth2"
	//scope := "snsapi_userinfo"
	//state := "103133"
	//redirectURL := mux.CreateRedirectURL(urlStr, scope, state)
	//glog.Debugf("redirectURL %s\n", redirectURL)
	//shortURL, err := mux.ShortURL(redirectURL)
	//glog.Debugf("shortURL %s, err %v\n", shortURL, err)
	http.Handle(pattern, mux) // 注册接收微信服务器数据的接口URI
	//mux.CreateHandlerFunc()
	//TODO fasthttp
	http.ListenAndServe(":6210", nil) // 启动接收微信数据服务器
}
