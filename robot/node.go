/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-11-19 13:12:54
 * Filename      : node.go
 * Description   : 机器人
 * *******************************************************/
package main

import (
	"errors"

	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
)

//NewRemote 启动远程连接服务
func (server *RobotServer) NewRemote(bind, kind string) {
	if bind == "" {
		glog.Panic("bind empty")
	}
	// Start the remote server
	remote.Start(bind)
	server.remoteRecv(kind) //接收远程消息
}

//接收远程消息
func (server *RobotServer) remoteRecv(kind string) {
	//create the channel
	server.channel = make(chan interface{}, 100) //protos中定义
	server.closeCh = make(chan struct{})

	//create an actor receiving messages and pushing them onto the channel
	props := actor.FromFunc(func(context actor.Context) {
		server.remoteSend(context.Message())
	})
	nodePid, err = actor.SpawnNamed(props, kind)
	if err != nil {
		glog.Panic(err)
	}
	server.remoteDbms()

	//consume the channel just like you use to
	go func() {
		for msg := range server.channel {
			err := server.remoteHandler(msg)
			if err != nil {
				//停止发送消息
				close(server.closeCh)
				break
			}
		}
		//channel closed
		server.disconnect()
	}()
}

func (server *RobotServer) remoteSend(message interface{}) {
	if server.channel == nil {
		glog.Errorf("server channel closed %#v", message)
		return
	}
	if len(server.channel) == cap(server.channel) {
		glog.Errorf("send msg channel full -> %d", len(server.channel))
		return
	}
	select {
	case <-server.closeCh:
		return
	default:
	}
	select {
	case <-server.closeCh:
		return
	case server.channel <- message:
	}
}

//处理
func (server *RobotServer) remoteHandler(message interface{}) error {
	switch message.(type) {
	case *pb.RobotMsg:
		msg := message.(*pb.RobotMsg)
		//分配机器人
		go Msg2Robots(msg, msg.Num)
		glog.Infof("node msg -> %v", msg)
	case closeFlag:
		return errors.New("msg channel closed")
	}
	return nil
}

func (server *RobotServer) remoteDbms() {
	//dbms
	bind := cfg.Section("dbms").Key("bind").Value()
	kind := cfg.Section("dbms").Key("kind").Value()
	dbmsPid = actor.NewPID(bind, kind)
	//role
	role := cfg.Section("dbms").Key("role").Value()
	rolePid = actor.NewPID(bind, role)
	//name
	server.Name = cfg.Section("robot").Name()
	connect := &pb.Connect{
		Name: server.Name,
	}
	dbmsPid.Request(connect, nodePid)
}

func (server *RobotServer) disconnect() {
	//dbms
	disconnect := &pb.Disconnect{
		Name: server.Name,
	}
	if dbmsPid != nil {
		dbmsPid.Tell(disconnect)
	}
	if nodePid != nil {
		nodePid.Stop()
	}
}
