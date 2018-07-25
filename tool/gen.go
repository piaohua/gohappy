package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	Init()
	//Gen()
}

//TODO 内部协议通信
//TODO 通过json配置

var (
	protoPacket = make(map[string]uint32) //响应协议
	protoUnpack = make(map[string]uint32) //请求协议
	protoSid    = make(map[uint32]uint32) //映射关系
	//
	packetPath  = "../pb/packet.go"  //打包协议文件路径
	unpackPath  = "../pb/unpack.go"  //解包协议文件路径
	rPacketPath = "../pb/rpacket.go" //机器人打包协议
	rUnpackPath = "../pb/runpack.go" //机器人解包协议
	//
	packetFunc  = "Packet"  //打包协议函数名字
	unpackFunc  = "Unpack"  //解包协议函数名字
	rPacketFunc = "Rpacket" //机器人打包协议函数名字
	rUnpackFunc = "Runpack" //机器人解包协议函数名字
	//
	luaPath  = "MsgID.lua"  //lua文件
	jsPath   = "MsgID.js"   //js文件
	jsonPath = "MsgID.json" //json文件
)

type proto struct {
	name string
	code uint32
}

var protosUnpack = map[string][]proto{
	//game
	"game": {
		{code: 1000, name: "CBuy"},
		{code: 1001, name: "CWxpayOrder"},
		{code: 1002, name: "CJtpayOrder"},
		{code: 1003, name: "CWxpayQuery"},
		{code: 1004, name: "CApplePay"},
		{code: 1005, name: "CShop"},
		//chat
		{code: 1006, name: "CChatText"},
		{code: 1007, name: "CChatVoice"},
		{code: 1008, name: "CNotice"},
		//login
		{code: 1009, name: "CLogin"},
		{code: 1010, name: "CRegist"},
		{code: 1011, name: "CWxLogin"},
		{code: 1012, name: "CResetPwd"},
		{code: 1013, name: "CTourist"},
		//user
		{code: 1014, name: "CUserData"},
		{code: 1015, name: "CGetCurrency"},
		{code: 1016, name: "CPing"},
		{code: 1017, name: "CBank"},
		{code: 1018, name: "CRank"},
		{code: 1019, name: "CTask"},
		{code: 1020, name: "CTaskPrize"},
		{code: 1021, name: "CLoginPrize"},
		{code: 1022, name: "CRoomRecord"},
		{code: 1023, name: "CSignature"},
		{code: 1024, name: "CLatLng"},
		{code: 1025, name: "CBankLog"},
		//agent
		{code: 1050, name: "CAgentJoin"},
		{code: 1051, name: "CMyAgent"},
		{code: 1052, name: "CAgentManage"},
		{code: 1053, name: "CAgentProfit"},
		{code: 1054, name: "CAgentProfitOrder"},
		{code: 1055, name: "CAgentProfitApply"},
		{code: 1056, name: "CAgentProfitRank"},
		{code: 1057, name: "CAgentPlayerManage"},
		{code: 1058, name: "CAgentPlayerApprove"},
		{code: 1059, name: "CAgentProfitReply"},
		{code: 1060, name: "CSetAgentProfitRate"},
	},
	//niu
	"niu": {
		{code: 2000, name: "CNNCoinEnterRoom"},
		{code: 2001, name: "CNNFreeEnterRoom"},
		{code: 2002, name: "CNNFreeDealer"},
		{code: 2003, name: "CNNFreeDealerList"},
		//{code: 2004, name: "CNNFreeSit"},
		{code: 2005, name: "CNNFreeBet"},
		{code: 2006, name: "CNNFreeTrend"},
		{code: 2007, name: "CNNRoomList"},
		{code: 2008, name: "CNNEnterRoom"},
		{code: 2009, name: "CNNCreateRoom"},
		{code: 2010, name: "CNNLeave"},
		{code: 2011, name: "CNNReady"},
		{code: 2012, name: "CNNDealer"},
		{code: 2013, name: "CNNBet"},
		{code: 2014, name: "CNNiu"},
		{code: 2015, name: "CNNGameRecord"},
		{code: 2016, name: "CNNLaunchVote"},
		{code: 2017, name: "CNNVote"},
		{code: 2018, name: "CNNFreeWiners"},
		{code: 2019, name: "CNNFreeRoles"},
		{code: 2020, name: "CNNSit"},
		{code: 2021, name: "CNNCoinChangeRoom"},
	},
	//san
	"san": {
		{code: 3000, name: "CSGCoinEnterRoom"},
		{code: 3001, name: "CSGFreeEnterRoom"},
		{code: 3002, name: "CSGFreeDealer"},
		{code: 3003, name: "CSGFreeDealerList"},
		//{code: 3004, name: "CSGFreeSit"},
		{code: 3005, name: "CSGFreeBet"},
		{code: 3006, name: "CSGFreeTrend"},
		{code: 3007, name: "CSGRoomList"},
		{code: 3008, name: "CSGEnterRoom"},
		{code: 3009, name: "CSGCreateRoom"},
		{code: 3010, name: "CSGLeave"},
		{code: 3011, name: "CSGReady"},
		{code: 3012, name: "CSGDealer"},
		{code: 3013, name: "CSGBet"},
		{code: 3014, name: "CSGiu"},
		{code: 3015, name: "CSGGameRecord"},
		{code: 3016, name: "CSGLaunchVote"},
		{code: 3017, name: "CSGVote"},
		{code: 3018, name: "CSGFreeWiners"},
		{code: 3019, name: "CSGFreeRoles"},
		{code: 3020, name: "CSGSit"},
		{code: 3021, name: "CSGCoinChangeRoom"},
	},
	//hua
	"hua": {
		{code: 4000, name: "CJHCoinEnterRoom"},
		{code: 4001, name: "CJHFreeEnterRoom"},
		{code: 4002, name: "CJHFreeDealer"},
		{code: 4003, name: "CJHFreeDealerList"},
		//{code: 4004, name: "CJHFreeSit"},
		{code: 4005, name: "CJHFreeBet"},
		{code: 4006, name: "CJHFreeTrend"},
		{code: 4007, name: "CJHRoomList"},
		{code: 4008, name: "CJHEnterRoom"},
		{code: 4009, name: "CJHCreateRoom"},
		{code: 4010, name: "CJHLeave"},
		{code: 4011, name: "CJHReady"},
		//{code: 4012, name: "CJHDealer"},
		//{code: 4013, name: "CJHBet"},
		//{code: 4014, name: "CJHiu"},
		{code: 4015, name: "CJHGameRecord"},
		{code: 4016, name: "CJHLaunchVote"},
		{code: 4017, name: "CJHVote"},
		{code: 4018, name: "CJHFreeWiners"},
		{code: 4019, name: "CJHFreeRoles"},
		{code: 4020, name: "CJHCoinSee"},
		//{code: 4021, name: "CJHCoinBet"},
		{code: 4022, name: "CJHCoinCall"},
		{code: 4023, name: "CJHCoinRaise"},
		{code: 4024, name: "CJHCoinFold"},
		{code: 4025, name: "CJHCoinBi"},
		{code: 4026, name: "CJHSit"},
		{code: 4027, name: "CJHCoinChangeRoom"},
	},
	//ebg
	"ebg": {
		{code: 5000, name: "CEBCoinEnterRoom"},
		{code: 5001, name: "CEBFreeEnterRoom"},
		{code: 5002, name: "CEBFreeDealer"},
		{code: 5003, name: "CEBFreeDealerList"},
		//{code: 5004, name: "CEBFreeSit"},
		{code: 5005, name: "CEBFreeBet"},
		{code: 5006, name: "CEBFreeTrend"},
		{code: 5007, name: "CEBRoomList"},
		{code: 5008, name: "CEBEnterRoom"},
		{code: 5009, name: "CEBCreateRoom"},
		{code: 5010, name: "CEBLeave"},
		{code: 5011, name: "CEBReady"},
		{code: 5012, name: "CEBDealer"},
		{code: 5013, name: "CEBBet"},
		{code: 5014, name: "CEBiu"},
		{code: 5015, name: "CEBGameRecord"},
		{code: 5016, name: "CEBLaunchVote"},
		{code: 5017, name: "CEBVote"},
		{code: 5018, name: "CEBFreeWiners"},
		{code: 5019, name: "CEBFreeRoles"},
		{code: 5020, name: "CEBSit"},
		{code: 5021, name: "CEBCoinChangeRoom"},
	},
}

