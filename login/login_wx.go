package main

import (
	"fmt"
	"net/http"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/valyala/fasthttp"
	"github.com/wizjin/weixin"
	"github.com/json-iterator/go"
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
	glog.Debugf("RequestURI is %q\n", ctx.RequestURI())
	code := string(ctx.QueryArgs().Peek("code"))
	state := string(ctx.QueryArgs().Peek("state"))
	if code == "" || state == "" {
		glog.Errorf("wxmpOauth2 failed: %s", ctx.RequestURI())
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	glog.Debugf("code: %s, state: %s", code, state)
	userAccessToken, err := mux.GetUserAccessToken(code)
	glog.Debugf("userAccessToken %#v, err %v", userAccessToken, err)
	if err != nil {
		glog.Errorf("GetUserAccessToken err: %v", err)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	if userAccessToken.OpenId == "" {
		glog.Errorf("GetUserAccessToken failed: %#v", userAccessToken)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	glog.Debugf("openid %s", userAccessToken.OpenId)
	userInfo, err := mux.GetSnsUserInfo(userAccessToken.OpenId, userAccessToken.AccessToken)
	glog.Debugf("userInfo %#v, err %v", userInfo, err)
	if err != nil {
		glog.Errorf("GetUserInfo err: %v", err)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	if userInfo.UnionId == "" {
		glog.Errorf("GetUserInfo failed: %#v", userInfo)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	//存储关系
	result, err := jsoniter.Marshal(userInfo)
	if err != nil {
		glog.Errorf("userinfo marshal err: %v", err)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	msg := new(pb.AgentOauth2Confirm)
	msg.Agentid = state
	msg.Userinfo = result
	res1, err1 := callNode(msg)
	if err1 != nil {
		glog.Errorf("shortURL err: %v", err1)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	if response1, ok := res1.(*pb.AgentOauth2Confirmed); ok {
		if response1.Error != pb.OK {
			glog.Errorf("Oauth2 err: %#v", response1)
			ctx.Error("failure", fasthttp.StatusBadRequest)
			return
		}
	} else {
		glog.Errorf("Oauth2 res1 err: %#v", res1)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	//重定向微信公众号设置的业务域名，防止被重新排版
	downloadURL := cfg.Section("domain").Key("download").Value()
	ctx.Redirect(downloadURL, fasthttp.StatusMovedPermanently)
}

// 获取链接地址, https://xxx.com/happy/wxmp/shorturl?userid=103133
func wxmpShortURL(ctx *fasthttp.RequestCtx) {
	glog.Debugf("RequestURI is %q\n", ctx.RequestURI())
	state := string(ctx.QueryArgs().Peek("userid"))
	if state == "" {
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	//确认代理
	msg := new(pb.AgentConfirm)
	msg.Userid = state
	res1, err1 := callNode(msg)
	if err1 != nil {
		glog.Errorf("shortURL err: %v", err1)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
	if response1, ok := res1.(*pb.AgentConfirmed); ok {
		if response1.Error != pb.OK {
			glog.Errorf("shortURL err: %#v", response1)
			ctx.Error("failure", fasthttp.StatusBadRequest)
			return
		}
	} else {
		glog.Errorf("shortURL res1 err: %#v", res1)
		ctx.Error("failure", fasthttp.StatusBadRequest)
		return
	}
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

// 下载页面, https://xxx.com/happy/download
func download(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-type", "text/html;charset=UTF-8")
	downloadHtml := downloadHtml()
	fmt.Fprintf(ctx, "%s", downloadHtml)
}

//TODO use html/template
func downloadHtml() (str string) {
	str = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1.0,maximum-scale=1.0,user-scalable=0">
<meta http-equiv="Pragma" content="no-cache">
<meta http-equiv="Cache-Control" content="no-cache">
<meta http-equiv="Expires" content="0">
<base href="/">
<title></title>
</head>
<body></body>
</html>
`
	return
}