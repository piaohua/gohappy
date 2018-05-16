package config

//配置从数据库初始化
func ConfigInit() {
	InitNotice() //公告服务
	InitShop()   //商城服务
	InitEnv()    //变量服务
	InitGame()   //游戏服务
}

//节点变量初始化, 节点连接时同步数据
func Init2Gate(id, secret, key, machid, pattern, notifyUrl string) {
	InitNotice2() //公告服务
	InitShop2()   //商城服务
	InitEnv2()    //变量服务
	InitGame2()   //游戏服务

	WxLoginInit(id, secret) //微信登录

	WxPayInit(id, key, machid, pattern, notifyUrl) //微信支付
}

//逻辑服初始化
func Init2Game() {
	InitNotice2() //公告服务
	InitShop2()   //商城服务
	InitEnv2()    //变量服务
	InitGame2()   //游戏服务
}
