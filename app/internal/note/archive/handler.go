package archive

import (
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/internal/shared/http"
	"HATCH_APP/pkg/telemetry"
	"errors"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(log *telemetry.Logger, api *gin.RouterGroup, uc UseCase) {
	api.PATCH(":id", handle(log, uc))
}

func handle(log *telemetry.Logger, uc UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

		id := c.Param("id")

		if err := uc.Execute(ctx, UseCaseInput{
			ID: id,
		}); err != nil {
			var notFoundErr *errs.NotFoundError
			if errors.As(err, &notFoundErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: not found error",
					"error", err.Error(),
				)

				http.SendNotFoundResponse(c, err.Error())
				return
			}

			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: database error",
					"error", databaseErr.Unwrap().Error(),
				)

				http.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: unknown error", "error", err.Error(),
			)

			http.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/ArchiveNote: finished successfully")

		http.SendEmptyResponse(c)
	}
}
