// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SetGroupMessagRule.proto

package SetGroupMessagRule

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

type SetGroupMessageRuleInner struct {
	Groupid              *uint64  `protobuf:"varint,1,req,name=groupid" json:"groupid,omitempty"`
	IsNotice             *uint32  `protobuf:"varint,2,opt,name=isNotice" json:"isNotice,omitempty"`
	IsApplePush          *uint32  `protobuf:"varint,3,opt,name=isApplePush" json:"isApplePush,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGroupMessageRuleInner) Reset()         { *m = SetGroupMessageRuleInner{} }
func (m *SetGroupMessageRuleInner) String() string { return proto.CompactTextString(m) }
func (*SetGroupMessageRuleInner) ProtoMessage()    {}
func (*SetGroupMessageRuleInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_50617db37d25a421, []int{0}
}

func (m *SetGroupMessageRuleInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGroupMessageRuleInner.Unmarshal(m, b)
}
func (m *SetGroupMessageRuleInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGroupMessageRuleInner.Marshal(b, m, deterministic)
}
func (m *SetGroupMessageRuleInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGroupMessageRuleInner.Merge(m, src)
}
func (m *SetGroupMessageRuleInner) XXX_Size() int {
	return xxx_messageInfo_SetGroupMessageRuleInner.Size(m)
}
func (m *SetGroupMessageRuleInner) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGroupMessageRuleInner.DiscardUnknown(m)
}

var xxx_messageInfo_SetGroupMessageRuleInner proto.InternalMessageInfo

func (m *SetGroupMessageRuleInner) GetGroupid() uint64 {
	if m != nil && m.Groupid != nil {
		return *m.Groupid
	}
	return 0
}

func (m *SetGroupMessageRuleInner) GetIsNotice() uint32 {
	if m != nil && m.IsNotice != nil {
		return *m.IsNotice
	}
	return 0
}

func (m *SetGroupMessageRuleInner) GetIsApplePush() uint32 {
	if m != nil && m.IsApplePush != nil {
		return *m.IsApplePush
	}
	return 0
}

func init() {
	proto.RegisterType((*SetGroupMessageRuleInner)(nil), "SetGroupMessageRuleInner")
}

func init() { proto.RegisterFile("SetGroupMessagRule.proto", fileDescriptor_50617db37d25a421) }

var fileDescriptor_50617db37d25a421 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x08, 0x4e, 0x2d, 0x71,
	0x2f, 0xca, 0x2f, 0x2d, 0xf0, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x0f, 0x2a, 0xcd, 0x49, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0x2a, 0x42, 0x97, 0x4b, 0x05, 0x49, 0x7a, 0xe6, 0xe5, 0xa5, 0x16,
	0x09, 0x49, 0x70, 0xb1, 0xa7, 0x83, 0x24, 0x32, 0x53, 0x24, 0x18, 0x15, 0x98, 0x34, 0x58, 0x82,
	0x60, 0x5c, 0x21, 0x29, 0x2e, 0x8e, 0xcc, 0x62, 0xbf, 0xfc, 0x92, 0xcc, 0xe4, 0x54, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xde, 0x20, 0x38, 0x5f, 0x48, 0x81, 0x8b, 0x3b, 0xb3, 0xd8, 0xb1, 0xa0, 0x20,
	0x27, 0x35, 0xa0, 0xb4, 0x38, 0x43, 0x82, 0x19, 0x2c, 0x8d, 0x2c, 0xe4, 0xa4, 0xc5, 0xa5, 0x96,
	0x9c, 0x9f, 0xab, 0x57, 0x59, 0x9a, 0x57, 0x92, 0x9f, 0x97, 0x5e, 0x51, 0x9a, 0xa7, 0x97, 0x9c,
	0x9f, 0x97, 0x97, 0x9a, 0x5c, 0x92, 0x5f, 0xa4, 0x97, 0x9c, 0x5f, 0x04, 0x75, 0x5a, 0x52, 0x69,
	0x9a, 0x07, 0x33, 0x20, 0x00, 0x00, 0xff, 0xff, 0x98, 0xeb, 0x3a, 0x13, 0xba, 0x00, 0x00, 0x00,
}
