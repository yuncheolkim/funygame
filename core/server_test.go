package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"funygame/pb"

	"github.com/golang/protobuf/proto"
	"net"
	"testing"
	"time"
)

type TestDataJson struct {
	Age  int
	Name string
}

func TestClient(t *testing.T) {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8900")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}

	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + " : Client connected!")
	msg := pb.Message{
		Seq: 11,
		Uid: 123,
	}

	for i := 0; i < 2; i++ {
		var b []byte
		l, _ := proto.Marshal(&msg)

		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, int32(len(l)))
		b = append(b, bytesBuffer.Bytes()...)
		b = append(b, l...)

		fmt.Printf("send %b\n", b)
		conn.Write(b)
	}

	time.Sleep(time.Hour)
	conn.Close()
}

func TestBit(t *testing.T) {
	b := []byte{0, 0, 0, 'J'}
	fmt.Println(int64(binary.BigEndian.Uint32(b)))
}
