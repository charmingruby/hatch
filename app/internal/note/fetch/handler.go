package fetch

import (
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/transport/http"
	"HATCH_APP/pkg/logger"
	"errors"

	"github.com/gin-gonic/gin"
)

func registerRoute(
	log *logger.Logger,
	api *gin.RouterGroup,
	uc UseCase,
) {
	api.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/ListNotes: request received")

		op, err := uc.Execute(ctx)
		if err != nil {
			var databaseErr *customerr.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ListNotes: database error",
					"error", databaseErr.Unwrap().Error(),
				)

				http.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ListNotes: unknown error", "error", err.Error(),
			)

			http.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/ListNotes: finished successfully")

		http.SendOKResponse(
			c,
			"",
			op.Notes,
		)
	})
}
