package gost

import (
	"gostonc/internal/model"
	"gostonc/internal/service"
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/go-log/log"
)

// Accepter represents a network endpoint that can accept connection from peer.
type Accepter interface {
	Accept() (net.Conn, error)
}

// Server is a proxy server.
type Server struct {
	Listener Listener
	Handler  Handler
	options  *ServerOptions
}

// Init intializes server with given options.
func (s *Server) Init(opts ...ServerOption) {
	if s.options == nil {
		s.options = &ServerOptions{}
	}
	for _, opt := range opts {
		opt(s.options)
	}
}

// Addr returns the address of the server
func (s *Server) Addr() net.Addr {
	return s.Listener.Addr()
}

// Close closes the server
func (s *Server) Close() error {
	return s.Listener.Close()
}

// Serve serves as a proxy server.
func (s *Server) Serve(h Handler, opts ...ServerOption) error {
	s.Init(opts...)

	if s.Listener == nil {
		ln, err := TCPListener("")
		if err != nil {
			return err
		}
		s.Listener = ln
	}

	if h == nil {
		h = s.Handler
	}
	if h == nil {
		h = HTTPHandler()
	}

	l := s.Listener
	var tempDelay time.Duration
	for {
		conn, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Logf("server: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		go h.Handle(conn)
	}
}

// Run starts to serve.
func (s *Server) Run() error {
	return s.Serve(s.Handler)
}

// ServerOptions holds the options for Server.
type ServerOptions struct {
}

// ServerOption allows a common way to set server options.
type ServerOption func(opts *ServerOptions)

// Listener is a proxy server listener, just like a net.Listener.
type Listener interface {
	net.Listener
}

func transport(user *model.User, rw1, rw2 io.ReadWriter) error {
	errc := make(chan error, 1)
	go func() {
		if user != nil {
			errc <- copyBuffer(user.Username, rw1, rw2)
		} else {
			errc <- copyBuffer2(rw1, rw2)
		}
	}()

	go func() {
		if user != nil {
			errc <- copyBuffer(user.Username, rw2, rw1)
		} else {
			errc <- copyBuffer2(rw2, rw1)
		}
	}()

	if err := <-errc; err != nil && err != io.EOF {
		return err
	}

	return nil
}

// 统计流量
func copyBuffer(uname string, dst io.Writer, src io.Reader) error {
	buf := lPool.Get().([]byte)
	defer lPool.Put(buf)

	n, err := io.CopyBuffer(dst, src, buf)

	if u, ok := service.GetCacheUserByUname(uname); ok {
		atomic.AddInt64(&u.TimeSpan.SpendFlow, n)
	}

	return err
}

func copyBuffer2(dst io.Writer, src io.Reader) error {
	buf := lPool.Get().([]byte)
	defer lPool.Put(buf)

	_, err := io.CopyBuffer(dst, src, buf)
	return err
}
