// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SearchGroupsResp.proto

package SearchGroupsResp

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

type SearchGroupsRespInner struct {
	Groups               []*GroupSearchInfo `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
	PageNo               *uint32            `protobuf:"varint,2,opt,name=pageNo" json:"pageNo,omitempty"`
	PageSize             *uint32            `protobuf:"varint,3,opt,name=pageSize" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SearchGroupsRespInner) Reset()         { *m = SearchGroupsRespInner{} }
func (m *SearchGroupsRespInner) String() string { return proto.CompactTextString(m) }
func (*SearchGroupsRespInner) ProtoMessage()    {}
func (*SearchGroupsRespInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7fc63e0e6498453, []int{0}
}

func (m *SearchGroupsRespInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchGroupsRespInner.Unmarshal(m, b)
}
func (m *SearchGroupsRespInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchGroupsRespInner.Marshal(b, m, deterministic)
}
func (m *SearchGroupsRespInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchGroupsRespInner.Merge(m, src)
}
func (m *SearchGroupsRespInner) XXX_Size() int {
	return xxx_messageInfo_SearchGroupsRespInner.Size(m)
}
func (m *SearchGroupsRespInner) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchGroupsRespInner.DiscardUnknown(m)
}

var xxx_messageInfo_SearchGroupsRespInner proto.InternalMessageInfo

func (m *SearchGroupsRespInner) GetGroups() []*GroupSearchInfo {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *SearchGroupsRespInner) GetPageNo() uint32 {
	if m != nil && m.PageNo != nil {
		return *m.PageNo
	}
	return 0
}

func (m *SearchGroupsRespInner) GetPageSize() uint32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

