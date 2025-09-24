package endpoint

import (
	"context"
	"errors"

	"PACK_APP/internal/note/dto"
	"PACK_APP/internal/shared/customerr"
	"PACK_APP/internal/shared/http/rest"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) CreateNote(c *gin.Context) {
	ctx := context.Background()

	var req dto.CreateNoteInput
	if err := c.BindJSON(&req); err != nil {
		rest.SendBadRequestResponse(c, err.Error())
		return
	}
	if err := e.val.Validate(req); err != nil {
		rest.SendBadRequestResponse(c, err.Error())
		return
	}

	op, err := e.service.CreateNote(ctx, dto.CreateNoteInput{
		Title:   req.Title,
		Content: req.Content,
	})
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

	rest.SendCreatedResponse(c, op.ID, "note")
}
