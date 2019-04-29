// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gate.proto

package gate

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

type GateMsgType int32

const (
	GateMsgType_gs GateMsgType = 1
	GateMsgType_us GateMsgType = 2
)

var GateMsgType_name = map[int32]string{
	1: "gs",
	2: "us",
}

var GateMsgType_value = map[string]int32{
	"gs": 1,
	"us": 2,
}

func (x GateMsgType) Enum() *GateMsgType {
	p := new(GateMsgType)
	*p = x
	return p
}

func (x GateMsgType) String() string {
	return proto.EnumName(GateMsgType_name, int32(x))
}

func (x *GateMsgType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GateMsgType_value, data, "GateMsgType")
	if err != nil {
		return err
	}
	*x = GateMsgType(value)
	return nil
}

func (GateMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}

type Msg int32

const (
	// 注册服务
	Msg_registerServerReq  Msg = 1
	Msg_registerServerResp Msg = 2
	Msg_g2s                Msg = 3
	Msg_s2g                Msg = 4
	Msg_s2u                Msg = 5
	Msg_u2s                Msg = 6
)

var Msg_name = map[int32]string{
	1: "registerServerReq",
	2: "registerServerResp",
	3: "g2s",
	4: "s2g",
	5: "s2u",
	6: "u2s",
}

var Msg_value = map[string]int32{
	"registerServerReq":  1,
	"registerServerResp": 2,
	"g2s":                3,
	"s2g":                4,
	"s2u":                5,
	"u2s":                6,
}

func (x Msg) Enum() *Msg {
	p := new(Msg)
	*p = x
	return p
}

func (x Msg) String() string {
	return proto.EnumName(Msg_name, int32(x))
}

func (x *Msg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Msg_value, data, "Msg")
	if err != nil {
		return err
	}
	*x = Msg(value)
	return nil
}

func (Msg) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{1}
}

type ErrorCode int32

const (
	ErrorCode_success ErrorCode = 0
	ErrorCode_failed  ErrorCode = 1
)

var ErrorCode_name = map[int32]string{
	0: "success",
	1: "failed",
}

var ErrorCode_value = map[string]int32{
	"success": 0,
	"failed":  1,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{2}
}

type ServerType int32

const (
	ServerType_service ServerType = 0
	ServerType_game    ServerType = 1
)

var ServerType_name = map[int32]string{
	0: "service",
	1: "game",
}

var ServerType_value = map[string]int32{
	"service": 0,
	"game":    1,
}

func (x ServerType) Enum() *ServerType {
	p := new(ServerType)
	*p = x
	return p
}

func (x ServerType) String() string {
	return proto.EnumName(ServerType_name, int32(x))
}

func (x *ServerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServerType_value, data, "ServerType")
	if err != nil {
		return err
	}
	*x = ServerType(value)
	return nil
}

func (ServerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{3}
}

