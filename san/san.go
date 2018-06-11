package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"time"

	"gohappy/game/config"
	"gohappy/glog"

	jsoniter "github.com/json-iterator/go"
	ini "gopkg.in/ini.v1"
)

var (
	cfg *ini.File
	sec *ini.Section
	err error

	node = flag.String("node", "", "If non-empty, start with this node")

	nodeName string

	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer glog.Flush()
	//日志定义
	glog.Init()
	//加载配置
	cfg, err = ini.Load("conf.ini")
	if err != nil {
		panic(err)
	}
	cfg.BlockMode = false //只读
	//启动服务
	if *node == "" {
		panic("unknown node")
	}
	nodeName = cfg.Section("game.san" + *node).Name()
	bind := cfg.Section(nodeName).Key("bind").Value()
	kind := cfg.Section(nodeName).Key("kind").Value()
	NewRemote(bind, kind)
	//初始化
	config.Init2Game()
	signalListen()
	//关闭服务
	Stop()
	//延迟等待
	<-time.After(10 * time.Second) //延迟关闭
}

func signalListen() {
	c := make(chan os.Signal)
	//signal.Notify(c)
	signal.Notify(c, os.Interrupt, os.Kill) //监听SIGINT和SIGKILL信号
	//signal.Stop(c)
	for {
		s := <-c
		glog.Error("get signal:", s)
		return
	}
}
