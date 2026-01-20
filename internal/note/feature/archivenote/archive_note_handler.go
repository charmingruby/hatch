package archivenote

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/http/rest"
	"HATCH_APP/pkg/o11y"
	"errors"

	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log := o11y.FromContext(ctx)

		log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

		id := c.Param("id")

		if err := svc.Execute(ctx, id); err != nil {
			if errors.Is(err, domain.ErrNoteNotFound) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: note not found",
					"error", err,
				)

				rest.SendNotFoundResponse(c, err.Error())
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: internal error",
				"error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/ArchiveNote: finished successfully")

		rest.SendEmptyResponse(c)
	}
}
