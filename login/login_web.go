package main

import (
	"fmt"
	"time"

	"gohappy/pb"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// web (protobuf格式请求响应)
func webHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "POST":
	default:
		fmt.Fprintf(ctx, "%s", "failed")
		return
	}
	switch getIP(ctx) {
	case "127.0.0.1":
	default:
		fmt.Fprintf(ctx, "%s", "failed")
		return
	}
	result := ctx.PostBody()
	msg1 := new(pb.WebRequest)
	err1 := msg1.Unmarshal(result)
	if err1 != nil {
		fmt.Fprintf(ctx, "%v", err1)
		return
	}
	timeout := 3 * time.Second
	res2, err2 := nodePid.RequestFuture(msg1, timeout).Result()
	if err2 != nil {
		fmt.Fprintf(ctx, "%v", err2)
		return
	}
	var response2 *pb.WebResponse
	var ok bool
	if response2, ok = res2.(*pb.WebResponse); !ok {
		fmt.Fprintf(ctx, "%s", "failed")
		return
	}
	body, err1 := response2.Marshal()
	if err1 != nil {
		fmt.Fprintf(ctx, "%v", err1)
		return
	}
	fmt.Fprintf(ctx, "%s", body)
}

// web (json格式请求响应)
func webJSONHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "POST":
	default:
		fmt.Fprintf(ctx, "%s", "failed")
		return
	}
	switch getIP(ctx) {
	case "127.0.0.1":
	default:
		//fmt.Fprintf(ctx, "%s", "failed")
		//return
	}
	//解析请求数据
	result := ctx.PostBody()
	msg1 := new(pb.WebRequest)
	err1 := jsoniter.Unmarshal(result, msg1)
	if err1 != nil {
		fmt.Fprintf(ctx, "%v", err1)
		return
	}
	//转换为pb格式,TODO 优化
	switch msg1.Code {
	case pb.WebGive:
		msg3 := new(pb.PayCurrency)
		err3 := jsoniter.Unmarshal(msg1.Data, msg3)
		if err3 != nil {
			fmt.Fprintf(ctx, "%v", err3)
			return
		}
		body, err4 := msg3.Marshal()
		if err4 != nil {
			fmt.Fprintf(ctx, "%v", err4)
			return
		}
		msg1.Data = body
	}
	//请求响应
	timeout := 3 * time.Second
	res2, err2 := nodePid.RequestFuture(msg1, timeout).Result()
	if err2 != nil {
		fmt.Fprintf(ctx, "%v", err2)
		return
	}
	var response2 *pb.WebResponse
	var ok bool
	if response2, ok = res2.(*pb.WebResponse); !ok {
		fmt.Fprintf(ctx, "%s", "failed")
		return
	}
	//result字符串格式响应
	resp2 := new(pb.WebResponse2)
	resp2.Code = response2.Code
	resp2.ErrCode = response2.ErrCode
	resp2.ErrMsg = response2.ErrMsg
	resp2.Result = string(response2.Result)
	body, err1 := jsoniter.Marshal(resp2)
	if err1 != nil {
		fmt.Fprintf(ctx, "%v", err1)
		return
	}
	fmt.Fprintf(ctx, "%s", body)
}
