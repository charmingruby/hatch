package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"HATCH_APP/config"
	"HATCH_APP/pkg/db/postgres"
	"HATCH_APP/pkg/telemetry"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port string

	http.Server
}

func NewServer(
	log *telemetry.Logger,
	cfg *config.Config,
	val *validator.Validator,
	db *postgres.Client,
) (*Server, *gin.Engine) {
	addr := fmt.Sprintf(":%s", cfg.RestServerPort)

	r := setupRouter(val)

	registerProbes(log, r, db)

	return &Server{
		Server: http.Server{
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         addr,
			Handler:      r,
		},
		port: cfg.RestServerPort,
	}, r
}

func setupRouter(val *validator.Validator) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(val.Middleware())

	return r
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
