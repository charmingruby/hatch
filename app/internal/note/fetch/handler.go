package fetch

import (
	"HATCH_APP/internal/shared/customerr"
	"HATCH_APP/internal/shared/transport/http"
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

	notes.PUT("/:id", h.handle)
}

func (h *handler) handle(c *gin.Context) {
	ctx := c.Request.Context()

	h.log.InfoContext(ctx, "endpoint/ListNotes: request received")

	op, err := h.svc.Execute(ctx)
	if err != nil {
		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			h.log.ErrorContext(
				ctx,
				"endpoint/ListNotes: database error",
				"error", databaseErr.Unwrap().Error(),
			)

			http.SendInternalServerErrorResponse(c)
			return
		}

		h.log.ErrorContext(
			ctx,
			"endpoint/ListNotes: unknown error", "error", err.Error(),
		)

		http.SendInternalServerErrorResponse(c)
		return
	}

	h.log.InfoContext(
		ctx,
		"endpoint/ListNotes: finished successfully",
	)

	http.SendOKResponse(
		c,
		"",
		op.Notes,
	)
}
