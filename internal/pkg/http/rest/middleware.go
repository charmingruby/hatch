package rest

import (
	"HATCH_APP/pkg/telemetry/logger"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
)

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Log.With(
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

		logger.WithLogger(c.Request.Context(), log)

		log.Info("request started")
		defer log.Info("request finished")

		c.Next()
	}
}

func validationMiddleware(v *validator.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		validator.WithValidator(c.Request.Context(), v)
		c.Next()
	}
}
