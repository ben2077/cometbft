// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/p2p/conn.proto

package p2p

import (
	fmt "fmt"
	crypto "github.com/ben2077/cometbft/proto/tendermint/crypto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PacketPing struct {
}

func (m *PacketPing) Reset()         { *m = PacketPing{} }
func (m *PacketPing) String() string { return proto.CompactTextString(m) }
func (*PacketPing) ProtoMessage()    {}
func (*PacketPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_22474b5527c8fa9f, []int{0}
}
func (m *PacketPing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketPing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketPing.Merge(m, src)
}
func (m *PacketPing) XXX_Size() int {
	return m.Size()
}
func (m *PacketPing) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketPing.DiscardUnknown(m)
}

var xxx_messageInfo_PacketPing proto.InternalMessageInfo

type PacketPong struct {
}

func (m *PacketPong) Reset()         { *m = PacketPong{} }
func (m *PacketPong) String() string { return proto.CompactTextString(m) }
func (*PacketPong) ProtoMessage()    {}
func (*PacketPong) Descriptor() ([]byte, []int) {
	return fileDescriptor_22474b5527c8fa9f, []int{1}
}
func (m *PacketPong) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketPong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketPong.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketPong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketPong.Merge(m, src)
}
func (m *PacketPong) XXX_Size() int {
	return m.Size()
}
func (m *PacketPong) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketPong.DiscardUnknown(m)
}

var xxx_messageInfo_PacketPong proto.InternalMessageInfo

type PacketMsg struct {
	ChannelID int32  `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	EOF       bool   `protobuf:"varint,2,opt,name=eof,proto3" json:"eof,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PacketMsg) Reset()         { *m = PacketMsg{} }
func (m *PacketMsg) String() string { return proto.CompactTextString(m) }
func (*PacketMsg) ProtoMessage()    {}
func (*PacketMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_22474b5527c8fa9f, []int{2}
}
func (m *PacketMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketMsg.Merge(m, src)
}
func (m *PacketMsg) XXX_Size() int {
	return m.Size()
}
func (m *PacketMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PacketMsg proto.InternalMessageInfo

func (m *PacketMsg) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *PacketMsg) GetEOF() bool {
	if m != nil {
		return m.EOF
	}
	return false
}

func (m *PacketMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Packet struct {
	// Types that are valid to be assigned to Sum:
	//
	//	*Packet_PacketPing
	//	*Packet_PacketPong
	//	*Packet_PacketMsg
	Sum isPacket_Sum `protobuf_oneof:"sum"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_22474b5527c8fa9f, []int{3}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return m.Size()
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

type isPacket_Sum interface {
	isPacket_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Packet_PacketPing struct {
	PacketPing *PacketPing `protobuf:"bytes,1,opt,name=packet_ping,json=packetPing,proto3,oneof" json:"packet_ping,omitempty"`
}
type Packet_PacketPong struct {
	PacketPong *PacketPong `protobuf:"bytes,2,opt,name=packet_pong,json=packetPong,proto3,oneof" json:"packet_pong,omitempty"`
}
type Packet_PacketMsg struct {
	PacketMsg *PacketMsg `protobuf:"bytes,3,opt,name=packet_msg,json=packetMsg,proto3,oneof" json:"packet_msg,omitempty"`
}

func (*Packet_PacketPing) isPacket_Sum() {}
func (*Packet_PacketPong) isPacket_Sum() {}
func (*Packet_PacketMsg) isPacket_Sum()  {}

func (m *Packet) GetSum() isPacket_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Packet) GetPacketPing() *PacketPing {
	if x, ok := m.GetSum().(*Packet_PacketPing); ok {
		return x.PacketPing
	}
	return nil
}

func (m *Packet) GetPacketPong() *PacketPong {
	if x, ok := m.GetSum().(*Packet_PacketPong); ok {
		return x.PacketPong
	}
	return nil
}

func (m *Packet) GetPacketMsg() *PacketMsg {
	if x, ok := m.GetSum().(*Packet_PacketMsg); ok {
		return x.PacketMsg
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Packet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Packet_PacketPing)(nil),
		(*Packet_PacketPong)(nil),
		(*Packet_PacketMsg)(nil),
	}
}

type AuthSigMessage struct {
	PubKey crypto.PublicKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	Sig    []byte           `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *AuthSigMessage) Reset()         { *m = AuthSigMessage{} }
func (m *AuthSigMessage) String() string { return proto.CompactTextString(m) }
func (*AuthSigMessage) ProtoMessage()    {}
func (*AuthSigMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_22474b5527c8fa9f, []int{4}
}
func (m *AuthSigMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthSigMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthSigMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthSigMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSigMessage.Merge(m, src)
}
func (m *AuthSigMessage) XXX_Size() int {
	return m.Size()
}
func (m *AuthSigMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSigMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSigMessage proto.InternalMessageInfo

func (m *AuthSigMessage) GetPubKey() crypto.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return crypto.PublicKey{}
}

func (m *AuthSigMessage) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*PacketPing)(nil), "tendermint.p2p.PacketPing")
	proto.RegisterType((*PacketPong)(nil), "tendermint.p2p.PacketPong")
	proto.RegisterType((*PacketMsg)(nil), "tendermint.p2p.PacketMsg")
	proto.RegisterType((*Packet)(nil), "tendermint.p2p.Packet")
	proto.RegisterType((*AuthSigMessage)(nil), "tendermint.p2p.AuthSigMessage")
}

