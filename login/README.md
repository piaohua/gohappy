# login

* http请求登录
* 返回网关信息
* 定时向中心分发器获取可用网关

## Installation

```
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
}

enum ConfigAtype {
    CONFIG_UPSERT  = 0; //插入或更新
    CONFIG_DELETE  = 1; //删除
}

//web请求
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

WebRequest 请求消息(json格式)
WebResponse 响应消息(json格式)
ErrCode 不为0 或者 ErrMsg 不为空 返回错误

1、WebOnline 请求Data数据 (json格式)
    msg1 := make([]string, 0)
    example: ["id1","id2"]

2、WebOnline 响应Data数据 (json格式)
    resp := make(map[string]int)
    //1表示在线,0表示离线
    example: {"id1":1,"id2":0}

3、WebNumber 请求Data数据 (json格式)
    example: {}

4、WebNumber 响应Data数据 (json格式)
    resp := make(map[int]int)
    //响应1 机器人,2 玩家
    example: {1: 10, 2: 11}

5、WebGive 请求Data数据 (json格式)
    msg2 := new(pb.PayCurrency)
    example: {}

6、WebGive 响应Data数据 (json格式)
    example: {}

7、WebShop 请求Data数据 (json格式)
    b := make(map[string]data.Shop)
    example: {"id1": {xxx}, "id2": {xxx}}

8、WebShop 响应Data数据 (json格式)
    example: {}

9、WebEnv 请求Data数据 (json格式)
    b := make(map[string]int32)
    example: {"id1": 1, "id2": 2}

10、WebEnv 响应Data数据 (json格式)
    example: {}

11、WebEnv 请求Data数据 (json格式)
    b := make(map[string]data.Notice)
    example: {"id1": {xxx}}

12、WebEnv 响应Data数据 (json格式)
    example: {}

13、WebGame 请求Data数据 (json格式)
    b := make(map[string]data.Game)
    example: {"id1": {xxx}}

14、WebGame 响应Data数据 (json格式)
    example: {}

15、WebVip 请求Data数据 (json格式)
    b := make(map[string]data.Vip)
    example: {"id1": {xxx}}

16、WebVip 响应Data数据 (json格式)
    example: {}

```
