// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game_type.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 协议文件名称格式:
// {gametype}_xxx.proto
type GameType int32

const (
	GAME GameType = 0
	NIU  GameType = 1
	SAN  GameType = 2
	DOU  GameType = 3
	HUA  GameType = 4
	PAI  GameType = 5
	EBG  GameType = 6
)

var GameType_name = map[int32]string{
	0: "GAME",
	1: "NIU",
	2: "SAN",
	3: "DOU",
	4: "HUA",
	5: "PAI",
	6: "EBG",
}
var GameType_value = map[string]int32{
	"GAME": 0,
	"NIU":  1,
	"SAN":  2,
	"DOU":  3,
	"HUA":  4,
	"PAI":  5,
	"EBG":  6,
}

func (GameType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{0} }

// 桌子状态
type DeskState int32

const (
	STATE_READY  DeskState = 0
	STATE_DEALER DeskState = 1
	STATE_NIU    DeskState = 2
	STATE_BET    DeskState = 3
	STATE_OVER   DeskState = 4
)

var DeskState_name = map[int32]string{
	0: "STATE_READY",
	1: "STATE_DEALER",
	2: "STATE_NIU",
	3: "STATE_BET",
	4: "STATE_OVER",
}
var DeskState_value = map[string]int32{
	"STATE_READY":  0,
	"STATE_DEALER": 1,
	"STATE_NIU":    2,
	"STATE_BET":    3,
	"STATE_OVER":   4,
}

func (DeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{1} }

// 房间类型
type RoomType int32

const (
	ROOM_TYPE0 RoomType = 0
	ROOM_TYPE1 RoomType = 1
	ROOM_TYPE2 RoomType = 2
)

var RoomType_name = map[int32]string{
	0: "ROOM_TYPE0",
	1: "ROOM_TYPE1",
	2: "ROOM_TYPE2",
}
var RoomType_value = map[string]int32{
	"ROOM_TYPE0": 0,
	"ROOM_TYPE1": 1,
	"ROOM_TYPE2": 2,
}

func (RoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{2} }

// 桌子类型
type DeskType int32

const (
	DESK_TYPE0 DeskType = 0
	DESK_TYPE1 DeskType = 1
	DESK_TYPE2 DeskType = 2
	DESK_TYPE3 DeskType = 3
)

var DeskType_name = map[int32]string{
	0: "DESK_TYPE0",
	1: "DESK_TYPE1",
	2: "DESK_TYPE2",
	3: "DESK_TYPE3",
}
var DeskType_value = map[string]int32{
	"DESK_TYPE0": 0,
	"DESK_TYPE1": 1,
	"DESK_TYPE2": 2,
	"DESK_TYPE3": 3,
}

func (DeskType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{3} }

// 房间等级
type RoomLevel int32

const (
	ROOM_LEVEL0 RoomLevel = 0
	ROOM_LEVEL1 RoomLevel = 1
	ROOM_LEVEL2 RoomLevel = 2
	ROOM_LEVEL3 RoomLevel = 3
	ROOM_LEVEL4 RoomLevel = 4
)

var RoomLevel_name = map[int32]string{
	0: "ROOM_LEVEL0",
	1: "ROOM_LEVEL1",
	2: "ROOM_LEVEL2",
	3: "ROOM_LEVEL3",
	4: "ROOM_LEVEL4",
}
var RoomLevel_value = map[string]int32{
	"ROOM_LEVEL0": 0,
	"ROOM_LEVEL1": 1,
	"ROOM_LEVEL2": 2,
	"ROOM_LEVEL3": 3,
	"ROOM_LEVEL4": 4,
}

func (RoomLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{4} }

// 桌子位置
type DeskSeat int32

const (
	DESK_SEAT0 DeskSeat = 0
	DESK_SEAT1 DeskSeat = 1
	DESK_SEAT2 DeskSeat = 2
	DESK_SEAT3 DeskSeat = 3
	DESK_SEAT4 DeskSeat = 4
	DESK_SEAT5 DeskSeat = 5
	DESK_SEAT6 DeskSeat = 6
	DESK_SEAT7 DeskSeat = 7
	DESK_SEAT8 DeskSeat = 8
	DESK_SEAT9 DeskSeat = 9
)