func init() { proto.RegisterFile("tendermint/p2p/conn.proto", fileDescriptor_22474b5527c8fa9f) }

var fileDescriptor_22474b5527c8fa9f = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xf5, 0xe2, 0x36, 0x25, 0x93, 0x50, 0xa1, 0x15, 0x87, 0x24, 0xaa, 0x9c, 0x28, 0xa7, 0x1c,
	0x90, 0x5d, 0x99, 0x43, 0x25, 0x10, 0x07, 0xcc, 0x87, 0x28, 0x55, 0x44, 0x64, 0x6e, 0x5c, 0x2c,
	0x7f, 0x6c, 0xd7, 0xab, 0xd4, 0xbb, 0xab, 0xec, 0xfa, 0xe0, 0x7f, 0xc1, 0xcf, 0x2a, 0xb7, 0x1e,
	0x39, 0x45, 0xc8, 0xf9, 0x23, 0xc8, 0xde, 0x40, 0x1c, 0x89, 0xde, 0xde, 0x9b, 0xd9, 0x37, 0xf3,
	0x9e, 0x76, 0x60, 0xac, 0x09, 0xcf, 0xc8, 0xa6, 0x60, 0x5c, 0x7b, 0xd2, 0x97, 0x5e, 0x2a, 0x38,
	0x77, 0xe5, 0x46, 0x68, 0x81, 0xcf, 0x0f, 0x2d, 0x57, 0xfa, 0x72, 0xf2, 0x82, 0x0a, 0x2a, 0xda,
	0x96, 0xd7, 0x20, 0xf3, 0x6a, 0x72, 0xd1, 0x19, 0x90, 0x6e, 0x2a, 0xa9, 0x85, 0xb7, 0x26, 0x95,
	0x32, 0xdd, 0xf9, 0x10, 0x60, 0x15, 0xa7, 0x6b, 0xa2, 0x57, 0x8c, 0xd3, 0x0e, 0x13, 0x9c, 0xce,
	0x73, 0xe8, 0x1b, 0xb6, 0x54, 0x14, 0xbf, 0x04, 0x48, 0xf3, 0x98, 0x73, 0x72, 0x17, 0xb1, 0x6c,
	0x84, 0x66, 0x68, 0x71, 0x1a, 0x3c, 0xab, 0xb7, 0xd3, 0xfe, 0x7b, 0x53, 0xbd, 0xfe, 0x10, 0xf6,
	0xf7, 0x0f, 0xae, 0x33, 0x3c, 0x06, 0x9b, 0x88, 0xdb, 0xd1, 0x93, 0x19, 0x5a, 0x3c, 0x0d, 0xce,
	0xea, 0xed, 0xd4, 0xfe, 0xf8, 0xf5, 0x53, 0xd8, 0xd4, 0x30, 0x86, 0x93, 0x2c, 0xd6, 0xf1, 0xc8,
	0x9e, 0xa1, 0xc5, 0x30, 0x6c, 0xf1, 0xfc, 0x27, 0x82, 0x9e, 0x59, 0x85, 0xdf, 0xc2, 0x40, 0xb6,
	0x28, 0x92, 0x8c, 0xd3, 0x76, 0xd1, 0xc0, 0x9f, 0xb8, 0xc7, 0x51, 0xdd, 0x83, 0xe7, 0xcf, 0x56,
	0x08, 0xf2, 0x1f, 0xeb, 0xca, 0x05, 0xa7, 0xad, 0x81, 0xc7, 0xe5, 0xe2, 0x48, 0x2e, 0x38, 0xc5,
	0xaf, 0x61, 0xcf, 0xa2, 0x42, 0xd1, 0xd6, 0xe2, 0xc0, 0x1f, 0xff, 0x5f, 0xbd, 0x54, 0x8d, 0xb8,
	0x2f, 0xff, 0x92, 0xe0, 0x14, 0x6c, 0x55, 0x16, 0xf3, 0x08, 0xce, 0xdf, 0x95, 0x3a, 0xff, 0xc6,
	0xe8, 0x92, 0x28, 0x15, 0x53, 0x82, 0xdf, 0xc0, 0x99, 0x2c, 0x93, 0x68, 0x4d, 0xaa, 0x7d, 0x9c,
	0x8b, 0xee, 0x44, 0xf3, 0x27, 0xee, 0xaa, 0x4c, 0xee, 0x58, 0x7a, 0x43, 0xaa, 0xe0, 0xe4, 0x7e,
	0x3b, 0xb5, 0xc2, 0x9e, 0x2c, 0x93, 0x1b, 0x52, 0xe1, 0xe7, 0x60, 0x2b, 0x66, 0x82, 0x0c, 0xc3,
	0x06, 0x06, 0x5f, 0xee, 0x6b, 0x07, 0x3d, 0xd4, 0x0e, 0xfa, 0x5d, 0x3b, 0xe8, 0xc7, 0xce, 0xb1,
	0x1e, 0x76, 0x8e, 0xf5, 0x6b, 0xe7, 0x58, 0xdf, 0x2f, 0x29, 0xd3, 0x79, 0x99, 0xb8, 0xa9, 0x28,
	0xbc, 0x84, 0x70, 0xff, 0xf2, 0xea, 0xca, 0x4b, 0x45, 0x41, 0x74, 0x72, 0xab, 0x3d, 0x73, 0x18,
	0xc7, 0xd7, 0x94, 0xf4, 0xda, 0xea, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xab, 0xf5,
	0x3a, 0x66, 0x02, 0x00, 0x00,
}

