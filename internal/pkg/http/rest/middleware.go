package rest

import (
	"HATCH_APP/pkg/o11y/logging"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
)

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logging.Log.With(
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

		ctx := logging.WithLogger(c.Request.Context(), log)
		c.Request = c.Request.WithContext(ctx)

		log.Info("request started")
		defer log.Info("request finished")

		c.Next()
	}
}

func validationMiddleware(v *validator.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := validator.WithValidator(c.Request.Context(), v)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
