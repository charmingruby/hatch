package archive

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
	api.PATCH(":id", func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

		id := c.Param("id")

		if err := uc.Execute(ctx, Input{
			ID: id,
		}); err != nil {
			var notFoundErr *customerr.NotFoundError
			if errors.As(err, &notFoundErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: not found error",
					"error", err.Error(),
				)

				http.SendNotFoundResponse(c, err.Error())
				return
			}

			var databaseErr *customerr.DatabaseError
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
	})
}
