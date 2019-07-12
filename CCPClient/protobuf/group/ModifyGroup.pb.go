// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ModifyGroup.proto

package ModifyGroup

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

type ModifyGroupInner struct {
	Modifier             *uint64  `protobuf:"varint,1,opt,name=modifier" json:"modifier,omitempty"`
	GroupId              *uint64  `protobuf:"varint,2,req,name=groupId" json:"groupId,omitempty"`
	GroupName            *string  `protobuf:"bytes,3,opt,name=groupName" json:"groupName,omitempty"`
	Type                 *uint32  `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	Province             *string  `protobuf:"bytes,5,opt,name=province" json:"province,omitempty"`
	City                 *string  `protobuf:"bytes,6,opt,name=city" json:"city,omitempty"`
	Scope                *uint32  `protobuf:"varint,7,opt,name=scope" json:"scope,omitempty"`
	Declared             *string  `protobuf:"bytes,8,opt,name=declared" json:"declared,omitempty"`
	Permission           *uint32  `protobuf:"varint,9,opt,name=permission" json:"permission,omitempty"`
	GroupDomain          *string  `protobuf:"bytes,10,opt,name=groupDomain" json:"groupDomain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyGroupInner) Reset()         { *m = ModifyGroupInner{} }
func (m *ModifyGroupInner) String() string { return proto.CompactTextString(m) }
func (*ModifyGroupInner) ProtoMessage()    {}
func (*ModifyGroupInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_92559a25a3e27d40, []int{0}
}

func (m *ModifyGroupInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyGroupInner.Unmarshal(m, b)
}
func (m *ModifyGroupInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyGroupInner.Marshal(b, m, deterministic)
}
func (m *ModifyGroupInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyGroupInner.Merge(m, src)
}
func (m *ModifyGroupInner) XXX_Size() int {
	return xxx_messageInfo_ModifyGroupInner.Size(m)
}
func (m *ModifyGroupInner) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyGroupInner.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyGroupInner proto.InternalMessageInfo

func (m *ModifyGroupInner) GetModifier() uint64 {
	if m != nil && m.Modifier != nil {
		return *m.Modifier
	}
	return 0
}

func (m *ModifyGroupInner) GetGroupId() uint64 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *ModifyGroupInner) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *ModifyGroupInner) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ModifyGroupInner) GetProvince() string {
	if m != nil && m.Province != nil {
		return *m.Province
	}
	return ""
}

func (m *ModifyGroupInner) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *ModifyGroupInner) GetScope() uint32 {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return 0
}

func (m *ModifyGroupInner) GetDeclared() string {
	if m != nil && m.Declared != nil {
		return *m.Declared
	}
	return ""
}

func (m *ModifyGroupInner) GetPermission() uint32 {
	if m != nil && m.Permission != nil {
		return *m.Permission
	}
	return 0
}

func (m *ModifyGroupInner) GetGroupDomain() string {
	if m != nil && m.GroupDomain != nil {
		return *m.GroupDomain
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyGroupInner)(nil), "ModifyGroupInner")
}

func init() { proto.RegisterFile("ModifyGroup.proto", fileDescriptor_92559a25a3e27d40) }

var fileDescriptor_92559a25a3e27d40 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0xb7, 0xeb, 0x6e, 0x47, 0x04, 0x1d, 0x3c, 0x0c, 0x22, 0x12, 0x3c, 0xf5, 0xe4,
	0x43, 0x88, 0xa0, 0x1e, 0xf4, 0x90, 0x37, 0x28, 0xc9, 0x28, 0x01, 0x93, 0x09, 0x69, 0x15, 0xfa,
	0x32, 0x3e, 0xab, 0x64, 0x96, 0x5d, 0x7b, 0xfb, 0xbf, 0x8f, 0xf9, 0xff, 0x04, 0xae, 0xde, 0xc4,
	0x87, 0x8f, 0xe5, 0xb9, 0xc8, 0x77, 0x7e, 0xc8, 0x45, 0x66, 0xb9, 0xff, 0x6d, 0xe1, 0x72, 0x65,
	0x5f, 0x53, 0xe2, 0x82, 0x37, 0xb0, 0x8f, 0xd5, 0x05, 0x2e, 0xd4, 0x98, 0x66, 0xe8, 0xec, 0x89,
	0x91, 0x60, 0xf7, 0xa9, 0x97, 0x9e, 0x5a, 0xd3, 0x0e, 0x9d, 0x3d, 0x22, 0xde, 0x42, 0xaf, 0xf1,
	0x7d, 0x8c, 0x4c, 0x1b, 0xd3, 0x0c, 0xbd, 0xfd, 0x17, 0x88, 0xd0, 0xcd, 0x4b, 0x66, 0xea, 0x4c,
	0x33, 0x5c, 0x58, 0xcd, 0xf5, 0x9d, 0x5c, 0xe4, 0x27, 0x24, 0xc7, 0xb4, 0xd5, 0xc2, 0x89, 0xeb,
	0xbd, 0x0b, 0xf3, 0x42, 0x67, 0xea, 0x35, 0xe3, 0x35, 0x6c, 0x27, 0x27, 0x99, 0x69, 0xa7, 0x23,
	0x07, 0xa8, 0x2b, 0x9e, 0xdd, 0xd7, 0x58, 0xd8, 0xd3, 0xfe, 0xb0, 0x72, 0x64, 0xbc, 0x03, 0xc8,
	0x5c, 0x62, 0x98, 0xa6, 0x20, 0x89, 0x7a, 0xad, 0xad, 0x0c, 0x1a, 0x38, 0xd7, 0x2f, 0x3e, 0x49,
	0x1c, 0x43, 0x22, 0xd0, 0xfa, 0x5a, 0x3d, 0xb6, 0x2f, 0x9b, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0x22, 0x6d, 0x36, 0x38, 0x01, 0x00, 0x00,
}
