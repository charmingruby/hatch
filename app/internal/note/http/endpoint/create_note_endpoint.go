package endpoint

import (
	"context"
	"errors"

	"github.com/charmingruby/pack/internal/note/usecase"
	"github.com/charmingruby/pack/internal/shared/customerr"
	"github.com/charmingruby/pack/internal/shared/http/rest"
	"github.com/gin-gonic/gin"
)

type CreateNoteRequest struct {
	Title   string `json:"title"   binding:"required" validate:"required,gt=0"`
	Content string `json:"content" binding:"required" validate:"required,gt=0"`
}

func (e *Endpoint) CreateNote(c *gin.Context) {
	ctx := context.Background()

	var req CreateNoteRequest
	if err := c.BindJSON(&req); err != nil {
		rest.SendBadRequestResponse(c, err.Error())
		return
	}
	if err := e.val.Validate(req); err != nil {
		rest.SendBadRequestResponse(c, err.Error())
		return
	}

	op, err := e.service.CreateNote(ctx, usecase.CreateNoteInput{
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