var protosPacket = map[string][]proto{
	//game
	"game": {
		{code: 1500, name: "SBuy"},
		{code: 1501, name: "SWxpayOrder"},
		{code: 1502, name: "SJtpayOrder"},
		{code: 1503, name: "SWxpayQuery"},
		{code: 1504, name: "SApplePay"},
		{code: 1505, name: "SShop"},
		//chat
		{code: 1506, name: "SChatText"},
		{code: 1507, name: "SChatVoice"},
		{code: 1508, name: "SBroadcast"},
		{code: 1509, name: "SNotice"},
		//login
		{code: 1510, name: "SLogin"},
		{code: 1511, name: "SRegist"},
		{code: 1512, name: "SWxLogin"},
		{code: 1513, name: "SLoginOut"},
		{code: 1514, name: "SResetPwd"},
		{code: 1515, name: "STourist"},
		//user
		{code: 1516, name: "SUserData"},
		{code: 1517, name: "SGetCurrency"},
		{code: 1518, name: "SPushCurrency"},
		{code: 1519, name: "SPing"},
		{code: 1520, name: "SBank"},
		{code: 1521, name: "SRank"},
		{code: 1522, name: "STask"},
		{code: 1523, name: "STaskPrize"},
		{code: 1524, name: "SLoginPrize"},
		{code: 1525, name: "SRoomRecord"},
		{code: 1526, name: "SPushNotice"},
		{code: 1527, name: "SSignature"},
		{code: 1528, name: "SLatLng"},
		{code: 1529, name: "SBankLog"},
		//agent
		{code: 1550, name: "SAgentJoin"},
		{code: 1551, name: "SMyAgent"},
		{code: 1552, name: "SAgentManage"},
		{code: 1553, name: "SAgentProfit"},
		{code: 1554, name: "SAgentProfitOrder"},
		{code: 1555, name: "SAgentProfitApply"},
		{code: 1556, name: "SAgentProfitRank"},
		{code: 1557, name: "SAgentPlayerManage"},
		{code: 1558, name: "SAgentPlayerApprove"},
		{code: 1559, name: "SAgentProfitReply"},
		{code: 1560, name: "SSetAgentProfitRate"},
	},
	//niu
	"niu": {
		{code: 2500, name: "SNNCoinEnterRoom"},
		{code: 2501, name: "SNNCoinGameover"},
		{code: 2502, name: "SNNFreeEnterRoom"},
		{code: 2503, name: "SNNFreeCamein"},
		{code: 2504, name: "SNNFreeDealer"},
		{code: 2505, name: "SNNFreeDealerList"},
		//{code: 2506, name: "SNNFreeSit"},
		{code: 2507, name: "SNNFreeBet"},
		{code: 2508, name: "SNNFreeGamestart"},
		{code: 2509, name: "SNNFreeGameover"},
		{code: 2510, name: "SNNFreeTrend"},
		{code: 2511, name: "SNNRoomList"},
		{code: 2512, name: "SNNEnterRoom"},
		{code: 2513, name: "SNNCreateRoom"},
		{code: 2514, name: "SNNCamein"},
		{code: 2515, name: "SNNLeave"},
		{code: 2516, name: "SNNReady"},
		{code: 2517, name: "SNNDraw"},
		{code: 2518, name: "SNNDealer"},
		{code: 2519, name: "SNNPushDealer"},
		{code: 2520, name: "SNNBet"},
		{code: 2521, name: "SNNiu"},
		{code: 2522, name: "SNNGameover"},
		{code: 2523, name: "SNNGameRecord"},
		{code: 2524, name: "SNNLaunchVote"},
		{code: 2525, name: "SNNVote"},
		{code: 2526, name: "SNNVoteResult"},
		{code: 2527, name: "SNNPushState"},
		{code: 2528, name: "SNNFreeWiners"},
		{code: 2529, name: "SNNFreeRoles"},
		{code: 2530, name: "SNNSit"},
		{code: 2531, name: "SNNPushOffline"},
		{code: 2532, name: "SNNCoinChangeRoom"},
		{code: 2533, name: "SNNPushDrawCoin"},
		{code: 2534, name: "SNNPushAward"},
	},
	//san
	"san": {
		{code: 3500, name: "SSGCoinEnterRoom"},
		{code: 3501, name: "SSGCoinGameover"},
		{code: 3502, name: "SSGFreeEnterRoom"},
		{code: 3503, name: "SSGFreeCamein"},
		{code: 3504, name: "SSGFreeDealer"},
		{code: 3505, name: "SSGFreeDealerList"},
		//{code: 3506, name: "SSGFreeSit"},
		{code: 3507, name: "SSGFreeBet"},
		{code: 3508, name: "SSGFreeGamestart"},
		{code: 3509, name: "SSGFreeGameover"},
		{code: 3510, name: "SSGFreeTrend"},
		{code: 3511, name: "SSGRoomList"},
		{code: 3512, name: "SSGEnterRoom"},
		{code: 3513, name: "SSGCreateRoom"},
		{code: 3514, name: "SSGCamein"},
		{code: 3515, name: "SSGLeave"},
		{code: 3516, name: "SSGReady"},
		{code: 3517, name: "SSGDraw"},
		{code: 3518, name: "SSGDealer"},
		{code: 3519, name: "SSGPushDealer"},
		{code: 3520, name: "SSGBet"},
		{code: 3521, name: "SSGiu"},
		{code: 3522, name: "SSGGameover"},
		{code: 3523, name: "SSGGameRecord"},
		{code: 3524, name: "SSGLaunchVote"},
		{code: 3525, name: "SSGVote"},
		{code: 3526, name: "SSGVoteResult"},
		{code: 3527, name: "SSGPushState"},
		{code: 3528, name: "SSGFreeWiners"},
		{code: 3529, name: "SSGFreeRoles"},
		{code: 3530, name: "SSGSit"},
		{code: 3531, name: "SSGPushOffline"},
		{code: 3532, name: "SSGCoinChangeRoom"},
		{code: 3533, name: "SSGPushDrawCoin"},
	},
	//hua
	"hua": {
		{code: 4500, name: "SJHCoinEnterRoom"},
		{code: 4501, name: "SJHCoinGameover"},
		{code: 4502, name: "SJHFreeEnterRoom"},
		{code: 4503, name: "SJHFreeCamein"},
		{code: 4504, name: "SJHFreeDealer"},
		{code: 4505, name: "SJHFreeDealerList"},
		//{code: 4506, name: "SJHFreeSit"},
		{code: 4507, name: "SJHFreeBet"},
		{code: 4508, name: "SJHFreeGamestart"},
		{code: 4509, name: "SJHFreeGameover"},
		{code: 4510, name: "SJHFreeTrend"},
		{code: 4511, name: "SJHRoomList"},
		{code: 4512, name: "SJHEnterRoom"},
		{code: 4513, name: "SJHCreateRoom"},
		{code: 4514, name: "SJHCamein"},
		{code: 4515, name: "SJHLeave"},
		{code: 4516, name: "SJHReady"},
		{code: 4517, name: "SJHDraw"},
		//{code: 4518, name: "SJHDealer"},
		{code: 4519, name: "SJHPushDealer"},
		//{code: 4520, name: "SJHBet"},
		//{code: 4521, name: "SJHiu"},
		{code: 4522, name: "SJHGameover"},
		{code: 4523, name: "SJHGameRecord"},
		{code: 4524, name: "SJHLaunchVote"},
		{code: 4525, name: "SJHVote"},
		{code: 4526, name: "SJHVoteResult"},
		{code: 4527, name: "SJHPushState"},
		{code: 4528, name: "SJHFreeWiners"},
		{code: 4529, name: "SJHFreeRoles"},
		{code: 4530, name: "SJHPushActState"},
		{code: 4531, name: "SJHCoinSee"},
		//{code: 4532, name: "SJHCoinBet"},
		{code: 4533, name: "SJHCoinCall"},
		{code: 4534, name: "SJHCoinRaise"},
		{code: 4535, name: "SJHCoinFold"},
		{code: 4536, name: "SJHCoinBi"},
		{code: 4537, name: "SJHSit"},
		{code: 4538, name: "SJHPushOffline"},
		{code: 4539, name: "SJHCoinChangeRoom"},
		{code: 4540, name: "SJHPushDrawCoin"},
	},
	//ebg
	"ebg": {
		{code: 5500, name: "SEBCoinEnterRoom"},
		{code: 5501, name: "SEBCoinGameover"},
		{code: 5502, name: "SEBFreeEnterRoom"},
		{code: 5503, name: "SEBFreeCamein"},
		{code: 5504, name: "SEBFreeDealer"},
		{code: 5505, name: "SEBFreeDealerList"},
		//{code: 5506, name: "SEBFreeSit"},
		{code: 5507, name: "SEBFreeBet"},
		{code: 5508, name: "SEBFreeGamestart"},
		{code: 5509, name: "SEBFreeGameover"},
		{code: 5510, name: "SEBFreeTrend"},
		{code: 5511, name: "SEBRoomList"},
		{code: 5512, name: "SEBEnterRoom"},
		{code: 5513, name: "SEBCreateRoom"},
		{code: 5514, name: "SEBCamein"},
		{code: 5515, name: "SEBLeave"},
		{code: 5516, name: "SEBReady"},
		{code: 5517, name: "SEBDraw"},
		{code: 5518, name: "SEBDealer"},
		{code: 5519, name: "SEBPushDealer"},
		{code: 5520, name: "SEBBet"},
		{code: 5521, name: "SEBiu"},
		{code: 5522, name: "SEBGameover"},
		{code: 5523, name: "SEBGameRecord"},
		{code: 5524, name: "SEBLaunchVote"},
		{code: 5525, name: "SEBVote"},
		{code: 5526, name: "SEBVoteResult"},
		{code: 5527, name: "SEBPushState"},
		{code: 5528, name: "SEBFreeWiners"},
		{code: 5529, name: "SEBFreeRoles"},
		{code: 5530, name: "SEBSit"},
		{code: 5531, name: "SEBPushOffline"},
		{code: 5532, name: "SEBCoinChangeRoom"},
		{code: 5533, name: "SEBPushDrawCoin"},
	},
}

