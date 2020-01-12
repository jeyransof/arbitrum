// Code generated by protoc-gen-go. DO NOT EDIT.
// source: valprotocol.proto

package valprotocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	value "github.com/offchainlabs/arbitrum/packages/arb-util/value"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PreconditionBuf struct {
	BeforeHash           *value.HashBuf                `protobuf:"bytes,1,opt,name=beforeHash,proto3" json:"beforeHash,omitempty"`
	TimeBounds           *protocol.TimeBoundsBlocksBuf `protobuf:"bytes,2,opt,name=timeBounds,proto3" json:"timeBounds,omitempty"`
	BeforeInbox          *value.HashBuf                `protobuf:"bytes,3,opt,name=beforeInbox,proto3" json:"beforeInbox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *PreconditionBuf) Reset()         { *m = PreconditionBuf{} }
func (m *PreconditionBuf) String() string { return proto.CompactTextString(m) }
func (*PreconditionBuf) ProtoMessage()    {}
func (*PreconditionBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_564cd27168d44311, []int{0}
}

func (m *PreconditionBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreconditionBuf.Unmarshal(m, b)
}
func (m *PreconditionBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreconditionBuf.Marshal(b, m, deterministic)
}
func (m *PreconditionBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreconditionBuf.Merge(m, src)
}
func (m *PreconditionBuf) XXX_Size() int {
	return xxx_messageInfo_PreconditionBuf.Size(m)
}
func (m *PreconditionBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_PreconditionBuf.DiscardUnknown(m)
}

var xxx_messageInfo_PreconditionBuf proto.InternalMessageInfo

func (m *PreconditionBuf) GetBeforeHash() *value.HashBuf {
	if m != nil {
		return m.BeforeHash
	}
	return nil
}

func (m *PreconditionBuf) GetTimeBounds() *protocol.TimeBoundsBlocksBuf {
	if m != nil {
		return m.TimeBounds
	}
	return nil
}

func (m *PreconditionBuf) GetBeforeInbox() *value.HashBuf {
	if m != nil {
		return m.BeforeInbox
	}
	return nil
}

type ExecutionAssertionStubBuf struct {
	AfterHash            *value.HashBuf `protobuf:"bytes,1,opt,name=afterHash,proto3" json:"afterHash,omitempty"`
	DidInboxInsn         bool           `protobuf:"varint,2,opt,name=didInboxInsn,proto3" json:"didInboxInsn,omitempty"`
	NumGas               uint64         `protobuf:"varint,3,opt,name=numGas,proto3" json:"numGas,omitempty"`
	FirstMessageHash     *value.HashBuf `protobuf:"bytes,4,opt,name=firstMessageHash,proto3" json:"firstMessageHash,omitempty"`
	LastMessageHash      *value.HashBuf `protobuf:"bytes,5,opt,name=lastMessageHash,proto3" json:"lastMessageHash,omitempty"`
	FirstLogHash         *value.HashBuf `protobuf:"bytes,6,opt,name=firstLogHash,proto3" json:"firstLogHash,omitempty"`
	LastLogHash          *value.HashBuf `protobuf:"bytes,7,opt,name=lastLogHash,proto3" json:"lastLogHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecutionAssertionStubBuf) Reset()         { *m = ExecutionAssertionStubBuf{} }
func (m *ExecutionAssertionStubBuf) String() string { return proto.CompactTextString(m) }
func (*ExecutionAssertionStubBuf) ProtoMessage()    {}
func (*ExecutionAssertionStubBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_564cd27168d44311, []int{1}
}

func (m *ExecutionAssertionStubBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionAssertionStubBuf.Unmarshal(m, b)
}
func (m *ExecutionAssertionStubBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionAssertionStubBuf.Marshal(b, m, deterministic)
}
func (m *ExecutionAssertionStubBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionAssertionStubBuf.Merge(m, src)
}
func (m *ExecutionAssertionStubBuf) XXX_Size() int {
	return xxx_messageInfo_ExecutionAssertionStubBuf.Size(m)
}
func (m *ExecutionAssertionStubBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionAssertionStubBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionAssertionStubBuf proto.InternalMessageInfo

func (m *ExecutionAssertionStubBuf) GetAfterHash() *value.HashBuf {
	if m != nil {
		return m.AfterHash
	}
	return nil
}

func (m *ExecutionAssertionStubBuf) GetDidInboxInsn() bool {
	if m != nil {
		return m.DidInboxInsn
	}
	return false
}

func (m *ExecutionAssertionStubBuf) GetNumGas() uint64 {
	if m != nil {
		return m.NumGas
	}
	return 0
}

func (m *ExecutionAssertionStubBuf) GetFirstMessageHash() *value.HashBuf {
	if m != nil {
		return m.FirstMessageHash
	}
	return nil
}

func (m *ExecutionAssertionStubBuf) GetLastMessageHash() *value.HashBuf {
	if m != nil {
		return m.LastMessageHash
	}
	return nil
}

func (m *ExecutionAssertionStubBuf) GetFirstLogHash() *value.HashBuf {
	if m != nil {
		return m.FirstLogHash
	}
	return nil
}

func (m *ExecutionAssertionStubBuf) GetLastLogHash() *value.HashBuf {
	if m != nil {
		return m.LastLogHash
	}
	return nil
}

type AddressBuf struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressBuf) Reset()         { *m = AddressBuf{} }
func (m *AddressBuf) String() string { return proto.CompactTextString(m) }
func (*AddressBuf) ProtoMessage()    {}
func (*AddressBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_564cd27168d44311, []int{2}
}

func (m *AddressBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBuf.Unmarshal(m, b)
}
func (m *AddressBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBuf.Marshal(b, m, deterministic)
}
func (m *AddressBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBuf.Merge(m, src)
}
func (m *AddressBuf) XXX_Size() int {
	return xxx_messageInfo_AddressBuf.Size(m)
}
func (m *AddressBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBuf.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBuf proto.InternalMessageInfo

func (m *AddressBuf) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type TokenTypeBuf struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenTypeBuf) Reset()         { *m = TokenTypeBuf{} }
func (m *TokenTypeBuf) String() string { return proto.CompactTextString(m) }
func (*TokenTypeBuf) ProtoMessage()    {}
func (*TokenTypeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_564cd27168d44311, []int{3}
}

func (m *TokenTypeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenTypeBuf.Unmarshal(m, b)
}
func (m *TokenTypeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenTypeBuf.Marshal(b, m, deterministic)
}
func (m *TokenTypeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTypeBuf.Merge(m, src)
}
func (m *TokenTypeBuf) XXX_Size() int {
	return xxx_messageInfo_TokenTypeBuf.Size(m)
}
func (m *TokenTypeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTypeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTypeBuf proto.InternalMessageInfo

func (m *TokenTypeBuf) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type MessageBuf struct {
	Value                *value.ValueBuf      `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	TokenType            *TokenTypeBuf        `protobuf:"bytes,2,opt,name=tokenType,proto3" json:"tokenType,omitempty"`
	Amount               *value.BigIntegerBuf `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender               *AddressBuf          `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MessageBuf) Reset()         { *m = MessageBuf{} }
func (m *MessageBuf) String() string { return proto.CompactTextString(m) }
func (*MessageBuf) ProtoMessage()    {}
func (*MessageBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_564cd27168d44311, []int{4}
}

func (m *MessageBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageBuf.Unmarshal(m, b)
}
func (m *MessageBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageBuf.Marshal(b, m, deterministic)
}
func (m *MessageBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageBuf.Merge(m, src)
}
func (m *MessageBuf) XXX_Size() int {
	return xxx_messageInfo_MessageBuf.Size(m)
}
func (m *MessageBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageBuf.DiscardUnknown(m)
}

var xxx_messageInfo_MessageBuf proto.InternalMessageInfo

func (m *MessageBuf) GetValue() *value.ValueBuf {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MessageBuf) GetTokenType() *TokenTypeBuf {
	if m != nil {
		return m.TokenType
	}
	return nil
}

func (m *MessageBuf) GetAmount() *value.BigIntegerBuf {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MessageBuf) GetSender() *AddressBuf {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*PreconditionBuf)(nil), "valprotocol.PreconditionBuf")
	proto.RegisterType((*ExecutionAssertionStubBuf)(nil), "valprotocol.ExecutionAssertionStubBuf")
	proto.RegisterType((*AddressBuf)(nil), "valprotocol.AddressBuf")
	proto.RegisterType((*TokenTypeBuf)(nil), "valprotocol.TokenTypeBuf")
	proto.RegisterType((*MessageBuf)(nil), "valprotocol.MessageBuf")
}

func init() { proto.RegisterFile("valprotocol.proto", fileDescriptor_564cd27168d44311) }

var fileDescriptor_564cd27168d44311 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xd9, 0xda, 0x46, 0x7b, 0x76, 0x71, 0x75, 0x28, 0x76, 0x5b, 0x10, 0x24, 0x28, 0x78,
	0x51, 0x13, 0x59, 0x2f, 0x14, 0xc1, 0x8b, 0x46, 0x8a, 0x2e, 0x28, 0x48, 0x5c, 0xbc, 0xf0, 0x6e,
	0x92, 0x9c, 0x64, 0x87, 0x4d, 0x66, 0x96, 0xf9, 0x53, 0xea, 0x6b, 0xf9, 0x1a, 0x3e, 0x85, 0x6f,
	0x22, 0x33, 0x99, 0xdd, 0x4d, 0x34, 0xf6, 0x26, 0x4c, 0xce, 0xf9, 0x7d, 0xdf, 0xf9, 0x38, 0x33,
	0xf0, 0xf0, 0x9a, 0xd6, 0x1b, 0x29, 0xb4, 0xc8, 0x45, 0x1d, 0xb9, 0x03, 0x19, 0x77, 0x4a, 0xe7,
	0xb6, 0x6f, 0x30, 0x76, 0xdf, 0xb6, 0x7f, 0x7e, 0xba, 0x6d, 0xc6, 0x7d, 0x61, 0xf8, 0x73, 0x04,
	0xd3, 0x2f, 0x12, 0x73, 0xc1, 0x0b, 0xa6, 0x99, 0xe0, 0x89, 0x29, 0x49, 0x04, 0x90, 0x61, 0x29,
	0x24, 0x7e, 0xa4, 0x6a, 0x35, 0x1b, 0x3d, 0x19, 0x3d, 0x1f, 0xcf, 0xef, 0x47, 0xad, 0x9d, 0x2d,
	0x25, 0xa6, 0x4c, 0x3b, 0x04, 0x79, 0x07, 0xa0, 0x59, 0x83, 0x89, 0x30, 0xbc, 0x50, 0xb3, 0x03,
	0xc7, 0x3f, 0x8e, 0x76, 0x83, 0x96, 0xbb, 0x5e, 0x52, 0x8b, 0x7c, 0xad, 0x9c, 0x7c, 0x2f, 0x20,
	0x2f, 0x61, 0xdc, 0x9a, 0x2d, 0x78, 0x26, 0x6e, 0x66, 0x77, 0x06, 0xe7, 0x75, 0x91, 0xf0, 0xf7,
	0x01, 0x9c, 0x5d, 0xdd, 0x60, 0x6e, 0x6c, 0xe2, 0x4b, 0xa5, 0x50, 0xda, 0xc3, 0x57, 0x6d, 0x32,
	0x1b, 0xff, 0x02, 0x8e, 0x69, 0xa9, 0x51, 0xde, 0x92, 0x7e, 0x0f, 0x90, 0x10, 0x26, 0x05, 0x2b,
	0x9c, 0xef, 0x82, 0x2b, 0xee, 0xe2, 0xdf, 0x4b, 0x7b, 0x35, 0xf2, 0x08, 0x02, 0x6e, 0x9a, 0x0f,
	0x54, 0xb9, 0x70, 0x87, 0xa9, 0xff, 0x23, 0x6f, 0xe1, 0x41, 0xc9, 0xa4, 0xd2, 0x9f, 0x51, 0x29,
	0x5a, 0xb5, 0xeb, 0x3a, 0x1c, 0x1c, 0xf8, 0x0f, 0x47, 0xde, 0xc0, 0xb4, 0xa6, 0x7d, 0xe9, 0xd1,
	0xa0, 0xf4, 0x6f, 0x8c, 0xcc, 0x61, 0xe2, 0xdc, 0x3e, 0x89, 0xca, 0xc9, 0x82, 0x41, 0x59, 0x8f,
	0xb1, 0x3b, 0xb6, 0x36, 0x5b, 0xc9, 0xdd, 0xe1, 0x1d, 0x77, 0x90, 0x30, 0x04, 0xb8, 0x2c, 0x0a,
	0x89, 0xca, 0xde, 0x17, 0x39, 0x81, 0x23, 0xc7, 0xba, 0x7d, 0x4e, 0xd2, 0xf6, 0x27, 0x7c, 0x0a,
	0x93, 0xa5, 0x58, 0x23, 0x5f, 0xfe, 0xd8, 0xe0, 0xff, 0xa9, 0x5f, 0x23, 0x00, 0x9f, 0xdf, 0x42,
	0xcf, 0xba, 0xd0, 0x78, 0x3e, 0xf5, 0x21, 0xbe, 0xd9, 0xaf, 0x4d, 0xd1, 0x76, 0xc9, 0x6b, 0x38,
	0xd6, 0x5b, 0x6f, 0xff, 0xa6, 0xce, 0xa2, 0xee, 0xc3, 0xef, 0x4e, 0x4e, 0xf7, 0x2c, 0xb9, 0x80,
	0x80, 0x36, 0xc2, 0x70, 0xed, 0x5f, 0xd2, 0x89, 0x1f, 0x90, 0xb0, 0x6a, 0xc1, 0x35, 0x56, 0x28,
	0xad, 0xc0, 0x33, 0x24, 0x86, 0x40, 0x21, 0x2f, 0x50, 0xfa, 0x8b, 0x3b, 0xed, 0xcd, 0xd8, 0x6f,
	0x20, 0xf5, 0x58, 0x72, 0xf5, 0xfd, 0x7d, 0xc5, 0xf4, 0xca, 0x64, 0x51, 0x2e, 0x9a, 0x58, 0x94,
	0x65, 0xbe, 0xa2, 0x8c, 0xd7, 0x34, 0x53, 0x31, 0x95, 0x19, 0xd3, 0xd2, 0x34, 0xf1, 0x86, 0xe6,
	0x6b, 0x5a, 0xa1, 0xab, 0xbc, 0xb8, 0xa6, 0x35, 0x2b, 0xa8, 0x16, 0x32, 0xee, 0x38, 0x67, 0x81,
	0x3b, 0xbd, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x65, 0xaf, 0x57, 0xcc, 0x03, 0x00, 0x00,
}