//message BackendMsgHeader {
//    optional uint32 cmd=1;
//}
//
//message BackendMsg {
//    optional BackendMsgHeader header=1;
//    optional bytes data=2;
//}
//
//message NetMsg {
//    optional uint32 msg_type=1;
//    optional uint32 msg_id=2;
//    optional bytes data=3;
//}
//
type RespTips struct {
	Code                 *ErrorCode `protobuf:"varint,1,opt,name=code,enum=gate.ErrorCode" json:"code,omitempty"`
	Msg                  *string    `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RespTips) Reset()         { *m = RespTips{} }
func (m *RespTips) String() string { return proto.CompactTextString(m) }
func (*RespTips) ProtoMessage()    {}
func (*RespTips) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}

func (m *RespTips) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespTips.Unmarshal(m, b)
}
func (m *RespTips) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespTips.Marshal(b, m, deterministic)
}
func (m *RespTips) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespTips.Merge(m, src)
}
func (m *RespTips) XXX_Size() int {
	return xxx_messageInfo_RespTips.Size(m)
}
func (m *RespTips) XXX_DiscardUnknown() {
	xxx_messageInfo_RespTips.DiscardUnknown(m)
}

var xxx_messageInfo_RespTips proto.InternalMessageInfo

func (m *RespTips) GetCode() ErrorCode {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ErrorCode_success
}

func (m *RespTips) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type RegisterServerReq struct {
	ServerType           *ServerType `protobuf:"varint,1,opt,name=server_type,json=serverType,enum=gate.ServerType" json:"server_type,omitempty"`
	MsgType              *uint32     `protobuf:"varint,2,opt,name=msg_type,json=msgType" json:"msg_type,omitempty"`
	ServerId             *uint32     `protobuf:"varint,3,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RegisterServerReq) Reset()         { *m = RegisterServerReq{} }
func (m *RegisterServerReq) String() string { return proto.CompactTextString(m) }
func (*RegisterServerReq) ProtoMessage()    {}
func (*RegisterServerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{1}
}

func (m *RegisterServerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterServerReq.Unmarshal(m, b)
}
func (m *RegisterServerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterServerReq.Marshal(b, m, deterministic)
}
func (m *RegisterServerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterServerReq.Merge(m, src)
}
func (m *RegisterServerReq) XXX_Size() int {
	return xxx_messageInfo_RegisterServerReq.Size(m)
}
func (m *RegisterServerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterServerReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterServerReq proto.InternalMessageInfo

func (m *RegisterServerReq) GetServerType() ServerType {
	if m != nil && m.ServerType != nil {
		return *m.ServerType
	}
	return ServerType_service
}

func (m *RegisterServerReq) GetMsgType() uint32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *RegisterServerReq) GetServerId() uint32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

type RegisterServerResp struct {
	ServerType           *ServerType `protobuf:"varint,1,opt,name=server_type,json=serverType,enum=gate.ServerType" json:"server_type,omitempty"`
	MsgType              *uint32     `protobuf:"varint,2,opt,name=msg_type,json=msgType" json:"msg_type,omitempty"`
	ServerId             *uint32     `protobuf:"varint,3,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ServerTime           *int64      `protobuf:"varint,4,opt,name=server_time,json=serverTime" json:"server_time,omitempty"`
	Ret                  *RespTips   `protobuf:"bytes,5,opt,name=ret" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RegisterServerResp) Reset()         { *m = RegisterServerResp{} }
func (m *RegisterServerResp) String() string { return proto.CompactTextString(m) }
func (*RegisterServerResp) ProtoMessage()    {}
func (*RegisterServerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{2}
}

func (m *RegisterServerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterServerResp.Unmarshal(m, b)
}
func (m *RegisterServerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterServerResp.Marshal(b, m, deterministic)
}
func (m *RegisterServerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterServerResp.Merge(m, src)
}
func (m *RegisterServerResp) XXX_Size() int {
	return xxx_messageInfo_RegisterServerResp.Size(m)
}
func (m *RegisterServerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterServerResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterServerResp proto.InternalMessageInfo

func (m *RegisterServerResp) GetServerType() ServerType {
	if m != nil && m.ServerType != nil {
		return *m.ServerType
	}
	return ServerType_service
}

func (m *RegisterServerResp) GetMsgType() uint32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *RegisterServerResp) GetServerId() uint32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *RegisterServerResp) GetServerTime() int64 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

func (m *RegisterServerResp) GetRet() *RespTips {
	if m != nil {
		return m.Ret
	}
	return nil
}

