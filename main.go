package main

import "net"

const defaultListenAddr = ":5000"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	ln net.Listener
}

func NewServer(cfg Config) *Server {

	/*
		 if the length of listen address is 0 or not available
		just assign the PORT number to default listen address PORT number
	*/

	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAddr
	}

	return &Server{
		Config: cfg,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)

	if err != nil {
		return err
	}

	s.ln = ln
	return nil
}