var sids = map[string]uint32{
	"game": 0,
	"niu":  1,
	"san":  2,
	"dou":  3,
	"hua":  4,
	"pai":  5,
	"ebg":  6,
}

//Init 初始化
func Init() {
	var packetStr string
	var unpackStr string
	var rpacketStr string
	var runpackStr string
	//request
	for _, k := range []string{"game", "niu", "san", "hua", "ebg"} {
		m := protosUnpack[k] //有序
		//for k, m := range protosUnpack {
		//初始化
		protoPacket = make(map[string]uint32) //响应协议
		protoUnpack = make(map[string]uint32) //请求协议
		protoSid = make(map[uint32]uint32)    //映射关系
		//组装
		for _, v := range m {
			registUnpack(v.name, v.code)
		}
		for _, v := range protosPacket[k] {
			registPacket(v.name, v.code)
			protoSid[v.code] = sids[k]
		}
		//最后生成文件
		prefix := k + "-" //文件前缀
		//genMsgID(prefix)
		//genjsMsgID(prefix)
		//genjsonMsgID(prefix)
		////打包组装
		//packetStr += bodyPacket()
		//unpackStr += bodyUnpack()
		//rpacketStr += bodyClientPacket()
		//runpackStr += bodyClientUnpack()
		//有序
		genMsgID2(prefix, protosPacket[k], m)
		genjsMsgID2(prefix, protosPacket[k], m)
		genjsonMsgID2(prefix, protosPacket[k], m)
		//打包组装
		packetStr += bodyPacket2(protosPacket[k])
		unpackStr += bodyUnpack2(m)
		rpacketStr += bodyClientPacket2(m)
		runpackStr += bodyClientUnpack2(protosPacket[k])
	}
	//server
	genPacket(packetStr)
	genUnpack(unpackStr)
	//client
	genClientPacket(rpacketStr)
	genClientUnpack(runpackStr)
}

