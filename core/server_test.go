package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"funygame/pb"
	"io"
	"io/ioutil"

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

	go readData(conn)

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

func readData(conn *net.TCPConn) error {
	for {

		var head []byte
		var total = 4
		for {
			r := make([]byte, total)
			n, err := conn.Read(r)
			Debug("读取内容:%v,%b,%v,%v", n, r, err, total)
			total -= n
			if n > 0 {
				head = append(head, r[0:n]...)
			}
			if total <= 0 {
				break
			}

			if err != nil {
				Debug("读发生错误%v", err)
				return err
			}
		}
		Debug("Head:%b,%s", head, head)
		l := int64(binary.BigEndian.Uint32(head))
		Logf("长度:%v", l)

		r := io.LimitReader(conn, l)
		p, e := ioutil.ReadAll(r)
		if e != nil {
			println(e)
		}
		test := &pb.Message{}
		proto.Unmarshal(p, test)
		fmt.Println(test)
	}
}

func TestBit(t *testing.T) {
	b := []byte{0, 0, 0, 'J'}
	fmt.Println(int64(binary.BigEndian.Uint32(b)))
}
