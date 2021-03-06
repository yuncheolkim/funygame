// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fightgame.proto

package pb

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

// 开始匹配
type StartMatchReq_10001 struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartMatchReq_10001) Reset()         { *m = StartMatchReq_10001{} }
func (m *StartMatchReq_10001) String() string { return proto.CompactTextString(m) }
func (*StartMatchReq_10001) ProtoMessage()    {}
func (*StartMatchReq_10001) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{0}
}

func (m *StartMatchReq_10001) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartMatchReq_10001.Unmarshal(m, b)
}
func (m *StartMatchReq_10001) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartMatchReq_10001.Marshal(b, m, deterministic)
}
func (m *StartMatchReq_10001) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartMatchReq_10001.Merge(m, src)
}
func (m *StartMatchReq_10001) XXX_Size() int {
	return xxx_messageInfo_StartMatchReq_10001.Size(m)
}
func (m *StartMatchReq_10001) XXX_DiscardUnknown() {
	xxx_messageInfo_StartMatchReq_10001.DiscardUnknown(m)
}

var xxx_messageInfo_StartMatchReq_10001 proto.InternalMessageInfo

func (m *StartMatchReq_10001) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type StartMatchRes_10001 struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	PlayerIndex          []int32  `protobuf:"varint,2,rep,packed,name=playerIndex,proto3" json:"playerIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartMatchRes_10001) Reset()         { *m = StartMatchRes_10001{} }
func (m *StartMatchRes_10001) String() string { return proto.CompactTextString(m) }
func (*StartMatchRes_10001) ProtoMessage()    {}
func (*StartMatchRes_10001) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{1}
}

func (m *StartMatchRes_10001) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartMatchRes_10001.Unmarshal(m, b)
}
func (m *StartMatchRes_10001) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartMatchRes_10001.Marshal(b, m, deterministic)
}
func (m *StartMatchRes_10001) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartMatchRes_10001.Merge(m, src)
}
func (m *StartMatchRes_10001) XXX_Size() int {
	return xxx_messageInfo_StartMatchRes_10001.Size(m)
}
func (m *StartMatchRes_10001) XXX_DiscardUnknown() {
	xxx_messageInfo_StartMatchRes_10001.DiscardUnknown(m)
}

var xxx_messageInfo_StartMatchRes_10001 proto.InternalMessageInfo

func (m *StartMatchRes_10001) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *StartMatchRes_10001) GetPlayerIndex() []int32 {
	if m != nil {
		return m.PlayerIndex
	}
	return nil
}

//离开游戏
type LeaveRoomReq_10005 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomReq_10005) Reset()         { *m = LeaveRoomReq_10005{} }
func (m *LeaveRoomReq_10005) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomReq_10005) ProtoMessage()    {}
func (*LeaveRoomReq_10005) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{2}
}

func (m *LeaveRoomReq_10005) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomReq_10005.Unmarshal(m, b)
}
func (m *LeaveRoomReq_10005) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomReq_10005.Marshal(b, m, deterministic)
}
func (m *LeaveRoomReq_10005) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomReq_10005.Merge(m, src)
}
func (m *LeaveRoomReq_10005) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomReq_10005.Size(m)
}
func (m *LeaveRoomReq_10005) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomReq_10005.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomReq_10005 proto.InternalMessageInfo

// 攻击别人
type AttackTell_20001 struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttackTell_20001) Reset()         { *m = AttackTell_20001{} }
func (m *AttackTell_20001) String() string { return proto.CompactTextString(m) }
func (*AttackTell_20001) ProtoMessage()    {}
func (*AttackTell_20001) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{3}
}

func (m *AttackTell_20001) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttackTell_20001.Unmarshal(m, b)
}
func (m *AttackTell_20001) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttackTell_20001.Marshal(b, m, deterministic)
}
func (m *AttackTell_20001) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttackTell_20001.Merge(m, src)
}
func (m *AttackTell_20001) XXX_Size() int {
	return xxx_messageInfo_AttackTell_20001.Size(m)
}
func (m *AttackTell_20001) XXX_DiscardUnknown() {
	xxx_messageInfo_AttackTell_20001.DiscardUnknown(m)
}

var xxx_messageInfo_AttackTell_20001 proto.InternalMessageInfo

func (m *AttackTell_20001) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AttackTell_20001) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

// 回血
type CureTell_20002 struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CureTell_20002) Reset()         { *m = CureTell_20002{} }
func (m *CureTell_20002) String() string { return proto.CompactTextString(m) }
func (*CureTell_20002) ProtoMessage()    {}
func (*CureTell_20002) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{4}
}

func (m *CureTell_20002) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CureTell_20002.Unmarshal(m, b)
}
func (m *CureTell_20002) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CureTell_20002.Marshal(b, m, deterministic)
}
func (m *CureTell_20002) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CureTell_20002.Merge(m, src)
}
func (m *CureTell_20002) XXX_Size() int {
	return xxx_messageInfo_CureTell_20002.Size(m)
}
func (m *CureTell_20002) XXX_DiscardUnknown() {
	xxx_messageInfo_CureTell_20002.DiscardUnknown(m)
}

var xxx_messageInfo_CureTell_20002 proto.InternalMessageInfo

func (m *CureTell_20002) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

// 增加护盾
type DefTell_20003 struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefTell_20003) Reset()         { *m = DefTell_20003{} }
func (m *DefTell_20003) String() string { return proto.CompactTextString(m) }
func (*DefTell_20003) ProtoMessage()    {}
func (*DefTell_20003) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{5}
}

func (m *DefTell_20003) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefTell_20003.Unmarshal(m, b)
}
func (m *DefTell_20003) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefTell_20003.Marshal(b, m, deterministic)
}
func (m *DefTell_20003) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefTell_20003.Merge(m, src)
}
func (m *DefTell_20003) XXX_Size() int {
	return xxx_messageInfo_DefTell_20003.Size(m)
}
func (m *DefTell_20003) XXX_DiscardUnknown() {
	xxx_messageInfo_DefTell_20003.DiscardUnknown(m)
}

var xxx_messageInfo_DefTell_20003 proto.InternalMessageInfo

func (m *DefTell_20003) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

// 游戏正式开始
type StartGamePush_30001 struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartGamePush_30001) Reset()         { *m = StartGamePush_30001{} }
func (m *StartGamePush_30001) String() string { return proto.CompactTextString(m) }
func (*StartGamePush_30001) ProtoMessage()    {}
func (*StartGamePush_30001) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{6}
}

func (m *StartGamePush_30001) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGamePush_30001.Unmarshal(m, b)
}
func (m *StartGamePush_30001) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGamePush_30001.Marshal(b, m, deterministic)
}
func (m *StartGamePush_30001) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGamePush_30001.Merge(m, src)
}
func (m *StartGamePush_30001) XXX_Size() int {
	return xxx_messageInfo_StartGamePush_30001.Size(m)
}
func (m *StartGamePush_30001) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGamePush_30001.DiscardUnknown(m)
}

var xxx_messageInfo_StartGamePush_30001 proto.InternalMessageInfo

func (m *StartGamePush_30001) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

// 玩家进入
type UserEnterPush_30002 struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserEnterPush_30002) Reset()         { *m = UserEnterPush_30002{} }
func (m *UserEnterPush_30002) String() string { return proto.CompactTextString(m) }
func (*UserEnterPush_30002) ProtoMessage()    {}
func (*UserEnterPush_30002) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{7}
}

func (m *UserEnterPush_30002) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEnterPush_30002.Unmarshal(m, b)
}
func (m *UserEnterPush_30002) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEnterPush_30002.Marshal(b, m, deterministic)
}
func (m *UserEnterPush_30002) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEnterPush_30002.Merge(m, src)
}
func (m *UserEnterPush_30002) XXX_Size() int {
	return xxx_messageInfo_UserEnterPush_30002.Size(m)
}
func (m *UserEnterPush_30002) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEnterPush_30002.DiscardUnknown(m)
}

var xxx_messageInfo_UserEnterPush_30002 proto.InternalMessageInfo

func (m *UserEnterPush_30002) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

// 血量变化
type BloodChangePush_30003 struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BloodChangePush_30003) Reset()         { *m = BloodChangePush_30003{} }
func (m *BloodChangePush_30003) String() string { return proto.CompactTextString(m) }
func (*BloodChangePush_30003) ProtoMessage()    {}
func (*BloodChangePush_30003) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{8}
}

func (m *BloodChangePush_30003) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BloodChangePush_30003.Unmarshal(m, b)
}
func (m *BloodChangePush_30003) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BloodChangePush_30003.Marshal(b, m, deterministic)
}
func (m *BloodChangePush_30003) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BloodChangePush_30003.Merge(m, src)
}
func (m *BloodChangePush_30003) XXX_Size() int {
	return xxx_messageInfo_BloodChangePush_30003.Size(m)
}
func (m *BloodChangePush_30003) XXX_DiscardUnknown() {
	xxx_messageInfo_BloodChangePush_30003.DiscardUnknown(m)
}

var xxx_messageInfo_BloodChangePush_30003 proto.InternalMessageInfo

func (m *BloodChangePush_30003) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BloodChangePush_30003) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

// 离开游戏
type LeaveRoomReq_30004 struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomReq_30004) Reset()         { *m = LeaveRoomReq_30004{} }
func (m *LeaveRoomReq_30004) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomReq_30004) ProtoMessage()    {}
func (*LeaveRoomReq_30004) Descriptor() ([]byte, []int) {
	return fileDescriptor_0455df22f6b044da, []int{9}
}

func (m *LeaveRoomReq_30004) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomReq_30004.Unmarshal(m, b)
}
func (m *LeaveRoomReq_30004) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomReq_30004.Marshal(b, m, deterministic)
}
func (m *LeaveRoomReq_30004) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomReq_30004.Merge(m, src)
}
func (m *LeaveRoomReq_30004) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomReq_30004.Size(m)
}
func (m *LeaveRoomReq_30004) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomReq_30004.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomReq_30004 proto.InternalMessageInfo

func (m *LeaveRoomReq_30004) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*StartMatchReq_10001)(nil), "proto.StartMatchReq_10001")
	proto.RegisterType((*StartMatchRes_10001)(nil), "proto.StartMatchRes_10001")
	proto.RegisterType((*LeaveRoomReq_10005)(nil), "proto.LeaveRoomReq_10005")
	proto.RegisterType((*AttackTell_20001)(nil), "proto.AttackTell_20001")
	proto.RegisterType((*CureTell_20002)(nil), "proto.CureTell_20002")
	proto.RegisterType((*DefTell_20003)(nil), "proto.DefTell_20003")
	proto.RegisterType((*StartGamePush_30001)(nil), "proto.StartGamePush_30001")
	proto.RegisterType((*UserEnterPush_30002)(nil), "proto.UserEnterPush_30002")
	proto.RegisterType((*BloodChangePush_30003)(nil), "proto.BloodChangePush_30003")
	proto.RegisterType((*LeaveRoomReq_30004)(nil), "proto.LeaveRoomReq_30004")
}

func init() { proto.RegisterFile("fightgame.proto", fileDescriptor_0455df22f6b044da) }

var fileDescriptor_0455df22f6b044da = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcb, 0x4c, 0xcf,
	0x28, 0x49, 0x4f, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0xba, 0x5c, 0xc2, 0xc1, 0x25, 0x89, 0x45, 0x25, 0xbe, 0x89, 0x25, 0xc9, 0x19, 0x41, 0xa9, 0x85,
	0xf1, 0x86, 0x06, 0x06, 0x06, 0x86, 0x42, 0x62, 0x5c, 0x6c, 0x45, 0xf9, 0xf9, 0xb9, 0x9e, 0x29,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x50, 0x9e, 0x92, 0x2f, 0xaa, 0xf2, 0x62, 0xa8, 0x72,
	0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0xb0, 0x6a, 0xd6, 0x20, 0x08, 0x47, 0x48, 0x81,
	0x8b, 0xbb, 0x20, 0x27, 0xb1, 0x32, 0xb5, 0xc8, 0x13, 0x2c, 0xc7, 0xa4, 0xc0, 0xac, 0xc1, 0x1a,
	0x84, 0x2c, 0xa4, 0x24, 0xc2, 0x25, 0xe4, 0x93, 0x9a, 0x58, 0x96, 0x1a, 0x94, 0x9f, 0x9f, 0x0b,
	0xb3, 0xdc, 0x54, 0xc9, 0x8a, 0x4b, 0xc0, 0xb1, 0xa4, 0x24, 0x31, 0x39, 0x3b, 0x24, 0x35, 0x27,
	0x27, 0xde, 0x08, 0x8f, 0x0d, 0x02, 0x5c, 0xcc, 0x79, 0xa5, 0xb9, 0x12, 0x4c, 0x60, 0x31, 0x10,
	0x53, 0x49, 0x89, 0x8b, 0xcf, 0xb9, 0xb4, 0x28, 0x15, 0xae, 0xd3, 0x08, 0xa6, 0x86, 0x11, 0xa1,
	0x46, 0x91, 0x8b, 0xd7, 0x25, 0x35, 0x0d, 0xae, 0xc4, 0x18, 0x8b, 0x12, 0x58, 0xb0, 0xb8, 0x27,
	0xe6, 0xa6, 0x06, 0x94, 0x16, 0x67, 0xc4, 0x1b, 0xe3, 0x0d, 0x16, 0x6d, 0x2e, 0xe1, 0xd0, 0xe2,
	0xd4, 0x22, 0xd7, 0xbc, 0x92, 0xd4, 0x22, 0xb8, 0x72, 0x23, 0xec, 0x8e, 0x56, 0xb2, 0xe7, 0x12,
	0x75, 0xca, 0xc9, 0xcf, 0x4f, 0x71, 0xce, 0x48, 0xcc, 0x4b, 0x47, 0x98, 0x6e, 0x4c, 0xb4, 0x1f,
	0xb5, 0xd0, 0x42, 0x0d, 0xa4, 0xdb, 0x04, 0xbb, 0x6e, 0x27, 0x96, 0x28, 0xa6, 0x82, 0xa4, 0x24,
	0x36, 0x70, 0x64, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0x75, 0xd8, 0x5f, 0x06, 0x02,
	0x00, 0x00,
}