func registUnpack(key string, code uint32) {
	if _, ok := protoUnpack[key]; ok {
		panic(fmt.Sprintf("%s registered: %d", key, code))
	}
	protoUnpack[key] = code
}

func registPacket(key string, code uint32) {
	if _, ok := protoPacket[key]; ok {
		panic(fmt.Sprintf("%s registered: %d", key, code))
	}
	protoPacket[key] = code
}

//Gen 生成文件
func Gen() {
	////server
	//genPacket()
	//genUnpack()
	////client
	//genClientPacket()
	//genClientUnpack()
}

//生成打包文件
func genPacket(body string) {
	var str string
	str += headPacket()
	//str += bodyPacket()
	str += body
	str += endPacket2()
	err := ioutil.WriteFile(packetPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成解包文件
func genUnpack(body string) {
	var str string
	str += headUnpack()
	//str += bodyUnpack()
	str += body
	str += endUnpack()
	err := ioutil.WriteFile(unpackPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

func bodyUnpack() string {
	var str string
	for k, v := range protoUnpack {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v, k, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v, k, bodyUnpackCode(v), resultUnpack())
	}
	return str
}

func bodyPacket() string {
	var str string
	for k, v := range protoPacket {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, k, resultPacket2(v))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, bodyPacketCode(v, k), k, resultPacket2(v))
	}
	return str
}

//有序
func bodyUnpack2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v.code, v.name, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v.code, v.name, bodyUnpackCode(v.name), resultUnpack())
	}
	return str
}

