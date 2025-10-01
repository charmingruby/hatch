package endpoint

import (
	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"
	"errors"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) ArchiveNote(c *gin.Context) {
	ctx := c.Request.Context()

	e.log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

	id := c.Param("id")

	if err := e.service.ArchiveNote(ctx, dto.ArchiveNoteInput{
		ID: id,
	}); err != nil {
		var notFoundErr *customerr.NotFoundError
		if errors.As(err, &notFoundErr) {
			e.log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: not found error",
				"error", err.Error(),
			)

			rest.SendNotFoundResponse(c, err.Error())
			return
		}

		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			e.log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: database error",
				"error", databaseErr.Unwrap().Error(),
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		e.log.ErrorContext(
			ctx,
			"endpoint/ArchiveNote: unknown error", "error", err.Error(),
		)

		rest.SendInternalServerErrorResponse(c)
		return
	}

	e.log.InfoContext(
		ctx,
		"endpoint/ArchiveNote: finished successfully",
	)

	rest.SendEmptyResponse(c)
}
