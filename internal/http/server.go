package http

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type HTTPPort string

type PageBuilder interface {
	BuildPage(ctx context.Context, videoURL string) (string, error)
}

type Server struct {
	port    string
	builder PageBuilder
	logger  *zap.Logger
}

func NewServer(p HTTPPort, builder PageBuilder) *Server {
	return &Server{
		port:    string(p),
		builder: builder,
	}
}

func (s *Server) Start() error {
	http.Handle("/{id}", s)
	return http.ListenAndServe(":"+s.port, nil)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	videoID := r.PathValue("id")
	if videoID == "" {
		s.logger.Error("video ID is required")
		http.Error(w, "video ID is required", http.StatusBadRequest)
		return
	}
	resp, err := s.builder.BuildPage(r.Context(), videoID)
	if err != nil {
		s.logger.Error("failed to build page", zap.Error(err))
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write([]byte(resp))
	if err != nil {
		s.logger.Error("failed to write response", zap.Error(err))
		http.Error(w, "failed to write response", http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}
