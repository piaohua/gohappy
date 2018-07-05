package main

import (
	"fmt"
	"net/http"

	"gohappy/glog"

	"github.com/valyala/fasthttp"
	"github.com/wizjin/weixin"
)

//Echo 文本消息的处理函数
func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content // 获取用户发送的消息
	glog.Debugf("echo txt %s", txt)
	w.ReplyText(txt)          // 回复一条文本消息
	w.PostText("Post:" + txt) // 发送一条文本消息
	//w.ReplyOK()
}

//Subscribe 关注事件的处理函数
func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	glog.Debugf("Subscribe txt %s", r.Content)
	w.ReplyText("欢迎关注") // 有新人关注，返回欢迎消息
}

var mux *weixin.Weixin

//Wxmp 微信公众号
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
	http.Handle(pattern, mux) // 注册接收微信服务器数据的接口URI
	//mux.CreateHandlerFunc()
	//TODO fasthttp
	http.ListenAndServe(":6210", nil) // 启动接收微信数据服务器
}

// 授权回调处理函数, https://xxx.com/happy/wxmp/oauth2?code=xxx&state=103133
func wxmpOauth2(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-type", "text/html;charset=UTF-8")
	glog.Debugf("RequestURI is %q\n", ctx.RequestURI())
	code := string(ctx.QueryArgs().Peek("code"))
	state := string(ctx.QueryArgs().Peek("state"))
	glog.Debugf("code: %s, state: %s", code, state)
	userAccessToken, err := mux.GetUserAccessToken(code)
	glog.Debugf("userAccessToken %#v, err %v", userAccessToken, err)
	glog.Debugf("openid %s", userAccessToken.OpenId)
	userInfo, err := mux.GetUserInfo(userAccessToken.OpenId)
	glog.Debugf("userInfo %#v, err %v", userInfo, err)
	//TODO 存储关系
	//TODO 展示页面
	fmt.Fprintf(ctx, "%s", "success")
}

// 获取链接地址, https://xxx.com/happy/wxmp/shorturl?userid=103133
func wxmpShortURL(ctx *fasthttp.RequestCtx) {
	glog.Debugf("RequestURI is %q\n", ctx.RequestURI())
	state := string(ctx.QueryArgs().Peek("userid"))
	if state == "" {
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	//TODO 确认代理
	//urlStr string, scope string, state string
	urlStr := cfg.Section("wxmp").Key("redirect_uri").Value()
	scope := cfg.Section("wxmp").Key("scope").Value()
	redirectURL := mux.CreateRedirectURL(urlStr, scope, state)
	glog.Debugf("redirectURL %s\n", redirectURL)
	shortURL, err := mux.ShortURL(redirectURL)
	glog.Debugf("shortURL %s, err %v\n", shortURL, err)
	if err != nil {
		glog.Error("shortURL err ", err)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	fmt.Fprintf(ctx, "%s", shortURL)
}

// 获取二维码, https://xxx.com/happy/wxmp/qrcode?userid=103133
func wxmpQRcode(ctx *fasthttp.RequestCtx) {
	//qr, err := mux.CreateQRLimitSceneByString(userid)
	//qrURL := qr.ToURL()
	fmt.Fprintf(ctx, "%s", "success")
}