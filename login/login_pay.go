package main

import (
	"fmt"
	"strings"

	"api/wxpay"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/valyala/fasthttp"
)

// 接收交易结果通知
func wxpayHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-type", "text/plain;charset=UTF-8")
	switch string(ctx.Method()) {
	case "POST":
	default:
		glog.Error("wxpay method err")
		//TODO failed
		fmt.Fprintf(ctx, wxpay.TradeRespXml())
		return
	}
	result := ctx.PostBody()
	glog.Debugf("result %s", string(result))
	//解析
	tradeResult, err := wxpay.ParseTradeResult(result)
	glog.Debugf("tradeResult %#v, err %v", tradeResult, err)
	if err != nil {
		glog.Errorf("trade result err: %v", err)
		//TODO failed
		fmt.Fprintf(ctx, wxpay.TradeRespXml())
		return
	}
	//发货
	msg := new(pb.WxpayCallback)
	msg.Result = string(result)
	nodePid.Tell(msg)
	fmt.Fprintf(ctx, wxpay.TradeRespXml())
}

// 接收交易结果通知
func jtpayHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-type", "text/plain;charset=UTF-8")
	switch string(ctx.Method()) {
	case "POST":
	default:
		//fmt.Fprintf(ctx, "%s", "failed")
		//return
	}
	ip := getIP(ctx)
	ipwhitelist := cfg.Section("jtpay").Key("ipwhitelist").Value()
	if !strings.Contains(ipwhitelist, ip) {
		//fmt.Fprintf(ctx, "%s", "failed")
		//return
	}
	glog.Errorf("ipwhitelist %v", ipwhitelist)
	glog.Errorf("ip %s, Method %s", ip, string(ctx.Method()))
	result := ctx.PostBody()
	glog.Debugf("result %s", string(result))
	glog.Errorf("result %s", string(result))
	//解析
	//tradeResult, err := wxpay.ParseTradeResult(result)
	//glog.Debugf("tradeResult %#v, err %v", tradeResult, err)
	//if err != nil {
	//	glog.Errorf("trade result err: %v", err)
	//	fmt.Fprintf(ctx, "%s", "failed")
	//	return
	//}
	//TODO 直接role服务中验证
	//发货
	//msg := new(pb.WxpayCallback)
	//msg.Result = string(result)
	//nodePid.Tell(msg)
	fmt.Fprintf(ctx, "%s", "success")
}
