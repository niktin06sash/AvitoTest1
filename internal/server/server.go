package server

import (
	"AvitoTest1/config"
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.ServerConfig, handler http.Handler) *Server {
	server := &Server{}
	server.httpServer = &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	return server
}
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
