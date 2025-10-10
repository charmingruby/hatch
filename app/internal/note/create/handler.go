package create

import (
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/logger"
	"HATCH_APP/pkg/validator"
	"errors"

	"github.com/gin-gonic/gin"
)

type handler struct {
	log *logger.Logger
	r   *gin.Engine
	val *validator.Validator
	svc Service
}

func registerRoute(h handler) {
	api := h.r.Group("/api")
	v1 := api.Group("/v1")
	notes := v1.Group("/notes")

	notes.POST("", h.handle)
}

func (e *handler) handle(c *gin.Context) {
	ctx := c.Request.Context()

	e.log.InfoContext(ctx, "endpoint/CreateNote: request received")

	var req Input
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

	op, err := e.svc.Execute(ctx, Input{
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
