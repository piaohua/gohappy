/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-11-19 11:32:23
 * Filename      : robot.go
 * Description   : 机器人
 * *******************************************************/
package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"time"

	"gohappy/glog"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
	ini "gopkg.in/ini.v1"
)

var (
	cfg *ini.File
	sec *ini.Section
	err error

	nodePid *actor.PID
	dbmsPid *actor.PID
	rolePid *actor.PID

	rbs *RobotServer

	aesEnc   *utils.AesEncrypt
	pbAesEnc *utils.AesEncrypt

	aesStatus   bool
	pbAesStatus bool

	node = flag.String("node", "", "If non-empty, start with this node")
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
	pbAesInit()
	//
	bind := cfg.Section("robot").Key("bind").Value()
	kind := cfg.Section("robot").Key("kind").Value()
	phone := cfg.Section("robot").Key("phone").Value()
	//启动服务
	rbs = new(RobotServer)
	rbs.phone = phone
	if rbs != nil {
		rbs.Start() //启动服务
		rbs.NewRemote(bind, kind)
	}
	//启动服务
	signalListen()
	//关闭服务
	//关闭websocket连接
	if rbs != nil {
		rbs.Close() //关闭服务
	}
	//延迟等待
	<-time.After(10 * time.Second)
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

//加密初始化
func pbAesInit() {
	pbAesEnc = new(utils.AesEncrypt)
	key := cfg.Section("gate").Key("key").Value()
	pbAesEnc.SetKey([]byte(key))
	pbAesStatus = cfg.Section("gate").Key("status").MustBool(false)
}

//加密
func pbAesEn(doc []byte) (arrEncrypt []byte) {
	arrEncrypt, err = pbAesEnc.Encrypt(doc)
	if err != nil {
		glog.Errorf("arrEncrypt: %s", string(doc))
	}
	return
}

//解密
func pbAesDe(arrEncrypt []byte) (bMsg []byte) {
	bMsg, err = pbAesEnc.Decrypt(arrEncrypt)
	if err != nil {
		glog.Errorf("arrEncrypt: %s", string(arrEncrypt))
	}
	return
}
