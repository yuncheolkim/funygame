package core

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

type Response struct {
	conn        *conn
	req         *Request
	cancelCtx   context.CancelFunc
	cw          responseWriter
	w           *bufio.Writer
	handlerDone atomicBool
}


func (r *Response) WritePb(data proto.Message) (n int, err error) {

	bytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	var b []byte
	b = append(append(b, IntToByte(int32(len(bytes)))...), bytes...)
	return r.w.Write(b)
}

func (r *Response) Write(data []byte) (n int, err error) {
	return r.w.Write(data)
}

func (r *Response) Flush() {
	r.w.Flush()
}
func (r *Response) finishRequest() {
	r.handlerDone.setTrue()

	r.Flush()
	putBufioWriter(r.w)
	r.conn.bufw.Flush()
}

//// responseWriter
type responseWriter struct {
	res *Response
}

func (rw *responseWriter) Write(p []byte) (n int, err error) {

	n, err = rw.res.conn.bufw.Write(p)
	if err != nil {
		rw.res.conn.close()
	}
	return
}



