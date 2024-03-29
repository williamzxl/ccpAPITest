// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GetGroupDetailResp.proto

package GetGroupDetailResp

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

type GetGroupDetailRespInner struct {
	Creator              *uint64  `protobuf:"varint,1,opt,name=creator" json:"creator,omitempty"`
	GroupName            *string  `protobuf:"bytes,2,opt,name=groupName" json:"groupName,omitempty"`
	Type                 *uint32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Province             *string  `protobuf:"bytes,4,opt,name=province" json:"province,omitempty"`
	City                 *string  `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	Scope                *uint32  `protobuf:"varint,6,opt,name=scope" json:"scope,omitempty"`
	Declared             *string  `protobuf:"bytes,7,opt,name=declared" json:"declared,omitempty"`
	DateCreated          *string  `protobuf:"bytes,8,opt,name=dateCreated" json:"dateCreated,omitempty"`
	Numbers              *uint32  `protobuf:"varint,9,opt,name=numbers" json:"numbers,omitempty"`
	IsNotice             *uint32  `protobuf:"varint,10,opt,name=isNotice" json:"isNotice,omitempty"`
	Permission           *uint32  `protobuf:"varint,11,opt,name=permission" json:"permission,omitempty"`
	GroupDomain          *string  `protobuf:"bytes,12,opt,name=groupDomain" json:"groupDomain,omitempty"`
	IsApplePush          *uint32  `protobuf:"varint,13,opt,name=isApplePush" json:"isApplePush,omitempty"`
	Target               *uint32  `protobuf:"varint,14,opt,name=target" json:"target,omitempty"`
	Anonymity            *uint32  `protobuf:"varint,15,opt,name=anonymity" json:"anonymity,omitempty"`
	IsForbid             *uint32  `protobuf:"varint,16,opt,name=isForbid" json:"isForbid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupDetailRespInner) Reset()         { *m = GetGroupDetailRespInner{} }
func (m *GetGroupDetailRespInner) String() string { return proto.CompactTextString(m) }
func (*GetGroupDetailRespInner) ProtoMessage()    {}
func (*GetGroupDetailRespInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec62d8af6f9a1979, []int{0}
}

func (m *GetGroupDetailRespInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupDetailRespInner.Unmarshal(m, b)
}
func (m *GetGroupDetailRespInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupDetailRespInner.Marshal(b, m, deterministic)
}
func (m *GetGroupDetailRespInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupDetailRespInner.Merge(m, src)
}
func (m *GetGroupDetailRespInner) XXX_Size() int {
	return xxx_messageInfo_GetGroupDetailRespInner.Size(m)
}
func (m *GetGroupDetailRespInner) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupDetailRespInner.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupDetailRespInner proto.InternalMessageInfo

func (m *GetGroupDetailRespInner) GetCreator() uint64 {
	if m != nil && m.Creator != nil {
		return *m.Creator
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *GetGroupDetailRespInner) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetProvince() string {
	if m != nil && m.Province != nil {
		return *m.Province
	}
	return ""
}

func (m *GetGroupDetailRespInner) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *GetGroupDetailRespInner) GetScope() uint32 {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetDeclared() string {
	if m != nil && m.Declared != nil {
		return *m.Declared
	}
	return ""
}

func (m *GetGroupDetailRespInner) GetDateCreated() string {
	if m != nil && m.DateCreated != nil {
		return *m.DateCreated
	}
	return ""
}

func (m *GetGroupDetailRespInner) GetNumbers() uint32 {
	if m != nil && m.Numbers != nil {
		return *m.Numbers
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetIsNotice() uint32 {
	if m != nil && m.IsNotice != nil {
		return *m.IsNotice
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetPermission() uint32 {
	if m != nil && m.Permission != nil {
		return *m.Permission
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetGroupDomain() string {
	if m != nil && m.GroupDomain != nil {
		return *m.GroupDomain
	}
	return ""
}

func (m *GetGroupDetailRespInner) GetIsApplePush() uint32 {
	if m != nil && m.IsApplePush != nil {
		return *m.IsApplePush
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetTarget() uint32 {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetAnonymity() uint32 {
	if m != nil && m.Anonymity != nil {
		return *m.Anonymity
	}
	return 0
}

func (m *GetGroupDetailRespInner) GetIsForbid() uint32 {
	if m != nil && m.IsForbid != nil {
		return *m.IsForbid
	}
	return 0
}

func init() {
	proto.RegisterType((*GetGroupDetailRespInner)(nil), "GetGroupDetailRespInner")
}

func init() { proto.RegisterFile("GetGroupDetailResp.proto", fileDescriptor_ec62d8af6f9a1979) }

var fileDescriptor_ec62d8af6f9a1979 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xdd, 0x4a, 0x3b, 0x31,
	0x10, 0xc5, 0xd9, 0x7e, 0x77, 0xfa, 0xef, 0x5f, 0x09, 0xa2, 0x83, 0x88, 0x2c, 0x5e, 0xf5, 0xca,
	0x77, 0x50, 0x8b, 0xd5, 0x9b, 0x22, 0xfb, 0x06, 0x69, 0x76, 0xa8, 0x81, 0xdd, 0x24, 0x24, 0x59,
	0x61, 0x1f, 0xd4, 0xf7, 0x91, 0xcc, 0xba, 0xed, 0x82, 0x77, 0xf9, 0x9d, 0x93, 0x39, 0xf3, 0x01,
	0xb8, 0xa3, 0xb8, 0xf3, 0xb6, 0x71, 0x5b, 0x8a, 0x52, 0x57, 0x05, 0x05, 0xf7, 0xe8, 0xbc, 0x8d,
	0xf6, 0xe1, 0x7b, 0x0c, 0x37, 0x7f, 0xcd, 0x77, 0x63, 0xc8, 0x0b, 0x84, 0xb9, 0xf2, 0x24, 0xa3,
	0xf5, 0x98, 0xe5, 0xd9, 0x66, 0x52, 0xf4, 0x28, 0xee, 0x60, 0x79, 0x4c, 0x15, 0x7b, 0x59, 0x13,
	0x8e, 0xf2, 0x6c, 0xb3, 0x2c, 0xce, 0x82, 0x10, 0x30, 0x89, 0xad, 0x23, 0x1c, 0xe7, 0xd9, 0x66,
	0x5d, 0xf0, 0x5b, 0xdc, 0xc2, 0xc2, 0x79, 0xfb, 0xa5, 0x8d, 0x22, 0x9c, 0x70, 0xc1, 0x89, 0xd3,
	0x7f, 0xa5, 0x63, 0x8b, 0x53, 0xd6, 0xf9, 0x2d, 0xae, 0x60, 0x1a, 0x94, 0x75, 0x84, 0x33, 0x0e,
	0xe9, 0x20, 0xa5, 0x94, 0xa4, 0x2a, 0xe9, 0xa9, 0xc4, 0x79, 0x97, 0xd2, 0xb3, 0xc8, 0x61, 0x55,
	0xca, 0x48, 0x2f, 0x69, 0x44, 0x2a, 0x71, 0xc1, 0xf6, 0x50, 0x4a, 0xfb, 0x98, 0xa6, 0x3e, 0x90,
	0x0f, 0xb8, 0xe4, 0xd4, 0x1e, 0x53, 0xae, 0x0e, 0x7b, 0x1b, 0xb5, 0x22, 0x04, 0xb6, 0x4e, 0x2c,
	0xee, 0x01, 0x1c, 0xf9, 0x5a, 0x87, 0xa0, 0xad, 0xc1, 0x15, 0xbb, 0x03, 0x25, 0xf5, 0xe5, 0xd5,
	0xb7, 0xb6, 0x96, 0xda, 0xe0, 0xbf, 0xae, 0xef, 0x40, 0x4a, 0x3f, 0x74, 0x78, 0x72, 0xae, 0xa2,
	0x8f, 0x26, 0x7c, 0xe2, 0x9a, 0x23, 0x86, 0x92, 0xb8, 0x86, 0x59, 0x94, 0xfe, 0x48, 0x11, 0xff,
	0xb3, 0xf9, 0x4b, 0xe9, 0xce, 0xd2, 0x58, 0xd3, 0xd6, 0xe9, 0x3c, 0x17, 0x6c, 0x9d, 0x85, 0x6e,
	0xea, 0x57, 0xeb, 0x0f, 0xba, 0xc4, 0xcb, 0x7e, 0xea, 0x8e, 0x9f, 0x47, 0x6f, 0xe3, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0xcf, 0xa3, 0x73, 0xf6, 0x01, 0x00, 0x00,
}
