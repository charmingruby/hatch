package endpoint

import (
	"errors"

	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) ArchiveNote(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	if err := e.service.ArchiveNote(ctx, dto.ArchiveNoteInput{
		ID: id,
	}); err != nil {
		var notFoundErr *customerr.NotFoundError
		if errors.As(err, &notFoundErr) {
			e.log.Error("not found error", "error", err.Error(), "request", c.Request)

			rest.SendNotFoundResponse(c, err.Error())
			return
		}

		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			e.log.Error("database error", "error", databaseErr.Unwrap().Error(), "request", c.Request)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		e.log.Error("unknown error", "error", err.Error(), "request", c.Request)

		rest.SendInternalServerErrorResponse(c)
		return
	}

	rest.SendEmptyResponse(c)
}
