package core

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const maxInt64 = 1<<63 - 1

//// error
var ErrServerClosed = errors.New("服务器已关闭")
var ErrAbortHandler = errors.New("net/http: abort Handler")

////

var ServerContextKey = &contextKey{"http-server"}
var LocalAddrContextKey = &contextKey{"local-addr"}

type contextKey struct {
	name string
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (l tcpKeepAliveListener) Accept() (net.Conn, error) {
	conn, e := l.AcceptTCP()
	if e != nil {
		return nil, e
	}
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Minute)

	return conn, nil
}

type Server struct {
	Addr    string
	Handler Handler

	// 消息头大小
	HeadBytes  int64
	inShutdown int32

	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	doneChan   chan struct{}
	activeConn map[*conn]struct{}

	onShutdown []func()
	onConnect  func(*Connection)
	onClose    func(string)
}

func (server *Server) initHeadByte() int64 {
	server.HeadBytes = 4
	return server.HeadBytes
}
func (server *Server) ListenAndServe() error {

	listener, e := net.Listen("tcp", server.Addr)

	if e != nil {
		return e
	}

	return server.Serve(tcpKeepAliveListener{listener.(*net.TCPListener)})
}

func (s *Server) Serve(l net.Listener) error {

	l = &onceCloseListener{Listener: l}
	defer l.Close()

	if !s.trackListener(&l, true) {
		return ErrServerClosed
	}
	defer s.trackListener(&l, false)

	var tempDelay time.Duration
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, ServerContextKey, s)
	for {
		rw, e := l.Accept()

		if e != nil {
			Logf("Accept error: %v", e)

			select {
			case <-s.doneChan:
				return ErrServerClosed
			default:
			}

			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := time.Second; tempDelay > max {
					tempDelay = max
				}

				Logf("Accept error: %v; sleep:%v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}

			return e
		}

		tempDelay = 0

		c := s.newConn(rw)
		c.setState(c.rwc, StateNew)

		if s.onConnect != nil {
			cc := &Connection{
				conn: c,
			}
			cc.cw.c = cc
			cc.w = newBufioWriterSize(&cc.cw, 2048)
			s.onConnect(cc)
		}
		go c.serve(ctx)
	}

}
func (s *Server) shuttingDown() bool {
	return atomic.LoadInt32(&s.inShutdown) != 0
}
func (s *Server) trackListener(ln *net.Listener, add bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listeners == nil {
		s.listeners = make(map[*net.Listener]struct{})

	}
	if add {
		if s.shuttingDown() {
			return false
		}
		s.listeners[ln] = struct{}{}
	} else {
		delete(s.listeners, ln)
	}

	return true
}

func (s *Server) newConn(c net.Conn) *conn {

	con := &conn{
		server:     s,
		rwc:        c,
		remoteAddr: c.RemoteAddr().String(),
	}
	return con
}

func (s *Server) trackConn(c *conn, add bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.activeConn == nil {
		s.activeConn = make(map[*conn]struct{})
	}
	if add {
		s.activeConn[c] = struct{}{}
	} else {
		delete(s.activeConn, c)
	}

}
func (s *Server) Shutdown(ctx context.Context) error {
	atomic.StoreInt32(&s.inShutdown, 1)
	s.mu.Lock()
	e := s.closeListenersLocked()
	s.closeDoneChanLocked()
	for _, f := range s.onShutdown {
		go f()
	}
	s.mu.Unlock()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		if s.closeIdleConns() {
			return e
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}

func (s *Server) closeListenersLocked() error {
	var err error
	for ln := range s.listeners {
		if cer := (*ln).Close(); cer != nil && err == nil {
			err = cer
		}
		delete(s.listeners, ln)
	}

	return err
}

func (s *Server) getDoneChan() <-chan struct{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.getDoneChanLocked()
}

func (server *Server) getDoneChanLocked() chan struct{} {
	if server.doneChan == nil {
		server.doneChan = make(chan struct{})
	}
	return server.doneChan
}

func (s *Server) closeDoneChanLocked() {
	ch := s.getDoneChanLocked()
	select {
	case <-ch:
	default:
		close(ch)
	}
}

func (s *Server) closeIdleConns() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	quiescent := true

	for c := range s.activeConn {
		state, unixSec := c.getState()

		if state == StateNew && unixSec < time.Now().Unix()-5 {
			state = StateIdle
		}
		if state != StateIdle || unixSec == 0 {
			quiescent = false
			continue
		}
		c.rwc.Close()
		delete(s.activeConn, c)
	}

	return quiescent

}

