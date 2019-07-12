// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UserAuthResp.proto

package auth

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type UserAuthRespInner struct {
	AuthState            *uint32                `protobuf:"varint,1,opt,name=authState" json:"authState,omitempty"`
	KickoffText          *string                `protobuf:"bytes,2,opt,name=kickoffText" json:"kickoffText,omitempty"`
	ConnectorId          *string                `protobuf:"bytes,3,opt,name=connectorId" json:"connectorId,omitempty"`
	Version              *uint64                `protobuf:"varint,4,opt,name=version" json:"version,omitempty"`
	Collect              *uint32                `protobuf:"varint,5,opt,name=collect" json:"collect,omitempty"`
	TransferPolicy       *TransferPolicy        `protobuf:"bytes,6,opt,name=transferPolicy" json:"transferPolicy,omitempty"`
	Pversion             *uint64                `protobuf:"varint,7,opt,name=pversion" json:"pversion,omitempty"`
	SoftVersion          *string                `protobuf:"bytes,8,opt,name=softVersion" json:"softVersion,omitempty"`
	UpdateMode           *uint32                `protobuf:"varint,9,opt,name=updateMode" json:"updateMode,omitempty"`
	Historyver           *uint64                `protobuf:"varint,10,opt,name=historyver" json:"historyver,omitempty"`
	AuthToken            *string                `protobuf:"bytes,11,opt,name=authToken" json:"authToken,omitempty"`
	IpSpeedTestPolicy    *IpSpeedTestPolicy     `protobuf:"bytes,12,opt,name=ipSpeedTestPolicy" json:"ipSpeedTestPolicy,omitempty"`
	LogUploadPolicy      *LogUploadPolicy       `protobuf:"bytes,13,opt,name=logUploadPolicy" json:"logUploadPolicy,omitempty"`
	AppleDeviceToken     *uint32                `protobuf:"varint,14,opt,name=appleDeviceToken" json:"appleDeviceToken,omitempty"`
	ConfigParams         *ConfigParams          `protobuf:"bytes,15,opt,name=configParams" json:"configParams,omitempty"`
	UpdateDesc           *string                `protobuf:"bytes,16,opt,name=updateDesc" json:"updateDesc,omitempty"`
	OfflineCall          *uint32                `protobuf:"varint,17,opt,name=offlineCall" json:"offlineCall,omitempty"`
	RedPacketToken       *string                `protobuf:"bytes,18,opt,name=redPacketToken" json:"redPacketToken,omitempty"`
	Timestamp            *uint64                `protobuf:"varint,19,opt,name=timestamp" json:"timestamp,omitempty"`
	Mediadatapolicy      *MediaDataUploadPolicy `protobuf:"bytes,20,opt,name=mediadatapolicy" json:"mediadatapolicy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UserAuthRespInner) Reset()         { *m = UserAuthRespInner{} }
func (m *UserAuthRespInner) String() string { return proto.CompactTextString(m) }
func (*UserAuthRespInner) ProtoMessage()    {}
func (*UserAuthRespInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa5c307a3ba97c9, []int{0}
}

func (m *UserAuthRespInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthRespInner.Unmarshal(m, b)
}
func (m *UserAuthRespInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthRespInner.Marshal(b, m, deterministic)
}
func (m *UserAuthRespInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthRespInner.Merge(m, src)
}
func (m *UserAuthRespInner) XXX_Size() int {
	return xxx_messageInfo_UserAuthRespInner.Size(m)
}
func (m *UserAuthRespInner) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthRespInner.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthRespInner proto.InternalMessageInfo

func (m *UserAuthRespInner) GetAuthState() uint32 {
	if m != nil && m.AuthState != nil {
		return *m.AuthState
	}
	return 0
}

func (m *UserAuthRespInner) GetKickoffText() string {
	if m != nil && m.KickoffText != nil {
		return *m.KickoffText
	}
	return ""
}

func (m *UserAuthRespInner) GetConnectorId() string {
	if m != nil && m.ConnectorId != nil {
		return *m.ConnectorId
	}
	return ""
}

func (m *UserAuthRespInner) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *UserAuthRespInner) GetCollect() uint32 {
	if m != nil && m.Collect != nil {
		return *m.Collect
	}
	return 0
}

func (m *UserAuthRespInner) GetTransferPolicy() *TransferPolicy {
	if m != nil {
		return m.TransferPolicy
	}
	return nil
}

func (m *UserAuthRespInner) GetPversion() uint64 {
	if m != nil && m.Pversion != nil {
		return *m.Pversion
	}
	return 0
}

func (m *UserAuthRespInner) GetSoftVersion() string {
	if m != nil && m.SoftVersion != nil {
		return *m.SoftVersion
	}
	return ""
}

func (m *UserAuthRespInner) GetUpdateMode() uint32 {
	if m != nil && m.UpdateMode != nil {
		return *m.UpdateMode
	}
	return 0
}

func (m *UserAuthRespInner) GetHistoryver() uint64 {
	if m != nil && m.Historyver != nil {
		return *m.Historyver
	}
	return 0
}

func (m *UserAuthRespInner) GetAuthToken() string {
	if m != nil && m.AuthToken != nil {
		return *m.AuthToken
	}
	return ""
}

func (m *UserAuthRespInner) GetIpSpeedTestPolicy() *IpSpeedTestPolicy {
	if m != nil {
		return m.IpSpeedTestPolicy
	}
	return nil
}

func (m *UserAuthRespInner) GetLogUploadPolicy() *LogUploadPolicy {
	if m != nil {
		return m.LogUploadPolicy
	}
	return nil
}

func (m *UserAuthRespInner) GetAppleDeviceToken() uint32 {
	if m != nil && m.AppleDeviceToken != nil {
		return *m.AppleDeviceToken
	}
	return 0
}

func (m *UserAuthRespInner) GetConfigParams() *ConfigParams {
	if m != nil {
		return m.ConfigParams
	}
	return nil
}

func (m *UserAuthRespInner) GetUpdateDesc() string {
	if m != nil && m.UpdateDesc != nil {
		return *m.UpdateDesc
	}
	return ""
}

func (m *UserAuthRespInner) GetOfflineCall() uint32 {
	if m != nil && m.OfflineCall != nil {
		return *m.OfflineCall
	}
	return 0
}

func (m *UserAuthRespInner) GetRedPacketToken() string {
	if m != nil && m.RedPacketToken != nil {
		return *m.RedPacketToken
	}
	return ""
}

func (m *UserAuthRespInner) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *UserAuthRespInner) GetMediadatapolicy() *MediaDataUploadPolicy {
	if m != nil {
		return m.Mediadatapolicy
	}
	return nil
}

type ServerAddr struct {
	Ip                   *string  `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port                 *uint32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Type                 *uint32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerAddr) Reset()         { *m = ServerAddr{} }
