package archivenote

import (
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/telemetry"
	"errors"

	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(log *telemetry.Logger, uc UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

		id := c.Param("id")

		if err := uc.Execute(ctx, id); err != nil {
			var notFoundErr *errs.NotFoundError
			if errors.As(err, &notFoundErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: not found error",
					"error", err,
				)

				rest.SendNotFoundResponse(c, err.Error())
				return
			}

			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: database error",
					"error", databaseErr.Unwrap().Error(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: unknown error", "error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/ArchiveNote: finished successfully")

		rest.SendEmptyResponse(c)
	}
}
