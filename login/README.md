# login

* http请求登录
* 返回网关信息
* 定时向中心分发器获取可用网关

## Installation

```
go get -u github.com/valyala/quicktemplate
go get -u github.com/valyala/quicktemplate/qtc

cd $GOPATH/bin
./qtc -file ../src/gohappy/login/templates/download.qtpl
./qtc -file ../src/gohappy/login/templates/jtpayorder.qtpl
./qtc -file ../src/gohappy/login/templates/jtpayreturn.qtpl

cd $GOPATH/src/gohappy/bin
./ctrl build login linux
```

## Usage:

```
./login-bin -log_dir=./logs > /dev/null 2&>1 &

./ctrl start login
./ctrl stop login
```

## TODO
    优化
    请求加密

## Document
```
//同步变动货币数据(充值或后台操作等)
message PayCurrency {
    string Userid = 1;
    int32 Type = 2;//变动类型
    int64 Coin = 3;//变动金币数量
    int64 Diamond = 4;//变动钻石数量
    int64 Chip = 5;//变动筹码数量
    int64 Card = 6;//变动房卡数量
    int64 Money = 7;//变动充值数量
}

enum WebCode {
    WebReqMsg   = 0; //正常
    WebOnline   = 1; //在线状态
    WebNotice   = 2; //公告
    WebBuild    = 3; //绑定
    WebGive     = 4; //赠送钻石
    WebShop     = 5; //商贸城
    WebEnv      = 6; //设置变量
    WebGame     = 7; //游戏列表
    WebVip      = 8; //VIP
    WebNumber   = 9; //在线人数
    WebTask     =10; //任务列表
    WebLogin    =11; //登录奖励列表
    WebRate     =12; //设置区域比例
    WebState    =13; //设置代理状态
    WebLucky    =14; //lucky列表
}

enum ConfigAtype {
    CONFIG_UPSERT  = 0; //插入或更新
    CONFIG_DELETE  = 1; //删除
}

//web请求 (post)
message WebRequest
{
    WebCode Code = 1; //协议号
    ConfigAtype Atype = 2 ;//操作类型
    bytes Data = 3; //数据
}

message WebResponse
{
    WebCode Code = 1; //协议号
    int32 ErrCode = 2; //错误码
    string ErrMsg = 3; //错误信息
    bytes Result = 4; //正常时返回信息
}

WebRequest 请求request消息(json格式)
    b = new(WebRequest)
    body, err := json.Marshal(b)
WebResponse 响应response消息(json格式)
    b := new(WebResponse)
    d, err := json.Unmarshal(body, b)
    ErrCode 不为0 或者 ErrMsg 不为空 返回错误

1、WebOnline 请求Data数据 (json格式)
    b := make([]string, 0) //玩家id列表
    d, err := json.Marshal(b)
    request.Data = d
    example: ["userid1","userid2"]

   WebOnline 响应Data数据 (json格式)
    b := make(map[string]int) //key:玩家id, value: 1表示在线,0表示离线
    d, err := json.Marshal(b)
    response.Result = d
    example: {"userid1":1,"userid2":0}

2、WebNumber 请求Data数据 (json格式)
    request.Data = []byte{} //空数据
    example: {}

   WebNumber 响应Data数据 (json格式)
    b := make(map[int]int) //key: 1 机器人,2 玩家, value: 数量
    d, err := json.Marshal(b)
    response.Result = d
    example: {"1": 10, "2": 11}

3、WebGive 请求Data数据 (json格式)
    b := new(pb.PayCurrency)
    d, err := json.Marshal(b)
    request.Data = d
    example: {"userid": "xxx"}

   WebGive 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

4、WebShop 请求Data数据 (json格式)
    b := make(map[string]data.Shop) //key: Shop.Id
    d, err := json.Marshal(b)
    request.Data = d
    example: {"id1": {xxx}, "id2": {xxx}}

   WebShop 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

5、WebEnv 请求Data数据 (json格式)
    b := make(map[string]int32) //key: GameEnv或自定义, value: 数值
    example: {"ENV1": 1, "ENV2": 2}

   WebEnv 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

6、WebNotice 请求Data数据 (json格式)
    b := make(map[string]data.Notice) //key: Notice.Id
    d, err := json.Marshal(b)
    request.Data = d
    example: {"id1": {xxx}}

   WebNotice 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

7、WebGame 请求Data数据 (json格式)
    b := make(map[string]data.Game) //key: Game.Id
    d, err := json.Marshal(b)
    request.Data = d
    example: {"id1": {xxx}}

   WebGame 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

8、WebVip 请求Data数据 (json格式)
    b := make(map[string]data.Vip) //key: Vip.Id
    d, err := json.Marshal(b)
    request.Data = d
    example: {"id1": {xxx}}

   WebVip 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

9、WebTask 请求Data数据 (json格式)
    b := make(map[int32]data.Task) //key: Task.Taskid
    d, err := json.Marshal(&b)
    request.Data = d
    example: {"taskid1": {xxx}, "taskid2": {xxx}}

   WebTask 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

10、WebLogin 请求Data数据 (json格式)
    b := make(map[uint32]data.LoginPrize) //key: LoginPrize.Day
    d, err := json.Marshal(&b)
    request.Result = d
    example: {"day1": {xxx}, "day2": {xxx}}

    WebLogin 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

11、WebRate 请求Data数据 (json格式)
    b := new(pb.SetAgentProfitRate)
    d, err := json.Marshal(b)
    request.Data = d
    example: {"userid": "xxx", "rate": xxx}

    WebRate 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

12、WebBuild 请求Data数据 (json格式)
    b := new(pb.SetAgentBuild)
    d, err := json.Marshal(b)
    request.Data = d
    example: {"userid": "xxx", "agent": "xxx"}

    WebBuild 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

13、WebState 请求Data数据 (json格式)
    b := new(pb.SetAgentState)
    d, err := json.Marshal(b)
    request.Data = d
    example: {"userid": "xxx", "state": xxx, "level": xxx}

    WebState 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

14、WebLucky 请求Data数据 (json格式)
    b := make(map[int32]data.Lucky) //key: Lucky.Luckyid
    d, err := json.Marshal(&b)
    request.Data = d
    example: {"luckyid1": {xxx}, "luckyid2": {xxx}}

   WebLucky 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

15、WebVaild 请求Data数据 (json格式)
    b := new(pb.AgentBuildUpdate)
    d, err := json.Marshal(b)
    request.Data = d
    example: {"userid": "xxx", "userid": "xxx", "AgentChild": 1, "BuildVaild": 1, "Build", 1}

    WebVaild 响应Data数据 (json格式)
    response.Result = []byte{} //空数据
    example: {}

```
