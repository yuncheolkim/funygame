package core

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"io"
)

type Response struct {
	conn        *conn
	req         *Request
	reqBody     io.ReadCloser
	cancelCtx   context.CancelFunc
	cw          responseWriter
	w           *bufio.Writer
	handlerDone atomicBool
}

func (r *Response) WriteJson(data interface{}) (n int, err error) {

	bytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	var b []byte
	b = append(append(b, IntToByte(int32(len(bytes)))...), bytes...)
	return r.w.Write(b)
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
	r.cw.flush()
}
func (r *Response) finishRequest() {
	r.handlerDone.setTrue()

	r.w.Flush()
	putBufioWriter(r.w)
	r.cw.close()
	r.conn.bufw.Flush()

	r.conn.r.abortPendingRead()
	r.reqBody.Close()

}

//// responseWriter

type responseWriter struct {
	res *Response
}

func (rw *responseWriter) Write(p []byte) (n int, err error) {

	n, err = rw.res.conn.rwc.Write(p)
	if err != nil {
		rw.res.conn.rwc.Close()
	}
	// todo
	return
}
func (rw *responseWriter) flush() {
	rw.res.conn.bufw.Flush()
}
func (rw *responseWriter) close() {

}

//// conn Writer

type ConnWriter struct {
	conn      *conn
	cancelCtx context.CancelFunc
	cw        responseWriter
	w         *bufio.Writer
}

func (r *ConnWriter) Write(data []byte) (n int, err error) {
	return r.w.Write(data)
}

func (r *ConnWriter) Flush() {
	r.w.Flush()
	r.cw.flush()
	putBufioWriter(r.w)
}
