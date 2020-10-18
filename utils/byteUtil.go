package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func MsgToBytes(msg proto.Message) []byte {
	var b []byte
	l, _ := proto.Marshal(msg)
	fmt.Printf("发送消息，长度:%v\n",len(l))
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, int32(len(l)))
	b = append(b, bytesBuffer.Bytes()...)
	b = append(b, l...)
	return b
}