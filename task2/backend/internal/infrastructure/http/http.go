package http

import (
	"context"
	"net/http"
	"software-engineering-2/internal/config"
)

type Server struct {
	server *http.Server
}

func NewServer(config config.HTTPConfig, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Handler:      handler,
			Addr:         config.Addr,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
			IdleTimeout:  config.IdleTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
