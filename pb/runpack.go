// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//解包消息
func Runpack(id uint32, b []byte) (interface{}, error) {
	switch id {
	case 1503:
		msg := new(SWxpayQuery)
		err := msg.Unmarshal(b)
		return msg, err
	case 1516:
		msg := new(SUserData)
		err := msg.Unmarshal(b)
		return msg, err
	case 1522:
		msg := new(STask)
		err := msg.Unmarshal(b)
		return msg, err
	case 1528:
		msg := new(SLatLng)
		err := msg.Unmarshal(b)
		return msg, err
	case 1504:
		msg := new(SApplePay)
		err := msg.Unmarshal(b)
		return msg, err
	case 1514:
		msg := new(SResetPwd)
		err := msg.Unmarshal(b)
		return msg, err
	case 1515:
		msg := new(STourist)
		err := msg.Unmarshal(b)
		return msg, err
	case 1517:
		msg := new(SGetCurrency)
		err := msg.Unmarshal(b)
		return msg, err
	case 1502:
		msg := new(SJtpayOrder)
		err := msg.Unmarshal(b)
		return msg, err
	case 1512:
		msg := new(SWxLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1529:
		msg := new(SBankLog)
		err := msg.Unmarshal(b)
		return msg, err
	case 1556:
		msg := new(SAgentProfitRank)
		err := msg.Unmarshal(b)
		return msg, err
	case 1501:
		msg := new(SWxpayOrder)
		err := msg.Unmarshal(b)
		return msg, err
	case 1507:
		msg := new(SChatVoice)
		err := msg.Unmarshal(b)
		return msg, err
	case 1520:
		msg := new(SBank)
		err := msg.Unmarshal(b)
		return msg, err
	case 1524:
		msg := new(SLoginPrize)
		err := msg.Unmarshal(b)
		return msg, err
	case 1526:
		msg := new(SPushNotice)
		err := msg.Unmarshal(b)
		return msg, err
	case 1553:
		msg := new(SAgentProfit)
		err := msg.Unmarshal(b)
		return msg, err
	case 1559:
		msg := new(SAgentProfitReply)
		err := msg.Unmarshal(b)
		return msg, err
	case 1508:
		msg := new(SBroadcast)
		err := msg.Unmarshal(b)
		return msg, err
	case 1510:
		msg := new(SLogin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1519:
		msg := new(SPing)
		err := msg.Unmarshal(b)
		return msg, err
	case 1551:
		msg := new(SMyAgent)
		err := msg.Unmarshal(b)
		return msg, err
	case 1557:
		msg := new(SAgentPlayerManage)
		err := msg.Unmarshal(b)
		return msg, err
	case 1558:
		msg := new(SAgentPlayerApprove)
		err := msg.Unmarshal(b)
		return msg, err
	case 1511:
		msg := new(SRegist)
		err := msg.Unmarshal(b)
		return msg, err
	case 1513:
		msg := new(SLoginOut)
		err := msg.Unmarshal(b)
		return msg, err
	case 1521:
		msg := new(SRank)
		err := msg.Unmarshal(b)
		return msg, err
	case 1523:
		msg := new(STaskPrize)
		err := msg.Unmarshal(b)
		return msg, err
	case 1555:
		msg := new(SAgentProfitApply)
		err := msg.Unmarshal(b)
		return msg, err
	case 1509:
		msg := new(SNotice)
		err := msg.Unmarshal(b)
		return msg, err
	case 1518:
		msg := new(SPushCurrency)
		err := msg.Unmarshal(b)
		return msg, err
	case 1550:
		msg := new(SAgentJoin)
		err := msg.Unmarshal(b)
		return msg, err
	case 1554:
		msg := new(SAgentProfitOrder)
		err := msg.Unmarshal(b)
		return msg, err
	case 1500:
		msg := new(SBuy)
		err := msg.Unmarshal(b)
		return msg, err
	case 1505:
		msg := new(SShop)
		err := msg.Unmarshal(b)
		return msg, err
	case 1506:
		msg := new(SChatText)
		err := msg.Unmarshal(b)
		return msg, err
	case 1525:
		msg := new(SRoomRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 1527:
		msg := new(SSignature)
		err := msg.Unmarshal(b)
		return msg, err
	case 1552:
		msg := new(SAgentManage)
		err := msg.Unmarshal(b)
		return msg, err
	case 2510:
		msg := new(SNNFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 2521:
		msg := new(SNNiu)
		err := msg.Unmarshal(b)
		return msg, err
	case 2522:
		msg := new(SNNGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 2523:
		msg := new(SNNGameRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 2524:
		msg := new(SNNLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 2507:
		msg := new(SNNFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 2508:
		msg := new(SNNFreeGamestart)
		err := msg.Unmarshal(b)
		return msg, err
	case 2509:
		msg := new(SNNFreeGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 2518:
		msg := new(SNNDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 2519:
		msg := new(SNNPushDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 2505:
		msg := new(SNNFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 2515:
		msg := new(SNNLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 2531:
		msg := new(SNNPushOffline)
		err := msg.Unmarshal(b)
		return msg, err
	case 2525:
		msg := new(SNNVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 2528:
		msg := new(SNNFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 2533:
		msg := new(SNNPushDrawCoin)
		err := msg.Unmarshal(b)
		return msg, err
	case 2530:
		msg := new(SNNSit)
		err := msg.Unmarshal(b)
		return msg, err
	case 2500:
		msg := new(SNNCoinEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2501:
		msg := new(SNNCoinGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 2503:
		msg := new(SNNFreeCamein)
		err := msg.Unmarshal(b)
		return msg, err
	case 2512:
		msg := new(SNNEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2529:
		msg := new(SNNFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 2502:
		msg := new(SNNFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2514:
		msg := new(SNNCamein)
		err := msg.Unmarshal(b)
		return msg, err
	case 2517:
		msg := new(SNNDraw)
		err := msg.Unmarshal(b)
		return msg, err
	case 2527:
		msg := new(SNNPushState)
		err := msg.Unmarshal(b)
		return msg, err
	case 2532:
		msg := new(SNNCoinChangeRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2504:
		msg := new(SNNFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 2511:
		msg := new(SNNRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 2520:
		msg := new(SNNBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 2513:
		msg := new(SNNCreateRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 2516:
		msg := new(SNNReady)
		err := msg.Unmarshal(b)
		return msg, err
	case 2526:
		msg := new(SNNVoteResult)
		err := msg.Unmarshal(b)
		return msg, err
	case 3501:
		msg := new(SSGCoinGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 3527:
		msg := new(SSGPushState)
		err := msg.Unmarshal(b)
		return msg, err
	case 3528:
		msg := new(SSGFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 3533:
		msg := new(SSGPushDrawCoin)
		err := msg.Unmarshal(b)
		return msg, err
	case 3509:
		msg := new(SSGFreeGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 3511:
		msg := new(SSGRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 3514:
		msg := new(SSGCamein)
		err := msg.Unmarshal(b)
		return msg, err
	case 3517:
		msg := new(SSGDraw)
		err := msg.Unmarshal(b)
		return msg, err
	case 3518:
		msg := new(SSGDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 3507:
		msg := new(SSGFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 3508:
		msg := new(SSGFreeGamestart)
		err := msg.Unmarshal(b)
		return msg, err
	case 3516:
		msg := new(SSGReady)
		err := msg.Unmarshal(b)
		return msg, err
	case 3524:
		msg := new(SSGLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 3510:
		msg := new(SSGFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 3502:
		msg := new(SSGFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3525:
		msg := new(SSGVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 3526:
		msg := new(SSGVoteResult)
		err := msg.Unmarshal(b)
		return msg, err
	case 3532:
		msg := new(SSGCoinChangeRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3512:
		msg := new(SSGEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3515:
		msg := new(SSGLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 3522:
		msg := new(SSGGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 3523:
		msg := new(SSGGameRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 3529:
		msg := new(SSGFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 3505:
		msg := new(SSGFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 3519:
		msg := new(SSGPushDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 3520:
		msg := new(SSGBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 3530:
		msg := new(SSGSit)
		err := msg.Unmarshal(b)
		return msg, err
	case 3531:
		msg := new(SSGPushOffline)
		err := msg.Unmarshal(b)
		return msg, err
	case 3500:
		msg := new(SSGCoinEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3503:
		msg := new(SSGFreeCamein)
		err := msg.Unmarshal(b)
		return msg, err
	case 3504:
		msg := new(SSGFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 3513:
		msg := new(SSGCreateRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 3521:
		msg := new(SSGiu)
		err := msg.Unmarshal(b)
		return msg, err
	case 4536:
		msg := new(SJHCoinBi)
		err := msg.Unmarshal(b)
		return msg, err
	case 4501:
		msg := new(SJHCoinGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 4503:
		msg := new(SJHFreeCamein)
		err := msg.Unmarshal(b)
		return msg, err
	case 4514:
		msg := new(SJHCamein)
		err := msg.Unmarshal(b)
		return msg, err
	case 4524:
		msg := new(SJHLaunchVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 4531:
		msg := new(SJHCoinSee)
		err := msg.Unmarshal(b)
		return msg, err
	case 4523:
		msg := new(SJHGameRecord)
		err := msg.Unmarshal(b)
		return msg, err
	case 4526:
		msg := new(SJHVoteResult)
		err := msg.Unmarshal(b)
		return msg, err
	case 4529:
		msg := new(SJHFreeRoles)
		err := msg.Unmarshal(b)
		return msg, err
	case 4533:
		msg := new(SJHCoinCall)
		err := msg.Unmarshal(b)
		return msg, err
	case 4504:
		msg := new(SJHFreeDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 4505:
		msg := new(SJHFreeDealerList)
		err := msg.Unmarshal(b)
		return msg, err
	case 4510:
		msg := new(SJHFreeTrend)
		err := msg.Unmarshal(b)
		return msg, err
	case 4519:
		msg := new(SJHPushDealer)
		err := msg.Unmarshal(b)
		return msg, err
	case 4525:
		msg := new(SJHVote)
		err := msg.Unmarshal(b)
		return msg, err
	case 4540:
		msg := new(SJHPushDrawCoin)
		err := msg.Unmarshal(b)
		return msg, err
	case 4535:
		msg := new(SJHCoinFold)
		err := msg.Unmarshal(b)
		return msg, err
	case 4538:
		msg := new(SJHPushOffline)
		err := msg.Unmarshal(b)
		return msg, err
	case 4507:
		msg := new(SJHFreeBet)
		err := msg.Unmarshal(b)
		return msg, err
	case 4512:
		msg := new(SJHEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4513:
		msg := new(SJHCreateRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4522:
		msg := new(SJHGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 4528:
		msg := new(SJHFreeWiners)
		err := msg.Unmarshal(b)
		return msg, err
	case 4508:
		msg := new(SJHFreeGamestart)
		err := msg.Unmarshal(b)
		return msg, err
	case 4539:
		msg := new(SJHCoinChangeRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4516:
		msg := new(SJHReady)
		err := msg.Unmarshal(b)
		return msg, err
	case 4527:
		msg := new(SJHPushState)
		err := msg.Unmarshal(b)
		return msg, err
	case 4537:
		msg := new(SJHSit)
		err := msg.Unmarshal(b)
		return msg, err
	case 4500:
		msg := new(SJHCoinEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4502:
		msg := new(SJHFreeEnterRoom)
		err := msg.Unmarshal(b)
		return msg, err
	case 4517:
		msg := new(SJHDraw)
		err := msg.Unmarshal(b)
		return msg, err
	case 4530:
		msg := new(SJHPushActState)
		err := msg.Unmarshal(b)
		return msg, err
	case 4509:
		msg := new(SJHFreeGameover)
		err := msg.Unmarshal(b)
		return msg, err
	case 4511:
		msg := new(SJHRoomList)
		err := msg.Unmarshal(b)
		return msg, err
	case 4515:
		msg := new(SJHLeave)
		err := msg.Unmarshal(b)
		return msg, err
	case 4534:
		msg := new(SJHCoinRaise)
		err := msg.Unmarshal(b)
		return msg, err
	default:
		return nil, errors.New("unknown message")
	}
}