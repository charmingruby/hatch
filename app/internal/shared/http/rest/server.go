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
	"go.uber.org/fx"
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
	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	addr := fmt.Sprintf(":%s", cfg.RestServerPort)

	router.Use(val.Middleware())

	registerProbes(log, router, db)

	return &Server{
		Server: http.Server{
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         addr,
			Handler:      router,
		},
		port: cfg.RestServerPort,
	}, router
}

func (s *Server) Start() error {
	if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Close(ctx context.Context) error {
	return s.Shutdown(ctx)
}

var Module = fx.Module("rest",
	fx.Provide(NewServer),
	fx.Invoke(func(lc fx.Lifecycle, srv *Server, router *gin.Engine, shutdowner fx.Shutdowner) {
		errChan := make(chan error, 1)

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := srv.Start(); err != nil {
						errChan <- err
						_ = shutdowner.Shutdown()
					}
				}()

				select {
				case err := <-errChan:
					return err
				case <-time.After(100 * time.Millisecond):
					return nil
				}
			},
			OnStop: func(ctx context.Context) error {
				if err := srv.Close(ctx); err != nil {
					return err
				}

				return nil
			},
		})
	}),
)
