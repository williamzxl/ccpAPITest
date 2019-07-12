// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GetGroupMembers.proto

package GetGroupMembers

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

type GetGroupMembersInner struct {
	GroupId              *uint64  `protobuf:"varint,1,req,name=groupId" json:"groupId,omitempty"`
	Asker                *uint64  `protobuf:"varint,2,opt,name=asker" json:"asker,omitempty"`
	MemAcc               *string  `protobuf:"bytes,3,opt,name=memAcc" json:"memAcc,omitempty"`
	PageSize             *uint32  `protobuf:"varint,4,opt,name=pageSize" json:"pageSize,omitempty"`
	PageNum              *uint32  `protobuf:"varint,5,opt,name=pageNum" json:"pageNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupMembersInner) Reset()         { *m = GetGroupMembersInner{} }
func (m *GetGroupMembersInner) String() string { return proto.CompactTextString(m) }
func (*GetGroupMembersInner) ProtoMessage()    {}
func (*GetGroupMembersInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abc595d02732660, []int{0}
}

func (m *GetGroupMembersInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupMembersInner.Unmarshal(m, b)
}
func (m *GetGroupMembersInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupMembersInner.Marshal(b, m, deterministic)
}
func (m *GetGroupMembersInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupMembersInner.Merge(m, src)
}
func (m *GetGroupMembersInner) XXX_Size() int {
	return xxx_messageInfo_GetGroupMembersInner.Size(m)
}
func (m *GetGroupMembersInner) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupMembersInner.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupMembersInner proto.InternalMessageInfo

func (m *GetGroupMembersInner) GetGroupId() uint64 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *GetGroupMembersInner) GetAsker() uint64 {
	if m != nil && m.Asker != nil {
		return *m.Asker
	}
	return 0
}

func (m *GetGroupMembersInner) GetMemAcc() string {
	if m != nil && m.MemAcc != nil {
		return *m.MemAcc
	}
	return ""
}

func (m *GetGroupMembersInner) GetPageSize() uint32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

func (m *GetGroupMembersInner) GetPageNum() uint32 {
	if m != nil && m.PageNum != nil {
		return *m.PageNum
	}
	return 0
}

func init() {
	proto.RegisterType((*GetGroupMembersInner)(nil), "GetGroupMembersInner")
}

func init() { proto.RegisterFile("GetGroupMembers.proto", fileDescriptor_0abc595d02732660) }

var fileDescriptor_0abc595d02732660 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x75, 0x4f, 0x2d, 0x71,
	0x2f, 0xca, 0x2f, 0x2d, 0xf0, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0x9a, 0xc2, 0xc8, 0x25, 0x82, 0x26, 0xe3, 0x99, 0x97, 0x97, 0x5a, 0x24, 0x24, 0xc1,
	0xc5, 0x9e, 0x0e, 0x12, 0xf4, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x09, 0x82, 0x71, 0x85,
	0x44, 0xb8, 0x58, 0x13, 0x8b, 0xb3, 0x53, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x20,
	0x1c, 0x21, 0x31, 0x2e, 0xb6, 0xdc, 0xd4, 0x5c, 0xc7, 0xe4, 0x64, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x28, 0x4f, 0x48, 0x8a, 0x8b, 0xa3, 0x20, 0x31, 0x3d, 0x35, 0x38, 0xb3, 0x2a, 0x55,
	0x82, 0x45, 0x81, 0x51, 0x83, 0x37, 0x08, 0xce, 0x07, 0xd9, 0x01, 0x62, 0xfb, 0x95, 0xe6, 0x4a,
	0xb0, 0x82, 0xa5, 0x60, 0x5c, 0x27, 0x26, 0x0f, 0x66, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c,
	0xe4, 0x80, 0xf4, 0xb2, 0x00, 0x00, 0x00,
}
