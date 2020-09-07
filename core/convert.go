package core

import (
	"bytes"
	"encoding/binary"
)

func IntToByte(x int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)

	return bytesBuffer.Bytes()
}