func (m *PacketPing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketPing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketPing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PacketPong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketPong) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketPong) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PacketMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintConn(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.EOF {
		i--
		if m.EOF {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.ChannelID != 0 {
		i = encodeVarintConn(dAtA, i, uint64(m.ChannelID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Packet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Packet_PacketPing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet_PacketPing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PacketPing != nil {
		{
			size, err := m.PacketPing.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConn(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Packet_PacketPong) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet_PacketPong) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PacketPong != nil {
		{
			size, err := m.PacketPong.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConn(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Packet_PacketMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet_PacketMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PacketMsg != nil {
		{
			size, err := m.PacketMsg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConn(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *AuthSigMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthSigMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthSigMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintConn(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConn(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintConn(dAtA []byte, offset int, v uint64) int {
	offset -= sovConn(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PacketPing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PacketPong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PacketMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChannelID != 0 {
		n += 1 + sovConn(uint64(m.ChannelID))
	}
	if m.EOF {
		n += 2
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovConn(uint64(l))
	}
	return n
}

func (m *Packet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Packet_PacketPing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketPing != nil {
		l = m.PacketPing.Size()
		n += 1 + l + sovConn(uint64(l))
	}
	return n
}
func (m *Packet_PacketPong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketPong != nil {
		l = m.PacketPong.Size()
		n += 1 + l + sovConn(uint64(l))
	}
	return n
}
func (m *Packet_PacketMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketMsg != nil {
		l = m.PacketMsg.Size()
		n += 1 + l + sovConn(uint64(l))
	}
	return n
}
func (m *AuthSigMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PubKey.Size()
	n += 1 + l + sovConn(uint64(l))
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovConn(uint64(l))
	}
	return n
}

func sovConn(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConn(x uint64) (n int) {
	return sovConn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PacketPing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConn
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
			return fmt.Errorf("proto: PacketPing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketPing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConn
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
func (m *PacketPong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConn
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
			return fmt.Errorf("proto: PacketPong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketPong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConn
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
func (m *PacketMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConn
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
			return fmt.Errorf("proto: PacketMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			m.ChannelID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EOF", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
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
			m.EOF = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
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
				return ErrInvalidLengthConn
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConn
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
			skippy, err := skipConn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConn
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
func (m *Packet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConn
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
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketPing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PacketPing{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Packet_PacketPing{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketPong", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PacketPong{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Packet_PacketPong{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PacketMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Packet_PacketMsg{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConn
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
func (m *AuthSigMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConn
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
			return fmt.Errorf("proto: AuthSigMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthSigMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConn
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
				return ErrInvalidLengthConn
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConn
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
func skipConn(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConn
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
					return 0, ErrIntOverflowConn
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConn
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
				return 0, ErrInvalidLengthConn
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConn
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConn
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConn        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConn          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConn = fmt.Errorf("proto: unexpected end of group")
)