var DeskSeat_name = map[int32]string{
	0: "DESK_SEAT0",
	1: "DESK_SEAT1",
	2: "DESK_SEAT2",
	3: "DESK_SEAT3",
	4: "DESK_SEAT4",
	5: "DESK_SEAT5",
	6: "DESK_SEAT6",
	7: "DESK_SEAT7",
	8: "DESK_SEAT8",
	9: "DESK_SEAT9",
}
var DeskSeat_value = map[string]int32{
	"DESK_SEAT0": 0,
	"DESK_SEAT1": 1,
	"DESK_SEAT2": 2,
	"DESK_SEAT3": 3,
	"DESK_SEAT4": 4,
	"DESK_SEAT5": 5,
	"DESK_SEAT6": 6,
	"DESK_SEAT7": 7,
	"DESK_SEAT8": 8,
	"DESK_SEAT9": 9,
}

func (DeskSeat) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{5} }

// 庄家操作类型
type DealerType int32

const (
	// 房间0下庄 1上庄 2补庄
	DEALER_DOWN DealerType = 0
	DEALER_UP   DealerType = 1
	DEALER_BU   DealerType = 2
)

var DealerType_name = map[int32]string{
	0: "DEALER_DOWN",
	1: "DEALER_UP",
	2: "DEALER_BU",
}
var DealerType_value = map[string]int32{
	"DEALER_DOWN": 0,
	"DEALER_UP":   1,
	"DEALER_BU":   2,
}

func (DealerType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{6} }

// 商城物品类型
type CurrencyType int32

const (
	CHEAP   CurrencyType = 0
	DIAMOND CurrencyType = 1
	COIN    CurrencyType = 2
	CARD    CurrencyType = 3
	CHIP    CurrencyType = 4
	VIP     CurrencyType = 5
)

var CurrencyType_name = map[int32]string{
	0: "CHEAP",
	1: "DIAMOND",
	2: "COIN",
	3: "CARD",
	4: "CHIP",
	5: "VIP",
}
var CurrencyType_value = map[string]int32{
	"CHEAP":   0,
	"DIAMOND": 1,
	"COIN":    2,
	"CARD":    3,
	"CHIP":    4,
	"VIP":     5,
}

func (CurrencyType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{7} }

// 商城支付方式类型
type PaymentType int32

const (
	PAY_COIN    PaymentType = 0
	PAY_RMB     PaymentType = 1
	PAY_DIAMOND PaymentType = 2
)

var PaymentType_name = map[int32]string{
	0: "PAY_COIN",
	1: "PAY_RMB",
	2: "PAY_DIAMOND",
}
var PaymentType_value = map[string]int32{
	"PAY_COIN":    0,
	"PAY_RMB":     1,
	"PAY_DIAMOND": 2,
}

func (PaymentType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{8} }

// 支付交易结果类型
type TradeType int32

const (
	TradeSuccess TradeType = 0
	TradeFail    TradeType = 1
	Tradeing     TradeType = 2
	TradeGoods   TradeType = 3
)

var TradeType_name = map[int32]string{
	0: "TradeSuccess",
	1: "TradeFail",
	2: "Tradeing",
	3: "TradeGoods",
}
var TradeType_value = map[string]int32{
	"TradeSuccess": 0,
	"TradeFail":    1,
	"Tradeing":     2,
	"TradeGoods":   3,
}

func (TradeType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{9} }

// 消息类型
type NoticeType int32

const (
	NOTICE_TYPE0 NoticeType = 0
	NOTICE_TYPE1 NoticeType = 1
	NOTICE_TYPE2 NoticeType = 2
	NOTICE_TYPE3 NoticeType = 3
)

var NoticeType_name = map[int32]string{
	0: "NOTICE_TYPE0",
	1: "NOTICE_TYPE1",
	2: "NOTICE_TYPE2",
	3: "NOTICE_TYPE3",
}
var NoticeType_value = map[string]int32{
	"NOTICE_TYPE0": 0,
	"NOTICE_TYPE1": 1,
	"NOTICE_TYPE2": 2,
	"NOTICE_TYPE3": 3,
}

func (NoticeType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{10} }

// 消息操作类型
type NoticeActType int32

const (
	NOTICE_ACT_TYPE0 NoticeActType = 0
	NOTICE_ACT_TYPE1 NoticeActType = 1
	NOTICE_ACT_TYPE2 NoticeActType = 2
)

var NoticeActType_name = map[int32]string{
	0: "NOTICE_ACT_TYPE0",
	1: "NOTICE_ACT_TYPE1",
	2: "NOTICE_ACT_TYPE2",
}
var NoticeActType_value = map[string]int32{
	"NOTICE_ACT_TYPE0": 0,
	"NOTICE_ACT_TYPE1": 1,
	"NOTICE_ACT_TYPE2": 2,
}

func (NoticeActType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{11} }

// 游戏变量
type GameEnv int32

