// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//打包消息
func Rpacket(msg interface{}) (uint32, []byte, error) {
	switch msg.(type) {
	case *CSGFreeBet:
		b, err := msg.(*CSGFreeBet).Marshal()
		return 3005, b, err
	case *CSGEnterRoom:
		b, err := msg.(*CSGEnterRoom).Marshal()
		return 3008, b, err
	case *CSGCreateRoom:
		b, err := msg.(*CSGCreateRoom).Marshal()
		return 3009, b, err
	case *CSGDealer:
		b, err := msg.(*CSGDealer).Marshal()
		return 3012, b, err
	case *CSGFreeRoles:
		b, err := msg.(*CSGFreeRoles).Marshal()
		return 3019, b, err
	case *CSGLeave:
		b, err := msg.(*CSGLeave).Marshal()
		return 3010, b, err
	case *CSGBet:
		b, err := msg.(*CSGBet).Marshal()
		return 3013, b, err
	case *CSGiu:
		b, err := msg.(*CSGiu).Marshal()
		return 3014, b, err
	case *CSGLaunchVote:
		b, err := msg.(*CSGLaunchVote).Marshal()
		return 3016, b, err
	case *CSGCoinChangeRoom:
		b, err := msg.(*CSGCoinChangeRoom).Marshal()
		return 3021, b, err
	case *CSGFreeEnterRoom:
		b, err := msg.(*CSGFreeEnterRoom).Marshal()
		return 3001, b, err
	case *CSGFreeTrend:
		b, err := msg.(*CSGFreeTrend).Marshal()
		return 3006, b, err
	case *CSGRoomList:
		b, err := msg.(*CSGRoomList).Marshal()
		return 3007, b, err
	case *CSGFreeWiners:
		b, err := msg.(*CSGFreeWiners).Marshal()
		return 3018, b, err
	case *CSGCoinEnterRoom:
		b, err := msg.(*CSGCoinEnterRoom).Marshal()
		return 3000, b, err
	case *CSGFreeDealer:
		b, err := msg.(*CSGFreeDealer).Marshal()
		return 3002, b, err
	case *CSGFreeDealerList:
		b, err := msg.(*CSGFreeDealerList).Marshal()
		return 3003, b, err
	case *CSGReady:
		b, err := msg.(*CSGReady).Marshal()
		return 3011, b, err
	case *CSGGameRecord:
		b, err := msg.(*CSGGameRecord).Marshal()
		return 3015, b, err
	case *CSGVote:
		b, err := msg.(*CSGVote).Marshal()
		return 3017, b, err
	case *CSGSit:
		b, err := msg.(*CSGSit).Marshal()
		return 3020, b, err
	case *CJHEnterRoom:
		b, err := msg.(*CJHEnterRoom).Marshal()
		return 4008, b, err
	case *CJHCreateRoom:
		b, err := msg.(*CJHCreateRoom).Marshal()
		return 4009, b, err
	case *CJHReady:
		b, err := msg.(*CJHReady).Marshal()
		return 4011, b, err
	case *CJHVote:
		b, err := msg.(*CJHVote).Marshal()
		return 4017, b, err
	case *CJHCoinRaise:
		b, err := msg.(*CJHCoinRaise).Marshal()
		return 4023, b, err
	case *CJHCoinFold:
		b, err := msg.(*CJHCoinFold).Marshal()
		return 4024, b, err
	case *CJHFreeDealerList:
		b, err := msg.(*CJHFreeDealerList).Marshal()
		return 4003, b, err
	case *CJHFreeTrend:
		b, err := msg.(*CJHFreeTrend).Marshal()
		return 4006, b, err
	case *CJHCoinBi:
		b, err := msg.(*CJHCoinBi).Marshal()
		return 4025, b, err
	case *CJHGameRecord:
		b, err := msg.(*CJHGameRecord).Marshal()
		return 4015, b, err
	case *CJHLaunchVote:
		b, err := msg.(*CJHLaunchVote).Marshal()
		return 4016, b, err
	case *CJHFreeWiners:
		b, err := msg.(*CJHFreeWiners).Marshal()
		return 4018, b, err
	case *CJHSit:
		b, err := msg.(*CJHSit).Marshal()
		return 4026, b, err
	case *CJHFreeEnterRoom:
		b, err := msg.(*CJHFreeEnterRoom).Marshal()
		return 4001, b, err
	case *CJHFreeDealer:
		b, err := msg.(*CJHFreeDealer).Marshal()
		return 4002, b, err
	case *CJHRoomList:
		b, err := msg.(*CJHRoomList).Marshal()
		return 4007, b, err
	case *CJHLeave:
		b, err := msg.(*CJHLeave).Marshal()
		return 4010, b, err
	case *CJHFreeRoles:
		b, err := msg.(*CJHFreeRoles).Marshal()
		return 4019, b, err
	case *CJHCoinSee:
		b, err := msg.(*CJHCoinSee).Marshal()
		return 4020, b, err
	case *CJHCoinCall:
		b, err := msg.(*CJHCoinCall).Marshal()
		return 4022, b, err
	case *CJHCoinChangeRoom:
		b, err := msg.(*CJHCoinChangeRoom).Marshal()
		return 4027, b, err
	case *CJHCoinEnterRoom:
		b, err := msg.(*CJHCoinEnterRoom).Marshal()
		return 4000, b, err
	case *CJHFreeBet:
		b, err := msg.(*CJHFreeBet).Marshal()
		return 4005, b, err
	case *CWxpayOrder:
		b, err := msg.(*CWxpayOrder).Marshal()
		return 1001, b, err
	case *CLoginPrize:
		b, err := msg.(*CLoginPrize).Marshal()
		return 1021, b, err
	case *CSignature:
		b, err := msg.(*CSignature).Marshal()
		return 1023, b, err
	case *CAgentJoin:
		b, err := msg.(*CAgentJoin).Marshal()
		return 1050, b, err
	case *CAgentManage:
		b, err := msg.(*CAgentManage).Marshal()
		return 1052, b, err
	case *CAgentProfitReply:
		b, err := msg.(*CAgentProfitReply).Marshal()
		return 1059, b, err
	case *CJtpayOrder:
		b, err := msg.(*CJtpayOrder).Marshal()
		return 1002, b, err
	case *CApplePay:
		b, err := msg.(*CApplePay).Marshal()
		return 1004, b, err
	case *CChatText:
		b, err := msg.(*CChatText).Marshal()
		return 1006, b, err
	case *CLogin:
		b, err := msg.(*CLogin).Marshal()
		return 1009, b, err
	case *CWxLogin:
		b, err := msg.(*CWxLogin).Marshal()
		return 1011, b, err
	case *CRoomRecord:
		b, err := msg.(*CRoomRecord).Marshal()
		return 1022, b, err
	case *CAgentProfit:
		b, err := msg.(*CAgentProfit).Marshal()
		return 1053, b, err
	case *CChatVoice:
		b, err := msg.(*CChatVoice).Marshal()
		return 1007, b, err
	case *CNotice:
		b, err := msg.(*CNotice).Marshal()
		return 1008, b, err
	case *CResetPwd:
		b, err := msg.(*CResetPwd).Marshal()
		return 1012, b, err
	case *CGetCurrency:
		b, err := msg.(*CGetCurrency).Marshal()
		return 1015, b, err
	case *CPing:
		b, err := msg.(*CPing).Marshal()
		return 1016, b, err
	case *CAgentProfitApply:
		b, err := msg.(*CAgentProfitApply).Marshal()
		return 1055, b, err
	case *CRank:
		b, err := msg.(*CRank).Marshal()
		return 1018, b, err
	case *CTask:
		b, err := msg.(*CTask).Marshal()
		return 1019, b, err
	case *CLatLng:
		b, err := msg.(*CLatLng).Marshal()
		return 1024, b, err
	case *CWxpayQuery:
		b, err := msg.(*CWxpayQuery).Marshal()
		return 1003, b, err
	case *CShop:
		b, err := msg.(*CShop).Marshal()
		return 1005, b, err
	case *CUserData:
		b, err := msg.(*CUserData).Marshal()
		return 1014, b, err
	case *CTaskPrize:
		b, err := msg.(*CTaskPrize).Marshal()
		return 1020, b, err
	case *CAgentPlayerApprove:
		b, err := msg.(*CAgentPlayerApprove).Marshal()
		return 1058, b, err
	case *CAgentProfitRank:
		b, err := msg.(*CAgentProfitRank).Marshal()
		return 1056, b, err
	case *CAgentPlayerManage:
		b, err := msg.(*CAgentPlayerManage).Marshal()
		return 1057, b, err
	case *CBuy:
		b, err := msg.(*CBuy).Marshal()
		return 1000, b, err
	case *CRegist:
		b, err := msg.(*CRegist).Marshal()
		return 1010, b, err
	case *CTourist:
		b, err := msg.(*CTourist).Marshal()
		return 1013, b, err
	case *CBank:
		b, err := msg.(*CBank).Marshal()
		return 1017, b, err
	case *CMyAgent:
		b, err := msg.(*CMyAgent).Marshal()
		return 1051, b, err
	case *CAgentProfitOrder:
		b, err := msg.(*CAgentProfitOrder).Marshal()
		return 1054, b, err
	case *CNNLeave:
		b, err := msg.(*CNNLeave).Marshal()
		return 2010, b, err
	case *CNNVote:
		b, err := msg.(*CNNVote).Marshal()
		return 2017, b, err
	case *CNNSit:
		b, err := msg.(*CNNSit).Marshal()
		return 2020, b, err
	case *CNNFreeDealer:
		b, err := msg.(*CNNFreeDealer).Marshal()
		return 2002, b, err
	case *CNNFreeDealerList:
		b, err := msg.(*CNNFreeDealerList).Marshal()
		return 2003, b, err
	case *CNNRoomList:
		b, err := msg.(*CNNRoomList).Marshal()
		return 2007, b, err
	case *CNNReady:
		b, err := msg.(*CNNReady).Marshal()
		return 2011, b, err
	case *CNNDealer:
		b, err := msg.(*CNNDealer).Marshal()
		return 2012, b, err
	case *CNNiu:
		b, err := msg.(*CNNiu).Marshal()
		return 2014, b, err
	case *CNNCoinChangeRoom:
		b, err := msg.(*CNNCoinChangeRoom).Marshal()
		return 2021, b, err
	case *CNNFreeEnterRoom:
		b, err := msg.(*CNNFreeEnterRoom).Marshal()
		return 2001, b, err
	case *CNNEnterRoom:
		b, err := msg.(*CNNEnterRoom).Marshal()
		return 2008, b, err
	case *CNNCreateRoom:
		b, err := msg.(*CNNCreateRoom).Marshal()
		return 2009, b, err
	case *CNNBet:
		b, err := msg.(*CNNBet).Marshal()
		return 2013, b, err
	case *CNNFreeWiners:
		b, err := msg.(*CNNFreeWiners).Marshal()
		return 2018, b, err
	case *CNNFreeRoles:
		b, err := msg.(*CNNFreeRoles).Marshal()
		return 2019, b, err
	case *CNNFreeBet:
		b, err := msg.(*CNNFreeBet).Marshal()
		return 2005, b, err
	case *CNNFreeTrend:
		b, err := msg.(*CNNFreeTrend).Marshal()
		return 2006, b, err
	case *CNNGameRecord:
		b, err := msg.(*CNNGameRecord).Marshal()
		return 2015, b, err
	case *CNNLaunchVote:
		b, err := msg.(*CNNLaunchVote).Marshal()
		return 2016, b, err
	case *CNNCoinEnterRoom:
		b, err := msg.(*CNNCoinEnterRoom).Marshal()
		return 2000, b, err
	default:
		return 0, []byte{}, errors.New("unknown message")
	}
}