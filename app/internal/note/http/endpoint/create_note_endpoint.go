package endpoint

import (
	"HATCH_APP/internal/note/dto"
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"
	"errors"

	"github.com/gin-gonic/gin"
)

type CreateNoteRequest struct {
	Title   string `json:"title"   binding:"required" validate:"required,gt=0"`
	Content string `json:"content" binding:"required" validate:"required,gt=0"`
}

func (e *Endpoint) CreateNote(c *gin.Context) {
	ctx := c.Request.Context()

	e.log.InfoContext(ctx, "endpoint/CreateNote: request received")

	var req CreateNoteRequest
	if err := c.BindJSON(&req); err != nil {
		e.log.ErrorContext(
			ctx,
			"endpoint/CreateNote: unable to parse payload",
			"error", err.Error(),
		)

		rest.SendBadRequestResponse(c, err.Error())
		return
	}
	if err := e.val.Validate(req); err != nil {
		e.log.ErrorContext(
			ctx,
			"endpoint/CreateNote: invalid payload",
			"error", err.Error(),
		)

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
			e.log.ErrorContext(
				ctx,
				"endpoint/CreateNote: database error",
				"error", databaseErr.Unwrap().Error(),
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		e.log.ErrorContext(
			ctx,
			"endpoint/CreateNote: unknown error", "error", err.Error(),
		)

		rest.SendInternalServerErrorResponse(c)
		return
	}

	e.log.InfoContext(
		ctx,
		"endpoint/CreateNote: finished successfully",
	)

	rest.SendCreatedResponse(c, op.ID, "note")
}
