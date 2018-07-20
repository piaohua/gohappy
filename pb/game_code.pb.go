// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game_code.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ErrCode int32

const (
	OK                     ErrCode = 0
	NotEnoughDiamond       ErrCode = 1
	NotEnoughCoin          ErrCode = 2
	NotInRoom              ErrCode = 3
	UsernameOrPwdError     ErrCode = 4
	PhoneNumberError       ErrCode = 5
	LoginError             ErrCode = 6
	UsernameEmpty          ErrCode = 7
	NameTooLong            ErrCode = 8
	PhoneNumberEnpty       ErrCode = 9
	PwdEmpty               ErrCode = 10
	PwdFormatError         ErrCode = 11
	PhoneRegisted          ErrCode = 12
	RegistError            ErrCode = 13
	UserDataNotExist       ErrCode = 14
	WechatLoingFailReAuth  ErrCode = 15
	GetWechatUserInfoFail  ErrCode = 16
	PayOrderFail           ErrCode = 17
	PayOrderError          ErrCode = 18
	RoomNotExist           ErrCode = 19
	RoomFull               ErrCode = 20
	CreateRoomFail         ErrCode = 21
	OperateError           ErrCode = 22
	NiuCardError           ErrCode = 23
	NiuValueError          ErrCode = 24
	BetValueError          ErrCode = 25
	GameStarted            ErrCode = 26
	NotInRoomCannotLeave   ErrCode = 27
	GameStartedCannotLeave ErrCode = 28
	StartedNotKick         ErrCode = 29
	RunningNotVote         ErrCode = 30
	VotingCantLaunchVote   ErrCode = 31
	NotVoteTime            ErrCode = 32
	NotInPrivateRoom       ErrCode = 33
	OtherLoginThisAccount  ErrCode = 34
	BeDealerNotEnough      ErrCode = 35
	SitNotEnough           ErrCode = 36
	SitDownFailed          ErrCode = 37
	BetDealerFailed        ErrCode = 38
	BetNotSeat             ErrCode = 39
	BetTopLimit            ErrCode = 40
	GameNotStart           ErrCode = 41
	StandUpFailed          ErrCode = 42
	DealerSitFailed        ErrCode = 43
	BeDealerAlreadySit     ErrCode = 44
	BeDealerAlready        ErrCode = 45
	DepositNumberError     ErrCode = 46
	DrawMoneyNumberError   ErrCode = 47
	GiveNumberError        ErrCode = 48
	GiveUseridError        ErrCode = 49
	GiveTooMuch            ErrCode = 50
	NotBankrupt            ErrCode = 51
	NotRelieves            ErrCode = 52
	NotPrizeDraw           ErrCode = 53
	NotGotPrizeDraw        ErrCode = 54
	BoxNotYet              ErrCode = 55
	NotBox                 ErrCode = 56
	NotTimes               ErrCode = 57
	AppleOrderFail         ErrCode = 58
	MatchClassicFail       ErrCode = 59
	EnterClassicNotEnough  ErrCode = 60
	NotWinning             ErrCode = 61
	AlreadyWinning         ErrCode = 62
	NotVip                 ErrCode = 63
	NotVipTimes            ErrCode = 64
	AlreadyInRoom          ErrCode = 65
	NotYourTurn            ErrCode = 66
	ErrorOperateValue      ErrCode = 67
	Failed                 ErrCode = 68
	RepeatLogin            ErrCode = 69
	VipTooLow              ErrCode = 70
	ChipNotEnough          ErrCode = 71
	BetSeatWrong           ErrCode = 72
	NotDealerRoom          ErrCode = 73
	SmsCodeEmpty           ErrCode = 74
	SmsCodeWrong           ErrCode = 75
	SmsCodeExpired         ErrCode = 76
	ResetPwdFaild          ErrCode = 77
	PhoneNotRegist         ErrCode = 78
	TouristInoperable      ErrCode = 79
	SafetycodeEmpty        ErrCode = 80
	SafetycodeNotExist     ErrCode = 81
	DealerDownFail         ErrCode = 82
	MatchFail              ErrCode = 83
	EnterFail              ErrCode = 84
	NotReady               ErrCode = 85
	AlreadyFold            ErrCode = 86
	AlreadyAllin           ErrCode = 87
	CallError              ErrCode = 88
	RaiseError             ErrCode = 89
	AlreadyAward           ErrCode = 90
	AwardFaild             ErrCode = 91
	AlreadyPrize           ErrCode = 92
	PwdError               ErrCode = 93
	BankNotOpen            ErrCode = 94
	BankAlreadyOpen        ErrCode = 95
	AlreadySitDown         ErrCode = 96
	SignTooLong            ErrCode = 97
	ChangeFailed           ErrCode = 98
	AlreadyBuild           ErrCode = 99
	ParamError             ErrCode = 100
	AgentNotExist          ErrCode = 101
	AgentLevelLow          ErrCode = 102
	NotAgent               ErrCode = 103
	AlreadyAgent           ErrCode = 104
	WaitForAudit           ErrCode = 105
	ProfitNotEnough        ErrCode = 106
	ProfitOrderNotExist    ErrCode = 107
	ProfitOrderReplied     ErrCode = 108
	ProfitLimit            ErrCode = 109
	AlreadySetRate         ErrCode = 110
	ProfitRateNotEnough    ErrCode = 111
)

