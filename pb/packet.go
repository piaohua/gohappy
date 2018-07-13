// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//打包消息
func Packet(msg interface{}) (uint32, uint32, []byte, error) {
	switch msg.(type) {
	case *SAgentProfitOrder:
		b, err := msg.(*SAgentProfitOrder).Marshal()
		return 1554, 0, b, err
	case *SAgentPlayerManage:
		b, err := msg.(*SAgentPlayerManage).Marshal()
		return 1557, 0, b, err
	case *SAgentProfitReply:
		b, err := msg.(*SAgentProfitReply).Marshal()
		return 1559, 0, b, err
	case *SShop:
		b, err := msg.(*SShop).Marshal()
		return 1505, 0, b, err
	case *SLoginOut:
		b, err := msg.(*SLoginOut).Marshal()
		return 1513, 0, b, err
	case *SAgentProfit:
		b, err := msg.(*SAgentProfit).Marshal()
		return 1553, 0, b, err
	case *SBuy:
		b, err := msg.(*SBuy).Marshal()
		return 1500, 0, b, err
	case *SUserData:
		b, err := msg.(*SUserData).Marshal()
		return 1516, 0, b, err
	case *SBank:
		b, err := msg.(*SBank).Marshal()
		return 1520, 0, b, err
	case *SLatLng:
		b, err := msg.(*SLatLng).Marshal()
		return 1528, 0, b, err
	case *SApplePay:
		b, err := msg.(*SApplePay).Marshal()
		return 1504, 0, b, err
	case *SBroadcast:
		b, err := msg.(*SBroadcast).Marshal()
		return 1508, 0, b, err
	case *SPushCurrency:
		b, err := msg.(*SPushCurrency).Marshal()
		return 1518, 0, b, err
	case *SWxpayQuery:
		b, err := msg.(*SWxpayQuery).Marshal()
		return 1503, 0, b, err
	case *SNotice:
		b, err := msg.(*SNotice).Marshal()
		return 1509, 0, b, err
	case *SBankLog:
		b, err := msg.(*SBankLog).Marshal()
		return 1529, 0, b, err
	case *SPushNotice:
		b, err := msg.(*SPushNotice).Marshal()
		return 1526, 0, b, err
	case *SAgentPlayerApprove:
		b, err := msg.(*SAgentPlayerApprove).Marshal()
		return 1558, 0, b, err
	case *SChatVoice:
		b, err := msg.(*SChatVoice).Marshal()
		return 1507, 0, b, err
	case *STaskPrize:
		b, err := msg.(*STaskPrize).Marshal()
		return 1523, 0, b, err
	case *SRoomRecord:
		b, err := msg.(*SRoomRecord).Marshal()
		return 1525, 0, b, err
	case *SLogin:
		b, err := msg.(*SLogin).Marshal()
		return 1510, 0, b, err
	case *SPing:
		b, err := msg.(*SPing).Marshal()
		return 1519, 0, b, err
	case *SAgentJoin:
		b, err := msg.(*SAgentJoin).Marshal()
		return 1550, 0, b, err
	case *SMyAgent:
		b, err := msg.(*SMyAgent).Marshal()
		return 1551, 0, b, err
	case *SWxpayOrder:
		b, err := msg.(*SWxpayOrder).Marshal()
		return 1501, 0, b, err
	case *SJtpayOrder:
		b, err := msg.(*SJtpayOrder).Marshal()
		return 1502, 0, b, err
	case *SChatText:
		b, err := msg.(*SChatText).Marshal()
		return 1506, 0, b, err
	case *SAgentProfitApply:
		b, err := msg.(*SAgentProfitApply).Marshal()
		return 1555, 0, b, err
	case *SAgentProfitRank:
		b, err := msg.(*SAgentProfitRank).Marshal()
		return 1556, 0, b, err
	case *SWxLogin:
		b, err := msg.(*SWxLogin).Marshal()
		return 1512, 0, b, err
	case *SGetCurrency:
		b, err := msg.(*SGetCurrency).Marshal()
		return 1517, 0, b, err
	case *SSignature:
		b, err := msg.(*SSignature).Marshal()
		return 1527, 0, b, err
	case *SRank:
		b, err := msg.(*SRank).Marshal()
		return 1521, 0, b, err
	case *STask:
		b, err := msg.(*STask).Marshal()
		return 1522, 0, b, err
	case *SLoginPrize:
		b, err := msg.(*SLoginPrize).Marshal()
		return 1524, 0, b, err
	case *SAgentManage:
		b, err := msg.(*SAgentManage).Marshal()
		return 1552, 0, b, err
	case *SRegist:
		b, err := msg.(*SRegist).Marshal()
		return 1511, 0, b, err
	case *SResetPwd:
		b, err := msg.(*SResetPwd).Marshal()
		return 1514, 0, b, err
	case *STourist:
		b, err := msg.(*STourist).Marshal()
		return 1515, 0, b, err
	case *SNNCoinEnterRoom:
		b, err := msg.(*SNNCoinEnterRoom).Marshal()
		return 2500, 1, b, err
	case *SNNFreeEnterRoom:
		b, err := msg.(*SNNFreeEnterRoom).Marshal()
		return 2502, 1, b, err
	case *SNNFreeBet:
		b, err := msg.(*SNNFreeBet).Marshal()
		return 2507, 1, b, err
	case *SNNFreeGamestart:
		b, err := msg.(*SNNFreeGamestart).Marshal()
		return 2508, 1, b, err
	case *SNNFreeGameover:
		b, err := msg.(*SNNFreeGameover).Marshal()
		return 2509, 1, b, err
	case *SNNEnterRoom:
		b, err := msg.(*SNNEnterRoom).Marshal()
		return 2512, 1, b, err
	case *SNNDraw:
		b, err := msg.(*SNNDraw).Marshal()
		return 2517, 1, b, err
	case *SNNPushDealer:
		b, err := msg.(*SNNPushDealer).Marshal()
		return 2519, 1, b, err
	case *SNNLeave:
		b, err := msg.(*SNNLeave).Marshal()
		return 2515, 1, b, err
	case *SNNBet:
		b, err := msg.(*SNNBet).Marshal()
		return 2520, 1, b, err
	case *SNNSit:
		b, err := msg.(*SNNSit).Marshal()
		return 2530, 1, b, err
	case *SNNFreeWiners:
		b, err := msg.(*SNNFreeWiners).Marshal()
		return 2528, 1, b, err
	case *SNNPushOffline:
		b, err := msg.(*SNNPushOffline).Marshal()
		return 2531, 1, b, err
	case *SNNPushDrawCoin:
		b, err := msg.(*SNNPushDrawCoin).Marshal()
		return 2533, 1, b, err
	case *SNNGameRecord:
		b, err := msg.(*SNNGameRecord).Marshal()
		return 2523, 1, b, err
	case *SNNLaunchVote:
		b, err := msg.(*SNNLaunchVote).Marshal()
		return 2524, 1, b, err
	case *SNNFreeRoles:
		b, err := msg.(*SNNFreeRoles).Marshal()
		return 2529, 1, b, err
	case *SNNFreeCamein:
		b, err := msg.(*SNNFreeCamein).Marshal()
		return 2503, 1, b, err
	case *SNNFreeDealer:
		b, err := msg.(*SNNFreeDealer).Marshal()
		return 2504, 1, b, err
	case *SNNFreeDealerList:
		b, err := msg.(*SNNFreeDealerList).Marshal()
		return 2505, 1, b, err
	case *SNNRoomList:
		b, err := msg.(*SNNRoomList).Marshal()
		return 2511, 1, b, err
	case *SNNGameover:
		b, err := msg.(*SNNGameover).Marshal()
		return 2522, 1, b, err
	case *SNNCoinChangeRoom:
		b, err := msg.(*SNNCoinChangeRoom).Marshal()
		return 2532, 1, b, err
	case *SNNCreateRoom:
		b, err := msg.(*SNNCreateRoom).Marshal()
		return 2513, 1, b, err
	case *SNNReady:
		b, err := msg.(*SNNReady).Marshal()
		return 2516, 1, b, err
	case *SNNVote:
		b, err := msg.(*SNNVote).Marshal()
		return 2525, 1, b, err
	case *SNNVoteResult:
		b, err := msg.(*SNNVoteResult).Marshal()
		return 2526, 1, b, err
	case *SNNCoinGameover:
		b, err := msg.(*SNNCoinGameover).Marshal()
		return 2501, 1, b, err
	case *SNNFreeTrend:
		b, err := msg.(*SNNFreeTrend).Marshal()
		return 2510, 1, b, err
	case *SNNCamein:
		b, err := msg.(*SNNCamein).Marshal()
		return 2514, 1, b, err
	case *SNNDealer:
		b, err := msg.(*SNNDealer).Marshal()
		return 2518, 1, b, err
	case *SNNiu:
		b, err := msg.(*SNNiu).Marshal()
		return 2521, 1, b, err
	case *SNNPushState:
		b, err := msg.(*SNNPushState).Marshal()
		return 2527, 1, b, err
	case *SSGCoinEnterRoom:
		b, err := msg.(*SSGCoinEnterRoom).Marshal()
		return 3500, 2, b, err
	case *SSGFreeGameover:
		b, err := msg.(*SSGFreeGameover).Marshal()
		return 3509, 2, b, err
	case *SSGReady:
		b, err := msg.(*SSGReady).Marshal()
		return 3516, 2, b, err
	case *SSGPushOffline:
		b, err := msg.(*SSGPushOffline).Marshal()
		return 3531, 2, b, err
	case *SSGEnterRoom:
		b, err := msg.(*SSGEnterRoom).Marshal()
		return 3512, 2, b, err
	case *SSGGameRecord:
		b, err := msg.(*SSGGameRecord).Marshal()
		return 3523, 2, b, err
	case *SSGPushState:
		b, err := msg.(*SSGPushState).Marshal()
		return 3527, 2, b, err
	case *SSGSit:
		b, err := msg.(*SSGSit).Marshal()
		return 3530, 2, b, err
	case *SSGCamein:
		b, err := msg.(*SSGCamein).Marshal()
		return 3514, 2, b, err
	case *SSGBet:
		b, err := msg.(*SSGBet).Marshal()
		return 3520, 2, b, err
	case *SSGRoomList:
		b, err := msg.(*SSGRoomList).Marshal()
		return 3511, 2, b, err
	case *SSGLeave:
		b, err := msg.(*SSGLeave).Marshal()
		return 3515, 2, b, err
	case *SSGPushDealer:
		b, err := msg.(*SSGPushDealer).Marshal()
		return 3519, 2, b, err
	case *SSGFreeEnterRoom:
		b, err := msg.(*SSGFreeEnterRoom).Marshal()
		return 3502, 2, b, err
	case *SSGFreeDealer:
		b, err := msg.(*SSGFreeDealer).Marshal()
		return 3504, 2, b, err
	case *SSGDraw:
		b, err := msg.(*SSGDraw).Marshal()
		return 3517, 2, b, err
	case *SSGDealer:
		b, err := msg.(*SSGDealer).Marshal()
		return 3518, 2, b, err
	case *SSGiu:
		b, err := msg.(*SSGiu).Marshal()
		return 3521, 2, b, err
	case *SSGFreeCamein:
		b, err := msg.(*SSGFreeCamein).Marshal()
		return 3503, 2, b, err
	case *SSGFreeBet:
		b, err := msg.(*SSGFreeBet).Marshal()
		return 3507, 2, b, err
	case *SSGFreeTrend:
		b, err := msg.(*SSGFreeTrend).Marshal()
		return 3510, 2, b, err
	case *SSGGameover:
		b, err := msg.(*SSGGameover).Marshal()
		return 3522, 2, b, err
	case *SSGFreeRoles:
		b, err := msg.(*SSGFreeRoles).Marshal()
		return 3529, 2, b, err
	case *SSGFreeDealerList:
		b, err := msg.(*SSGFreeDealerList).Marshal()
		return 3505, 2, b, err
	case *SSGLaunchVote:
		b, err := msg.(*SSGLaunchVote).Marshal()
		return 3524, 2, b, err
	case *SSGFreeWiners:
		b, err := msg.(*SSGFreeWiners).Marshal()
		return 3528, 2, b, err
	case *SSGCoinChangeRoom:
		b, err := msg.(*SSGCoinChangeRoom).Marshal()
		return 3532, 2, b, err
	case *SSGCoinGameover:
		b, err := msg.(*SSGCoinGameover).Marshal()
		return 3501, 2, b, err
	case *SSGFreeGamestart:
		b, err := msg.(*SSGFreeGamestart).Marshal()
		return 3508, 2, b, err
	case *SSGCreateRoom:
		b, err := msg.(*SSGCreateRoom).Marshal()
		return 3513, 2, b, err
	case *SSGVote:
		b, err := msg.(*SSGVote).Marshal()
		return 3525, 2, b, err
	case *SSGVoteResult:
		b, err := msg.(*SSGVoteResult).Marshal()
		return 3526, 2, b, err
	case *SSGPushDrawCoin:
		b, err := msg.(*SSGPushDrawCoin).Marshal()
		return 3533, 2, b, err
	case *SJHCoinGameover:
		b, err := msg.(*SJHCoinGameover).Marshal()
		return 4501, 4, b, err
	case *SJHFreeCamein:
		b, err := msg.(*SJHFreeCamein).Marshal()
		return 4503, 4, b, err
	case *SJHFreeBet:
		b, err := msg.(*SJHFreeBet).Marshal()
		return 4507, 4, b, err
	case *SJHPushDealer:
		b, err := msg.(*SJHPushDealer).Marshal()
		return 4519, 4, b, err
	case *SJHGameover:
		b, err := msg.(*SJHGameover).Marshal()
		return 4522, 4, b, err
	case *SJHCoinCall:
		b, err := msg.(*SJHCoinCall).Marshal()
		return 4533, 4, b, err
	case *SJHDraw:
		b, err := msg.(*SJHDraw).Marshal()
		return 4517, 4, b, err
	case *SJHPushDrawCoin:
		b, err := msg.(*SJHPushDrawCoin).Marshal()
		return 4540, 4, b, err
	case *SJHCamein:
		b, err := msg.(*SJHCamein).Marshal()
		return 4514, 4, b, err
	case *SJHReady:
		b, err := msg.(*SJHReady).Marshal()
		return 4516, 4, b, err
	case *SJHCoinChangeRoom:
		b, err := msg.(*SJHCoinChangeRoom).Marshal()
		return 4539, 4, b, err
	case *SJHSit:
		b, err := msg.(*SJHSit).Marshal()
		return 4537, 4, b, err
	case *SJHFreeDealer:
		b, err := msg.(*SJHFreeDealer).Marshal()
		return 4504, 4, b, err
	case *SJHFreeGamestart:
		b, err := msg.(*SJHFreeGamestart).Marshal()
		return 4508, 4, b, err
	case *SJHLaunchVote:
		b, err := msg.(*SJHLaunchVote).Marshal()
		return 4524, 4, b, err
	case *SJHPushActState:
		b, err := msg.(*SJHPushActState).Marshal()
		return 4530, 4, b, err
	case *SJHCoinRaise:
		b, err := msg.(*SJHCoinRaise).Marshal()
		return 4534, 4, b, err
	case *SJHCoinFold:
		b, err := msg.(*SJHCoinFold).Marshal()
		return 4535, 4, b, err
	case *SJHCoinBi:
		b, err := msg.(*SJHCoinBi).Marshal()
		return 4536, 4, b, err
	case *SJHRoomList:
		b, err := msg.(*SJHRoomList).Marshal()
		return 4511, 4, b, err
	case *SJHCreateRoom:
		b, err := msg.(*SJHCreateRoom).Marshal()
		return 4513, 4, b, err
	case *SJHLeave:
		b, err := msg.(*SJHLeave).Marshal()
		return 4515, 4, b, err
	case *SJHVote:
		b, err := msg.(*SJHVote).Marshal()
		return 4525, 4, b, err
	case *SJHPushState:
		b, err := msg.(*SJHPushState).Marshal()
		return 4527, 4, b, err
	case *SJHCoinSee:
		b, err := msg.(*SJHCoinSee).Marshal()
		return 4531, 4, b, err
	case *SJHCoinEnterRoom:
		b, err := msg.(*SJHCoinEnterRoom).Marshal()
		return 4500, 4, b, err
	case *SJHFreeEnterRoom:
		b, err := msg.(*SJHFreeEnterRoom).Marshal()
		return 4502, 4, b, err
	case *SJHFreeDealerList:
		b, err := msg.(*SJHFreeDealerList).Marshal()
		return 4505, 4, b, err
	case *SJHFreeWiners:
		b, err := msg.(*SJHFreeWiners).Marshal()
		return 4528, 4, b, err
	case *SJHPushOffline:
		b, err := msg.(*SJHPushOffline).Marshal()
		return 4538, 4, b, err
	case *SJHFreeRoles:
		b, err := msg.(*SJHFreeRoles).Marshal()
		return 4529, 4, b, err
	case *SJHFreeGameover:
		b, err := msg.(*SJHFreeGameover).Marshal()
		return 4509, 4, b, err
	case *SJHFreeTrend:
		b, err := msg.(*SJHFreeTrend).Marshal()
		return 4510, 4, b, err
	case *SJHEnterRoom:
		b, err := msg.(*SJHEnterRoom).Marshal()
		return 4512, 4, b, err
	case *SJHGameRecord:
		b, err := msg.(*SJHGameRecord).Marshal()
		return 4523, 4, b, err
	case *SJHVoteResult:
		b, err := msg.(*SJHVoteResult).Marshal()
		return 4526, 4, b, err
	default:
		return 0, 0, []byte{}, errors.New("unknown message")
	}
}