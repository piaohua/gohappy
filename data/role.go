package data

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

//Role 角色在线临时数据
type Role struct {
	Gate string     //玩家所在节点
	Pid  *actor.PID //进程ID
}
