package rest

import (
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/validator"

	"github.com/gin-gonic/gin"
)

func o11yMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := o11y.Log.With(
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

		ctx := o11y.WithLogger(c.Request.Context(), log)
		c.Request = c.Request.WithContext(ctx)

		log.InfoContext(ctx, "request started")
		defer log.InfoContext(ctx, "request finished")

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
