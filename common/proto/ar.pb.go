// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ar.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Type int32

const (
	Type_Hello            Type = 0
	Type_CantFind         Type = 1
	Type_OK               Type = 2
	Type_PassWay          Type = 3
	Type_ConnectClosed    Type = 4
	Type_ConnectionFrom   Type = 5
	Type_Ping             Type = 6
	Type_CantConnectSocks Type = 7
	Type_WantToConnection Type = 8
)

var Type_name = map[int32]string{
	0: "Hello",
	1: "CantFind",
	2: "OK",
	3: "PassWay",
	4: "ConnectClosed",
	5: "ConnectionFrom",
	6: "Ping",
	7: "CantConnectSocks",
	8: "WantToConnection",
}

var Type_value = map[string]int32{
	"Hello":            0,
	"CantFind":         1,
	"OK":               2,
	"PassWay":          3,
	"ConnectClosed":    4,
	"ConnectionFrom":   5,
	"Ping":             6,
	"CantConnectSocks": 7,
	"WantToConnection": 8,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c2a5d8a61e0e8f0, []int{0}
}

type Pkg struct {
	From                 string   `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	Type                 Type     `protobuf:"varint,3,opt,name=Type,proto3,enum=anywhere_go.common.proto.Type" json:"Type,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pkg) Reset()         { *m = Pkg{} }
func (m *Pkg) String() string { return proto.CompactTextString(m) }
func (*Pkg) ProtoMessage()    {}
func (*Pkg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c2a5d8a61e0e8f0, []int{0}
}
func (m *Pkg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pkg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pkg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pkg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pkg.Merge(m, src)
}
func (m *Pkg) XXX_Size() int {
	return m.Size()
}
func (m *Pkg) XXX_DiscardUnknown() {
	xxx_messageInfo_Pkg.DiscardUnknown(m)
}

var xxx_messageInfo_Pkg proto.InternalMessageInfo

func (m *Pkg) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Pkg) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Pkg) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_Hello
}

