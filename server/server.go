package server

import "net/http"

type Server struct {
	Server *http.Server
}

type ServerOptions struct {
	Port int
}

func (s *Server) Start(options *ServerOptions) error {
	return nil
}
