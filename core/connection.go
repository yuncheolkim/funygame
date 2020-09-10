package core

type Connection struct {
	conn *conn
}

func (r *Connection) RemoteAddr() string {
	return r.conn.remoteAddr
}
func (r *Connection) Write(data []byte) (n int, err error) {
	return r.conn.bufw.Write(data)
}

func (r *Connection) Flush() {
	if r.conn.bufw != nil {
		r.conn.bufw.Flush()
	}
}

func (r *Connection) WriteAndFlush(data []byte) (n int, err error) {
	n, err = r.Write(data)
	if err == nil {
		r.Flush()
	}
	return
}
func (r *Connection) Close() {
	r.conn.close()
}