func (m *Pkg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Msg struct {
	Chan                 int32    `protobuf:"varint,1,opt,name=Chan,proto3" json:"Chan,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Close                bool     `protobuf:"varint,3,opt,name=close,proto3" json:"close,omitempty"`
	From                 string   `protobuf:"bytes,4,opt,name=From,proto3" json:"From,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c2a5d8a61e0e8f0, []int{1}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return m.Size()
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetChan() int32 {
	if m != nil {
		return m.Chan
	}
	return 0
}

func (m *Msg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Msg) GetClose() bool {
	if m != nil {
		return m.Close
	}
	return false
}

func (m *Msg) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func init() {
	proto.RegisterEnum("anywhere_go.common.proto.Type", Type_name, Type_value)
	proto.RegisterType((*Pkg)(nil), "anywhere_go.common.proto.Pkg")
	proto.RegisterType((*Msg)(nil), "anywhere_go.common.proto.Msg")
}

func init() { proto.RegisterFile("ar.proto", fileDescriptor_1c2a5d8a61e0e8f0) }

var fileDescriptor_1c2a5d8a61e0e8f0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x3b, 0x93, 0x49, 0x9b, 0xfe, 0xd6, 0x32, 0xfe, 0x14, 0xcc, 0x2a, 0x94, 0xae, 0x8a,
	0x8b, 0x2c, 0xea, 0x0d, 0x8c, 0x14, 0x41, 0xc4, 0x12, 0x03, 0x05, 0x37, 0x32, 0xa6, 0xa1, 0x2d,
	0x6d, 0xe7, 0xaf, 0x49, 0x40, 0x7a, 0x0c, 0x77, 0x1e, 0xc9, 0xa5, 0x47, 0x90, 0x78, 0x11, 0x99,
	0x49, 0x69, 0x56, 0xae, 0xf2, 0xf2, 0xe6, 0xcd, 0xf7, 0x1e, 0x03, 0x9e, 0xca, 0xc3, 0x7d, 0x4e,
	0x25, 0xa1, 0xaf, 0xf4, 0xe1, 0x7d, 0x95, 0xe5, 0xd9, 0xcb, 0x92, 0xc2, 0x94, 0x76, 0x3b, 0xd2,
	0xf5, 0xc9, 0xe8, 0x0d, 0x9c, 0xd9, 0x66, 0x89, 0x08, 0x62, 0x9a, 0xd3, 0xce, 0x67, 0x43, 0x36,
	0xee, 0xc6, 0x56, 0x63, 0x1f, 0x78, 0x42, 0x3e, 0xb7, 0x0e, 0x4f, 0x08, 0x27, 0x20, 0x92, 0xc3,
	0x3e, 0xf3, 0x9d, 0x21, 0x1b, 0xf7, 0x27, 0x41, 0xf8, 0x1f, 0x33, 0x34, 0xa9, 0xd8, 0x66, 0x0d,
	0xf7, 0x56, 0x95, 0xca, 0x17, 0x43, 0x36, 0xee, 0xc5, 0x56, 0x8f, 0xe6, 0xe0, 0x3c, 0x14, 0xb6,
	0x32, 0x5a, 0x29, 0x6d, 0x2b, 0xdd, 0xd8, 0xea, 0x53, 0x9c, 0x37, 0x71, 0x1c, 0x80, 0x9b, 0x6e,
	0xa9, 0xa8, 0x7b, 0xbd, 0xb8, 0xfe, 0x39, 0x0d, 0x16, 0xcd, 0xe0, 0xab, 0x0f, 0x56, 0x2f, 0xc4,
	0x2e, 0xb8, 0x77, 0xd9, 0x76, 0x4b, 0xb2, 0x85, 0x3d, 0xf0, 0x22, 0xa5, 0xcb, 0xe9, 0x5a, 0x2f,
	0x24, 0xc3, 0x36, 0xf0, 0xc7, 0x7b, 0xc9, 0xf1, 0x0c, 0x3a, 0x33, 0x55, 0x14, 0x73, 0x75, 0x90,
	0x0e, 0x5e, 0xc0, 0x79, 0x44, 0x5a, 0x67, 0x69, 0x19, 0x19, 0xf4, 0x42, 0x0a, 0x44, 0xe8, 0x1f,
	0xad, 0x35, 0x69, 0xc3, 0x96, 0x2e, 0x7a, 0x20, 0x66, 0x6b, 0xbd, 0x94, 0x6d, 0x1c, 0x80, 0x34,
	0xcc, 0x63, 0xe2, 0x89, 0xd2, 0x4d, 0x21, 0x3b, 0xc6, 0x9d, 0x2b, 0x5d, 0x26, 0xd4, 0xdc, 0x94,
	0xde, 0xcd, 0xe5, 0x57, 0x15, 0xb0, 0xef, 0x2a, 0x60, 0x3f, 0x55, 0xc0, 0x3e, 0x7f, 0x83, 0xd6,
	0xb3, 0x6b, 0x1f, 0xe9, 0xb5, 0x6d, 0x3f, 0xd7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x78,
	0x66, 0x3d, 0xa5, 0x01, 0x00, 0x00,
}

func (m *Pkg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pkg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAr(dAtA, i, uint64(len(m.From)))
		i += copy(dAtA[i:], m.From)
	}
	if len(m.To) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAr(dAtA, i, uint64(len(m.To)))
		i += copy(dAtA[i:], m.To)
	}
	if m.Type != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAr(dAtA, i, uint64(m.Type))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAr(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Msg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Msg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Chan != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAr(dAtA, i, uint64(m.Chan))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAr(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.Close {
		dAtA[i] = 0x18
		i++
		if m.Close {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.From) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAr(dAtA, i, uint64(len(m.From)))
		i += copy(dAtA[i:], m.From)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAr(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Pkg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovAr(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovAr(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovAr(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAr(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Msg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chan != 0 {
		n += 1 + sovAr(uint64(m.Chan))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAr(uint64(l))
	}
	if m.Close {
		n += 2
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovAr(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAr(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAr(x uint64) (n int) {
	return sovAr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pkg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAr
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pkg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pkg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAr
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAr
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAr
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Msg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAr
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Msg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Msg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chan", wireType)
			}
			m.Chan = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chan |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAr
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Close", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Close = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAr
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAr
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAr
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
					return 0, ErrIntOverflowAr
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
					return 0, ErrIntOverflowAr
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
			if length < 0 {
				return 0, ErrInvalidLengthAr
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAr
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAr
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
				next, err := skipAr(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAr
				}
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
	ErrInvalidLengthAr = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAr   = fmt.Errorf("proto: integer overflow")
)