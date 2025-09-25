package endpoint

import (
	"errors"

	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) ListNotes(c *gin.Context) {
	ctx := c.Request.Context()

	op, err := e.service.ListNotes(ctx)
	if err != nil {
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

	rest.SendOKResponse(
		c,
		"",
		op.Notes,
	)
}