type SGMsg struct {
	Gmt                  *GateMsgType `protobuf:"varint,1,opt,name=gmt,enum=gate.GateMsgType" json:"gmt,omitempty"`
	MsgType              *uint32      `protobuf:"varint,2,opt,name=msg_type,json=msgType" json:"msg_type,omitempty"`
	MsgId                *uint32      `protobuf:"varint,3,opt,name=msg_id,json=msgId" json:"msg_id,omitempty"`
	Data                 []byte       `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SGMsg) Reset()         { *m = SGMsg{} }
func (m *SGMsg) String() string { return proto.CompactTextString(m) }
func (*SGMsg) ProtoMessage()    {}
func (*SGMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{3}
}

func (m *SGMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SGMsg.Unmarshal(m, b)
}
func (m *SGMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SGMsg.Marshal(b, m, deterministic)
}
func (m *SGMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SGMsg.Merge(m, src)
}
func (m *SGMsg) XXX_Size() int {
	return xxx_messageInfo_SGMsg.Size(m)
}
func (m *SGMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SGMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SGMsg proto.InternalMessageInfo

func (m *SGMsg) GetGmt() GateMsgType {
	if m != nil && m.Gmt != nil {
		return *m.Gmt
	}
	return GateMsgType_gs
}

func (m *SGMsg) GetMsgType() uint32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *SGMsg) GetMsgId() uint32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *SGMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GSMsg struct {
	Gmt                  *GateMsgType `protobuf:"varint,1,opt,name=gmt,enum=gate.GateMsgType" json:"gmt,omitempty"`
	FromMsgType          *uint32      `protobuf:"varint,2,opt,name=from_msg_type,json=fromMsgType" json:"from_msg_type,omitempty"`
	ToMsgType            *uint32      `protobuf:"varint,3,opt,name=to_msg_type,json=toMsgType" json:"to_msg_type,omitempty"`
	MsgId                *uint32      `protobuf:"varint,4,opt,name=msg_id,json=msgId" json:"msg_id,omitempty"`
	Data                 []byte       `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GSMsg) Reset()         { *m = GSMsg{} }
func (m *GSMsg) String() string { return proto.CompactTextString(m) }
func (*GSMsg) ProtoMessage()    {}
func (*GSMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{4}
}