func bodyPacket2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, v.name, resultPacket2(v.code))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, bodyPacketCode(v.code, v.name), v.name, resultPacket2(v.code))
	}
	return str
}

func bodyUnpackCode(code uint32) (str string) {
	str = fmt.Sprintf("//msg.Code = %d", code)
	return
}

func bodyPacketCode(code uint32, name string) (str string) {
	str = fmt.Sprintf("//msg.(*%s).Code = %d", name, code)
	return
}

func headPacket() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Packet 打包消息
func Packet(msg interface{}) (uint32, uint32, []byte, error) {
	switch msg.(type) {
	`)
}

func headUnpack() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Unpack 解包消息
func Unpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	`)
}

func headRpacket() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Rpacket 打包消息
func Rpacket(msg interface{}) (uint32, []byte, error) {
	switch msg.(type) {
	`)
}

func headRunpack() string {
	return fmt.Sprintf(`// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Runpack 解包消息
func Runpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	`)
}

func resultPacket(code uint32) string {
	return fmt.Sprintf("return %d, b, err", code)
}

func resultPacket2(code uint32) string {
	return fmt.Sprintf("return %d, %d, b, err", code, protoSid[code])
}

func resultUnpack() string {
	return fmt.Sprintf(`err := msg.Unmarshal(b)
		return msg, err`)
}

func endPacket() string {
	return fmt.Sprintf(`default:
		return 0, []byte{}, errors.New("unknown message")
	}
}`)
}

func endPacket2() string {
	return fmt.Sprintf(`default:
		return 0, 0, []byte{}, errors.New("unknown message")
	}
}`)
}

func endUnpack() string {
	return fmt.Sprintf(`default:
		return nil, errors.New("unknown message")
	}
}`)
}

//生成lua文件
func genMsgID(prefix string) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for k, v := range protoUnpack {
		str += fmt.Sprintf("\n\t%s = %d,", k, v)
	}
	str += fmt.Sprintf("\n")
	for k, v := range protoPacket {
		str += fmt.Sprintf("\n\t%s = %d,", k, v)
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+luaPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成js文件
func genjsMsgID(prefix string) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for k, v := range protoUnpack {
		str += fmt.Sprintf("\n\t%s : %d,", k, v)
	}
	str += fmt.Sprintf("\n")
	length := len(protoPacket)
	var i int
	for k, v := range protoPacket {
		i++
		if i == length {
			str += fmt.Sprintf("\n\t%s : %d", k, v)
		} else {
			str += fmt.Sprintf("\n\t%s : %d,", k, v)
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//{
//	3028:{type:"room",        sendType:"protocol.CChat",            revType:"protocol.SChat",           },
//}
func genjsonMsgID(prefix string) {
	var str string
	str += fmt.Sprintf("{")
	//每条协议id唯一
	for k, v := range protoUnpack { //响应
		rsp := ""
		for k2, v2 := range protoPacket { //请求
			if v == v2 {
				rsp = k2
				break
			}
		}
		if len(rsp) == 0 {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"%s\",\t\t},", v, k, rsp)
		} else {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v, k, rsp)
		}
	}
	//
	length := len(protoPacket)
	var i int
	for k, v := range protoPacket { //响应
		rsp := ""
		for k2, v2 := range protoUnpack { //请求
			if v == v2 {
				rsp = k2
				break
			}
		}
		i++
		if i == length {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t}", v, rsp, k)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t}", v, rsp, k)
			}
		} else {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t},", v, rsp, k)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v, rsp, k)
			}
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsonPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//有序
//生成lua文件
func genMsgID2(prefix string, ps, unps []proto) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for _, v := range unps {
		str += fmt.Sprintf("\n\t%s = %d,", v.name, v.code)
	}
	str += fmt.Sprintf("\n")
	for _, v := range ps {
		str += fmt.Sprintf("\n\t%s = %d,", v.name, v.code)
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+luaPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成js文件
func genjsMsgID2(prefix string, ps, unps []proto) {
	var str string
	str += fmt.Sprintf("msgID = {")
	for _, v := range unps {
		str += fmt.Sprintf("\n\t%s : %d,", v.name, v.code)
	}
	str += fmt.Sprintf("\n")
	length := len(ps)
	var i int
	for _, v := range ps {
		i++
		if i == length {
			str += fmt.Sprintf("\n\t%s : %d", v.name, v.code)
		} else {
			str += fmt.Sprintf("\n\t%s : %d,", v.name, v.code)
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//{
//	3028:{type:"room",        sendType:"protocol.CChat",            revType:"protocol.SChat",           },
//}
func genjsonMsgID2(prefix string, ps, unps []proto) {
	var str string
	str += fmt.Sprintf("{")
	//每条协议id唯一
	for _, v := range unps { //响应
		rsp := ""
		for _, v2 := range ps { //请求
			if v.code == v2.code {
				rsp = v2.name
				break
			}
		}
		if len(rsp) == 0 {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"%s\",\t\t},", v.code, v.name, rsp)
		} else {
			str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v.code, v.name, rsp)
		}
	}
	//
	length := len(ps)
	var i int
	for _, v := range ps { //响应
		rsp := ""
		for _, v2 := range unps { //请求
			if v.code == v2.code {
				rsp = v2.name
				break
			}
		}
		i++
		if i == length {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t}", v.code, rsp, v.name)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t}", v.code, rsp, v.name)
			}
		} else {
			if len(rsp) == 0 {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"%s\",\t\trevType:\"pb.%s\",\t\t},", v.code, rsp, v.name)
			} else {
				str += fmt.Sprintf("\n\t%d:{type:\"room\",\t\tsendType:\"pb.%s\",\t\trevType:\"pb.%s\",\t\t},", v.code, rsp, v.name)
			}
		}
	}
	str += fmt.Sprintf("\n}")
	err := ioutil.WriteFile(prefix+jsonPath, []byte(str), 0666)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成机器人打包文件
func genClientPacket(body string) {
	var str string
	str += headRpacket()
	//str += bodyClientPacket()
	str += body
	str += endPacket()
	err := ioutil.WriteFile(rPacketPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

//生成机器人解包文件
func genClientUnpack(body string) {
	var str string
	str += headRunpack()
	//str += bodyClientUnpack()
	str += body
	str += endUnpack()
	err := ioutil.WriteFile(rUnpackPath, []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("write file err -> %v\n", err))
	}
}

func bodyClientPacket() string {
	var str string
	for k, v := range protoUnpack {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, k, resultPacket(v))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", k, bodyClientPacketCode(v, k), k, resultPacket(v))
	}
	return str
}

func bodyClientUnpack() string {
	var str string
	for k, v := range protoPacket {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v, k, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v, k, bodyClientUnpackCode(v), resultUnpack())
	}
	return str
}

//有序
func bodyClientPacket2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case *%s:\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, v.name, resultPacket(v.code))
		//str += fmt.Sprintf("case *%s:\n\t\t%s\n\t\tb, err := msg.(*%s).Marshal()\n\t\t%s\n\t", v.name, bodyClientPacketCode(v.code, v.name), v.name, resultPacket(v.code))
	}
	return str
}

func bodyClientUnpack2(ps []proto) string {
	var str string
	for _, v := range ps {
		str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t", v.code, v.name, resultUnpack())
		//str += fmt.Sprintf("case %d:\n\t\tmsg := new(%s)\n\t\t%s\n\t\t%s\n\t", v.code, v.name, bodyClientUnpackCode(v.code), resultUnpack())
	}
	return str
}

func bodyClientUnpackCode(code uint32) (str string) {
	str = fmt.Sprintf("//msg.Code = %d", code)
	return
}

func bodyClientPacketCode(code uint32, name string) (str string) {
	str = fmt.Sprintf("//msg.(*%s).Code = %d", name, code)
	return
}
