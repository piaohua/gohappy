// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: actor_agent.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		actor_agent.proto
		actor_desk.proto
		actor_game.proto
		actor_gate.proto
		actor_logger.proto
		actor_node.proto
		actor_pay.proto
		actor_robot.proto
		actor_role.proto
		actor_web.proto
		game_agent.proto
		game_buy.proto
		game_chat.proto
		game_code.proto
		game_login.proto
		game_pub.proto
		game_type.proto
		game_user.proto
		hua_coin.proto
		hua_free.proto
		hua_pub.proto
		hua_room.proto
		hua_vote.proto
		niu_coin.proto
		niu_free.proto
		niu_pub.proto
		niu_room.proto
		niu_vote.proto
		san_coin.proto
		san_free.proto
		san_pub.proto
		san_room.proto
		san_vote.proto

	It has these top-level messages:
		AgentJoin
		AgentJoined
		MatchDesk
		MatchedDesk
		GenDesk
		GenedDesk
		AddDesk
		AddedDesk
		CloseDesk
		ClosedDesk
		EnterDesk
		EnteredDesk
		JoinDesk
		JoinedDesk
		LeaveDesk
		LeftDesk
		OfflineDesk
		PrintDesk
		SetRecord
		CreateDesk
		CreatedDesk
		GetRoomList
		GotRoomList
		ChangeDesk
		ChangedDesk
		GetGames
		GotGames
		SyncGames
		LoginGate
		LoginedGate
		Login2Gate
		Logined2Gate
		SelectGate
		SelectedGate
		LoginSuccess
		LogoutGate
		LogoutedGate
		OfflineStop
		LogRegist
		LogLogin
		LogLogout
		LogDiamond
		LogCoin
		LogCard
		LogChip
		LogBuildAgency
		LogOnline
		LogExpect
		LogNotice
		LogTask
		Request
		Response
		Connect
		Connected
		Disconnect
		Disconnected
		ServeStart
		ServeStarted
		ServeStop
		ServeStoped
		ServeClose
		Tick
		ApplePay
		ApplePaid
		WxpayCallback
		WxpayGoods
		RobotMsg
		RobotLogin
		RobotReLogin
		RobotLogout
		RobotStop
		RobotRoomList
		RobotEnterRoom
		RobotAllot
		RobotFake
		SetLogin
		SetLogined
		LoginHall
		LoginedHall
		Login
		Logined
		GetUser
		GotUser
		LoginElse
		LoginedElse
		Logout
		Logouted
		SyncUser
		ChangeCurrency
		OfflineCurrency
		PayCurrency
		RoleLogin
		RoleLogined
		RoleRegist
		RoleRegisted
		WxLogin
		WxLogined
		GetUserData
		GotUserData
		SmscodeRegist
		SmscodeRegisted
		RoleBuild
		RoleBuilded
		TouristLogin
		TouristLogined
		BankGive
		BankGiven
		BankCheck
		BankChecked
		BankChange
		TaskUpdate
		LoginPrizeUpdate
		GetRoomRecord
		SyncConfig
		GetConfig
		WebRequest
		WebResponse
		WebRequest2
		WebResponse2
		GetNumber
		GotNumber
		CAgentJoin
		SAgentJoin
		CMyAgent
		SMyAgent
		CAgentManage
		SAgentManage
		AgentManage
		CAgentProfit
		SAgentProfit
		AgentProfitDetail
		CAgentProfitOrder
		SAgentProfitOrder
		AgentProfitOrder
		CAgentProfitApply
		SAgentProfitApply
		CAgentProfitRank
		SAgentProfitRank
		AgentProfit
		CAgentPlayerManage
		SAgentPlayerManage
		AgentPlayerManage
		CAgentPlayerApprove
		SAgentPlayerApprove
		CBuy
		SBuy
		CWxpayOrder
		SWxpayOrder
		CWxpayQuery
		SWxpayQuery
		CApplePay
		SApplePay
		CShop
		SShop
		Shop
		CJtpayOrder
		SJtpayOrder
		CChatText
		SChatText
		CChatVoice
		SChatVoice
		SBroadcast
		CNotice
		SNotice
		SPushNotice
		Notice
		CLogin
		SLogin
		CRegist
		SRegist
		CWxLogin
		SWxLogin
		SLoginOut
		CResetPwd
		SResetPwd
		CTourist
		STourist
		UserData
		Currency
		TopInfo
		Rank
		Task
		LoginPrize
		RoomRecord
		RoomRecordInfo
		RoundRecord
		RoundRoleRecord
		RoleRecord
		CPing
		SPing
		CUserData
		SUserData
		CGetCurrency
		SGetCurrency
		SPushCurrency
		CBank
		SBank
		CRank
		SRank
		CTask
		STask
		CTaskPrize
		STaskPrize
		CLoginPrize
		SLoginPrize
		CRoomRecord
		SRoomRecord
		CSignature
		SSignature
		CLatLng
		SLatLng
		CJHCoinEnterRoom
		SJHCoinEnterRoom
		SJHCoinGameover
		SJHPushActState
		CJHCoinSee
		SJHCoinSee
		CJHCoinCall
		SJHCoinCall
		CJHCoinRaise
		SJHCoinRaise
		CJHCoinFold
		SJHCoinFold
		CJHCoinBi
		SJHCoinBi
		CJHCoinChangeRoom
		SJHCoinChangeRoom
		CJHFreeEnterRoom
		SJHFreeEnterRoom
		SJHFreeCamein
		CJHFreeDealer
		SJHFreeDealer
		CJHFreeDealerList
		SJHFreeDealerList
		CJHSit
		SJHSit
		CJHFreeBet
		SJHFreeBet
		SJHFreeGamestart
		SJHFreeGameover
		CJHFreeTrend
		SJHFreeTrend
		CJHFreeWiners
		SJHFreeWiners
		CJHFreeRoles
		SJHFreeRoles
		JHRoomUser
		JHRoomData
		JHRoomBets
		JHFreeUser
		JHFreeRoom
		JHRoomOver
		JHFreeRoomOver
		JHFreeSeatOver
		JHRoomScore
		JHCoinOver
		JHOverList
		JHRoomVote
		JHDealerList
		JHFreeTrend
		JHFreeWiner
		JHFreeRole
		JHRecordList
		CJHRoomList
		SJHRoomList
		CJHEnterRoom
		SJHEnterRoom
		CJHCreateRoom
		SJHCreateRoom
		SJHCamein
		CJHLeave
		SJHLeave
		SJHPushOffline
		CJHReady
		SJHReady
		SJHDraw
		SJHPushDealer
		SJHPushState
		SJHGameover
		CJHGameRecord
		SJHGameRecord
		CJHLaunchVote
		SJHLaunchVote
		CJHVote
		SJHVote
		SJHVoteResult
		CNNCoinEnterRoom
		SNNCoinEnterRoom
		SNNCoinGameover
		CNNCoinChangeRoom
		SNNCoinChangeRoom
		CNNFreeEnterRoom
		SNNFreeEnterRoom
		SNNFreeCamein
		CNNFreeDealer
		SNNFreeDealer
		CNNFreeDealerList
		SNNFreeDealerList
		CNNSit
		SNNSit
		CNNFreeBet
		SNNFreeBet
		SNNFreeGamestart
		SNNFreeGameover
		CNNFreeTrend
		SNNFreeTrend
		CNNFreeWiners
		SNNFreeWiners
		CNNFreeRoles
		SNNFreeRoles
		NNRoomUser
		NNRoomData
		NNRoomBets
		NNFreeUser
		NNFreeRoom
		NNRoomOver
		NNFreeRoomOver
		NNFreeSeatOver
		NNRoomScore
		NNCoinOver
		NNOverList
		NNRoomVote
		NNDealerList
		NNFreeTrend
		NNFreeWiner
		NNFreeRole
		NNRecordList
		CNNRoomList
		SNNRoomList
		CNNEnterRoom
		SNNEnterRoom
		CNNCreateRoom
		SNNCreateRoom
		SNNCamein
		CNNLeave
		SNNLeave
		SNNPushOffline
		CNNReady
		SNNReady
		SNNDraw
		CNNDealer
		SNNDealer
		SNNPushDealer
		SNNPushState
		CNNBet
		SNNBet
		CNNiu
		SNNiu
		SNNGameover
		CNNGameRecord
		SNNGameRecord
		CNNLaunchVote
		SNNLaunchVote
		CNNVote
		SNNVote
		SNNVoteResult
		CSGCoinEnterRoom
		SSGCoinEnterRoom
		SSGCoinGameover
		CSGCoinChangeRoom
		SSGCoinChangeRoom
		CSGFreeEnterRoom
		SSGFreeEnterRoom
		SSGFreeCamein
		CSGFreeDealer
		SSGFreeDealer
		CSGFreeDealerList
		SSGFreeDealerList
		CSGSit
		SSGSit
		CSGFreeBet
		SSGFreeBet
		SSGFreeGamestart
		SSGFreeGameover
		CSGFreeTrend
		SSGFreeTrend
		CSGFreeWiners
		SSGFreeWiners
		CSGFreeRoles
		SSGFreeRoles
		SGRoomUser
		SGRoomData
		SGRoomBets
		SGFreeUser
		SGFreeRoom
		SGRoomOver
		SGFreeRoomOver
		SGFreeSeatOver
		SGRoomScore
		SGCoinOver
		SGOverList
		SGRoomVote
		SGDealerList
		SGFreeTrend
		SGFreeWiner
		SGFreeRole
		SGRecordList
		CSGRoomList
		SSGRoomList
		CSGEnterRoom
		SSGEnterRoom
		CSGCreateRoom
		SSGCreateRoom
		SSGCamein
		CSGLeave
		SSGLeave
		SSGPushOffline
		CSGReady
		SSGReady
		SSGDraw
		CSGDealer
		SSGDealer
		SSGPushDealer
		SSGPushState
		CSGBet
		SSGBet
		CSGiu
		SSGiu
		SSGGameover
		CSGGameRecord
		SSGGameRecord
		CSGLaunchVote
		SSGLaunchVote
		CSGVote
		SSGVote
		SSGVoteResult
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// 申请加入
type AgentJoin struct {
	Agentname string `protobuf:"bytes,1,opt,name=agentname,proto3" json:"agentname,omitempty"`
	Agentid   string `protobuf:"bytes,2,opt,name=agentid,proto3" json:"agentid,omitempty"`
	Realname  string `protobuf:"bytes,3,opt,name=realname,proto3" json:"realname,omitempty"`
	Weixin    string `protobuf:"bytes,4,opt,name=weixin,proto3" json:"weixin,omitempty"`
}

