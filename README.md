# gohappy
* this sample is a game server
* based on [protoactor-go](https://github.com/AsynkronIT/protoactor-go)
* [go-logging](https://github.com/piaohua/go-logging)
* [utils](https://github.com/piaohua/utils)
* [api](https://github.com/piaohua/api)

## Installation
```
cd $GOPATH/src
go clone github.com/piaohua/gohappy
```

## Usage:
```
cd $GOPATH/bin/ctrl
./ctrl

cp conf.local.ini conf.ini

./ctrl proto
./ctrl protojson

./ctrl build login linux
./ctrl build dbms linux
./ctrl build gate linux
./ctrl build niu linux
./ctrl build web

cp gate-bin gate1-bin
cp gate-bin gate2-bin

./ctrl start dbms
./ctrl start login
./ctrl start niu1 -node=1
./ctrl start gate1 -node=1
./ctrl start gate2 -node=2

./ctrl stop niu1 -node=1
./ctrl stop gate1 -node=1
./ctrl stop gate2 -node=2
./ctrl stop login
./ctrl stop dbms
```

## Document
```
包引用规范
import (
    标准库包

    自定义包

    第三方包
)

src/gohappy/
protocol   协议文件目录 Google Protobuf version 3.0
pb         生成协议文件目录, packet unpack 文件
data       数据结构定义, 数据库连接操作 mongodb
tool       生成协议工具
game       逻辑处理 (algo, config, handler, login...),多个子目录
glog       日志记录
test       测试文件

gate       协议转发, 多个, 处理客户端连接, scoket packet unpack
dbms       大厅服务,单个,处理服务注册
robot      机器人, 模拟客户端, 请求顺序login-gate
login      http请求登录,返回网关信息,单个,定时向中心分发器获取可用网关

web        后台 web 服务, based on beego

gate      缓存玩家进程,优化?

xxxxx     游戏逻辑

```

## package & program
```
protocol (proto)
    协议文件目录
    Google Protobuf version 3.0

pb  (package)
    生成协议文件目录
    packet unpack 文件
    rpacket runpack 机器人协议文件
    工具自动生成文件,无需手动修改

data (package)
    数据库操作
    数据结构
    参数定义

dbms (program)
    玩家数据缓存
    玩家数据中心
    logger日志中心
    唯一id管理
    房间列表管理
    房间基础信息
    后台配置数据加载
    邮箱管理
    投注活动
    桌子匹配路由
    处理服务注册
    处理请求转发
    处理网关信息
    玩家节点路由
    桌子节点路由

login (program)
    http请求节点信息
    登录节点分配
    支付回调请求

gate (program)
    websocket连接
    处理消息转发
    消息打包解包
    处理业务逻辑
    响应请求结果

game (package)
    处理业务逻辑
    创建房间逻辑
```

## TODO
* 配置文件动态加载,读取配置服务独立
* data数据库操作服务独立
* crontab服务
* dockerfile
* data数据库mgo操作优化
* dbms数据管理系统,玩家数据中心,优化拆分
* logging,mail,bets等服务拆分
* 版本控制
* game逻辑处理区分dbms,gate操作
* 返回消息按大小拆分发送,客户端处理粘包获取
* dbms分配登录节点,游戏服节点规则
* game开奖种类、开奖时间、开奖间隔等可配置
* ip注册限制
* 最好部署在一个内网环境下,否则超时可能会比较严重
* dbms服务中注册节点时检测节点是否已经注册,如果已经注册则启动失败
* 服务节点后台管理
* 房间后台管理.后台只添加不同房间的配置,游戏内查找配置生成房间
* 前后端协议添加version控制
* 玩家进入游戏时扣除携带

## 启动顺序
    dbms
    game (xxxxx ...)
    gate
    login
    robot

## 停服顺序
    robot
    login
    game (xxxxx ...)
    gate
    dbms
