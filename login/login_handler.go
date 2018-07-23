package main

import (
	"fmt"

	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/valyala/fasthttp"
)

//Start 启动监听服务
func Start(addr string) {
	if err = fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		glog.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func getIP(ctx *fasthttp.RequestCtx) (ip string) {
	ip = string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if ip == "" {
		ip = ctx.RemoteIP().String()
	}
	return
}

func fooHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())
	fmt.Fprintf(ctx, "header X-Forwarded-For %q\n\n", ctx.Request.Header.Peek("X-Forwarded-For"))
}

func barHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
}

func logHandler(ctx *fasthttp.RequestCtx) {
	glog.Debugf("Request method is %q\n", ctx.Method())
	glog.Debugf("RequestURI is %q\n", ctx.RequestURI())
	glog.Debugf("Requested path is %q\n", ctx.Path())
	glog.Debugf("Host is %q\n", ctx.Host())
	glog.Debugf("Query string is %q\n", ctx.QueryArgs())
	glog.Debugf("User-Agent is %q\n", ctx.UserAgent())
	glog.Debugf("Connection has been established at %s\n", ctx.ConnTime())
	glog.Debugf("Request has been started at %s\n", ctx.Time())
	glog.Debugf("Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	glog.Debugf("Your ip is %q\n\n", ctx.RemoteIP())
	glog.Debugf("Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
}

// 短信验证码
func smsHandler(ctx *fasthttp.RequestCtx) {
	phone := string(ctx.QueryArgs().Peek("phone"))
	glog.Debugf("phone %s", phone)
	if !utils.PhoneValidate(phone) {
		//fmt.Fprintf(ctx, "%s", "1")
		return
	}
	//生成
	arg := new(pb.SmscodeRegist)
	arg.Phone = phone
	arg.Type = 1
	arg.Ipaddr = getIP(ctx)
	nodePid.Tell(arg)
	//fmt.Fprintf(ctx, "%s", "0")
}

//TODO version rule, json格式返回,包含状态或规则
func gateHandler(ctx *fasthttp.RequestCtx) {
	r := "gate.node"
	r += string(ctx.QueryArgs().Peek("version"))
	glog.Debugf("gate node : %s", r)
	sec1, err1 := cfg.GetSection(r)
	if err1 != nil {
		glog.Error("Unknwon version ", err1)
		ctx.Error("Unknwon version", fasthttp.StatusBadRequest)
		return
	}
	key2, err2 := sec1.GetKey("host")
	if err2 != nil {
		glog.Error("Unknwon version ", err2)
		ctx.Error("Unknwon version", fasthttp.StatusBadRequest)
		return
	}
	host := key2.Value()
	logHandler(ctx)
	glog.Debugf("gate host %s", host)
	//fmt.Fprintf(ctx, "%s", string(aesEn(host)))
	fmt.Fprintf(ctx, "%s", string(host))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	defer ctx.SetConnectionClose()
	switch string(ctx.Path()) {
	case "/happy/web":
		webHandler(ctx)
	case "/happy/webjson":
		webJSONHandler(ctx)
	case "/happy/sms":
		smsHandler(ctx)
	case "/happy/gate":
		gateHandler(ctx)
	case "/happy/wxpay/notice":
		wxpayHandler(ctx)
	case "/happy/jtpay/order":
		jtpayOrder(ctx)
	case "/happy/jtpay/return":
		jtpayReturn(ctx)
	case "/happy/jtpay/notify":
		jtpayNotify(ctx)
	case "/happy/wxmp/oauth2":
		wxmpOauth2(ctx)
	case "/happy/wxmp/shorturl":
		wxmpShortURL(ctx)
	case "/happy/wxmp/qrcode":
		wxmpQRcode(ctx)
	case "/happy/download":
		//download(ctx)
        downloadPageHandler(ctx)
	//case "/wxmp/wx":
	case "/happy/foo":
		fooHandler(ctx)
	case "/happy/bar":
		barHandler(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}
