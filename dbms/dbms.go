package main

import (
	"os"
	"os/signal"
	"runtime"
	"time"

	"gohappy/data"
	"gohappy/game/config"
	"gohappy/glog"

	jsoniter "github.com/json-iterator/go"
	ini "gopkg.in/ini.v1"
)

var (
	cfg *ini.File
	sec *ini.Section
	err error

	json = jsoniter.ConfigCompatibleWithStandardLibrary

	smsusername    string
	smspassword    string
	sms253account  string
	sms253password string
	sms253URL      string
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
	//数据库连接
	host := cfg.Section("mongod").Key("host").Value()
	port := cfg.Section("mongod").Key("port").Value()
	user := cfg.Section("mongod").Key("user").Value()
	passwd := cfg.Section("mongod").Key("passwd").Value()
	dbname := cfg.Section("mongod").Key("name").Value()
	data.InitMgo(host, port, user, passwd, dbname)
	//配置初始化
	config.ConfigInit()
	smsusername = cfg.Section("smsbao").Key("username").Value()
	smspassword = cfg.Section("smsbao").Key("password").Value()
	sms253account = cfg.Section("sms253").Key("account").Value()
	sms253password = cfg.Section("sms253").Key("password").Value()
	sms253URL = cfg.Section("sms253").Key("url").Value()
	//启动服务
	bind := cfg.Section("dbms").Key("bind").Value()
	kind := cfg.Section("dbms").Key("kind").Value()
	room := cfg.Section("dbms").Key("room").Value()
	role := cfg.Section("dbms").Key("role").Value()
	logger := cfg.Section("dbms").Key("logger").Value()
	NewRemote(bind, kind, room, role, logger)
	signalListen()
	//关闭服务
	Stop()
	data.Close() //数据库断开
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
