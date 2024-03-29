// Code generated by protoc-gen-go. DO NOT EDIT.
// source: QueryGroupMemberCard.proto

package QueryGroupMemberCard

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

type QueryGroupMemberCardInner struct {
	MemberId             *uint64  `protobuf:"varint,1,opt,name=memberId" json:"memberId,omitempty"`
	Belong               *uint64  `protobuf:"varint,2,req,name=belong" json:"belong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryGroupMemberCardInner) Reset()         { *m = QueryGroupMemberCardInner{} }
func (m *QueryGroupMemberCardInner) String() string { return proto.CompactTextString(m) }
func (*QueryGroupMemberCardInner) ProtoMessage()    {}
func (*QueryGroupMemberCardInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3e742dda6839157, []int{0}
}

func (m *QueryGroupMemberCardInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryGroupMemberCardInner.Unmarshal(m, b)
}
func (m *QueryGroupMemberCardInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryGroupMemberCardInner.Marshal(b, m, deterministic)
}
func (m *QueryGroupMemberCardInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGroupMemberCardInner.Merge(m, src)
}
func (m *QueryGroupMemberCardInner) XXX_Size() int {
	return xxx_messageInfo_QueryGroupMemberCardInner.Size(m)
}
func (m *QueryGroupMemberCardInner) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGroupMemberCardInner.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGroupMemberCardInner proto.InternalMessageInfo

func (m *QueryGroupMemberCardInner) GetMemberId() uint64 {
	if m != nil && m.MemberId != nil {
		return *m.MemberId
	}
	return 0
}

func (m *QueryGroupMemberCardInner) GetBelong() uint64 {
	if m != nil && m.Belong != nil {
		return *m.Belong
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryGroupMemberCardInner)(nil), "QueryGroupMemberCardInner")
}

func init() { proto.RegisterFile("QueryGroupMemberCard.proto", fileDescriptor_b3e742dda6839157) }

var fileDescriptor_b3e742dda6839157 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x0a, 0x2c, 0x4d, 0x2d,
	0xaa, 0x74, 0x2f, 0xca, 0x2f, 0x2d, 0xf0, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d, 0x72, 0x4e, 0x2c, 0x4a,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2, 0xe7, 0x92, 0xc4, 0x26, 0xeb, 0x99, 0x97, 0x97,
	0x5a, 0x24, 0x24, 0xc5, 0xc5, 0x91, 0x0b, 0x16, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x09, 0x82, 0xf3, 0x85, 0xc4, 0xb8, 0xd8, 0x92, 0x52, 0x73, 0xf2, 0xf3, 0xd2, 0x25, 0x98, 0x14,
	0x98, 0x34, 0x58, 0x82, 0xa0, 0x3c, 0x27, 0x26, 0x0f, 0x66, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x76, 0xc7, 0x56, 0x09, 0x71, 0x00, 0x00, 0x00,
}
