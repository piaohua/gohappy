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
	// 1249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x56, 0xc9, 0x7a, 0x1b, 0x45,
	0x10, 0x96, 0x0c, 0x71, 0x92, 0xc9, 0xe2, 0xf2, 0x64, 0x0f, 0x20, 0x08, 0x3b, 0x26, 0x84, 0x25,
	0xec, 0xbb, 0x16, 0xdb, 0x31, 0x91, 0x25, 0x21, 0xc9, 0x36, 0x61, 0x0b, 0x6d, 0x4d, 0x79, 0xd4,
	0x78, 0xa6, 0x7b, 0xbe, 0x9e, 0x1a, 0xd9, 0xe6, 0xc4, 0x23, 0xf0, 0x18, 0xbc, 0x09, 0x1c, 0x73,
	0xe4, 0x48, 0xc4, 0x85, 0x63, 0x1e, 0x81, 0xaf, 0xba, 0x5b, 0xf2, 0x84, 0x9b, 0xe6, 0x9f, 0xaa,
	0xbf, 0xb6, 0xbf, 0x6a, 0x14, 0x2c, 0xc5, 0x22, 0xc5, 0xfb, 0x23, 0x1d, 0xe1, 0xad, 0xcc, 0x68,
	0xd2, 0xe1, 0x42, 0xb6, 0xbb, 0xf2, 0x07, 0x04, 0x27, 0x57, 0x8d, 0x69, 0xea, 0x08, 0xc3, 0xc5,
	0x60, 0xa1, 0x7b, 0x17, 0x2a, 0xe1, 0xc5, 0x00, 0x3a, 0x9a, 0x56, 0x95, 0x2e, 0xe2, 0x71, 0x4b,
	0x8a, 0x54, 0xab, 0x08, 0xaa, 0xe1, 0x72, 0x70, 0x6e, 0x8e, 0x36, 0xb5, 0x54, 0xb0, 0x10, 0x9e,
	0x0b, 0x4e, 0x77, 0x34, 0x6d, 0xa8, 0xbe, 0xd6, 0x29, 0x3c, 0x11, 0x5e, 0x0e, 0xc2, 0xad, 0x1c,
	0x8d, 0x12, 0x29, 0x76, 0x4d, 0xef, 0x20, 0x5a, 0x35, 0x46, 0x1b, 0x78, 0x92, 0xf9, 0x7a, 0x63,
	0xad, 0xb0, 0x53, 0xa4, 0xbb, 0x68, 0x1c, 0x7a, 0x22, 0x3c, 0x1f, 0x04, 0x6d, 0x1d, 0x4b, 0xe5,
	0x9e, 0x17, 0x99, 0x7f, 0xe6, 0xbd, 0x9a, 0x66, 0x74, 0x04, 0x27, 0xc3, 0xa5, 0xe0, 0x4c, 0x47,
	0xa4, 0x38, 0xd4, 0xba, 0xad, 0x55, 0x0c, 0xa7, 0xfe, 0xcf, 0xa4, 0xd8, 0xec, 0x74, 0x78, 0x36,
	0x38, 0xc5, 0xd1, 0xac, 0x53, 0x10, 0x86, 0xc1, 0xf9, 0xde, 0x41, 0xb4, 0xa6, 0x4d, 0x2a, 0xc8,
	0x71, 0x9f, 0x61, 0x6e, 0xeb, 0xd7, 0xc7, 0x58, 0xe6, 0x84, 0x11, 0x9c, 0x65, 0x6e, 0xf7, 0xe4,
	0x6c, 0xce, 0x31, 0x37, 0xc7, 0x6f, 0x09, 0x12, 0x5c, 0xe7, 0xa1, 0xcc, 0x09, 0xce, 0x87, 0xd7,
	0x82, 0x4b, 0x3b, 0x38, 0x1a, 0x0b, 0x6a, 0x6b, 0xa9, 0xe2, 0x35, 0x21, 0x93, 0x3e, 0xd6, 0x0b,
	0x1a, 0xc3, 0x12, 0xbf, 0x5a, 0x47, 0x72, 0x6f, 0xd9, 0x73, 0x43, 0xed, 0x69, 0x36, 0x00, 0x08,
	0x21, 0x38, 0xdb, 0x13, 0x47, 0x5d, 0x13, 0xa1, 0xb1, 0xc8, 0xb2, 0xcd, 0xc0, 0x23, 0x2e, 0x60,
	0xc8, 0x46, 0xdc, 0xb8, 0x79, 0xb0, 0x0b, 0x5c, 0x08, 0x23, 0x6b, 0x45, 0x92, 0xc0, 0x45, 0x2e,
	0xa4, 0x69, 0x50, 0x10, 0x5a, 0x8c, 0x69, 0x2e, 0xb1, 0x4f, 0x37, 0x43, 0x23, 0x08, 0x1d, 0xcb,
	0x65, 0x46, 0x3a, 0xb2, 0x68, 0x0a, 0xe3, 0xdb, 0x7d, 0xc5, 0x0e, 0x4a, 0x16, 0xdb, 0x22, 0x29,
	0xbc, 0xd1, 0x55, 0x86, 0x1a, 0x48, 0x25, 0xe8, 0x1a, 0xd7, 0xbf, 0x2e, 0x52, 0x1c, 0x90, 0x30,
	0xdc, 0x90, 0xeb, 0xe1, 0xd5, 0xe0, 0xe2, 0x7c, 0x98, 0x4d, 0xa1, 0x94, 0xa6, 0x36, 0x8a, 0x09,
	0xc2, 0x53, 0xe1, 0xf5, 0xe0, 0x72, 0xc9, 0xb4, 0xfc, 0xee, 0x69, 0x4e, 0xd2, 0xe3, 0x1d, 0x4d,
	0x77, 0xe5, 0x68, 0x1f, 0x9e, 0x61, 0xac, 0x5f, 0x28, 0x25, 0x55, 0xdc, 0xd1, 0xb4, 0xad, 0x09,
	0xa1, 0xc6, 0xec, 0xdb, 0x9a, 0xa4, 0x8a, 0x9b, 0x42, 0x51, 0x5b, 0x14, 0x6a, 0x34, 0xb6, 0x6f,
	0x9e, 0xb5, 0x43, 0x76, 0x66, 0x43, 0x99, 0x22, 0x3c, 0xe7, 0xe5, 0xb7, 0xa1, 0x7a, 0x46, 0x4e,
	0x7c, 0xf5, 0x70, 0x83, 0xbb, 0xdd, 0xa5, 0x31, 0x1a, 0xab, 0x99, 0xe1, 0x58, 0xe6, 0xf5, 0xd1,
	0x48, 0x17, 0x8a, 0xe0, 0xf9, 0xf0, 0x52, 0xb0, 0xdc, 0xc0, 0x16, 0x8a, 0x04, 0xcd, 0x5c, 0xa1,
	0xf0, 0x02, 0x77, 0x66, 0x20, 0xe9, 0x18, 0x79, 0x91, 0xdb, 0x30, 0x90, 0xd4, 0xd2, 0x07, 0x8a,
	0xdb, 0x89, 0x11, 0xbc, 0x14, 0x5e, 0x08, 0x96, 0x1a, 0x48, 0xce, 0xd9, 0x83, 0x2f, 0xb3, 0x34,
	0x1b, 0xc8, 0x9e, 0x03, 0x14, 0x04, 0xaf, 0x70, 0x8a, 0x0d, 0xa4, 0xa1, 0xce, 0xda, 0x32, 0x95,
	0x04, 0xaf, 0x32, 0x35, 0x77, 0x84, 0x2d, 0xb8, 0x78, 0x78, 0xcd, 0x52, 0x93, 0x50, 0xd1, 0x56,
	0xe6, 0x59, 0x56, 0x98, 0xda, 0xf1, 0x0e, 0x24, 0x79, 0xf0, 0x75, 0xde, 0x91, 0x59, 0xae, 0xf5,
	0xc4, 0xa0, 0x88, 0x8e, 0x06, 0x92, 0xe0, 0xa6, 0xcb, 0xe3, 0x31, 0x1c, 0xde, 0x60, 0xe3, 0x16,
	0x66, 0x3a, 0x97, 0x54, 0x5e, 0x9d, 0x5b, 0xdc, 0xcc, 0x96, 0x11, 0x07, 0x9b, 0x5a, 0xe1, 0x51,
	0xf9, 0xcd, 0x9b, 0x4c, 0xb3, 0x2e, 0x27, 0x8f, 0x6d, 0xda, 0x5b, 0x33, 0x90, 0x35, 0x2a, 0xbd,
	0x4a, 0xde, 0xb6, 0xf3, 0x97, 0x13, 0xde, 0xad, 0xcd, 0x62, 0x34, 0x86, 0x77, 0xfc, 0x1c, 0x1a,
	0x42, 0xed, 0x9b, 0x22, 0x23, 0xb8, 0xed, 0x81, 0x3e, 0x26, 0x12, 0x27, 0x98, 0xc3, 0xbb, 0x56,
	0x6a, 0x9a, 0x7a, 0x46, 0xfe, 0x82, 0x1c, 0x1e, 0xde, 0x63, 0xe6, 0x8e, 0xa6, 0xf5, 0x32, 0xf8,
	0x3e, 0x5f, 0x85, 0x86, 0x3e, 0xec, 0x68, 0xba, 0x87, 0x04, 0x1f, 0x84, 0x41, 0xb0, 0xc8, 0xbc,
	0xfa, 0x10, 0x3e, 0x64, 0x81, 0x77, 0x34, 0xf1, 0x9c, 0x73, 0xf8, 0x88, 0x75, 0x52, 0xcf, 0xb2,
	0x04, 0x8f, 0xf7, 0xe4, 0x63, 0x1e, 0xfe, 0xa6, 0xa0, 0xd1, 0xb8, 0x99, 0x88, 0x3c, 0x97, 0x23,
	0x8b, 0x7e, 0xc2, 0xc3, 0x5f, 0x55, 0x84, 0xc6, 0xa3, 0xc7, 0x33, 0xfd, 0x94, 0x67, 0xd5, 0xd1,
	0xb4, 0x23, 0xad, 0xde, 0xe0, 0x33, 0x4b, 0xea, 0x1a, 0x38, 0xc3, 0x3e, 0xf7, 0x29, 0x6c, 0xcb,
	0x0c, 0xbe, 0x98, 0xc9, 0x4d, 0x66, 0x2e, 0x8b, 0x2f, 0x79, 0x72, 0xde, 0xc1, 0x1f, 0xb2, 0xba,
	0xb7, 0xb9, 0xa7, 0x0b, 0x33, 0x2c, 0x8c, 0x82, 0x06, 0x2b, 0xcc, 0xf6, 0xcd, 0xef, 0x9e, 0x5d,
	0x24, 0x68, 0x32, 0xaf, 0x1f, 0x6c, 0xcb, 0xdd, 0x93, 0x0c, 0xf9, 0x50, 0xc4, 0x52, 0xc1, 0x2a,
	0xb7, 0x81, 0xa3, 0xf0, 0xed, 0x3a, 0x80, 0x35, 0x0e, 0xd3, 0x1c, 0xcb, 0xec, 0x38, 0xf5, 0x75,
	0xee, 0x67, 0x03, 0xad, 0xc6, 0x76, 0x0c, 0xdf, 0xb7, 0x3b, 0xfe, 0xc6, 0x3a, 0x19, 0xd8, 0x5c,
	0x36, 0xac, 0x8a, 0xd3, 0x9c, 0xef, 0xb3, 0x3b, 0x70, 0x5f, 0x95, 0x10, 0xe7, 0x76, 0xd7, 0x2e,
	0xa1, 0xb7, 0x39, 0xcc, 0xa4, 0xc1, 0x08, 0xda, 0x4c, 0xd5, 0xc7, 0x1c, 0x89, 0x6f, 0xa1, 0x90,
	0x49, 0x04, 0x9b, 0xf6, 0x32, 0xda, 0xeb, 0xc9, 0x53, 0xe5, 0xd3, 0x07, 0x1d, 0xae, 0x6c, 0xa8,
	0x0b, 0x23, 0x73, 0xda, 0x50, 0x9a, 0xab, 0xdb, 0x4d, 0x10, 0xba, 0x3c, 0xd8, 0x81, 0xd8, 0x43,
	0x3a, 0x1a, 0xcd, 0x03, 0xf7, 0x58, 0x8e, 0xc7, 0xe0, 0xfc, 0x6c, 0x7d, 0xcd, 0xbc, 0x2e, 0xe5,
	0xd9, 0x66, 0x41, 0x9f, 0xab, 0xb7, 0x73, 0xb4, 0x8f, 0x03, 0x7e, 0xb4, 0x03, 0xb4, 0x8f, 0x43,
	0xaf, 0x83, 0xbe, 0x95, 0xf9, 0x16, 0xb7, 0xce, 0x4f, 0x60, 0x4d, 0x27, 0x11, 0x6c, 0x73, 0x85,
	0x1e, 0xa8, 0x27, 0x89, 0x54, 0xb0, 0xc3, 0xfe, 0x4d, 0x91, 0x24, 0x4e, 0xbc, 0xdf, 0xf0, 0xd0,
	0xfb, 0x42, 0xe6, 0xfe, 0x98, 0xdd, 0x2b, 0x3b, 0x1c, 0x08, 0x13, 0xc1, 0xb7, 0x6c, 0x61, 0x7f,
	0xba, 0xda, 0xbf, 0x2b, 0x59, 0x58, 0xa9, 0xc2, 0xf7, 0xb3, 0xaf, 0x86, 0x65, 0xf8, 0xc1, 0xae,
	0xb8, 0x50, 0xfb, 0x1d, 0x4d, 0xdd, 0x0c, 0x15, 0xfc, 0x68, 0x17, 0x52, 0xa8, 0x7d, 0xef, 0x64,
	0xc1, 0xfb, 0x25, 0x71, 0xf9, 0x3b, 0x02, 0x3f, 0xb1, 0xe7, 0x40, 0xc6, 0x6a, 0xf6, 0x91, 0x12,
	0x1c, 0xaa, 0x39, 0x16, 0x2a, 0x46, 0xaf, 0x8d, 0xdd, 0x52, 0xf0, 0x46, 0xc1, 0xe9, 0x8c, 0x38,
	0xbd, 0x9e, 0x30, 0x22, 0x75, 0xe1, 0x23, 0x2b, 0xc2, 0x18, 0x15, 0xcd, 0xbb, 0x8a, 0x73, 0xa8,
	0x8d, 0x13, 0x4c, 0x58, 0x43, 0x7b, 0xbe, 0x6d, 0x16, 0x85, 0xb8, 0x5c, 0xb4, 0x45, 0xc6, 0x8c,
	0xec, 0x08, 0x49, 0x6b, 0xda, 0xd4, 0x8b, 0x48, 0x12, 0x48, 0xae, 0xa2, 0x67, 0xf4, 0x5e, 0xf9,
	0x0c, 0xfe, 0x1c, 0x5e, 0x09, 0x2e, 0x38, 0xd0, 0x2e, 0xde, 0x3c, 0xe4, 0x3e, 0x0f, 0xb8, 0xf4,
	0xa2, 0x8f, 0x59, 0x22, 0x31, 0x82, 0x84, 0x4b, 0x74, 0xb8, 0xbb, 0x7f, 0x69, 0xb9, 0x0f, 0x48,
	0x7d, 0x41, 0x08, 0x6a, 0xe5, 0x46, 0x70, 0x72, 0x20, 0x69, 0x78, 0x94, 0x61, 0x78, 0xc6, 0xfe,
	0xb4, 0xfd, 0xa9, 0x84, 0xa7, 0x83, 0x13, 0x03, 0x49, 0x5b, 0x19, 0x54, 0x57, 0xee, 0x04, 0xcb,
	0x36, 0xd5, 0x7a, 0x96, 0x19, 0x3d, 0xe1, 0x0f, 0x0a, 0xa1, 0x2d, 0xa3, 0x04, 0x42, 0xc5, 0xea,
	0xc1, 0x22, 0xb1, 0x41, 0xe4, 0xbf, 0x1e, 0x33, 0x93, 0x3e, 0xee, 0x15, 0x39, 0x46, 0xb0, 0xd0,
	0xb8, 0xf9, 0xe0, 0x61, 0xad, 0xf2, 0xd7, 0xc3, 0x5a, 0xe5, 0xd1, 0xc3, 0x5a, 0xf5, 0xd7, 0x69,
	0xad, 0xfa, 0xfb, 0xb4, 0x56, 0xfd, 0x73, 0x5a, 0xab, 0x3e, 0x98, 0xd6, 0xaa, 0x7f, 0x4f, 0x6b,
	0xd5, 0x7f, 0xa7, 0xb5, 0xca, 0xa3, 0x69, 0xad, 0xfa, 0xdb, 0x3f, 0xb5, 0xca, 0xee, 0xa2, 0xfd,
	0xbf, 0x73, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xa2, 0x55, 0x45, 0x02, 0x09, 0x00,
	0x00,
}
