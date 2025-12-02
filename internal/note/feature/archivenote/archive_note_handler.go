package archivenote

import (
	"HATCH_APP/internal/pkg/errs"
	"HATCH_APP/internal/pkg/http/rest"
	"HATCH_APP/pkg/o11y/logging"
	"errors"

	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log := logging.FromContext(ctx)

		log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

		id := c.Param("id")

		if err := svc.Execute(ctx, id); err != nil {
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
