/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-11-19 13:11:36
 * Filename      : client.go
 * Description   : 机器人
 * *******************************************************/
package main

import (
	"errors"
	"time"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 1024                // Maximum message size allowed from peer.
	waitForLogin   = 20 * time.Second    // 连接建立后5秒内没有收到登陆请求,断开socket
)

//WebsocketConnSet 连接集合
type WebsocketConnSet map[*websocket.Conn]struct{}

//Robot 机器人连接数据
type Robot struct {
	conn *websocket.Conn // websocket连接

	stopCh chan struct{}    // 关闭通道
	msgCh  chan interface{} // 消息通道

	maxMsgLen uint32 // 最大消息长度

	//玩家游戏数据
	data *user //数据
	//游戏桌子数据
	code   string //邀请码
	gtype  int32  //game type
	rtype  int32  //room type
	dtype  int32  //desk type
	ltype  int32  //room type
	roomid string //房间id
	envBet int32  //下注规则
	seat   uint32 //位置
	//临时数据
	round  uint32 //玩牌局数
	sits   uint32 //尝试坐下次数
	bits   uint32 //尝试下注次数
	bitNum uint32 //尝试下注数量
	regist bool   //注册标识
	timer  uint32 //在线时间
}

// 基本数据
type user struct {
	Userid   string // 用户id
	Nickname string // 用户昵称
	Sex      uint32 // 用户性别,男1 女2 非男非女3
	Phone    string // 绑定的手机号码
	Coin     int64  // 金币
	Diamond  int64  // 钻石
	Card     int64  // 房卡
	Chip     int64  // 筹码
	Vip      uint32 // vip
}

//WSPING 心跳
type WSPING int

//通道关闭信号
type closeFlag int

//创建连接
func newRobot(conn *websocket.Conn, pendingWriteNum int, maxMsgLen uint32) *Robot {
	return &Robot{
		maxMsgLen: maxMsgLen,

		conn: conn,
		data: new(user),

		msgCh:  make(chan interface{}, pendingWriteNum),
		stopCh: make(chan struct{}),
	}
}

//Close 断开连接
func (ws *Robot) Close() {
	select {
	case <-ws.stopCh:
		return
	default:
		//关闭消息通道
		ws.Sender(closeFlag(1))
		//关闭连接
		ws.conn.Close()
		//Logout message
		Logout(ws.roomid, ws.data.Phone, ws.code, ws.data.Chip)
	}
}

//Router 接收消息路由
func (ws *Robot) Router(id uint32, body []byte) {
	//body = pbAesDe(body) //解密
	msg, err := pb.Runpack(id, body)
	if err != nil {
		glog.Error("protocol unpack err:", id, err)
		return
	}
	ws.receive(msg)
}

//Sender 发送消息
func (ws *Robot) Sender(msg interface{}) {
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
		glog.Info("ticker closed")
		return
	default: //防止阻塞
	}
	select {
	case <-ws.stopCh:
		return
	case ws.msgCh <- msg:
	}
}

//时钟
func (ws *Robot) ticker() {
	tick := time.Tick(5 * time.Second)
	for {
		select {
		case <-ws.stopCh:
			glog.Info("ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-ws.stopCh:
			glog.Info("ticker closed")
			return
		case <-tick:
			ws.SendPing()
		}
	}
}

func (ws *Robot) readPump() {
	defer ws.Close()
	ws.conn.SetReadLimit(maxMessageSize)
	ws.conn.SetReadDeadline(time.Now().Add(pongWait))
	ws.conn.SetPongHandler(func(string) error { ws.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		n, message, err := ws.conn.ReadMessage()
		if err != nil {
			glog.Errorf("Read error: %s, %d\n", err, n)
			break
		}
		//if len(message) > 1024 {
		//	glog.Errorf("message too long -> %d", len(message))
		//	return
		//}
		if len(message) < 5 {
			glog.Errorf("message error: %#v\n", message)
			return
		}
		//路由
		ws.Router(decodeUint32(message[:4]), message[5:])
	}
}

//消息写入
func (ws *Robot) writePump() {
	for {
		select {
		case message, ok := <-ws.msgCh:
			if !ok {
				ws.write(websocket.CloseMessage, []byte{})
				return
			}
			err := ws.write(websocket.BinaryMessage, message)
			if err != nil {
				//停止发送消息
				close(ws.stopCh)
				return
			}
		}
	}
}

//Send pings
func (ws *Robot) pingPump() {
	tick := time.Tick(pingPeriod)
	for {
		select {
		case <-tick:
			ws.Sender(WSPING(1))
		case <-ws.stopCh:
			return
		}
	}
}

//写入
func (ws *Robot) write(mt int, msg interface{}) error {
	var message []byte
	switch msg.(type) {
	case closeFlag:
		return errors.New("msg channel closed")
	case WSPING:
		mt = websocket.PingMessage
	case []byte:
		message = msg.([]byte)
	default:
		code, body, err := pb.Rpacket(msg)
		if err != nil {
			glog.Errorf("write msg err %v", msg)
			return err
		}
		message = pack(code, body)
	}
	if uint32(len(message)) > ws.maxMsgLen {
		glog.Errorf("write msg too long -> %d", len(message))
		return errors.New("write msg too long")
	}
	ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return ws.conn.WriteMessage(mt, message)
}

func decodeUint32(b []byte) (i uint32) {
	i = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	return
}

func encodeUint32(i uint32) (b []byte) {
	b = append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
	return
}

//封包
func pack(code uint32, msg []byte) []byte {
	//msg = pbAesEn(msg) //加密
	return append(encodeUint32(code), msg...)
}