const (
	Regist_diamond  GameEnv = 0
	Regist_coin     GameEnv = 1
	Regist_chip     GameEnv = 2
	Regist_card     GameEnv = 3
	Build           GameEnv = 4
	First_pay_multi GameEnv = 5
	First_pay_coin  GameEnv = 6
	Relieve         GameEnv = 7
	Prizedraw       GameEnv = 8
	Bankrupt_coin   GameEnv = 9
	Relieve_coin    GameEnv = 10
	Robot_num       GameEnv = 11
	Robot_allot1    GameEnv = 12
	Robot_allot2    GameEnv = 13
	Robot_bet       GameEnv = 14
)

var GameEnv_name = map[int32]string{
	0:  "Regist_diamond",
	1:  "Regist_coin",
	2:  "Regist_chip",
	3:  "Regist_card",
	4:  "Build",
	5:  "First_pay_multi",
	6:  "First_pay_coin",
	7:  "Relieve",
	8:  "Prizedraw",
	9:  "Bankrupt_coin",
	10: "Relieve_coin",
	11: "Robot_num",
	12: "Robot_allot1",
	13: "Robot_allot2",
	14: "Robot_bet",
}
var GameEnv_value = map[string]int32{
	"Regist_diamond":  0,
	"Regist_coin":     1,
	"Regist_chip":     2,
	"Regist_card":     3,
	"Build":           4,
	"First_pay_multi": 5,
	"First_pay_coin":  6,
	"Relieve":         7,
	"Prizedraw":       8,
	"Bankrupt_coin":   9,
	"Relieve_coin":    10,
	"Robot_num":       11,
	"Robot_allot1":    12,
	"Robot_allot2":    13,
	"Robot_bet":       14,
}

func (GameEnv) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{12} }

// 登出类型
type LogoutType int32

const (
	LOGOUT_TYPE0 LogoutType = 0
	LOGOUT_TYPE1 LogoutType = 1
	LOGOUT_TYPE2 LogoutType = 2
	LOGOUT_TYPE3 LogoutType = 3
	LOGOUT_TYPE4 LogoutType = 4
)

var LogoutType_name = map[int32]string{
	0: "LOGOUT_TYPE0",
	1: "LOGOUT_TYPE1",
	2: "LOGOUT_TYPE2",
	3: "LOGOUT_TYPE3",
	4: "LOGOUT_TYPE4",
}
var LogoutType_value = map[string]int32{
	"LOGOUT_TYPE0": 0,
	"LOGOUT_TYPE1": 1,
	"LOGOUT_TYPE2": 2,
	"LOGOUT_TYPE3": 3,
	"LOGOUT_TYPE4": 4,
}

func (LogoutType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{13} }

// 日志操作类型
type LogType int32

const (
	LOG_TYPE0  LogType = 0
	LOG_TYPE1  LogType = 1
	LOG_TYPE2  LogType = 2
	LOG_TYPE3  LogType = 3
	LOG_TYPE4  LogType = 4
	LOG_TYPE5  LogType = 5
	LOG_TYPE6  LogType = 6
	LOG_TYPE7  LogType = 7
	LOG_TYPE8  LogType = 8
	LOG_TYPE9  LogType = 9
	LOG_TYPE10 LogType = 10
	LOG_TYPE11 LogType = 11
	LOG_TYPE12 LogType = 12
	LOG_TYPE13 LogType = 13
	LOG_TYPE14 LogType = 14
	LOG_TYPE15 LogType = 15
	LOG_TYPE18 LogType = 18
	LOG_TYPE24 LogType = 24
	LOG_TYPE25 LogType = 25
	LOG_TYPE26 LogType = 26
	LOG_TYPE44 LogType = 44
	LOG_TYPE45 LogType = 45
	LOG_TYPE46 LogType = 46
	LOG_TYPE47 LogType = 47
	LOG_TYPE48 LogType = 48
	LOG_TYPE49 LogType = 49
	LOG_TYPE50 LogType = 50
	LOG_TYPE51 LogType = 51
)