func (m *AgentJoin) Reset()                    { *m = AgentJoin{} }
func (*AgentJoin) ProtoMessage()               {}
func (*AgentJoin) Descriptor() ([]byte, []int) { return fileDescriptorActorAgent, []int{0} }

func (m *AgentJoin) GetAgentname() string {
	if m != nil {
		return m.Agentname
	}
	return ""
}

func (m *AgentJoin) GetAgentid() string {
	if m != nil {
		return m.Agentid
	}
	return ""
}

func (m *AgentJoin) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *AgentJoin) GetWeixin() string {
	if m != nil {
		return m.Weixin
	}
	return ""
}

type AgentJoined struct {
	Error ErrCode `protobuf:"varint,1,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *AgentJoined) Reset()                    { *m = AgentJoined{} }
func (*AgentJoined) ProtoMessage()               {}
func (*AgentJoined) Descriptor() ([]byte, []int) { return fileDescriptorActorAgent, []int{1} }

func (m *AgentJoined) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

func init() {
	proto.RegisterType((*AgentJoin)(nil), "pb.AgentJoin")
	proto.RegisterType((*AgentJoined)(nil), "pb.AgentJoined")
}
func (this *AgentJoin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AgentJoin)
	if !ok {
		that2, ok := that.(AgentJoin)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Agentname != that1.Agentname {
		return false
	}
	if this.Agentid != that1.Agentid {
		return false
	}
	if this.Realname != that1.Realname {
		return false
	}
	if this.Weixin != that1.Weixin {
		return false
	}
	return true
}
func (this *AgentJoined) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AgentJoined)
	if !ok {
		that2, ok := that.(AgentJoined)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *AgentJoin) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.AgentJoin{")
	s = append(s, "Agentname: "+fmt.Sprintf("%#v", this.Agentname)+",\n")
	s = append(s, "Agentid: "+fmt.Sprintf("%#v", this.Agentid)+",\n")
	s = append(s, "Realname: "+fmt.Sprintf("%#v", this.Realname)+",\n")
	s = append(s, "Weixin: "+fmt.Sprintf("%#v", this.Weixin)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AgentJoined) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.AgentJoined{")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringActorAgent(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AgentJoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgentJoin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Agentname) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorAgent(dAtA, i, uint64(len(m.Agentname)))
		i += copy(dAtA[i:], m.Agentname)
	}
	if len(m.Agentid) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintActorAgent(dAtA, i, uint64(len(m.Agentid)))
		i += copy(dAtA[i:], m.Agentid)
	}
	if len(m.Realname) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintActorAgent(dAtA, i, uint64(len(m.Realname)))
		i += copy(dAtA[i:], m.Realname)
	}
	if len(m.Weixin) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintActorAgent(dAtA, i, uint64(len(m.Weixin)))
		i += copy(dAtA[i:], m.Weixin)
	}
	return i, nil
}

func (m *AgentJoined) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgentJoined) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintActorAgent(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func encodeVarintActorAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AgentJoin) Size() (n int) {
	var l int
	_ = l
	l = len(m.Agentname)
	if l > 0 {
		n += 1 + l + sovActorAgent(uint64(l))
	}
	l = len(m.Agentid)
	if l > 0 {
		n += 1 + l + sovActorAgent(uint64(l))
	}
	l = len(m.Realname)
	if l > 0 {
		n += 1 + l + sovActorAgent(uint64(l))
	}
	l = len(m.Weixin)
	if l > 0 {
		n += 1 + l + sovActorAgent(uint64(l))
	}
	return n
}

func (m *AgentJoined) Size() (n int) {
	var l int
	_ = l
	if m.Error != 0 {
		n += 1 + sovActorAgent(uint64(m.Error))
	}
	return n
}

func sovActorAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozActorAgent(x uint64) (n int) {
	return sovActorAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AgentJoin) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AgentJoin{`,
		`Agentname:` + fmt.Sprintf("%v", this.Agentname) + `,`,
		`Agentid:` + fmt.Sprintf("%v", this.Agentid) + `,`,
		`Realname:` + fmt.Sprintf("%v", this.Realname) + `,`,
		`Weixin:` + fmt.Sprintf("%v", this.Weixin) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AgentJoined) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AgentJoined{`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringActorAgent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AgentJoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AgentJoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgentJoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Agentname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Agentname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Agentid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Agentid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Realname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Realname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weixin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthActorAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Weixin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AgentJoined) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AgentJoined: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgentJoined: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= (ErrCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActorAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipActorAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActorAgent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowActorAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowActorAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthActorAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowActorAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipActorAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthActorAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActorAgent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("actor_agent.proto", fileDescriptorActorAgent) }

var fileDescriptorActorAgent = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4c, 0x2e, 0xc9,
	0x2f, 0x8a, 0x4f, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x92, 0xe2, 0x4f, 0x4f, 0xcc, 0x4d, 0x8d, 0x4f, 0xce, 0x4f, 0x49, 0x85, 0x08, 0x2a, 0x95,
	0x73, 0x71, 0x3a, 0x82, 0xd4, 0x78, 0xe5, 0x67, 0xe6, 0x09, 0xc9, 0x70, 0x71, 0x82, 0x35, 0xe4,
	0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0x24, 0xb8, 0xd8,
	0xc1, 0x9c, 0xcc, 0x14, 0x09, 0x26, 0xb0, 0x1c, 0x8c, 0x2b, 0x24, 0xc5, 0xc5, 0x51, 0x94, 0x9a,
	0x98, 0x03, 0xd6, 0xc6, 0x0c, 0x96, 0x82, 0xf3, 0x85, 0xc4, 0xb8, 0xd8, 0xca, 0x53, 0x33, 0x2b,
	0x32, 0xf3, 0x24, 0x58, 0xc0, 0x32, 0x50, 0x9e, 0x92, 0x01, 0x17, 0x37, 0xdc, 0xe2, 0xd4, 0x14,
	0x21, 0x45, 0x2e, 0xd6, 0xd4, 0xa2, 0xa2, 0xfc, 0x22, 0xb0, 0xb5, 0x7c, 0x46, 0xdc, 0x7a, 0x05,
	0x49, 0x7a, 0xae, 0x45, 0x45, 0xce, 0xf9, 0x29, 0xa9, 0x41, 0x10, 0x19, 0x27, 0x9d, 0x0b, 0x0f,
	0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a,
	0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xf6,
	0x9f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xde, 0x00, 0x3c, 0x25, 0x09, 0x01, 0x00, 0x00,
}