func (m *GSMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GSMsg.Unmarshal(m, b)
}
func (m *GSMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GSMsg.Marshal(b, m, deterministic)
}
func (m *GSMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GSMsg.Merge(m, src)
}
func (m *GSMsg) XXX_Size() int {
	return xxx_messageInfo_GSMsg.Size(m)
}
func (m *GSMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GSMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GSMsg proto.InternalMessageInfo

func (m *GSMsg) GetGmt() GateMsgType {
	if m != nil && m.Gmt != nil {
		return *m.Gmt
	}
	return GateMsgType_gs
}

func (m *GSMsg) GetFromMsgType() uint32 {
	if m != nil && m.FromMsgType != nil {
		return *m.FromMsgType
	}
	return 0
}

func (m *GSMsg) GetToMsgType() uint32 {
	if m != nil && m.ToMsgType != nil {
		return *m.ToMsgType
	}
	return 0
}

func (m *GSMsg) GetMsgId() uint32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *GSMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type S2GUserMsg struct {
	ConnIds              []uint64 `protobuf:"varint,1,rep,name=conn_ids,json=connIds" json:"conn_ids,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2GUserMsg) Reset()         { *m = S2GUserMsg{} }
func (m *S2GUserMsg) String() string { return proto.CompactTextString(m) }
func (*S2GUserMsg) ProtoMessage()    {}
func (*S2GUserMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{5}
}

func (m *S2GUserMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2GUserMsg.Unmarshal(m, b)
}
func (m *S2GUserMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2GUserMsg.Marshal(b, m, deterministic)
}
func (m *S2GUserMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2GUserMsg.Merge(m, src)
}
func (m *S2GUserMsg) XXX_Size() int {
	return xxx_messageInfo_S2GUserMsg.Size(m)
}
func (m *S2GUserMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_S2GUserMsg.DiscardUnknown(m)
}

var xxx_messageInfo_S2GUserMsg proto.InternalMessageInfo

func (m *S2GUserMsg) GetConnIds() []uint64 {
	if m != nil {
		return m.ConnIds
	}
	return nil
}

func (m *S2GUserMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GUMsg struct {
	MsgType              *uint32  `protobuf:"varint,1,opt,name=msg_type,json=msgType" json:"msg_type,omitempty"`
	MsgId                *uint32  `protobuf:"varint,2,opt,name=msg_id,json=msgId" json:"msg_id,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GUMsg) Reset()         { *m = GUMsg{} }
func (m *GUMsg) String() string { return proto.CompactTextString(m) }
func (*GUMsg) ProtoMessage()    {}
func (*GUMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{6}
}

func (m *GUMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GUMsg.Unmarshal(m, b)
}
func (m *GUMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GUMsg.Marshal(b, m, deterministic)
}
func (m *GUMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GUMsg.Merge(m, src)
}
func (m *GUMsg) XXX_Size() int {
	return xxx_messageInfo_GUMsg.Size(m)
}
func (m *GUMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GUMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GUMsg proto.InternalMessageInfo

func (m *GUMsg) GetMsgType() uint32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *GUMsg) GetMsgId() uint32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *GUMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("gate.GateMsgType", GateMsgType_name, GateMsgType_value)
	proto.RegisterEnum("gate.Msg", Msg_name, Msg_value)
	proto.RegisterEnum("gate.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("gate.ServerType", ServerType_name, ServerType_value)
	proto.RegisterType((*RespTips)(nil), "gate.RespTips")
	proto.RegisterType((*RegisterServerReq)(nil), "gate.RegisterServerReq")
	proto.RegisterType((*RegisterServerResp)(nil), "gate.RegisterServerResp")
	proto.RegisterType((*SGMsg)(nil), "gate.SGMsg")
	proto.RegisterType((*GSMsg)(nil), "gate.GSMsg")
	proto.RegisterType((*S2GUserMsg)(nil), "gate.S2GUserMsg")
	proto.RegisterType((*GUMsg)(nil), "gate.GUMsg")
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor_743bb58a714d8b7d) }

var fileDescriptor_743bb58a714d8b7d = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0x2b, 0xcb, 0xce, 0x9f, 0xd7, 0x6b, 0xa7, 0x0a, 0x3a, 0x5c, 0xc6, 0x36, 0xe3, 0xec,
	0x10, 0x72, 0x28, 0xcc, 0xd7, 0x9d, 0xc6, 0x18, 0x26, 0x87, 0x5c, 0x94, 0x16, 0x76, 0x0b, 0xc6,
	0x7e, 0x2b, 0x0c, 0x73, 0xed, 0xe9, 0x55, 0x0a, 0xbd, 0xec, 0x8b, 0xec, 0xdb, 0xec, 0x93, 0x0d,
	0xd9, 0xa9, 0x13, 0xd2, 0x6e, 0xec, 0xb4, 0x93, 0x1f, 0xeb, 0x79, 0xf4, 0xbc, 0x3f, 0x84, 0x04,
	0xa0, 0x73, 0x8b, 0x57, 0xad, 0x69, 0x6c, 0x23, 0x7d, 0xa7, 0x93, 0x4f, 0x30, 0x51, 0x48, 0xed,
	0x75, 0xd5, 0x92, 0x9c, 0x81, 0x5f, 0x34, 0x25, 0x46, 0x2c, 0x66, 0xf3, 0xb3, 0xf4, 0xe5, 0x55,
	0x17, 0xfe, 0x62, 0x4c, 0x63, 0x3e, 0x37, 0x25, 0xaa, 0xce, 0x94, 0x02, 0x78, 0x4d, 0x3a, 0xf2,
	0x62, 0x36, 0x9f, 0x2a, 0x27, 0x93, 0x1f, 0x70, 0xae, 0x50, 0x57, 0x64, 0xd1, 0xac, 0xd1, 0xdc,
	0xa3, 0x51, 0xf8, 0x5d, 0x7e, 0x80, 0x90, 0xba, 0x9f, 0x8d, 0x7d, 0x68, 0x1f, 0x2b, 0x45, 0x5f,
	0xd9, 0xa7, 0xae, 0x1f, 0x5a, 0x54, 0x40, 0x83, 0x96, 0x97, 0x30, 0xa9, 0x49, 0xf7, 0x79, 0x57,
	0x7f, 0xaa, 0xc6, 0x35, 0xe9, 0xce, 0x7a, 0x0d, 0xd3, 0x5d, 0x5b, 0x55, 0x46, 0xbc, 0xf3, 0x26,
	0xfd, 0xc2, 0xb2, 0x4c, 0x7e, 0x31, 0x90, 0xc7, 0x00, 0xd4, 0xfe, 0x47, 0x02, 0xf9, 0x6e, 0x3f,
	0xaa, 0xaa, 0x31, 0xf2, 0x63, 0x36, 0xe7, 0x43, 0x71, 0x55, 0xa3, 0x8c, 0x81, 0x1b, 0xb4, 0x51,
	0x10, 0xb3, 0x79, 0x98, 0x9e, 0xf5, 0x0c, 0x8f, 0xc7, 0xae, 0x9c, 0x95, 0x18, 0x08, 0xd6, 0xd9,
	0x8a, 0xb4, 0x9c, 0x01, 0xd7, 0xb5, 0xdd, 0xe1, 0x9e, 0xf7, 0xd1, 0x2c, 0xb7, 0xb8, 0xea, 0x41,
	0x94, 0x73, 0xff, 0x06, 0x7a, 0x01, 0x23, 0x67, 0x0d, 0x94, 0x41, 0x4d, 0x7a, 0x59, 0x4a, 0x09,
	0x7e, 0x99, 0xdb, 0xbc, 0x63, 0x7b, 0xa1, 0x3a, 0x9d, 0xfc, 0x64, 0x10, 0x64, 0xeb, 0x7f, 0x1e,
	0x9a, 0xc0, 0xe9, 0xad, 0x69, 0xea, 0xcd, 0xd1, 0xe4, 0xd0, 0x2d, 0xee, 0x82, 0xf2, 0x2d, 0x84,
	0xb6, 0xd9, 0x27, 0x7a, 0x84, 0xa9, 0x6d, 0x56, 0x4f, 0xe8, 0xfc, 0xe7, 0xe8, 0x82, 0x03, 0xba,
	0x8f, 0x00, 0xeb, 0x34, 0xbb, 0x21, 0x34, 0x8e, 0xf0, 0x12, 0x26, 0x45, 0x73, 0x77, 0xb7, 0xa9,
	0x4a, 0x8a, 0x58, 0xcc, 0xe7, 0xbe, 0x1a, 0xbb, 0xff, 0x65, 0x49, 0xc3, 0x66, 0xef, 0x60, 0xf3,
	0x0a, 0x82, 0xec, 0x66, 0xb7, 0x6f, 0xa0, 0x61, 0x7f, 0x3a, 0x29, 0xef, 0x39, 0x16, 0xbe, 0xaf,
	0x5b, 0xbc, 0x81, 0xf0, 0xe0, 0x38, 0xe4, 0x08, 0x3c, 0x4d, 0x82, 0xb9, 0xef, 0x96, 0x84, 0xb7,
	0xf8, 0x0a, 0xdc, 0xcd, 0xba, 0x80, 0x73, 0x73, 0xfc, 0x10, 0x04, 0x93, 0xaf, 0x40, 0x9a, 0x27,
	0xd7, 0x53, 0x78, 0x72, 0x0c, 0x5c, 0xa7, 0x24, 0xb8, 0x13, 0x94, 0x6a, 0xe1, 0xf7, 0x62, 0x2b,
	0x02, 0x27, 0xb6, 0x29, 0x89, 0xd1, 0xe2, 0x3d, 0x4c, 0x87, 0x07, 0x28, 0x43, 0x18, 0xd3, 0xb6,
	0x28, 0x90, 0x48, 0x9c, 0x48, 0x80, 0xd1, 0x6d, 0x5e, 0x7d, 0xc3, 0x52, 0xb0, 0xc5, 0x0c, 0x60,
	0x7f, 0xa3, 0xbb, 0x18, 0x9a, 0xfb, 0xaa, 0x40, 0x71, 0x22, 0x27, 0xe0, 0xeb, 0xbc, 0x46, 0xc1,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xec, 0xd9, 0xcc, 0x31, 0xfc, 0x03, 0x00, 0x00,
}
