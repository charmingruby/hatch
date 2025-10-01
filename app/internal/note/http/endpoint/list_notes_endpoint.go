package endpoint

import (
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"
	"errors"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) ListNotes(c *gin.Context) {
	ctx := c.Request.Context()

	e.log.InfoContext(ctx, "endpoint/ListNotes: request received")

	op, err := e.service.ListNotes(ctx)
	if err != nil {
		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			e.log.ErrorContext(
				ctx,
				"endpoint/ListNotes: database error",
				"error", databaseErr.Unwrap().Error(),
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		e.log.ErrorContext(
			ctx,
			"endpoint/ListNotes: unknown error", "error", err.Error(),
		)

		rest.SendInternalServerErrorResponse(c)
		return
	}

	e.log.InfoContext(
		ctx,
		"endpoint/ListNotes: finished successfully",
	)

	rest.SendOKResponse(
		c,
		"",
		op.Notes,
	)
}
