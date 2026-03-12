package httpx

import (
	"HATCH_APP/pkg/validator"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
}

func NewServer(port string, v *validator.Validator, ext External) (*Server, *chi.Mux) {
	addr := ":" + port

	r := chi.NewRouter()

	r.Use(withO11y, withValidator(v))

	registerProbes(r, ext)

	return &Server{
		Server: http.Server{
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         addr,
			Handler:      r,
		},
	}, r
}

func (s *Server) Start() error {
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Close(ctx context.Context) error {
	return s.Shutdown(ctx)
}
