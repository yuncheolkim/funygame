package utils

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
)

func MsgToBytes(msg proto.Message) []byte {
	var b []byte
	l, _ := proto.Marshal(msg)

	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, int32(len(l)))
	b = append(b, bytesBuffer.Bytes()...)
	b = append(b, l...)
	return b
}