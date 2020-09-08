package core

import (
	"bufio"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"sync"
)

var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")

type Request struct {
	Length     int64
	Body       io.ReadCloser
	ctx        context.Context
	RemoteAddr string
}

// 消息类型
func (r *Request) MesType() string {
	return "pf"
}
func (r *Request) ReadJson(obj interface{}) ([]byte, error) {

	p, _ := ioutil.ReadAll(r.Body)

	Debug("ReadJson %b", p)
	return p, json.Unmarshal(p, obj)
}

func (r *Request) ReadPb(obj proto.Message) error {
	p, e := ioutil.ReadAll(r.Body)
	if e != nil {
		println(e)
	}

	return proto.Unmarshal(p, obj)
}

func readRequest(b *bufio.Reader) (request *Request, err error) {

	request = new(Request)

	var head []byte
	var total = 4
	for {
		r := make([]byte, total)
		n, err := b.Read(r)
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
			return nil, err
		}
	}
	Debug("Head:%b,%s", head, head)
	request.Length = int64(binary.BigEndian.Uint32(head))
	Logf("长度:%v", request.Length)

	requestBody(request, b)

	return
}

func requestBody(request *Request, r *bufio.Reader) {
	request.Body = &body{src: io.LimitReader(r, request.Length), r: r}
}

type body struct {
	src    io.Reader
	r      *bufio.Reader
	remain int

	mu       sync.Mutex
	sawEof   bool
	closed   bool
	onHitEOF func()
}

func (b *body) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return nil
	}
	var err error

	b.closed = true

	return err
}

func (b *body) Read(p []byte) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return 0, ErrBodyReadAfterClose
	}
	return b.readLocked(p)

}

func (b *body) readLocked(p []byte) (n int, err error) {
	if b.sawEof {
		return 0, io.EOF
	}
	Debug("读取长度:%v", len(p))
	n, err = b.src.Read(p)
	Debug("读取:%v,err:%v", n, err)
	if err == io.EOF {
		b.sawEof = true

		if lr, ok := b.src.(*io.LimitedReader); ok && lr.N > 0 {
			Debug("ErrUnexpectedEOF")
			err = io.ErrUnexpectedEOF
		}
	}

	if err == nil && n > 0 {
		if lr, ok := b.src.(*io.LimitedReader); ok && lr.N == 0 {
			Debug("结束body")
			err = io.EOF
			b.sawEof = true
		}
	}

	if b.sawEof && b.onHitEOF != nil {
		Debug("结束")
		b.onHitEOF()
	}
	return
}

func (b *body) bodyRemains() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return !b.sawEof
}

func (b *body) registerOnHitEOF(fn func()) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.onHitEOF = fn
}