var LogType_name = map[int32]string{
	0:  "LOG_TYPE0",
	1:  "LOG_TYPE1",
	2:  "LOG_TYPE2",
	3:  "LOG_TYPE3",
	4:  "LOG_TYPE4",
	5:  "LOG_TYPE5",
	6:  "LOG_TYPE6",
	7:  "LOG_TYPE7",
	8:  "LOG_TYPE8",
	9:  "LOG_TYPE9",
	10: "LOG_TYPE10",
	11: "LOG_TYPE11",
	12: "LOG_TYPE12",
	13: "LOG_TYPE13",
	14: "LOG_TYPE14",
	15: "LOG_TYPE15",
	18: "LOG_TYPE18",
	24: "LOG_TYPE24",
	25: "LOG_TYPE25",
	26: "LOG_TYPE26",
	44: "LOG_TYPE44",
	45: "LOG_TYPE45",
	46: "LOG_TYPE46",
	47: "LOG_TYPE47",
	48: "LOG_TYPE48",
	49: "LOG_TYPE49",
	50: "LOG_TYPE50",
	51: "LOG_TYPE51",
}
var LogType_value = map[string]int32{
	"LOG_TYPE0":  0,
	"LOG_TYPE1":  1,
	"LOG_TYPE2":  2,
	"LOG_TYPE3":  3,
	"LOG_TYPE4":  4,
	"LOG_TYPE5":  5,
	"LOG_TYPE6":  6,
	"LOG_TYPE7":  7,
	"LOG_TYPE8":  8,
	"LOG_TYPE9":  9,
	"LOG_TYPE10": 10,
	"LOG_TYPE11": 11,
	"LOG_TYPE12": 12,
	"LOG_TYPE13": 13,
	"LOG_TYPE14": 14,
	"LOG_TYPE15": 15,
	"LOG_TYPE18": 18,
	"LOG_TYPE24": 24,
	"LOG_TYPE25": 25,
	"LOG_TYPE26": 26,
	"LOG_TYPE44": 44,
	"LOG_TYPE45": 45,
	"LOG_TYPE46": 46,
	"LOG_TYPE47": 47,
	"LOG_TYPE48": 48,
	"LOG_TYPE49": 49,
	"LOG_TYPE50": 50,
	"LOG_TYPE51": 51,
}

func (LogType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{14} }

// 放看比加跟掩码
type ActType int32

const (
	ACT_NUL   ActType = 0
	ACT_FOLD  ActType = 1
	ACT_SEE   ActType = 2
	ACT_BI    ActType = 4
	ACT_RAISE ActType = 8
	ACT_CALL  ActType = 16
)

var ActType_name = map[int32]string{
	0:  "ACT_NUL",
	1:  "ACT_FOLD",
	2:  "ACT_SEE",
	4:  "ACT_BI",
	8:  "ACT_RAISE",
	16: "ACT_CALL",
}
var ActType_value = map[string]int32{
	"ACT_NUL":   0,
	"ACT_FOLD":  1,
	"ACT_SEE":   2,
	"ACT_BI":    4,
	"ACT_RAISE": 8,
	"ACT_CALL":  16,
}

func (ActType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{15} }

// 任务类型
type TaskType int32

const (
	TASK_TYPE0  TaskType = 0
	TASK_TYPE1  TaskType = 1
	TASK_TYPE2  TaskType = 2
	TASK_TYPE3  TaskType = 3
	TASK_TYPE4  TaskType = 4
	TASK_TYPE5  TaskType = 5
	TASK_TYPE6  TaskType = 6
	TASK_TYPE7  TaskType = 7
	TASK_TYPE8  TaskType = 8
	TASK_TYPE9  TaskType = 9
	TASK_TYPE10 TaskType = 10
	TASK_TYPE11 TaskType = 11
	TASK_TYPE12 TaskType = 12
	TASK_TYPE13 TaskType = 13
	TASK_TYPE14 TaskType = 14
	TASK_TYPE15 TaskType = 15
	TASK_TYPE16 TaskType = 16
	TASK_TYPE17 TaskType = 17
	TASK_TYPE18 TaskType = 18
	TASK_TYPE19 TaskType = 19
	TASK_TYPE20 TaskType = 20
	TASK_TYPE21 TaskType = 21
	TASK_TYPE22 TaskType = 22
	TASK_TYPE23 TaskType = 23
	TASK_TYPE24 TaskType = 24
)