func (s *Server) RegisterOnShutdown(f func()) {
	s.mu.Lock()
	s.onShutdown = append(s.onShutdown, f)
	s.mu.Unlock()
}
func (s *Server) RegisterOnConnection(f func(connection *Connection)) {
	s.mu.Lock()
	s.onConnect = f
	s.mu.Unlock()
}
func (s *Server) RegisterOnClose(f func(connection string)) {
	s.mu.Lock()
	s.onClose = f
	s.mu.Unlock()
}

type onceCloseListener struct {
	net.Listener
	once     sync.Once
	closeErr error
}
type ConnState int

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateClosed
)

// 一个tcp连接
type conn struct {
	server *Server  // 服务
	rwc    net.Conn //一个真正网络链接

	curRes     atomic.Value
	r          *connReader
	remoteAddr string
	cancelCtx  context.CancelFunc
	bufr       *bufio.Reader
	bufw       *bufio.Writer
	werr       error

	curState struct{ atomic uint64 }
}

type connReader struct {
	conn *conn

	mu      sync.Mutex
	cond    *sync.Cond
	inRead  bool
	remain  int64
	aborted bool
	hasByte bool
	byteBuf [1]byte
}

func (cr *connReader) lock() {
	cr.mu.Lock()
	if cr.cond == nil {
		cr.cond = sync.NewCond(&cr.mu)
	}
}

func (cr *connReader) unlock() {
	cr.mu.Unlock()
}

func (cr *connReader) Read(p []byte) (n int, err error) {
	cr.lock()
	if cr.inRead {
		panic("concurrent Read call")
	}

	if cr.hitReadLimit() {
		cr.unlock()
		return 0, io.EOF
	}

	if len(p) == 0 {
		return 0, nil
	}

	if int64(len(p)) > cr.remain {
		p = p[:cr.remain]
	}

	if cr.hasByte {
		p[0] = cr.byteBuf[0]
		cr.hasByte = false
		cr.unlock()
		return 1, nil
	}

	cr.inRead = true
	cr.unlock()
	n, err = cr.conn.rwc.Read(p)

	cr.lock()
	cr.inRead = false
	if err != nil {
		cr.handleReadErr(err)
	}
	cr.remain -= int64(n)
	cr.unlock()
	cr.cond.Broadcast()

	return
}

func (cr *connReader) hitReadLimit() bool {
	return cr.remain <= 0
}

func (cr *connReader) handleReadErr(e error) {

	cr.conn.cancelCtx()

}

func (cr *connReader) setInfiniteReadLimit() {
	cr.remain = maxInt64
}

type checkConnErrorWriter struct {
	c *conn
}

func (w checkConnErrorWriter) Write(p []byte) (n int, err error) {

	n, err = w.c.rwc.Write(p)

	if err != nil && w.c.werr == nil {
		w.c.werr = err
		w.c.cancelCtx()
	}
	return
}

var (
	bufioReaderPool   sync.Pool
	bufioWriter2kPool sync.Pool
	bufioWriter4kPool sync.Pool
)

func bufioWriterPool(size int) *sync.Pool {
	switch size {
	case 2 << 10:
		return &bufioWriter2kPool
	case 4 << 10:
		return &bufioWriter4kPool
	}
	return nil
}

///// read

func newBufioReader(r io.Reader) *bufio.Reader {
	if v := bufioReaderPool.Get(); v != nil {
		br := v.(*bufio.Reader)
		br.Reset(r)
		return br
	}
	return bufio.NewReader(r)
}

func putBufioReader(br *bufio.Reader) {
	br.Reset(nil)
	bufioReaderPool.Put(br)
}

///// write

func newBufioWriterSize(w io.Writer, size int) *bufio.Writer {
	pool := bufioWriterPool(size)
	if pool != nil {
		if v := pool.Get(); v != nil {
			bw := v.(*bufio.Writer)
			bw.Reset(w)
			return bw
		}
	}
	return bufio.NewWriterSize(w, size)
}
func putBufioWriter(bw *bufio.Writer) {
	bw.Reset(nil)
	if pool := bufioWriterPool(bw.Available()); pool != nil {
		pool.Put(bw)
	}
}

