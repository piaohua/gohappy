// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: niu_vote.proto

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

// 私人局,发起投票申请解散房间
type CNNLaunchVote struct {
}

func (m *CNNLaunchVote) Reset()                    { *m = CNNLaunchVote{} }
func (*CNNLaunchVote) ProtoMessage()               {}
func (*CNNLaunchVote) Descriptor() ([]byte, []int) { return fileDescriptorNiuVote, []int{0} }

type SNNLaunchVote struct {
	Seat  uint32  `protobuf:"varint,1,opt,name=seat,proto3" json:"seat,omitempty"`
	Error ErrCode `protobuf:"varint,2,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *SNNLaunchVote) Reset()                    { *m = SNNLaunchVote{} }
func (*SNNLaunchVote) ProtoMessage()               {}
func (*SNNLaunchVote) Descriptor() ([]byte, []int) { return fileDescriptorNiuVote, []int{1} }

func (m *SNNLaunchVote) GetSeat() uint32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

func (m *SNNLaunchVote) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

// 私人局,发起投票,投票解散房间
type CNNVote struct {
	Vote uint32 `protobuf:"varint,1,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (m *CNNVote) Reset()                    { *m = CNNVote{} }
func (*CNNVote) ProtoMessage()               {}
func (*CNNVote) Descriptor() ([]byte, []int) { return fileDescriptorNiuVote, []int{2} }

func (m *CNNVote) GetVote() uint32 {
	if m != nil {
		return m.Vote
	}
	return 0
}

type SNNVote struct {
	Vote  uint32  `protobuf:"varint,1,opt,name=vote,proto3" json:"vote,omitempty"`
	Seat  uint32  `protobuf:"varint,2,opt,name=seat,proto3" json:"seat,omitempty"`
	Error ErrCode `protobuf:"varint,3,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *SNNVote) Reset()                    { *m = SNNVote{} }
func (*SNNVote) ProtoMessage()               {}
func (*SNNVote) Descriptor() ([]byte, []int) { return fileDescriptorNiuVote, []int{3} }

func (m *SNNVote) GetVote() uint32 {
	if m != nil {
		return m.Vote
	}
	return 0
}

func (m *SNNVote) GetSeat() uint32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

func (m *SNNVote) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

// 投票解散房间事件结果,服务器主动推送
type SNNVoteResult struct {
	// 0半数通过马上解散房间,
	// 1半数以上不通过终止解散房间
	Vote uint32 `protobuf:"varint,1,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (m *SNNVoteResult) Reset()                    { *m = SNNVoteResult{} }
func (*SNNVoteResult) ProtoMessage()               {}
func (*SNNVoteResult) Descriptor() ([]byte, []int) { return fileDescriptorNiuVote, []int{4} }

func (m *SNNVoteResult) GetVote() uint32 {
	if m != nil {
		return m.Vote
	}
	return 0
}

func init() {
	proto.RegisterType((*CNNLaunchVote)(nil), "pb.CNNLaunchVote")
	proto.RegisterType((*SNNLaunchVote)(nil), "pb.SNNLaunchVote")
	proto.RegisterType((*CNNVote)(nil), "pb.CNNVote")
	proto.RegisterType((*SNNVote)(nil), "pb.SNNVote")
	proto.RegisterType((*SNNVoteResult)(nil), "pb.SNNVoteResult")
}
func (this *CNNLaunchVote) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CNNLaunchVote)
	if !ok {
		that2, ok := that.(CNNLaunchVote)
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
	return true
}
func (this *SNNLaunchVote) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SNNLaunchVote)
	if !ok {
		that2, ok := that.(SNNLaunchVote)
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
	if this.Seat != that1.Seat {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *CNNVote) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CNNVote)
	if !ok {
		that2, ok := that.(CNNVote)
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
	if this.Vote != that1.Vote {
		return false
	}
	return true
}
func (this *SNNVote) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SNNVote)
	if !ok {
		that2, ok := that.(SNNVote)
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
	if this.Vote != that1.Vote {
		return false
	}
	if this.Seat != that1.Seat {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *SNNVoteResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SNNVoteResult)
	if !ok {
		that2, ok := that.(SNNVoteResult)
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
	if this.Vote != that1.Vote {
		return false
	}
	return true
}
func (this *CNNLaunchVote) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&pb.CNNLaunchVote{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SNNLaunchVote) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.SNNLaunchVote{")
	s = append(s, "Seat: "+fmt.Sprintf("%#v", this.Seat)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CNNVote) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.CNNVote{")
	s = append(s, "Vote: "+fmt.Sprintf("%#v", this.Vote)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SNNVote) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pb.SNNVote{")
	s = append(s, "Vote: "+fmt.Sprintf("%#v", this.Vote)+",\n")
	s = append(s, "Seat: "+fmt.Sprintf("%#v", this.Seat)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SNNVoteResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.SNNVoteResult{")
	s = append(s, "Vote: "+fmt.Sprintf("%#v", this.Vote)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNiuVote(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CNNLaunchVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CNNLaunchVote) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SNNLaunchVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SNNLaunchVote) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Seat != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNiuVote(dAtA, i, uint64(m.Seat))
	}
	if m.Error != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNiuVote(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func (m *CNNVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CNNVote) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNiuVote(dAtA, i, uint64(m.Vote))
	}
	return i, nil
}

func (m *SNNVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SNNVote) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNiuVote(dAtA, i, uint64(m.Vote))
	}
	if m.Seat != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNiuVote(dAtA, i, uint64(m.Seat))
	}
	if m.Error != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNiuVote(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func (m *SNNVoteResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SNNVoteResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNiuVote(dAtA, i, uint64(m.Vote))
	}
	return i, nil
}

func encodeVarintNiuVote(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CNNLaunchVote) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SNNLaunchVote) Size() (n int) {
	var l int
	_ = l
	if m.Seat != 0 {
		n += 1 + sovNiuVote(uint64(m.Seat))
	}
	if m.Error != 0 {
		n += 1 + sovNiuVote(uint64(m.Error))
	}
	return n
}

func (m *CNNVote) Size() (n int) {
	var l int
	_ = l
	if m.Vote != 0 {
		n += 1 + sovNiuVote(uint64(m.Vote))
	}
	return n
}

func (m *SNNVote) Size() (n int) {
	var l int
	_ = l
	if m.Vote != 0 {
		n += 1 + sovNiuVote(uint64(m.Vote))
	}
	if m.Seat != 0 {
		n += 1 + sovNiuVote(uint64(m.Seat))
	}
	if m.Error != 0 {
		n += 1 + sovNiuVote(uint64(m.Error))
	}
	return n
}

func (m *SNNVoteResult) Size() (n int) {
	var l int
	_ = l
	if m.Vote != 0 {
		n += 1 + sovNiuVote(uint64(m.Vote))
	}
	return n
}

func sovNiuVote(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNiuVote(x uint64) (n int) {
	return sovNiuVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CNNLaunchVote) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CNNLaunchVote{`,
		`}`,
	}, "")
	return s
}
func (this *SNNLaunchVote) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SNNLaunchVote{`,
		`Seat:` + fmt.Sprintf("%v", this.Seat) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CNNVote) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CNNVote{`,
		`Vote:` + fmt.Sprintf("%v", this.Vote) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SNNVote) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SNNVote{`,
		`Vote:` + fmt.Sprintf("%v", this.Vote) + `,`,
		`Seat:` + fmt.Sprintf("%v", this.Seat) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SNNVoteResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SNNVoteResult{`,
		`Vote:` + fmt.Sprintf("%v", this.Vote) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNiuVote(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CNNLaunchVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNiuVote
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
			return fmt.Errorf("proto: CNNLaunchVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CNNLaunchVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNiuVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNiuVote
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
func (m *SNNLaunchVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNiuVote
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
			return fmt.Errorf("proto: SNNLaunchVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SNNLaunchVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seat", wireType)
			}
			m.Seat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNiuVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seat |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNiuVote
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
			skippy, err := skipNiuVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNiuVote
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
func (m *CNNVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNiuVote
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
			return fmt.Errorf("proto: CNNVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CNNVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNiuVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNiuVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNiuVote
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
func (m *SNNVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNiuVote
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
			return fmt.Errorf("proto: SNNVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SNNVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNiuVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seat", wireType)
			}
			m.Seat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNiuVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seat |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNiuVote
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
			skippy, err := skipNiuVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNiuVote
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
func (m *SNNVoteResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNiuVote
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
			return fmt.Errorf("proto: SNNVoteResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SNNVoteResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNiuVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNiuVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNiuVote
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
func skipNiuVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNiuVote
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
					return 0, ErrIntOverflowNiuVote
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
					return 0, ErrIntOverflowNiuVote
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
				return 0, ErrInvalidLengthNiuVote
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNiuVote
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
				next, err := skipNiuVote(dAtA[start:])
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
	ErrInvalidLengthNiuVote = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNiuVote   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("niu_vote.proto", fileDescriptorNiuVote) }

var fileDescriptorNiuVote = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcb, 0x2c, 0x8d,
	0x2f, 0xcb, 0x2f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2,
	0x4f, 0x4f, 0xcc, 0x4d, 0x8d, 0x4f, 0xce, 0x4f, 0x81, 0x0a, 0x2a, 0xf1, 0x73, 0xf1, 0x3a, 0xfb,
	0xf9, 0xf9, 0x24, 0x96, 0xe6, 0x25, 0x67, 0x84, 0xe5, 0x97, 0xa4, 0x2a, 0xb9, 0x71, 0xf1, 0x06,
	0x23, 0x0b, 0x08, 0x09, 0x71, 0xb1, 0x14, 0xa7, 0x26, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x06, 0x81, 0xd9, 0x42, 0x8a, 0x5c, 0xac, 0xa9, 0x45, 0x45, 0xf9, 0x45, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x7c, 0x46, 0xdc, 0x7a, 0x05, 0x49, 0x7a, 0xae, 0x45, 0x45, 0xce, 0xf9, 0x29, 0xa9, 0x41,
	0x10, 0x19, 0x25, 0x59, 0x2e, 0x76, 0x67, 0x3f, 0x3f, 0x98, 0x09, 0x20, 0x67, 0xc0, 0x4c, 0x00,
	0xb1, 0x95, 0x42, 0xb8, 0xd8, 0x83, 0x71, 0x4b, 0xc3, 0x2d, 0x65, 0xc2, 0x66, 0x29, 0x33, 0x4e,
	0x4b, 0x95, 0xc1, 0x8e, 0x07, 0x99, 0x1a, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x82, 0xcd, 0x6c, 0x27,
	0x9d, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91,
	0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x90,
	0xc4, 0x06, 0x0e, 0x27, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xdc, 0x2f, 0xfd, 0x4e,
	0x01, 0x00, 0x00,
}