var ErrCode_name = map[int32]string{
	0:   "OK",
	1:   "NotEnoughDiamond",
	2:   "NotEnoughCoin",
	3:   "NotInRoom",
	4:   "UsernameOrPwdError",
	5:   "PhoneNumberError",
	6:   "LoginError",
	7:   "UsernameEmpty",
	8:   "NameTooLong",
	9:   "PhoneNumberEnpty",
	10:  "PwdEmpty",
	11:  "PwdFormatError",
	12:  "PhoneRegisted",
	13:  "RegistError",
	14:  "UserDataNotExist",
	15:  "WechatLoingFailReAuth",
	16:  "GetWechatUserInfoFail",
	17:  "PayOrderFail",
	18:  "PayOrderError",
	19:  "RoomNotExist",
	20:  "RoomFull",
	21:  "CreateRoomFail",
	22:  "OperateError",
	23:  "NiuCardError",
	24:  "NiuValueError",
	25:  "BetValueError",
	26:  "GameStarted",
	27:  "NotInRoomCannotLeave",
	28:  "GameStartedCannotLeave",
	29:  "StartedNotKick",
	30:  "RunningNotVote",
	31:  "VotingCantLaunchVote",
	32:  "NotVoteTime",
	33:  "NotInPrivateRoom",
	34:  "OtherLoginThisAccount",
	35:  "BeDealerNotEnough",
	36:  "SitNotEnough",
	37:  "SitDownFailed",
	38:  "BetDealerFailed",
	39:  "BetNotSeat",
	40:  "BetTopLimit",
	41:  "GameNotStart",
	42:  "StandUpFailed",
	43:  "DealerSitFailed",
	44:  "BeDealerAlreadySit",
	45:  "BeDealerAlready",
	46:  "DepositNumberError",
	47:  "DrawMoneyNumberError",
	48:  "GiveNumberError",
	49:  "GiveUseridError",
	50:  "GiveTooMuch",
	51:  "NotBankrupt",
	52:  "NotRelieves",
	53:  "NotPrizeDraw",
	54:  "NotGotPrizeDraw",
	55:  "BoxNotYet",
	56:  "NotBox",
	57:  "NotTimes",
	58:  "AppleOrderFail",
	59:  "MatchClassicFail",
	60:  "EnterClassicNotEnough",
	61:  "NotWinning",
	62:  "AlreadyWinning",
	63:  "NotVip",
	64:  "NotVipTimes",
	65:  "AlreadyInRoom",
	66:  "NotYourTurn",
	67:  "ErrorOperateValue",
	68:  "Failed",
	69:  "RepeatLogin",
	70:  "VipTooLow",
	71:  "ChipNotEnough",
	72:  "BetSeatWrong",
	73:  "NotDealerRoom",
	74:  "SmsCodeEmpty",
	75:  "SmsCodeWrong",
	76:  "SmsCodeExpired",
	77:  "ResetPwdFaild",
	78:  "PhoneNotRegist",
	79:  "TouristInoperable",
	80:  "SafetycodeEmpty",
	81:  "SafetycodeNotExist",
	82:  "DealerDownFail",
	83:  "MatchFail",
	84:  "EnterFail",
	85:  "NotReady",
	86:  "AlreadyFold",
	87:  "AlreadyAllin",
	88:  "CallError",
	89:  "RaiseError",
	90:  "AlreadyAward",
	91:  "AwardFaild",
	92:  "AlreadyPrize",
	93:  "PwdError",
	94:  "BankNotOpen",
	95:  "BankAlreadyOpen",
	96:  "AlreadySitDown",
	97:  "SignTooLong",
	98:  "ChangeFailed",
	99:  "AlreadyBuild",
	100: "ParamError",
	101: "AgentNotExist",
	102: "AgentLevelLow",
	103: "NotAgent",
	104: "AlreadyAgent",
	105: "WaitForAudit",
	106: "ProfitNotEnough",
	107: "ProfitOrderNotExist",
	108: "ProfitOrderReplied",
	109: "ProfitLimit",
	110: "AlreadySetRate",
	111: "ProfitRateNotEnough",
}
var ErrCode_value = map[string]int32{
	"OK":                     0,
	"NotEnoughDiamond":       1,
	"NotEnoughCoin":          2,
	"NotInRoom":              3,
	"UsernameOrPwdError":     4,
	"PhoneNumberError":       5,
	"LoginError":             6,
	"UsernameEmpty":          7,
	"NameTooLong":            8,
	"PhoneNumberEnpty":       9,
	"PwdEmpty":               10,
	"PwdFormatError":         11,
	"PhoneRegisted":          12,
	"RegistError":            13,
	"UserDataNotExist":       14,
	"WechatLoingFailReAuth":  15,
	"GetWechatUserInfoFail":  16,
	"PayOrderFail":           17,
	"PayOrderError":          18,
	"RoomNotExist":           19,
	"RoomFull":               20,
	"CreateRoomFail":         21,
	"OperateError":           22,
	"NiuCardError":           23,
	"NiuValueError":          24,
	"BetValueError":          25,
	"GameStarted":            26,
	"NotInRoomCannotLeave":   27,
	"GameStartedCannotLeave": 28,
	"StartedNotKick":         29,
	"RunningNotVote":         30,
	"VotingCantLaunchVote":   31,
	"NotVoteTime":            32,
	"NotInPrivateRoom":       33,
	"OtherLoginThisAccount":  34,
	"BeDealerNotEnough":      35,
	"SitNotEnough":           36,
	"SitDownFailed":          37,
	"BetDealerFailed":        38,
	"BetNotSeat":             39,
	"BetTopLimit":            40,
	"GameNotStart":           41,
	"StandUpFailed":          42,
	"DealerSitFailed":        43,
	"BeDealerAlreadySit":     44,
	"BeDealerAlready":        45,
	"DepositNumberError":     46,
	"DrawMoneyNumberError":   47,
	"GiveNumberError":        48,
	"GiveUseridError":        49,
	"GiveTooMuch":            50,
	"NotBankrupt":            51,
	"NotRelieves":            52,
	"NotPrizeDraw":           53,
	"NotGotPrizeDraw":        54,
	"BoxNotYet":              55,
	"NotBox":                 56,
	"NotTimes":               57,
	"AppleOrderFail":         58,
	"MatchClassicFail":       59,
	"EnterClassicNotEnough":  60,
	"NotWinning":             61,
	"AlreadyWinning":         62,
	"NotVip":                 63,
	"NotVipTimes":            64,
	"AlreadyInRoom":          65,
	"NotYourTurn":            66,
	"ErrorOperateValue":      67,
	"Failed":                 68,
	"RepeatLogin":            69,
	"VipTooLow":              70,
	"ChipNotEnough":          71,
	"BetSeatWrong":           72,
	"NotDealerRoom":          73,
	"SmsCodeEmpty":           74,
	"SmsCodeWrong":           75,
	"SmsCodeExpired":         76,
	"ResetPwdFaild":          77,
	"PhoneNotRegist":         78,
	"TouristInoperable":      79,
	"SafetycodeEmpty":        80,
	"SafetycodeNotExist":     81,
	"DealerDownFail":         82,
	"MatchFail":              83,
	"EnterFail":              84,
	"NotReady":               85,
	"AlreadyFold":            86,
	"AlreadyAllin":           87,
	"CallError":              88,
	"RaiseError":             89,
	"AlreadyAward":           90,
	"AwardFaild":             91,
	"AlreadyPrize":           92,
	"PwdError":               93,
	"BankNotOpen":            94,
	"BankAlreadyOpen":        95,
	"AlreadySitDown":         96,
	"SignTooLong":            97,
	"ChangeFailed":           98,
	"AlreadyBuild":           99,
	"ParamError":             100,
	"AgentNotExist":          101,
	"AgentLevelLow":          102,
	"NotAgent":               103,
	"AlreadyAgent":           104,
	"WaitForAudit":           105,
	"ProfitNotEnough":        106,
	"ProfitOrderNotExist":    107,
	"ProfitOrderReplied":     108,
	"ProfitLimit":            109,
	"AlreadySetRate":         110,
	"ProfitRateNotEnough":    111,
}