func (m *ServerAddr) String() string { return proto.CompactTextString(m) }
func (*ServerAddr) ProtoMessage()    {}
func (*ServerAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa5c307a3ba97c9, []int{1}
}

func (m *ServerAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerAddr.Unmarshal(m, b)
}
func (m *ServerAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerAddr.Marshal(b, m, deterministic)
}
func (m *ServerAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerAddr.Merge(m, src)
}
func (m *ServerAddr) XXX_Size() int {
	return xxx_messageInfo_ServerAddr.Size(m)
}
func (m *ServerAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerAddr.DiscardUnknown(m)
}

var xxx_messageInfo_ServerAddr proto.InternalMessageInfo

func (m *ServerAddr) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *ServerAddr) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ServerAddr) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type ConfigParams struct {
	Lots                 *uint32  `protobuf:"varint,1,opt,name=lots" json:"lots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigParams) Reset()         { *m = ConfigParams{} }
func (m *ConfigParams) String() string { return proto.CompactTextString(m) }
func (*ConfigParams) ProtoMessage()    {}
func (*ConfigParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa5c307a3ba97c9, []int{2}
}

func (m *ConfigParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigParams.Unmarshal(m, b)
}
func (m *ConfigParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigParams.Marshal(b, m, deterministic)
}
func (m *ConfigParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigParams.Merge(m, src)
}
func (m *ConfigParams) XXX_Size() int {
	return xxx_messageInfo_ConfigParams.Size(m)
}
func (m *ConfigParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigParams.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigParams proto.InternalMessageInfo

func (m *ConfigParams) GetLots() uint32 {
	if m != nil && m.Lots != nil {
		return *m.Lots
	}
	return 0
}

type TransferPolicy struct {
	Addr                 []*ServerAddr `protobuf:"bytes,1,rep,name=addr" json:"addr,omitempty"`
	Type                 *uint32       `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TransferPolicy) Reset()         { *m = TransferPolicy{} }
