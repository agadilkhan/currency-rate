package httpserver

import "time"

type Option func(s *Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.server.Addr = host
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
