// Code generated by protoc-gen-go. DO NOT EDIT.
// source: PushMsgNotify.proto

package PushMsgNotify

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PushMsgNotifyInner struct {
	Version              *uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMsgNotifyInner) Reset()         { *m = PushMsgNotifyInner{} }
func (m *PushMsgNotifyInner) String() string { return proto.CompactTextString(m) }
func (*PushMsgNotifyInner) ProtoMessage()    {}
func (*PushMsgNotifyInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_6049aa891b23ef61, []int{0}
}

func (m *PushMsgNotifyInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMsgNotifyInner.Unmarshal(m, b)
}
func (m *PushMsgNotifyInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMsgNotifyInner.Marshal(b, m, deterministic)
}
func (m *PushMsgNotifyInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgNotifyInner.Merge(m, src)
}
func (m *PushMsgNotifyInner) XXX_Size() int {
	return xxx_messageInfo_PushMsgNotifyInner.Size(m)
}
func (m *PushMsgNotifyInner) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgNotifyInner.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgNotifyInner proto.InternalMessageInfo

func (m *PushMsgNotifyInner) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*PushMsgNotifyInner)(nil), "PushMsgNotifyInner")
}

func init() { proto.RegisterFile("PushMsgNotify.proto", fileDescriptor_6049aa891b23ef61) }

var fileDescriptor_6049aa891b23ef61 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0e, 0x28, 0x2d, 0xce,
	0xf0, 0x2d, 0x4e, 0xf7, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57,
	0xd2, 0xe3, 0x12, 0x42, 0x11, 0xf6, 0xcc, 0xcb, 0x4b, 0x2d, 0x12, 0x92, 0xe0, 0x62, 0x2f, 0x4b,
	0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x82, 0x71, 0x9d, 0xb4,
	0xb8, 0xd4, 0x92, 0xf3, 0x73, 0xf5, 0x2a, 0x4b, 0xf3, 0x4a, 0xf2, 0xf3, 0xd2, 0x2b, 0x4a, 0xf3,
	0xf4, 0x92, 0xf3, 0xf3, 0xf2, 0x52, 0x93, 0x4b, 0xf2, 0x8b, 0xf4, 0x92, 0xf3, 0x8b, 0x52, 0x21,
	0xc6, 0x26, 0x95, 0xa6, 0x79, 0x30, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x42, 0xb9, 0x50,
	0x71, 0x00, 0x00, 0x00,
}