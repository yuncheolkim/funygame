// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

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

type Message struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	MsgNo                int32    `protobuf:"varint,2,opt,name=msgNo,proto3" json:"msgNo,omitempty"`
	BroadcastUid         []int64  `protobuf:"varint,3,rep,packed,name=broadcastUid,proto3" json:"broadcastUid,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Uid                  int64    `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Message) GetMsgNo() int32 {
	if m != nil {
		return m.MsgNo
	}
	return 0
}

func (m *Message) GetBroadcastUid() []int64 {
	if m != nil {
		return m.BroadcastUid
	}
	return nil
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Message) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "proto.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd5, 0x5c,
	0xec, 0xbe, 0x10, 0x71, 0x21, 0x01, 0x2e, 0xe6, 0xe2, 0xd4, 0x42, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xe6, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b, 0x35, 0xb7, 0x38, 0xdd, 0x2f, 0x5f, 0x82, 0x49, 0x81,
	0x51, 0x83, 0x35, 0x08, 0xc2, 0x11, 0x52, 0xe2, 0xe2, 0x49, 0x2a, 0xca, 0x4f, 0x4c, 0x49, 0x4e,
	0x2c, 0x2e, 0x09, 0xcd, 0x4c, 0x91, 0x60, 0x56, 0x60, 0xd6, 0x60, 0x0e, 0x42, 0x11, 0x13, 0x12,
	0xe2, 0x62, 0x49, 0xca, 0x4f, 0xa9, 0x94, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x41,
	0xe6, 0x97, 0x66, 0xa6, 0x48, 0xb0, 0x42, 0xcc, 0x2f, 0xcd, 0x4c, 0x71, 0x62, 0x89, 0x62, 0x2a,
	0x48, 0x4a, 0x62, 0x03, 0xbb, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x24, 0x03, 0x4a, 0xed,
	0xa1, 0x00, 0x00, 0x00,
}
