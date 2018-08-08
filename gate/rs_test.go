package main

import (
	"fmt"
	"testing"
	"time"

	"gohappy/pb"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/remote"
)

func TestRole(t *testing.T) {
	remote.Start("127.0.0.1:0")
	activate2("hello")
	<-time.After(time.Minute)
	activate2("hello")
	console.ReadLine()
}

func activate2(name string) {
	timeout := 1 * time.Second
	pid, err := remote.SpawnNamed("127.0.0.1:8081", "remote1", name, timeout)
	if err != nil {
		//fmt.Println(err)
		//return
	}
	res, _ := pid.GetPid().RequestFuture(new(pb.Request), timeout).Result()
	fmt.Println("res ", res)
	response := res.(*pb.Response)
	fmt.Println(response)
	//pid.Stop()
	//
	//pid, _ = remote.SpawnNamed("127.0.0.1:8080", "remote2", name, timeout)
	res, _ = pid.GetPid().RequestFuture(new(pb.Request), timeout).Result()
	response = res.(*pb.Response)
	fmt.Println(response)
}
