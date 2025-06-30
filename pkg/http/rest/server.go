// Package rest provides an HTTP server.
//
// It defines a Server struct with methods to start and gracefully stop the HTTP server.
package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server wraps an http.Server configured to serve RESTful APIs.
type Server struct {
	server http.Server
}

// New initializes a new Server instance and returns it along with a preconfigured *gin.Engine router.
//
// The server is configured with:
//   - Read timeout: 5 seconds
//   - Write timeout: 10 seconds
//   - Idle timeout: 120 seconds
func New(port string) (*Server, *gin.Engine) {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	addr := fmt.Sprintf(":%s", port)

	return &Server{
		server: http.Server{
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         addr,
			Handler:      router,
		},
	}, router
}

// Start runs the HTTP server and begins handling requests.
//
// It blocks the calling goroutine. If the port is already in use or the server fails to start, it returns an error.
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Stop gracefully shuts down the HTTP server with the provided context.
//
// This allows in-flight requests to complete before terminating.
// Returns an error if the shutdown fails or times out.
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
