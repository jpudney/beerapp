package http

import (
	"net"
	"net/http"
)

const DefaultAddr = ":5000"

type Server struct {
	ln net.Listener

	Handler *Handler

	Addr string
}

func NewServer(addr string, h *Handler) *Server {
	if addr == "" {
		addr = DefaultAddr
	}

	return &Server{
		Addr:    addr,
		Handler: h,
	}
}

func (s *Server) Open() error {

	ln, err := net.Listen("tcp", s.Addr)

	if err != nil {
		return err
	}

	s.ln = ln

	return http.Serve(s.ln, s.Handler)

}

func (s *Server) Close() error {
	if s.ln != nil {
		return s.ln.Close()
	}

	return nil
}
