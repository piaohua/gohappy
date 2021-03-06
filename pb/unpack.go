// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//Unpack 解包消息
func Unpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	case 1000:
		msg := new(CBuy)
		err := msg.Unmarshal(b)
		return msg, err
	case 1001:
		msg := new(CWxpayOrder)
		err := msg.Unmarshal(b)
		return msg, err
	case 1002:
		msg := new(CJtpayOrder)
		err := msg.Unmarshal(b)
		return msg, err
	case 1003:
		msg := new(CWxpayQuery)
		err := msg.Unmarshal(b)
		return msg, err
	case 1004:
		msg := new(CApplePay)
		err := msg.Unmarshal(b)
		return msg, err
	case 1005:
		msg := new(CShop)
		err := msg.Unmarshal(b)
		return msg, err
	case 1006:
		msg := new(CChatText)
		err := msg.Unmarshal(b)
		return msg, err
	case 1007:
		msg := new(CChatVoice)
		err := msg.Unmarshal(b)
		return msg, err
	case 1008:
		msg := new(CNotice)
		err := msg.Unmarshal(b)
		return msg, err
	case 1009:
		msg := new(CLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1010:
		msg := new(CRegist)
		err := msg.Unmarshal(b)
		return msg, err
	case 1011:
		msg := new(CWxLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1012:
		msg := new(CResetPwd)
		err := msg.Unmarshal(b)
		return msg, err
	case 1013:
		msg := new(CTourist)
		err := msg.Unmarshal(b)
		return msg, err
	case 1014:
		msg := new(CUserData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1015:
		msg := new(CGetCurrency)
		err := msg.Unmarshal(b)
		return msg, err
	case 1016:
		msg := new(CPing)
		err := msg.Unmarshal(b)
		return msg, err
	case 1017:
		msg := new(CBank)
		err := msg.Unmarshal(b)
		return msg, err
	case 1018:
		msg := new(CRank)
		err := msg.Unmarshal(b)
		return msg, err
	case 1019:
		msg := new(CTask)
		err := msg.Unmarshal(b)
		return msg, err
	case 1020:
		msg := new(CTaskPrize)
		err := msg.Unmarshal(b)
		return msg, err
	case 1021:
		msg := new(CLoginPrize)
		err := msg.Unmarshal(b)
		return msg, err
	case 1022:
		msg := new(CRoomRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 1023:
		msg := new(CSignature)
		err := msg.Unmarshal(b)
		return msg, err
	case 1024:
		msg := new(CLatLng)
		err := msg.Unmarshal(b)
		return msg, err
	case 1025:
		msg := new(CBankLog)
		err := msg.Unmarshal(b)
		return msg, err
	case 1026:
		msg := new(CLucky)
		err := msg.Unmarshal(b)
		return msg, err
	case 1027:
		msg := new(CActivity)
		err := msg.Unmarshal(b)
		return msg, err
	case 1028:
		msg := new(CJoinActivity)
		err := msg.Unmarshal(b)
		return msg, err
	case 1050:
		msg := new(CAgentJoin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1051:
		msg := new(CMyAgent)
		err := msg.Unmarshal(b)
		return msg, err
	case 1052:
		msg := new(CAgentManage)
		err := msg.Unmarshal(b)
		return msg, err
	case 1053:
		msg := new(CAgentProfit)
		err := msg.Unmarshal(b)
		return msg, err
	case 1054:
		msg := new(CAgentProfitOrder)
		err := msg.Unmarshal(b)
		return msg, err
	case 1055:
		msg := new(CAgentProfitApply)
		err := msg.Unmarshal(b)
		return msg, err
	case 1056:
		msg := new(CAgentProfitRank)
		err := msg.Unmarshal(b)
		return msg, err
	case 1057:
		msg := new(CAgentPlayerManage)
		err := msg.Unmarshal(b)
		return msg, err
	case 1058:
		msg := new(CAgentPlayerApprove)
		err := msg.Unmarshal(b)
		return msg, err
	case 1059:
		msg := new(CAgentProfitReply)
		err := msg.Unmarshal(b)
		return msg, err
	case 1060:
		msg := new(CSetAgentProfitRate)
		err := msg.Unmarshal(b)
		return msg, err
	case 1061:
		msg := new(CGetAgent)
		err := msg.Unmarshal(b)
		return msg, err
	case 1062:
		msg := new(CAgentProfitManage)
		err := msg.Unmarshal(b)
		return msg, err
	case 1063:
		msg := new(CSetAgentNote)
		err := msg.Unmarshal(b)
		return msg, err
	case 1064:
		msg := new(CAgentDayProfit)
		err := msg.Unmarshal(b)
		return msg, err
	case 1065:
		msg := new(CChatLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 1066:
		msg := new(CChatVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 1067:
		msg := new(CChatVoiceJoin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1068:
		msg := new(CChatVoiceLeft)
		err := msg.Unmarshal(b)
		return msg, err
	case 2000:
		msg := new(CNNCoinEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2001:
		msg := new(CNNFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2002:
		msg := new(CNNFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 2003:
		msg := new(CNNFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 2005:
		msg := new(CNNFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 2006:
		msg := new(CNNFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 2007:
		msg := new(CNNRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 2008:
		msg := new(CNNEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2009:
		msg := new(CNNCreateRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2010:
		msg := new(CNNLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 2011:
		msg := new(CNNReady)
		err := msg.Unmarshal(b)
		return msg, err
	case 2012:
		msg := new(CNNDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 2013:
		msg := new(CNNBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 2014:
		msg := new(CNNiu)
		err := msg.Unmarshal(b)
		return msg, err
	case 2015:
		msg := new(CNNGameRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 2016:
		msg := new(CNNLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 2017:
		msg := new(CNNVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 2018:
		msg := new(CNNFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 2019:
		msg := new(CNNFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 2020:
		msg := new(CNNSit)
		err := msg.Unmarshal(b)
		return msg, err
	case 2021:
		msg := new(CNNCoinChangeRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3000:
		msg := new(CSGCoinEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3001:
		msg := new(CSGFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3002:
		msg := new(CSGFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 3003:
		msg := new(CSGFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 3005:
		msg := new(CSGFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 3006:
		msg := new(CSGFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 3007:
		msg := new(CSGRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 3008:
		msg := new(CSGEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3009:
		msg := new(CSGCreateRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3010:
		msg := new(CSGLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 3011:
		msg := new(CSGReady)
		err := msg.Unmarshal(b)
		return msg, err
	case 3012:
		msg := new(CSGDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 3013:
		msg := new(CSGBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 3014:
		msg := new(CSGiu)
		err := msg.Unmarshal(b)
		return msg, err
	case 3015:
		msg := new(CSGGameRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 3016:
		msg := new(CSGLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 3017:
		msg := new(CSGVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 3018:
		msg := new(CSGFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 3019:
		msg := new(CSGFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 3020:
		msg := new(CSGSit)
		err := msg.Unmarshal(b)
		return msg, err
	case 3021:
		msg := new(CSGCoinChangeRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4000:
		msg := new(CJHCoinEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4001:
		msg := new(CJHFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4002:
		msg := new(CJHFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 4003:
		msg := new(CJHFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 4005:
		msg := new(CJHFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 4006:
		msg := new(CJHFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 4007:
		msg := new(CJHRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 4008:
		msg := new(CJHEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4009:
		msg := new(CJHCreateRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4010:
		msg := new(CJHLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 4011:
		msg := new(CJHReady)
		err := msg.Unmarshal(b)
		return msg, err
	case 4015:
		msg := new(CJHGameRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 4016:
		msg := new(CJHLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 4017:
		msg := new(CJHVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 4018:
		msg := new(CJHFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 4019:
		msg := new(CJHFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 4020:
		msg := new(CJHCoinSee)
		err := msg.Unmarshal(b)
		return msg, err
	case 4022:
		msg := new(CJHCoinCall)
		err := msg.Unmarshal(b)
		return msg, err
	case 4023:
		msg := new(CJHCoinRaise)
		err := msg.Unmarshal(b)
		return msg, err
	case 4024:
		msg := new(CJHCoinFold)
		err := msg.Unmarshal(b)
		return msg, err
	case 4025:
		msg := new(CJHCoinBi)
		err := msg.Unmarshal(b)
		return msg, err
	case 4026:
		msg := new(CJHSit)
		err := msg.Unmarshal(b)
		return msg, err
	case 4027:
		msg := new(CJHCoinChangeRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 5000:
		msg := new(CEBCoinEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 5001:
		msg := new(CEBFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 5002:
		msg := new(CEBFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 5003:
		msg := new(CEBFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 5005:
		msg := new(CEBFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 5006:
		msg := new(CEBFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 5007:
		msg := new(CEBRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 5008:
		msg := new(CEBEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 5009:
		msg := new(CEBCreateRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 5010:
		msg := new(CEBLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 5011:
		msg := new(CEBReady)
		err := msg.Unmarshal(b)
		return msg, err
	case 5012:
		msg := new(CEBDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 5013:
		msg := new(CEBBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 5014:
		msg := new(CEBiu)
		err := msg.Unmarshal(b)
		return msg, err
	case 5015:
		msg := new(CEBGameRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 5016:
		msg := new(CEBLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 5017:
		msg := new(CEBVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 5018:
		msg := new(CEBFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 5019:
		msg := new(CEBFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 5020:
		msg := new(CEBSit)
		err := msg.Unmarshal(b)
		return msg, err
	case 5021:
		msg := new(CEBCoinChangeRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 6001:
		msg := new(CLHFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 6002:
		msg := new(CLHFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 6003:
		msg := new(CLHFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 6005:
		msg := new(CLHFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 6006:
		msg := new(CLHFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 6007:
		msg := new(CLHRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 6010:
		msg := new(CLHLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 6018:
		msg := new(CLHFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 6019:
		msg := new(CLHFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 6020:
		msg := new(CLHSit)
		err := msg.Unmarshal(b)
		return msg, err
	default:
		return nil, errors.New("unknown message")
	}
}