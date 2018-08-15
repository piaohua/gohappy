package main

import (
	"os"
	"os/signal"
	"runtime"
	"time"

	"gohappy/game/config"
	"gohappy/glog"
	"utils"

	ini "gopkg.in/ini.v1"
)

var (
	cfg *ini.File
	sec *ini.Section
	err error

	aesEnc *utils.AesEncrypt

	aesStatus bool
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
	//初始化
	aesInit()
	config.Init2Game()
	JtPayInit()
	//启动服务
	bind := cfg.Section("login").Key("bind").Value()
	kind := cfg.Section("login").Key("kind").Value()
	NewRemote(bind, kind)
	//监听地址
	addr := cfg.Section("login").Key("addr").Value()
	//启动监听
	go Start(addr)
	go Wxmp()
	signalListen() //监听关闭信号
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

//加密初始化
func aesInit() {
	aesEnc = new(utils.AesEncrypt)
	key := cfg.Section("login").Key("key").Value()
	aesEnc.SetKey([]byte(key))
	aesStatus = cfg.Section("login").Key("status").MustBool(false)
}

//加密
func aesEn(doc string) (arrEncrypt []byte) {
	arrEncrypt, err = aesEnc.Encrypt([]byte(doc))
	if err != nil {
		glog.Errorf("arrEncrypt: %s", doc)
	}
	return
}

//解密
func aesDe(arrEncrypt []byte) (strMsg string) {
	bMsg, err := aesEnc.Decrypt(arrEncrypt)
	if err != nil {
		glog.Errorf("arrEncrypt: %s", string(arrEncrypt))
	}
	strMsg = string(bMsg)
	return
}
