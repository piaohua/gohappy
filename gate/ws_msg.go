package main

import (
	"gohappy/glog"
	"gohappy/pb"
)

//Router 路由
func (ws *WSConn) Router(id uint32, body []byte) {
	if aesStatus {
		body = aesDe(body) //解密
	}
	msg, err := pb.Unpack(id, body)
	if err != nil {
		glog.Error("protocol unpack err:", id, err)
		return
	}
	ws.pid.Tell(msg)
}

//Send 发送消息
func (ws *WSConn) Send(msg interface{}) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		glog.Errorf("msg %#v, err %#v", msg, err)
	//		glog.Error(string(debug.Stack()))
	//	}
	//}()
	if ws.msgCh == nil {
		glog.Errorf("WSConn msg channel closed %x", msg)
		return
	}
	if len(ws.msgCh) == cap(ws.msgCh) {
		glog.Errorf("send msg channel full -> %d", len(ws.msgCh))
		return
	}
	select {
	case <-ws.stopCh:
		return
	default:
	}
	select {
	case <-ws.stopCh:
		return
	case ws.msgCh <- msg:
		//glog.Debugf("message %#v", msg)
	}
}

//封包
func pack(code, sid uint32, msg []byte) []byte {
	if aesStatus {
		msg = aesEn(msg) //加密
	}
	return append(append(encodeUint32(code), byte(sid)), msg...)
}

func encodeUint32(i uint32) (b []byte) {
	b = append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
	return
}

func decodeUint32(b []byte) (i uint32) {
	i = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return
}
