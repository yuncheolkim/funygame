package core

import (
	"bufio"
	"encoding/json"
)

type Connection struct {
	conn *conn
	w    *bufio.Writer
	cw   connectionWriter
}

func (r *Connection) RemoteAddr() string {
	return r.conn.remoteAddr
}
func (r *Connection) Write(data []byte) (n int, err error) {
	return r.w.Write(data)
}

func (r *Connection) Flush() {
	if r.conn.bufw != nil && r.w != nil {
		Debug("r.w : %v", r.cw)

		r.w.Flush()
		r.conn.bufw.Flush()
	}
	//putBufioWriter(r.w)
}

func (r *Connection) WriteAndFlush(data []byte) (n int, err error) {
	n, err = r.Write(data)
	if err == nil {
		r.Flush()
	}
	return
}

func (r *Connection) WriteJson(data interface{}) (n int, err error) {

	bytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	var b []byte
	b = append(append(b, IntToByte(int32(len(bytes)))...), bytes...)
	return r.w.Write(b)
}

type connectionWriter struct {
	c *Connection
}

////
func (rw *connectionWriter) Write(p []byte) (n int, err error) {

	n, err = rw.c.conn.rwc.Write(p)
	if err != nil {
		Debug("%v ", err)
		rw.c.conn.rwc.Close()
	}
	// todo
	return
}
func (rw *connectionWriter) flush() {
	rw.c.conn.bufw.Flush()
}
func (rw *connectionWriter) close() {

}
