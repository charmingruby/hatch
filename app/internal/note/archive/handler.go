package archive

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

	h.log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

	id := c.Param("id")

	if err := h.svc.Execute(ctx, Input{
		ID: id,
	}); err != nil {
		var notFoundErr *customerr.NotFoundError
		if errors.As(err, &notFoundErr) {
			h.log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: not found error",
				"error", err.Error(),
			)

			http.SendNotFoundResponse(c, err.Error())
			return
		}

		var databaseErr *customerr.DatabaseError
		if errors.As(err, &databaseErr) {
			h.log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: database error",
				"error", databaseErr.Unwrap().Error(),
			)

			http.SendInternalServerErrorResponse(c)
			return
		}

		h.log.ErrorContext(
			ctx,
			"endpoint/ArchiveNote: unknown error", "error", err.Error(),
		)

		http.SendInternalServerErrorResponse(c)
		return
	}

	h.log.InfoContext(
		ctx,
		"endpoint/ArchiveNote: finished successfully",
	)

	http.SendEmptyResponse(c)
}