var TaskType_name = map[int32]string{
	0:  "TASK_TYPE0",
	1:  "TASK_TYPE1",
	2:  "TASK_TYPE2",
	3:  "TASK_TYPE3",
	4:  "TASK_TYPE4",
	5:  "TASK_TYPE5",
	6:  "TASK_TYPE6",
	7:  "TASK_TYPE7",
	8:  "TASK_TYPE8",
	9:  "TASK_TYPE9",
	10: "TASK_TYPE10",
	11: "TASK_TYPE11",
	12: "TASK_TYPE12",
	13: "TASK_TYPE13",
	14: "TASK_TYPE14",
	15: "TASK_TYPE15",
	16: "TASK_TYPE16",
	17: "TASK_TYPE17",
	18: "TASK_TYPE18",
	19: "TASK_TYPE19",
	20: "TASK_TYPE20",
	21: "TASK_TYPE21",
	22: "TASK_TYPE22",
	23: "TASK_TYPE23",
	24: "TASK_TYPE24",
}
var TaskType_value = map[string]int32{
	"TASK_TYPE0":  0,
	"TASK_TYPE1":  1,
	"TASK_TYPE2":  2,
	"TASK_TYPE3":  3,
	"TASK_TYPE4":  4,
	"TASK_TYPE5":  5,
	"TASK_TYPE6":  6,
	"TASK_TYPE7":  7,
	"TASK_TYPE8":  8,
	"TASK_TYPE9":  9,
	"TASK_TYPE10": 10,
	"TASK_TYPE11": 11,
	"TASK_TYPE12": 12,
	"TASK_TYPE13": 13,
	"TASK_TYPE14": 14,
	"TASK_TYPE15": 15,
	"TASK_TYPE16": 16,
	"TASK_TYPE17": 17,
	"TASK_TYPE18": 18,
	"TASK_TYPE19": 19,
	"TASK_TYPE20": 20,
	"TASK_TYPE21": 21,
	"TASK_TYPE22": 22,
	"TASK_TYPE23": 23,
	"TASK_TYPE24": 24,
}

func (TaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGameType, []int{16} }

func init() {
	proto.RegisterEnum("pb.GameType", GameType_name, GameType_value)
	proto.RegisterEnum("pb.DeskState", DeskState_name, DeskState_value)
	proto.RegisterEnum("pb.RoomType", RoomType_name, RoomType_value)
	proto.RegisterEnum("pb.DeskType", DeskType_name, DeskType_value)
	proto.RegisterEnum("pb.RoomLevel", RoomLevel_name, RoomLevel_value)
	proto.RegisterEnum("pb.DeskSeat", DeskSeat_name, DeskSeat_value)
	proto.RegisterEnum("pb.DealerType", DealerType_name, DealerType_value)
	proto.RegisterEnum("pb.CurrencyType", CurrencyType_name, CurrencyType_value)
	proto.RegisterEnum("pb.PaymentType", PaymentType_name, PaymentType_value)
	proto.RegisterEnum("pb.TradeType", TradeType_name, TradeType_value)
	proto.RegisterEnum("pb.NoticeType", NoticeType_name, NoticeType_value)
	proto.RegisterEnum("pb.NoticeActType", NoticeActType_name, NoticeActType_value)
	proto.RegisterEnum("pb.GameEnv", GameEnv_name, GameEnv_value)
	proto.RegisterEnum("pb.LogoutType", LogoutType_name, LogoutType_value)
	proto.RegisterEnum("pb.LogType", LogType_name, LogType_value)
	proto.RegisterEnum("pb.ActType", ActType_name, ActType_value)
	proto.RegisterEnum("pb.TaskType", TaskType_name, TaskType_value)
}
func (x GameType) String() string {
	s, ok := GameType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x DeskState) String() string {
	s, ok := DeskState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x RoomType) String() string {
	s, ok := RoomType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x DeskType) String() string {
	s, ok := DeskType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x RoomLevel) String() string {
	s, ok := RoomLevel_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x DeskSeat) String() string {
	s, ok := DeskSeat_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x DealerType) String() string {
	s, ok := DealerType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CurrencyType) String() string {
	s, ok := CurrencyType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x PaymentType) String() string {
	s, ok := PaymentType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TradeType) String() string {
	s, ok := TradeType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x NoticeType) String() string {
	s, ok := NoticeType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x NoticeActType) String() string {
	s, ok := NoticeActType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x GameEnv) String() string {
	s, ok := GameEnv_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x LogoutType) String() string {
	s, ok := LogoutType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x LogType) String() string {
	s, ok := LogType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ActType) String() string {
	s, ok := ActType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TaskType) String() string {
	s, ok := TaskType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("game_type.proto", fileDescriptorGameType) }