func (m *TransferPolicy) String() string { return proto.CompactTextString(m) }
func (*TransferPolicy) ProtoMessage()    {}
func (*TransferPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa5c307a3ba97c9, []int{3}
}

func (m *TransferPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferPolicy.Unmarshal(m, b)
}
func (m *TransferPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferPolicy.Marshal(b, m, deterministic)
}
func (m *TransferPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferPolicy.Merge(m, src)
}
func (m *TransferPolicy) XXX_Size() int {
	return xxx_messageInfo_TransferPolicy.Size(m)
}
func (m *TransferPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_TransferPolicy proto.InternalMessageInfo

func (m *TransferPolicy) GetAddr() []*ServerAddr {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *TransferPolicy) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type LogUploadPolicy struct {
	EnableLog            *uint32  `protobuf:"varint,1,opt,name=enableLog" json:"enableLog,omitempty"`
	Level                *uint32  `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	Policy               *uint32  `protobuf:"varint,3,opt,name=policy" json:"policy,omitempty"`
	Timeout              *uint32  `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogUploadPolicy) Reset()         { *m = LogUploadPolicy{} }
func (m *LogUploadPolicy) String() string { return proto.CompactTextString(m) }
func (*LogUploadPolicy) ProtoMessage()    {}
func (*LogUploadPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa5c307a3ba97c9, []int{4}
}

func (m *LogUploadPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogUploadPolicy.Unmarshal(m, b)
}
func (m *LogUploadPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogUploadPolicy.Marshal(b, m, deterministic)
}
func (m *LogUploadPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogUploadPolicy.Merge(m, src)
}
func (m *LogUploadPolicy) XXX_Size() int {
	return xxx_messageInfo_LogUploadPolicy.Size(m)
}
func (m *LogUploadPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_LogUploadPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_LogUploadPolicy proto.InternalMessageInfo

func (m *LogUploadPolicy) GetEnableLog() uint32 {
	if m != nil && m.EnableLog != nil {
		return *m.EnableLog
	}
	return 0
}

func (m *LogUploadPolicy) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *LogUploadPolicy) GetPolicy() uint32 {
	if m != nil && m.Policy != nil {
		return *m.Policy
	}
	return 0
}

func (m *LogUploadPolicy) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

type IpSpeedTestPolicy struct {
	PolicyType           *uint32       `protobuf:"varint,1,opt,name=policyType" json:"policyType,omitempty"`
	ServerAddr           []*ServerAddr `protobuf:"bytes,2,rep,name=serverAddr" json:"serverAddr,omitempty"`
	Count                *uint32       `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Interval             *uint32       `protobuf:"varint,4,opt,name=interval" json:"interval,omitempty"`
	Timeout              *uint32       `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
	UploadPolicy         *uint32       `protobuf:"varint,6,opt,name=uploadPolicy" json:"uploadPolicy,omitempty"`
	RemoteHost           *string       `protobuf:"bytes,7,opt,name=remoteHost" json:"remoteHost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IpSpeedTestPolicy) Reset()         { *m = IpSpeedTestPolicy{} }
func (m *IpSpeedTestPolicy) String() string { return proto.CompactTextString(m) }
func (*IpSpeedTestPolicy) ProtoMessage()    {}
func (*IpSpeedTestPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa5c307a3ba97c9, []int{5}
}

func (m *IpSpeedTestPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpSpeedTestPolicy.Unmarshal(m, b)
}
func (m *IpSpeedTestPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpSpeedTestPolicy.Marshal(b, m, deterministic)
}
func (m *IpSpeedTestPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpSpeedTestPolicy.Merge(m, src)
}
func (m *IpSpeedTestPolicy) XXX_Size() int {
	return xxx_messageInfo_IpSpeedTestPolicy.Size(m)
}
func (m *IpSpeedTestPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_IpSpeedTestPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_IpSpeedTestPolicy proto.InternalMessageInfo

func (m *IpSpeedTestPolicy) GetPolicyType() uint32 {
	if m != nil && m.PolicyType != nil {
		return *m.PolicyType
	}
	return 0
}

func (m *IpSpeedTestPolicy) GetServerAddr() []*ServerAddr {
	if m != nil {
		return m.ServerAddr
	}
	return nil
}

func (m *IpSpeedTestPolicy) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *IpSpeedTestPolicy) GetInterval() uint32 {
	if m != nil && m.Interval != nil {
		return *m.Interval
	}
	return 0
}

func (m *IpSpeedTestPolicy) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *IpSpeedTestPolicy) GetUploadPolicy() uint32 {
	if m != nil && m.UploadPolicy != nil {
		return *m.UploadPolicy
	}
	return 0
}

func (m *IpSpeedTestPolicy) GetRemoteHost() string {
	if m != nil && m.RemoteHost != nil {
		return *m.RemoteHost
	}
	return ""
}

type MediaDataUploadPolicy struct {
	SummaryRule          *uint32  `protobuf:"varint,1,opt,name=summaryRule" json:"summaryRule,omitempty"`
	SummaryTime          *uint32  `protobuf:"varint,2,opt,name=summaryTime" json:"summaryTime,omitempty"`
	DetailRule           *uint32  `protobuf:"varint,3,opt,name=detailRule" json:"detailRule,omitempty"`
	DetailTime           *uint32  `protobuf:"varint,4,opt,name=detailTime" json:"detailTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MediaDataUploadPolicy) Reset()         { *m = MediaDataUploadPolicy{} }
func (m *MediaDataUploadPolicy) String() string { return proto.CompactTextString(m) }
func (*MediaDataUploadPolicy) ProtoMessage()    {}
func (*MediaDataUploadPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa5c307a3ba97c9, []int{6}
}

func (m *MediaDataUploadPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaDataUploadPolicy.Unmarshal(m, b)
}
func (m *MediaDataUploadPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaDataUploadPolicy.Marshal(b, m, deterministic)
}
func (m *MediaDataUploadPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaDataUploadPolicy.Merge(m, src)
}
func (m *MediaDataUploadPolicy) XXX_Size() int {
	return xxx_messageInfo_MediaDataUploadPolicy.Size(m)
}
func (m *MediaDataUploadPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaDataUploadPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MediaDataUploadPolicy proto.InternalMessageInfo

func (m *MediaDataUploadPolicy) GetSummaryRule() uint32 {
	if m != nil && m.SummaryRule != nil {
		return *m.SummaryRule
	}
	return 0
}

func (m *MediaDataUploadPolicy) GetSummaryTime() uint32 {
	if m != nil && m.SummaryTime != nil {
		return *m.SummaryTime
	}
	return 0
}

func (m *MediaDataUploadPolicy) GetDetailRule() uint32 {
	if m != nil && m.DetailRule != nil {
		return *m.DetailRule
	}
	return 0
}

func (m *MediaDataUploadPolicy) GetDetailTime() uint32 {
	if m != nil && m.DetailTime != nil {
		return *m.DetailTime
	}
	return 0
}

func init() {
	proto.RegisterType((*UserAuthRespInner)(nil), "UserAuthRespInner")
	proto.RegisterType((*ServerAddr)(nil), "ServerAddr")
	proto.RegisterType((*ConfigParams)(nil), "ConfigParams")
	proto.RegisterType((*TransferPolicy)(nil), "TransferPolicy")
	proto.RegisterType((*LogUploadPolicy)(nil), "LogUploadPolicy")
	proto.RegisterType((*IpSpeedTestPolicy)(nil), "IpSpeedTestPolicy")
	proto.RegisterType((*MediaDataUploadPolicy)(nil), "MediaDataUploadPolicy")
}

func init() { proto.RegisterFile("UserAuthResp.proto", fileDescriptor_0aa5c307a3ba97c9) }

var fileDescriptor_0aa5c307a3ba97c9 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0xc7, 0x21, 0xc7, 0x71, 0xe2, 0xf1, 0x57, 0xcc, 0xa6, 0x01, 0x51, 0x14, 0xad, 0xa1, 0x43,
	0x61, 0xb4, 0x80, 0x81, 0xe6, 0x52, 0xa0, 0xa7, 0xa4, 0x71, 0x81, 0x04, 0x48, 0x80, 0x80, 0x71,
	0xf6, 0xce, 0x95, 0xc6, 0x09, 0x61, 0x4a, 0x24, 0x28, 0xca, 0xbb, 0x7e, 0x96, 0x3d, 0xef, 0xfb,
	0xed, 0x23, 0x2c, 0x48, 0x49, 0x36, 0x6d, 0xef, 0x4d, 0xf3, 0x9b, 0xe1, 0xcc, 0x7f, 0x86, 0x1c,
	0x01, 0x79, 0x2d, 0xd0, 0xdc, 0x96, 0xf6, 0x9d, 0x61, 0xa1, 0x67, 0xda, 0x28, 0xab, 0xe2, 0xaf,
	0x1d, 0x18, 0x87, 0xf8, 0x21, 0xcf, 0xd1, 0x90, 0x5f, 0xa1, 0xcb, 0x4b, 0xfb, 0xfe, 0x62, 0xb9,
	0x45, 0x1a, 0x4d, 0xa2, 0xe9, 0x80, 0xed, 0x00, 0x99, 0x40, 0x6f, 0x25, 0x92, 0x95, 0x5a, 0x2e,
	0x17, 0xf8, 0xd9, 0xd2, 0xd6, 0x24, 0x9a, 0x76, 0x59, 0x88, 0x5c, 0x44, 0xa2, 0xf2, 0x1c, 0x13,
	0xab, 0xcc, 0x43, 0x4a, 0x4f, 0xaa, 0x88, 0x00, 0x11, 0x0a, 0x67, 0x6b, 0x34, 0x85, 0x50, 0x39,
	0x6d, 0x4f, 0xa2, 0x69, 0x9b, 0x35, 0xa6, 0xf3, 0x24, 0x4a, 0x4a, 0x4c, 0x2c, 0x3d, 0xf5, 0x95,
	0x1b, 0x93, 0xfc, 0x03, 0x43, 0x6b, 0x78, 0x5e, 0x2c, 0xd1, 0x3c, 0x2b, 0x29, 0x92, 0x0d, 0xed,
	0x4c, 0xa2, 0x69, 0xef, 0x7a, 0x34, 0x5b, 0xec, 0x61, 0x76, 0x10, 0x46, 0x7e, 0x81, 0x73, 0xdd,
	0x54, 0x3b, 0xf3, 0xd5, 0xb6, 0xb6, 0x93, 0x5a, 0xa8, 0xa5, 0xfd, 0x50, 0xbb, 0xcf, 0x2b, 0xa9,
	0x01, 0x22, 0xbf, 0x01, 0x94, 0x3a, 0xe5, 0x16, 0x9f, 0x54, 0x8a, 0xb4, 0xeb, 0x35, 0x05, 0xc4,
	0xf9, 0xdf, 0x45, 0x61, 0x95, 0xd9, 0xac, 0xd1, 0x50, 0xf0, 0xf9, 0x03, 0xd2, 0x0c, 0x73, 0xa1,
	0x56, 0x98, 0xd3, 0x9e, 0xcf, 0xbf, 0x03, 0xe4, 0x06, 0xc6, 0x42, 0xbf, 0x68, 0xc4, 0x74, 0x81,
	0x85, 0xad, 0xfb, 0xea, 0xfb, 0xbe, 0xc8, 0xec, 0xe1, 0xd0, 0xc3, 0x8e, 0x83, 0xc9, 0xbf, 0x30,
	0x92, 0xea, 0xed, 0x55, 0x4b, 0xc5, 0xd3, 0xfa, 0xfc, 0xc0, 0x9f, 0xbf, 0x98, 0x3d, 0xee, 0x73,
	0x76, 0x18, 0x48, 0xfe, 0x84, 0x0b, 0xae, 0xb5, 0xc4, 0x39, 0xae, 0x45, 0x82, 0x95, 0xc4, 0xa1,
	0xef, 0xf0, 0x88, 0x93, 0xbf, 0xa1, 0x9f, 0xa8, 0x7c, 0x29, 0xde, 0x9e, 0xb9, 0xe1, 0x59, 0x41,
	0x47, 0xbe, 0xc8, 0x60, 0x76, 0x17, 0x40, 0xb6, 0x17, 0xb2, 0x1b, 0xdd, 0x1c, 0x8b, 0x84, 0x5e,
	0xf8, 0xde, 0x03, 0xe2, 0x86, 0xaf, 0x96, 0x4b, 0x29, 0x72, 0xbc, 0xe3, 0x52, 0xd2, 0xb1, 0xaf,
	0x1c, 0x22, 0xf2, 0x07, 0x0c, 0x0d, 0xa6, 0xcf, 0x3c, 0x59, 0xa1, 0xad, 0xe4, 0x11, 0x9f, 0xe5,
	0x80, 0xba, 0x21, 0x5b, 0x91, 0x61, 0x61, 0x79, 0xa6, 0xe9, 0x4f, 0xfe, 0x0e, 0x76, 0x80, 0xdc,
	0xc0, 0x28, 0xc3, 0x54, 0xf0, 0x94, 0x5b, 0xae, 0xab, 0x11, 0x5d, 0x7a, 0xf5, 0x57, 0xb3, 0x27,
	0xc7, 0xe7, 0xdc, 0xf2, 0xfd, 0x41, 0x1d, 0x84, 0xc7, 0x73, 0x80, 0x17, 0x34, 0x6b, 0x34, 0xb7,
	0x69, 0x6a, 0xc8, 0x10, 0x5a, 0x42, 0xfb, 0xc5, 0xe8, 0xb2, 0x96, 0xd0, 0x84, 0x40, 0x5b, 0x2b,
	0x53, 0xad, 0xc2, 0x80, 0xf9, 0x6f, 0xc7, 0xec, 0x46, 0xa3, 0x7f, 0xfc, 0x03, 0xe6, 0xbf, 0xe3,
	0x18, 0xfa, 0xe1, 0xb4, 0x5c, 0x8c, 0x54, 0xb6, 0xa8, 0x57, 0xcc, 0x7f, 0xc7, 0xff, 0xc3, 0x70,
	0xff, 0x39, 0x93, 0xdf, 0xa1, 0xcd, 0xd3, 0xd4, 0xd0, 0x68, 0x72, 0x32, 0xed, 0x5d, 0xf7, 0x66,
	0x3b, 0x21, 0xcc, 0x3b, 0xb6, 0xa5, 0x5a, 0x41, 0xa9, 0x4f, 0x30, 0x3a, 0xb8, 0x7d, 0x37, 0x23,
	0xcc, 0xf9, 0x47, 0x89, 0x8f, 0xea, 0xad, 0xd9, 0xea, 0x2d, 0x20, 0x97, 0x70, 0x2a, 0x71, 0x8d,
	0xb2, 0xce, 0x52, 0x19, 0xe4, 0x0a, 0x3a, 0xf5, 0xc0, 0xaa, 0x3e, 0x6a, 0xcb, 0x6d, 0xa9, 0x1b,
	0xaf, 0x2a, 0xad, 0xdf, 0xdf, 0x01, 0x6b, 0xcc, 0xf8, 0x5b, 0x04, 0xe3, 0xa3, 0x77, 0xeb, 0x5e,
	0x42, 0x75, 0x72, 0xe1, 0x84, 0x56, 0xc5, 0x03, 0x42, 0xfe, 0x02, 0x28, 0xb6, 0x6d, 0xd1, 0xd6,
	0x71, 0xa7, 0x81, 0xdb, 0x49, 0x4d, 0x54, 0x99, 0xdb, 0x5a, 0x53, 0x65, 0xb8, 0x2d, 0x17, 0xb9,
	0x45, 0xb3, 0xe6, 0xb2, 0xd6, 0xb4, 0xb5, 0x43, 0xb9, 0xa7, 0x7b, 0x72, 0x49, 0x0c, 0xfd, 0x32,
	0x5c, 0x9d, 0x8e, 0x77, 0xef, 0x31, 0x27, 0xde, 0x60, 0xa6, 0x2c, 0xde, 0xab, 0xc2, 0xfa, 0x3f,
	0x48, 0x97, 0x05, 0x24, 0xfe, 0x12, 0xc1, 0xcf, 0x3f, 0x7c, 0x47, 0xfe, 0xef, 0x52, 0x66, 0x19,
	0x37, 0x1b, 0x56, 0xca, 0xa6, 0xef, 0x10, 0x05, 0x11, 0x0b, 0x91, 0x35, 0x57, 0x18, 0x22, 0x57,
	0x3d, 0x45, 0xcb, 0x85, 0xf4, 0x29, 0xaa, 0x96, 0x03, 0xb2, 0xf3, 0xfb, 0x04, 0xed, 0xd0, 0xef,
	0xc8, 0x7f, 0xad, 0xfb, 0x93, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xed, 0x44, 0x63, 0xfb,
	0x05, 0x00, 0x00,
}
