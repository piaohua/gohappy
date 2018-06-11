// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: actor_game.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import actor "github.com/AsynkronIT/protoactor-go/actor"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SyncGames_SyncGamesType int32

const (
	TYPE_UPSERT SyncGames_SyncGamesType = 0
	TYPE_DELETE SyncGames_SyncGamesType = 1
)

var SyncGames_SyncGamesType_name = map[int32]string{
	0: "TYPE_UPSERT",
	1: "TYPE_DELETE",
}
var SyncGames_SyncGamesType_value = map[string]int32{
	"TYPE_UPSERT": 0,
	"TYPE_DELETE": 1,
}

func (SyncGames_SyncGamesType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorActorGame, []int{2, 0}
}

// 获取节点
type GetGames struct {
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type uint32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *GetGames) Reset()                    { *m = GetGames{} }
func (*GetGames) ProtoMessage()               {}
func (*GetGames) Descriptor() ([]byte, []int) { return fileDescriptorActorGame, []int{0} }

func (m *GetGames) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetGames) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type GotGames struct {
	Name    string     `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	GamePid *actor.PID `protobuf:"bytes,2,opt,name=GamePid" json:"GamePid,omitempty"`
}

func (m *GotGames) Reset()                    { *m = GotGames{} }
func (*GotGames) ProtoMessage()               {}
func (*GotGames) Descriptor() ([]byte, []int) { return fileDescriptorActorGame, []int{1} }

func (m *GotGames) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GotGames) GetGamePid() *actor.PID {
	if m != nil {
		return m.GamePid
	}
	return nil
}

// 同步节点
type SyncGames struct {
	Type    SyncGames_SyncGamesType `protobuf:"varint,1,opt,name=Type,proto3,enum=pb.SyncGames_SyncGamesType" json:"Type,omitempty"`
	Name    string                  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	GamePid *actor.PID              `protobuf:"bytes,3,opt,name=GamePid" json:"GamePid,omitempty"`
}

func (m *SyncGames) Reset()                    { *m = SyncGames{} }
func (*SyncGames) ProtoMessage()               {}
func (*SyncGames) Descriptor() ([]byte, []int) { return fileDescriptorActorGame, []int{2} }

func (m *SyncGames) GetType() SyncGames_SyncGamesType {
	if m != nil {
		return m.Type
	}
	return TYPE_UPSERT
}

func (m *SyncGames) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SyncGames) GetGamePid() *actor.PID {
	if m != nil {
		return m.GamePid
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGames)(nil), "pb.GetGames")
	proto.RegisterType((*GotGames)(nil), "pb.GotGames")
	proto.RegisterType((*SyncGames)(nil), "pb.SyncGames")
	proto.RegisterEnum("pb.SyncGames_SyncGamesType", SyncGames_SyncGamesType_name, SyncGames_SyncGamesType_value)
}
func (x SyncGames_SyncGamesType) String() string {
	s, ok := SyncGames_SyncGamesType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *GetGames) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetGames)
	if !ok {
		that2, ok := that.(GetGames)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *GotGames) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GotGames)
	if !ok {
		that2, ok := that.(GotGames)
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
	if this.Name != that1.Name {
		return false
	}
	if !this.GamePid.Equal(that1.GamePid) {
		return false
	}
	return true
}
func (this *SyncGames) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SyncGames)
	if !ok {
		that2, ok := that.(SyncGames)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.GamePid.Equal(that1.GamePid) {
		return false
	}
	return true
}
func (this *GetGames) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.GetGames{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GotGames) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.GotGames{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	if this.GamePid != nil {
		s = append(s, "GamePid: "+fmt.Sprintf("%#v", this.GamePid)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SyncGames) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pb.SyncGames{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	if this.GamePid != nil {
		s = append(s, "GamePid: "+fmt.Sprintf("%#v", this.GamePid)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringActorGame(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GetGames) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetGames) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorGame(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintActorGame(dAtA, i, uint64(m.Type))
	}
	return i, nil
}

func (m *GotGames) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GotGames) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintActorGame(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.GamePid != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintActorGame(dAtA, i, uint64(m.GamePid.Size()))
		n1, err := m.GamePid.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SyncGames) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncGames) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintActorGame(dAtA, i, uint64(m.Type))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintActorGame(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.GamePid != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintActorGame(dAtA, i, uint64(m.GamePid.Size()))
		n2, err := m.GamePid.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintActorGame(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetGames) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovActorGame(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovActorGame(uint64(m.Type))
	}
	return n
}

func (m *GotGames) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovActorGame(uint64(l))
	}
	if m.GamePid != nil {
		l = m.GamePid.Size()
		n += 1 + l + sovActorGame(uint64(l))
	}
	return n
}

func (m *SyncGames) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovActorGame(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovActorGame(uint64(l))
	}
	if m.GamePid != nil {
		l = m.GamePid.Size()
		n += 1 + l + sovActorGame(uint64(l))
	}
	return n
}

func sovActorGame(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozActorGame(x uint64) (n int) {
	return sovActorGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetGames) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetGames{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GotGames) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GotGames{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`GamePid:` + strings.Replace(fmt.Sprintf("%v", this.GamePid), "PID", "actor.PID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SyncGames) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SyncGames{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`GamePid:` + strings.Replace(fmt.Sprintf("%v", this.GamePid), "PID", "actor.PID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringActorGame(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetGames) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorGame
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
			return fmt.Errorf("proto: GetGames: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetGames: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorGame
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
				return ErrInvalidLengthActorGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActorGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorGame
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
func (m *GotGames) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorGame
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
			return fmt.Errorf("proto: GotGames: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GotGames: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorGame
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
				return ErrInvalidLengthActorGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GamePid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActorGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GamePid == nil {
				m.GamePid = &actor.PID{}
			}
			if err := m.GamePid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorGame
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
func (m *SyncGames) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActorGame
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
			return fmt.Errorf("proto: SyncGames: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncGames: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (SyncGames_SyncGamesType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorGame
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
				return ErrInvalidLengthActorGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GamePid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActorGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActorGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GamePid == nil {
				m.GamePid = &actor.PID{}
			}
			if err := m.GamePid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActorGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActorGame
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
func skipActorGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActorGame
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
					return 0, ErrIntOverflowActorGame
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
					return 0, ErrIntOverflowActorGame
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
				return 0, ErrInvalidLengthActorGame
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowActorGame
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
				next, err := skipActorGame(dAtA[start:])
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
	ErrInvalidLengthActorGame = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActorGame   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("actor_game.proto", fileDescriptorActorGame) }

var fileDescriptorActorGame = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4c, 0x2e, 0xc9,
	0x2f, 0x8a, 0x4f, 0x4f, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x2c, 0xae,
	0xcc, 0xcb, 0x2e, 0xca, 0xcf, 0xf3, 0x0c, 0xd1, 0x07, 0x2b, 0x00, 0x6b, 0xd0, 0x4d, 0xcf, 0xd7,
	0x07, 0x33, 0x20, 0x62, 0xc5, 0x10, 0xbd, 0x4a, 0x46, 0x5c, 0x1c, 0xee, 0xa9, 0x25, 0xee, 0x89,
	0xb9, 0xa9, 0xc5, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x60, 0x36, 0x48, 0x2c, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x37,
	0x08, 0xcc, 0x56, 0x72, 0xe1, 0xe2, 0x70, 0xcf, 0xc7, 0xa3, 0x47, 0x85, 0x8b, 0x1d, 0x24, 0x19,
	0x90, 0x99, 0x02, 0xd6, 0xc6, 0x6d, 0xc4, 0xa5, 0x07, 0xb6, 0x59, 0x2f, 0xc0, 0xd3, 0x25, 0x08,
	0x26, 0xa5, 0xb4, 0x92, 0x91, 0x8b, 0x33, 0xb8, 0x32, 0x2f, 0x19, 0x62, 0x8e, 0x3e, 0xd4, 0x1e,
	0x90, 0x39, 0x7c, 0x46, 0xd2, 0x7a, 0x05, 0x49, 0x7a, 0x70, 0x49, 0x04, 0x0b, 0xa4, 0x04, 0xe2,
	0x08, 0xb8, 0xc5, 0x4c, 0xd8, 0x2d, 0x66, 0xc6, 0x6d, 0xb1, 0x21, 0x17, 0x2f, 0x8a, 0x81, 0x42,
	0xfc, 0x5c, 0xdc, 0x21, 0x91, 0x01, 0xae, 0xf1, 0xa1, 0x01, 0xc1, 0xae, 0x41, 0x21, 0x02, 0x0c,
	0x70, 0x01, 0x17, 0x57, 0x1f, 0xd7, 0x10, 0x57, 0x01, 0x46, 0x27, 0x9d, 0x0b, 0x0f, 0xe5, 0x18,
	0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72,
	0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47,
	0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x0e, 0x5a, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xca, 0x49, 0xd3, 0xaa, 0x01, 0x00, 0x00,
}
