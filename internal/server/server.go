package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, r *chi.Mux) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