///// conn
func (c *conn) serve(ctx context.Context) {
	ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())

	defer func() {

		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			Debug("发生错误 %v####%s\n", err, buf)
		}

		if c.server.onClose != nil {

			c.server.onClose(c.remoteAddr)
		}
		c.close()
		c.setState(c.rwc, StateClosed)
	}()

	ctx, cancel := context.WithCancel(ctx)
	c.cancelCtx = cancel
	defer cancel()

	c.r = &connReader{conn: c}
	c.bufr = newBufioReader(c.r)
	c.bufw = newBufioWriterSize(checkConnErrorWriter{c}, 4<<10)

	for {
		Debug("开始读取请求")
		response, e := c.readRequest(ctx)
		if c.r.remain != c.server.initHeadByte() {
			c.setState(c.rwc, StateActive)
		}
		if e != nil {
			Logf("err 结束读取:%v", e)
			return
		}

		c.curRes.Store(response)

		// 处理业务
		serverHandler{s: c.server}.ServeHandler(response.req, response)

		response.finishRequest()

		c.setState(c.rwc, StateIdle)
		c.curRes.Store((*Response)(nil)) // 直接nil不可以？

		// 读不超时
		c.rwc.SetReadDeadline(time.Time{})

	}

}
func requestBodyRemains(rc io.ReadCloser) bool {

	switch v := rc.(type) {
	case *body:
		return v.bodyRemains()
	}
	return false
}
func registerOnHitEOF(rc io.ReadCloser, fn func()) {
	switch v := rc.(type) {
	case *body:
		v.registerOnHitEOF(fn)
	}
}

func (c *conn) readRequest(ctx context.Context) (*Response, error) {
	c.r.remain = c.server.initHeadByte()
	req, err := readRequest(c.bufr)
	if err != nil {
		return nil, err
	}
	c.r.setInfiniteReadLimit()
	ctx, cancelCtx := context.WithCancel(ctx)
	req.ctx = ctx
	req.RemoteAddr = c.remoteAddr
	response := &Response{
		conn:      c,
		req:       req,
		reqBody:   req.Body,
		cancelCtx: cancelCtx,
	}

	response.cw.res = response
	response.w = newBufioWriterSize(&response.cw, 2048)

	return response, err
}

func (c *conn) close() {
	c.finalFlush()
	Debug("Close conn")
	c.rwc.Close()
}

func (c *conn) finalFlush() {

	if c.bufr != nil {
		putBufioReader(c.bufr)
		c.bufr = nil
	}
	if c.bufw != nil {
		putBufioWriter(c.bufw)
		Debug("close conn bufw")
		c.bufw = nil
	}

}

func (c *conn) setState(rwc net.Conn, state ConnState) {
	srv := c.server
	switch state {
	case StateNew:
		srv.trackConn(c, true)
	case StateClosed:
		srv.trackConn(c, false)
	}
	packedState := uint64(time.Now().Unix()<<8) | uint64(state)
	atomic.StoreUint64(&c.curState.atomic, packedState)

}

func (c *conn) getState() (state ConnState, unixSec int64) {
	packedState := atomic.LoadUint64(&c.curState.atomic)
	return ConnState(packedState & 0xff), int64(packedState >> 8)
}

/////// handler
type Handler interface {
	ServeHandler(request *Request, response *Response)
}

type ServeMux struct {
	mu sync.RWMutex
	m  map[string]muxEntry
}

type muxEntry struct {
	h       Handler
	pattern string
}

// 注册匿名函数的时候使用
type HandlerFunc func(*Request, *Response)

func (f HandlerFunc) ServeHandler(req *Request, res *Response) {
	f(req, res)
}

func (mu *ServeMux) Handler(request *Request) (h Handler) {

	h = mu.match(request.MesType())
	if h == nil {
		Logf("没有handler:%s", request.MesType())
		return
	}

	return
}

func (mu *ServeMux) match(path string) (h Handler) {
	v, ok := mu.m[path]
	if ok {
		return v.h
	}

	return nil
}

func (mu *ServeMux) ServeHandler(req *Request, res *Response) {

	h := mu.Handler(req)

	h.ServeHandler(req, res)
}

// 注册处理过程
func (mu *ServeMux) Handle(patten string, handler Handler) {
	mu.mu.Lock()
	defer mu.mu.Unlock()

	if mu.m == nil {
		mu.m = make(map[string]muxEntry)
	}

	h := muxEntry{handler, patten}
	mu.m[patten] = h
}

// 注册处理方法
func (mu *ServeMux) HandleFunc(patten string, handler func(*Request, *Response)) {
	mu.Handle(patten, HandlerFunc(handler))
}

type serverHandler struct {
	s *Server
}

func (s serverHandler) ServeHandler(request *Request, response *Response) {
	s.s.Handler.ServeHandler(request, response)
}