type GroupSearchInfo struct {
	GroupId              *uint64  `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Owner                *string  `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	Permission           *uint32  `protobuf:"varint,4,opt,name=permission" json:"permission,omitempty"`
	Declared             *string  `protobuf:"bytes,5,opt,name=declared" json:"declared,omitempty"`
	MemberCount          *uint32  `protobuf:"varint,6,opt,name=memberCount" json:"memberCount,omitempty"`
	Scope                *uint32  `protobuf:"varint,7,opt,name=scope" json:"scope,omitempty"`
	GroupDomain          *string  `protobuf:"bytes,8,opt,name=groupDomain" json:"groupDomain,omitempty"`
	IsForbid             *uint32  `protobuf:"varint,9,opt,name=isForbid" json:"isForbid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupSearchInfo) Reset()         { *m = GroupSearchInfo{} }
func (m *GroupSearchInfo) String() string { return proto.CompactTextString(m) }
func (*GroupSearchInfo) ProtoMessage()    {}
func (*GroupSearchInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7fc63e0e6498453, []int{1}
}

func (m *GroupSearchInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupSearchInfo.Unmarshal(m, b)
}
func (m *GroupSearchInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupSearchInfo.Marshal(b, m, deterministic)
}
func (m *GroupSearchInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupSearchInfo.Merge(m, src)
}
func (m *GroupSearchInfo) XXX_Size() int {
	return xxx_messageInfo_GroupSearchInfo.Size(m)
}
func (m *GroupSearchInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupSearchInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupSearchInfo proto.InternalMessageInfo

func (m *GroupSearchInfo) GetGroupId() uint64 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *GroupSearchInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GroupSearchInfo) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *GroupSearchInfo) GetPermission() uint32 {
	if m != nil && m.Permission != nil {
		return *m.Permission
	}
	return 0
}

func (m *GroupSearchInfo) GetDeclared() string {
	if m != nil && m.Declared != nil {
		return *m.Declared
	}
	return ""
}

func (m *GroupSearchInfo) GetMemberCount() uint32 {
	if m != nil && m.MemberCount != nil {
		return *m.MemberCount
	}
	return 0
}

func (m *GroupSearchInfo) GetScope() uint32 {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return 0
}

func (m *GroupSearchInfo) GetGroupDomain() string {
	if m != nil && m.GroupDomain != nil {
		return *m.GroupDomain
	}
	return ""
}

func (m *GroupSearchInfo) GetIsForbid() uint32 {
	if m != nil && m.IsForbid != nil {
		return *m.IsForbid
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchGroupsRespInner)(nil), "SearchGroupsRespInner")
	proto.RegisterType((*GroupSearchInfo)(nil), "GroupSearchInfo")
}

func init() { proto.RegisterFile("SearchGroupsResp.proto", fileDescriptor_d7fc63e0e6498453) }

var fileDescriptor_d7fc63e0e6498453 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcb, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0x71, 0xee, 0x99, 0xf0, 0xf3, 0x17, 0xd1, 0x06, 0xd1, 0x45, 0x31, 0x59, 0x14, 0xaf,
	0xbc, 0xe8, 0x23, 0xb4, 0xa5, 0x6d, 0x36, 0x5d, 0x28, 0x4f, 0xe0, 0xc8, 0x13, 0x57, 0x10, 0x69,
	0x84, 0x14, 0xd1, 0xcb, 0x0b, 0xf4, 0xb5, 0x8b, 0xc7, 0x49, 0x30, 0xd9, 0xcd, 0x77, 0xce, 0xcc,
	0x99, 0x41, 0x82, 0xe5, 0x06, 0xab, 0xa0, 0x3f, 0x5e, 0x03, 0x25, 0x1f, 0x15, 0x46, 0x5f, 0xfa,
	0x40, 0x07, 0x5a, 0x25, 0xb8, 0xb9, 0x74, 0xd6, 0xce, 0x61, 0x10, 0x05, 0x4c, 0x1a, 0x96, 0x64,
	0x96, 0x0f, 0x8b, 0xc5, 0xc3, 0x55, 0xc9, 0x1d, 0x5d, 0xf3, 0xda, 0xed, 0x48, 0x1d, 0x7d, 0xb1,
	0x84, 0x89, 0xaf, 0x1a, 0x7c, 0x27, 0x39, 0xc8, 0xb3, 0xe2, 0x9f, 0x3a, 0x92, 0xb8, 0x85, 0x59,
	0x5b, 0x6d, 0xcc, 0x0f, 0xca, 0x21, 0x3b, 0x67, 0x5e, 0xfd, 0x0e, 0xe0, 0xff, 0x45, 0x9e, 0x90,
	0x30, 0xe5, 0xc4, 0x75, 0x2d, 0xb3, 0x3c, 0x2b, 0x46, 0xea, 0x84, 0x42, 0xc0, 0xc8, 0x55, 0x16,
	0x39, 0x7f, 0xae, 0xb8, 0x16, 0xd7, 0x30, 0xa6, 0x4f, 0x87, 0x81, 0xa3, 0xe7, 0xaa, 0x03, 0x71,
	0x07, 0xe0, 0x31, 0x58, 0x13, 0xa3, 0x21, 0x27, 0x47, 0xbc, 0xb5, 0xa7, 0xb4, 0x37, 0xd5, 0xa8,
	0xf7, 0x55, 0xc0, 0x5a, 0x8e, 0x79, 0xf0, 0xcc, 0x22, 0x87, 0x85, 0x45, 0xbb, 0xc5, 0xf0, 0x44,
	0xc9, 0x1d, 0xe4, 0x84, 0x87, 0xfb, 0x52, 0xbb, 0x33, 0x6a, 0xf2, 0x28, 0xa7, 0xec, 0x75, 0xd0,
	0xce, 0xf1, 0xa1, 0xcf, 0x64, 0x2b, 0xe3, 0xe4, 0x8c, 0x63, 0xfb, 0x52, 0xbb, 0xd5, 0xc4, 0x17,
	0x0a, 0x5b, 0x53, 0xcb, 0x79, 0xf7, 0x12, 0x27, 0x7e, 0xbc, 0x87, 0x95, 0x26, 0x5b, 0x7e, 0x27,
	0x77, 0x20, 0xd7, 0x7c, 0x25, 0x57, 0xee, 0xa9, 0x31, 0xba, 0xd4, 0x14, 0xb0, 0xfb, 0xa3, 0x6d,
	0xda, 0xbd, 0x0d, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xa2, 0x8b, 0x67, 0xc1, 0x01, 0x00,
	0x00,
}