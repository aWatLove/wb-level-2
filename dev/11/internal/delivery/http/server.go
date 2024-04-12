package http

import (
	"context"
	"net/http"
	"time"
)

// Server - основная структура сервера
type Server struct {
	httpServer *http.Server
}

// Run - метод Server, который запускает сервер
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown - метод Server, который останавливает сервер
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
