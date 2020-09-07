package core

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"testing"
	"time"
)

type TestDataJson struct {
	Age  int
	Name string
}

func TestServer_Serve(t *testing.T) {

	sm := &ServeMux{}

	sm.HandleFunc("test", func(r *Request, w *Response) {

		test := TestDataJson{}

		r.ReadJson(&test)
		fmt.Println(test)
		w.WriteJson(test)

	})

	s := &Server{
		Addr: "127.0.0.1:8900",
	}

	s.Handler = sm

	s.ListenAndServe()

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
	msg := TestDataJson{
		Age:  123,
		Name: "sst",
	}

	for i := 0; i < 2; i++ {
		var b []byte
		l, _ := json.Marshal(msg)

		x := int32(len(l))
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, x)
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