var fileDescriptorGameType = []byte{
	// 1051 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x96, 0xcf, 0x77, 0xe2, 0x54,
	0x14, 0xc7, 0x09, 0xd0, 0x02, 0x17, 0x28, 0x77, 0x32, 0xa3, 0x8e, 0x2e, 0xf2, 0x07, 0xe4, 0x8c,
	0x63, 0xf9, 0xd1, 0x76, 0xaa, 0xab, 0x00, 0x29, 0x45, 0x53, 0x82, 0x21, 0xd4, 0x53, 0x8f, 0xe7,
	0x60, 0x0a, 0xb1, 0xe6, 0x0c, 0x10, 0x4e, 0x08, 0xf5, 0xd4, 0x95, 0x3b, 0xb7, 0xfe, 0x19, 0xea,
	0x3f, 0xe1, 0xd6, 0xe5, 0x2c, 0x5d, 0x5a, 0xdc, 0xb8, 0x9c, 0x3f, 0xc1, 0x73, 0x93, 0x47, 0xc8,
	0x75, 0x76, 0xef, 0xf3, 0x6d, 0xde, 0x7d, 0xf7, 0xbd, 0xfb, 0xbd, 0x97, 0x42, 0xed, 0xce, 0x59,
	0xb8, 0x93, 0xf0, 0x61, 0xe5, 0xbe, 0x5c, 0x05, 0x7e, 0xe8, 0xcb, 0xd9, 0xd5, 0xad, 0xfa, 0x39,
	0x14, 0x7b, 0xce, 0xc2, 0xb5, 0x1f, 0x56, 0xae, 0x5c, 0x84, 0x7c, 0x4f, 0xbb, 0xd2, 0x31, 0x23,
	0x17, 0x20, 0x37, 0xe8, 0x8f, 0x51, 0xa2, 0xc5, 0x48, 0x1b, 0x60, 0x96, 0x16, 0x5d, 0x73, 0x8c,
	0x39, 0x5a, 0x5c, 0x8e, 0x35, 0xcc, 0xd3, 0x62, 0xa8, 0xf5, 0xf1, 0x80, 0x16, 0x7a, 0xbb, 0x87,
	0x87, 0xea, 0x37, 0x50, 0xea, 0xba, 0xeb, 0xd7, 0xa3, 0xd0, 0x09, 0x5d, 0xb9, 0x06, 0xe5, 0x91,
	0xad, 0xd9, 0xfa, 0xc4, 0xd2, 0xb5, 0xee, 0x0d, 0x66, 0x64, 0x84, 0x4a, 0x2c, 0x74, 0x75, 0xcd,
	0xd0, 0x2d, 0x94, 0xe4, 0x2a, 0x94, 0x62, 0x85, 0xce, 0xca, 0xee, 0xb1, 0xad, 0xdb, 0x98, 0x93,
	0x8f, 0x00, 0x62, 0x34, 0xaf, 0x75, 0x0b, 0xf3, 0xea, 0xa7, 0x50, 0xb4, 0x7c, 0x7f, 0x11, 0x65,
	0x7a, 0x04, 0x60, 0x99, 0xe6, 0xd5, 0xc4, 0xbe, 0x19, 0xea, 0xc7, 0x98, 0x61, 0x5c, 0x47, 0x89,
	0x71, 0x03, 0xb3, 0x74, 0x4b, 0xca, 0x6c, 0xb7, 0xb7, 0xab, 0x8f, 0xbe, 0x48, 0xef, 0x4d, 0x58,
	0xec, 0x4d, 0xb8, 0x81, 0x59, 0xc6, 0x4d, 0xcc, 0xa9, 0xdf, 0x42, 0x89, 0xf2, 0x30, 0xdc, 0x7b,
	0x77, 0x4e, 0xb7, 0x8c, 0x0e, 0x32, 0xf4, 0x6b, 0xdd, 0xa0, 0x68, 0x4c, 0xa0, 0x70, 0x4c, 0xa0,
	0x78, 0x4c, 0x68, 0x62, 0x8e, 0x0b, 0x2d, 0xcc, 0xab, 0xbf, 0x4b, 0x71, 0xba, 0x23, 0xd7, 0x09,
	0x93, 0xe3, 0x47, 0xba, 0x66, 0xa7, 0xd3, 0x25, 0x4e, 0xa7, 0x4b, 0x9c, 0x4e, 0x97, 0xb8, 0x19,
	0x3f, 0x63, 0xc2, 0x2d, 0xcc, 0x33, 0x3e, 0xc1, 0x03, 0xc6, 0xa7, 0x78, 0xc8, 0xf8, 0x0c, 0x0b,
	0x8c, 0x5f, 0x61, 0x91, 0xf1, 0x39, 0x96, 0xd4, 0xcf, 0x00, 0xba, 0xae, 0x33, 0x77, 0x83, 0xe8,
	0x71, 0x6b, 0x50, 0x8e, 0xcb, 0x3b, 0xe9, 0x9a, 0x5f, 0x0d, 0x30, 0x43, 0x45, 0x15, 0xc2, 0x78,
	0x18, 0x97, 0x5c, 0x60, 0x7b, 0x8c, 0x59, 0xf5, 0x0a, 0x2a, 0x9d, 0x4d, 0x10, 0xb8, 0xcb, 0xe9,
	0x43, 0xb4, 0xbd, 0x04, 0x07, 0x9d, 0x4b, 0x5d, 0x1b, 0x62, 0x46, 0x2e, 0x43, 0xa1, 0xdb, 0xd7,
	0xae, 0xcc, 0x41, 0x17, 0x25, 0x72, 0x66, 0xc7, 0xec, 0x93, 0x0f, 0x69, 0xa5, 0x59, 0x5d, 0xcc,
	0x45, 0xab, 0xcb, 0xfe, 0x30, 0x76, 0xe2, 0x75, 0x7f, 0x88, 0x07, 0xea, 0x39, 0x94, 0x87, 0xce,
	0xc3, 0xc2, 0x5d, 0x86, 0x51, 0xb4, 0x0a, 0x14, 0x87, 0xda, 0xcd, 0x24, 0xda, 0x19, 0x05, 0x24,
	0xb2, 0xae, 0xda, 0x71, 0x55, 0x08, 0x76, 0x27, 0x90, 0x43, 0x4a, 0x76, 0xe0, 0xcc, 0xe2, 0x46,
	0x40, 0xa8, 0x44, 0x30, 0xda, 0x4c, 0xa7, 0xee, 0x7a, 0x1d, 0x5f, 0x23, 0x52, 0x2e, 0x1c, 0x6f,
	0x8e, 0x12, 0x45, 0x8e, 0xd0, 0x5b, 0xde, 0xc5, 0x4f, 0x1e, 0x51, 0xcf, 0xf7, 0x67, 0x6b, 0xcc,
	0xa9, 0x36, 0xc0, 0xc0, 0x0f, 0xbd, 0x69, 0x12, 0x6c, 0x60, 0xda, 0xfd, 0x8e, 0x9e, 0x38, 0x8e,
	0x2b, 0x54, 0x44, 0xae, 0x50, 0x19, 0xb9, 0x42, 0xbe, 0xfb, 0x12, 0xaa, 0x71, 0x54, 0x6d, 0x1a,
	0x5f, 0xef, 0x19, 0xa0, 0xf8, 0x44, 0xeb, 0xd8, 0x49, 0xf0, 0x77, 0x55, 0x3a, 0xe0, 0x5d, 0x95,
	0xda, 0xe2, 0xe7, 0x2c, 0x14, 0xa8, 0xfb, 0xf5, 0xe5, 0xbd, 0x2c, 0xc3, 0x91, 0xe5, 0xde, 0x79,
	0xeb, 0x70, 0x32, 0xf3, 0x9c, 0x85, 0xbf, 0x9c, 0x09, 0x33, 0xc7, 0xda, 0xd4, 0xf7, 0x96, 0xc2,
	0xcc, 0x42, 0xf8, 0xde, 0x5b, 0x09, 0x33, 0x0b, 0xc1, 0x09, 0x66, 0x98, 0xa3, 0x0a, 0xb6, 0x37,
	0xde, 0x7c, 0x86, 0x79, 0xf9, 0x29, 0xd4, 0x2e, 0xbc, 0x60, 0x1d, 0x4e, 0x56, 0xce, 0xc3, 0x64,
	0xb1, 0x99, 0x87, 0x1e, 0x1e, 0xd0, 0x31, 0x7b, 0x31, 0x8a, 0x7a, 0x48, 0x95, 0xb1, 0xdc, 0xb9,
	0xe7, 0xde, 0xbb, 0x58, 0xa0, 0x97, 0x1e, 0x06, 0xde, 0x8f, 0xee, 0x2c, 0x70, 0x7e, 0xc0, 0xa2,
	0xfc, 0x04, 0xaa, 0x6d, 0x67, 0xf9, 0x3a, 0xd8, 0xac, 0x44, 0x12, 0x25, 0x7a, 0x1a, 0xf1, 0x79,
	0xac, 0x00, 0xed, 0xb1, 0xfc, 0x5b, 0x3f, 0x9c, 0x2c, 0x37, 0x0b, 0x2c, 0x47, 0x1f, 0x44, 0xe8,
	0xcc, 0xe7, 0x7e, 0x58, 0xc7, 0xca, 0xff, 0x94, 0x06, 0x56, 0xf7, 0x5b, 0x6e, 0xdd, 0x10, 0x8f,
	0xd4, 0xef, 0x00, 0x0c, 0xff, 0xce, 0xdf, 0x84, 0xbb, 0x92, 0x19, 0x66, 0xcf, 0x1c, 0xdb, 0xe9,
	0x92, 0xa5, 0x14, 0x51, 0xb2, 0x94, 0x22, 0x4a, 0x96, 0x52, 0xa8, 0xf7, 0xb8, 0x42, 0xad, 0xfd,
	0x47, 0x0e, 0x0a, 0x86, 0x7f, 0x17, 0x9d, 0x52, 0x85, 0x92, 0x61, 0xf6, 0x92, 0x23, 0x52, 0x58,
	0x8f, 0x3b, 0x65, 0x87, 0x8d, 0x78, 0x38, 0xee, 0x90, 0x22, 0xa7, 0x90, 0x9a, 0x3a, 0x85, 0xd4,
	0xd3, 0x29, 0xa4, 0x96, 0x4e, 0xe1, 0x59, 0xfc, 0xc2, 0x3b, 0xa4, 0x86, 0x4e, 0xe1, 0x39, 0x96,
	0xc8, 0xcc, 0x49, 0x1a, 0xc7, 0x08, 0x8c, 0xeb, 0x58, 0x66, 0xdc, 0xc0, 0x0a, 0xe3, 0x26, 0x56,
	0x19, 0xb7, 0xf0, 0x88, 0xf1, 0x09, 0xd6, 0x18, 0xbf, 0x42, 0x39, 0xcd, 0x8d, 0x16, 0x3e, 0x67,
	0x7c, 0x82, 0x1f, 0x32, 0x3e, 0xc5, 0x8f, 0xd2, 0xdc, 0x6a, 0xe1, 0x0b, 0xc6, 0x27, 0xf8, 0x31,
	0xe3, 0x53, 0x7c, 0xc9, 0xf8, 0x0c, 0x3f, 0x61, 0xfc, 0x0a, 0x8f, 0x19, 0x9f, 0x63, 0x3d, 0xcd,
	0x27, 0xc7, 0xd8, 0x60, 0x5c, 0xc7, 0xa6, 0xfa, 0x35, 0x14, 0x76, 0x0d, 0x58, 0x86, 0x02, 0x75,
	0xd3, 0x60, 0x6c, 0x60, 0x86, 0x46, 0x02, 0xc1, 0x85, 0x69, 0xd0, 0xc0, 0x12, 0x7f, 0x1a, 0xe9,
	0x3a, 0x66, 0x65, 0x80, 0x43, 0x82, 0x76, 0x3f, 0xae, 0x14, 0xad, 0x2d, 0xad, 0x3f, 0xd2, 0xb1,
	0xb8, 0xdb, 0xd5, 0xd1, 0x0c, 0x03, 0x51, 0xfd, 0x2d, 0x07, 0x45, 0xdb, 0xd9, 0xff, 0x4e, 0xd9,
	0x1a, 0xff, 0x9d, 0x4a, 0x58, 0x0c, 0xfe, 0x84, 0xc5, 0xe0, 0x4f, 0x58, 0x0c, 0xfe, 0x84, 0xc5,
	0xe0, 0x4f, 0x58, 0x0c, 0xfe, 0x84, 0xc5, 0xe0, 0x4f, 0x58, 0x0c, 0xfe, 0x84, 0xc5, 0xe0, 0x4f,
	0x98, 0x8c, 0x52, 0x83, 0xf2, 0x3e, 0x1f, 0x72, 0x0a, 0x13, 0xc8, 0x2a, 0x4c, 0x20, 0xaf, 0x30,
	0x81, 0xcc, 0xc2, 0x04, 0x72, 0x0b, 0x13, 0xc8, 0x2e, 0x4c, 0x38, 0x45, 0xe4, 0xc2, 0x19, 0x3e,
	0xe1, 0x02, 0x39, 0x8a, 0x09, 0xe7, 0xf8, 0x94, 0x09, 0x8d, 0x63, 0x7c, 0xc6, 0x85, 0x3a, 0xbe,
	0xc7, 0x85, 0x06, 0xbe, 0xcf, 0x85, 0x26, 0x7e, 0xc0, 0x85, 0x16, 0x3e, 0x6f, 0xbf, 0x78, 0xf3,
	0xa8, 0x64, 0xfe, 0x7a, 0x54, 0x32, 0x6f, 0x1f, 0x15, 0xe9, 0xa7, 0xad, 0x22, 0xfd, 0xba, 0x55,
	0xa4, 0x3f, 0xb7, 0x8a, 0xf4, 0x66, 0xab, 0x48, 0x7f, 0x6f, 0x15, 0xe9, 0xdf, 0xad, 0x92, 0x79,
	0xbb, 0x55, 0xa4, 0x5f, 0xfe, 0x51, 0x32, 0xb7, 0x87, 0xd1, 0x7f, 0x5c, 0xcd, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb8, 0xb8, 0xcd, 0x60, 0x84, 0x09, 0x00, 0x00,
}