func (ErrCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameCode, []int{0} }

type SitType int32

const (
	SitDown SitType = 0
	SitUp   SitType = 1
)

var SitType_name = map[int32]string{
	0: "SitDown",
	1: "SitUp",
}
var SitType_value = map[string]int32{
	"SitDown": 0,
	"SitUp":   1,
}

func (SitType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameCode, []int{1} }

// 代理玩家管理状态
type AgentApproveState int32

const (
	AgentApprove AgentApproveState = 0
	AgentAgreed  AgentApproveState = 1
	AgentRefused AgentApproveState = 2
)

var AgentApproveState_name = map[int32]string{
	0: "AgentApprove",
	1: "AgentAgreed",
	2: "AgentRefused",
}
var AgentApproveState_value = map[string]int32{
	"AgentApprove": 0,
	"AgentAgreed":  1,
	"AgentRefused": 2,
}

func (AgentApproveState) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameCode, []int{2} }

func init() {
	proto.RegisterEnum("pb.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterEnum("pb.SitType", SitType_name, SitType_value)
	proto.RegisterEnum("pb.AgentApproveState", AgentApproveState_name, AgentApproveState_value)
}
func (x ErrCode) String() string {
	s, ok := ErrCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x SitType) String() string {
	s, ok := SitType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x AgentApproveState) String() string {
	s, ok := AgentApproveState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("game_code.proto", fileDescriptorGameCode) }

var fileDescriptorGameCode = []byte{
	// 1254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x56, 0xd9, 0x7a, 0xdb, 0x54,
	0x10, 0xb6, 0x03, 0x4d, 0x5b, 0x75, 0xc9, 0x44, 0xdd, 0x0b, 0x18, 0xca, 0x4e, 0x28, 0x65, 0x29,
	0xfb, 0xee, 0x25, 0x49, 0x43, 0x1d, 0xd9, 0xd8, 0x4e, 0x42, 0xd9, 0xca, 0x89, 0x35, 0x91, 0x0f,
	0x91, 0xce, 0xd1, 0x77, 0x34, 0x72, 0x62, 0xae, 0x78, 0x04, 0x1e, 0x83, 0x47, 0xe1, 0xb2, 0x97,
	0x5c, 0xb6, 0xe6, 0x86, 0xcb, 0x3e, 0x02, 0xdf, 0x1c, 0x1d, 0xdb, 0x2a, 0x77, 0xd6, 0xaf, 0x99,
	0x7f, 0xb6, 0x7f, 0x46, 0xf6, 0x56, 0x22, 0x91, 0xe0, 0xfd, 0xa1, 0x0e, 0xf1, 0x56, 0x6a, 0x34,
	0x69, 0x7f, 0x29, 0xdd, 0x5f, 0x7b, 0x08, 0xde, 0xc9, 0x75, 0x63, 0x9a, 0x3a, 0x44, 0x7f, 0xd9,
	0x5b, 0xea, 0xdc, 0x85, 0x8a, 0x7f, 0xd1, 0x83, 0x40, 0xd3, 0xba, 0xd2, 0x79, 0x34, 0x6a, 0x49,
	0x91, 0x68, 0x15, 0x42, 0xd5, 0x5f, 0xf5, 0xce, 0xcd, 0xd1, 0xa6, 0x96, 0x0a, 0x96, 0xfc, 0x73,
	0xde, 0xe9, 0x40, 0xd3, 0x96, 0xea, 0x69, 0x9d, 0xc0, 0x53, 0xfe, 0x65, 0xcf, 0xdf, 0xc9, 0xd0,
	0x28, 0x91, 0x60, 0xc7, 0x74, 0x8f, 0xc2, 0x75, 0x63, 0xb4, 0x81, 0xa7, 0x99, 0xaf, 0x3b, 0xd2,
	0x0a, 0x83, 0x3c, 0xd9, 0x47, 0x53, 0xa0, 0x27, 0xfc, 0xf3, 0x9e, 0xd7, 0xd6, 0x91, 0x54, 0xc5,
	0xf3, 0x32, 0xf3, 0xcf, 0xbc, 0xd7, 0x93, 0x94, 0x26, 0x70, 0xd2, 0x5f, 0xf1, 0xce, 0x04, 0x22,
	0xc1, 0x81, 0xd6, 0x6d, 0xad, 0x22, 0x38, 0xf5, 0x7f, 0x26, 0xc5, 0x66, 0xa7, 0xfd, 0xb3, 0xde,
	0x29, 0x8e, 0x66, 0x9d, 0x3c, 0xdf, 0xf7, 0xce, 0x77, 0x8f, 0xc2, 0x0d, 0x6d, 0x12, 0x41, 0x05,
	0xf7, 0x19, 0xe6, 0xb6, 0x7e, 0x3d, 0x8c, 0x64, 0x46, 0x18, 0xc2, 0x59, 0xe6, 0x2e, 0x9e, 0x0a,
	0x9b, 0x73, 0xcc, 0xcd, 0xf1, 0x5b, 0x82, 0x04, 0xd7, 0x79, 0x2c, 0x33, 0x82, 0xf3, 0xfe, 0x35,
	0xef, 0xd2, 0x1e, 0x0e, 0x47, 0x82, 0xda, 0x5a, 0xaa, 0x68, 0x43, 0xc8, 0xb8, 0x87, 0xf5, 0x9c,
	0x46, 0xb0, 0xc2, 0xaf, 0x36, 0x91, 0x8a, 0xb7, 0xec, 0xb9, 0xa5, 0x0e, 0x34, 0x1b, 0x00, 0xf8,
	0xe0, 0x9d, 0xed, 0x8a, 0x49, 0xc7, 0x84, 0x68, 0x2c, 0xb2, 0x6a, 0x33, 0x70, 0x48, 0x11, 0xd0,
	0x67, 0x23, 0x6e, 0xdc, 0x3c, 0xd8, 0x05, 0x2e, 0x84, 0x91, 0x8d, 0x3c, 0x8e, 0xe1, 0x22, 0x17,
	0xd2, 0x34, 0x28, 0x08, 0x2d, 0xc6, 0x34, 0x97, 0xd8, 0xa7, 0x93, 0xa2, 0x11, 0x84, 0x05, 0xcb,
	0x65, 0x46, 0x02, 0x99, 0x37, 0x85, 0x71, 0xed, 0xbe, 0x62, 0x07, 0x25, 0xf3, 0x5d, 0x11, 0xe7,
	0xce, 0xe8, 0x2a, 0x43, 0x0d, 0xa4, 0x12, 0x74, 0x8d, 0xeb, 0xdf, 0x14, 0x09, 0xf6, 0x49, 0x18,
	0x6e, 0xc8, 0x75, 0xff, 0xaa, 0x77, 0x71, 0x3e, 0xcc, 0xa6, 0x50, 0x4a, 0x53, 0x1b, 0xc5, 0x18,
	0xe1, 0x19, 0xff, 0xba, 0x77, 0xb9, 0x64, 0x5a, 0x7e, 0xf7, 0x2c, 0x27, 0xe9, 0xf0, 0x40, 0xd3,
	0x5d, 0x39, 0x3c, 0x84, 0xe7, 0x18, 0xeb, 0xe5, 0x4a, 0x49, 0x15, 0x05, 0x9a, 0x76, 0x35, 0x21,
	0xd4, 0x98, 0x7d, 0x57, 0x93, 0x54, 0x51, 0x53, 0x28, 0x6a, 0x8b, 0x5c, 0x0d, 0x47, 0xf6, 0xcd,
	0xf3, 0x76, 0xc8, 0x85, 0xd9, 0x40, 0x26, 0x08, 0x2f, 0x38, 0xf9, 0x6d, 0xa9, 0xae, 0x91, 0x63,
	0x57, 0x3d, 0xdc, 0xe0, 0x6e, 0x77, 0x68, 0x84, 0xc6, 0x6a, 0x66, 0x30, 0x92, 0x59, 0x7d, 0x38,
	0xd4, 0xb9, 0x22, 0x78, 0xd1, 0xbf, 0xe4, 0xad, 0x36, 0xb0, 0x85, 0x22, 0x46, 0x33, 0x57, 0x28,
	0xbc, 0xc4, 0x9d, 0xe9, 0x4b, 0x5a, 0x20, 0x2f, 0x73, 0x1b, 0xfa, 0x92, 0x5a, 0xfa, 0x48, 0x71,
	0x3b, 0x31, 0x84, 0x57, 0xfc, 0x0b, 0xde, 0x4a, 0x03, 0xa9, 0x70, 0x76, 0xe0, 0xab, 0x2c, 0xcd,
	0x06, 0xb2, 0x67, 0x1f, 0x05, 0xc1, 0x6b, 0x9c, 0x62, 0x03, 0x69, 0xa0, 0xd3, 0xb6, 0x4c, 0x24,
	0xc1, 0xeb, 0x4c, 0xcd, 0x1d, 0x61, 0x0b, 0x2e, 0x1e, 0xde, 0xb0, 0xd4, 0x24, 0x54, 0xb8, 0x93,
	0x3a, 0x96, 0x35, 0xa6, 0x2e, 0x78, 0xfb, 0x92, 0x1c, 0xf8, 0x26, 0xef, 0xc8, 0x2c, 0xd7, 0x7a,
	0x6c, 0x50, 0x84, 0x93, 0xbe, 0x24, 0xb8, 0x59, 0xe4, 0xf1, 0x04, 0x0e, 0x6f, 0xb1, 0x71, 0x0b,
	0x53, 0x9d, 0x49, 0x2a, 0xaf, 0xce, 0x2d, 0x6e, 0x66, 0xcb, 0x88, 0xa3, 0x6d, 0xad, 0x70, 0x52,
	0x7e, 0xf3, 0x36, 0xd3, 0x6c, 0xca, 0xf1, 0x13, 0x9b, 0xf6, 0xce, 0x0c, 0x64, 0x8d, 0x4a, 0xa7,
	0x92, 0x77, 0xed, 0xfc, 0xe5, 0x98, 0x77, 0x6b, 0x3b, 0x1f, 0x8e, 0xe0, 0x3d, 0x37, 0x87, 0x86,
	0x50, 0x87, 0x26, 0x4f, 0x09, 0x6e, 0x3b, 0xa0, 0x87, 0xb1, 0xc4, 0x31, 0x66, 0xf0, 0xbe, 0x95,
	0x9a, 0xa6, 0xae, 0x91, 0xbf, 0x21, 0x87, 0x87, 0x0f, 0x98, 0x39, 0xd0, 0xb4, 0x59, 0x06, 0x3f,
	0xe4, 0xab, 0xd0, 0xd0, 0xc7, 0x81, 0xa6, 0x7b, 0x48, 0xf0, 0x91, 0xef, 0x79, 0xcb, 0xcc, 0xab,
	0x8f, 0xe1, 0x63, 0x16, 0x78, 0xa0, 0x89, 0xe7, 0x9c, 0xc1, 0x27, 0xac, 0x93, 0x7a, 0x9a, 0xc6,
	0xb8, 0xd8, 0x93, 0x4f, 0x79, 0xf8, 0xdb, 0x82, 0x86, 0xa3, 0x66, 0x2c, 0xb2, 0x4c, 0x0e, 0x2d,
	0xfa, 0x19, 0x0f, 0x7f, 0x5d, 0x11, 0x1a, 0x87, 0x2e, 0x66, 0xfa, 0x39, 0xcf, 0x2a, 0xd0, 0xb4,
	0x27, 0xad, 0xde, 0xe0, 0x0b, 0x4b, 0x5a, 0x34, 0x70, 0x86, 0x7d, 0xe9, 0x52, 0xd8, 0x95, 0x29,
	0x7c, 0x35, 0x93, 0x9b, 0x4c, 0x8b, 0x2c, 0xbe, 0xe6, 0xc9, 0x39, 0x07, 0x77, 0xc8, 0xea, 0xce,
	0xe6, 0x9e, 0xce, 0xcd, 0x20, 0x37, 0x0a, 0x1a, 0xac, 0x30, 0xdb, 0x37, 0xb7, 0x7b, 0x76, 0x91,
	0xa0, 0xc9, 0xbc, 0x6e, 0xb0, 0xad, 0xe2, 0x9e, 0xa4, 0xc8, 0x87, 0x22, 0x92, 0x0a, 0xd6, 0xb9,
	0x0d, 0x1c, 0x85, 0x6f, 0xd7, 0x11, 0x6c, 0x70, 0x98, 0xe6, 0x48, 0xa6, 0x8b, 0xd4, 0x37, 0xb9,
	0x9f, 0x0d, 0xb4, 0x1a, 0xdb, 0x33, 0x7c, 0xdf, 0xee, 0xb8, 0x1b, 0x5b, 0xc8, 0xc0, 0xe6, 0xb2,
	0x65, 0x55, 0x9c, 0x64, 0x7c, 0x9f, 0x8b, 0x03, 0xf7, 0x4d, 0x09, 0x29, 0xdc, 0xee, 0xda, 0x25,
	0x74, 0x36, 0xc7, 0xa9, 0x34, 0x18, 0x42, 0x9b, 0xa9, 0x7a, 0x98, 0x21, 0xf1, 0x2d, 0x14, 0x32,
	0x0e, 0x61, 0xdb, 0x5e, 0x46, 0x7b, 0x3d, 0x79, 0xaa, 0x7c, 0xfa, 0x20, 0xe0, 0xca, 0x06, 0x3a,
	0x37, 0x32, 0xa3, 0x2d, 0xa5, 0xb9, 0xba, 0xfd, 0x18, 0xa1, 0xc3, 0x83, 0xed, 0x8b, 0x03, 0xa4,
	0xc9, 0x70, 0x1e, 0xb8, 0xcb, 0x72, 0x5c, 0x80, 0xf3, 0xb3, 0xf5, 0x2d, 0xf3, 0x16, 0x29, 0xcf,
	0x36, 0x0b, 0x7a, 0x5c, 0xbd, 0x9d, 0xa3, 0x7d, 0xec, 0xf3, 0xa3, 0x1d, 0xa0, 0x7d, 0x1c, 0x38,
	0x1d, 0xf4, 0xac, 0xcc, 0x77, 0xb8, 0x75, 0x6e, 0x02, 0x1b, 0x3a, 0x0e, 0x61, 0x97, 0x2b, 0x74,
	0x40, 0x3d, 0x8e, 0xa5, 0x82, 0x3d, 0xf6, 0x6f, 0x8a, 0x38, 0x2e, 0xc4, 0xfb, 0x1d, 0x0f, 0xbd,
	0x27, 0x64, 0xe6, 0x8e, 0xd9, 0xbd, 0xb2, 0xc3, 0x91, 0x30, 0x21, 0x7c, 0xcf, 0x16, 0xf6, 0x67,
	0x51, 0xfb, 0x0f, 0x25, 0x0b, 0x2b, 0x55, 0xf8, 0x71, 0xf6, 0xd5, 0xb0, 0x0c, 0x3f, 0xd9, 0x15,
	0x17, 0xea, 0x30, 0xd0, 0xd4, 0x49, 0x51, 0xc1, 0xcf, 0x76, 0x21, 0x85, 0x3a, 0x74, 0x4e, 0x16,
	0xbc, 0x5f, 0x12, 0x97, 0xbb, 0x23, 0xf0, 0x0b, 0x7b, 0xf6, 0x65, 0xa4, 0x66, 0x1f, 0x29, 0xc1,
	0xa1, 0x9a, 0x23, 0xa1, 0x22, 0x74, 0xda, 0xd8, 0x2f, 0x05, 0x6f, 0xe4, 0x9c, 0xce, 0x90, 0xd3,
	0xeb, 0x0a, 0x23, 0x92, 0x22, 0x7c, 0x68, 0x45, 0x18, 0xa1, 0xa2, 0x79, 0x57, 0x71, 0x0e, 0xb5,
	0x71, 0x8c, 0x31, 0x6b, 0xe8, 0xc0, 0xb5, 0xcd, 0xa2, 0x10, 0x95, 0x8b, 0xb6, 0xc8, 0x88, 0x91,
	0x3d, 0x21, 0x69, 0x43, 0x9b, 0x7a, 0x1e, 0x4a, 0x02, 0xc9, 0x55, 0x74, 0x8d, 0x3e, 0x28, 0x9f,
	0xc1, 0x5f, 0xfd, 0x2b, 0xde, 0x85, 0x02, 0xb4, 0x8b, 0x37, 0x0f, 0x79, 0xc8, 0x03, 0x2e, 0xbd,
	0xe8, 0x61, 0x1a, 0x4b, 0x0c, 0x21, 0xe6, 0x12, 0x0b, 0xbc, 0xb8, 0x7f, 0x49, 0xb9, 0x0f, 0x48,
	0x3d, 0x41, 0x08, 0x6a, 0xc1, 0xca, 0xcf, 0x8b, 0x70, 0x7a, 0xed, 0x86, 0x77, 0xb2, 0x2f, 0x69,
	0x30, 0x49, 0xd1, 0x3f, 0x63, 0x7f, 0xda, 0xc6, 0x55, 0xfc, 0xd3, 0xde, 0x89, 0xbe, 0xa4, 0x9d,
	0x14, 0xaa, 0x6b, 0x77, 0xbc, 0x55, 0x5b, 0x43, 0x3d, 0x4d, 0x8d, 0x1e, 0xf3, 0x97, 0x86, 0xd0,
	0xd6, 0x57, 0x02, 0xa1, 0x62, 0x85, 0x62, 0x91, 0xc8, 0x20, 0xf2, 0x7f, 0x92, 0x99, 0x49, 0x0f,
	0x0f, 0xf2, 0x0c, 0x43, 0x58, 0x6a, 0xdc, 0x7c, 0xf0, 0xa8, 0x56, 0xf9, 0xfb, 0x51, 0xad, 0xf2,
	0xf8, 0x51, 0xad, 0xfa, 0xfb, 0xb4, 0x56, 0xfd, 0x73, 0x5a, 0xab, 0xfe, 0x35, 0xad, 0x55, 0x1f,
	0x4c, 0x6b, 0xd5, 0x87, 0xd3, 0x5a, 0xf5, 0xdf, 0x69, 0xad, 0xf2, 0x78, 0x5a, 0xab, 0xfe, 0xf1,
	0x4f, 0xad, 0xb2, 0xbf, 0x6c, 0xff, 0x08, 0xdd, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xc6,
	0x7b, 0x71, 0x1b, 0x09, 0x00, 0x00,
}
