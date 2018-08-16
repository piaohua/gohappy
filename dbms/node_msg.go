package main

import (
	"time"

	"gohappy/game/handler"
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
)

//Handler 消息处理
func (a *DBMSActor) Handler(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.Connected:
		arg := msg.(*pb.Connected)
		glog.Infof("Connected %s", arg.Name)
	case *pb.Disconnected:
		arg := msg.(*pb.Disconnected)
		glog.Infof("Disconnected %s", arg.Name)
	case *pb.Connect:
		arg := msg.(*pb.Connect)
		//服务注册
		a.serve[arg.Name] = ctx.Sender()
		//响应
		connected := &pb.Connected{
			Name: a.Name,
		}
		ctx.Respond(connected)
		glog.Infof("Connect %s", arg.Name)
		//同步配置到gate,game
		a.syncConfig(arg.Name)
	case *pb.Disconnect:
		arg := msg.(*pb.Disconnect)
		//服务注销
		delete(a.serve, arg.Name)
		//响应
		//disconnected := &pb.Disconnected{
		//	Name: a.Name,
		//}
		//ctx.Respond(disconnected)
		glog.Infof("Disconnect %s", arg.Name)
	case *pb.ServeStop:
		//关闭服务
		a.handlerStop(ctx)
		//响应登录
		rsp := new(pb.ServeStoped)
		ctx.Respond(rsp)
	case *pb.ServeStart:
		a.start(ctx)
		//响应
		//rsp := new(pb.ServeStarted)
		//ctx.Respond(rsp)
	case *pb.Tick:
		a.ding(ctx)
	case proto.Message:
		a.handlerUser(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//启动服务
func (a *DBMSActor) start(ctx actor.Context) {
	glog.Infof("dbms start: %v", ctx.Self().String())
	//TODO 设置测试数据,正式后台配置
	/*handler.SetTaskList3()
	handler.SetEBGCoinGame()
	handler.SetActivityList()
	handler.SetTaskList2()
	handler.SetNiuCoinGame()
	handler.SetLuckyList()
	//handler.SetGameList()
	handler.SetShopList()
	handler.SetTaskList()
	handler.SetLoginPrizeList()
	head := cfg.Section("domain").Key("headimag").Value()
	passwd := cfg.Section("robot").Key("passwd").Value()
	phone := cfg.Section("robot").Key("phone").Value()
	rs := data.RegistRobots4(head, passwd, phone)
	for _, v := range rs {
		rolePid.Tell(v)
	}
	list, err := data.GetAgentDayProfit(&pb.CAgentDayProfit{Selfid:"105757",Page:1})
	glog.Debugf("list %#v, err %v", list, err)*/
	//启动
	go a.ticker(ctx)
	//统计服务
	a.statPid = NewStat()
	statStart(a.statPid)
}

//时钟
func (a *DBMSActor) ticker(ctx actor.Context) {
	tick := time.Tick(30 * time.Second)
	msg := new(pb.Tick)
	for {
		select {
		case <-a.stopCh:
			glog.Info("dbms ticker closed")
			return
		default: //防止阻塞
		}
		select {
		case <-a.stopCh:
			glog.Info("dbms ticker closed")
			return
		case <-tick:
			ctx.Self().Tell(msg)
		}
	}
}

//钟声
func (a *DBMSActor) ding(ctx actor.Context) {
	//glog.Debugf("ding: %v", ctx.Self().String())
	//TODO
}

//关闭时钟
func (a *DBMSActor) closeTick() {
	select {
	case <-a.stopCh:
		return
	default:
		//停止发送消息
		close(a.stopCh)
	}
}

func (a *DBMSActor) handlerStop(ctx actor.Context) {
	glog.Debugf("handlerStop: %s", a.Name)
	//关闭
	a.closeTick()
	//回存数据
	for k := range a.serve {
		glog.Debugf("Stop gate: %s", k)
	}
	//关闭统计服务
	statStop(a.statPid)
}

//同步配置
func (a *DBMSActor) syncConfig(key string) {
	if _, ok := a.serve[key]; !ok {
		glog.Errorf("gate not exists: %s", key)
		return
	}
	pid := a.serve[key]
	a.syncConfig2(pid)
}

//同步配置
func (a *DBMSActor) syncConfig2(pid *actor.PID) {
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_ENV))
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_NOTICE))
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_SHOP))
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_GAMES))
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_TASK))
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_LOGIN))
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_LUCKY))
	pid.Tell(handler.GetSyncConfig(pb.CONFIG_ACT))
}
